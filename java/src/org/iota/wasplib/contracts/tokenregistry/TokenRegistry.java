// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.tokenregistry;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
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

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("mint_supply", TokenRegistry::mintSupply);
		exports.AddCall("update_metadata", TokenRegistry::updateMetadata);
		exports.AddCall("transfer_ownership", TokenRegistry::transferOwnership);
	}

	public static void mintSupply(ScCallContext sc) {
		ScColor minted = sc.Incoming().Minted();
		if (minted.equals(ScColor.MINT)) {
			sc.Panic("TokenRegistry: No newly minted tokens found");
		}
		ScMutableMap state = sc.State();
		ScMutableBytes registry = state.GetMap(keyRegistry).GetBytes(minted);
		if (registry.Exists()) {
			sc.Panic("TokenRegistry: Color already exists");
		}
		ScImmutableMap params = sc.Params();
		TokenInfo token = new TokenInfo();
		{
			token.supply = sc.Incoming().Balance(minted);
			token.mintedBy = sc.Caller();
			token.owner = sc.Caller();
			token.created = sc.Timestamp();
			token.updated = sc.Timestamp();
			token.description = params.GetString(keyDescription).Value();
			token.userDefined = params.GetString(keyUserDefined).Value();
		}
		if (token.supply <= 0) {
			sc.Panic("TokenRegistry: Insufficient supply");
		}
		if (token.description.isEmpty()) {
			token.description += "no dscr";
		}
		registry.SetValue(TokenInfo.encode(token));
		ScMutableColorArray colors = state.GetColorArray(keyColorList);
		colors.GetColor(colors.Length()).SetValue(minted);
	}

	public static void updateMetadata(ScCallContext _sc) {
		//TODO
	}

	public static void transferOwnership(ScCallContext _sc) {
		//TODO
	}
}
