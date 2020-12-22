// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairroulette;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class BetInfo {
	public long amount;
	public ScAgent better;
	public long color;

	public static byte[] encode(BetInfo o) {
		return new BytesEncoder().
				Int(o.amount).
				Agent(o.better).
				Int(o.color).
				Data();
	}

	public static BetInfo decode(byte[] bytes) {
		BytesDecoder d = new BytesDecoder(bytes);
		BetInfo data = new BetInfo();
		data.amount = d.Int();
		data.better = d.Agent();
		data.color = d.Int();
		return data;
	}
}
