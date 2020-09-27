#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScContext;

const NUM_COLORS: i64 = 5;
const PLAY_PERIOD: i64 = 120;

struct BetInfo {
    id: String,
    sender: String,
    color: i64,
    amount: i64,
}

#[no_mangle]
pub fn placeBet() {
    let ctx = ScContext::new();
    let request = ctx.request();
    let amount = request.balance("iota");
    if amount == 0 {
        ctx.log("Empty bet...");
        return;
    }
    let color = request.params().get_int("color").value();
    if color == 0 {
        ctx.log("No color...");
        return;
    }
    if color < 1 || color > NUM_COLORS {
        ctx.log("Invalid color...");
        return;
    }

    let bet = BetInfo {
        id: request.id(),
        sender: request.address(),
        color,
        amount,
    };

    let state = ctx.state();
    let bets = state.get_bytes_array("bets");
    let bet_nr = bets.length();
    let bet_data = encodeBetInfo(&bet);
    bets.get_bytes(bet_nr).set_value(&bet_data);
    if bet_nr == 0 {
        let mut play_period = state.get_int("playPeriod").value();
        if play_period < 10 {
            play_period = PLAY_PERIOD;
        }
        ctx.event("", "lockBets", play_period);
    }
}

#[no_mangle]
pub fn lockBets() {
    // can only be sent by SC itself
    let ctx = ScContext::new();
    if ctx.request().address() != ctx.contract().address() {
        ctx.log("Cancel spoofed request");
        return;
    }

    // move all current bets to the locked_bets array
    let state = ctx.state();
    let bets = state.get_bytes_array("bets");
    let locked_bets = state.get_bytes_array("lockedBets");
    for i in 0..bets.length() {
        let bet = bets.get_bytes(i).value();
        locked_bets.get_bytes(i).set_value(&bet);
    }
    bets.clear();

    ctx.event("", "payWinners", 0);
}

#[no_mangle]
pub fn payWinners() {
    // can only be sent by SC itself
    let ctx = ScContext::new();
    let sc_address = ctx.contract().address();
    if ctx.request().address() != sc_address {
        ctx.log("Cancel spoofed request");
        return;
    }

    let winning_color = ctx.utility().random(5) + 1;
    let state = ctx.state();
    state.get_int("lastWinningColor").set_value(winning_color);

    // gather all winners and calculate some totals
    let mut total_bet_amount: i64 = 0;
    let mut total_win_amount: i64 = 0;
    let locked_bets = state.get_bytes_array("lockedBets");
    let mut winners: Vec<BetInfo> = Vec::new();
    for i in 0..locked_bets.length() {
        let data = locked_bets.get_bytes(i).value();
        let bet = decodeBetInfo(&data);
        total_bet_amount += bet.amount;
        if bet.color == winning_color {
            total_win_amount += bet.amount;
            winners.push(bet);
        }
    }
    locked_bets.clear();

    if winners.is_empty() {
        ctx.log("Nobody wins!");
        // compact separate UTXOs into a single one
        ctx.transfer(&sc_address, "iota", total_bet_amount);
        return;
    }

    // pay out the winners proportionally to their bet amount
    let mut total_payout: i64 = 0;
    for i in 0..winners.len() {
        let bet = &winners[i];
        let payout = total_bet_amount * bet.amount / total_win_amount;
        if payout != 0 {
            total_payout += payout;
            ctx.transfer(&bet.sender, "iota", payout);
        }
        let text = "Pay ".to_string() + &payout.to_string() + " to " + &bet.sender;
        ctx.log(&text);
    }

    // any truncation left-overs are fair picking for the smart contract
    if total_payout != total_bet_amount {
        let remainder = total_bet_amount - total_payout;
        let text = "Remainder is ".to_string() + &remainder.to_string();
        ctx.log(&text);
        ctx.transfer(&sc_address, "iota", remainder);
    }
}

#[no_mangle]
pub fn playPeriod() {
    // can only be sent by SC owner
    let ctx = ScContext::new();
    let request = ctx.request();
    if request.address() != ctx.contract().owner() {
        ctx.log("Cancel spoofed request");
        return;
    }

    let play_period = request.params().get_int("playPeriod").value();
    if play_period < 10 {
        ctx.log("Invalid play period...");
        return;
    }

    ctx.state().get_int("playPeriod").set_value(play_period);
}

fn decodeBetInfo(data: &[u8]) -> BetInfo {
    let mut decoder = BytesDecoder::new(data);
    BetInfo {
        id: decoder.string(),
        sender: decoder.string(),
        amount: decoder.int(),
        color: decoder.int(),
    }
}

fn encodeBetInfo(data: &BetInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.string(&data.id);
    encoder.string(&data.sender);
    encoder.int(data.amount);
    encoder.int(data.color);
    encoder.data()
}
