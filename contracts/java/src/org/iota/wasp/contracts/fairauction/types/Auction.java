// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairauction.types;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class Auction {
	//@formatter:off
	public ScColor   Color;         // color of tokens for sale
	public ScAgentId Creator;       // issuer of start_auction transaction
	public long      Deposit;       // deposit by auction owner to cover the SC fees
	public String    Description;   // auction description
	public long      Duration;      // auction duration in minutes
	public long      HighestBid;    // the current highest bid amount
	public ScAgentId HighestBidder; // the current highest bidder
	public long      MinimumBid;    // minimum bid amount
	public long      NumTokens;     // number of tokens for sale
	public long      OwnerMargin;   // auction owner's margin in promilles
	public long      WhenStarted;   // timestamp when auction started
	//@formatter:on

	public Auction() {
	}

	public Auction(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		Color = decode.Color();
		Creator = decode.AgentId();
		Deposit = decode.Int();
		Description = decode.String();
		Duration = decode.Int();
		HighestBid = decode.Int();
		HighestBidder = decode.AgentId();
		MinimumBid = decode.Int();
		NumTokens = decode.Int();
		OwnerMargin = decode.Int();
		WhenStarted = decode.Int();
	}

	public byte[] toBytes() {
		return new BytesEncoder().
				Color(Color).
				AgentId(Creator).
				Int(Deposit).
				String(Description).
				Int(Duration).
				Int(HighestBid).
				AgentId(HighestBidder).
				Int(MinimumBid).
				Int(NumTokens).
				Int(OwnerMargin).
				Int(WhenStarted).
				Data();
	}
}
