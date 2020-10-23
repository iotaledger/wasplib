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

//export onLoad
func onLoadFairRoulette() {
	exports := client.NewScExports()
	exports.Add("placeBet")
	exports.Add("lockBets")   //TODO sc internal only
	exports.Add("payWinners") //TODO sc internal only
	exports.AddProtected("playPeriod")
	exports.Add("nothing")
}

//export placeBet
func placeBet() {
	sc := client.NewScContext()
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
		sender: request.Address(),
		color:  color,
		amount: amount,
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
		sc.PostRequest(sc.Contract().Address(), "lockBets", playPeriod)
	}
}

//export lockBets
func lockBets() {
	// can only be sent by SC itself
	sc := client.NewScContext()
	scAddress := sc.Contract().Address()
	if sc.Request().Address() != scAddress {
		sc.Log("Cancel spoofed request")
		return
	}

	state := sc.State()
	bets := state.GetStringArray("bets")
	lockedBets := state.GetStringArray("lockedBets")
	for i := int32(0); i < bets.Length(); i++ {
		bytes := bets.GetString(i).Value()
		lockedBets.GetString(i).SetValue(bytes)
	}
	bets.Clear()

	sc.PostRequest(scAddress, "payWinners", 0)
}

//export payWinners
func payWinners() {
	// can only be sent by SC itself
	sc := client.NewScContext()
	scAddress := sc.Contract().Address()
	if sc.Request().Address() != scAddress {
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
	for i := int32(0); i < lockedBets.Length(); i++ {
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
		sc.Transfer(scAddress, client.IOTA, totalBetAmount)
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
		text := "Pay " + strconv.FormatInt(payout, 10) + " to " + bet.sender
		sc.Log(text)
	}

	if totalPayout != totalBetAmount {
		remainder := totalBetAmount - totalPayout
		text := "Remainder is " + strconv.FormatInt(remainder, 10)
		sc.Log(text)
		sc.Transfer(scAddress, client.IOTA, remainder)
	}
}

//export playPeriod
func playPeriod() {
	// can only be sent by SC owner
	sc := client.NewScContext()
	if sc.Request().Address() != sc.Contract().Owner() {
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
		id:     decoder.String(),
		sender: decoder.String(),
		amount: decoder.Int(),
		color:  decoder.Int(),
	}
}

func encodeBetInfo(bet *BetInfo) []byte {
	return client.NewBytesEncoder().
		String(bet.id).
		String(bet.sender).
		Int(bet.amount).
		Int(bet.color).
		Data()
}
