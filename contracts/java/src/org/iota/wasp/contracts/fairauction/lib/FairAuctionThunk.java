// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairauction.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.fairauction.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class FairAuctionThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc(Consts.FuncFinalizeAuction, FairAuctionThunk::funcFinalizeAuctionThunk);
        exports.AddFunc(Consts.FuncPlaceBid, FairAuctionThunk::funcPlaceBidThunk);
        exports.AddFunc(Consts.FuncSetOwnerMargin, FairAuctionThunk::funcSetOwnerMarginThunk);
        exports.AddFunc(Consts.FuncStartAuction, FairAuctionThunk::funcStartAuctionThunk);
        exports.AddView(Consts.ViewGetInfo, FairAuctionThunk::viewGetInfoThunk);
    }

    private static void funcFinalizeAuctionThunk(ScFuncContext ctx) {
        ctx.Log("fairauction.funcFinalizeAuction");
        // only SC itself can invoke this function
        ctx.Require(ctx.Caller().equals(ctx.ContractId().AsAgentId()), "no permission");

        var p = ctx.Params();
        var params = new FuncFinalizeAuctionParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        FairAuction.funcFinalizeAuction(ctx, params);
        ctx.Log("fairauction.funcFinalizeAuction ok");
    }

    private static void funcPlaceBidThunk(ScFuncContext ctx) {
        ctx.Log("fairauction.funcPlaceBid");
        var p = ctx.Params();
        var params = new FuncPlaceBidParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        FairAuction.funcPlaceBid(ctx, params);
        ctx.Log("fairauction.funcPlaceBid ok");
    }

    private static void funcSetOwnerMarginThunk(ScFuncContext ctx) {
        ctx.Log("fairauction.funcSetOwnerMargin");
        // only SC creator can set owner margin
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncSetOwnerMarginParams();
        params.OwnerMargin = p.GetInt64(Consts.ParamOwnerMargin);
        ctx.Require(params.OwnerMargin.Exists(), "missing mandatory ownerMargin");
        FairAuction.funcSetOwnerMargin(ctx, params);
        ctx.Log("fairauction.funcSetOwnerMargin ok");
    }

    private static void funcStartAuctionThunk(ScFuncContext ctx) {
        ctx.Log("fairauction.funcStartAuction");
        var p = ctx.Params();
        var params = new FuncStartAuctionParams();
        params.Color = p.GetColor(Consts.ParamColor);
        params.Description = p.GetString(Consts.ParamDescription);
        params.Duration = p.GetInt64(Consts.ParamDuration);
        params.MinimumBid = p.GetInt64(Consts.ParamMinimumBid);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        ctx.Require(params.MinimumBid.Exists(), "missing mandatory minimumBid");
        FairAuction.funcStartAuction(ctx, params);
        ctx.Log("fairauction.funcStartAuction ok");
    }

    private static void viewGetInfoThunk(ScViewContext ctx) {
        ctx.Log("fairauction.viewGetInfo");
        var p = ctx.Params();
        var params = new ViewGetInfoParams();
        params.Color = p.GetColor(Consts.ParamColor);
        ctx.Require(params.Color.Exists(), "missing mandatory color");
        FairAuction.viewGetInfo(ctx, params);
        ctx.Log("fairauction.viewGetInfo ok");
    }
}
