// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.fairauction;

import org.iota.wasp.contracts.fairauction.lib.*;
import org.iota.wasp.contracts.fairauction.types.*;
import org.iota.wasp.wasmlib.builders.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class FairAuction {
	private static final Key KeyAuctions = new Key("auctions");
	private static final Key KeyBidders = new Key("bidders");
	private static final Key KeyBidderList = new Key("bidder_list");
	private static final Key KeyColor = new Key("color");
	private static final Key KeyCreator = new Key("creator");
	private static final Key KeyDeposit = new Key("deposit");
	private static final Key KeyDescription = new Key("description");
	private static final Key KeyDuration = new Key("duration");
	private static final Key KeyHighestBid = new Key("highest_bid");
	private static final Key KeyHighestBidder = new Key("highest_bidder");
	private static final Key KeyInfo = new Key("info");
	private static final Key KeyMinimumBid = new Key("minimum");
	private static final Key KeyNumTokens = new Key("num_tokens");
	private static final Key KeyOwnerMargin = new Key("owner_margin");
	private static final Key KeyWhenStarted = new Key("when_started");

	private static final int DurationDefault = 60;
	private static final int DurationMin = 1;
	private static final int DurationMax = 120;
	private static final int MaxDescriptionLength = 150;
	private static final int OwnerMarginDefault = 50;
	private static final int OwnerMarginMin = 5;
	private static final int OwnerMarginMax = 100;

	public static void FuncFinalizeAuction(ScFuncContext ctx, FuncFinalizeAuctionParams params) {
		// can only be sent by SC itself
		if (!ctx.Caller().equals(ctx.ContractId().AsAgentId())) {
			ctx.Panic("Cancel spoofed request");
		}

		ScImmutableColor colorParam = ctx.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			ctx.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = ctx.State();
		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			ctx.Panic("Missing auction info");
		}
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
		ScMutableMap bidders = currentAuction.GetMap(KeyBidders);
		ScMutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(KeyBidderList);
		int size = bidderList.Length();
		for (int i = 0; i < size; i++) {
			ScAgentId bidder = bidderList.GetAgentId(i).Value();
			if (!bidder.equals(auction.HighestBidder)) {
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

	public static void FuncPlaceBid(ScFuncContext ctx, FuncPlaceBidParams params) {
		long bidAmount = ctx.Incoming().Balance(ScColor.IOTA);
		if (bidAmount == 0) {
			ctx.Panic("Missing bid amount");
		}

		ScImmutableColor colorParam = ctx.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			ctx.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = ctx.State();
		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			ctx.Panic("Missing auction info");
		}

		Auction auction = new Auction(auctionInfo.Value());
		ScMutableMap bidders = currentAuction.GetMap(KeyBidders);
		ScMutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(KeyBidderList);
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
			if (bidAmount < auction.MinimumBid) {
				ctx.Panic("Insufficient bid amount");
			}
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

	public static void FuncSetOwnerMargin(ScFuncContext ctx, FuncSetOwnerMarginParams params) {
		// can only be sent by SC creator
		if (!ctx.Caller().equals(ctx.ContractCreator())) {
			ctx.Panic("Cancel spoofed request");
		}

		long ownerMargin = ctx.Params().GetInt(KeyOwnerMargin).Value();
		if (ownerMargin < OwnerMarginMin) {
			ownerMargin = OwnerMarginMin;
		}
		if (ownerMargin > OwnerMarginMax) {
			ownerMargin = OwnerMarginMax;
		}
		ctx.State().GetInt(KeyOwnerMargin).SetValue(ownerMargin);
		ctx.Log("Updated owner margin");
	}

	public static void FuncStartAuction(ScFuncContext ctx, FuncStartAuctionParams params) {
		ScImmutableMap p = ctx.Params();
		ScImmutableColor colorParam = p.GetColor(KeyColor);
		if (!colorParam.Exists()) {
			ctx.Panic("Missing auction token color");
		}
		ScColor color = colorParam.Value();
		if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
			ctx.Panic("Reserved auction token color");
		}
		long numTokens = ctx.Incoming().Balance(color);
		if (numTokens == 0) {
			ctx.Panic("Missing auction tokens");
		}

		long minimumBid = p.GetInt(KeyMinimumBid).Value();
		if (minimumBid == 0) {
			ctx.Panic("Missing minimum bid");
		}

		// duration in minutes
		long duration = p.GetInt(KeyDuration).Value();
		if (duration == 0) {
			duration = DurationDefault;
		}
		if (duration < DurationMin) {
			duration = DurationMin;
		}
		if (duration > DurationMax) {
			duration = DurationMax;
		}

		String description = p.GetString(KeyDescription).Value();
		if (description == "") {
			description = "N/A";
		}
		if (description.length() > MaxDescriptionLength) {
			description = description.substring(0, MaxDescriptionLength) + "[...]";
		}

		ScMutableMap state = ctx.State();
		long ownerMargin = state.GetInt(KeyOwnerMargin).Value();
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

		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
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

		ScPostBuilder finalizeRequest = ctx.Post("finalize_auction");
		ScMutableMap finalizeParams = finalizeRequest.Params();
		finalizeParams.GetColor(KeyColor).SetValue(auction.Color);
		finalizeRequest.Post(duration * 60);
		ctx.Log("New auction started");
	}

	public static void ViewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
		ScImmutableColor colorParam = ctx.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			ctx.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScImmutableMap state = ctx.State();
		ScImmutableMap auctions = state.GetMap(KeyAuctions);
		ScImmutableMap currentAuction = auctions.GetMap(color);
		ScImmutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			ctx.Panic("Missing auction info");
		}

		Auction auction = new Auction(auctionInfo.Value());
		ScMutableMap results = ctx.Results();
		results.GetColor(KeyColor).SetValue(auction.Color);
		results.GetAgentId(KeyCreator).SetValue(auction.Creator);
		results.GetInt(KeyDeposit).SetValue(auction.Deposit);
		results.GetString(KeyDescription).SetValue(auction.Description);
		results.GetInt(KeyDuration).SetValue(auction.Duration);
		results.GetInt(KeyHighestBid).SetValue(auction.HighestBid);
		results.GetAgentId(KeyHighestBidder).SetValue(auction.HighestBidder);
		results.GetInt(KeyMinimumBid).SetValue(auction.MinimumBid);
		results.GetInt(KeyNumTokens).SetValue(auction.NumTokens);
		results.GetInt(KeyOwnerMargin).SetValue(auction.OwnerMargin);
		results.GetInt(KeyWhenStarted).SetValue(auction.WhenStarted);

		ScImmutableAgentIdArray bidderList = currentAuction.GetAgentIdArray(KeyBidderList);
		results.GetInt(KeyBidders).SetValue(bidderList.Length());
	}

	private static void transfer(ScFuncContext ctx, ScAgentId agent, ScColor color, long amount) {
		if (!agent.IsAddress()) {
			// not an address, deposit into account on chain
			ctx.Transfer(agent, color, amount);
			return;
		}

		// send to original Tangle address
		ctx.TransferToAddress(agent.Address()).Transfer(color, amount).Send();
	}
}
