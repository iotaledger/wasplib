// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.request.ScPostInfo;

import java.util.ArrayList;

public class FairAuction {
	private static final int durationDefault = 60;
	private static final int durationMin = 1;
	private static final int durationMax = 120;
	private static final int maxDescriptionLength = 150;
	private static final int ownerMarginDefault = 50;
	private static final int ownerMarginMin = 5;
	private static final int ownerMarginMax = 100;

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("startAuction", FairAuction::startAuction);
		exports.AddCall("finalizeAuction", FairAuction::finalizeAuction);
		exports.AddCall("placeBid", FairAuction::placeBid);
		exports.AddCall("setOwnerMargin", FairAuction::setOwnerMargin);
	}

	public static void startAuction(ScCallContext sc) {
		long deposit = sc.Balances().Balance(ScColor.IOTA);
		if (deposit < 1) {
			sc.Log("Empty deposit...");
			return;
		}

		ScMutableMap state = sc.State();
		long ownerMargin = state.GetInt("ownerMargin").Value();
		if (ownerMargin == 0) {
			ownerMargin = ownerMarginDefault;
		}

		ScImmutableMap params = sc.Params();
		ScImmutableColor colorParam = params.GetColor("color");
		if (!colorParam.Exists()) {
			refund(sc, deposit / 2, "Missing token color...");
			return;
		}
		ScColor color = colorParam.Value();

		if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
			refund(sc, deposit / 2, "Reserved token color...");
			return;
		}

		long numTokens = sc.Balances().Balance(color);
		if (numTokens == 0) {
			refund(sc, deposit / 2, "Auction tokens missing from request...");
			return;
		}

		long minimumBid = params.GetInt("minimum").Value();
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
		long duration = params.GetInt("duration").Value();
		if (duration == 0) {
			duration = durationDefault;
		}
		if (duration < durationMin) {
			duration = durationMin;
		}
		if (duration > durationMax) {
			duration = durationMax;
		}

		String description = params.GetString("dscr").Value();
		if (description.isEmpty()) {
			description = "N/A";
		}
		if (description.length() > maxDescriptionLength) {
			description = description.substring(0, maxDescriptionLength) + "[...]";
		}

		ScMutableKeyMap auctions = state.GetKeyMap("auctions");
		ScMutableBytes currentAuction = auctions.GetBytes(color.toBytes());
		if (currentAuction.Value().length != 0) {
			refund(sc, deposit / 2, "Auction for this token already exists...");
			return;
		}

		AuctionInfo auction = new AuctionInfo();
		auction.color = color;
		auction.numTokens = numTokens;
		auction.minimumBid = minimumBid;
		auction.description = description;
		auction.whenStarted = sc.Timestamp();
		auction.duration = duration;
		auction.auctionOwner = sc.Caller();
		auction.deposit = deposit;
		auction.ownerMargin = ownerMargin;
		byte[] bytes2 = encodeAuctionInfo(auction);
		currentAuction.SetValue(bytes2);

		ScPostInfo finalizeRequest = sc.Post("finalizeAuction");
		ScMutableMap finalizeParams = finalizeRequest.Params();
		finalizeParams.GetColor("color").SetValue(auction.color);
		finalizeRequest.Post(auction.duration * 60);
		sc.Log("New auction started...");
	}

	public static void finalizeAuction(ScCallContext sc) {
		// can only be sent by SC itself
		if (!sc.From(sc.Contract().Id())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScImmutableColor colorParam = sc.Params().GetColor("color");
		if (!colorParam.Exists()) {
			sc.Log("INTERNAL INCONSISTENCY: missing color");
			return;
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableKeyMap auctions = state.GetKeyMap("auctions");
		ScMutableBytes currentAuction = auctions.GetBytes(color.toBytes());
		byte[] bytes2 = currentAuction.Value();
		if (bytes2.length == 0) {
			sc.Log("INTERNAL INCONSISTENCY missing auction info");
			return;
		}
		AuctionInfo auction = decodeAuctionInfo(bytes2);
		long ownerFee;
		if (auction.bids.size() == 0) {
			sc.Log("No one bid on " + color);
			ownerFee = auction.minimumBid * auction.ownerMargin / 1000;
			if (ownerFee == 0) {
				ownerFee = 1;
			}
			// finalizeAuction request token was probably not confirmed yet
			sc.Transfer(sc.Contract().Owner(), ScColor.IOTA, ownerFee - 1);
			sc.Transfer(auction.auctionOwner, auction.color, auction.numTokens);
			sc.Transfer(auction.auctionOwner, ScColor.IOTA, auction.deposit - ownerFee);
			return;
		}

		BidInfo winner = new BidInfo();
		for (BidInfo bidder : auction.bids) {
			if (bidder.amount >= winner.amount) {
				if (bidder.amount > winner.amount || bidder.when < winner.when) {
					winner = bidder;
				}
			}
		}
		ownerFee = winner.amount * auction.ownerMargin / 1000;
		if (ownerFee == 0) {
			ownerFee = 1;
		}

		// return staked bids to losers
		for (BidInfo bidder : auction.bids) {
			if (bidder != winner) {
				sc.Transfer(bidder.bidder, ScColor.IOTA, bidder.amount);
			}
		}

		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Owner(), ScColor.IOTA, ownerFee - 1);
		sc.Transfer(winner.bidder, auction.color, auction.numTokens);
		sc.Transfer(auction.auctionOwner, ScColor.IOTA, auction.deposit + winner.amount - ownerFee);
	}

	public static void placeBid(ScCallContext sc) {
		long bidAmount = sc.Balances().Balance(ScColor.IOTA);
		if (bidAmount == 0) {
			sc.Log("Insufficient bid amount");
			return;
		}

		ScImmutableColor colorParam = sc.Params().GetColor("color");
		if (!colorParam.Exists()) {
			refund(sc, bidAmount, "Missing token color");
			return;
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableKeyMap auctions = state.GetKeyMap("auctions");
		ScMutableBytes currentAuction = auctions.GetBytes(color.toBytes());
		byte[] bytes = currentAuction.Value();
		if (bytes.length == 0) {
			refund(sc, bidAmount, "Missing auction");
			return;
		}

		ScAgent caller = sc.Caller();
		AuctionInfo auction = decodeAuctionInfo(bytes);
		BidInfo bid = null;
		for (BidInfo bidder : auction.bids) {
			if (bidder.bidder == caller) {
				bid = bidder;
				break;
			}
		}
		if (bid == null) {
			sc.Log("New bid from: " + caller);
			bid = new BidInfo();
			bid.bidder = caller;
			auction.bids.add(bid);
		}
		bid.amount += bidAmount;
		bid.when = sc.Timestamp();

		bytes = encodeAuctionInfo(auction);
		currentAuction.SetValue(bytes);
		sc.Log("Updated auction with bid...");
	}

	public static void setOwnerMargin(ScCallContext sc) {
		// can only be sent by SC owner
		if (!sc.From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long ownerMargin = sc.Params().GetInt("ownerMargin").Value();
		if (ownerMargin < ownerMarginMin) {
			ownerMargin = ownerMarginMin;
		}
		if (ownerMargin > ownerMarginMax) {
			ownerMargin = ownerMarginMax;
		}
		sc.State().GetInt("ownerMargin").SetValue(ownerMargin);
		sc.Log("Updated owner margin...");
	}

	public static AuctionInfo decodeAuctionInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		AuctionInfo auction = new AuctionInfo();
		auction.color = decoder.Color();
		auction.numTokens = decoder.Int();
		auction.minimumBid = decoder.Int();
		auction.description = decoder.String();
		auction.whenStarted = decoder.Int();
		auction.duration = decoder.Int();
		auction.auctionOwner = decoder.Agent();
		auction.deposit = decoder.Int();
		auction.ownerMargin = decoder.Int();
		return auction;
	}

	public static BidInfo decodeBidInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		BidInfo bid = new BidInfo();
		bid.bidder = decoder.Agent();
		bid.amount = decoder.Int();
		bid.when = decoder.Int();
		return bid;
	}

	public static byte[] encodeAuctionInfo(AuctionInfo auction) {
		BytesEncoder encoder = new BytesEncoder().
				Color(auction.color).
				Int(auction.numTokens).
				Int(auction.minimumBid).
				String(auction.description).
				Int(auction.whenStarted).
				Int(auction.duration).
				Agent(auction.auctionOwner).
				Int(auction.deposit).
				Int(auction.ownerMargin).
				Int(auction.bids.size());
		for (BidInfo bid : auction.bids) {
			byte[] bytes = encodeBidInfo(bid);
			encoder.Bytes(bytes);
		}
		return encoder.Data();
	}

	public static byte[] encodeBidInfo(BidInfo bid) {
		return new BytesEncoder().
				Agent(bid.bidder).
				Int(bid.amount).
				Int(bid.when).
				Data();
	}

	public static void refund(ScCallContext sc, long amount, String reason) {
		sc.Log(reason);
		ScAgent caller = sc.Caller();
		if (amount != 0) {
			sc.Transfer(caller, ScColor.IOTA, amount);
		}
		long deposit = sc.Balances().Balance(ScColor.IOTA);
		if (deposit - amount != 0) {
			sc.Transfer(sc.Contract().Owner(), ScColor.IOTA, deposit - amount);
		}

//TODO
//		// refund all other token colors, don't keep tokens that were to be auctioned
//		ScImmutableColorArray colors = request.Colors();
//		int items = colors.Length();
//		for (int i = 0; i < items; i++) {
//			ScColor color = colors.GetColor(i).Value();
//			if (!color.equals(ScColor.IOTA)) {
//				sc.Transfer(caller, color, sc.Balances().Balance(color));
//			}
//		}
	}

	public static class AuctionInfo {
		// color of tokens for sale
		ScColor color;
		// number of tokens for sale
		long numTokens;
		// minimum bid. Set by the auction initiator
		long minimumBid;
		// any text, like "AuctionOwner of the token have a right to call me for a date". Set by auction initiator
		String description;
		// timestamp when auction started
		long whenStarted;
		// duration of the auctions in minutes. Should be >= MinAuctionDurationMinutes
		long duration;
		// address which issued StartAuction transaction
		ScAgent auctionOwner;
		// deposit by the auction owner. Iotas sent by the auction owner together with the tokens for sale in the same
		// transaction.
		long deposit;
		// AuctionOwner's margin in promilles, taken at the moment of creation of smart contract
		long ownerMargin;
		// list of bids to the auction
		ArrayList<BidInfo> bids = new ArrayList<>();
	}

	public static class BidInfo {
		// originator of the bid
		ScAgent bidder;
		// the amount is a cumulative sum of all bids from the same bidder
		long amount;
		// most recent bid update time
		long when;
	}
}
