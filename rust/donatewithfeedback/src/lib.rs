#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScContext;

struct DonationInfo {
    seq: i64,
    id: String,
    amount: i64,
    sender: String,
    feedback: String,
    error: String,
}

#[no_mangle]
pub fn donate() {
    let ctx = ScContext::new();
    let tlog = ctx.timestamped_log("l");
    let request = ctx.request();
    let mut di = DonationInfo {
        seq: tlog.length() as i64,
        id: request.id(),
        amount: request.balance("iota"),
        sender: request.address(),
        feedback: request.params().get_string("f").value(),
        error: String::new(),
    };
    if di.amount == 0 || di.feedback.len() == 0 {
        di.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)".to_string();
        if di.amount > 0 {
            ctx.transfer(&di.sender, "iota", di.amount);
            di.amount = 0;
        }
    }
    let data = encodeDonationInfo(&di);
    tlog.append(request.timestamp(), &data);

    let state = ctx.state();
    let maxd = state.get_int("maxd");
    let total = state.get_int("total");
    if di.amount > maxd.value() { maxd.set_value(di.amount); }
    total.set_value(total.value() + di.amount);
}

#[no_mangle]
pub fn withdraw() {
    let ctx = ScContext::new();
    let owner = ctx.contract().owner();
    let request = ctx.request();
    if request.address() != owner {
        ctx.log("Cancel spoofed request");
        return;
    }

    let account = ctx.account();
    let bal = account.balance("iota");
    let mut withdrawSum = request.params().get_int("s").value();
    if withdrawSum == 0 || withdrawSum > bal {
        withdrawSum = bal;
    }
    if withdrawSum == 0 {
        ctx.log("DonateWithFeedback: withdraw. nothing to withdraw");
        return;
    }

    ctx.transfer(&owner, "iota", withdrawSum);
}

fn decodeDonationInfo(data: &[u8]) -> DonationInfo {
    let mut decoder = BytesDecoder::new(data);
    DonationInfo {
        seq: decoder.int(),
        id: decoder.string(),
        amount: decoder.int(),
        sender: decoder.string(),
        feedback: decoder.string(),
        error: decoder.string(),
    }
}

fn encodeDonationInfo(data: &DonationInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(data.seq);
    encoder.string(&data.id);
    encoder.int(data.amount);
    encoder.string(&data.feedback);
    encoder.string(&data.error);
    encoder.data()
}
