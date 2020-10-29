#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

const NUM_COLORS: i64 = 5;
const PLAY_PERIOD: i64 = 120;

struct BetInfo {
    id: ScRequestId,
    sender: ScAddress,
    amount: i64,
    color: i64,
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
    let amount = request.balance(&ScColor::IOTA);
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
        amount,
        color,
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
        sc.post_request(&sc.contract().address(), "lockBets", play_period);
    }
}

#[no_mangle]
pub fn lockBets() {
    // can only be sent by SC itself
    let sc = ScContext::new();
    let sc_address = sc.contract().address();
    if !sc.request().from(&sc_address) {
        sc.log("Cancel spoofed request");
        return;
    }

    // move all current bets to the locked_bets array
    let state = sc.state();
    let bets = state.get_bytes_array("bets");
    let locked_bets = state.get_bytes_array("lockedBets");
    let nrBets = bets.length();
    for i in 0..nrBets {
        let bytes = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bytes);
    }
    bets.clear();

    sc.post_request(&sc_address, "payWinners", 0);
}

#[no_mangle]
pub fn payWinners() {
    // can only be sent by SC itself
    let sc = ScContext::new();
    let sc_address = sc.contract().address();
    if !sc.request().from(&sc_address) {
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
        sc.transfer(&sc_address, &ScColor::IOTA, total_bet_amount);
        return;
    }

    // pay out the winners proportionally to their bet amount
    let mut total_payout = 0_i64;
    for i in 0..winners.len() {
        let bet = &winners[i];
        let payout = total_bet_amount * bet.amount / total_win_amount;
        if payout != 0 {
            total_payout += payout;
            sc.transfer(&bet.sender, &ScColor::IOTA, payout);
        }
        let text = "Pay ".to_string() + &payout.to_string() + " to " + &bet.sender.to_string();
        sc.log(&text);
    }

    // any truncation left-overs are fair picking for the smart contract
    if total_payout != total_bet_amount {
        let remainder = total_bet_amount - total_payout;
        let text = "Remainder is ".to_string() + &remainder.to_string();
        sc.log(&text);
        sc.transfer(&sc_address, &ScColor::IOTA, remainder);
    }
}

#[no_mangle]
pub fn playPeriod() {
    // can only be sent by SC owner
    let sc = ScContext::new();
    let request = sc.request();
    if !request.from(&sc.contract().owner()) {
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
        id: decoder.request_id(),
        sender: decoder.address(),
        amount: decoder.int(),
        color: decoder.int(),
    }
}

fn encodeBetInfo(bet: &BetInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.request_id(&bet.id);
    encoder.address(&bet.sender);
    encoder.int(bet.amount);
    encoder.int(bet.color);
    encoder.data()
}
