#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

struct DonationInfo {
    seq: i64,
    id: ScRequestId,
    amount: i64,
    sender: ScAddress,
    error: String,
    feedback: String,
}

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("donate");
    exports.add_protected("withdraw");
}

#[no_mangle]
pub fn donate() {
    let sc = ScContext::new();
    let tlog = sc.timestamped_log("l");
    let request = sc.request();
    let mut donation = DonationInfo {
        seq: tlog.length() as i64,
        id: request.id(),
        amount: request.balance(&ScColor::IOTA),
        sender: request.address(),
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

#[no_mangle]
pub fn withdraw() {
    let sc = ScContext::new();
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

fn decodeDonationInfo(bytes: &[u8]) -> DonationInfo {
    let mut decoder = BytesDecoder::new(bytes);
    DonationInfo {
        seq: decoder.int(),
        id: decoder.request_id(),
        amount: decoder.int(),
        sender: decoder.address(),
        error: decoder.string(),
        feedback: decoder.string(),
    }
}

fn encodeDonationInfo(donation: &DonationInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(donation.seq);
    encoder.request_id(&donation.id);
    encoder.int(donation.amount);
    encoder.address(&donation.sender);
    encoder.string(&donation.error);
    encoder.string(&donation.feedback);
    encoder.data()
}
