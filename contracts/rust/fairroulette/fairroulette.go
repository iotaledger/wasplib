// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// This example implements 'fairroulette', a simple smart contract that can automatically handle
// an unlimited amount of bets bets on a number during a timed betting round. Once a betting round
// is over the contract will automatically pay out the winners proportionally to their bet amount.
// The intent is to showcase basic functionality of WasmLib and timed calling of functions
// through a minimal implementation and not to come up with a complete real-world solution.

package fairroulette

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

// define some default configuration parameters

// the maximum number one can bet on. The range of numbers starts at 1.
const MaxNumber = 5
// the default playing period of one betting round in minutes
const DefaultPlayPeriod = 120

// 'placeBet' is used by betters to place a bet on a number from 1 to MAX_NUMBER. The first
// incoming bet triggers a betting round of configurable duration. After the playing period
// expires the smart contract will automatically pay out any winners and start a new betting
// round upon arrival of a new bet.
// The 'placeBet' function takes 1 mandatory parameter:
// - 'number', which must be s an Int64 number from 1 to MAX_NUMBER
// The 'member' function will save the number together with the address of the better and
// the amount of incoming iotas as the bet amount in its state.
func funcPlaceBet(ctx wasmlib.ScFuncContext, params *FuncPlaceBetParams) {

    // Since we are sure that the 'number' parameter actually exists we can
    // retrieve its actual value into an i64.
    number := params.Number.Value()
    // require that the number is a valid number to bet on, otherwise panic out.
    ctx.Require(number >= 1 && number <= MaxNumber, "invalid number")

    // Create ScBalances proxy to the incoming balances for this Func request.
    // Note that ScBalances wraps an ScImmutableMap of token color/amount combinations
    // in a simpler to use interface.
    incoming := ctx.Incoming()

    // Retrieve the amount of plain iota tokens from the incoming balance
    amount := incoming.Balance(wasmlib.IOTA)

    // require that there are actually some iotas there
    ctx.Require(amount > 0, "empty bet")

    // Now we gather all information together into a single serializable struct
    // Note that we use the caller() method of the function context to determine
    // the address of the better. This is the address where a pay-out will be sent.
    bet := &Bet {
        Better: ctx.Caller(),
        Amount: amount,
        Number: number,
    }

    // Create an ScMutableMap proxy to the state storage map on the host.
    state := ctx.State()

    // Create an ScMutableBytesArray proxy to a bytes array named "bets" in the state storage.
    bets := state.GetBytesArray(VarBets)

    // Determine what the next bet number is by retrieving the length of the bets array.
    betNr := bets.Length()

    // Append the bet data to the bets array. We get an ScBytes proxy to the bytes stored
    // using the bet number as index. Then we set the bytes value in the best array on the
    // host to the result of serializing the bet data into a bytes representation.
    bets.GetBytes(betNr).SetValue(bet.Bytes())

    // Was this the first bet of this round?
    if betNr == 0 {
        // Yes it was, query the state for the length of the playing period in seconds by
        // retrieving the "playPeriod" from state storage
        playPeriod := state.GetInt64(VarPlayPeriod).Value()

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
        transfer := wasmlib.NewScTransferIotas(1)
        ctx.PostSelf(HFuncLockBets, nil, transfer, playPeriod)
    }
}

