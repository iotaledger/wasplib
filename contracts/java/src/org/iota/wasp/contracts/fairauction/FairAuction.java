// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.fairauction;

import org.iota.wasp.contracts.fairauction.lib.*;
import org.iota.wasp.contracts.fairauction.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class FairAuction {

    private static final int DurationDefault = 60;
    private static final int DurationMin = 1;
    private static final int DurationMax = 120;
    private static final int MaxDescriptionLength = 150;
    private static final int OwnerMarginDefault = 50;
    private static final int OwnerMarginMin = 5;
    private static final int OwnerMarginMax = 100;

    public static void funcFinalizeAuction(ScFuncContext ctx, FuncFinalizeAuctionParams params) {
        var color = params.Color.Value();
        var state = ctx.State();
        var auctions = state.GetMap(Consts.VarAuctions);
        var currentAuction = auctions.GetMap(color);
        var auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
        ctx.Require(auctionInfo.Exists(), "Missing auction info");
        var auction = new Auction(auctionInfo.Value());
        if (auction.HighestBid < 0) {
            ctx.Log("No one bid on " + color);
            var ownerFee = auction.MinimumBid * auction.OwnerMargin / 1000;
            if (ownerFee == 0) {
                ownerFee = 1;
            }
            // finalizeAuction request token was probably not confirmed yet
            transfer(ctx, ctx.ContractCreator(), ScColor.IOTA, ownerFee - 1);
            transfer(ctx, auction.Creator, auction.Color, auction.NumTokens);
            transfer(ctx, auction.Creator, ScColor.IOTA, auction.Deposit - ownerFee);
            return;
        }

        var ownerFee = auction.HighestBid * auction.OwnerMargin / 1000;
        if (ownerFee == 0) {
            ownerFee = 1;
        }

        // return staked bids to losers
        var bidders = currentAuction.GetMap(Consts.VarBidders);
        var bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
        var size = bidderList.Length();
        for (var i = 0; i < size; i++) {
            var bidder = bidderList.GetAgentId(i).Value();
            if (!bidder.equals(auction.HighestBidder)) {
                var loser = bidders.GetBytes(bidder);
                var bid = new Bid(loser.Value());
                transfer(ctx, bidder, ScColor.IOTA, bid.Amount);
            }
        }

        // finalizeAuction request token was probably not confirmed yet
        transfer(ctx, ctx.ContractCreator(), ScColor.IOTA, ownerFee - 1);
        transfer(ctx, auction.HighestBidder, auction.Color, auction.NumTokens);
        transfer(ctx, auction.Creator, ScColor.IOTA, auction.Deposit + auction.HighestBid - ownerFee);
    }

    public static void funcPlaceBid(ScFuncContext ctx, FuncPlaceBidParams params) {
        var bidAmount = ctx.Incoming().Balance(ScColor.IOTA);
        ctx.Require(bidAmount > 0, "Missing bid amount");

        var color = params.Color.Value();
        var state = ctx.State();
        var auctions = state.GetMap(Consts.VarAuctions);
        var currentAuction = auctions.GetMap(color);
        var auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
        ctx.Require(auctionInfo.Exists(), "Missing auction info");

        var auction = new Auction(auctionInfo.Value());
        var bidders = currentAuction.GetMap(Consts.VarBidders);
        var bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
        var caller = ctx.Caller();
        var bidder = bidders.GetBytes(caller);
        if (bidder.Exists()) {
            ctx.Log("Upped bid from: " + caller);
            var bid = new Bid(bidder.Value());
            bidAmount += bid.Amount;
            bid.Amount = bidAmount;
            bid.Timestamp = ctx.Timestamp();
            bidder.SetValue(bid.toBytes());
        } else {
            ctx.Require(bidAmount >= auction.MinimumBid, "Insufficient bid amount");
            ctx.Log("New bid from: " + caller);
            var index = bidderList.Length();
            bidderList.GetAgentId(index).SetValue(caller);
            var bid = new Bid();
            {
                bid.Index = index;
                bid.Amount = bidAmount;
                bid.Timestamp = ctx.Timestamp();
            }
            bidder.SetValue(bid.toBytes());
        }
        if (bidAmount > auction.HighestBid) {
            ctx.Log("New highest bidder");
            auction.HighestBid = bidAmount;
            auction.HighestBidder = caller;
            auctionInfo.SetValue(auction.toBytes());
        }
    }

    public static void funcSetOwnerMargin(ScFuncContext ctx, FuncSetOwnerMarginParams params) {
        var ownerMargin = params.OwnerMargin.Value();
        if (ownerMargin < OwnerMarginMin) {
            ownerMargin = OwnerMarginMin;
        }
        if (ownerMargin > OwnerMarginMax) {
            ownerMargin = OwnerMarginMax;
        }
        ctx.State().GetInt64(Consts.VarOwnerMargin).SetValue(ownerMargin);
    }

    public static void funcStartAuction(ScFuncContext ctx, FuncStartAuctionParams params) {
        var color = params.Color.Value();
        if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
            ctx.Panic("Reserved auction token color");
        }
        var numTokens = ctx.Incoming().Balance(color);
        if (numTokens == 0) {
            ctx.Panic("Missing auction tokens");
        }

        var minimumBid = params.MinimumBid.Value();

        // duration in minutes
        var duration = params.Duration.Value();
        if (duration == 0) {
            duration = DurationDefault;
        }
        if (duration < DurationMin) {
            duration = DurationMin;
        }
        if (duration > DurationMax) {
            duration = DurationMax;
        }

        var description = params.Description.Value();
        if (description.equals("")) {
            description = "N/A";
        }
        if (description.length() > MaxDescriptionLength) {
            var ss = description.substring(0, MaxDescriptionLength);
            description = ss + "[...]";
        }

        var state = ctx.State();
        var ownerMargin = state.GetInt64(Consts.VarOwnerMargin).Value();
        if (ownerMargin == 0) {
            ownerMargin = OwnerMarginDefault;
        }

        // need at least 1 iota to run SC
        var margin = minimumBid * ownerMargin / 1000;
        if (margin == 0) {
            margin = 1;
        }
        var deposit = ctx.Incoming().Balance(ScColor.IOTA);
        if (deposit < margin) {
            ctx.Panic("Insufficient deposit");
        }

        var auctions = state.GetMap(Consts.VarAuctions);
        var currentAuction = auctions.GetMap(color);
        var auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
        if (auctionInfo.Exists()) {
            ctx.Panic("Auction for this token color already exists");
        }

        var auction = new Auction();
        {
            auction.Creator = ctx.Caller();
            auction.Color = color;
            auction.Deposit = deposit;
            auction.Description = description;
            auction.Duration = duration;
            auction.HighestBid = -1;
            auction.HighestBidder = new ScAgentId(new byte[37]);
            auction.MinimumBid = minimumBid;
            auction.NumTokens = numTokens;
            auction.OwnerMargin = ownerMargin;
            auction.WhenStarted = ctx.Timestamp();
        }
        auctionInfo.SetValue(auction.toBytes());

        var finalizeParams = new ScMutableMap();
        finalizeParams.GetColor(Consts.VarColor).SetValue(auction.Color);
        var transfer = ScTransfers.iotas(1);
        ctx.PostSelf(Consts.HFuncFinalizeAuction, finalizeParams, transfer, duration * 60);
    }

    public static void viewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
        var color = params.Color.Value();
        var state = ctx.State();
        var auctions = state.GetMap(Consts.VarAuctions);
        var currentAuction = auctions.GetMap(color);
        var auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
        if (!auctionInfo.Exists()) {
            ctx.Panic("Missing auction info");
        }

        var auction = new Auction(auctionInfo.Value());
        var results = ctx.Results();
        results.GetColor(Consts.VarColor).SetValue(auction.Color);
        results.GetAgentId(Consts.VarCreator).SetValue(auction.Creator);
        results.GetInt64(Consts.VarDeposit).SetValue(auction.Deposit);
        results.GetString(Consts.VarDescription).SetValue(auction.Description);
        results.GetInt64(Consts.VarDuration).SetValue(auction.Duration);
        results.GetInt64(Consts.VarHighestBid).SetValue(auction.HighestBid);
        results.GetAgentId(Consts.VarHighestBidder).SetValue(auction.HighestBidder);
        results.GetInt64(Consts.VarMinimumBid).SetValue(auction.MinimumBid);
        results.GetInt64(Consts.VarNumTokens).SetValue(auction.NumTokens);
        results.GetInt64(Consts.VarOwnerMargin).SetValue(auction.OwnerMargin);
        results.GetInt64(Consts.VarWhenStarted).SetValue(auction.WhenStarted);

        var bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
        results.GetInt64(Consts.VarBidders).SetValue(bidderList.Length());
    }

    public static void transfer(ScFuncContext ctx, ScAgentId agent, ScColor color, long amount) {
        if (agent.IsAddress()) {
            // send back to original Tangle address
            ctx.TransferToAddress(agent.Address(), new ScTransfers(color, amount));
            return;
        }

        // TODO not an address, deposit into account on chain
        ctx.TransferToAddress(agent.Address(), new ScTransfers(color, amount));
    }
}
