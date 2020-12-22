// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

const keyBets = client.Key("bets")
const keyColor = client.Key("color")
const keyLastWinningColor = client.Key("last_winning_color")
const keyLockedBets = client.Key("locked_bets")
const keyPlayPeriod = client.Key("play_period")

const numColors = 5
const defaultPlayPeriod = 120

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("place_bet", placeBet)
	exports.AddCall("lock_bets", lockBets)
	exports.AddCall("pay_winners", payWinners)
	exports.AddCall("play_period", playPeriod)
	exports.AddCall("nothing", client.Nothing)
}

func placeBet(sc *client.ScCallContext) {
	amount := sc.Incoming().Balance(client.IOTA)
	if amount == 0 {
		sc.Log("Empty bet...")
		return
	}
	color := sc.Params().GetInt(keyColor).Value()
	if color == 0 {
		sc.Log("No color...")
		return
	}
	if color < 1 || color > numColors {
		sc.Log("Invalid color...")
		return
	}

	bet := &BetInfo{
		better: sc.Caller(),
		amount: amount,
		color:  color,
	}

	state := sc.State()
	bets := state.GetBytesArray(keyBets)
	betNr := bets.Length()
	bets.GetBytes(betNr).SetValue(encodeBetInfo(bet))
	if betNr == 0 {
		playPeriod := state.GetInt(keyPlayPeriod).Value()
		if playPeriod < 10 {
			playPeriod = defaultPlayPeriod
		}
		sc.Post("lock_bets").Post(playPeriod)
	}
}

func lockBets(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.From(sc.Contract().Id()) {
		sc.Log("Cancel spoofed request")
		return
	}

	// move all current bets to the locked_bets array
	state := sc.State()
	bets := state.GetBytesArray(keyBets)
	lockedBets := state.GetBytesArray(keyLockedBets)
	nrBets := bets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := bets.GetBytes(i).Value()
		lockedBets.GetBytes(i).SetValue(bytes)
	}
	bets.Clear()

	sc.Post("pay_winners").Post(0)
}

func payWinners(sc *client.ScCallContext) {
	// can only be sent by SC itself
	scId := sc.Contract().Id()
	if !sc.From(scId) {
		sc.Log("Cancel spoofed request")
		return
	}

	winningColor := sc.Utility().Random(5) + 1
	state := sc.State()
	state.GetInt(keyLastWinningColor).SetValue(winningColor)

	// gather all winners and calculate some totals
	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	lockedBets := state.GetBytesArray(keyLockedBets)
	winners := make([]*BetInfo, 0)
	nrBets := lockedBets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := lockedBets.GetBytes(i).Value()
		bet := decodeBetInfo(bytes)
		totalBetAmount += bet.amount
		if bet.color == winningColor {
			totalWinAmount += bet.amount
			winners = append(winners, bet)
		}
	}
	lockedBets.Clear()

	if len(winners) == 0 {
		sc.Log("Nobody wins!")
		// compact separate UTXOs into a single one
		sc.Transfer(scId, client.IOTA, totalBetAmount)
		return
	}

	// pay out the winners proportionally to their bet amount
	totalPayout := int64(0)
	size := len(winners)
	for i := 0; i < size; i++ {
		bet := winners[i]
		payout := totalBetAmount * bet.amount / totalWinAmount
		if payout != 0 {
			totalPayout += payout
			sc.Transfer(bet.better, client.IOTA, payout)
		}
		text := "Pay " + sc.Utility().String(payout) + " to " + bet.better.String()
		sc.Log(text)
	}

	// any truncation left-overs are fair picking for the smart contract
	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + sc.Utility().String(remainder)
		sc.Log(text)
		sc.Transfer(scId, client.IOTA, remainder)
	}
}

func playPeriod(sc *client.ScCallContext) {
	// can only be sent by SC owner
	if !sc.From(sc.Contract().Owner()) {
		sc.Log("Cancel spoofed request")
		return
	}

	playPeriod := sc.Params().GetInt(keyPlayPeriod).Value()
	if playPeriod < 10 {
		sc.Log("Invalid play period...")
		return
	}

	sc.State().GetInt(keyPlayPeriod).SetValue(playPeriod)
}
