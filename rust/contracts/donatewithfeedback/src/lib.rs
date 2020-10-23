#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScColor;
use wasplib::client::ScContext;
use wasplib::client::ScExports;

struct DonationInfo {
    seq: i64,
    id: String,
    amount: i64,
    sender: String,
    feedback: String,
    error: String,
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
        amount: request.balance(&ScColor::iota()),
        sender: request.address(),
        feedback: request.params().get_string("f").value(),
        error: String::new(),
    };
    if donation.amount == 0 || donation.feedback.len() == 0 {
        donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)".to_string();
        if donation.amount > 0 {
            sc.transfer(&donation.sender, &ScColor::iota(), donation.amount);
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
    let owner = sc.contract().owner();
    let request = sc.request();
    if request.address() != owner {
        sc.log("Cancel spoofed request");
        return;
    }

    let account = sc.account();
    let amount = account.balance(&ScColor::iota());
    let mut withdraw_amount = request.params().get_int("s").value();
    if withdraw_amount == 0 || withdraw_amount > amount {
        withdraw_amount = amount;
    }
    if withdraw_amount == 0 {
        sc.log("DonateWithFeedback: withdraw. nothing to withdraw");
        return;
    }

    sc.transfer(&owner, &ScColor::iota(), withdraw_amount);
}

fn decodeDonationInfo(bytes: &[u8]) -> DonationInfo {
    let mut decoder = BytesDecoder::new(bytes);
    DonationInfo {
        seq: decoder.int(),
        id: decoder.string(),
        amount: decoder.int(),
        sender: decoder.string(),
        feedback: decoder.string(),
        error: decoder.string(),
    }
}

fn encodeDonationInfo(donation: &DonationInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(donation.seq);
    encoder.string(&donation.id);
    encoder.int(donation.amount);
    encoder.string(&donation.feedback);
    encoder.string(&donation.error);
    encoder.data()
}
