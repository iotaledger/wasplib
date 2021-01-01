// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.builders.ScPostBuilder;
import org.iota.wasplib.client.context.ScBalances;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableColor;
import org.iota.wasplib.client.immutable.ScImmutableColorArray;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableAgentArray;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class FairAuction {
	private static final Key keyAuctions = new Key("auctions");
	private static final Key keyBidders = new Key("bidders");
	private static final Key keyBidderList = new Key("bidder_list");
	private static final Key keyColor = new Key("color");
	private static final Key keyDescription = new Key("description");
	private static final Key keyDuration = new Key("duration");
	private static final Key keyInfo = new Key("info");
	private static final Key keyMinimumBid = new Key("minimum");
	private static final Key keyOwnerMargin = new Key("owner_margin");

	private static final int durationDefault = 60;
	private static final int durationMin = 1;
	private static final int durationMax = 120;
	private static final int maxDescriptionLength = 150;
	private static final int ownerMarginDefault = 50;
	private static final int ownerMarginMin = 5;
	private static final int ownerMarginMax = 100;

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("start_auction", FairAuction::startAuction);
		exports.AddCall("finalize_auction", FairAuction::finalizeAuction);
		exports.AddCall("place_bid", FairAuction::placeBid);
		exports.AddCall("set_owner_margin", FairAuction::setOwnerMargin);
	}

	public static void startAuction(ScCallContext sc) {
		long deposit = sc.Incoming().Balance(ScColor.IOTA);
		if (deposit < 1) {
			sc.Panic("Empty deposit...");
		}

		ScMutableMap state = sc.State();
		long ownerMargin = state.GetInt(keyOwnerMargin).Value();
		if (ownerMargin == 0) {
			ownerMargin = ownerMarginDefault;
		}

		ScImmutableMap params = sc.Params();
		ScImmutableColor colorParam = params.GetColor(keyColor);
		if (!colorParam.Exists()) {
			refund(sc, deposit / 2, "Missing token color...");
			return;
		}
		ScColor color = colorParam.Value();

		if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
			refund(sc, deposit / 2, "Reserved token color...");
			return;
		}

		long numTokens = sc.Incoming().Balance(color);
		if (numTokens == 0) {
			refund(sc, deposit / 2, "Auction tokens missing from request...");
			return;
		}

		long minimumBid = params.GetInt(keyMinimumBid).Value();
		if (minimumBid == 0) {
			refund(sc, deposit / 2, "Missing minimum bid...");
			return;
		}

		// need at least 1 iota to run SC
		long margin = minimumBid * ownerMargin / 1000;
		if (margin == 0) {
			margin = 1;
		}
		if (deposit < margin) {
			refund(sc, deposit / 2, "Insufficient deposit...");
			return;
		}

		// duration in minutes
		long duration = params.GetInt(keyDuration).Value();
		if (duration == 0) {
			duration = durationDefault;
		}
		if (duration < durationMin) {
			duration = durationMin;
		}
		if (duration > durationMax) {
			duration = durationMax;
		}

		String description = params.GetString(keyDescription).Value();
		if (description.equals("")) {
			description = "N/A";
		}
		if (description.length() > maxDescriptionLength) {
			description = description.substring(0, maxDescriptionLength) + "[...]";
		}

		ScMutableMap auctions = state.GetMap(keyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes currentInfo = currentAuction.GetBytes(keyInfo);
		if (currentInfo.Exists()) {
			refund(sc, deposit / 2, "Auction for this token already exists...");
			return;
		}

		AuctionInfo auction = new AuctionInfo();
		{
			auction.auctionOwner = sc.Caller();
			auction.color = color;
			auction.deposit = deposit;
			auction.description = description;
			auction.duration = duration;
			auction.highestBid = -1;
			auction.highestBidder = ScAgent.NONE;
			auction.minimumBid = minimumBid;
			auction.numTokens = numTokens;
			auction.ownerMargin = ownerMargin;
			auction.whenStarted = sc.Timestamp();
		}
		currentInfo.SetValue(AuctionInfo.encode(auction));

		ScPostBuilder finalizeRequest = sc.Post("finalize_auction");
		ScMutableMap finalizeParams = finalizeRequest.Params();
		finalizeParams.GetColor(keyColor).SetValue(auction.color);
		finalizeRequest.Post(duration * 60);
		sc.Log("New auction started...");
	}

	public static void finalizeAuction(ScCallContext sc) {
		// can only be sent by SC itself
		if (!sc.From(sc.Contract().Id())) {
			sc.Panic("Cancel spoofed request");
		}

		ScImmutableColor colorParam = sc.Params().GetColor(keyColor);
		if (!colorParam.Exists()) {
			sc.Panic("Internal inconsistency: missing color");
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableMap auctions = state.GetMap(keyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes currentInfo = currentAuction.GetBytes(keyInfo);
		if (!currentInfo.Exists()) {
			sc.Panic("Internal inconsistency: missing auction info");
		}
		AuctionInfo auction = AuctionInfo.decode(currentInfo.Value());
		if (auction.highestBid < 0) {
			sc.Log("No one bid on " + color);
			long ownerFee = auction.minimumBid * auction.ownerMargin / 1000;
			if (ownerFee == 0) {
				ownerFee = 1;
			}
			// finalizeAuction request token was probably not confirmed yet
			sc.Transfer(sc.Contract().Creator(), ScColor.IOTA, ownerFee - 1);
			sc.Transfer(auction.auctionOwner, auction.color, auction.numTokens);
			sc.Transfer(auction.auctionOwner, ScColor.IOTA, auction.deposit - ownerFee);
			return;
		}

		long ownerFee = auction.highestBid * auction.ownerMargin / 1000;
		if (ownerFee == 0) {
			ownerFee = 1;
		}

		// return staked bids to losers
		ScMutableMap bidders = currentAuction.GetMap(keyBidders);
		ScMutableAgentArray bidderList = currentAuction.GetAgentArray(keyBidderList);
		int size = bidderList.Length();
		for (int i = 0; i < size; i++) {
			ScAgent bidder = bidderList.GetAgent(i).Value();
			if (!bidder.equals(auction.highestBidder)) {
				ScMutableBytes loser = bidders.GetBytes(bidder);
				BidInfo bid = BidInfo.decode(loser.Value());
				sc.Transfer(bidder, ScColor.IOTA, bid.amount);
			}
		}

		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Creator(), ScColor.IOTA, ownerFee - 1);
		sc.Transfer(auction.highestBidder, auction.color, auction.numTokens);
		sc.Transfer(auction.auctionOwner, ScColor.IOTA, auction.deposit + auction.highestBid - ownerFee);
	}

	public static void placeBid(ScCallContext sc) {
		long bidAmount = sc.Incoming().Balance(ScColor.IOTA);
		if (bidAmount == 0) {
			sc.Panic("Insufficient bid amount");
		}

		ScImmutableColor colorParam = sc.Params().GetColor(keyColor);
		if (!colorParam.Exists()) {
			refund(sc, bidAmount, "Missing token color");
			return;
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableMap auctions = state.GetMap(keyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes currentInfo = currentAuction.GetBytes(keyInfo);
		if (!currentInfo.Exists()) {
			refund(sc, bidAmount, "Missing auction");
			return;
		}

		AuctionInfo auction = AuctionInfo.decode(currentInfo.Value());
		ScMutableMap bidders = currentAuction.GetMap(keyBidders);
		ScMutableAgentArray bidderList = currentAuction.GetAgentArray(keyBidderList);
		ScAgent caller = sc.Caller();
		ScMutableBytes bidder = bidders.GetBytes(caller);
		BidInfo bid;
		if (bidder.Exists()) {
			sc.Log("Upped bid from: " + caller);
			bid = BidInfo.decode(bidder.Value());
			bidAmount += bid.amount;
			bid.amount = bidAmount;
			bid.timestamp = sc.Timestamp();
			bidder.SetValue(BidInfo.encode(bid));
		} else {
			sc.Log("New bid from: " + caller);
			int index = bidderList.Length();
			bidderList.GetAgent(index).SetValue(caller);
			bid = new BidInfo();
			{
				bid.index = index;
				bid.amount = bidAmount;
				bid.timestamp = sc.Timestamp();
			}
			bidder.SetValue(BidInfo.encode(bid));
		}
		if (bidAmount > auction.highestBid) {
			sc.Log("New highest bidder...");
			auction.highestBid = bidAmount;
			auction.highestBidder = caller;
			currentInfo.SetValue(AuctionInfo.encode(auction));
		}
	}

	public static void setOwnerMargin(ScCallContext sc) {
		// can only be sent by SC creator
		if (!sc.From(sc.Contract().Creator())) {
			sc.Panic("Cancel spoofed request");
		}

		long ownerMargin = sc.Params().GetInt(keyOwnerMargin).Value();
		if (ownerMargin < ownerMarginMin) {
			ownerMargin = ownerMarginMin;
		}
		if (ownerMargin > ownerMarginMax) {
			ownerMargin = ownerMarginMax;
		}
		sc.State().GetInt(keyOwnerMargin).SetValue(ownerMargin);
		sc.Log("Updated owner margin...");
	}

	public static void refund(ScCallContext sc, long amount, String reason) {
		sc.Log(reason);
		ScAgent caller = sc.Caller();
		if (amount != 0) {
			sc.Transfer(caller, ScColor.IOTA, amount);
		}
		ScBalances incoming = sc.Incoming();
		long deposit = incoming.Balance(ScColor.IOTA);
		if (deposit - amount != 0) {
			sc.Transfer(sc.Contract().Creator(), ScColor.IOTA, deposit - amount);
		}

		// refund all other token colors, don't keep tokens that were to be auctioned
		ScImmutableColorArray colors = incoming.Colors();
		int size = colors.Length();
		for (int i = 0; i < size; i++) {
			ScColor color = colors.GetColor(i).Value();
			if (color != ScColor.IOTA) {
				sc.Transfer(caller, color, incoming.Balance(color));
			}
		}
	}
}
