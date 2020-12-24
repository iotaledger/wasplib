// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairroulette;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class BetInfo{
	//@formatter:off
	public long    amount;
	public ScAgent better;
	public long    color;
	//@formatter:on

	public static byte[] encode(BetInfo o){
		return new BytesEncoder().
				Int(o.amount).
				Agent(o.better).
				Int(o.color).
				Data();
	}

	public static BetInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
        BetInfo data = new BetInfo();
		data.amount = decode.Int();
		data.better = decode.Agent();
		data.color = decode.Int();
		return data;
	}
}
