// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const KEY_AUCTIONS: &str = "auctions";
const KEY_BIDDERS: &str = "bidders";
const KEY_BIDDER_LIST: &str = "bidder_list";
const KEY_COLOR: &str = "color";
const KEY_DESCRIPTION: &str = "description";
const KEY_DURATION: &str = "duration";
const KEY_INFO: &str = "info";
const KEY_MINIMUM_BID: &str = "minimum";
const KEY_OWNER_MARGIN: &str = "owner_margin";

const DURATION_DEFAULT: i64 = 60;
const DURATION_MIN: i64 = 1;
const DURATION_MAX: i64 = 120;
const MAX_DESCRIPTION_LENGTH: usize = 150;
const OWNER_MARGIN_DEFAULT: i64 = 50;
const OWNER_MARGIN_MIN: i64 = 5;
const OWNER_MARGIN_MAX: i64 = 100;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("start_auction", start_auction);
    exports.add_call("finalize_auction", finalize_auction);
    exports.add_call("place_bid", place_bid);
    exports.add_call("set_owner_margin", set_owner_margin);
}

fn start_auction(sc: &ScCallContext) {
    let deposit = sc.incoming().balance(&ScColor::IOTA);
    if deposit < 1 {
        sc.log("Empty deposit...");
        return;
    }

    let state = sc.state();
    let mut owner_margin = state.get_int(KEY_OWNER_MARGIN).value();
    if owner_margin == 0 {
        owner_margin = OWNER_MARGIN_DEFAULT;
    }

    let params = sc.params();
    let color_param = params.get_color(KEY_COLOR);
    if !color_param.exists() {
        refund(sc, deposit / 2, "Missing token color...");
        return;
    }
    let color = color_param.value();

    if color == ScColor::IOTA || color == ScColor::MINT {
        refund(sc, deposit / 2, "Reserved token color...");
        return;
    }

    let num_tokens = sc.incoming().balance(&color);
    if num_tokens == 0 {
        refund(sc, deposit / 2, "Auction tokens missing from request...");
        return;
    }

    let minimum_bid = params.get_int(KEY_MINIMUM_BID).value();
    if minimum_bid == 0 {
        refund(sc, deposit / 2, "Missing minimum bid...");
        return;
    }

    // need at least 1 iota to run SC
    let mut margin = minimum_bid * owner_margin / 1000;
    if margin == 0 {
        margin = 1;
    }
    if deposit < margin {
        refund(sc, deposit / 2, "Insufficient deposit...");
        return;
    }

    // duration in minutes
    let mut duration = params.get_int(KEY_DURATION).value();
    if duration == 0 {
        duration = DURATION_DEFAULT;
    }
    if duration < DURATION_MIN {
        duration = DURATION_MIN;
    }
    if duration > DURATION_MAX {
        duration = DURATION_MAX;
    }

    let mut description = params.get_string(KEY_DESCRIPTION).value();
    if description == "" {
        description = "N/A".to_string()
    }
    if description.len() > MAX_DESCRIPTION_LENGTH {
        let ss: String = description.chars().take(MAX_DESCRIPTION_LENGTH).collect();
        description = ss + "[...]";
    }

    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let current_info = current_auction.get_bytes(KEY_INFO);
    if current_info.exists() {
        refund(sc, deposit / 2, "Auction for this token already exists...");
        return;
    }

    let auction = &AuctionInfo {
        auction_owner: sc.caller(),
        color: color,
        deposit: deposit,
        description: description,
        duration: duration,
        highest_bid: -1,
        highest_bidder: ScAgent::NONE,
        minimum_bid: minimum_bid,
        num_tokens: num_tokens,
        owner_margin: owner_margin,
        when_started: sc.timestamp(),
    };
    current_info.set_value(&encode_auction_info(auction));

    let finalize_request = sc.post("finalize_auction");
    let finalize_params = finalize_request.params();
    finalize_params.get_color(KEY_COLOR).set_value(&auction.color);
    finalize_request.post(duration * 60);
    sc.log("New auction started...");
}

