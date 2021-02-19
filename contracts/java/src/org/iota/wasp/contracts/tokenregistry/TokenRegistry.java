// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.tokenregistry;

import org.iota.wasp.contracts.tokenregistry.lib.FuncMintSupplyParams;
import org.iota.wasp.contracts.tokenregistry.lib.FuncTransferOwnershipParams;
import org.iota.wasp.contracts.tokenregistry.lib.FuncUpdateMetadataParams;
import org.iota.wasp.contracts.tokenregistry.lib.ViewGetInfoParams;
import org.iota.wasp.contracts.tokenregistry.types.Token;
import org.iota.wasp.wasmlib.context.ScFuncContext;
import org.iota.wasp.wasmlib.context.ScViewContext;
import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.immutable.ScImmutableMap;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.mutable.ScMutableBytes;
import org.iota.wasp.wasmlib.mutable.ScMutableColorArray;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

public class TokenRegistry {
	private static final Key KeyColorList = new Key("color_list");
	private static final Key KeyDescription = new Key("description");
	private static final Key KeyRegistry = new Key("registry");
	private static final Key KeyUserDefined = new Key("user_defined");

	public static void FuncMintSupply(ScFuncContext ctx, FuncMintSupplyParams params) {
		ScColor minted = ctx.Incoming().Minted();
		if (minted.equals(ScColor.MINT)) {
			ctx.Panic("TokenRegistry: No newly minted tokens found");
		}
		ScMutableMap state = ctx.State();
		ScMutableBytes registry = state.GetMap(KeyRegistry).GetBytes(minted);
		if (registry.Exists()) {
			ctx.Panic("TokenRegistry: Color already exists");
		}
		ScImmutableMap p = ctx.Params();
		Token token = new Token();
		{
			token.Supply = ctx.Incoming().Balance(minted);
			token.MintedBy = ctx.Caller();
			token.Owner = ctx.Caller();
			token.Created = ctx.Timestamp();
			token.Updated = ctx.Timestamp();
			token.Description = p.GetString(KeyDescription).Value();
			token.UserDefined = p.GetString(KeyUserDefined).Value();
		}
		if (token.Supply <= 0) {
			ctx.Panic("TokenRegistry: Insufficient supply");
		}
		if (token.Description.isEmpty()) {
			token.Description += "no dscr";
		}
		registry.SetValue(token.toBytes());
		ScMutableColorArray colors = state.GetColorArray(KeyColorList);
		colors.GetColor(colors.Length()).SetValue(minted);
	}

	public static void FuncTransferOwnership(ScFuncContext ctx, FuncTransferOwnershipParams params) {
		//TODO
	}

	public static void FuncUpdateMetadata(ScFuncContext ctx, FuncUpdateMetadataParams params) {
		//TODO
	}

	public static void ViewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
		//TODO
	}
}
