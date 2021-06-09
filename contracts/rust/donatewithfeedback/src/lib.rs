// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use donatewithfeedback::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::state::*;

mod consts;
mod keys;
mod state;
mod types;
mod donatewithfeedback;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DONATE, func_donate_thunk);
    exports.add_func(FUNC_WITHDRAW, func_withdraw_thunk);
    exports.add_view(VIEW_DONATION, view_donation_thunk);
    exports.add_view(VIEW_DONATION_INFO, view_donation_info_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct FuncDonateParams {
    pub feedback: ScImmutableString, // feedback for the person you donate to
}

pub struct FuncDonateContext {
    params: FuncDonateParams,
    state:  DonateWithFeedbackFuncState,
}

fn func_donate_thunk(ctx: &ScFuncContext) {
    ctx.log("donatewithfeedback.funcDonate");
    let p = ctx.params().map_id();
    let f = FuncDonateContext {
        params: FuncDonateParams {
            feedback: ScImmutableString::new(p, idx_map(IDX_PARAM_FEEDBACK)),
        },
        state: DonateWithFeedbackFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_donate(ctx, &f);
    ctx.log("donatewithfeedback.funcDonate ok");
}

pub struct FuncWithdrawParams {
    pub amount: ScImmutableInt64, // amount to withdraw
}

pub struct FuncWithdrawContext {
    params: FuncWithdrawParams,
    state:  DonateWithFeedbackFuncState,
}

fn func_withdraw_thunk(ctx: &ScFuncContext) {
    ctx.log("donatewithfeedback.funcWithdraw");
    // only SC creator can withdraw donated funds
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params().map_id();
    let f = FuncWithdrawContext {
        params: FuncWithdrawParams {
            amount: ScImmutableInt64::new(p, idx_map(IDX_PARAM_AMOUNT)),
        },
        state: DonateWithFeedbackFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_withdraw(ctx, &f);
    ctx.log("donatewithfeedback.funcWithdraw ok");
}

pub struct ViewDonationParams {
    pub nr: ScImmutableInt64,
}

pub struct ViewDonationResults {
    pub amount:    ScMutableInt64,   // amount donated
    pub donator:   ScMutableAgentId, // who donated
    pub error:     ScMutableString,  // error to be reported to donator if anything goes wrong
    pub feedback:  ScMutableString,  // the feedback for the person donated to
    pub timestamp: ScMutableInt64,   // when the donation took place
}

pub struct ViewDonationContext {
    params:  ViewDonationParams,
    results: ViewDonationResults,
    state:   DonateWithFeedbackViewState,
}

fn view_donation_thunk(ctx: &ScViewContext) {
    ctx.log("donatewithfeedback.viewDonation");
    let p = ctx.params().map_id();
    let r = ctx.results().map_id();
    let f = ViewDonationContext {
        params: ViewDonationParams {
            nr: ScImmutableInt64::new(p, idx_map(IDX_PARAM_NR)),
        },
        results: ViewDonationResults {
            amount:    ScMutableInt64::new(r, idx_map(IDX_RESULT_AMOUNT)),
            donator:   ScMutableAgentId::new(r, idx_map(IDX_RESULT_DONATOR)),
            error:     ScMutableString::new(r, idx_map(IDX_RESULT_ERROR)),
            feedback:  ScMutableString::new(r, idx_map(IDX_RESULT_FEEDBACK)),
            timestamp: ScMutableInt64::new(r, idx_map(IDX_RESULT_TIMESTAMP)),
        },
        state: DonateWithFeedbackViewState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.nr.exists(), "missing mandatory nr");
    view_donation(ctx, &f);
    ctx.log("donatewithfeedback.viewDonation ok");
}

pub struct ViewDonationInfoResults {
    pub count:          ScMutableInt64,
    pub max_donation:   ScMutableInt64,
    pub total_donation: ScMutableInt64,
}

pub struct ViewDonationInfoContext {
    results: ViewDonationInfoResults,
    state:   DonateWithFeedbackViewState,
}

fn view_donation_info_thunk(ctx: &ScViewContext) {
    ctx.log("donatewithfeedback.viewDonationInfo");
    let r = ctx.results().map_id();
    let f = ViewDonationInfoContext {
        results: ViewDonationInfoResults {
            count:          ScMutableInt64::new(r, idx_map(IDX_RESULT_COUNT)),
            max_donation:   ScMutableInt64::new(r, idx_map(IDX_RESULT_MAX_DONATION)),
            total_donation: ScMutableInt64::new(r, idx_map(IDX_RESULT_TOTAL_DONATION)),
        },
        state: DonateWithFeedbackViewState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_donation_info(ctx, &f);
    ctx.log("donatewithfeedback.viewDonationInfo ok");
}

//@formatter:on
