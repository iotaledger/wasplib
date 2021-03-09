// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.tokenregistry.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.tokenregistry.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class TokenRegistryThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc("mintSupply", TokenRegistryThunk::funcMintSupplyThunk);
        exports.AddFunc("transferOwnership", TokenRegistryThunk::funcTransferOwnershipThunk);
        exports.AddFunc("updateMetadata", TokenRegistryThunk::funcUpdateMetadataThunk);
        exports.AddView("getInfo", TokenRegistryThunk::viewGetInfoThunk);
    }

    private static void funcMintSupplyThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncMintSupplyParams();
        params.Description = p.GetString(Consts.ParamDescription);
        params.UserDefined = p.GetString(Consts.ParamUserDefined);
        ctx.Log("tokenregistry.funcMintSupply");
        TokenRegistry.funcMintSupply(ctx, params);
        ctx.Log("tokenregistry.funcMintSupply ok");
    }

    private static void funcTransferOwnershipThunk(ScFuncContext ctx) {
        //TODO the one who can transfer token ownership
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncTransferOwnershipParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        ctx.Log("tokenregistry.funcTransferOwnership");
        TokenRegistry.funcTransferOwnership(ctx, params);
        ctx.Log("tokenregistry.funcTransferOwnership ok");
    }

    private static void funcUpdateMetadataThunk(ScFuncContext ctx) {
        //TODO the one who can change the token info
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncUpdateMetadataParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        ctx.Log("tokenregistry.funcUpdateMetadata");
        TokenRegistry.funcUpdateMetadata(ctx, params);
        ctx.Log("tokenregistry.funcUpdateMetadata ok");
    }

    private static void viewGetInfoThunk(ScViewContext ctx) {
        var p = ctx.Params();
        var params = new ViewGetInfoParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        ctx.Log("tokenregistry.viewGetInfo");
        TokenRegistry.viewGetInfo(ctx, params);
        ctx.Log("tokenregistry.viewGetInfo ok");
    }
}