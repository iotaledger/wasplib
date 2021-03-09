// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairroulette.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.fairroulette.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class FairRouletteThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc("lockBets", FairRouletteThunk::funcLockBetsThunk);
        exports.AddFunc("payWinners", FairRouletteThunk::funcPayWinnersThunk);
        exports.AddFunc("placeBet", FairRouletteThunk::funcPlaceBetThunk);
        exports.AddFunc("playPeriod", FairRouletteThunk::funcPlayPeriodThunk);
    }

    private static void funcLockBetsThunk(ScFuncContext ctx) {
        // only SC itself can invoke this function
        ctx.Require(ctx.Caller().equals(ctx.ContractId().AsAgentId()), "no permission");

        var params = new FuncLockBetsParams();
        ctx.Log("fairroulette.funcLockBets");
        FairRoulette.funcLockBets(ctx, params);
        ctx.Log("fairroulette.funcLockBets ok");
    }

    private static void funcPayWinnersThunk(ScFuncContext ctx) {
        // only SC itself can invoke this function
        ctx.Require(ctx.Caller().equals(ctx.ContractId().AsAgentId()), "no permission");

        var params = new FuncPayWinnersParams();
        ctx.Log("fairroulette.funcPayWinners");
        FairRoulette.funcPayWinners(ctx, params);
        ctx.Log("fairroulette.funcPayWinners ok");
    }

    private static void funcPlaceBetThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncPlaceBetParams();
        params.Number = p.GetInt64(Consts.ParamNumber);
        ctx.Require(params.Number.Exists(), "missing mandatory number");
        ctx.Log("fairroulette.funcPlaceBet");
        FairRoulette.funcPlaceBet(ctx, params);
        ctx.Log("fairroulette.funcPlaceBet ok");
    }

    private static void funcPlayPeriodThunk(ScFuncContext ctx) {
        // only SC creator can update the play period
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncPlayPeriodParams();
        params.PlayPeriod = p.GetInt64(Consts.ParamPlayPeriod);
        ctx.Require(params.PlayPeriod.Exists(), "missing mandatory playPeriod");
        ctx.Log("fairroulette.funcPlayPeriod");
        FairRoulette.funcPlayPeriod(ctx, params);
        ctx.Log("fairroulette.funcPlayPeriod ok");
    }
}