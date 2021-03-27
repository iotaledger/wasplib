// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.testcore.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.testcore.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class TestCoreThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc(Consts.FuncCallOnChain, TestCoreThunk::funcCallOnChainThunk);
        exports.AddFunc(Consts.FuncCheckContextFromFullEP, TestCoreThunk::funcCheckContextFromFullEPThunk);
        exports.AddFunc(Consts.FuncDoNothing, TestCoreThunk::funcDoNothingThunk);
        exports.AddFunc(Consts.FuncGetMintedSupply, TestCoreThunk::funcGetMintedSupplyThunk);
        exports.AddFunc(Consts.FuncIncCounter, TestCoreThunk::funcIncCounterThunk);
        exports.AddFunc(Consts.FuncInit, TestCoreThunk::funcInitThunk);
        exports.AddFunc(Consts.FuncPassTypesFull, TestCoreThunk::funcPassTypesFullThunk);
        exports.AddFunc(Consts.FuncRunRecursion, TestCoreThunk::funcRunRecursionThunk);
        exports.AddFunc(Consts.FuncSendToAddress, TestCoreThunk::funcSendToAddressThunk);
        exports.AddFunc(Consts.FuncSetInt, TestCoreThunk::funcSetIntThunk);
        exports.AddFunc(Consts.FuncTestCallPanicFullEP, TestCoreThunk::funcTestCallPanicFullEPThunk);
        exports.AddFunc(Consts.FuncTestCallPanicViewEPFromFull, TestCoreThunk::funcTestCallPanicViewEPFromFullThunk);
        exports.AddFunc(Consts.FuncTestChainOwnerIDFull, TestCoreThunk::funcTestChainOwnerIDFullThunk);
        exports.AddFunc(Consts.FuncTestEventLogDeploy, TestCoreThunk::funcTestEventLogDeployThunk);
        exports.AddFunc(Consts.FuncTestEventLogEventData, TestCoreThunk::funcTestEventLogEventDataThunk);
        exports.AddFunc(Consts.FuncTestEventLogGenericData, TestCoreThunk::funcTestEventLogGenericDataThunk);
        exports.AddFunc(Consts.FuncTestPanicFullEP, TestCoreThunk::funcTestPanicFullEPThunk);
        exports.AddFunc(Consts.FuncWithdrawToChain, TestCoreThunk::funcWithdrawToChainThunk);
        exports.AddView(Consts.ViewCheckContextFromViewEP, TestCoreThunk::viewCheckContextFromViewEPThunk);
        exports.AddView(Consts.ViewFibonacci, TestCoreThunk::viewFibonacciThunk);
        exports.AddView(Consts.ViewGetCounter, TestCoreThunk::viewGetCounterThunk);
        exports.AddView(Consts.ViewGetInt, TestCoreThunk::viewGetIntThunk);
        exports.AddView(Consts.ViewJustView, TestCoreThunk::viewJustViewThunk);
        exports.AddView(Consts.ViewPassTypesView, TestCoreThunk::viewPassTypesViewThunk);
        exports.AddView(Consts.ViewTestCallPanicViewEPFromView, TestCoreThunk::viewTestCallPanicViewEPFromViewThunk);
        exports.AddView(Consts.ViewTestChainOwnerIDView, TestCoreThunk::viewTestChainOwnerIDViewThunk);
        exports.AddView(Consts.ViewTestPanicViewEP, TestCoreThunk::viewTestPanicViewEPThunk);
        exports.AddView(Consts.ViewTestSandboxCall, TestCoreThunk::viewTestSandboxCallThunk);
    }

    private static void funcCallOnChainThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcCallOnChain");
        var p = ctx.Params();
        var params = new FuncCallOnChainParams();
        params.HnameContract = p.GetHname(Consts.ParamHnameContract);
        params.HnameEP = p.GetHname(Consts.ParamHnameEP);
        params.IntValue = p.GetInt64(Consts.ParamIntValue);
        ctx.Require(params.IntValue.Exists(), "missing mandatory intValue");
        TestCore.funcCallOnChain(ctx, params);
        ctx.Log("testcore.funcCallOnChain ok");
    }

    private static void funcCheckContextFromFullEPThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcCheckContextFromFullEP");
        var p = ctx.Params();
        var params = new FuncCheckContextFromFullEPParams();
        params.AgentId = p.GetAgentId(Consts.ParamAgentId);
        params.Caller = p.GetAgentId(Consts.ParamCaller);
        params.ChainId = p.GetChainId(Consts.ParamChainId);
        params.ChainOwnerId = p.GetAgentId(Consts.ParamChainOwnerId);
        params.ContractCreator = p.GetAgentId(Consts.ParamContractCreator);
        ctx.Require(params.AgentId.Exists(), "missing mandatory agentId");
        ctx.Require(params.Caller.Exists(), "missing mandatory caller");
        ctx.Require(params.ChainId.Exists(), "missing mandatory chainId");
        ctx.Require(params.ChainOwnerId.Exists(), "missing mandatory chainOwnerId");
        ctx.Require(params.ContractCreator.Exists(), "missing mandatory contractCreator");
        TestCore.funcCheckContextFromFullEP(ctx, params);
        ctx.Log("testcore.funcCheckContextFromFullEP ok");
    }

    private static void funcDoNothingThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcDoNothing");
        var params = new FuncDoNothingParams();
        TestCore.funcDoNothing(ctx, params);
        ctx.Log("testcore.funcDoNothing ok");
    }

    private static void funcGetMintedSupplyThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcGetMintedSupply");
        var params = new FuncGetMintedSupplyParams();
        TestCore.funcGetMintedSupply(ctx, params);
        ctx.Log("testcore.funcGetMintedSupply ok");
    }

    private static void funcIncCounterThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcIncCounter");
        var params = new FuncIncCounterParams();
        TestCore.funcIncCounter(ctx, params);
        ctx.Log("testcore.funcIncCounter ok");
    }

    private static void funcInitThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcInit");
        var params = new FuncInitParams();
        TestCore.funcInit(ctx, params);
        ctx.Log("testcore.funcInit ok");
    }

    private static void funcPassTypesFullThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcPassTypesFull");
        var p = ctx.Params();
        var params = new FuncPassTypesFullParams();
        params.Hash = p.GetHash(Consts.ParamHash);
        params.Hname = p.GetHname(Consts.ParamHname);
        params.HnameZero = p.GetHname(Consts.ParamHnameZero);
        params.Int64 = p.GetInt64(Consts.ParamInt64);
        params.Int64Zero = p.GetInt64(Consts.ParamInt64Zero);
        params.String = p.GetString(Consts.ParamString);
        params.StringZero = p.GetString(Consts.ParamStringZero);
        ctx.Require(params.Hash.Exists(), "missing mandatory hash");
        ctx.Require(params.Hname.Exists(), "missing mandatory hname");
        ctx.Require(params.HnameZero.Exists(), "missing mandatory hnameZero");
        ctx.Require(params.Int64.Exists(), "missing mandatory int64");
        ctx.Require(params.Int64Zero.Exists(), "missing mandatory int64Zero");
        ctx.Require(params.String.Exists(), "missing mandatory string");
        ctx.Require(params.StringZero.Exists(), "missing mandatory stringZero");
        TestCore.funcPassTypesFull(ctx, params);
        ctx.Log("testcore.funcPassTypesFull ok");
    }

    private static void funcRunRecursionThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcRunRecursion");
        var p = ctx.Params();
        var params = new FuncRunRecursionParams();
        params.IntValue = p.GetInt64(Consts.ParamIntValue);
        ctx.Require(params.IntValue.Exists(), "missing mandatory intValue");
        TestCore.funcRunRecursion(ctx, params);
        ctx.Log("testcore.funcRunRecursion ok");
    }

    private static void funcSendToAddressThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcSendToAddress");
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncSendToAddressParams();
        params.Address = p.GetAddress(Consts.ParamAddress);
        ctx.Require(params.Address.Exists(), "missing mandatory address");
        TestCore.funcSendToAddress(ctx, params);
        ctx.Log("testcore.funcSendToAddress ok");
    }

    private static void funcSetIntThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcSetInt");
        var p = ctx.Params();
        var params = new FuncSetIntParams();
        params.IntValue = p.GetInt64(Consts.ParamIntValue);
        params.Name = p.GetString(Consts.ParamName);
        ctx.Require(params.IntValue.Exists(), "missing mandatory intValue");
        ctx.Require(params.Name.Exists(), "missing mandatory name");
        TestCore.funcSetInt(ctx, params);
        ctx.Log("testcore.funcSetInt ok");
    }

    private static void funcTestCallPanicFullEPThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestCallPanicFullEP");
        var params = new FuncTestCallPanicFullEPParams();
        TestCore.funcTestCallPanicFullEP(ctx, params);
        ctx.Log("testcore.funcTestCallPanicFullEP ok");
    }

    private static void funcTestCallPanicViewEPFromFullThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestCallPanicViewEPFromFull");
        var params = new FuncTestCallPanicViewEPFromFullParams();
        TestCore.funcTestCallPanicViewEPFromFull(ctx, params);
        ctx.Log("testcore.funcTestCallPanicViewEPFromFull ok");
    }

    private static void funcTestChainOwnerIDFullThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestChainOwnerIDFull");
        var params = new FuncTestChainOwnerIDFullParams();
        TestCore.funcTestChainOwnerIDFull(ctx, params);
        ctx.Log("testcore.funcTestChainOwnerIDFull ok");
    }

    private static void funcTestEventLogDeployThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestEventLogDeploy");
        var params = new FuncTestEventLogDeployParams();
        TestCore.funcTestEventLogDeploy(ctx, params);
        ctx.Log("testcore.funcTestEventLogDeploy ok");
    }

    private static void funcTestEventLogEventDataThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestEventLogEventData");
        var params = new FuncTestEventLogEventDataParams();
        TestCore.funcTestEventLogEventData(ctx, params);
        ctx.Log("testcore.funcTestEventLogEventData ok");
    }

    private static void funcTestEventLogGenericDataThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestEventLogGenericData");
        var p = ctx.Params();
        var params = new FuncTestEventLogGenericDataParams();
        params.Counter = p.GetInt64(Consts.ParamCounter);
        ctx.Require(params.Counter.Exists(), "missing mandatory counter");
        TestCore.funcTestEventLogGenericData(ctx, params);
        ctx.Log("testcore.funcTestEventLogGenericData ok");
    }

    private static void funcTestPanicFullEPThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcTestPanicFullEP");
        var params = new FuncTestPanicFullEPParams();
        TestCore.funcTestPanicFullEP(ctx, params);
        ctx.Log("testcore.funcTestPanicFullEP ok");
    }

    private static void funcWithdrawToChainThunk(ScFuncContext ctx) {
        ctx.Log("testcore.funcWithdrawToChain");
        var p = ctx.Params();
        var params = new FuncWithdrawToChainParams();
        params.ChainId = p.GetChainId(Consts.ParamChainId);
        ctx.Require(params.ChainId.Exists(), "missing mandatory chainId");
        TestCore.funcWithdrawToChain(ctx, params);
        ctx.Log("testcore.funcWithdrawToChain ok");
    }

    private static void viewCheckContextFromViewEPThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewCheckContextFromViewEP");
        var p = ctx.Params();
        var params = new ViewCheckContextFromViewEPParams();
        params.AgentId = p.GetAgentId(Consts.ParamAgentId);
        params.ChainId = p.GetChainId(Consts.ParamChainId);
        params.ChainOwnerId = p.GetAgentId(Consts.ParamChainOwnerId);
        params.ContractCreator = p.GetAgentId(Consts.ParamContractCreator);
        ctx.Require(params.AgentId.Exists(), "missing mandatory agentId");
        ctx.Require(params.ChainId.Exists(), "missing mandatory chainId");
        ctx.Require(params.ChainOwnerId.Exists(), "missing mandatory chainOwnerId");
        ctx.Require(params.ContractCreator.Exists(), "missing mandatory contractCreator");
        TestCore.viewCheckContextFromViewEP(ctx, params);
        ctx.Log("testcore.viewCheckContextFromViewEP ok");
    }

    private static void viewFibonacciThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewFibonacci");
        var p = ctx.Params();
        var params = new ViewFibonacciParams();
        params.IntValue = p.GetInt64(Consts.ParamIntValue);
        ctx.Require(params.IntValue.Exists(), "missing mandatory intValue");
        TestCore.viewFibonacci(ctx, params);
        ctx.Log("testcore.viewFibonacci ok");
    }

    private static void viewGetCounterThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewGetCounter");
        var params = new ViewGetCounterParams();
        TestCore.viewGetCounter(ctx, params);
        ctx.Log("testcore.viewGetCounter ok");
    }

    private static void viewGetIntThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewGetInt");
        var p = ctx.Params();
        var params = new ViewGetIntParams();
        params.Name = p.GetString(Consts.ParamName);
        ctx.Require(params.Name.Exists(), "missing mandatory name");
        TestCore.viewGetInt(ctx, params);
        ctx.Log("testcore.viewGetInt ok");
    }

    private static void viewJustViewThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewJustView");
        var params = new ViewJustViewParams();
        TestCore.viewJustView(ctx, params);
        ctx.Log("testcore.viewJustView ok");
    }

    private static void viewPassTypesViewThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewPassTypesView");
        var p = ctx.Params();
        var params = new ViewPassTypesViewParams();
        params.Hash = p.GetHash(Consts.ParamHash);
        params.Hname = p.GetHname(Consts.ParamHname);
        params.HnameZero = p.GetHname(Consts.ParamHnameZero);
        params.Int64 = p.GetInt64(Consts.ParamInt64);
        params.Int64Zero = p.GetInt64(Consts.ParamInt64Zero);
        params.String = p.GetString(Consts.ParamString);
        params.StringZero = p.GetString(Consts.ParamStringZero);
        ctx.Require(params.Hash.Exists(), "missing mandatory hash");
        ctx.Require(params.Hname.Exists(), "missing mandatory hname");
        ctx.Require(params.HnameZero.Exists(), "missing mandatory hnameZero");
        ctx.Require(params.Int64.Exists(), "missing mandatory int64");
        ctx.Require(params.Int64Zero.Exists(), "missing mandatory int64Zero");
        ctx.Require(params.String.Exists(), "missing mandatory string");
        ctx.Require(params.StringZero.Exists(), "missing mandatory stringZero");
        TestCore.viewPassTypesView(ctx, params);
        ctx.Log("testcore.viewPassTypesView ok");
    }

    private static void viewTestCallPanicViewEPFromViewThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewTestCallPanicViewEPFromView");
        var params = new ViewTestCallPanicViewEPFromViewParams();
        TestCore.viewTestCallPanicViewEPFromView(ctx, params);
        ctx.Log("testcore.viewTestCallPanicViewEPFromView ok");
    }

    private static void viewTestChainOwnerIDViewThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewTestChainOwnerIDView");
        var params = new ViewTestChainOwnerIDViewParams();
        TestCore.viewTestChainOwnerIDView(ctx, params);
        ctx.Log("testcore.viewTestChainOwnerIDView ok");
    }

    private static void viewTestPanicViewEPThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewTestPanicViewEP");
        var params = new ViewTestPanicViewEPParams();
        TestCore.viewTestPanicViewEP(ctx, params);
        ctx.Log("testcore.viewTestPanicViewEP ok");
    }

    private static void viewTestSandboxCallThunk(ScViewContext ctx) {
        ctx.Log("testcore.viewTestSandboxCall");
        var params = new ViewTestSandboxCallParams();
        TestCore.viewTestSandboxCall(ctx, params);
        ctx.Log("testcore.viewTestSandboxCall ok");
    }
}
