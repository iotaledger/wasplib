// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.tokenregistry;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class TokenInfo {
	public long created;
	public String description;
	public ScAgent mintedBy;
	public ScAgent owner;
	public long supply;
	public long updated;
	public String userDefined;

	public static byte[] encode(TokenInfo o) {
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
		BytesDecoder d = new BytesDecoder(bytes);
		TokenInfo data = new TokenInfo();
		data.created = d.Int();
		data.description = d.String();
		data.mintedBy = d.Agent();
		data.owner = d.Agent();
		data.supply = d.Int();
		data.updated = d.Int();
		data.userDefined = d.String();
		return data;
	}
}
