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
		ScColor color = params.Color.Value();
		ScMutableMap state = ctx.State();
		ScMutableMap auctions = state.GetMap(Consts.VarAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
		ctx.Require(auctionInfo.Exists(), "Missing auction info");
		Auction auction = new Auction(auctionInfo.Value());
		long ownerFee;
		if (auction.HighestBid < 0) {
			ctx.Log("No one bid on " + color);
			ownerFee = auction.MinimumBid * auction.OwnerMargin / 1000;
			if (ownerFee == 0) {
				ownerFee = 1;
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
		ScMutableMap bidders = currentAuction.GetMap(Consts.VarBidders);
		ScMutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
		int size = bidderList.Length();
		for (int i = 0; i < size; i++) {
			ScAgentId bidder = bidderList.GetAgentId(i).Value();
			if (bidder != auction.HighestBidder) {
				ScMutableBytes loser = bidders.GetBytes(bidder);
				Bid bid = new Bid(loser.Value());
				transfer(ctx, bidder, ScColor.IOTA, bid.Amount);
			}
		}

		// finalizeAuction request token was probably not confirmed yet
		transfer(ctx, ctx.ContractCreator(), ScColor.IOTA, ownerFee - 1);
		transfer(ctx, auction.HighestBidder, auction.Color, auction.NumTokens);
		transfer(ctx, auction.Creator, ScColor.IOTA, auction.Deposit + auction.HighestBid - ownerFee);
	}

	public static void funcPlaceBid(ScFuncContext ctx, FuncPlaceBidParams params) {
		long bidAmount = ctx.Incoming().Balance(ScColor.IOTA);
		ctx.Require(bidAmount > 0, "Missing bid amount");

		ScColor color = params.Color.Value();
		ScMutableMap state = ctx.State();
		ScMutableMap auctions = state.GetMap(Consts.VarAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
		ctx.Require(auctionInfo.Exists(), "Missing auction info");

		Auction auction = new Auction(auctionInfo.Value());
		ScMutableMap bidders = currentAuction.GetMap(Consts.VarBidders);
		ScMutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
		ScAgentId caller = ctx.Caller();
		ScMutableBytes bidder = bidders.GetBytes(caller);
		if (bidder.Exists()) {
			ctx.Log("Upped bid from: " + caller);
			Bid bid = new Bid(bidder.Value());
			bidAmount += bid.Amount;
			bid.Amount = bidAmount;
			bid.Timestamp = ctx.Timestamp();
			bidder.SetValue(bid.toBytes());
		} else {
			ctx.Require(bidAmount >= auction.MinimumBid, "Insufficient bid amount");
			ctx.Log("New bid from: " + caller);
			int index = bidderList.Length();
			bidderList.GetAgentId(index).SetValue(caller);
			Bid bid = new Bid();
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
		long ownerMargin = params.OwnerMargin.Value();
		if (ownerMargin < OwnerMarginMin) {
			ownerMargin = OwnerMarginMin;
		}
		if (ownerMargin > OwnerMarginMax) {
			ownerMargin = OwnerMarginMax;
		}
		ctx.State().GetInt(Consts.VarOwnerMargin).SetValue(ownerMargin);
		ctx.Log("Updated owner margin");
	}

	public static void funcStartAuction(ScFuncContext ctx, FuncStartAuctionParams params) {
		ScColor color = params.Color.Value();
		if (color == ScColor.IOTA || color == ScColor.MINT) {
			ctx.Panic("Reserved auction token color");
		}
		long numTokens = ctx.Incoming().Balance(color);
		if (numTokens == 0) {
			ctx.Panic("Missing auction tokens");
		}

		long minimumBid = params.MinimumBid.Value();

		// duration in minutes
		long duration = params.Duration.Value();
		if (duration == 0) {
			duration = DurationDefault;
		}
		if (duration < DurationMin) {
			duration = DurationMin;
		}
		if (duration > DurationMax) {
			duration = DurationMax;
		}

		String description = params.Description.Value();
		if (description == "") {
			description = "N/A";
		}
		if (description.length() > MaxDescriptionLength) {
			description = description.substring(0, MaxDescriptionLength) + "[...]";
		}

		ScMutableMap state = ctx.State();
		long ownerMargin = state.GetInt(Consts.VarOwnerMargin).Value();
		if (ownerMargin == 0) {
			ownerMargin = OwnerMarginDefault;
		}

		// need at least 1 iota to run SC
		long margin = minimumBid * ownerMargin / 1000;
		if (margin == 0) {
			margin = 1;
		}
		long deposit = ctx.Incoming().Balance(ScColor.IOTA);
		if (deposit < margin) {
			ctx.Panic("Insufficient deposit");
		}

		ScMutableMap auctions = state.GetMap(Consts.VarAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
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
			auction.HighestBidder = new ScAgentId(new byte[37]);
			auction.MinimumBid = minimumBid;
			auction.NumTokens = numTokens;
			auction.OwnerMargin = ownerMargin;
			auction.WhenStarted = ctx.Timestamp();
		}
		auctionInfo.SetValue(auction.toBytes());

		ScMutableMap finalizeParams = new ScMutableMap();
		finalizeParams.GetColor(Consts.VarColor).SetValue(auction.Color);
		PostRequestParams par = new PostRequestParams();
		par.ContractId = ctx.ContractId();
		par.Function = Consts.HFuncFinalizeAuction;
		par.Params = finalizeParams;
		par.Transfer = null;
		par.Delay = duration * 60;
		ctx.Post(par);
		ctx.Log("New auction started");
	}

	public static void viewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
		ScColor color = params.Color.Value();
		ScImmutableMap state = ctx.State();
		ScImmutableMap auctions = state.GetMap(Consts.VarAuctions);
		ScImmutableMap currentAuction = auctions.GetMap(color);
		ScImmutableBytes auctionInfo = currentAuction.GetBytes(Consts.VarInfo);
		if (!auctionInfo.Exists()) {
			ctx.Panic("Missing auction info");
		}

		Auction auction = new Auction(auctionInfo.Value());
		ScMutableMap results = ctx.Results();
		results.GetColor(Consts.VarColor).SetValue(auction.Color);
		results.GetAgentId(Consts.VarCreator).SetValue(auction.Creator);
		results.GetInt(Consts.VarDeposit).SetValue(auction.Deposit);
		results.GetString(Consts.VarDescription).SetValue(auction.Description);
		results.GetInt(Consts.VarDuration).SetValue(auction.Duration);
		results.GetInt(Consts.VarHighestBid).SetValue(auction.HighestBid);
		results.GetAgentId(Consts.VarHighestBidder).SetValue(auction.HighestBidder);
		results.GetInt(Consts.VarMinimumBid).SetValue(auction.MinimumBid);
		results.GetInt(Consts.VarNumTokens).SetValue(auction.NumTokens);
		results.GetInt(Consts.VarOwnerMargin).SetValue(auction.OwnerMargin);
		results.GetInt(Consts.VarWhenStarted).SetValue(auction.WhenStarted);

		ScImmutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(Consts.VarBidderList);
		results.GetInt(Consts.VarBidders).SetValue(bidderList.Length());
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
