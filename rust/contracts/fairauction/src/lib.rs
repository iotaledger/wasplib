#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScColor;
use wasplib::client::ScContext;
use wasplib::client::ScExports;

const DURATION_DEFAULT: i64 = 60;
const DURATION_MIN: i64 = 1;
const DURATION_MAX: i64 = 120;
const MAX_DESCRIPTION_LENGTH: usize = 150;
const OWNER_MARGIN_DEFAULT: i64 = 50;
const OWNER_MARGIN_MIN: i64 = 5;
const OWNER_MARGIN_MAX: i64 = 100;

struct AuctionInfo {
    // color of tokens for sale
    color: ScColor,
    // number of tokens for sale
    numTokens: i64,
    // minimum bid. Set by the auction initiator
    minimumBid: i64,
    // any text, like "AuctionOwner of the token have a right to call me for a date". Set by auction initiator
    description: String,
    // timestamp when auction started
    whenStarted: i64,
    // duration of the auctions in minutes. Should be >= MinAuctionDurationMinutes
    duration: i64,
    // address which issued StartAuction transaction
    auctionOwner: String,
    // deposit by the auction owner. Iotas sent by the auction owner together with the tokens for sale in the same
    // transaction.
    deposit: i64,
    // AuctionOwner's margin in promilles, taken at the moment of creation of smart contract
    ownerMargin: i64,
    // list of bids to the auction
    bids: Vec<BidInfo>,
}

struct BidInfo {
    // originator of the bid
    address: String,
    // the amount is a cumulative sum of all bids from the same bidder
    amount: i64,
    // most recent bid update time
    when: i64,
}

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("startAuction");
    exports.add("finalizeAuction");
    exports.add("placeBid");
    exports.add_protected("setOwnerMargin");
}

#[no_mangle]
pub fn startAuction() {
    let sc = ScContext::new();
    let request = sc.request();
    let deposit = request.balance(&ScColor::iota());
    if deposit < 1 {
        sc.log("Empty deposit...");
        return;
    }

    let state = sc.state();
    let mut ownerMargin = state.get_int("ownerMargin").value();
    if ownerMargin == 0 {
        ownerMargin = OWNER_MARGIN_DEFAULT;
    }

    let params = request.params();
    let bytes = params.get_string("color").value();
    if bytes.len() == 0 {
        refund(deposit / 2, "Missing token color...");
        return;
    }
    let color = ScColor::from_bytes(&bytes);

    if color == ScColor::iota() || color == ScColor::mint() {
        refund(deposit / 2, "Reserved token color...");
        return;
    }

    let numTokens = request.balance(&color);
    if numTokens == 0 {
        refund(deposit / 2, "Auction tokens missing from request...");
        return;
    }

    let minimumBid = params.get_int("minimum").value();
    if minimumBid == 0 {
        refund(deposit / 2, "Missing minimum bid...");
        return;
    }

    // need at least 1 iota to run SC
    let mut margin = minimumBid * ownerMargin / 1000;
    if margin == 0 {
        margin = 1;
    }
    if deposit < margin {
        refund(deposit / 2, "Insufficient deposit...");
        return;
    }

    // duration in minutes
    let mut duration = params.get_int("duration").value();
    if duration == 0 {
        duration = DURATION_DEFAULT;
    }
    if duration < DURATION_MIN {
        duration = DURATION_MIN;
    }
    if duration > DURATION_MAX {
        duration = DURATION_MAX;
    }

    let mut description = params.get_string("dscr").value();
    if description == "" {
        description = "N/A".to_string()
    }
    if description.len() > MAX_DESCRIPTION_LENGTH {
        let ss: String = description.chars().take(MAX_DESCRIPTION_LENGTH).collect();
        description = ss + "[...]";
    }

    let auctions = state.get_map("auctions");
    let currentAuction = auctions.get_bytes(&color.as_bytes());
    if currentAuction.value().len() != 0 {
        refund(deposit / 2, "Auction for this token already exists...");
        return;
    }

    let auction = &AuctionInfo {
        color: color,
        numTokens: numTokens,
        minimumBid: minimumBid,
        description: description,
        whenStarted: request.timestamp(),
        duration: duration,
        auctionOwner: request.address(),
        deposit: deposit,
        ownerMargin: ownerMargin,
        bids: Vec::new(),
    };
    let bytes = encodeAuctionInfo(auction);
    currentAuction.set_value(&bytes);

    let finalizeparams = sc.post_request(&sc.contract().address(), "finalizeAuction", duration * 60);
    finalizeparams.get_string("color").set_value(&auction.color.as_bytes());
    sc.log("New auction started...");
}

