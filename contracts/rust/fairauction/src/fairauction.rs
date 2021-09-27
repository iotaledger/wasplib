// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasmlib::*;

use crate::*;
use crate::contract::*;

const DURATION_DEFAULT: i32 = 60;
const DURATION_MIN: i32 = 1;
const DURATION_MAX: i32 = 120;
const MAX_DESCRIPTION_LENGTH: usize = 150;
const OWNER_MARGIN_DEFAULT: i64 = 50;
const OWNER_MARGIN_MIN: i64 = 5;
const OWNER_MARGIN_MAX: i64 = 100;

pub fn func_finalize_auction(ctx: &ScFuncContext, f: &FinalizeAuctionContext) {
    let color = f.params.color().value();
    let current_auction = f.state.auctions().get_auction(&color);
    ctx.require(current_auction.exists(), "Missing auction info");
    let auction = current_auction.value();
    if auction.highest_bid < 0 {
        ctx.log(&("No one bid on ".to_string() + &color.to_string()));
        let mut owner_fee = auction.minimum_bid * auction.owner_margin / 1000;
        if owner_fee == 0 {
            owner_fee = 1;
        }
        // finalizeAuction request token was probably not confirmed yet
        transfer_tokens(ctx, &ctx.contract_creator(), &ScColor::IOTA, owner_fee - 1);
        transfer_tokens(ctx, &auction.creator, &auction.color, auction.num_tokens);
        transfer_tokens(ctx, &auction.creator, &ScColor::IOTA, auction.deposit - owner_fee);
        return;
    }

    let mut owner_fee = auction.highest_bid * auction.owner_margin / 1000;
    if owner_fee == 0 {
        owner_fee = 1;
    }

    // return staked bids to losers
    let bids = f.state.bids().get_bids(&color);
    let bidder_list = f.state.bidder_list().get_bidder_list(&color);
    let size = bidder_list.length();
    for i in 0..size {
        let loser = bidder_list.get_agent_id(i).value();
        if loser != auction.highest_bidder {
            let bid = bids.get_bid(&loser).value();
            transfer_tokens(ctx, &loser, &ScColor::IOTA, bid.amount);
        }
    }

    // finalizeAuction request token was probably not confirmed yet
    transfer_tokens(ctx, &ctx.contract_creator(), &ScColor::IOTA, owner_fee - 1);
    transfer_tokens(ctx, &auction.highest_bidder, &auction.color, auction.num_tokens);
    transfer_tokens(ctx, &auction.creator, &ScColor::IOTA, auction.deposit + auction.highest_bid - owner_fee);
}

pub fn func_place_bid(ctx: &ScFuncContext, f: &PlaceBidContext) {
    let mut bid_amount = ctx.incoming().balance(&ScColor::IOTA);
    ctx.require(bid_amount > 0, "Missing bid amount");

    let color = f.params.color().value();
    let current_auction = f.state.auctions().get_auction(&color);
    ctx.require(current_auction.exists(), "Missing auction info");

    let mut auction = current_auction.value();
    let bids = f.state.bids().get_bids(&color);
    let bidder_list = f.state.bidder_list().get_bidder_list(&color);
    let caller = ctx.caller();
    let current_bid = bids.get_bid(&caller);
    if current_bid.exists() {
        ctx.log(&("Upped bid from: ".to_string() + &caller.to_string()));
        let mut bid = current_bid.value();
        bid_amount += bid.amount;
        bid.amount = bid_amount;
        bid.timestamp = ctx.timestamp();
        current_bid.set_value(&bid);
    } else {
        ctx.require(bid_amount >= auction.minimum_bid, "Insufficient bid amount");
        ctx.log(&("New bid from: ".to_string() + &caller.to_string()));
        let index = bidder_list.length();
        bidder_list.get_agent_id(index).set_value(&caller);
        let bid = Bid {
            index: index,
            amount: bid_amount,
            timestamp: ctx.timestamp(),
        };
        current_bid.set_value(&bid);
    }
    if bid_amount > auction.highest_bid {
        ctx.log("New highest bidder");
        auction.highest_bid = bid_amount;
        auction.highest_bidder = caller;
        current_auction.set_value(&auction);
    }
}

