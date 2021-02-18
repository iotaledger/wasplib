// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.testcore;

public class Testcore {

const ContractNameDeployed: &str = "exampleDeployTR";
const MsgFullPanic: &str = "========== panic FULL ENTRY POINT =========";
const MsgViewPanic: &str = "========== panic VIEW =========";

public static void funcCallOnChain(ScFuncContext ctx, FuncCallOnChainParams params) {
    ctx.Log("calling callOnChain");

    paramInt = params.IntValue.Value();

    targetContract = ctx.ContractId().Hname();
    if (params.HnameContract.Exists()) {
        targetContract = params.HnameContract.Value()
    }

    targetEp = HFuncCallOnChain;
    if (params.HnameEp.Exists()) {
        targetEp = params.HnameEp.Value()
    }

    varCounter = ctx.State().GetInt(VarCounter);
    counter = varCounter.Value();
    varCounter.SetValue(counter + 1);

    ctx.Log(format!("call depth = {} hnameContract = {} hnameEP = {} counter = {}",
                     paramInt, targetContract.toString(), targetEp.toString(), counter));

    parms = new ScMutableMapp();
    parms.GetInt(ParamIntValue).SetValue(paramInt);
    ret = ctx.Call(targetContract, targetEp, Some(parms), null);

    retVal = ret.GetInt(ParamIntValue);
    ctx.Results().GetInt(ParamIntValue).SetValue(retVal.Value());
}

public static void funcCheckContextFromFullEp(ScFuncContext ctx, FuncCheckContextFromFullEPParams params) {
    ctx.Log("calling checkContextFromFullEP");

    ctx.Require(params.ChainId.Value() == ctx.ContractId().ChainId(), "fail: chainID");
    ctx.Require(params.ChainOwnerId.Value() == ctx.ChainOwnerId(), "fail: chainOwnerID");
    ctx.Require(params.Caller.Value() == ctx.Caller(), "fail: caller");
    ctx.Require(params.ContractId.Value() == ctx.ContractId(), "fail: contractID");
    ctx.Require(params.AgentId.Value() == ctx.ContractId().AsAgentId(), "fail: agentID");
    ctx.Require(params.ContractCreator.Value() == ctx.ContractCreator(), "fail: contractCreator");
}

public static void funcDoNothing(ScFuncContext ctx, FuncDoNothingParams params) {
    ctx.Log("calling doNothing");
}

public static void funcInit(ScFuncContext ctx, FuncInitParams params) {
    ctx.Log("calling init");
}

public static void funcPassTypesFull(ScFuncContext ctx, FuncPassTypesFullParams params) {
    ctx.Log("calling passTypesFull");

    hash = ctx.Utility().HashBlake2b(ParamHash.AsBytes());
    ctx.Require(params.Hash.Value() == hash, "Hash wrong");
    ctx.Require(params.Int64.Value() == 42, "int64 wrong");
    ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
    ctx.Require(params.String.Value() == ParamString, "string wrong");
    ctx.Require(params.Hname.Value() == new ScHnamee(ParamHname), "Hname wrong");
    ctx.Require(params.HnameZero.Value() == ScHname(0), "Hname-0 wrong");
}

public static void funcRunRecursion(ScFuncContext ctx, FuncRunRecursionParams params) {
    ctx.Log("calling runRecursion");

    depth = params.IntValue.Value();
    if (depth <= 0) {
        return;
    }

    parms = new ScMutableMapp();
    parms.GetInt(ParamIntValue).SetValue(depth - 1);
    parms.GetHname(ParamHnameEp).SetValue(HFuncRunRecursion);
    ctx.CallSelf(HFuncCallOnChain, Some(parms), null);
    // TODO how would I return result of the call ???
    ctx.Results().GetInt(ParamIntValue).SetValue(depth - 1);
}

public static void funcSendToAddress(ScFuncContext ctx, FuncSendToAddressParams params) {
    ctx.Log("calling sendToAddress");
    ctx.TransferToAddress(params.Address.Value(), ctx.Balances());
}

public static void funcSetInt(ScFuncContext ctx, FuncSetIntParams params) {
    ctx.Log("calling setInt");
    ctx.State().GetInt(params.Name.Value()).SetValue(params.IntValue.Value());
}

public static void funcTestCallPanicFullEp(ScFuncContext ctx, FuncTestCallPanicFullEPParams params) {
    ctx.Log("calling testCallPanicFullEP");
    ctx.CallSelf(HFuncTestPanicFullEp, null, null);
}

public static void funcTestCallPanicViewEpFromFull(ScFuncContext ctx, FuncTestCallPanicViewEPFromFullParams params) {
    ctx.Log("calling testCallPanicViewEPFromFull");
    ctx.CallSelf(HViewTestPanicViewEp, null, null);
}

public static void funcTestChainOwnerIdFull(ScFuncContext ctx, FuncTestChainOwnerIDFullParams params) {
    ctx.Log("calling testChainOwnerIDFull");
    ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

public static void funcTestContractIdFull(ScFuncContext ctx, FuncTestContractIDFullParams params) {
    ctx.Log("calling testContractIDFull");
    ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId());
}

public static void funcTestEventLogDeploy(ScFuncContext ctx, FuncTestEventLogDeployParams params) {
    ctx.Log("calling testEventLogDeploy");
    //Deploy the same contract with another name
    programHash = ctx.Utility().HashBlake2b("test_sandbox".AsBytes());
    ctx.Deploy(programHash, ContractNameDeployed,
               "test contract deploy log", null)
}

public static void funcTestEventLogEventData(ScFuncContext ctx, FuncTestEventLogEventDataParams params) {
    ctx.Log("calling testEventLogEventData");
    ctx.Event("[Event] - Testing Event...");
}

public static void funcTestEventLogGenericData(ScFuncContext ctx, FuncTestEventLogGenericDataParams params) {
    ctx.Log("calling testEventLogGenericData");
    event = "[GenericData] Counter Number: " + params.Counter;
    ctx.Event(event)
}

public static void funcTestPanicFullEp(ScFuncContext ctx, FuncTestPanicFullEPParams params) {
    ctx.Log("calling testPanicFullEP");
    ctx.Panic(MsgFullPanic)
}

public static void funcWithdrawToChain(ScFuncContext ctx, FuncWithdrawToChainParams params) {
    ctx.Log("calling withdrawToChain");

    //Deploy the same contract with another name
    targetContractId = new ScContractId(params.ChainId.Value(), CoreAccounts);
    transfers = new ScTransfers(ScColor.IOTA, 2);
    ctx.Post(PostRequestParams {
        request.ContractId = targetContractId;
        request.Function = CoreAccountsFuncWithdrawToChain;
        request.Params = null;
        request.Transfer = Some(Box::new(transfers));
        request.Delay = 0;
    });
    ctx.Log("====  success ====");
    // TODO how to check if post was successful
}

public static void viewCheckContextFromViewEp(ScViewContext ctx, ViewCheckContextFromViewEPParams params) {
    ctx.Log("calling checkContextFromViewEP");

    ctx.Require(params.ChainId.Value() == ctx.ContractId().ChainId(), "fail: chainID");
    ctx.Require(params.ChainOwnerId.Value() == ctx.ChainOwnerId(), "fail: chainOwnerID");
    ctx.Require(params.ContractId.Value() == ctx.ContractId(), "fail: contractID");
    ctx.Require(params.AgentId.Value() == ctx.ContractId().AsAgentId(), "fail: agentID");
    ctx.Require(params.ContractCreator.Value() == ctx.ContractCreator(), "fail: contractCreator");
}

public static void viewFibonacci(ScViewContext ctx, ViewFibonacciParams params) {
    ctx.Log("calling fibonacci");

    n = params.IntValue.Value();
    if (n == 0 || n == 1) {
        ctx.Results().GetInt(ParamIntValue).SetValue(n);
        return;
    }
    parms1 = new ScMutableMapp();
    parms1.GetInt(ParamIntValue).SetValue(n - 1);
    results1 = ctx.CallSelf(HViewFibonacci, Some(parms1));
    n1 = results1.GetInt(ParamIntValue).Value();

    parms2 = new ScMutableMapp();
    parms2.GetInt(ParamIntValue).SetValue(n - 2);
    results2 = ctx.CallSelf(HViewFibonacci, Some(parms2));
    n2 = results2.GetInt(ParamIntValue).Value();

    ctx.Results().GetInt(ParamIntValue).SetValue(n1 + n2);
}

public static void viewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
    ctx.Log("calling getCounter");
    counter = ctx.State().GetInt(VarCounter);
    ctx.Results().GetInt(VarCounter).SetValue(counter.Value());
}

