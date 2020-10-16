#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScContext;
use wasplib::client::ScExports;

const NUM_COLORS: i64 = 5;
const PLAY_PERIOD: i64 = 120;

struct BetInfo {
    id: String,
    sender: String,
    color: i64,
    amount: i64,
}

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("placeBet");
    exports.add("lockBets");
    exports.add("payWinners");
    exports.add_protected("playPeriod");
    exports.add("nothing");
}

#[no_mangle]
pub fn placeBet() {
    let sc = ScContext::new();
    let request = sc.request();
    let amount = request.balance("iota");
    if amount == 0 {
        sc.log("Empty bet...");
        return;
    }
    let color = request.params().get_int("color").value();
    if color == 0 {
        sc.log("No color...");
        return;
    }
    if color < 1 || color > NUM_COLORS {
        sc.log("Invalid color...");
        return;
    }

    let bet = BetInfo {
        id: request.id(),
        sender: request.address(),
        color,
        amount,
    };

    let state = sc.state();
    let bets = state.get_bytes_array("bets");
    let bet_nr = bets.length();
    let bytes = encodeBetInfo(&bet);
    bets.get_bytes(bet_nr).set_value(&bytes);
    if bet_nr == 0 {
        let mut play_period = state.get_int("playPeriod").value();
        if play_period < 10 {
            play_period = PLAY_PERIOD;
        }
        sc.event("", "lockBets", play_period);
    }
}

#[no_mangle]
pub fn lockBets() {
    // can only be sent by SC itself
    let sc = ScContext::new();
    if sc.request().address() != sc.contract().address() {
        sc.log("Cancel spoofed request");
        return;
    }

    // move all current bets to the locked_bets array
    let state = sc.state();
    let bets = state.get_bytes_array("bets");
    let locked_bets = state.get_bytes_array("lockedBets");
    for i in 0..bets.length() {
        let bytes = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bytes);
    }
    bets.clear();

    sc.event("", "payWinners", 0);
}

#[no_mangle]
pub fn payWinners() {
    // can only be sent by SC itself
    let sc = ScContext::new();
    let sc_address = sc.contract().address();
    if sc.request().address() != sc_address {
        sc.log("Cancel spoofed request");
        return;
    }

    let winning_color = sc.utility().random(5) + 1;
    let state = sc.state();
    state.get_int("lastWinningColor").set_value(winning_color);

    // gather all winners and calculate some totals
    let mut total_bet_amount = 0_i64;
    let mut total_win_amount = 0_i64;
    let locked_bets = state.get_bytes_array("lockedBets");
    let mut winners: Vec<BetInfo> = Vec::new();
    for i in 0..locked_bets.length() {
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
        sc.transfer(&sc_address, "iota", total_bet_amount);
        return;
    }

    // pay out the winners proportionally to their bet amount
    let mut total_payout = 0_i64;
    for i in 0..winners.len() {
        let bet = &winners[i];
        let payout = total_bet_amount * bet.amount / total_win_amount;
        if payout != 0 {
            total_payout += payout;
            sc.transfer(&bet.sender, "iota", payout);
        }
        let text = "Pay ".to_string() + &payout.to_string() + " to " + &bet.sender;
        sc.log(&text);
    }

    // any truncation left-overs are fair picking for the smart contract
    if total_payout != total_bet_amount {
        let remainder = total_bet_amount - total_payout;
        let text = "Remainder is ".to_string() + &remainder.to_string();
        sc.log(&text);
        sc.transfer(&sc_address, "iota", remainder);
    }
}

#[no_mangle]
pub fn playPeriod() {
    // can only be sent by SC owner
    let sc = ScContext::new();
    let request = sc.request();
    if request.address() != sc.contract().owner() {
        sc.log("Cancel spoofed request");
        return;
    }

    let play_period = request.params().get_int("playPeriod").value();
    if play_period < 10 {
        sc.log("Invalid play period...");
        return;
    }

    sc.state().get_int("playPeriod").set_value(play_period);
}

fn decodeBetInfo(bytes: &[u8]) -> BetInfo {
    let mut decoder = BytesDecoder::new(bytes);
    BetInfo {
        id: decoder.string(),
        sender: decoder.string(),
        amount: decoder.int(),
        color: decoder.int(),
    }
}

fn encodeBetInfo(bet: &BetInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.string(&bet.id);
    encoder.string(&bet.sender);
    encoder.int(bet.amount);
    encoder.int(bet.color);
    encoder.data()
}
