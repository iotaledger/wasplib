package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScExports;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableColor;
import org.iota.wasplib.client.immutable.ScImmutableColorArray;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

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
		exports.Add("startAuction");
		exports.Add("finalizeAuction");
		exports.Add("placeBid");
		exports.AddProtected("setOwnerMargin");
	}

	//export startAuction
	public static void startAuction() {
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		long deposit = request.Balance(ScColor.IOTA);
		if (deposit < 1) {
			sc.Log("Empty deposit...");
			return;
		}

		ScMutableMap state = sc.State();
		long ownerMargin = state.GetInt("ownerMargin").Value();
		if (ownerMargin == 0) {
			ownerMargin = ownerMarginDefault;
		}

		ScImmutableMap params = request.Params();
		ScImmutableColor colorParam = params.GetColor("color");
		if (!colorParam.Exists()) {
			refund(deposit / 2, "Missing token color...");
			return;
		}
		ScColor color = colorParam.Value();

		if (color.equals(ScColor.IOTA) || color.equals(ScColor.MINT)) {
			refund(deposit / 2, "Reserved token color...");
			return;
		}

		long numTokens = request.Balance(color);
		if (numTokens == 0) {
			refund(deposit / 2, "Auction tokens missing from request...");
			return;
		}

		long minimumBid = params.GetInt("minimum").Value();
		if (minimumBid == 0) {
			refund(deposit / 2, "Missing minimum bid...");
			return;
		}

		// need at least 1 iota to run SC
		long margin = minimumBid * ownerMargin / 1000;
		if (margin == 0) {
			margin = 1;
		}
		if (deposit < margin) {
			refund(deposit / 2, "Insufficient deposit...");
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
			refund(deposit / 2, "Auction for this token already exists...");
			return;
		}

		AuctionInfo auction = new AuctionInfo();
		auction.color = color;
		auction.numTokens = numTokens;
		auction.minimumBid = minimumBid;
		auction.description = description;
		auction.whenStarted = request.Timestamp();
		auction.duration = duration;
		auction.auctionOwner = request.Address();
		auction.deposit = deposit;
		auction.ownerMargin = ownerMargin;
		byte[] bytes2 = encodeAuctionInfo(auction);
		currentAuction.SetValue(bytes2);

		ScMutableMap finalizeParams = sc.PostRequest(sc.Contract().Address(), "finalizeAuction", auction.duration * 60);
		finalizeParams.GetColor("color").SetValue(auction.color);
		sc.Log("New auction started...");
	}

	//export finalizeAuction
	public static void finalizeAuction() {
		// can only be sent by SC itself
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		if ((!request.From(sc.Contract().Address()))) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScImmutableColor colorParam = request.Params().GetColor("color");
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
				sc.Transfer(bidder.address, ScColor.IOTA, bidder.amount);
			}
		}

		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Owner(), ScColor.IOTA, ownerFee - 1);
		sc.Transfer(winner.address, auction.color, auction.numTokens);
		sc.Transfer(auction.auctionOwner, ScColor.IOTA, auction.deposit + winner.amount - ownerFee);
	}

	//export placeBid
	public static void placeBid() {
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		long bidAmount = request.Balance(ScColor.IOTA);
		if (bidAmount == 0) {
			sc.Log("Insufficient bid amount");
			return;
		}

		ScImmutableColor colorParam = request.Params().GetColor("color");
		if (!colorParam.Exists()) {
			refund(bidAmount, "Missing token color");
			return;
		}
		ScColor color = colorParam.Value();

		ScMutableMap state = sc.State();
		ScMutableKeyMap auctions = state.GetKeyMap("auctions");
		ScMutableBytes currentAuction = auctions.GetBytes(color.toBytes());
		byte[] bytes = currentAuction.Value();
		if (bytes.length == 0) {
			refund(bidAmount, "Missing auction");
			return;
		}

		ScAddress sender = request.Address();
		AuctionInfo auction = decodeAuctionInfo(bytes);
		BidInfo bid = null;
		for (BidInfo bidder : auction.bids) {
			if (bidder.address == sender) {
				bid = bidder;
				break;
			}
		}
		if (bid == null) {
			sc.Log("New bid from: " + sender);
			bid = new BidInfo();
			bid.address = sender;
			auction.bids.add(bid);
		}
		bid.amount += bidAmount;
		bid.when = request.Timestamp();

		bytes = encodeAuctionInfo(auction);
		currentAuction.SetValue(bytes);
		sc.Log("Updated auction with bid...");
	}

	//export setOwnerMargin
	public static void setOwnerMargin() {
		// can only be sent by SC owner
		ScContext sc = new ScContext();
		if (!sc.Request().From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long ownerMargin = sc.Request().Params().GetInt("ownerMargin").Value();
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
		auction.auctionOwner = decoder.Address();
		auction.deposit = decoder.Int();
		auction.ownerMargin = decoder.Int();
		return auction;
	}

	public static BidInfo decodeBidInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		BidInfo bid = new BidInfo();
		bid.address = decoder.Address();
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
				Address(auction.auctionOwner).
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
				Address(bid.address).
				Int(bid.amount).
				Int(bid.when).
				Data();
	}

	public static void refund(long amount, String reason) {
		ScContext sc = new ScContext();
		sc.Log(reason);
		ScRequest request = sc.Request();
		ScAddress sender = request.Address();
		if (amount != 0) {
			sc.Transfer(sender, ScColor.IOTA, amount);
		}
		long deposit = request.Balance(ScColor.IOTA);
		if (deposit - amount != 0) {
			sc.Transfer(sc.Contract().Owner(), ScColor.IOTA, deposit - amount);
		}

		// refund all other token colors, don't keep tokens that were to be auctioned
		ScImmutableColorArray colors = request.Colors();
		int items = colors.Length();
		for (int i = 0; i < items; i++) {
			ScColor color = colors.GetColor(i).Value();
			if (!color.equals(ScColor.IOTA)) {
				sc.Transfer(sender, color, request.Balance(color));
			}
		}
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
		ScAddress auctionOwner;
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
		ScAddress address;
		// the amount is a cumulative sum of all bids from the same bidder
		long amount;
		// most recent bid update time
		long when;
	}
}
