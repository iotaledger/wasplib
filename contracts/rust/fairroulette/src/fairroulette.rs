// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

const MAX_NUMBER: i64 = 5;
const DEFAULT_PLAY_PERIOD: i64 = 120;

pub fn func_lock_bets(ctx: &ScFuncContext) {
    // only SC itself can invoke this function
    ctx.require(ctx.caller() == ctx.contract_id().as_agent_id(), "no permission");

    // move all current bets to the locked_bets array
    let state = ctx.state();
    let bets = state.get_bytes_array(VAR_BETS);
    let locked_bets = state.get_bytes_array(VAR_LOCKED_BETS);
    let nr_bets = bets.length();
    for i in 0..nr_bets {
        let bytes = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bytes);
    }
    bets.clear();

    ctx.post(&PostRequestParams {
        contract_id: ctx.contract_id(),
        function: HFUNC_PAY_WINNERS,
        params: None,
        transfer: None,
        delay: 0,
    });
}

pub fn func_pay_winners(ctx: &ScFuncContext) {
    // only SC itself can invoke this function
    ctx.require(ctx.caller() == ctx.contract_id().as_agent_id(), "no permission");

    let sc_id = ctx.contract_id().as_agent_id();
    let winning_number = ctx.utility().random(5) + 1;
    let state = ctx.state();
    state.get_int(VAR_LAST_WINNING_NUMBER).set_value(winning_number);

    // gather all winners and calculate some totals
    let mut total_bet_amount = 0_i64;
    let mut total_win_amount = 0_i64;
    let locked_bets = state.get_bytes_array(VAR_LOCKED_BETS);
    let mut winners: Vec<Bet> = Vec::new();
    let nr_bets = locked_bets.length();
    for i in 0..nr_bets {
        let bet = Bet::from_bytes(&locked_bets.get_bytes(i).value());
        total_bet_amount += bet.amount;
        if bet.number == winning_number {
            total_win_amount += bet.amount;
            winners.push(bet);
        }
    }
    locked_bets.clear();

    if winners.is_empty() {
        ctx.log("Nobody wins!");
        // compact separate bet deposit UTXOs into a single one
        ctx.transfer_to_address(&sc_id.address(), &ScTransfers::new(&ScColor::IOTA, total_bet_amount));
        return;
    }

    // pay out the winners proportionally to their bet amount
    let mut total_payout = 0_i64;
    let size = winners.len();
    for i in 0..size {
        let bet = &winners[i];
        let payout = total_bet_amount * bet.amount / total_win_amount;
        if payout != 0 {
            total_payout += payout;
            ctx.transfer_to_address(&bet.better.address(), &ScTransfers::new(&ScColor::IOTA, payout));
        }
        let text = "Pay ".to_string() + &payout.to_string() +
            " to " + &bet.better.to_string();
        ctx.log(&text);
    }

    // any truncation left-overs are fair picking for the smart contract
    if total_payout != total_bet_amount {
        let remainder = total_bet_amount - total_payout;
        let text = "Remainder is ".to_string() + &remainder.to_string();
        ctx.log(&text);
        ctx.transfer_to_address(&sc_id.address(), &ScTransfers::new(&ScColor::IOTA, remainder));
    }
}

pub fn func_place_bet(ctx: &ScFuncContext) {
    let p = ctx.params();
    let param_number = p.get_int(PARAM_NUMBER);

    ctx.require(param_number.exists(), "missing mandatory number");

    let amount = ctx.incoming().balance(&ScColor::IOTA);
    if amount == 0 {
        ctx.panic("Empty bet...");
    }
    let number = param_number.value();
    if number < 1 || number > MAX_NUMBER {
        ctx.panic("Invalid number...");
    }

    let bet = Bet {
        better: ctx.caller(),
        amount: amount,
        number: number,
    };

    let state = ctx.state();
    let bets = state.get_bytes_array(VAR_BETS);
    let bet_nr = bets.length();
    bets.get_bytes(bet_nr).set_value(&bet.to_bytes());
    if bet_nr == 0 {
        let mut play_period = state.get_int(VAR_PLAY_PERIOD).value();
        if play_period < 10 {
            play_period = DEFAULT_PLAY_PERIOD;
        }
        ctx.post(&PostRequestParams {
            contract_id: ctx.contract_id(),
            function: HFUNC_LOCK_BETS,
            params: None,
            transfer: None,
            delay: play_period,
        });
    }
}

pub fn func_play_period(ctx: &ScFuncContext) {
    // only SC creator can update the play period
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let param_play_period = p.get_int(PARAM_PLAY_PERIOD);

    ctx.require(param_play_period.exists(), "missing mandatory playPeriod");

    let play_period = param_play_period.value();
    if play_period < 10 {
        ctx.panic("Invalid play period...");
    }

    ctx.state().get_int(VAR_PLAY_PERIOD).set_value(play_period);
}