public static void viewGetInt(ScViewContext ctx, ViewGetIntParams params) {
    ctx.Log("calling getInt");

    name = params.Name.Value();
    value = ctx.State().GetInt(name);
    ctx.Require(value.Exists(), "param 'value' not found");
    ctx.Results().GetInt(name).SetValue(value.Value());
}

public static void viewJustView(ScViewContext ctx, ViewJustViewParams params) {
    ctx.Log("calling justView");
}

public static void viewPassTypesView(ScViewContext ctx, ViewPassTypesViewParams params) {
    ctx.Log("calling passTypesView");

    hash = ctx.Utility().HashBlake2b(ParamHash.AsBytes());
    ctx.Require(params.Hash.Value() == hash, "Hash wrong");
    ctx.Require(params.Int64.Value() == 42, "int64 wrong");
    ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
    ctx.Require(params.String.Value() == ParamString, "string wrong");
    ctx.Require(params.StringZero.Value() == "", "string-0 wrong");
    ctx.Require(params.Hname.Value() == new ScHnamee(ParamHname), "Hname wrong");
    ctx.Require(params.HnameZero.Value() == ScHname(0), "Hname-0 wrong");
}

public static void viewTestCallPanicViewEpFromView(ScViewContext ctx, ViewTestCallPanicViewEPFromViewParams params) {
    ctx.Log("calling testCallPanicViewEPFromView");
    ctx.CallSelf(HViewTestPanicViewEp, null);
}

public static void viewTestChainOwnerIdView(ScViewContext ctx, ViewTestChainOwnerIDViewParams params) {
    ctx.Log("calling testChainOwnerIDView");
    ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

public static void viewTestContractIdView(ScViewContext ctx, ViewTestContractIDViewParams params) {
    ctx.Log("calling testContractIDView");
    ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId());
}

public static void viewTestPanicViewEp(ScViewContext ctx, ViewTestPanicViewEPParams params) {
    ctx.Log("calling testPanicViewEP");
    ctx.Panic(MsgViewPanic)
}

public static void viewTestSandboxCall(ScViewContext ctx, ViewTestSandboxCallParams params) {
    ctx.Log("calling testSandboxCall");
    ret = ctx.Call(CoreRoot, CoreRootViewGetChainInfo, null);
    desc = ret.GetString("d").Value();
    ctx.Results().GetString("sandboxCall").SetValue(desc);
}
}