fn finalize_auction(sc: &ScCallContext) {
    // can only be sent by SC itself
    if !sc.from(&sc.contract().id()) {
        sc.log("Cancel spoofed request");
        return;
    }

    let color_param = sc.params().get_color(KEY_COLOR);
    if !color_param.exists() {
        sc.log("Internal inconsistency: missing color");
        return;
    }
    let color = color_param.value();

    let state = sc.state();
    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let current_info = current_auction.get_bytes(KEY_INFO);
    if !current_info.exists() {
        sc.log("Internal inconsistency: missing auction info");
        return;
    }
    let auction = decode_auction_info(&current_info.value());
    if auction.highest_bid < 0 {
        sc.log(&("No one bid on ".to_string() + &color.to_string()));
        let mut owner_fee = auction.minimum_bid * auction.owner_margin / 1000;
        if owner_fee == 0 {
            owner_fee = 1
        }
        // finalizeAuction request token was probably not confirmed yet
        sc.transfer(&sc.contract().creator(), &ScColor::IOTA, owner_fee - 1);
        sc.transfer(&auction.auction_owner, &auction.color, auction.num_tokens);
        sc.transfer(&auction.auction_owner, &ScColor::IOTA, auction.deposit - owner_fee);
        return;
    }

    let mut owner_fee = auction.highest_bid * auction.owner_margin / 1000;
    if owner_fee == 0 {
        owner_fee = 1;
    }

    // return staked bids to losers
    let bidders = current_auction.get_map(KEY_BIDDERS);
    let bidder_list = current_auction.get_agent_array(KEY_BIDDER_LIST);
    let size = bidder_list.length();
    for i in 0..size {
        let bidder = bidder_list.get_agent(i).value();
        if bidder != auction.highest_bidder {
            let loser = bidders.get_bytes(&bidder);
            let bid = decode_bid_info(&loser.value());
            sc.transfer(&bidder, &ScColor::IOTA, bid.amount);
        }
    }

    // finalizeAuction request token was probably not confirmed yet
    sc.transfer(&sc.contract().creator(), &ScColor::IOTA, owner_fee - 1);
    sc.transfer(&auction.highest_bidder, &auction.color, auction.num_tokens);
    sc.transfer(&auction.auction_owner, &ScColor::IOTA, auction.deposit + auction.highest_bid - owner_fee);
}

fn place_bid(sc: &ScCallContext) {
    let mut bid_amount = sc.incoming().balance(&ScColor::IOTA);
    if bid_amount == 0 {
        sc.log("Insufficient bid amount");
        return;
    }

    let color_param = sc.params().get_color(KEY_COLOR);
    if !color_param.exists() {
        refund(sc, bid_amount, "Missing token color");
        return;
    }
    let color = color_param.value();

    let state = sc.state();
    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let current_info = current_auction.get_bytes(KEY_INFO);
    if !current_info.exists() {
        refund(sc, bid_amount, "Missing auction");
        return;
    }

    let mut auction = decode_auction_info(&current_info.value());
    let bidders = current_auction.get_map(KEY_BIDDERS);
    let bidder_list = current_auction.get_agent_array(KEY_BIDDER_LIST);
    let caller = sc.caller();
    let bidder = bidders.get_bytes(&caller);
    if bidder.exists() {
        sc.log(&("Upped bid from: ".to_string() + &caller.to_string()));
        let mut bid = decode_bid_info(&bidder.value());
        bid_amount += bid.amount;
        bid.amount = bid_amount;
        bid.timestamp = sc.timestamp();
        bidder.set_value(&encode_bid_info(&bid));
    } else {
        sc.log(&("New bid from: ".to_string() + &caller.to_string()));
        let index = bidder_list.length();
        bidder_list.get_agent(index).set_value(&caller);
        let bid = BidInfo {
            index: index as i64,
            amount: bid_amount,
            timestamp: sc.timestamp(),
        };
        bidder.set_value(&encode_bid_info(&bid));
    }
    if bid_amount > auction.highest_bid {
        sc.log("New highest bidder...");
        auction.highest_bid = bid_amount;
        auction.highest_bidder = caller;
        current_info.set_value(&encode_auction_info(&auction));
    }
}

fn set_owner_margin(sc: &ScCallContext) {
    // can only be sent by SC owner
    if !sc.from(&sc.contract().creator()) {
        sc.log("Cancel spoofed request");
        return;
    }

    let mut owner_margin = sc.params().get_int(KEY_OWNER_MARGIN).value();
    if owner_margin < OWNER_MARGIN_MIN {
        owner_margin = OWNER_MARGIN_MIN;
    }
    if owner_margin > OWNER_MARGIN_MAX {
        owner_margin = OWNER_MARGIN_MAX;
    }
    sc.state().get_int(KEY_OWNER_MARGIN).set_value(owner_margin);
    sc.log("Updated owner margin...");
}

fn refund(sc: &ScCallContext, amount: i64, reason: &str) {
    sc.log(reason);
    let caller = sc.caller();
    if amount != 0 {
        sc.transfer(&caller, &ScColor::IOTA, amount);
    }
    let incoming = sc.incoming();
    let deposit = incoming.balance(&ScColor::IOTA);
    if deposit - amount != 0 {
        sc.transfer(&sc.contract().creator(), &ScColor::IOTA, deposit - amount);
    }

    // refund all other token colors, don't keep tokens that were to be auctioned
    let colors = incoming.colors();
    let size = colors.length();
    for i in 0..size {
        let color = colors.get_color(i).value();
        if color != ScColor::IOTA {
            sc.transfer(&caller, &color, incoming.balance(&color));
        }
    }
}
