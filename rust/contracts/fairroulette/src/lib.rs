// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use schema::*;
use types::*;
use wasplib::client::*;

mod schema;
mod types;

const NUM_COLORS: i64 = 5;
const DEFAULT_PLAY_PERIOD: i64 = 120;

fn func_place_bet(ctx: &ScCallContext) {
    let amount = ctx.incoming().balance(&ScColor::IOTA);
    if amount == 0 {
        ctx.panic("Empty bet...");
    }
    let color = ctx.params().get_int(PARAM_COLOR).value();
    if color == 0 {
        ctx.panic("No color...");
    }
    if color < 1 || color > NUM_COLORS {
        ctx.panic("Invalid color...");
    }

    let bet = BetInfo {
        better: ctx.caller(),
        amount: amount,
        color: color,
    };

    let state = ctx.state();
    let bets = state.get_bytes_array(VAR_BETS);
    let bet_nr = bets.length();
    bets.get_bytes(bet_nr).set_value(&encode_bet_info(&bet));
    if bet_nr == 0 {
        let mut play_period = state.get_int(VAR_PLAY_PERIOD).value();
        if play_period < 10 {
            play_period = DEFAULT_PLAY_PERIOD;
        }
        ctx.post(&PostRequestParams {
            contract: ctx.contract_id(),
            function: Hname::new("lock_bets"),
            params: None,
            transfer: None,
            delay: play_period,
        });
    }
}

fn func_lock_bets(ctx: &ScCallContext) {
    // can only be sent by SC itself
    if !ctx.from(&ctx.contract_id().as_agent()) {
        ctx.panic("Cancel spoofed request");
    }

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
        contract: ctx.contract_id(),
        function: Hname::new("pay_winners"),
        params: None,
        transfer: None,
        delay: 0,
    });
}

fn func_pay_winners(ctx: &ScCallContext) {
    // can only be sent by SC itself
    let sc_id = ctx.contract_id().as_agent();
    if !ctx.from(&sc_id) {
        ctx.panic("Cancel spoofed request");
    }

    let winning_color = ctx.utility().random(5) + 1;
    let state = ctx.state();
    state.get_int(VAR_LAST_WINNING_COLOR).set_value(winning_color);

    // gather all winners and calculate some totals
    let mut total_bet_amount = 0_i64;
    let mut total_win_amount = 0_i64;
    let locked_bets = state.get_bytes_array(VAR_LOCKED_BETS);
    let mut winners: Vec<BetInfo> = Vec::new();
    let nr_bets = locked_bets.length();
    for i in 0..nr_bets {
        let bet = decode_bet_info(&locked_bets.get_bytes(i).value());
        total_bet_amount += bet.amount;
        if bet.color == winning_color {
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

fn func_play_period(ctx: &ScCallContext) {
    // can only be sent by SC creator
    if !ctx.from(&ctx.contract_creator()) {
        ctx.panic("Cancel spoofed request");
    }

    let play_period = ctx.params().get_int(PARAM_PLAY_PERIOD).value();
    if play_period < 10 {
        ctx.panic("Invalid play period...");
    }

    ctx.state().get_int(VAR_PLAY_PERIOD).set_value(play_period);
}
