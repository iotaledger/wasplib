// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/iotaledger/wasplib/client"
	"strconv"
)

const NUM_COLORS int64 = 5
const PLAY_PERIOD int64 = 120

type BetInfo struct {
	id     *client.ScRequestId
	sender *client.ScAgent
	amount int64
	color  int64
}

func main() {
}

//export onLoad
func onLoadFairRoulette() {
	exports := client.NewScExports()
	exports.AddCall("placeBet", placeBet)
	exports.AddCall("lockBets", lockBets)     //TODO sc internal only
	exports.AddCall("payWinners", payWinners) //TODO sc internal only
	exports.AddCall("playPeriod", playPeriod)
	exports.AddCall("nothing", client.Nothing)
}

func placeBet(sc *client.ScCallContext) {
	request := sc.Request()
	amount := request.Balance(client.IOTA)
	if amount == 0 {
		sc.Log("Empty bet...")
		return
	}
	color := request.Params().GetInt("color").Value()
	if color == 0 {
		sc.Log("No color...")
		return
	}
	if color < 1 || color > NUM_COLORS {
		sc.Log("Invalid color...")
		return
	}

	bet := BetInfo{
		id:     request.Id(),
		sender: request.Sender(),
		amount: amount,
		color:  color,
	}

	state := sc.State()
	bets := state.GetBytesArray("bets")
	betNr := bets.Length()
	bytes := encodeBetInfo(&bet)
	bets.GetBytes(betNr).SetValue(bytes)
	if betNr == 0 {
		playPeriod := state.GetInt("playPeriod").Value()
		if playPeriod < 10 {
			playPeriod = PLAY_PERIOD
		}
		sc.PostSelf("lockBets").Post(playPeriod)
	}
}

func lockBets(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.Request().From(sc.Contract().Id()) {
		sc.Log("Cancel spoofed request")
		return
	}

	state := sc.State()
	bets := state.GetBytesArray("bets")
	lockedBets := state.GetBytesArray("lockedBets")
	nrBets := bets.Length()
	for i := int32(0); i < nrBets; i++ {
		bytes := bets.GetBytes(i).Value()
		lockedBets.GetBytes(i).SetValue(bytes)
	}
	bets.Clear()

	sc.PostSelf("payWinners").Post(0)
}

func payWinners(sc *client.ScCallContext) {
	// can only be sent by SC itself
	scId := sc.Contract().Id()
	if !sc.Request().From(scId) {
		sc.Log("Cancel spoofed request")
		return
	}

	winningColor := sc.Utility().Random(5) + 1
	state := sc.State()
	state.GetInt("lastWinningColor").SetValue(winningColor)

	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	lockedBets := state.GetBytesArray("lockedBets")
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

	totalPayout := int64(0)
	for i := 0; i < len(winners); i++ {
		bet := winners[i]
		payout := totalBetAmount * bet.amount / totalWinAmount
		if payout != 0 {
			totalPayout += payout
			sc.Transfer(bet.sender, client.IOTA, payout)
		}
		text := "Pay " + strconv.FormatInt(payout, 10) + " to " + bet.sender.String()
		sc.Log(text)
	}

	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + strconv.FormatInt(remainder, 10)
		sc.Log(text)
		sc.Transfer(scId, client.IOTA, remainder)
	}
}

func playPeriod(sc *client.ScCallContext) {
	// can only be sent by SC owner
	if !sc.Request().From(sc.Contract().Owner()) {
		sc.Log("Cancel spoofed request")
		return
	}

	playPeriod := sc.Request().Params().GetInt("playPeriod").Value()
	if playPeriod < 10 {
		sc.Log("Invalid play period...")
		return
	}

	sc.State().GetInt("playPeriod").SetValue(playPeriod)
}

func decodeBetInfo(bytes []byte) *BetInfo {
	decoder := client.NewBytesDecoder(bytes)
	return &BetInfo{
		id:     decoder.RequestId(),
		sender: decoder.Agent(),
		amount: decoder.Int(),
		color:  decoder.Int(),
	}
}

func encodeBetInfo(bet *BetInfo) []byte {
	return client.NewBytesEncoder().
		RequestId(bet.id).
		Agent(bet.sender).
		Int(bet.amount).
		Int(bet.color).
		Data()
}
