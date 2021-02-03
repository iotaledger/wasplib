// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;
use super::*;

pub const SC_NAME: &str = "fairroulette";
pub const SC_HNAME: ScHname = ScHname(0xdf79d138);

pub const PARAM_NUMBER: &str = "number";
pub const PARAM_PLAY_PERIOD: &str = "playPeriod";

pub const VAR_BETS: &str = "bets";
pub const VAR_LAST_WINNING_NUMBER: &str = "lastWinningNumber";
pub const VAR_LOCKED_BETS: &str = "lockedBets";
pub const VAR_PLAY_PERIOD: &str = "playPeriod";

pub const FUNC_LOCK_BETS: &str = "lockBets";
pub const FUNC_PAY_WINNERS: &str = "payWinners";
pub const FUNC_PLACE_BET: &str = "placeBet";
pub const FUNC_PLAY_PERIOD: &str = "playPeriod";

pub const HFUNC_LOCK_BETS: ScHname = ScHname(0xe163b43c);
pub const HFUNC_PAY_WINNERS: ScHname = ScHname(0xfb2b0144);
pub const HFUNC_PLACE_BET: ScHname = ScHname(0xdfba7d1b);
pub const HFUNC_PLAY_PERIOD: ScHname = ScHname(0xcb94b293);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_LOCK_BETS, func_lock_bets);
    exports.add_call(FUNC_PAY_WINNERS, func_pay_winners);
    exports.add_call(FUNC_PLACE_BET, func_place_bet);
    exports.add_call(FUNC_PLAY_PERIOD, func_play_period);
}
