// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairauction;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;

public class BidInfo {
	public long amount; // cumulative amount of bids from same bidder
	public long index; // index of bidder in bidder list
	public long timestamp; // timestamp of most recent bid

	public static byte[] encode(BidInfo o) {
		return new BytesEncoder().
				Int(o.amount).
				Int(o.index).
				Int(o.timestamp).
				Data();
	}

	public static BidInfo decode(byte[] bytes) {
		BytesDecoder d = new BytesDecoder(bytes);
		BidInfo data = new BidInfo();
		data.amount = d.Int();
		data.index = d.Int();
		data.timestamp = d.Int();
		return data;
	}
}
