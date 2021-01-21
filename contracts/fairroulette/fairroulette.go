// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

const KeyBets = client.Key("bets")
const KeyColor = client.Key("color")
const KeyLastWinningColor = client.Key("last_winning_color")
const KeyLockedBets = client.Key("locked_bets")
const KeyPlayPeriod = client.Key("play_period")

const NumColors = 5
const DefaultPlayPeriod = 120

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
		sc.Panic("Empty bet...")
	}
	color := sc.Params().GetInt(KeyColor).Value()
	if color == 0 {
		sc.Panic("No color...")
	}
	if color < 1 || color > NumColors {
		sc.Panic("Invalid color...")
	}

	bet := &BetInfo{
		Better: sc.Caller(),
		Amount: amount,
		Color:  color,
	}

	state := sc.State()
	bets := state.GetBytesArray(KeyBets)
	betNr := bets.Length()
	bets.GetBytes(betNr).SetValue(EncodeBetInfo(bet))
	if betNr == 0 {
		playPeriod := state.GetInt(KeyPlayPeriod).Value()
		if playPeriod < 10 {
			playPeriod = DefaultPlayPeriod
		}
		sc.Post(nil, client.Hname(0), client.NewHname("lock_bets"), nil, nil, playPeriod)
	}
}

func lockBets(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.From(sc.ContractId()) {
		sc.Panic("Cancel spoofed request")
	}

	// move all current bets to the locked_bets array
	state := sc.State()
	bets := state.GetBytesArray(KeyBets)
	lockedBets := state.GetBytesArray(KeyLockedBets)
	nrBets := bets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := bets.GetBytes(i).Value()
		lockedBets.GetBytes(i).SetValue(bytes)
	}
	bets.Clear()

	sc.Post(nil, client.Hname(0), client.NewHname("pay_winners"), nil, nil, 0)
}

func payWinners(sc *client.ScCallContext) {
	// can only be sent by SC itself
	scId := sc.ContractId()
	if !sc.From(scId) {
		sc.Panic("Cancel spoofed request")
	}

	winningColor := sc.Utility().Random(5) + 1
	state := sc.State()
	state.GetInt(KeyLastWinningColor).SetValue(winningColor)

	// gather all winners and calculate some totals
	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	lockedBets := state.GetBytesArray(KeyLockedBets)
	winners := make([]*BetInfo, 0)
	nrBets := lockedBets.Length()
	for i := int32(0); i < nrBets; i++ {
		bet := DecodeBetInfo(lockedBets.GetBytes(i).Value())
		totalBetAmount += bet.Amount
		if bet.Color == winningColor {
			totalWinAmount += bet.Amount
			winners = append(winners, bet)
		}
	}
	lockedBets.Clear()

	if len(winners) == 0 {
		sc.Log("Nobody wins!")
		// compact separate UTXOs into a single one
		sc.TransferToAddress(scId.Address(), client.IOTA, totalBetAmount)
		return
	}

	// pay out the winners proportionally to their bet amount
	totalPayout := int64(0)
	size := len(winners)
	for i := 0; i < size; i++ {
		bet := winners[i]
		payout := totalBetAmount * bet.Amount / totalWinAmount
		if payout != 0 {
			totalPayout += payout
			sc.TransferToAddress(bet.Better.Address(), client.IOTA, payout)
		}
		text := "Pay " + sc.Utility().String(payout) + " to " + bet.Better.String()
		sc.Log(text)
	}

	// any truncation left-overs are fair picking for the smart contract
	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + sc.Utility().String(remainder)
		sc.Log(text)
		sc.TransferToAddress(scId.Address(), client.IOTA, remainder)
	}
}

func playPeriod(sc *client.ScCallContext) {
	// can only be sent by SC creator
	if !sc.From(sc.ContractCreator()) {
		sc.Panic("Cancel spoofed request")
	}

	playPeriod := sc.Params().GetInt(KeyPlayPeriod).Value()
	if playPeriod < 10 {
		sc.Panic("Invalid play period...")
	}

	sc.State().GetInt(KeyPlayPeriod).SetValue(playPeriod)
}