pub fn func_set_owner_margin(_ctx: &ScFuncContext, f: &SetOwnerMarginContext) {
    let mut owner_margin = f.params.owner_margin().value();
    if owner_margin < OWNER_MARGIN_MIN {
        owner_margin = OWNER_MARGIN_MIN;
    }
    if owner_margin > OWNER_MARGIN_MAX {
        owner_margin = OWNER_MARGIN_MAX;
    }
    f.state.owner_margin().set_value(owner_margin);
}

pub fn func_start_auction(ctx: &ScFuncContext, f: &StartAuctionContext) {
    let color = f.params.color().value();
    if color == ScColor::IOTA || color == ScColor::MINT {
        ctx.panic("Reserved auction token color");
    }
    let num_tokens = ctx.incoming().balance(&color);
    if num_tokens == 0 {
        ctx.panic("Missing auction tokens");
    }

    let minimum_bid = f.params.minimum_bid().value();

    // duration in minutes
    let mut duration = f.params.duration().value();
    if duration == 0 {
        duration = DURATION_DEFAULT;
    }
    if duration < DURATION_MIN {
        duration = DURATION_MIN;
    }
    if duration > DURATION_MAX {
        duration = DURATION_MAX;
    }

    let mut description = f.params.description().value();
    if description == "" {
        description = "N/A".to_string();
    }
    if description.len() > MAX_DESCRIPTION_LENGTH {
        let ss: String = description.chars().take(MAX_DESCRIPTION_LENGTH).collect();
        description = ss + "[...]";
    }

    let mut owner_margin = f.state.owner_margin().value();
    if owner_margin == 0 {
        owner_margin = OWNER_MARGIN_DEFAULT;
    }

    // need at least 1 iota to run SC
    let mut margin = minimum_bid * owner_margin / 1000;
    if margin == 0 {
        margin = 1;
    }
    let deposit = ctx.incoming().balance(&ScColor::IOTA);
    if deposit < margin {
        ctx.panic("Insufficient deposit");
    }

    let current_auction = f.state.auctions().get_auction(&color);
    if current_auction.exists() {
        ctx.panic("Auction for this token color already exists");
    }

    let auction = Auction {
        creator: ctx.caller(),
        color: color,
        deposit: deposit,
        description: description,
        duration: duration,
        highest_bid: -1,
        highest_bidder: ScAgentID::from_bytes(&[0; 37]),
        minimum_bid: minimum_bid,
        num_tokens: num_tokens,
        owner_margin: owner_margin,
        when_started: ctx.timestamp(),
    };
    current_auction.set_value(&auction);

    let fa = ScFuncs::finalize_auction(ctx);
    fa.params.color().set_value(&auction.color);
    fa.func.delay(duration * 60).transfer_iotas(1).post();
}

pub fn view_get_info(ctx: &ScViewContext, f: &GetInfoContext) {
    let color = f.params.color().value();
    let current_auction = f.state.auctions().get_auction(&color);
    ctx.require(current_auction.exists(), "Missing auction info");

    let auction = current_auction.value();
    f.results.color().set_value(&auction.color);
    f.results.creator().set_value(&auction.creator);
    f.results.deposit().set_value(auction.deposit);
    f.results.description().set_value(&auction.description);
    f.results.duration().set_value(auction.duration);
    f.results.highest_bid().set_value(auction.highest_bid);
    f.results.highest_bidder().set_value(&auction.highest_bidder);
    f.results.minimum_bid().set_value(auction.minimum_bid);
    f.results.num_tokens().set_value(auction.num_tokens);
    f.results.owner_margin().set_value(auction.owner_margin);
    f.results.when_started().set_value(auction.when_started);

    let bidder_list = f.state.bidder_list().get_bidder_list(&color);
    f.results.bidders().set_value(bidder_list.length());
}

fn transfer_tokens(ctx: &ScFuncContext, agent: &ScAgentID, color: &ScColor, amount: i64) {
    if agent.is_address() {
        // send back to original Tangle address
        ctx.transfer_to_address(&agent.address(), ScTransfers::new(color, amount));
        return;
    }

    // TODO not an address, deposit into account on chain
    ctx.transfer_to_address(&agent.address(), ScTransfers::new(color, amount));
}
