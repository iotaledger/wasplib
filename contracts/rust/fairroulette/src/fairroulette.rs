// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// This example implements 'fairroulette', a simple smart contract that can automatically handle
// an unlimited amount of bets bets on a number during a timed betting round. Once a betting round
// is over the contract will automatically pay out the winners proportionally to their bet amount.
// The intent is to showcase basic functionality of WasmLib and timed calling of functions
// through a minimal implementation and not to come up with a complete real-world solution.

use wasmlib::*;

use crate::*;
use crate::types::*;

// define some default configuration parameters

// the maximum number one can bet on. The range of numbers starts at 1.
const MAX_NUMBER: i64 = 5;
// the default playing period of one betting round in minutes
const DEFAULT_PLAY_PERIOD: i64 = 120;

// 'placeBet' is used by betters to place a bet on a number from 1 to MAX_NUMBER. The first
// incoming bet triggers a betting round of configurable duration. After the playing period
// expires the smart contract will automatically pay out any winners and start a new betting
// round upon arrival of a new bet.
// The 'placeBet' function takes 1 mandatory parameter:
// - 'number', which must be s an Int64 number from 1 to MAX_NUMBER
// The 'member' function will save the number together with the address of the better and
// the amount of incoming iotas as the bet amount in its state.
pub fn func_place_bet(ctx: &ScFuncContext, f: &FuncPlaceBetContext) {

    // Since we are sure that the 'number' parameter actually exists we can
    // retrieve its actual value into an i64.
    let number: i64 = f.params.number.value();
    // require that the number is a valid number to bet on, otherwise panic out.
    ctx.require(number >= 1 && number <= MAX_NUMBER, "invalid number");

    // Create ScBalances proxy to the incoming balances for this Func request.
    // Note that ScBalances wraps an ScImmutableMap of token color/amount combinations
    // in a simpler to use interface.
    let incoming: ScBalances = ctx.incoming();

    // Retrieve the amount of plain iota tokens from the incoming balance
    let amount: i64 = incoming.balance(&ScColor::IOTA);

    // require that there are actually some iotas there
    ctx.require(amount > 0, "empty bet");

    // Now we gather all information together into a single serializable struct
    // Note that we use the caller() method of the function context to determine
    // the address of the better. This is the address where a pay-out will be sent.
    let bet = Bet {
        better: ctx.caller(),
        amount: amount,
        number: number,
    };

    // Create an ScMutableMap proxy to the state storage map on the host.

    // Create an ScMutableBytesArray proxy to a bytes array named "bets" in the state storage.
    let bets: ArrayOfMutableBet = f.state.bets();

    // Determine what the next bet number is by retrieving the length of the bets array.
    let bet_nr: i32 = bets.length();

    // Append the bet data to the bets array. We get an ScBytes proxy to the bytes stored
    // using the bet number as index. Then we set the bytes value in the best array on the
    // host to the result of serializing the bet data into a bytes representation.
    bets.get_bet(bet_nr).set_value(&bet);

    // Was this the first bet of this round?
    if bet_nr == 0 {
        // Yes it was, query the state for the length of the playing period in seconds by
        // retrieving the "playPeriod" from state storage
        let mut play_period: i64 = f.state.play_period().value();

        // if the play period is less than 10 seconds we override it with the default duration.
        // Note that this will also happen when the play period was not set yet because in that
        // case a zero value was returned.
        if play_period < 10 {
            play_period = DEFAULT_PLAY_PERIOD;
        }

        // And now for our next trick we post a delayed request to ourselves on the Tangle.
        // We are requesting to call the 'lockBets' function, but delay it for the play_period
        // amount of seconds. This will lock in the playing period, during which more bets can
        // be placed. Once the 'lockBets' function gets triggered by the ISCP it will gather all
        // bets up to that moment as the ones to consider for determining the winner.
        let transfer = ScTransfers::iotas(1);
        ctx.post_self(HFUNC_LOCK_BETS, None, transfer, play_period);
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
pub fn func_lock_bets(ctx: &ScFuncContext, f: &FuncLockBetsContext) {

    // Create an ScMutableMap proxy to the state storage map on the host.

    // Create an ScMutableBytesArray proxy to the bytes array named 'bets' in state storage.
    let bets: ArrayOfMutableBet = f.state.bets();

    // Create an ScMutableBytesArray proxy to a bytes array named 'lockedBets' in state storage.
    let locked_bets: ArrayOfMutableBet = f.state.locked_bets();

    // Determine the amount of bets in the 'bets' array.
    let nr_bets: i32 = bets.length();

    // Copy all bet data from the 'bets' array to the 'lockedBets' array by
    // looping through all indexes of the array and copying the best one by one.
    for i in 0..nr_bets {

        // Get the bytes stored at the next index in the 'bets' array.
        // Save the bytes at the next index in the 'lockedBets' array.
        locked_bets.get_bet(i).set_value(&bets.get_bet(i).value());
    }

    // Now that we have a copy of all bets it is safe to clear the 'bets' array
    // This will reset the length to zero, so that the next incoming bet will once
    // again trigger the delayed 'lockBets' request.
    bets.clear();

    // Next we trigger an immediate request to the 'payWinners' function
    // See more explanation of the why below.
    let transfer = ScTransfers::iotas(1);
    ctx.post_self(HFUNC_PAY_WINNERS, None, transfer, 0);
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
pub fn func_pay_winners(ctx: &ScFuncContext, f: &FuncPayWinnersContext) {

    // Use the built-in random number generator which has been automatically initialized by
    // using the transaction hash as initial entropy data. Note that the pseudo-random number
    // generator will use the next 8 bytes from the hash as its random Int64 number and once
    // it runs out of data it simply hashes the previous hash for a next psuedo-random sequence.
    // Here we determine the winning number for this round in the range of 1 thru 5 (inclusive).
    let winning_number: i64 = ctx.utility().random(5) + 1;

    // Create an ScMutableMap proxy to the state storage map on the host.

    // Save the last winning number in state storage under 'lastWinningNumber' so that there is
    // (limited) time for people to call the 'getLastWinningNumber' View to verify the last winning
    // number if they wish. Note that this is just a silly example. We could log much more extensive
    // statistics information about each playing round in state storage and make that data available
    // through views for anyone to see.
    f.state.last_winning_number().set_value(winning_number);

    // Gather all winners and calculate some totals at the same time.
    // Keep track of the total bet amount, the total win amount, and all the winners
    let mut total_bet_amount: i64 = 0_i64;
    let mut total_win_amount: i64 = 0_i64;
    let mut winners: Vec<Bet> = Vec::new();

    // Create an ScMutableBytesArray proxy to the 'lockedBets' bytes array in state storage.
    let locked_bets: ArrayOfMutableBet = f.state.locked_bets();

    // Determine the amount of bets in the 'lockedBets' array.
    let nr_bets: i32 = locked_bets.length();

    // Loop through all indexes of the 'lockedBets' array.
    for i in 0..nr_bets {
        // Retrieve the bytes stored at the next index
        // Deserialize the bytes into the original Bet structure
        let bet: Bet = locked_bets.get_bet(i).value();

        // Add this bet amount to the running total bet ammount
        total_bet_amount += bet.amount;

        // Did this better bet on the winning number?
        if bet.number == winning_number {
            // Yes, add this bet amount to the running total win amount
            total_win_amount += bet.amount;

            // And save this bet in the winners vector
            winners.push(bet);
        }
    }

    // Now that we preprocessed all bets we can get rid of the data in state storage so that
    // the 'lockedBets' array is available for the next betting round.
    locked_bets.clear();

    // Did we have any winners at all?
    if winners.is_empty() {
        // No winners, log this fact to the log on the host.
        ctx.log("Nobody wins!");
    }

    // Pay out the winners proportionally to their bet amount. Note that we could configure
    // a small percentage that would go to the owner of the smart contract as hosting payment.

    // Keep track of the total payout so we can calculate the remainder after truncation
    let mut total_payout: i64 = 0_i64;

    // Loop through all winners
    let size: usize = winners.len();
    for i in 0..size {

        // Get the next winner
        let bet: &Bet = &winners[i];

        // Determine the proportional winning (we could take our percentage here)
        let payout: i64 = total_bet_amount * bet.amount / total_win_amount;

        // Anything to pay to the winner?
        if payout != 0 {
            // Yep, keep track of the running total payout
            total_payout += payout;

            // Set up an ScTransfers proxy that transfers the correct amount of iotas.
            // Note that ScTransfers wraps an ScMutableMap of token color/amount combinations
            // in a simpler to use interface. The constructor we use here creates and initializes
            // a single token color transfer in a single statement. The actual color and amount
            // values passed in will be stored in a new map on the host.
            let transfers: ScTransfers = ScTransfers::iotas(payout);

            // Perform the actual transfer of tokens from the smart contract to the better
            // address. The transfer_to_address() method receives the address value and
            // the proxy to the new transfers map on the host, and will call the corresponding
            // host sandbox function with these values.
            ctx.transfer_to_address(&bet.better.address(), transfers);
        }

        // Log who got sent what in the log on the host
        let text: String = "Pay ".to_string() + &payout.to_string() + " to " + &bet.better.to_string();
        ctx.log(&text);
    }

    // This is where we transfer the remainder after payout to the creator of the smart contract.
    // The bank always wins :-P
    let remainder: i64 = total_bet_amount - total_payout;
    if remainder != 0 {
        // We have a remainder First create a transfer for the remainder.
        let transfers: ScTransfers = ScTransfers::iotas(remainder);

        // Send the remainder to the contract creator.
        ctx.transfer_to_address(&ctx.contract_creator().address(), transfers);
    }
}

// 'playPeriod' can be used by the contract creator to set the length of a betting round
// to a different value than the default value, which is 120 seconds..
pub fn func_play_period(ctx: &ScFuncContext, f: &FuncPlayPeriodContext) {

    // Since we are sure that the 'playPeriod' parameter actually exists we can
    // retrieve its actual value into an i64 value.
    let play_period: i64 = f.params.play_period.value();

    // Require that the play period (in seconds) is not ridiculously low.
    // Otherwise panic out with an error message.
    ctx.require(play_period >= 10, "invalid play period");

    // Now we set the corresponding state variable 'playPeriod' through the state
    // map proxy to the value we just got.
    f.state.play_period().set_value(play_period);
}

pub fn view_last_winning_number(_ctx: &ScViewContext, f: &ViewLastWinningNumberContext) {

    // Create an ScImmutableMap proxy to the state storage map on the host.

    // Get the 'lastWinningNumber' int64 value from state storage through
    // an ScImmutableInt64 proxy.
    let last_winning_number: i64 = f.state.last_winning_number().value();

    // Create an ScMutableMap proxy to the map on the host that will store the
    // key/value pairs that we want to return from this View function

    // Set the value associated with the 'lastWinningNumber' key to the value
    // we got from state storage
    f.results.last_winning_number.set_value(last_winning_number);
}
