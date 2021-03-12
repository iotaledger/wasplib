// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

const MaxNumber = 5
const DefaultPlayPeriod = 120

func funcLockBets(ctx wasmlib.ScFuncContext, params *FuncLockBetsParams) {
	// move all current bets to the locked_bets array
	state := ctx.State()
	bets := state.GetBytesArray(VarBets)
	lockedBets := state.GetBytesArray(VarLockedBets)
	nrBets := bets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := bets.GetBytes(i).Value()
		lockedBets.GetBytes(i).SetValue(bytes)
	}
	bets.Clear()

	ctx.PostSelf(HFuncPayWinners, nil, nil, 0)
}

func funcPayWinners(ctx wasmlib.ScFuncContext, params *FuncPayWinnersParams) {
	scId := ctx.ContractId().AsAgentId()
	winningNumber := ctx.Utility().Random(5) + 1
	state := ctx.State()
	state.GetInt64(VarLastWinningNumber).SetValue(winningNumber)

	// gather all winners and calculate some totals
	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	lockedBets := state.GetBytesArray(VarLockedBets)
	winners := make([]*Bet, 0)
	nrBets := lockedBets.Length()
	for i := int32(0); i < nrBets; i++ {
		bet := NewBetFromBytes(lockedBets.GetBytes(i).Value())
		totalBetAmount += bet.Amount
		if bet.Number == winningNumber {
			totalWinAmount += bet.Amount
			winners = append(winners, bet)
		}
	}
	lockedBets.Clear()

	if len(winners) == 0 {
		ctx.Log("Nobody wins!")
		// compact separate bet deposit UTXOs into a single one
		ctx.TransferToAddress(scId.Address(), wasmlib.NewScTransfer(wasmlib.IOTA, totalBetAmount))
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
			ctx.TransferToAddress(bet.Better.Address(), wasmlib.NewScTransfer(wasmlib.IOTA, payout))
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
		ctx.TransferToAddress(scId.Address(), wasmlib.NewScTransfer(wasmlib.IOTA, remainder))
	}
}

func funcPlaceBet(ctx wasmlib.ScFuncContext, params *FuncPlaceBetParams) {
	amount := ctx.Incoming().Balance(wasmlib.IOTA)
	if amount == 0 {
		ctx.Panic("Empty bet...")
	}
	number := params.Number.Value()
	if number < 1 || number > MaxNumber {
		ctx.Panic("Invalid number...")
	}

	bet := &Bet{
		Better: ctx.Caller(),
		Amount: amount,
		Number: number,
	}

	state := ctx.State()
	bets := state.GetBytesArray(VarBets)
	betNr := bets.Length()
	bets.GetBytes(betNr).SetValue(bet.Bytes())
	if betNr == 0 {
		playPeriod := state.GetInt64(VarPlayPeriod).Value()
		if playPeriod < 10 {
			playPeriod = DefaultPlayPeriod
		}
		ctx.PostSelf(HFuncLockBets, nil, nil, playPeriod)
	}
}

func funcPlayPeriod(ctx wasmlib.ScFuncContext, params *FuncPlayPeriodParams) {
	playPeriod := params.PlayPeriod.Value()
	ctx.Require(playPeriod >= 10, "invalid play period")
	ctx.State().GetInt64(VarPlayPeriod).SetValue(playPeriod)
}

func viewLastWinningNumber(ctx wasmlib.ScViewContext, params *ViewLastWinningNumberParams) {
	// Create an ScImmutableMap proxy to the state storage map on the host.
	state := ctx.State()

	// Get the 'lastWinningNumber' int64 value from state storage through
	// an ScImmutableInt64 proxy.
	lastWinningNumber := state.GetInt64(VarLastWinningNumber).Value()

	// Create an ScMutableMap proxy to the map on the host that will store the
	// key/value pairs that we want to return from this View function
	results := ctx.Results()

	// Set the value associated with the 'lastWinningNumber' key to the value
	// we got from state storage
	results.GetInt64(VarLastWinningNumber).SetValue(lastWinningNumber)
}
