// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.dividend;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class Member{
	//@formatter:off
	public ScAddress address;
	public long      factor;
	//@formatter:on

	public static byte[] encode(Member o){
		return new BytesEncoder().
				Address(o.address).
				Int(o.factor).
				Data();
	}

	public static Member decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
        Member data = new Member();
		data.address = decode.Address();
		data.factor = decode.Int();
		return data;
	}
}
