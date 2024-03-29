// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// This example implements 'fairroulette', a simple smart contract that can automatically handle
// an unlimited amount of bets on a number during a timed betting round. Once a betting round
// is over the contract will automatically pay out the winners proportionally to their bet amount.
// The intent is to showcase basic functionality of WasmLib and timed calling of functions
// through a minimal implementation and not to come up with a complete real-world solution.

package fairroulette

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

// Define some default configuration parameters.

// The maximum number one can bet on. The range of numbers starts at 1.
const MaxNumber = 5

// The default playing period of one betting round in seconds.
const DefaultPlayPeriod = 1200

// 'placeBet' is used by betters to place a bet on a number from 1 to MAX_NUMBER. The first
// incoming bet triggers a betting round of configurable duration. After the playing period
// expires the smart contract will automatically pay out any winners and start a new betting
// round upon arrival of a new bet.
// The 'placeBet' function takes 1 mandatory parameter:
// - 'number', which must be an Int64 number from 1 to MAX_NUMBER
// The 'member' function will save the number together with the address of the better and
// the amount of incoming iotas as the bet amount in its state.
func funcPlaceBet(ctx wasmlib.ScFuncContext, f *PlaceBetContext) {
	// Since we are sure that the 'number' parameter actually exists we can
	// retrieve its actual value into an i64.
	number := f.Params.Number().Value()

	// Require that the number is a valid number to bet on, otherwise panic out.
	ctx.Require(number >= 1 && number <= MaxNumber, "invalid number")

	// Create ScBalances proxy to the incoming balances for this request.
	// Note that ScBalances wraps an ScImmutableMap of token color/amount combinations
	// in a simpler to use interface.
	incoming := ctx.Incoming()

	// Retrieve the amount of plain iota tokens that are part of the incoming balance.
	amount := incoming.Balance(wasmlib.IOTA)

	// Require that there are actually some plain iotas there
	ctx.Require(amount > 0, "empty bet")

	// Now we gather all information together into a single serializable struct
	// Note that we use the caller() method of the function context to determine
	// the agent id of the better. This is where a potential pay-out will be sent.
	bet := &Bet{
		Better: ctx.Caller(),
		Amount: amount,
		Number: number,
	}

	// Get the array of current bets from state storage.
	bets := f.State.Bets()

	// Determine what the next bet number is by retrieving the length of the bets array.
	betNr := bets.Length()

	// Append the bet data to the bets array. The bet array will automatically take care
	// of serializing the bet struct into a bytes representation.
	bets.GetBet(betNr).SetValue(bet)

	// Was this the first bet of this round?
	if betNr == 0 {
		// Yes it was, query the state for the length of the playing period in seconds by
		// retrieving the playPeriod value from state storage
		playPeriod := f.State.PlayPeriod().Value()

		// if the play period is less than 10 seconds we override it with the default duration.
		// Note that this will also happen when the play period was not set yet because in that
		// case a zero value was returned.
		if playPeriod < 10 {
			playPeriod = DefaultPlayPeriod
		}

		// And now for our next trick we post a delayed request to ourselves on the Tangle.
		// We are requesting to call the 'lockBets' function, but delay it for the play_period
		// amount of seconds. This will lock in the playing period, during which more bets can
		// be placed. Once the 'lockBets' function gets triggered by the ISCP it will gather all
		// bets up to that moment as the ones to consider for determining the winner.
		ScFuncs.LockBets(ctx).Func.Delay(playPeriod).TransferIotas(1).Post()
	}
}

// 'lockBets' is a function whose execution gets initiated by the 'placeBets' function as soon as
// the first bet comes in and will be triggered after a configurable number of seconds that defines
// the length of the playing round that started with that first bet. While this function is waiting
// to get triggered by the ISCP at the correct time any other incoming bets are added to the "bets"
// array in state storage. Once the 'lockBets' function gets triggered it will move all bets to a
// second state storage array called "lockedBets", after which it will request the 'payWinners'
// function to be run. Note that any bets coming in after that moment will start the cycle from
// scratch, with the first incoming bet triggering a new delayed execution of 'lockBets'.
func funcLockBets(ctx wasmlib.ScFuncContext, f *LockBetsContext) {
	// Get the bets array in state storage.
	bets := f.State.Bets()

	// Get the lockedBets array in state storage.
	lockedBets := f.State.LockedBets()

	// Determine the amount of bets in the 'bets' array.
	nrBets := bets.Length()

	// Copy all bet data from the 'bets' array to the 'lockedBets' array by
	// looping through all indexes of the array and copying the bets one by one.
	for i := int32(0); i < nrBets; i++ {
		// Get the bet data stored at the next index in the 'bets' array.
		bet := bets.GetBet(i).Value()

		// Save the bet data at the next index in the 'lockedBets' array.
		lockedBets.GetBet(i).SetValue(bet)
	}

	// Now that we have a copy of all bets it is safe to clear the 'bets' array
	// This will reset the length to zero, so that the next incoming bet will once
	// again trigger the delayed 'lockBets' request.
	bets.Clear()

	// Next we trigger an immediate request to the 'payWinners' function
	ScFuncs.PayWinners(ctx).Func.TransferIotas(1).Post()
}

