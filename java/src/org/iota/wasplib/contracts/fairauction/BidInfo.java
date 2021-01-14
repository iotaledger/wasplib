// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;

public class BidInfo {
	//@formatter:off
	public long Amount;    // cumulative amount of bids from same bidder
	public long Index;     // index of bidder in bidder list
	public long Timestamp; // timestamp of most recent bid
	//@formatter:on

	public static byte[] encode(BidInfo o) {
		return new BytesEncoder().
				Int(o.Amount).
				Int(o.Index).
				Int(o.Timestamp).
				Data();
	}

	public static BidInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		BidInfo data = new BidInfo();
		data.Amount = decode.Int();
		data.Index = decode.Int();
		data.Timestamp = decode.Int();
		return data;
	}
}
