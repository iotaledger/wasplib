// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;

public class BidInfo {
	//@formatter:off
	public long amount;    // cumulative amount of bids from same bidder
	public long index;     // index of bidder in bidder list
	public long timestamp; // timestamp of most recent bid
	//@formatter:on

	public static byte[] encode(BidInfo o) {
		return new BytesEncoder().
				Int(o.amount).
				Int(o.index).
				Int(o.timestamp).
				Data();
	}

	public static BidInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		BidInfo data = new BidInfo();
		data.amount = decode.Int();
		data.index = decode.Int();
		data.timestamp = decode.Int();
		return data;
	}
}