// 'lockBets' is a function whose execution gets initiated by the 'placeBets' function as soon as
// the first bet comes in and will be triggered after a configurable number of seconds that defines
// the length of the playing round started with that first bet. While this function is waiting to
// get triggered by the ISCP at the correct time any other incoming bets are added to the "bets"
// array in state storage. Once the 'lockBets' function gets triggered it will move all bets to a
// second state storage array called "lockedBets", after which it will request the 'payWinners'
// function to be run. Note that any bets coming in after that moment will start the cycle from
// scratch, with the first incoming bet triggering a new delayed execution of 'lockBets'.
func funcLockBets(ctx wasmlib.ScFuncContext, params *FuncLockBetsParams) {

    // Create an ScMutableMap proxy to the state storage map on the host.
    state := ctx.State()

    // Create an ScMutableBytesArray proxy to the bytes array named 'bets' in state storage.
    bets := state.GetBytesArray(VarBets)

    // Create an ScMutableBytesArray proxy to a bytes array named 'lockedBets' in state storage.
    lockedBets := state.GetBytesArray(VarLockedBets)

    // Determine the amount of bets in the 'bets' array.
    nrBets := bets.Length()

    // Copy all bet data from the 'bets' array to the 'lockedBets' array by
    // looping through all indexes of the array and copying the best one by one.
    for i := int32(0); i < nrBets; i++ {

        // Get the bytes stored at the next index in the 'bets' array.
        bytes:= bets.GetBytes(i).Value()

        // Save the bytes at the next index in the 'lockedBets' array.
        lockedBets.GetBytes(i).SetValue(bytes)
    }

    // Now that we have a copy of all bets it is safe to clear the 'bets' array
    // This will reset the length to zero, so that the next incoming bet will once
    // again trigger the delayed 'lockBets' request.
    bets.Clear()

    // Next we trigger an immediate request to the 'payWinners' function
    // See more explanation of the why below.
    transfer := wasmlib.NewScTransferIotas(1)
    ctx.PostSelf(HFuncPayWinners, nil, transfer, 0)
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
func funcPayWinners(ctx wasmlib.ScFuncContext, params *FuncPayWinnersParams) {

    // Use the built-in random number generator which has been automatically initialized by
    // using the transaction hash as initial entropy data. Note that the pseudo-random number
    // generator will use the next 8 bytes from the hash as its random Int64 number and once
    // it runs out of data it simply hashes the previous hash for a next psuedo-random sequence.
    // Here we determine the winning number for this round in the range of 1 thru 5 (inclusive).
    winningNumber := ctx.Utility().Random(5) + 1

    // Create an ScMutableMap proxy to the state storage map on the host.
    state := ctx.State()

    // Save the last winning number in state storage under 'lastWinningNumber' so that there is
    // (limited) time for people to call the 'getLastWinningNumber' View to verify the last winning
    // number if they wish. Note that this is just a silly example. We could log much more extensive
    // statistics information about each playing round in state storage and make that data available
    // through views for anyone to see.
    state.GetInt64(VarLastWinningNumber).SetValue(winningNumber)

    // Gather all winners and calculate some totals at the same time.
    // Keep track of the total bet amount, the total win amount, and all the winners
    totalBetAmount := int64(0)
    totalWinAmount := int64(0)
    winners:= make([]*Bet,0)

    // Create an ScMutableBytesArray proxy to the 'lockedBets' bytes array in state storage.
    lockedBets := state.GetBytesArray(VarLockedBets)

    // Determine the amount of bets in the 'lockedBets' array.
    nrBets := lockedBets.Length()

    // Loop through all indexes of the 'lockedBets' array.
    for i := int32(0); i < nrBets; i++ {
        // Retrieve the bytes stored at the next index
        bytes:= lockedBets.GetBytes(i).Value()

        // Deserialize the bytes into the original Bet structure
        bet := NewBetFromBytes(bytes)

        // Add this bet amount to the running total bet ammount
        totalBetAmount += bet.Amount

        // Did this better bet on the winning number?
        if bet.Number == winningNumber {
            // Yes, add this bet amount to the running total win amount
            totalWinAmount += bet.Amount

            // And save this bet in the winners vector
            winners = append(winners, bet)
        }
    }

    // Now that we preprocessed all bets we can get rid of the data in state storage so that
    // the 'lockedBets' array is available for the next betting round.
    lockedBets.Clear()

    // Did we have any winners at all?
    if len(winners) == 0 {
        // No winners, log this fact to the log on the host.
        ctx.Log("Nobody wins!")
    }

    // Pay out the winners proportionally to their bet amount. Note that we could configure
    // a small percentage that would go to the owner of the smart contract as hosting payment.

    // Keep track of the total payout so we can calculate the remainder after truncation
    totalPayout := int64(0)

    // Loop through all winners
    size := len(winners)
    for i := 0; i < size; i++ {

        // Get the next winner
        bet := winners[i]

        // Determine the proportional winning (we could take our percentage here)
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

            // Perform the actual transfer of tokens from the smart contract to the better
            // address. The transfer_to_address() method receives the address value and
            // the proxy to the new transfers map on the host, and will call the corresponding
            // host sandbox function with these values.
            ctx.TransferToAddress(bet.Better.Address(), transfers)
        }

        // Log who got sent what in the log on the host
        text := "Pay " + ctx.Utility().String(payout) + " to " + bet.Better.String()
        ctx.Log(text)
    }

    // This is where we transfer the remainder after payout to the creator of the smart contract.
    // The bank always wins :-P
    remainder := totalBetAmount - totalPayout
    if remainder != 0 {
        // We have a remainder First create a transfer for the remainder.
        transfers := wasmlib.NewScTransferIotas(remainder)

        // Send the remainder to the contract creator.
        ctx.TransferToAddress(ctx.ContractCreator().Address(), transfers)
    }
}

// 'playPeriod' can be used by the contract creator to set the length of a betting round
// to a different value than the default value, which is 120 seconds..
func funcPlayPeriod(ctx wasmlib.ScFuncContext, params *FuncPlayPeriodParams) {

    // Since we are sure that the 'playPeriod' parameter actually exists we can
    // retrieve its actual value into an i64 value.
    playPeriod := params.PlayPeriod.Value()

    // Require that the play period (in seconds) is not ridiculously low.
    // Otherwise panic out with an error message.
    ctx.Require(playPeriod >= 10, "invalid play period")

    // Now we set the corresponding state variable 'playPeriod' through the state
    // map proxy to the value we just got.
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
