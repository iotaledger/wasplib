// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.tokenregistry;

import org.iota.wasp.contracts.tokenregistry.lib.*;
import org.iota.wasp.contracts.tokenregistry.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class TokenRegistry {

    public static void funcMintSupply(ScFuncContext ctx, FuncMintSupplyParams params) {
        var mintedSupply = ctx.MintedSupply();
        if (mintedSupply == 0) {
            ctx.Panic("TokenRegistry: No newly minted tokens found");
        }
        var mintedColor = ctx.MintedColor();
        var state = ctx.State();
        var registry = state.GetMap(Consts.VarRegistry).GetBytes(mintedColor);
        if (registry.Exists()) {
            // should never happen, because transaction id is unique
            ctx.Panic("TokenRegistry: registry for color already exists");
        }
        var token = new Token();
        {
            token.Supply = mintedSupply;
            token.MintedBy = ctx.Caller();
            token.Owner = ctx.Caller();
            token.Created = ctx.Timestamp();
            token.Updated = ctx.Timestamp();
            token.Description = params.Description.Value();
            token.UserDefined = params.UserDefined.Value();
        }
        if (token.Description.isEmpty()) {
            token.Description += "no dscr";
        }
        registry.SetValue(token.toBytes());
        var colors = state.GetColorArray(Consts.VarColorList);
        colors.GetColor(colors.Length()).SetValue(mintedColor);
        ctx.Log("tokenregistry.mintSupply ok");
    }

    public static void funcTransferOwnership(ScFuncContext ctx, FuncTransferOwnershipParams params) {
        //TODO
    }

    public static void funcUpdateMetadata(ScFuncContext ctx, FuncUpdateMetadataParams params) {
        //TODO
    }

    public static void viewGetInfo(ScViewContext ctx, ViewGetInfoParams params) {
        //TODO
    }
}
