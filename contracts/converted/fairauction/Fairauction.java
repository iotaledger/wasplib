// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

public class Fairauction {

private static final int DurationDefault = 60;
private static final int DurationMin = 1;
private static final int DurationMax = 120;
private static final int MaxDescriptionLength = 150;
private static final int OwnerMarginDefault = 50;
private static final int OwnerMarginMin = 5;
private static final int OwnerMarginMax = 100;

public static void funcFinalizeAuction(ScFuncContext ctx, FuncFinalizeAuctionParams params) {
    color = params.Color.Value();
    state = ctx.State();
    auctions = state.GetMap(VarAuctions);
    currentAuction = auctions.GetMap(color);
    auctionInfo = currentAuction.GetBytes(VarInfo);
    ctx.Require(auctionInfo.Exists(), "Missing auction info");
    auction = Auction::fromBytes(auctionInfo.Value());
    if (auction.HighestBid < 0) {
        ctx.Log("No one bid on " + color);
        ownerFee = auction.MinimumBid * auction.OwnerMargin / 1000;
        if (ownerFee == 0) {
            ownerFee = 1
        }
        // finalizeAuction request token was probably not confirmed yet
        transfer(ctx, ctx.ContractCreator(), ScColor.IOTA, ownerFee - 1);
        transfer(ctx, auction.Creator, auction.Color, auction.NumTokens);
        transfer(ctx, auction.Creator, ScColor.IOTA, auction.Deposit - ownerFee);
        return;
    }

    ownerFee = auction.HighestBid * auction.OwnerMargin / 1000;
    if (ownerFee == 0) {
        ownerFee = 1;
    }

    // return staked bids to losers
    bidders = currentAuction.GetMap(VarBidders);
    bidderList = currentAuction.GetAgentIdArray(VarBidderList);
    size = bidderList.Length();
    for (int i = 0; i < size; i++) {
        bidder = bidderList.GetAgentId(i).Value();
        if (bidder != auction.HighestBidder) {
            loser = bidders.GetBytes(bidder);
            bid = Bid::fromBytes(loser.Value());
            transfer(ctx, bidder, ScColor.IOTA, bid.Amount);
        }
    }

    // finalizeAuction request token was probably not confirmed yet
    transfer(ctx, ctx.ContractCreator(), ScColor.IOTA, ownerFee - 1);
    transfer(ctx, auction.HighestBidder, auction.Color, auction.NumTokens);
    transfer(ctx, auction.Creator, ScColor.IOTA, auction.Deposit + auction.HighestBid - ownerFee);
}

public static void funcPlaceBid(ScFuncContext ctx, FuncPlaceBidParams params) {
    bidAmount = ctx.Incoming().Balance(ScColor.IOTA);
    ctx.Require(bidAmount > 0, "Missing bid amount");

    color = params.Color.Value();
    state = ctx.State();
    auctions = state.GetMap(VarAuctions);
    currentAuction = auctions.GetMap(color);
    auctionInfo = currentAuction.GetBytes(VarInfo);
    ctx.Require(auctionInfo.Exists(), "Missing auction info");

    auction = Auction::fromBytes(auctionInfo.Value());
    bidders = currentAuction.GetMap(VarBidders);
    bidderList = currentAuction.GetAgentIdArray(VarBidderList);
    caller = ctx.Caller();
    bidder = bidders.GetBytes(caller);
    if (bidder.Exists()) {
        ctx.Log("Upped bid from: " + caller);
        bid = Bid::fromBytes(bidder.Value());
        bidAmount += bid.Amount;
        bid.Amount = bidAmount;
        bid.Timestamp = ctx.Timestamp();
        bidder.SetValue(bid.ToBytes());
    } else {
        ctx.Require(bidAmount >= auction.MinimumBid, "Insufficient bid amount");
        ctx.Log("New bid from: " + caller);
        index = bidderList.Length();
        bidderList.GetAgentId(index).SetValue(caller);
        Bid bid = new Bid();
         {
            bid.Index = index as i64;
            bid.Amount = bidAmount;
            bid.Timestamp = ctx.Timestamp();
        }
        bidder.SetValue(bid.ToBytes());
    }
    if (bidAmount > auction.HighestBid) {
        ctx.Log("New highest bidder");
        auction.HighestBid = bidAmount;
        auction.HighestBidder = caller;
        auctionInfo.SetValue(auction.ToBytes());
    }
}

public static void funcSetOwnerMargin(ScFuncContext ctx, FuncSetOwnerMarginParams params) {
    ownerMargin = params.OwnerMargin.Value();
    if (ownerMargin < OwnerMarginMin) {
        ownerMargin = OwnerMarginMin;
    }
    if (ownerMargin > OwnerMarginMax) {
        ownerMargin = OwnerMarginMax;
    }
    ctx.State().GetInt(VarOwnerMargin).SetValue(ownerMargin);
    ctx.Log("Updated owner margin");
}

public static void funcStartAuction(ScFuncContext ctx, FuncStartAuctionParams params) {
    color = params.Color.Value();
    if (color == ScColor.IOTA || color == ScColor.MINT) {
        ctx.Panic("Reserved auction token color");
    }
    numTokens = ctx.Incoming().Balance(color);
    if (numTokens == 0) {
        ctx.Panic("Missing auction tokens");
    }

    minimumBid = params.MinimumBid.Value();

    // duration in minutes
    duration = params.Duration.Value();
    if (duration == 0) {
        duration = DurationDefault;
    }
    if (duration < DurationMin) {
        duration = DurationMin;
    }
    if (duration > DurationMax) {
        duration = DurationMax;
    }

    description = params.Description.Value();
    if (description == "") {
        description = "N/A"
    }
    if (description.Len() > MaxDescriptionLength) {
        let ss: String = description.Chars().Take(MaxDescriptionLength).Collect();
        description = ss + "[...]";
    }

    state = ctx.State();
    ownerMargin = state.GetInt(VarOwnerMargin).Value();
    if (ownerMargin == 0) {
        ownerMargin = OwnerMarginDefault;
    }

    // need at least 1 iota to run SC
    margin = minimumBid * ownerMargin / 1000;
    if (margin == 0) {
        margin = 1;
    }
    deposit = ctx.Incoming().Balance(ScColor.IOTA);
    if (deposit < margin) {
        ctx.Panic("Insufficient deposit");
    }

    auctions = state.GetMap(VarAuctions);
    currentAuction = auctions.GetMap(color);
    auctionInfo = currentAuction.GetBytes(VarInfo);
    if (auctionInfo.Exists()) {
        ctx.Panic("Auction for this token color already exists");
    }

    Auction auction = new Auction();
         {
        auction.Creator = ctx.Caller();
        auction.Color = color;
        auction.Deposit = deposit;
        auction.Description = description;
        auction.Duration = duration;
        auction.HighestBid = -1;
        auction.HighestBidder = ScAgentId::fromBytes([0; 37]);
        auction.MinimumBid = minimumBid;
        auction.NumTokens = numTokens;
        auction.OwnerMargin = ownerMargin;
        auction.WhenStarted = ctx.Timestamp();
    }
    auctionInfo.SetValue(auction.ToBytes());

    finalizeParams = new ScMutableMapp();
    finalizeParams.GetColor(VarColor).SetValue(auction.Color);
    ctx.Post(PostRequestParams {
        auction.ContractId = ctx.ContractId();
        auction.Function = HFuncFinalizeAuction;
        auction.Params = Some(finalizeParams);
        auction.Transfer = null;
        auction.Delay = duration * 60;
    });
    ctx.Log("New auction started");
}

public static void viewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
    color = params.Color.Value();
    state = ctx.State();
    auctions = state.GetMap(VarAuctions);
    currentAuction = auctions.GetMap(color);
    auctionInfo = currentAuction.GetBytes(VarInfo);
    if (!auctionInfo.Exists()) {
        ctx.Panic("Missing auction info");
    }

