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

func placeBet(ctx *client.ScCallContext) {
	amount := ctx.Incoming().Balance(client.IOTA)
	if amount == 0 {
		ctx.Panic("Empty bet...")
	}
	color := ctx.Params().GetInt(KeyColor).Value()
	if color == 0 {
		ctx.Panic("No color...")
	}
	if color < 1 || color > NumColors {
		ctx.Panic("Invalid color...")
	}

	bet := &BetInfo{
		Better: ctx.Caller(),
		Amount: amount,
		Color:  color,
	}

	state := ctx.State()
	bets := state.GetBytesArray(KeyBets)
	betNr := bets.Length()
	bets.GetBytes(betNr).SetValue(EncodeBetInfo(bet))
	if betNr == 0 {
		playPeriod := state.GetInt(KeyPlayPeriod).Value()
		if playPeriod < 10 {
			playPeriod = DefaultPlayPeriod
		}
		ctx.Post(&client.PostRequestParams{
			Contract: ctx.ContractId(),
			Function: client.NewHname("lock_bets"),
			Params:   nil,
			Transfer: nil,
			Delay:    playPeriod,
		})
	}
}

func lockBets(ctx *client.ScCallContext) {
	// can only be sent by SC itself
	if !ctx.From(ctx.ContractId().AsAgent()) {
		ctx.Panic("Cancel spoofed request")
	}

	// move all current bets to the locked_bets array
	state := ctx.State()
	bets := state.GetBytesArray(KeyBets)
	lockedBets := state.GetBytesArray(KeyLockedBets)
	nrBets := bets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := bets.GetBytes(i).Value()
		lockedBets.GetBytes(i).SetValue(bytes)
	}
	bets.Clear()

	ctx.Post(&client.PostRequestParams{
		Contract: ctx.ContractId(),
		Function: client.NewHname("pay_winners"),
		Params:   nil,
		Transfer: nil,
		Delay:    0,
	})
}

func payWinners(ctx *client.ScCallContext) {
	// can only be sent by SC itself
	scId := ctx.ContractId().AsAgent()
	if !ctx.From(scId) {
		ctx.Panic("Cancel spoofed request")
	}

	winningColor := ctx.Utility().Random(5) + 1
	state := ctx.State()
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
		ctx.Log("Nobody wins!")
		// compact separate bet deposit UTXOs into a single one
		ctx.TransferToAddress(scId.Address(), client.NewScTransfer(client.IOTA, totalBetAmount))
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
			ctx.TransferToAddress(bet.Better.Address(), client.NewScTransfer(client.IOTA, payout))
		}
		text := "Pay " + ctx.Utility().String(payout) +
			" to " + bet.Better.String()
		ctx.Log(text)
	}

	// any truncation left-overs are fair picking for the smart contract
	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + ctx.Utility().String(remainder)
		ctx.Log(text)
		ctx.TransferToAddress(scId.Address(), client.NewScTransfer(client.IOTA, remainder))
	}
}

func playPeriod(ctx *client.ScCallContext) {
	// can only be sent by SC creator
	if !ctx.From(ctx.ContractCreator()) {
		ctx.Panic("Cancel spoofed request")
	}

	playPeriod := ctx.Params().GetInt(KeyPlayPeriod).Value()
	if playPeriod < 10 {
		ctx.Panic("Invalid play period...")
	}

	ctx.State().GetInt(KeyPlayPeriod).SetValue(playPeriod)
}
