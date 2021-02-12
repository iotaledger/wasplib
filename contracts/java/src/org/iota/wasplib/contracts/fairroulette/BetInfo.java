// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairroulette;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class BetInfo {
	//@formatter:off
	public long    Amount;
	public ScAgent Better;
	public long    Color;
	//@formatter:on

	public static byte[] encode(BetInfo o) {
		return new BytesEncoder().
				Int(o.Amount).
				Agent(o.Better).
				Int(o.Color).
				Data();
	}

	public static BetInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		BetInfo data = new BetInfo();
		data.Amount = decode.Int();
		data.Better = decode.Agent();
		data.Color = decode.Int();
		return data;
	}
}