    auction = Auction::fromBytes(auctionInfo.Value());
    results = ctx.Results();
    results.GetColor(VarColor).SetValue(auction.Color);
    results.GetAgentId(VarCreator).SetValue(auction.Creator);
    results.GetInt(VarDeposit).SetValue(auction.Deposit);
    results.GetString(VarDescription).SetValue(auction.Description);
    results.GetInt(VarDuration).SetValue(auction.Duration);
    results.GetInt(VarHighestBid).SetValue(auction.HighestBid);
    results.GetAgentId(VarHighestBidder).SetValue(auction.HighestBidder);
    results.GetInt(VarMinimumBid).SetValue(auction.MinimumBid);
    results.GetInt(VarNumTokens).SetValue(auction.NumTokens);
    results.GetInt(VarOwnerMargin).SetValue(auction.OwnerMargin);
    results.GetInt(VarWhenStarted).SetValue(auction.WhenStarted);

    bidderList = currentAuction.GetAgentIdArray(VarBidderList);
    results.GetInt(VarBidders).SetValue(bidderList.Length() as i64);
}

public static void transfer(ScFuncContext ctx, ScAgentId agent, ScColor color, i64 amount) {
    if (agent.IsAddress()) {
        // send back to original Tangle address
        ctx.TransferToAddress(agent.Address(), new ScTransfers(color, amount));
        return;
    }

    // TODO not an address, deposit into account on chain
    ctx.TransferToAddress(agent.Address(), new ScTransfers(color, amount));
}
}
