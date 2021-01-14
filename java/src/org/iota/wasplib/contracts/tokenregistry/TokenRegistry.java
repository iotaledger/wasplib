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
	private static final Key KeyColorList = new Key("color_list");
	private static final Key KeyDescription = new Key("description");
	private static final Key KeyRegistry = new Key("registry");
	private static final Key KeyUserDefined = new Key("user_defined");

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
		ScMutableBytes registry = state.GetMap(KeyRegistry).GetBytes(minted);
		if (registry.Exists()) {
			sc.Panic("TokenRegistry: Color already exists");
		}
		ScImmutableMap params = sc.Params();
		TokenInfo token = new TokenInfo();
		{
			token.Supply = sc.Incoming().Balance(minted);
			token.MintedBy = sc.Caller();
			token.Owner = sc.Caller();
			token.Created = sc.Timestamp();
			token.Updated = sc.Timestamp();
			token.Description = params.GetString(KeyDescription).Value();
			token.UserDefined = params.GetString(KeyUserDefined).Value();
		}
		if (token.Supply <= 0) {
			sc.Panic("TokenRegistry: Insufficient supply");
		}
		if (token.Description.isEmpty()) {
			token.Description += "no dscr";
		}
		registry.SetValue(TokenInfo.encode(token));
		ScMutableColorArray colors = state.GetColorArray(KeyColorList);
		colors.GetColor(colors.Length()).SetValue(minted);
	}

	public static void updateMetadata(ScCallContext _sc) {
		//TODO
	}

	public static void transferOwnership(ScCallContext _sc) {
		//TODO
	}
}
