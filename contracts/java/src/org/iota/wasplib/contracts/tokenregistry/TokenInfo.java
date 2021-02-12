// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.tokenregistry;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class TokenInfo {
	//@formatter:off
	public long    Created;
	public String  Description;
	public ScAgent MintedBy;
	public ScAgent Owner;
	public long    Supply;
	public long    Updated;
	public String  UserDefined;
	//@formatter:on

	public static byte[] encode(TokenInfo o) {
		return new BytesEncoder().
				Int(o.Created).
				String(o.Description).
				Agent(o.MintedBy).
				Agent(o.Owner).
				Int(o.Supply).
				Int(o.Updated).
				String(o.UserDefined).
				Data();
	}

	public static TokenInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		TokenInfo data = new TokenInfo();
		data.Created = decode.Int();
		data.Description = decode.String();
		data.MintedBy = decode.Agent();
		data.Owner = decode.Agent();
		data.Supply = decode.Int();
		data.Updated = decode.Int();
		data.UserDefined = decode.String();
		return data;
	}
}
