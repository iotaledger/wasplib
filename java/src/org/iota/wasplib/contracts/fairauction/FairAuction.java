// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.builders.ScPostBuilder;
import org.iota.wasplib.client.context.ScFuncContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableAgentArray;
import org.iota.wasplib.client.immutable.ScImmutableBytes;
import org.iota.wasplib.client.immutable.ScImmutableColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableAgentArray;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableMap;

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

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddFunc("start_auction", FairAuction::startAuction);
		exports.AddFunc("finalize_auction", FairAuction::finalizeAuction);
		exports.AddFunc("place_bid", FairAuction::placeBid);
		exports.AddFunc("set_owner_margin", FairAuction::setOwnerMargin);
		exports.AddView("get_info", FairAuction::getInfo);
	}

	public static void startAuction(ScFuncContext sc) {
		ScImmutableMap params = sc.Params();
		ScImmutableColor colorParam = params.GetColor(KeyColor);
		if (!colorParam.Exists()) {
			sc.Panic("Missing auction token color");
		}
		ScColor color = colorParam.Value();
		if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
			sc.Panic("Reserved auction token color");
		}
		long numTokens = sc.Incoming().Balance(color);
		if (numTokens == 0) {
			sc.Panic("Missing auction tokens");
		}

		long minimumBid = params.GetInt(KeyMinimumBid).Value();
		if (minimumBid == 0) {
			sc.Panic("Missing minimum bid");
		}

		// duration in minutes
		long duration = params.GetInt(KeyDuration).Value();
		if (duration == 0) {
			duration = DurationDefault;
		}
		if (duration < DurationMin) {
			duration = DurationMin;
		}
		if (duration > DurationMax) {
			duration = DurationMax;
		}

		String description = params.GetString(KeyDescription).Value();
		if (description == "") {
			description = "N/A";
		}
		if (description.length() > MaxDescriptionLength) {
			description = description.substring(0, MaxDescriptionLength) + "[...]";
		}

		ScMutableMap state = sc.State();
		long ownerMargin = state.GetInt(KeyOwnerMargin).Value();
		if (ownerMargin == 0) {
			ownerMargin = OwnerMarginDefault;
		}

		// need at least 1 iota to run SC
		long margin = minimumBid * ownerMargin / 1000;
		if (margin == 0) {
			margin = 1;
		}
		long deposit = sc.Incoming().Balance(ScColor.IOTA);
		if (deposit < margin) {
			sc.Panic("Insufficient deposit");
		}

		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (auctionInfo.Exists()) {
			sc.Panic("Auction for this token color already exists");
		}

		AuctionInfo auction = new AuctionInfo();
		{
			auction.Creator = sc.Caller();
			auction.Color = color;
			auction.Deposit = deposit;
			auction.Description = description;
			auction.Duration = duration;
			auction.HighestBid = -1;
			auction.HighestBidder = new ScAgent(new byte[37]);
			auction.MinimumBid = minimumBid;
			auction.NumTokens = numTokens;
			auction.OwnerMargin = ownerMargin;
			auction.WhenStarted = sc.Timestamp();
		}
		auctionInfo.SetValue(AuctionInfo.encode(auction));

		ScPostBuilder finalizeRequest = sc.Post("finalize_auction");
		ScMutableMap finalizeParams = finalizeRequest.Params();
		finalizeParams.GetColor(KeyColor).SetValue(auction.Color);
		finalizeRequest.Post(duration * 60);
		sc.Log("New auction started");
	}

	public static void finalizeAuction(ScFuncContext sc) {
		// can only be sent by SC itself
		if (!sc.From(sc.ContractId().AsAgent())) {
			sc.Panic("Cancel spoofed request");
		}

		ScImmutableColor colorParam = sc.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			sc.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			sc.Panic("Missing auction info");
		}
		AuctionInfo auction = AuctionInfo.decode(auctionInfo.Value());
		long ownerFee;
		if (auction.HighestBid < 0) {
			sc.Log("No one bid on " + color);
			ownerFee = auction.MinimumBid * auction.OwnerMargin / 1000;
			if (ownerFee == 0) {
				ownerFee = 1;
			}
			// finalizeAuction request token was probably not confirmed yet
			transfer(sc, sc.ContractCreator(), ScColor.IOTA, ownerFee - 1);
			transfer(sc, auction.Creator, auction.Color, auction.NumTokens);
			transfer(sc, auction.Creator, ScColor.IOTA, auction.Deposit - ownerFee);
			return;
		}

		ownerFee = auction.HighestBid * auction.OwnerMargin / 1000;
		if (ownerFee == 0) {
			ownerFee = 1;
		}

		// return staked bids to losers
		ScMutableMap bidders = currentAuction.GetMap(KeyBidders);
		ScMutableAgentArray bidderList = currentAuction.GetAgentArray(KeyBidderList);
		int size = bidderList.Length();
		for (int i = 0; i < size; i++) {
			ScAgent bidder = bidderList.GetAgent(i).Value();
			if (!bidder.equals(auction.HighestBidder)) {
				ScMutableBytes loser = bidders.GetBytes(bidder);
				BidInfo bid = BidInfo.decode(loser.Value());
				transfer(sc, bidder, ScColor.IOTA, bid.Amount);
			}
		}

		// finalizeAuction request token was probably not confirmed yet
		transfer(sc, sc.ContractCreator(), ScColor.IOTA, ownerFee - 1);
		transfer(sc, auction.HighestBidder, auction.Color, auction.NumTokens);
		transfer(sc, auction.Creator, ScColor.IOTA, auction.Deposit + auction.HighestBid - ownerFee);
	}

	public static void placeBid(ScFuncContext sc) {
		long bidAmount = sc.Incoming().Balance(ScColor.IOTA);
		if (bidAmount == 0) {
			sc.Panic("Missing bid amount");
		}

		ScImmutableColor colorParam = sc.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			sc.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableMap auctions = state.GetMap(KeyAuctions);
		ScMutableMap currentAuction = auctions.GetMap(color);
		ScMutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			sc.Panic("Missing auction info");
		}

		AuctionInfo auction = AuctionInfo.decode(auctionInfo.Value());
		ScMutableMap bidders = currentAuction.GetMap(KeyBidders);
		ScMutableAgentArray bidderList = currentAuction.GetAgentArray(KeyBidderList);
		ScAgent caller = sc.Caller();
		ScMutableBytes bidder = bidders.GetBytes(caller);
		if (bidder.Exists()) {
			sc.Log("Upped bid from: " + caller);
			BidInfo bid = BidInfo.decode(bidder.Value());
			bidAmount += bid.Amount;
			bid.Amount = bidAmount;
			bid.Timestamp = sc.Timestamp();
			bidder.SetValue(BidInfo.encode(bid));
		} else {
			if (bidAmount < auction.MinimumBid) {
				sc.Panic("Insufficient bid amount");
			}
			sc.Log("New bid from: " + caller);
			int index = bidderList.Length();
			bidderList.GetAgent(index).SetValue(caller);
			BidInfo bid = new BidInfo();
			{
				bid.Index = index;
				bid.Amount = bidAmount;
				bid.Timestamp = sc.Timestamp();
			}
			bidder.SetValue(BidInfo.encode(bid));
		}
		if (bidAmount > auction.HighestBid) {
			sc.Log("New highest bidder");
			auction.HighestBid = bidAmount;
			auction.HighestBidder = caller;
			auctionInfo.SetValue(AuctionInfo.encode(auction));
		}
	}

	public static void setOwnerMargin(ScFuncContext sc) {
		// can only be sent by SC creator
		if (!sc.From(sc.ContractCreator())) {
			sc.Panic("Cancel spoofed request");
		}

		long ownerMargin = sc.Params().GetInt(KeyOwnerMargin).Value();
		if (ownerMargin < OwnerMarginMin) {
			ownerMargin = OwnerMarginMin;
		}
		if (ownerMargin > OwnerMarginMax) {
			ownerMargin = OwnerMarginMax;
		}
		sc.State().GetInt(KeyOwnerMargin).SetValue(ownerMargin);
		sc.Log("Updated owner margin");
	}

	public static void getInfo(ScViewContext sc) {
		ScImmutableColor colorParam = sc.Params().GetColor(KeyColor);
		if (!colorParam.Exists()) {
			sc.Panic("Missing token color");
		}
		ScColor color = colorParam.Value();

		ScImmutableMap state = sc.State();
		ScImmutableMap auctions = state.GetMap(KeyAuctions);
		ScImmutableMap currentAuction = auctions.GetMap(color);
		ScImmutableBytes auctionInfo = currentAuction.GetBytes(KeyInfo);
		if (!auctionInfo.Exists()) {
			sc.Panic("Missing auction info");
		}

		AuctionInfo auction = AuctionInfo.decode(auctionInfo.Value());
		ScMutableMap results = sc.Results();
		results.GetColor(KeyColor).SetValue(auction.Color);
		results.GetAgent(KeyCreator).SetValue(auction.Creator);
		results.GetInt(KeyDeposit).SetValue(auction.Deposit);
		results.GetString(KeyDescription).SetValue(auction.Description);
		results.GetInt(KeyDuration).SetValue(auction.Duration);
		results.GetInt(KeyHighestBid).SetValue(auction.HighestBid);
		results.GetAgent(KeyHighestBidder).SetValue(auction.HighestBidder);
		results.GetInt(KeyMinimumBid).SetValue(auction.MinimumBid);
		results.GetInt(KeyNumTokens).SetValue(auction.NumTokens);
		results.GetInt(KeyOwnerMargin).SetValue(auction.OwnerMargin);
		results.GetInt(KeyWhenStarted).SetValue(auction.WhenStarted);

		ScImmutableAgentArray bidderList = currentAuction.GetAgentArray(KeyBidderList);
		results.GetInt(KeyBidders).SetValue(bidderList.Length());
	}

	private static void transfer(ScFuncContext sc, ScAgent agent, ScColor color, long amount) {
		if (!agent.IsAddress()) {
			// not an address, deposit into account on chain
			sc.Transfer(agent, color, amount);
			return;
		}

		// send to original Tangle address
		sc.TransferToAddress(agent.Address()).Transfer(color, amount).Send();
	}
}
