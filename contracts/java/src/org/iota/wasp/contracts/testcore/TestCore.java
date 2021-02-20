package org.iota.wasp.contracts.testcore;

import org.iota.wasp.contracts.testcore.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.nio.charset.*;

public class TestCore {

	private static String ContractNameDeployed = "exampleDeployTR";
	private static String MsgFullPanic = "========== panic FULL ENTRY POINT =========";
	private static String MsgViewPanic = "========== panic VIEW =========";

	public static void funcCallOnChain(ScFuncContext ctx, FuncCallOnChainParams params) {
		ctx.Log("calling callOnChain");

		long paramInt = params.IntValue.Value();

		ScHname targetContract = ctx.ContractId().Hname();
		if (params.HnameContract.Exists()) {
			targetContract = params.HnameContract.Value();
		}

		ScHname targetEp = Consts.HFuncCallOnChain;
		if (params.HnameEP.Exists()) {
			targetEp = params.HnameEP.Value();
		}

		ScMutableInt varCounter = ctx.State().GetInt(Consts.VarCounter);
		long counter = varCounter.Value();
		varCounter.SetValue(counter + 1);

		ctx.Log("call depth = " + paramInt + " hnameContract = " + targetContract + " hnameEP = " + targetEp + " counter = " + counter);

		ScMutableMap parms = new ScMutableMap();
		parms.GetInt(Consts.ParamIntValue).SetValue(paramInt);
		ScImmutableMap ret = ctx.Call(targetContract, targetEp, parms, null);

		ScImmutableInt retVal = ret.GetInt(Consts.ParamIntValue);
		ctx.Results().GetInt(Consts.ParamIntValue).SetValue(retVal.Value());
	}

	public static void funcCheckContextFromFullEP(ScFuncContext ctx, FuncCheckContextFromFullEPParams params) {
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

		ScHash hash = ctx.Utility().HashBlake2b(Consts.ParamHash.toString().getBytes(StandardCharsets.UTF_8));
		ctx.Require(params.Hash.Value().equals(hash), "Hash wrong");
		ctx.Require(params.Int64.Value() == 42, "int64 wrong");
		ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
		ctx.Require(params.String.Value().equals(Consts.ParamString.toString()), "string wrong");
		ctx.Require(params.StringZero.Value().equals(""), "string-0 wrong");
		ctx.Require(params.Hname.Value().equals(new ScHname(Consts.ParamHname.toString())), "Hname wrong");
		ctx.Require(params.HnameZero.Value().equals(new ScHname(0)), "Hname-0 wrong");
	}

	public static void funcRunRecursion(ScFuncContext ctx, FuncRunRecursionParams params) {
		ctx.Log("calling runRecursion");

		long depth = params.IntValue.Value();
		if (depth <= 0) {
			return;
		}

		ScMutableMap parms = new ScMutableMap();
		parms.GetInt(Consts.ParamIntValue).SetValue(depth - 1);
		parms.GetHname(Consts.ParamHnameEP).SetValue(Consts.HFuncRunRecursion);
		ctx.CallSelf(Consts.HFuncCallOnChain, parms, null);
		// TODO how would I return result of the call ???
		ctx.Results().GetInt(Consts.ParamIntValue).SetValue(depth - 1);
	}

	public static void funcSendToAddress(ScFuncContext ctx, FuncSendToAddressParams params) {
		ctx.Log("calling sendToAddress");
		ScTransfers balances = new ScTransfers(ctx.Balances());
		ctx.TransferToAddress(params.Address.Value(), balances);
	}

	public static void funcSetInt(ScFuncContext ctx, FuncSetIntParams params) {
		ctx.Log("calling setInt");
		ctx.State().GetInt(new Key(params.Name.Value())).SetValue(params.IntValue.Value());
	}

