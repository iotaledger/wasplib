// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class AuctionInfo {
	//@formatter:off
	public ScAgent auctionOwner;  // issuer of start_auction transaction
	public ScColor color;         // color of tokens for sale
	public long    deposit;       // deposit by auction owner to cover the SC fees
	public String  description;   // auction description
	public long    duration;      // auction duration in minutes
	public long    highestBid;    // the current highest bid amount
	public ScAgent highestBidder; // the current highest bidder
	public long    minimumBid;    // minimum bid amount
	public long    numTokens;     // number of tokens for sale
	public long    ownerMargin;   // auction owner's margin in promilles
	public long    whenStarted;   // timestamp when auction started
	//@formatter:on

	public static byte[] encode(AuctionInfo o) {
		return new BytesEncoder().
				Agent(o.auctionOwner).
				Color(o.color).
				Int(o.deposit).
				String(o.description).
				Int(o.duration).
				Int(o.highestBid).
				Agent(o.highestBidder).
				Int(o.minimumBid).
				Int(o.numTokens).
				Int(o.ownerMargin).
				Int(o.whenStarted).
				Data();
	}

	public static AuctionInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		AuctionInfo data = new AuctionInfo();
		data.auctionOwner = decode.Agent();
		data.color = decode.Color();
		data.deposit = decode.Int();
		data.description = decode.String();
		data.duration = decode.Int();
		data.highestBid = decode.Int();
		data.highestBidder = decode.Agent();
		data.minimumBid = decode.Int();
		data.numTokens = decode.Int();
		data.ownerMargin = decode.Int();
		data.whenStarted = decode.Int();
		return data;
	}
}
