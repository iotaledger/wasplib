// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const KEY_BETS: &str = "bets";
const KEY_COLOR: &str = "color";
const KEY_LAST_WINNING_COLOR: &str = "last_winning_color";
const KEY_LOCKED_BETS: &str = "locked_bets";
const KEY_PLAY_PERIOD: &str = "play_period";

const NUM_COLORS: i64 = 5;
const DEFAULT_PLAY_PERIOD: i64 = 120;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("place_bet", place_bet);
    exports.add_call("lock_bets", lock_bets);
    exports.add_call("pay_winners", pay_winners);
    exports.add_call("play_period", play_period);
    exports.add_call("nothing", ScExports::nothing);
}

fn place_bet(sc: &ScCallContext) {
    let amount = sc.incoming().balance(&ScColor::IOTA);
    if amount == 0 {
        sc.panic("Empty bet...");
    }
    let color = sc.params().get_int(KEY_COLOR).value();
    if color == 0 {
        sc.panic("No color...");
    }
    if color < 1 || color > NUM_COLORS {
        sc.panic("Invalid color...");
    }

    let bet = BetInfo {
        better: sc.caller(),
        amount: amount,
        color: color,
    };

    let state = sc.state();
    let bets = state.get_bytes_array(KEY_BETS);
    let bet_nr = bets.length();
    bets.get_bytes(bet_nr).set_value(&encode_bet_info(&bet));
    if bet_nr == 0 {
        let mut play_period = state.get_int(KEY_PLAY_PERIOD).value();
        if play_period < 10 {
            play_period = DEFAULT_PLAY_PERIOD;
        }
        sc.post(&ScAddress::NULL,
                Hname::SELF,
                Hname::new("lock_bets"),
                ScMutableMap::NONE,
                ScTransfers::NONE,
                play_period);
    }
}

fn lock_bets(sc: &ScCallContext) {
    // can only be sent by SC itself
    if !sc.from(&sc.contract_id()) {
        sc.panic("Cancel spoofed request");
    }

    // move all current bets to the locked_bets array
    let state = sc.state();
    let bets = state.get_bytes_array(KEY_BETS);
    let locked_bets = state.get_bytes_array(KEY_LOCKED_BETS);
    let nr_bets = bets.length();
    for i in 0..nr_bets {
        let bytes = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bytes);
    }
    bets.clear();

    sc.post(&ScAddress::NULL,
            Hname::SELF,
            Hname::new("pay_winners"),
            ScMutableMap::NONE,
            ScTransfers::NONE,
            0);
}

fn pay_winners(sc: &ScCallContext) {
    // can only be sent by SC itself
    let sc_id = sc.contract_id();
    if !sc.from(&sc_id) {
        sc.panic("Cancel spoofed request");
    }

    let winning_color = sc.utility().random(5) + 1;
    let state = sc.state();
    state.get_int(KEY_LAST_WINNING_COLOR).set_value(winning_color);

    // gather all winners and calculate some totals
    let mut total_bet_amount = 0_i64;
    let mut total_win_amount = 0_i64;
    let locked_bets = state.get_bytes_array(KEY_LOCKED_BETS);
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
        sc.log("Nobody wins!");
        // compact separate bet deposit UTXOs into a single one
        sc.transfer(&sc_id, &ScColor::IOTA, total_bet_amount);
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
            sc.transfer(&bet.better, &ScColor::IOTA, payout);
        }
        let text = "Pay ".to_string() + &payout.to_string() + " to " + &bet.better.to_string();
        sc.log(&text);
    }

    // any truncation left-overs are fair picking for the smart contract
    if total_payout != total_bet_amount {
        let remainder = total_bet_amount - total_payout;
        let text = "Remainder is ".to_string() + &remainder.to_string();
        sc.log(&text);
        sc.transfer(&sc_id, &ScColor::IOTA, remainder);
    }
}

fn play_period(sc: &ScCallContext) {
    // can only be sent by SC creator
    if !sc.from(&sc.contract_creator()) {
        sc.panic("Cancel spoofed request");
    }

    let play_period = sc.params().get_int(KEY_PLAY_PERIOD).value();
    if play_period < 10 {
        sc.panic("Invalid play period...");
    }

    sc.state().get_int(KEY_PLAY_PERIOD).set_value(play_period);
}