	public static void funcTestCallPanicFullEP(ScFuncContext ctx, FuncTestCallPanicFullEPParams params) {
		ctx.Log("calling testCallPanicFullEP");
		ctx.CallSelf(Consts.HFuncTestPanicFullEP, null, null);
	}

	public static void funcTestCallPanicViewEPFromFull(ScFuncContext ctx, FuncTestCallPanicViewEPFromFullParams params) {
		ctx.Log("calling testCallPanicViewEPFromFull");
		ctx.CallSelf(Consts.HViewTestPanicViewEP, null, null);
	}

	public static void funcTestChainOwnerIDFull(ScFuncContext ctx, FuncTestChainOwnerIDFullParams params) {
		ctx.Log("calling testChainOwnerIDFull");
		ctx.Results().GetAgentId(Consts.ParamChainOwnerId).SetValue(ctx.ChainOwnerId());
	}

	public static void funcTestContractIDFull(ScFuncContext ctx, FuncTestContractIDFullParams params) {
		ctx.Log("calling testContractIDFull");
		ctx.Results().GetContractId(Consts.ParamContractId).SetValue(ctx.ContractId());
	}

	public static void funcTestEventLogDeploy(ScFuncContext ctx, FuncTestEventLogDeployParams params) {
		ctx.Log("calling testEventLogDeploy");
		//Deploy the same contract with another name
		ScHash programHash = ctx.Utility().HashBlake2b("test_sandbox".getBytes(StandardCharsets.UTF_8));
		ctx.Deploy(programHash, ContractNameDeployed,
				"test contract deploy log", null);
	}

	public static void funcTestEventLogEventData(ScFuncContext ctx, FuncTestEventLogEventDataParams params) {
		ctx.Log("calling testEventLogEventData");
		ctx.Event("[Event] - Testing Event...");
	}

	public static void funcTestEventLogGenericData(ScFuncContext ctx, FuncTestEventLogGenericDataParams params) {
		ctx.Log("calling testEventLogGenericData");
		String event = "[GenericData] Counter Number: " + params.Counter;
		ctx.Event(event);
	}

	public static void funcTestPanicFullEP(ScFuncContext ctx, FuncTestPanicFullEPParams params) {
		ctx.Log("calling testPanicFullEP");
		ctx.Panic(MsgFullPanic);
	}

	public static void funcWithdrawToChain(ScFuncContext ctx, FuncWithdrawToChainParams params) {
		ctx.Log("calling withdrawToChain");

		//Deploy the same contract with another name
		ScContractId targetContractId = new ScContractId(params.ChainId.Value(), Core.Accounts);
		ScTransfers transfers = new ScTransfers(ScColor.IOTA, 2);
		PostRequestParams req = new PostRequestParams();
		req.ContractId = targetContractId;
		req.Function = Core.AccountsFuncWithdrawToChain;
		req.Transfer = transfers;
		ctx.Post(req);
		ctx.Log("====  success ====");
		// TODO how to check if post was successful
	}

	public static void viewCheckContextFromViewEP(ScViewContext ctx, ViewCheckContextFromViewEPParams params) {
		ctx.Log("calling checkContextFromViewEP");

		ctx.Require(params.ChainId.Value() == ctx.ContractId().ChainId(), "fail: chainID");
		ctx.Require(params.ChainOwnerId.Value() == ctx.ChainOwnerId(), "fail: chainOwnerID");
		ctx.Require(params.ContractId.Value() == ctx.ContractId(), "fail: contractID");
		ctx.Require(params.AgentId.Value() == ctx.ContractId().AsAgentId(), "fail: agentID");
		ctx.Require(params.ContractCreator.Value() == ctx.ContractCreator(), "fail: contractCreator");
	}

