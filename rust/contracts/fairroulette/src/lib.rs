// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

const KEY_BETS: &str = "bets";
const KEY_COLOR: &str = "color";
const KEY_LAST_WINNING_COLOR: &str = "lastWinningColor";
const KEY_LOCKED_BETS: &str = "lockedBets";
const KEY_PLAY_PERIOD: &str = "playPeriod";

const NUM_COLORS: i64 = 5;
const PLAY_PERIOD: i64 = 120;

struct BetInfo {
    better: ScAgent,
    amount: i64,
    color: i64,
}

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("placeBet", placeBet);
    exports.add_call("lockBets", lockBets);
    exports.add_call("payWinners", payWinners);
    exports.add_call("playPeriod", playPeriod);
    exports.add_call("nothing", ScExports::nothing);
}

fn placeBet(sc: &ScCallContext) {
    let amount = sc.incoming().balance(&ScColor::IOTA);
    if amount == 0 {
        sc.log("Empty bet...");
        return;
    }
    let color = sc.params().get_int(KEY_COLOR).value();
    if color == 0 {
        sc.log("No color...");
        return;
    }
    if color < 1 || color > NUM_COLORS {
        sc.log("Invalid color...");
        return;
    }

    let bet = BetInfo {
        better: sc.caller(),
        amount,
        color,
    };

    let state = sc.state();
    let bets = state.get_bytes_array(KEY_BETS);
    let bet_nr = bets.length();
    let bytes = encodeBetInfo(&bet);
    bets.get_bytes(bet_nr).set_value(&bytes);
    if bet_nr == 0 {
        let mut play_period = state.get_int(KEY_PLAY_PERIOD).value();
        if play_period < 10 {
            play_period = PLAY_PERIOD;
        }
        sc.post("lockBets").post(play_period);
    }
}

fn lockBets(sc: &ScCallContext) {
    // can only be sent by SC itself
    if !sc.from(&sc.contract().id()) {
        sc.log("Cancel spoofed request");
        return;
    }

    // move all current bets to the locked_bets array
    let state = sc.state();
    let bets = state.get_bytes_array(KEY_BETS);
    let locked_bets = state.get_bytes_array(KEY_LOCKED_BETS);
    let nrBets = bets.length();
    for i in 0..nrBets {
        let bytes = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bytes);
    }
    bets.clear();

    sc.post("payWinners").post(0);
}

fn payWinners(sc: &ScCallContext) {
    // can only be sent by SC itself
    let sc_id = sc.contract().id();
    if !sc.from(&sc_id) {
        sc.log("Cancel spoofed request");
        return;
    }

    let winning_color = sc.utility().random(5) + 1;
    let state = sc.state();
    state.get_int(KEY_LAST_WINNING_COLOR).set_value(winning_color);

    // gather all winners and calculate some totals
    let mut total_bet_amount = 0_i64;
    let mut total_win_amount = 0_i64;
    let locked_bets = state.get_bytes_array(KEY_LOCKED_BETS);
    let mut winners: Vec<BetInfo> = Vec::new();
    let nrBets = locked_bets.length();
    for i in 0..nrBets {
        let bytes = locked_bets.get_bytes(i).value();
        let bet = decodeBetInfo(&bytes);
        total_bet_amount += bet.amount;
        if bet.color == winning_color {
            total_win_amount += bet.amount;
            winners.push(bet);
        }
    }
    locked_bets.clear();

    if winners.is_empty() {
        sc.log("Nobody wins!");
        // compact separate UTXOs into a single one
        sc.transfer(&sc_id, &ScColor::IOTA, total_bet_amount);
        return;
    }

    // pay out the winners proportionally to their bet amount
    let mut total_payout = 0_i64;
    for i in 0..winners.len() {
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

fn playPeriod(sc: &ScCallContext) {
    // can only be sent by SC owner
    if !sc.from(&sc.contract().owner()) {
        sc.log("Cancel spoofed request");
        return;
    }

    let play_period = sc.params().get_int(KEY_PLAY_PERIOD).value();
    if play_period < 10 {
        sc.log("Invalid play period...");
        return;
    }

    sc.state().get_int(KEY_PLAY_PERIOD).set_value(play_period);
}

fn decodeBetInfo(bytes: &[u8]) -> BetInfo {
    let mut decoder = BytesDecoder::new(bytes);
    BetInfo {
        better: decoder.agent(),
        amount: decoder.int(),
        color: decoder.int(),
    }
}

fn encodeBetInfo(bet: &BetInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.agent(&bet.better);
    encoder.int(bet.amount);
    encoder.int(bet.color);
    encoder.data()
}