#[no_mangle]
pub fn finalizeAuction() {
    // can only be sent by SC itself
    let sc = ScContext::new();
    let request = sc.request();
    if request.address() != sc.contract().address() {
        sc.log("Cancel spoofed request");
        return;
    }

    let bytes = request.params().get_string("color").value();
    if bytes.len() == 0 {
        sc.log("INTERNAL INCONSISTENCY: missing color");
        return;
    }
    let color = ScColor::from_bytes(&bytes);

    let state = sc.state();
    let auctions = state.get_map("auctions");
    let currentAuction = auctions.get_bytes(&color.as_bytes());
    let bytes = currentAuction.value();
    if bytes.len() == 0 {
        sc.log("INTERNAL INCONSISTENCY missing auction info");
        return;
    }
    let auction = decodeAuctionInfo(&bytes);
    if auction.bids.len() == 0 {
        sc.log(&("No one bid on ".to_string() + &color.as_string()));
        let mut ownerFee = auction.minimumBid * auction.ownerMargin / 1000;
        if ownerFee == 0 {
            ownerFee = 1
        }
        // finalizeAuction request token was probably not confirmed yet
        sc.transfer(&sc.contract().owner(), &ScColor::iota(), ownerFee - 1);
        sc.transfer(&auction.auctionOwner, &ScColor::iota(), auction.deposit - ownerFee);
        return;
    }

    let mut winner = BidInfo {
        amount: 0,
        address: String::new(),
        when: 0,
    };
    for bidder in &auction.bids {
        if bidder.amount >= winner.amount {
            if bidder.amount > winner.amount || bidder.when < winner.when {
                winner.amount = bidder.amount;
                winner.address = bidder.address.to_string();
                winner.when = bidder.when;
            }
        }
    }
    let mut ownerFee = winner.amount * auction.ownerMargin / 1000;
    if ownerFee == 0 {
        ownerFee = 1;
    }

    // return staked bids to losers
    for bidder in auction.bids {
        if bidder.address != winner.address {
            sc.transfer(&bidder.address, &ScColor::iota(), bidder.amount);
        }
    }

    // finalizeAuction request token was probably not confirmed yet
    sc.transfer(&sc.contract().owner(), &ScColor::iota(), ownerFee - 1);
    sc.transfer(&winner.address, &auction.color, auction.numTokens);
    sc.transfer(&auction.auctionOwner, &ScColor::iota(), auction.deposit + winner.amount - ownerFee);
}

#[no_mangle]
pub fn placeBid() {
    let sc = ScContext::new();
    let request = sc.request();
    let bidAmount = request.balance(&ScColor::iota());
    if bidAmount == 0 {
        sc.log("Insufficient bid amount");
        return;
    }

    let bytes = request.params().get_string("color").value();
    if bytes.len() == 0 {
        refund(bidAmount, "Missing token color");
        return;
    }
    let color = ScColor::from_bytes(&bytes);

    let state = sc.state();
    let auctions = state.get_map("auctions");
    let currentAuction = auctions.get_bytes(&color.as_bytes());
    let bytes = currentAuction.value();
    if bytes.len() == 0 {
        refund(bidAmount, "Missing auction");
        return;
    }

    let sender = request.address();
    let mut auction = decodeAuctionInfo(&bytes);
    let mut bidIndex = auction.bids.iter().position(|b| b.address == sender);
    if bidIndex == None {
        sc.log(&("New bid from: ".to_string() + &sender));
        let bid = BidInfo { address: sender, amount: 0, when: 0 };
        bidIndex = Some(auction.bids.len());
        auction.bids.push(bid);
    }
    let mut bid = auction.bids.get_mut(bidIndex.unwrap()).unwrap();
    bid.amount += bidAmount;
    bid.when = request.timestamp();

    let bytes = encodeAuctionInfo(&auction);
    currentAuction.set_value(&bytes);
    sc.log("Updated auction with bid...");
}

