// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

struct DonationInfo {
    seq: i64,
    id: ScRequestId,
    amount: i64,
    sender: ScAgent,
    error: String,
    feedback: String,
}

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("donate", donate);
    exports.add_call("withdraw", withdraw);
    exports.add_view("viewDonations", viewDonations);
}

fn donate(sc: &ScCallContext) {
    let tlog = sc.timestamped_log("l");
    let request = sc.request();
    let mut donation = DonationInfo {
        seq: tlog.length() as i64,
        id: request.id(),
        amount: request.balance(&ScColor::IOTA),
        sender: request.sender(),
        error: String::new(),
        feedback: request.params().get_string("f").value(),
    };
    if donation.amount == 0 || donation.feedback.len() == 0 {
        donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)".to_string();
        if donation.amount > 0 {
            sc.transfer(&donation.sender, &ScColor::IOTA, donation.amount);
            donation.amount = 0;
        }
    }
    let bytes = encodeDonationInfo(&donation);
    tlog.append(request.timestamp(), &bytes);

    let state = sc.state();
    let largest_donation = state.get_int("maxd");
    let total_donated = state.get_int("total");
    if donation.amount > largest_donation.value() { largest_donation.set_value(donation.amount); }
    total_donated.set_value(total_donated.value() + donation.amount);
}

fn withdraw(sc: &ScCallContext) {
    let sc_owner = sc.contract().owner();
    let request = sc.request();
    if !request.from(&sc_owner) {
        sc.log("Cancel spoofed request");
        return;
    }

    let account = sc.account();
    let amount = account.balance(&ScColor::IOTA);
    let mut withdraw_amount = request.params().get_int("s").value();
    if withdraw_amount == 0 || withdraw_amount > amount {
        withdraw_amount = amount;
    }
    if withdraw_amount == 0 {
        sc.log("DonateWithFeedback: withdraw. nothing to withdraw");
        return;
    }

    sc.transfer(&sc_owner, &ScColor::IOTA, withdraw_amount);
}

fn viewDonations(sc: &ScViewContext) {
    let state = sc.state();
    let largestDonation = state.get_int("maxd");
    let totalDonated = state.get_int("total");
    let tlog = sc.timestamped_log("l");
    let results = sc.results();
    results.get_int("largest").set_value(largestDonation.value());
    results.get_int("total").set_value(totalDonated.value());
    let donations = results.get_map_array("donations");
    let size = tlog.length();
    for i in 0..size {
        let log = tlog.get_map(i);
        let donation = donations.get_map(i);
        donation.get_int("timestamp").set_value(log.get_int("timestamp").value());
        let bytes = log.get_bytes("data").value();
        let di = decodeDonationInfo(&bytes);
        donation.get_int("amount").set_value(di.amount);
        donation.get_string("feedback").set_value(&di.feedback);
        donation.get_string("sender").set_value(&di.sender.to_string());
        donation.get_string("error").set_value(&di.error);
    }
}

fn decodeDonationInfo(bytes: &[u8]) -> DonationInfo {
    let mut decoder = BytesDecoder::new(bytes);
    DonationInfo {
        seq: decoder.int(),
        id: decoder.request_id(),
        amount: decoder.int(),
        sender: decoder.agent(),
        error: decoder.string(),
        feedback: decoder.string(),
    }
}

fn encodeDonationInfo(donation: &DonationInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(donation.seq);
    encoder.request_id(&donation.id);
    encoder.int(donation.amount);
    encoder.agent(&donation.sender);
    encoder.string(&donation.error);
    encoder.string(&donation.feedback);
    encoder.data()
}
