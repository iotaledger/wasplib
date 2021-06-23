// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_donate(ctx: &ScFuncContext, f: &DonateContext) {
    let mut donation = Donation {
        amount: ctx.incoming().balance(&ScColor::IOTA),
        donator: ctx.caller(),
        error: String::new(),
        feedback: f.params.feedback().value(),
        timestamp: ctx.timestamp(),
    };
    if donation.amount == 0 || donation.feedback.len() == 0 {
        donation.error = "error: empty feedback or donated amount = 0".to_string();
        if donation.amount > 0 {
            ctx.transfer_to_address(&donation.donator.address(), ScTransfers::iotas(donation.amount));
            donation.amount = 0;
        }
    }
    let log = f.state.log();
    log.get_donation(log.length()).set_value(&donation);

    let largest_donation = f.state.max_donation();
    let total_donated = f.state.total_donation();
    if donation.amount > largest_donation.value() {
        largest_donation.set_value(donation.amount);
    }
    total_donated.set_value(total_donated.value() + donation.amount);
}

pub fn func_withdraw(ctx: &ScFuncContext, f: &WithdrawContext) {
    let balance = ctx.balances().balance(&ScColor::IOTA);
    let mut amount = f.params.amount().value();
    if amount == 0 || amount > balance {
        amount = balance;
    }
    if amount == 0 {
        ctx.log("dwf.withdraw: nothing to withdraw");
        return;
    }

    let sc_creator = ctx.contract_creator().address();
    ctx.transfer_to_address(&sc_creator, ScTransfers::iotas(amount));
}

pub fn view_donation(_ctx: &ScViewContext, f: &DonationContext) {
    let nr = (f.params.nr().value()) as i32;
    let donation = f.state.log().get_donation(nr).value();
    f.results.amount().set_value(donation.amount);
    f.results.donator().set_value(&donation.donator);
    f.results.error().set_value(&donation.error);
    f.results.feedback().set_value(&donation.feedback);
    f.results.timestamp().set_value(donation.timestamp);
}

pub fn view_donation_info(_ctx: &ScViewContext, f: &DonationInfoContext) {
    f.results.max_donation().set_value(f.state.max_donation().value());
    f.results.total_donation().set_value(f.state.total_donation().value());
    f.results.count().set_value(f.state.log().length() as i64);
}