	public static void viewFibonacci(ScViewContext ctx, ViewFibonacciParams params) {
		ctx.Log("calling fibonacci");

		long n = params.IntValue.Value();
		if (n == 0 || n == 1) {
			ctx.Results().GetInt(Consts.ParamIntValue).SetValue(n);
			return;
		}
		ScMutableMap parms1 = new ScMutableMap();
		parms1.GetInt(Consts.ParamIntValue).SetValue(n - 1);
		ScImmutableMap results1 = ctx.CallSelf(Consts.HViewFibonacci, parms1);
		long n1 = results1.GetInt(Consts.ParamIntValue).Value();

		ScMutableMap parms2 = new ScMutableMap();
		parms2.GetInt(Consts.ParamIntValue).SetValue(n - 2);
		ScImmutableMap results2 = ctx.CallSelf(Consts.HViewFibonacci, parms2);
		long n2 = results2.GetInt(Consts.ParamIntValue).Value();

		ctx.Results().GetInt(Consts.ParamIntValue).SetValue(n1 + n2);
	}

	public static void viewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
		ctx.Log("calling getCounter");
		ScImmutableInt counter = ctx.State().GetInt(Consts.VarCounter);
		ctx.Results().GetInt(Consts.VarCounter).SetValue(counter.Value());
	}

	public static void viewGetInt(ScViewContext ctx, ViewGetIntParams params) {
		ctx.Log("calling getInt");

		String name = params.Name.Value();
		ScImmutableInt value = ctx.State().GetInt(new Key(name));
		ctx.Require(value.Exists(), "param 'value' not found");
		ctx.Results().GetInt(new Key(name)).SetValue(value.Value());
	}

	public static void viewJustView(ScViewContext ctx, ViewJustViewParams params) {
		ctx.Log("calling justView");
	}

	public static void viewPassTypesView(ScViewContext ctx, ViewPassTypesViewParams params) {
		ctx.Log("calling passTypesView");

		ScHash hash = ctx.Utility().HashBlake2b(Consts.ParamHash.toString().getBytes(StandardCharsets.UTF_8));
		ctx.Require(params.Hash.Value().equals(hash), "Hash wrong");
		ctx.Require(params.Int64.Value() == 42, "int64 wrong");
		ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong");
		ctx.Require(params.String.Value().equals(Consts.ParamString.toString()), "string wrong");
		ctx.Require(params.StringZero.Value().equals(""), "string-0 wrong");
		ctx.Require(params.Hname.Value().equals(new ScHname(Consts.ParamHname.toString())), "Hname wrong");
		ctx.Require(params.HnameZero.Value().equals(new ScHname(0)), "Hname-0 wrong");
	}

	public static void viewTestCallPanicViewEPFromView(ScViewContext ctx, ViewTestCallPanicViewEPFromViewParams params) {
		ctx.Log("calling testCallPanicViewEPFromView");
		ctx.CallSelf(Consts.HViewTestPanicViewEP, null);
	}

	public static void viewTestChainOwnerIDView(ScViewContext ctx, ViewTestChainOwnerIDViewParams params) {
		ctx.Log("calling testChainOwnerIDView");
		ctx.Results().GetAgentId(Consts.ParamChainOwnerId).SetValue(ctx.ChainOwnerId());
	}

	public static void viewTestContractIDView(ScViewContext ctx, ViewTestContractIDViewParams params) {
		ctx.Log("calling testContractIDView");
		ctx.Results().GetContractId(Consts.ParamContractId).SetValue(ctx.ContractId());
	}

	public static void viewTestPanicViewEP(ScViewContext ctx, ViewTestPanicViewEPParams params) {
		ctx.Log("calling testPanicViewEP");
		ctx.Panic(MsgViewPanic);
	}

	public static void viewTestSandboxCall(ScViewContext ctx, ViewTestSandboxCallParams params) {
		ctx.Log("calling testSandboxCall");
		ScImmutableMap ret = ctx.Call(Core.Root, Core.RootViewGetChainInfo, null);
		String desc = ret.GetString(new Key("d")).Value();
		ctx.Results().GetString(new Key("sandboxCall")).SetValue(desc);
	}
}
