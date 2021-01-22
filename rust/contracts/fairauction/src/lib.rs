// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const KEY_AUCTIONS: &str = "auctions";
const KEY_BIDDERS: &str = "bidders";
const KEY_BIDDER_LIST: &str = "bidder_list";
const KEY_COLOR: &str = "color";
const KEY_CREATOR: &str = "creator";
const KEY_DEPOSIT: &str = "deposit";
const KEY_DESCRIPTION: &str = "description";
const KEY_DURATION: &str = "duration";
const KEY_HIGHEST_BID: &str = "highest_bid";
const KEY_HIGHEST_BIDDER: &str = "highest_bidder";
const KEY_INFO: &str = "info";
const KEY_MINIMUM_BID: &str = "minimum";
const KEY_NUM_TOKENS: &str = "num_tokens";
const KEY_OWNER_MARGIN: &str = "owner_margin";
const KEY_WHEN_STARTED: &str = "when_started";

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
    exports.add_view("get_info", get_info);
}

fn start_auction(sc: &ScCallContext) {
    let params = sc.params();
    let color_param = params.get_color(KEY_COLOR);
    if !color_param.exists() {
        sc.panic("Missing auction token color");
    }
    let color = color_param.value();
    if color.equals(&ScColor::IOTA) || color.equals(&ScColor::MINT) {
        sc.panic("Reserved auction token color");
    }
    let num_tokens = sc.incoming().balance(&color);
    if num_tokens == 0 {
        sc.panic("Missing auction tokens");
    }

    let minimum_bid = params.get_int(KEY_MINIMUM_BID).value();
    if minimum_bid == 0 {
        sc.panic("Missing minimum bid");
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

    let state = sc.state();
    let mut owner_margin = state.get_int(KEY_OWNER_MARGIN).value();
    if owner_margin == 0 {
        owner_margin = OWNER_MARGIN_DEFAULT;
    }

    // need at least 1 iota to run SC
    let mut margin = minimum_bid * owner_margin / 1000;
    if margin == 0 {
        margin = 1;
    }
    let deposit = sc.incoming().balance(&ScColor::IOTA);
    if deposit < margin {
        sc.panic("Insufficient deposit");
    }

    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let auction_info = current_auction.get_bytes(KEY_INFO);
    if auction_info.exists() {
        sc.panic("Auction for this token color already exists");
    }

    let auction = &AuctionInfo {
        creator: sc.caller(),
        color: color,
        deposit: deposit,
        description: description,
        duration: duration,
        highest_bid: -1,
        highest_bidder: ScAgent::NULL,
        minimum_bid: minimum_bid,
        num_tokens: num_tokens,
        owner_margin: owner_margin,
        when_started: sc.timestamp(),
    };
    auction_info.set_value(&encode_auction_info(auction));

    let finalize_params = ScMutableMap::new();
    finalize_params.get_color(KEY_COLOR).set_value(&auction.color);
    sc.post(&ScAddress::NULL,
            Hname::SELF,
            Hname::new("finalize_auction"),
            finalize_params,
            &ScTransfers::NONE,
            duration * 60);
    sc.log("New auction started");
}

fn finalize_auction(sc: &ScCallContext) {
    // can only be sent by SC itself
    if !sc.from(&sc.contract_id()) {
        sc.panic("Cancel spoofed request");
    }

    let color_param = sc.params().get_color(KEY_COLOR);
    if !color_param.exists() {
        sc.panic("Missing token color");
    }
    let color = color_param.value();

    let state = sc.state();
    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let auction_info = current_auction.get_bytes(KEY_INFO);
    if !auction_info.exists() {
        sc.panic("Missing auction info");
    }
    let auction = decode_auction_info(&auction_info.value());
    if auction.highest_bid < 0 {
        sc.log(&("No one bid on ".to_string() + &color.to_string()));
        let mut owner_fee = auction.minimum_bid * auction.owner_margin / 1000;
        if owner_fee == 0 {
            owner_fee = 1
        }
        // finalizeAuction request token was probably not confirmed yet
        transfer(sc, &sc.contract_creator(), &ScColor::IOTA, owner_fee - 1);
        transfer(sc, &auction.creator, &auction.color, auction.num_tokens);
        transfer(sc, &auction.creator, &ScColor::IOTA, auction.deposit - owner_fee);
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
        if !bidder.equals(&auction.highest_bidder) {
            let loser = bidders.get_bytes(&bidder);
            let bid = decode_bid_info(&loser.value());
            transfer(sc, &bidder, &ScColor::IOTA, bid.amount);
        }
    }

    // finalizeAuction request token was probably not confirmed yet
    transfer(sc, &sc.contract_creator(), &ScColor::IOTA, owner_fee - 1);
    transfer(sc, &auction.highest_bidder, &auction.color, auction.num_tokens);
    transfer(sc, &auction.creator, &ScColor::IOTA, auction.deposit + auction.highest_bid - owner_fee);
}

fn place_bid(sc: &ScCallContext) {
    let mut bid_amount = sc.incoming().balance(&ScColor::IOTA);
    if bid_amount == 0 {
        sc.panic("Missing bid amount");
    }

    let color_param = sc.params().get_color(KEY_COLOR);
    if !color_param.exists() {
        sc.panic("Missing token color");
    }
    let color = color_param.value();

    let state = sc.state();
    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let auction_info = current_auction.get_bytes(KEY_INFO);
    if !auction_info.exists() {
        sc.panic("Missing auction info");
    }

    let mut auction = decode_auction_info(&auction_info.value());
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
        if bid_amount < auction.minimum_bid {
            sc.panic("Insufficient bid amount");
        }
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
        sc.log("New highest bidder");
        auction.highest_bid = bid_amount;
        auction.highest_bidder = caller;
        auction_info.set_value(&encode_auction_info(&auction));
    }
}

