// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

//nolint:dupl
package testcore

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/corecontracts/coreaccounts"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/corecontracts/coregovernance"
)

const (
	ContractNameDeployed = "exampleDeployTR"
	MsgCoreOnlyPanic     = "========== core only ========="
	MsgFullPanic         = "========== panic FULL ENTRY POINT ========="
	MsgViewPanic         = "========== panic VIEW ========="
)

func funcCallOnChain(ctx wasmlib.ScFuncContext, f *CallOnChainContext) {
	paramIn := f.Params.IntValue().Value()

	hnameContract := ctx.Contract()
	if f.Params.HnameContract().Exists() {
		hnameContract = f.Params.HnameContract().Value()
	}

	hnameEP := HFuncCallOnChain
	if f.Params.HnameEP().Exists() {
		hnameEP = f.Params.HnameEP().Value()
	}

	counter := f.State.Counter()

	ctx.Log("call depth = " + f.Params.IntValue().String() +
		", hnameContract = " + hnameContract.String() +
		", hnameEP = " + hnameEP.String() +
		", counter = " + counter.String())

	counter.SetValue(counter.Value() + 1)

	params := wasmlib.NewScMutableMap()
	params.GetInt64(ParamIntValue).SetValue(paramIn)
	ret := ctx.Call(hnameContract, hnameEP, params, nil)
	retVal := ret.GetInt64(ResultIntValue).Value()
	f.Results.IntValue().SetValue(retVal)
}

