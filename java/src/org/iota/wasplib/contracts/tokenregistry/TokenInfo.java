// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.tokenregistry;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;

public class TokenInfo{
	//@formatter:off
	public long    created;
	public String  description;
	public ScAgent mintedBy;
	public ScAgent owner;
	public long    supply;
	public long    updated;
	public String  userDefined;
	//@formatter:on

	public static byte[] encode(TokenInfo o){
		return new BytesEncoder().
				Int(o.created).
				String(o.description).
				Agent(o.mintedBy).
				Agent(o.owner).
				Int(o.supply).
				Int(o.updated).
				String(o.userDefined).
				Data();
	}

	public static TokenInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
        TokenInfo data = new TokenInfo();
		data.created = decode.Int();
		data.description = decode.String();
		data.mintedBy = decode.Agent();
		data.owner = decode.Agent();
		data.supply = decode.Int();
		data.updated = decode.Int();
		data.userDefined = decode.String();
		return data;
	}
}
