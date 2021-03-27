// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.testcore;

import org.iota.wasp.contracts.testcore.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.nio.charset.*;

public class TestCore {

    static String ContractNameDeployed = "exampleDeployTR";
    static String MsgFullPanic = "========== panic FULL ENTRY POINT =========";
    static String MsgViewPanic = "========== panic VIEW =========";

    public static void funcCallOnChain(ScFuncContext ctx, FuncCallOnChainParams params) {
        var paramInt = params.IntValue.Value();

        var targetContract = ctx.Contract();
        if (params.HnameContract.Exists()) {
            targetContract = params.HnameContract.Value();
        }

        var targetEp = Consts.HFuncCallOnChain;
        if (params.HnameEP.Exists()) {
            targetEp = params.HnameEP.Value();
        }

        var varCounter = ctx.State().GetInt64(Consts.VarCounter);
        var counter = varCounter.Value();
        varCounter.SetValue(counter + 1);

        ctx.Log("call depth = " + paramInt +
                " hnameContract = " + targetContract +
                " hnameEP = " + targetEp +
                " counter = " + counter);

        var parms = new ScMutableMap();
        parms.GetInt64(Consts.ParamIntValue).SetValue(paramInt);
        var ret = ctx.Call(targetContract, targetEp, parms, null);

        var retVal = ret.GetInt64(Consts.ParamIntValue);
        ctx.Results().GetInt64(Consts.ParamIntValue).SetValue(retVal.Value());
    }

    public static void funcCheckContextFromFullEP(ScFuncContext ctx, FuncCheckContextFromFullEPParams params) {
        ctx.Require(params.AgentId.Value().equals(ctx.AccountId()), "fail: agentID");
        ctx.Require(params.Caller.Value().equals(ctx.Caller()), "fail: caller");
        ctx.Require(params.ChainId.Value().equals(ctx.ChainId()), "fail: chainID");
        ctx.Require(params.ChainOwnerId.Value().equals(ctx.ChainOwnerId()), "fail: chainOwnerID");
        ctx.Require(params.ContractCreator.Value().equals(ctx.ContractCreator()), "fail: contractCreator");
    }

    public static void funcDoNothing(ScFuncContext ctx, FuncDoNothingParams params) {
        ctx.Log("doing nothing...");
    }

    public static void funcGetMintedSupply(ScFuncContext ctx, FuncGetMintedSupplyParams params) {
        var minted = ctx.Minted();
        var mintedColors = minted.Colors();
        ctx.Require(mintedColors.Length() == 1, "test only supports one minted color");
        var color = mintedColors.GetColor(0).Value();
        var amount = minted.Balance(color);
        ctx.Results().GetInt64(Consts.VarMintedSupply).SetValue(amount);
        ctx.Results().GetColor(Consts.VarMintedColor).SetValue(color);
    }

    public static void funcIncCounter(ScFuncContext ctx, FuncIncCounterParams params) {
        ctx.State().GetInt64(Consts.VarCounter).SetValue(ctx.State().GetInt64(Consts.VarCounter).Value() + 1);
    }

    public static void funcInit(ScFuncContext ctx, FuncInitParams params) {
        ctx.Log("doing nothing...");
    }

    public static void funcPassTypesFull(ScFuncContext ctx, FuncPassTypesFullParams params) {
        var hash = ctx.Utility().HashBlake2b(Consts.ParamHash.toString().getBytes(StandardCharsets.UTF_8));
        ctx.Require(params.Hash.Value().equals(hash), "Hash wrong");
        ctx.Require(params.Int64.Value() == 42, "int64 wrong");
        ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
        ctx.Require(params.String.Value().equals(Consts.ParamString.toString()), "string wrong");
        ctx.Require(params.StringZero.Value().equals(""), "string-0 wrong");
        ctx.Require(params.Hname.Value().equals(new ScHname(Consts.ParamHname.toString())), "Hname wrong");
        ctx.Require(params.HnameZero.Value().equals(new ScHname(0)), "Hname-0 wrong");
    }

    public static void funcRunRecursion(ScFuncContext ctx, FuncRunRecursionParams params) {
        var depth = params.IntValue.Value();
        if (depth <= 0) {
            return;
        }

        var parms = new ScMutableMap();
        parms.GetInt64(Consts.ParamIntValue).SetValue(depth - 1);
        parms.GetHname(Consts.ParamHnameEP).SetValue(Consts.HFuncRunRecursion);
        ctx.CallSelf(Consts.HFuncCallOnChain, parms, null);
        // TODO how would I return result of the call ???
        ctx.Results().GetInt64(Consts.ParamIntValue).SetValue(depth - 1);
    }

    public static void funcSendToAddress(ScFuncContext ctx, FuncSendToAddressParams params) {
        var balances = new ScTransfers(ctx.Balances());
        ctx.TransferToAddress(params.Address.Value(), balances);
    }

    public static void funcSetInt(ScFuncContext ctx, FuncSetIntParams params) {
        ctx.State().GetInt64(new Key(params.Name.Value())).SetValue(params.IntValue.Value());
    }

    public static void funcTestCallPanicFullEP(ScFuncContext ctx, FuncTestCallPanicFullEPParams params) {
        ctx.CallSelf(Consts.HFuncTestPanicFullEP, null, null);
    }

    public static void funcTestCallPanicViewEPFromFull(ScFuncContext ctx, FuncTestCallPanicViewEPFromFullParams params) {
        ctx.CallSelf(Consts.HViewTestPanicViewEP, null, null);
    }