fn set_owner_margin(sc: &ScCallContext) {
    // can only be sent by SC creator
    if !sc.from(&sc.contract_creator()) {
        sc.panic("Cancel spoofed request");
    }

    let mut owner_margin = sc.params().get_int(KEY_OWNER_MARGIN).value();
    if owner_margin < OWNER_MARGIN_MIN {
        owner_margin = OWNER_MARGIN_MIN;
    }
    if owner_margin > OWNER_MARGIN_MAX {
        owner_margin = OWNER_MARGIN_MAX;
    }
    sc.state().get_int(KEY_OWNER_MARGIN).set_value(owner_margin);
    sc.log("Updated owner margin");
}

fn get_info(sc: &ScViewContext) {
    let color_param = sc.params().get_color(KEY_COLOR);
    if !color_param.exists() {
        sc.panic("Missing token color");
    }
    let color = color_param.value();

    let state = sc.state();
    let auctions = state.get_map(KEY_AUCTIONS);
    let current_auction = auctions.get_map(&color);
    let auction_info = current_auction.get_bytes(KEY_INFO);
    if !auction_info.exists() {
        sc.panic("Missing auction info");
    }

    let auction = decode_auction_info(&auction_info.value());
    let results = sc.results();
    results.get_color(KEY_COLOR).set_value(&auction.color);
    results.get_agent(KEY_CREATOR).set_value(&auction.creator);
    results.get_int(KEY_DEPOSIT).set_value(auction.deposit);
    results.get_string(KEY_DESCRIPTION).set_value(&auction.description);
    results.get_int(KEY_DURATION).set_value(auction.duration);
    results.get_int(KEY_HIGHEST_BID).set_value(auction.highest_bid);
    results.get_agent(KEY_HIGHEST_BIDDER).set_value(&auction.highest_bidder);
    results.get_int(KEY_MINIMUM_BID).set_value(auction.minimum_bid);
    results.get_int(KEY_NUM_TOKENS).set_value(auction.num_tokens);
    results.get_int(KEY_OWNER_MARGIN).set_value(auction.owner_margin);
    results.get_int(KEY_WHEN_STARTED).set_value(auction.when_started);

    let bidder_list = current_auction.get_agent_array(KEY_BIDDER_LIST);
    results.get_int(KEY_BIDDERS).set_value(bidder_list.length() as i64);
}

fn transfer(sc: &ScCallContext, agent: &ScAgent, color: &ScColor, amount: i64) {
    if agent.is_address() {
        // send to original Tangle address
        sc.transfer_to_address(&agent.address(), &ScTransfers::new(color, amount));
        return;
    }

    // TODO not an address, deposit into account on chain
    sc.transfer_to_address(&agent.address(), &ScTransfers::new(color, amount));
}