// 'payWinners' is a function whose execution gets initiated by the 'lockBets' function.
// The reason that the 'lockBets' function does not immediately take care of paying the winners
// itself is that we need to introduce some unpredictability in the outcome of the randomizer
// used in selecting the winning number. To prevent people from observing the 'lockBets' request
// and potentially calculating the winning value in advance the 'lockBets' function instead asks
// the 'payWinners' function to do this once the bets have been locked. This will generate a new
// transaction with completely unpredictable transaction hash. This hash is what we will use as
// a deterministic source of entropy for the random number generator. In this way every node in
// the committee will be using the same pseudo-random value sequence, which in turn makes sure
// that all nodes can agree on the outcome.
func funcPayWinners(ctx wasmlib.ScFuncContext, f *PayWinnersContext) {
	// Use the built-in random number generator which has been automatically initialized by
	// using the transaction hash as initial entropy data. Note that the pseudo-random number
	// generator will use the next 8 bytes from the hash as its random Int64 number and once
	// it runs out of data it simply hashes the previous hash for a next pseudo-random sequence.
	// Here we determine the winning number for this round in the range of 1 thru MaxNumber.
	winningNumber := ctx.Utility().Random(MaxNumber) + 1

	// Save the last winning number in state storage under 'lastWinningNumber' so that there
	// is (limited) time for people to call the 'getLastWinningNumber' View to verify the last
	// winning number if they wish. Note that this is just a silly example. We could log much
	// more extensive statistics information about each playing round in state storage and
	// make that data available through views for anyone to see.
	f.State.LastWinningNumber().SetValue(winningNumber)

	// Gather all winners and calculate some totals at the same time.
	// Keep track of the total bet amount, the total win amount, and all the winners.
	// Note how we decided to keep the winners in a local vector instead of creating
	// yet another array in state storage or having to go through lockedBets again.
	totalBetAmount := int64(0)
	totalWinAmount := int64(0)
	winners := make([]*Bet, 0)

	// Get the 'lockedBets' array in state storage.
	lockedBets := f.State.LockedBets()

	// Determine the amount of bets in the 'lockedBets' array.
	nrBets := lockedBets.Length()

	// Loop through all indexes of the 'lockedBets' array.
	for i := int32(0); i < nrBets; i++ {
		// Retrieve the bet stored at the next index
		bet := lockedBets.GetBet(i).Value()

		// Add this bet's amount to the running total bet amount
		totalBetAmount += bet.Amount

		// Did this better bet on the winning number?
		if bet.Number == winningNumber {
			// Yes, add this bet amount to the running total win amount.
			totalWinAmount += bet.Amount

			// And save this bet in the winners vector.
			winners = append(winners, bet)
		}
	}

	// Now that we preprocessed all bets we can get rid of the data in state storage so that
	// the 'lockedBets' array becomes available for when the next betting round ends.
	lockedBets.Clear()

	payWinnersProportionally(ctx, winners, totalBetAmount, totalWinAmount)
}

func payWinnersProportionally(ctx wasmlib.ScFuncContext, winners []*Bet, totalBetAmount, totalWinAmount int64) {
	// Did we have any winners at all?
	if len(winners) == 0 {
		// No winners, log this fact to the log on the host.
		ctx.Log("Nobody wins!")
	}

	// Pay out the winners proportionally to their bet amount. Note that we could configure
	// a small percentage that would go to the owner of the smart contract as hosting payment.

	// Keep track of the total payout so we can calculate the remainder after truncation.
	totalPayout := int64(0)

	// Loop through all winners.
	size := len(winners)
	for i := 0; i < size; i++ {
		// Get the next winner.
		bet := winners[i]

		// Determine the proportional win amount (we could take our percentage here)
		payout := totalBetAmount * bet.Amount / totalWinAmount

		// Anything to pay to the winner?
		if payout != 0 {
			// Yep, keep track of the running total payout
			totalPayout += payout

			// Set up an ScTransfers proxy that transfers the correct amount of iotas.
			// Note that ScTransfers wraps an ScMutableMap of token color/amount combinations
			// in a simpler to use interface. The constructor we use here creates and initializes
			// a single token color transfer in a single statement. The actual color and amount
			// values passed in will be stored in a new map on the host.
			transfers := wasmlib.NewScTransferIotas(payout)

			// Perform the actual transfer of tokens from the smart contract to the address
			// of the winner. The transfer_to_address() method receives the address value and
			// the proxy to the new transfers map on the host, and will call the corresponding
			// host sandbox function with these values.
			ctx.TransferToAddress(bet.Better.Address(), transfers)
		}

		// Log who got sent what in the log on the host.
		text := "Pay " + ctx.Utility().String(payout) + " to " + bet.Better.String()
		ctx.Log(text)
	}

	// This is where we transfer the remainder after payout to the creator of the smart contract.
	// The bank always wins :-P
	remainder := totalBetAmount - totalPayout
	if remainder != 0 {
		// We have a remainder. First create a transfer for the remainder.
		transfers := wasmlib.NewScTransferIotas(remainder)

		// Send the remainder to the contract creator.
		ctx.TransferToAddress(ctx.ContractCreator().Address(), transfers)
	}
}

// 'playPeriod' can be used by the contract creator to set the length of a betting round
// to a different value than the default value, which is 120 seconds.
func funcPlayPeriod(ctx wasmlib.ScFuncContext, f *PlayPeriodContext) {
	// Since we are sure that the 'playPeriod' parameter actually exists we can
	// retrieve its actual value into an i32 value.
	playPeriod := f.Params.PlayPeriod().Value()

	// Require that the play period (in seconds) is not ridiculously low.
	// Otherwise panic out with an error message.
	ctx.Require(playPeriod >= 10, "invalid play period")

	// Now we set the corresponding variable 'playPeriod' in state storage.
	f.State.PlayPeriod().SetValue(playPeriod)
}

func viewLastWinningNumber(ctx wasmlib.ScViewContext, f *LastWinningNumberContext) {
	// Get the 'lastWinningNumber' int64 value from state storage.
	lastWinningNumber := f.State.LastWinningNumber().Value()

	// Set the 'lastWinningNumber' in results to the value from state storage.
	f.Results.LastWinningNumber().SetValue(lastWinningNumber)
}
