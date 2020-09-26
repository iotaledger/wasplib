package main

import (
	"github.com/iotaledger/wasplib/client"
	"strconv"
)

const NUM_COLORS int64 = 5
const PLAY_PERIOD int64 = 120

type BetInfo struct {
	id     string
	sender string
	color  int64
	amount int64
}

func main() {
}

//export placeBet
func placeBet() {
	ctx := client.NewScContext()
	request := ctx.Request()
	amount := request.Balance("iota")
	if amount == 0 {
		ctx.Log("Empty bet...")
		return
	}
	color := request.Params().GetInt("color").Value()
	if color == 0 {
		ctx.Log("No color...")
		return
	}
	if color < 1 || color > NUM_COLORS {
		ctx.Log("Invalid color...")
		return
	}

	bet := BetInfo{
		id:     request.Id(),
		sender: request.Address(),
		color:  color,
		amount: amount,
	}

	state := ctx.State()
	bets := state.GetBytesArray("bets")
	betNr := bets.Length()
	data := encodeBetInfo(&bet)
	bets.GetBytes(betNr).SetValue(data)
	if betNr == 0 {
		playPeriod := state.GetInt("playPeriod").Value()
		if playPeriod < 10 {
			playPeriod = PLAY_PERIOD
		}
		ctx.Event("", "lockBets", playPeriod)
	}
}

//export lockBets
func lockBets() {
	// can only be sent by SC itself
	ctx := client.NewScContext()
	if ctx.Request().Address() != ctx.Contract().Address() {
		ctx.Log("Cancel spoofed request")
		return
	}

	state := ctx.State()
	bets := state.GetStringArray("bets")
	lockedBets := state.GetStringArray("lockedBets")
	for i := int32(0); i < bets.Length(); i++ {
		bet := bets.GetString(i).Value()
		lockedBets.GetString(i).SetValue(bet)
	}
	bets.Clear()

	ctx.Event("", "payWinners", 0)
}

//export payWinners
func payWinners() {
	// can only be sent by SC itself
	ctx := client.NewScContext()
	scAddress := ctx.Contract().Address()
	if ctx.Request().Address() != scAddress {
		ctx.Log("Cancel spoofed request")
		return
	}

	winningColor := ctx.Random(5) + 1
	state := ctx.State()
	state.GetInt("lastWinningColor").SetValue(winningColor)

	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	lockedBets := state.GetBytesArray("lockedBets")
	winners := make([]*BetInfo, 0)
	for i := int32(0); i < lockedBets.Length(); i++ {
		betData := lockedBets.GetBytes(i).Value()
		bet := decodeBetInfo(betData)
		totalBetAmount += bet.amount
		if bet.color == winningColor {
			totalWinAmount += bet.amount
			winners = append(winners, bet)
		}
	}
	lockedBets.Clear()

	if len(winners) == 0 {
		ctx.Log("Nobody wins!")
		// compact separate UTXOs into a single one
		ctx.Transfer(scAddress, "iota", totalBetAmount)
		return
	}

	totalPayout := int64(0)
	for i := 0; i < len(winners); i++ {
		bet := winners[i]
		payout := totalBetAmount * bet.amount / totalWinAmount
		if payout != 0 {
			totalPayout += payout
			ctx.Transfer(bet.sender, "iota", payout)
		}
		text := "Pay " + strconv.FormatInt(payout, 10) + " to " + bet.sender
		ctx.Log(text)
	}

	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + strconv.FormatInt(remainder, 10)
		ctx.Log(text)
		ctx.Transfer(scAddress, "iota", remainder)
	}
}

//export playPeriod
func playPeriod() {
	// can only be sent by SC owner
	ctx := client.NewScContext()
	if ctx.Request().Address() != ctx.Contract().Owner() {
		ctx.Log("Cancel spoofed request")
		return
	}

	playPeriod := ctx.Request().Params().GetInt("playPeriod").Value()
	if playPeriod < 10 {
		ctx.Log("Invalid play period...")
		return
	}

	ctx.State().GetInt("playPeriod").SetValue(playPeriod)
}

func decodeBetInfo(bytes []byte) *BetInfo {
	decoder := client.NewBytesDecoder(bytes)
	return &BetInfo{
		id:     decoder.String(),
		sender: decoder.String(),
		amount: decoder.Int(),
		color:  decoder.Int(),
	}
}

func encodeBetInfo(data *BetInfo) []byte {
	return client.NewBytesEncoder().
		String(data.id).
		String(data.sender).
		Int(data.amount).
		Int(data.color).
		Data()
}
