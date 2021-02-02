// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;

use super::*;

pub const SC_NAME: &str = "fairroulette";
pub const SC_HNAME: Hname = Hname(0xdf79d138);

pub const PARAM_COLOR: &str = "color";
pub const PARAM_PLAY_PERIOD: &str = "play_period";

pub const VAR_BETS: &str = "bets";
pub const VAR_LAST_WINNING_COLOR: &str = "last_winning_color";
pub const VAR_LOCKED_BETS: &str = "locked_bets";
pub const VAR_PLAY_PERIOD: &str = "play_period";

pub const FUNC_LOCK_BETS: &str = "lock_bets";
pub const FUNC_PAY_WINNERS: &str = "pay_winners";
pub const FUNC_PLACE_BET: &str = "place_bet";
pub const FUNC_PLAY_PERIOD: &str = "play_period";

pub const HFUNC_LOCK_BETS: Hname = Hname(0x853da2a7);
pub const HFUNC_PAY_WINNERS: Hname = Hname(0x3df139de);
pub const HFUNC_PLACE_BET: Hname = Hname(0x575b51d2);
pub const HFUNC_PLAY_PERIOD: Hname = Hname(0xf534dac1);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_LOCK_BETS, func_lock_bets);
    exports.add_call(FUNC_PAY_WINNERS, func_pay_winners);
    exports.add_call(FUNC_PLACE_BET, func_place_bet);
    exports.add_call(FUNC_PLAY_PERIOD, func_play_period);
}
