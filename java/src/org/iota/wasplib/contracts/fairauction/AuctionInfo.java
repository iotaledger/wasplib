// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class AuctionInfo {
	//@formatter:off
	public ScColor Color;         // color of tokens for sale
	public ScAgent Creator;       // issuer of start_auction transaction
	public long    Deposit;       // deposit by auction owner to cover the SC fees
	public String  Description;   // auction description
	public long    Duration;      // auction duration in minutes
	public long    HighestBid;    // the current highest bid amount
	public ScAgent HighestBidder; // the current highest bidder
	public long    MinimumBid;    // minimum bid amount
	public long    NumTokens;     // number of tokens for sale
	public long    OwnerMargin;   // auction owner's margin in promilles
	public long    WhenStarted;   // timestamp when auction started
	//@formatter:on

	public static byte[] encode(AuctionInfo o) {
		return new BytesEncoder().
				Color(o.Color).
				Agent(o.Creator).
				Int(o.Deposit).
				String(o.Description).
				Int(o.Duration).
				Int(o.HighestBid).
				Agent(o.HighestBidder).
				Int(o.MinimumBid).
				Int(o.NumTokens).
				Int(o.OwnerMargin).
				Int(o.WhenStarted).
				Data();
	}

	public static AuctionInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		AuctionInfo data = new AuctionInfo();
		data.Color = decode.Color();
		data.Creator = decode.Agent();
		data.Deposit = decode.Int();
		data.Description = decode.String();
		data.Duration = decode.Int();
		data.HighestBid = decode.Int();
		data.HighestBidder = decode.Agent();
		data.MinimumBid = decode.Int();
		data.NumTokens = decode.Int();
		data.OwnerMargin = decode.Int();
		data.WhenStarted = decode.Int();
		return data;
	}
}