    public static void funcTestChainOwnerIDFull(ScFuncContext ctx, FuncTestChainOwnerIDFullParams params) {
        ctx.Results().GetAgentId(Consts.ParamChainOwnerId).SetValue(ctx.ChainOwnerId());
    }

    public static void funcTestEventLogDeploy(ScFuncContext ctx, FuncTestEventLogDeployParams params) {
        //Deploy the same contract with another name
        var programHash = ctx.Utility().HashBlake2b("test_sandbox".getBytes(StandardCharsets.UTF_8));
        ctx.Deploy(programHash, ContractNameDeployed, "test contract deploy log", null);
    }

    public static void funcTestEventLogEventData(ScFuncContext ctx, FuncTestEventLogEventDataParams params) {
        ctx.Event("[Event] - Testing Event...");
    }

    public static void funcTestEventLogGenericData(ScFuncContext ctx, FuncTestEventLogGenericDataParams params) {
        var event = "[GenericData] Counter Number: " + params.Counter;
        ctx.Event(event);
    }

    public static void funcTestPanicFullEP(ScFuncContext ctx, FuncTestPanicFullEPParams params) {
        ctx.Panic(MsgFullPanic);
    }

    public static void funcWithdrawToChain(ScFuncContext ctx, FuncWithdrawToChainParams params) {
        var transfer = ScTransfers.iotas(1);
        ctx.Post(params.ChainId.Value(), Core.Accounts, Core.AccountsFuncWithdraw, null, transfer, 0);
        ctx.Log("====  success ====");
    }

    public static void viewCheckContextFromViewEP(ScViewContext ctx, ViewCheckContextFromViewEPParams params) {
        ctx.Require(params.AgentId.Value().equals(ctx.AccountId()), "fail: agentID");
        ctx.Require(params.ChainId.Value().equals(ctx.ChainId()), "fail: chainID");
        ctx.Require(params.ChainOwnerId.Value().equals(ctx.ChainOwnerId()), "fail: chainOwnerID");
        ctx.Require(params.ContractCreator.Value().equals(ctx.ContractCreator()), "fail: contractCreator");
    }

    public static void viewFibonacci(ScViewContext ctx, ViewFibonacciParams params) {
        var n = params.IntValue.Value();
        if (n == 0 || n == 1) {
            ctx.Results().GetInt64(Consts.ParamIntValue).SetValue(n);
            return;
        }
        var parms1 = new ScMutableMap();
        parms1.GetInt64(Consts.ParamIntValue).SetValue(n - 1);
        var results1 = ctx.CallSelf(Consts.HViewFibonacci, parms1);
        var n1 = results1.GetInt64(Consts.ParamIntValue).Value();

        var parms2 = new ScMutableMap();
        parms2.GetInt64(Consts.ParamIntValue).SetValue(n - 2);
        var results2 = ctx.CallSelf(Consts.HViewFibonacci, parms2);
        var n2 = results2.GetInt64(Consts.ParamIntValue).Value();

        ctx.Results().GetInt64(Consts.ParamIntValue).SetValue(n1 + n2);
    }

    public static void viewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        ctx.Results().GetInt64(Consts.VarCounter).SetValue(counter.Value());
    }

    public static void viewGetInt(ScViewContext ctx, ViewGetIntParams params) {
        var name = params.Name.Value();
        var value = ctx.State().GetInt64(new Key(name));
        ctx.Require(value.Exists(), "param 'value' not found");
        ctx.Results().GetInt64(new Key(name)).SetValue(value.Value());
    }

    public static void viewJustView(ScViewContext ctx, ViewJustViewParams params) {
        ctx.Log("doing nothing...");
    }

    public static void viewPassTypesView(ScViewContext ctx, ViewPassTypesViewParams params) {
        var hash = ctx.Utility().HashBlake2b(Consts.ParamHash.toString().getBytes(StandardCharsets.UTF_8));
        ctx.Require(params.Hash.Value().equals(hash), "Hash wrong");
        ctx.Require(params.Int64.Value() == 42, "int64 wrong");
        ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
        ctx.Require(params.String.Value().equals(Consts.ParamString.toString()), "string wrong");
        ctx.Require(params.StringZero.Value().equals(""), "string-0 wrong");
        ctx.Require(params.Hname.Value().equals(new ScHname(Consts.ParamHname.toString())), "Hname wrong");
        ctx.Require(params.HnameZero.Value().equals(new ScHname(0)), "Hname-0 wrong");
    }

    public static void viewTestCallPanicViewEPFromView(ScViewContext ctx, ViewTestCallPanicViewEPFromViewParams params) {
        ctx.CallSelf(Consts.HViewTestPanicViewEP, null);
    }

    public static void viewTestChainOwnerIDView(ScViewContext ctx, ViewTestChainOwnerIDViewParams params) {
        ctx.Results().GetAgentId(Consts.ParamChainOwnerId).SetValue(ctx.ChainOwnerId());
    }

    public static void viewTestPanicViewEP(ScViewContext ctx, ViewTestPanicViewEPParams params) {
        ctx.Panic(MsgViewPanic);
    }

    public static void viewTestSandboxCall(ScViewContext ctx, ViewTestSandboxCallParams params) {
        var ret = ctx.Call(Core.Root, Core.RootViewGetChainInfo, null);
        var desc = ret.GetString(new Key("d")).Value();
        ctx.Results().GetString(new Key("sandboxCall")).SetValue(desc);
    }
}
