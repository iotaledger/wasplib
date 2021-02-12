// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_donate(ctx: &ScFuncContext, params: &FuncDonateParams) {
    let mut donation = Donation {
        amount: ctx.incoming().balance(&ScColor::IOTA),
        donator: ctx.caller(),
        error: String::new(),
        feedback: params.feedback.value(),
        timestamp: ctx.timestamp(),
    };
    if donation.amount == 0 || donation.feedback.len() == 0 {
        donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)".to_string();
        if donation.amount > 0 {
            ctx.transfer_to_address(&donation.donator.address(), &ScTransfers::new(&ScColor::IOTA, donation.amount));
            donation.amount = 0;
        }
    }
    let state = ctx.state();
    let log = state.get_bytes_array(VAR_LOG);
    log.get_bytes(log.length()).set_value(&donation.to_bytes());

    let largest_donation = state.get_int(VAR_MAX_DONATION);
    let total_donated = state.get_int(VAR_TOTAL_DONATION);
    if donation.amount > largest_donation.value() {
        largest_donation.set_value(donation.amount);
    }
    total_donated.set_value(total_donated.value() + donation.amount);
}

pub fn func_withdraw(ctx: &ScFuncContext, params: &FuncWithdrawParams) {
    let balance = ctx.balances().balance(&ScColor::IOTA);
    let mut amount = params.amount.value();
    if amount == 0 || amount > balance {
        amount = balance;
    }
    if amount == 0 {
        ctx.log("DonateWithFeedback: nothing to withdraw");
        return;
    }

    let sc_creator = ctx.contract_creator().address();
    ctx.transfer_to_address(&sc_creator, &ScTransfers::new(&ScColor::IOTA, amount));
}

pub fn view_donations(ctx: &ScViewContext, _params: &ViewDonationsParams) {
    let state = ctx.state();
    let largest_donation = state.get_int(VAR_MAX_DONATION);
    let total_donated = state.get_int(VAR_TOTAL_DONATION);
    let log = state.get_bytes_array(VAR_LOG);
    let results = ctx.results();
    results.get_int(VAR_MAX_DONATION).set_value(largest_donation.value());
    results.get_int(VAR_TOTAL_DONATION).set_value(total_donated.value());
    let donations = results.get_map_array(VAR_DONATIONS);
    let size = log.length();
    for i in 0..size {
        let di = Donation::from_bytes(&log.get_bytes(i).value());
        let donation = donations.get_map(i);
        donation.get_int(VAR_AMOUNT).set_value(di.amount);
        donation.get_string(VAR_DONATOR).set_value(&di.donator.to_string());
        donation.get_string(VAR_ERROR).set_value(&di.error);
        donation.get_string(VAR_FEEDBACK).set_value(&di.feedback);
        donation.get_int(VAR_TIMESTAMP).set_value(di.timestamp);
    }
}
