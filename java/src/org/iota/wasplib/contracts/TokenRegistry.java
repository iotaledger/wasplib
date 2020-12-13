// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableColorArray;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class TokenRegistry {
	private static final Key keyColorList = new Key("color_list");
	private static final Key keyDescription = new Key("description");
	private static final Key keyRegistry = new Key("registry");
	private static final Key keyUserDefined = new Key("user_defined");

	//export on_load
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("mint_supply", TokenRegistry::mintSupply);
		exports.AddCall("update_metadata", TokenRegistry::updateMetadata);
		exports.AddCall("transfer_ownership", TokenRegistry::transferOwnership);
	}

	public static void mintSupply(ScCallContext sc) {
		ScColor minted = sc.Incoming().Minted();
		if (minted.equals(ScColor.MINT)) {
			sc.Log("TokenRegistry: No newly minted tokens found");
			return;
		}
		ScMutableMap state = sc.State();
		ScMutableBytes registry = state.GetMap(keyRegistry).GetBytes(minted);
		if (registry.Exists()) {
			sc.Log("TokenRegistry: Color already exists");
			return;
		}
		ScImmutableMap params = sc.Params();
		TokenInfo token = new TokenInfo();
		token.supply = sc.Balances().Balance(minted);
		token.mintedBy = sc.Caller();
		token.owner = sc.Caller();
		token.created = sc.Timestamp();
		token.updated = sc.Timestamp();
		token.description = params.GetString(keyDescription).Value();
		token.userDefined = params.GetString(keyUserDefined).Value();
		if (token.supply <= 0) {
			sc.Log("TokenRegistry: Insufficient supply");
			return;
		}
		if (token.description.isEmpty()) {
			token.description += "no dscr";
		}
		byte[] bytes = encodeTokenInfo(token);
		registry.SetValue(bytes);
		ScMutableColorArray colors = state.GetColorArray(keyColorList);
		colors.GetColor(colors.Length()).SetValue(minted);
	}

	public static void updateMetadata(ScCallContext sc) {
		//TODO
	}

	public static void transferOwnership(ScCallContext sc) {
		//TODO
	}

	public static TokenInfo decodeTokenInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		TokenInfo token = new TokenInfo();
		token.supply = decoder.Int();
		token.mintedBy = decoder.Agent();
		token.owner = decoder.Agent();
		token.created = decoder.Int();
		token.updated = decoder.Int();
		token.description = decoder.String();
		token.userDefined = decoder.String();
		return token;
	}

	public static byte[] encodeTokenInfo(TokenInfo token) {
		return new BytesEncoder().
				Int(token.supply).
				Agent(token.mintedBy).
				Agent(token.owner).
				Int(token.created).
				Int(token.updated).
				String(token.description).
				String(token.userDefined).
				Data();
	}

	public static class TokenInfo {
		long supply;
		ScAgent mintedBy;
		ScAgent owner;
		long created;
		long updated;
		String description;
		String userDefined;
	}
}