#[no_mangle]
pub fn setOwnerMargin() {
    // can only be sent by SC owner
    let sc = ScContext::new();
    let request = sc.request();
    if request.address() != sc.contract().owner() {
        sc.log("Cancel spoofed request");
        return;
    }

    let mut ownerMargin = sc.request().params().get_int("ownerMargin").value();
    if ownerMargin < OWNER_MARGIN_MIN {
        ownerMargin = OWNER_MARGIN_MIN;
    }
    if ownerMargin > OWNER_MARGIN_MAX {
        ownerMargin = OWNER_MARGIN_MAX;
    }
    sc.state().get_int("ownerMargin").set_value(ownerMargin);
    sc.log("Updated owner margin...");
}

fn decodeAuctionInfo(bytes: &[u8]) -> AuctionInfo {
    let mut decoder = BytesDecoder::new(bytes);
    let mut auction = AuctionInfo {
        color: ScColor::from_bytes(&decoder.string()),
        numTokens: decoder.int(),
        minimumBid: decoder.int(),
        description: decoder.string(),
        whenStarted: decoder.int(),
        duration: decoder.int(),
        auctionOwner: decoder.string(),
        deposit: decoder.int(),
        ownerMargin: decoder.int(),
        bids: Vec::new(),
    };
    let bids = decoder.int();
    for _ in 0..bids {
        let bytes = decoder.bytes();
        let bid = decodeBidInfo(&bytes);
        auction.bids.push(bid);
    }
    return auction;
}

fn decodeBidInfo(bytes: &[u8]) -> BidInfo {
    let mut decoder = BytesDecoder::new(bytes);
    BidInfo {
        address: decoder.string(),
        amount: decoder.int(),
        when: decoder.int(),
    }
}

fn encodeAuctionInfo(auction: &AuctionInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.string(&auction.color.as_bytes());
    encoder.int(auction.numTokens);
    encoder.int(auction.minimumBid);
    encoder.string(&auction.description);
    encoder.int(auction.whenStarted);
    encoder.int(auction.duration);
    encoder.string(&auction.auctionOwner);
    encoder.int(auction.deposit);
    encoder.int(auction.ownerMargin);
    encoder.int(auction.bids.len() as i64);
    for bid in &auction.bids {
        let bytes = encodeBidInfo(&bid);
        encoder.bytes(&bytes);
    }
    return encoder.data();
}

fn encodeBidInfo(bid: &BidInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.string(&bid.address);
    encoder.int(bid.amount);
    encoder.int(bid.when);
    encoder.data()
}

fn refund(amount: i64, reason: &str) {
    let sc = ScContext::new();
    sc.log(reason);
    let request = sc.request();
    let sender = request.address();
    if amount != 0 {
        sc.transfer(&sender, &ScColor::iota(), amount);
    }
    let deposit = request.balance(&ScColor::iota());
    if deposit - amount != 0 {
        sc.transfer(&sc.contract().owner(), &ScColor::iota(), deposit - amount);
    }

    // refund all other token colors, don't keep tokens that were to be auctioned
    let colors = request.colors();
    let items = colors.length();
    for i in 0..items {
        let color = colors.get_color(i);
        if color != ScColor::iota() {
            sc.transfer(&sender, &color, request.balance(&color));
        }
    }
}