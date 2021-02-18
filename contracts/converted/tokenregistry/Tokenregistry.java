// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.tokenregistry;

public class Tokenregistry {

public static void funcMintSupply(ScFuncContext ctx, FuncMintSupplyParams params) {
    minted = ctx.Incoming().Minted();
    if (minted == ScColor.MINT) {
        ctx.Panic("TokenRegistry: No newly minted tokens found");
    }
    state = ctx.State();
    registry = state.GetMap(VarRegistry).GetBytes(minted);
    if (registry.Exists()) {
        ctx.Panic("TokenRegistry: Color already exists");
    }
    Token token = new Token();
         {
        token.Supply = ctx.Incoming().Balance(minted);
        token.MintedBy = ctx.Caller();
        token.Owner = ctx.Caller();
        token.Created = ctx.Timestamp();
        token.Updated = ctx.Timestamp();
        token.Description = params.Description.Value();
        token.UserDefined = params.UserDefined.Value();
    }
    if (token.Supply <= 0) {
        ctx.Panic("TokenRegistry: Insufficient supply");
    }
    if (token.Description.IsEmpty()) {
        token.Description += "no dscr";
    }
    registry.SetValue(token.ToBytes());
    colors = state.GetColorArray(VarColorList);
    colors.GetColor(colors.Length()).SetValue(minted);
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
