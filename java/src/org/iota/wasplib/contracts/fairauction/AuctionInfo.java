// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class AuctionInfo {
	public ScAgent auctionOwner; // issuer of start_auction transaction
	public ScColor color; // color of tokens for sale
	public long deposit; // deposit by auction owner to cover the SC fees
	public String description; // auction description
	public long duration; // auction duration in minutes
	public long highestBid; // the current highest bid amount
	public ScAgent highestBidder; // the current highest bidder
	public long minimumBid; // minimum bid amount
	public long numTokens; // number of tokens for sale
	public long ownerMargin; // auction owner's margin in promilles
	public long whenStarted; // timestamp when auction started

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
		BytesDecoder d = new BytesDecoder(bytes);
		AuctionInfo data = new AuctionInfo();
		data.auctionOwner = d.Agent();
		data.color = d.Color();
		data.deposit = d.Int();
		data.description = d.String();
		data.duration = d.Int();
		data.highestBid = d.Int();
		data.highestBidder = d.Agent();
		data.minimumBid = d.Int();
		data.numTokens = d.Int();
		data.ownerMargin = d.Int();
		data.whenStarted = d.Int();
		return data;
	}
}
