// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.dividend;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAddress;

public class Member {
	//@formatter:off
	public ScAddress Address;
	public long      Factor;
	//@formatter:on

	public static byte[] encode(Member o) {
		return new BytesEncoder().
				Address(o.Address).
				Int(o.Factor).
				Data();
	}

	public static Member decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		Member data = new Member();
		data.Address = decode.Address();
		data.Factor = decode.Int();
		return data;
	}
}
