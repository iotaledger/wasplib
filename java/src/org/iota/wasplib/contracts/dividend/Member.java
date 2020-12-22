// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.dividend;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAddress;

public class Member {
	public ScAddress address;
	public long factor;

	public static byte[] encode(Member o) {
		return new BytesEncoder().
				Address(o.address).
				Int(o.factor).
				Data();
	}

	public static Member decode(byte[] bytes) {
		BytesDecoder d = new BytesDecoder(bytes);
		Member data = new Member();
		data.address = d.Address();
		data.factor = d.Int();
		return data;
	}
}