func funcCheckContextFromFullEP(ctx wasmlib.ScFuncContext, f *CheckContextFromFullEPContext) {
	ctx.Require(f.Params.AgentID().Value() == ctx.AccountID(), "fail: agentID")
	ctx.Require(f.Params.Caller().Value() == ctx.Caller(), "fail: caller")
	ctx.Require(f.Params.ChainID().Value() == ctx.ChainID(), "fail: chainID")
	ctx.Require(f.Params.ChainOwnerID().Value() == ctx.ChainOwnerID(), "fail: chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Value() == ctx.ContractCreator(), "fail: contractCreator")
}

//nolint:unparam
func funcDoNothing(ctx wasmlib.ScFuncContext, f *DoNothingContext) {
	ctx.Log("doing nothing...")
}

func funcGetMintedSupply(ctx wasmlib.ScFuncContext, f *GetMintedSupplyContext) {
	minted := ctx.Minted()
	mintedColors := minted.Colors()
	ctx.Require(mintedColors.Length() == 1, "test only supports one minted color")
	color := mintedColors.GetColor(0).Value()
	amount := minted.Balance(color)
	f.Results.MintedColor().SetValue(color)
	f.Results.MintedSupply().SetValue(amount)
}

func funcIncCounter(ctx wasmlib.ScFuncContext, f *IncCounterContext) {
	counter := f.State.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Fail().Exists() {
		ctx.Panic("failing on purpose")
	}
}

func funcPassTypesFull(ctx wasmlib.ScFuncContext, f *PassTypesFullContext) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(f.Params.Hash().Value() == hash, "Hash wrong")
	ctx.Require(f.Params.Int64().Value() == 42, "int64 wrong")
	ctx.Require(f.Params.Int64Zero().Value() == 0, "int64-0 wrong")
	ctx.Require(f.Params.String().Value() == string(ParamString), "string wrong")
	ctx.Require(f.Params.StringZero().Value() == "", "string-0 wrong")
	ctx.Require(f.Params.Hname().Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(f.Params.HnameZero().Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

func funcRunRecursion(ctx wasmlib.ScFuncContext, f *RunRecursionContext) {
	depth := f.Params.IntValue().Value()
	if depth <= 0 {
		return
	}

	callOnChain := ScFuncs.CallOnChain(ctx)
	callOnChain.Params.IntValue().SetValue(depth - 1)
	callOnChain.Params.HnameEP().SetValue(HFuncRunRecursion)
	callOnChain.Func.Call()
	retVal := callOnChain.Results.IntValue().Value()
	f.Results.IntValue().SetValue(retVal)
}

func funcSendToAddress(ctx wasmlib.ScFuncContext, f *SendToAddressContext) {
	balances := wasmlib.NewScTransfersFromBalances(ctx.Balances())
	ctx.TransferToAddress(f.Params.Address().Value(), balances)
}

func funcSetInt(ctx wasmlib.ScFuncContext, f *SetIntContext) {
	f.State.Ints().GetInt64(f.Params.Name().Value()).SetValue(f.Params.IntValue().Value())
}

//nolint:unparam
func funcTestCallPanicFullEP(ctx wasmlib.ScFuncContext, f *TestCallPanicFullEPContext) {
	ScFuncs.TestPanicFullEP(ctx).Func.Call()
}

//nolint:unparam
func funcTestCallPanicViewEPFromFull(ctx wasmlib.ScFuncContext, f *TestCallPanicViewEPFromFullContext) {
	ScFuncs.TestPanicViewEP(ctx).Func.Call()
}

func funcTestChainOwnerIDFull(ctx wasmlib.ScFuncContext, f *TestChainOwnerIDFullContext) {
	f.Results.ChainOwnerID().SetValue(ctx.ChainOwnerID())
}

//nolint:unparam
func funcTestEventLogDeploy(ctx wasmlib.ScFuncContext, f *TestEventLogDeployContext) {
	// deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("testcore"))
	ctx.Deploy(programHash, ContractNameDeployed, "test contract deploy log", nil)
}

//nolint:unparam
func funcTestEventLogEventData(ctx wasmlib.ScFuncContext, f *TestEventLogEventDataContext) {
	ctx.Event("[Event] - Testing Event...")
}

func funcTestEventLogGenericData(ctx wasmlib.ScFuncContext, f *TestEventLogGenericDataContext) {
	event := "[GenericData] Counter Number: " + f.Params.Counter().String()
	ctx.Event(event)
}

//nolint:unparam
func funcTestPanicFullEP(ctx wasmlib.ScFuncContext, f *TestPanicFullEPContext) {
	ctx.Panic(MsgFullPanic)
}

func funcWithdrawToChain(ctx wasmlib.ScFuncContext, f *WithdrawToChainContext) {
	coreaccounts.ScFuncs.Withdraw(ctx).Func.TransferIotas(1).PostToChain(f.Params.ChainID().Value())
}

func viewCheckContextFromViewEP(ctx wasmlib.ScViewContext, f *CheckContextFromViewEPContext) {
	ctx.Require(f.Params.AgentID().Value() == ctx.AccountID(), "fail: agentID")
	ctx.Require(f.Params.ChainID().Value() == ctx.ChainID(), "fail: chainID")
	ctx.Require(f.Params.ChainOwnerID().Value() == ctx.ChainOwnerID(), "fail: chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Value() == ctx.ContractCreator(), "fail: contractCreator")
}

func viewFibonacci(ctx wasmlib.ScViewContext, f *FibonacciContext) {
	n := f.Params.IntValue().Value()
	if n == 0 || n == 1 {
		f.Results.IntValue().SetValue(n)
		return
	}

	fib := ScFuncs.Fibonacci(ctx)
	fib.Params.IntValue().SetValue(n - 1)
	fib.Func.Call()
	n1 := fib.Results.IntValue().Value()

	fib.Params.IntValue().SetValue(n - 2)
	fib.Func.Call()
	n2 := fib.Results.IntValue().Value()

	f.Results.IntValue().SetValue(n1 + n2)
}

func viewGetCounter(ctx wasmlib.ScViewContext, f *GetCounterContext) {
	f.Results.Counter().SetValue(f.State.Counter().Value())
}

func viewGetInt(ctx wasmlib.ScViewContext, f *GetIntContext) {
	name := f.Params.Name().Value()
	value := f.State.Ints().GetInt64(name)
	ctx.Require(value.Exists(), "param 'value' not found")
	f.Results.Values().GetInt64(name).SetValue(value.Value())
}

//nolint:unparam
func viewJustView(ctx wasmlib.ScViewContext, f *JustViewContext) {
	ctx.Log("doing nothing...")
}

func viewPassTypesView(ctx wasmlib.ScViewContext, f *PassTypesViewContext) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(f.Params.Hash().Value() == hash, "Hash wrong")
	ctx.Require(f.Params.Int64().Value() == 42, "int64 wrong")
	ctx.Require(f.Params.Int64Zero().Value() == 0, "int64-0 wrong")
	ctx.Require(f.Params.String().Value() == string(ParamString), "string wrong")
	ctx.Require(f.Params.StringZero().Value() == "", "string-0 wrong")
	ctx.Require(f.Params.Hname().Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(f.Params.HnameZero().Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

//nolint:unparam
func viewTestCallPanicViewEPFromView(ctx wasmlib.ScViewContext, f *TestCallPanicViewEPFromViewContext) {
	ScFuncs.TestPanicViewEP(ctx).Func.Call()
}

func viewTestChainOwnerIDView(ctx wasmlib.ScViewContext, f *TestChainOwnerIDViewContext) {
	f.Results.ChainOwnerID().SetValue(ctx.ChainOwnerID())
}

//nolint:unparam
func viewTestPanicViewEP(ctx wasmlib.ScViewContext, f *TestPanicViewEPContext) {
	ctx.Panic(MsgViewPanic)
}

func viewTestSandboxCall(ctx wasmlib.ScViewContext, f *TestSandboxCallContext) {
	getChainInfo := coregovernance.ScFuncs.GetChainInfo(ctx)
	getChainInfo.Func.Call()
	f.Results.SandboxCall().SetValue(getChainInfo.Results.Description().Value())
}

//nolint:unparam
func funcTestBlockContext1(ctx wasmlib.ScFuncContext, f *TestBlockContext1Context) {
	ctx.Panic(MsgCoreOnlyPanic)
}

//nolint:unparam
func funcTestBlockContext2(ctx wasmlib.ScFuncContext, f *TestBlockContext2Context) {
	ctx.Panic(MsgCoreOnlyPanic)
}

//nolint:unparam
func viewGetStringValue(ctx wasmlib.ScViewContext, f *GetStringValueContext) {
	ctx.Panic(MsgCoreOnlyPanic)
}

func funcSpawn(ctx wasmlib.ScFuncContext, f *SpawnContext) {
	spawnName := ScName + "_spawned"
	spawnDescr := "spawned contract description"
	ctx.Deploy(f.Params.ProgHash().Value(), spawnName, spawnDescr, nil)

	spawnHname := wasmlib.NewScHname(spawnName)
	for i := 0; i < 5; i++ {
		ctx.Call(spawnHname, HFuncIncCounter, nil, nil)
	}
}
