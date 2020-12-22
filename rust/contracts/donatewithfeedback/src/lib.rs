// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const KEY_AMOUNT: &str = "amount";
const KEY_DONATIONS: &str = "donations";
const KEY_DONATOR: &str = "donator";
const KEY_ERROR: &str = "error";
const KEY_FEEDBACK: &str = "feedback";
const KEY_LOG: &str = "log";
const KEY_MAX_DONATION: &str = "max_donation";
const KEY_TIMESTAMP: &str = "timestamp";
const KEY_TOTAL_DONATION: &str = "total_donation";
const KEY_WITHDRAW_AMOUNT: &str = "withdraw";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("donate", donate);
    exports.add_call("withdraw", withdraw);
    exports.add_view("view_donations", view_donations);
}

fn donate(sc: &ScCallContext) {
    let mut donation = DonationInfo {
        amount: sc.incoming().balance(&ScColor::IOTA),
        donator: sc.caller(),
        error: String::new(),
        feedback: sc.params().get_string(KEY_FEEDBACK).value(),
        timestamp: sc.timestamp(),
    };
    if donation.amount == 0 || donation.feedback.len() == 0 {
        donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)".to_string();
        if donation.amount > 0 {
            sc.transfer(&donation.donator, &ScColor::IOTA, donation.amount);
            donation.amount = 0;
        }
    }
    let state = sc.state();
    let log = state.get_bytes_array(KEY_LOG);
    log.get_bytes(log.length()).set_value(&encode_donation_info(&donation));

    let largest_donation = state.get_int(KEY_MAX_DONATION);
    let total_donated = state.get_int(KEY_TOTAL_DONATION);
    if donation.amount > largest_donation.value() {
        largest_donation.set_value(donation.amount);
    }
    total_donated.set_value(total_donated.value() + donation.amount);
}

fn withdraw(sc: &ScCallContext) {
    let sc_owner = sc.contract().owner();
    if !sc.from(&sc_owner) {
        sc.log("Cancel spoofed request");
        return;
    }

    let amount = sc.balances().balance(&ScColor::IOTA);
    let mut withdraw_amount = sc.params().get_int(KEY_WITHDRAW_AMOUNT).value();
    if withdraw_amount == 0 || withdraw_amount > amount {
        withdraw_amount = amount;
    }
    if withdraw_amount == 0 {
        sc.log("DonateWithFeedback: withdraw. nothing to withdraw");
        return;
    }

    sc.transfer(&sc_owner, &ScColor::IOTA, withdraw_amount);
}

fn view_donations(sc: &ScViewContext) {
    let state = sc.state();
    let largest_donation = state.get_int(KEY_MAX_DONATION);
    let total_donated = state.get_int(KEY_TOTAL_DONATION);
    let log = state.get_bytes_array(KEY_LOG);
    let results = sc.results();
    results.get_int(KEY_MAX_DONATION).set_value(largest_donation.value());
    results.get_int(KEY_TOTAL_DONATION).set_value(total_donated.value());
    let donations = results.get_map_array(KEY_DONATIONS);
    let size = log.length();
    for i in 0..size {
        let di = decode_donation_info(&log.get_bytes(i).value());
        let donation = donations.get_map(i);
        donation.get_int(KEY_AMOUNT).set_value(di.amount);
        donation.get_string(KEY_DONATOR).set_value(&di.donator.to_string());
        donation.get_string(KEY_ERROR).set_value(&di.error);
        donation.get_string(KEY_FEEDBACK).set_value(&di.feedback);
        donation.get_int(KEY_TIMESTAMP).set_value(di.timestamp);
    }
}
