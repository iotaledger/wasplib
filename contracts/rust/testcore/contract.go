// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type CallOnChainCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCallOnChainParams
	Results ImmutableCallOnChainResults
}

func NewCallOnChainCall(ctx wasmlib.ScFuncContext) *CallOnChainCall {
	f := &CallOnChainCall{Func: wasmlib.NewScFunc(HScName, HFuncCallOnChain)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

type CheckContextFromFullEPCall struct {
	Func   *wasmlib.ScFunc
	Params MutableCheckContextFromFullEPParams
}

func NewCheckContextFromFullEPCall(ctx wasmlib.ScFuncContext) *CheckContextFromFullEPCall {
	f := &CheckContextFromFullEPCall{Func: wasmlib.NewScFunc(HScName, HFuncCheckContextFromFullEP)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type DoNothingCall struct {
	Func *wasmlib.ScFunc
}

func NewDoNothingCall(ctx wasmlib.ScFuncContext) *DoNothingCall {
	return &DoNothingCall{Func: wasmlib.NewScFunc(HScName, HFuncDoNothing)}
}

type GetMintedSupplyCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableGetMintedSupplyResults
}

func NewGetMintedSupplyCall(ctx wasmlib.ScFuncContext) *GetMintedSupplyCall {
	f := &GetMintedSupplyCall{Func: wasmlib.NewScFunc(HScName, HFuncGetMintedSupply)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

type IncCounterCall struct {
	Func *wasmlib.ScFunc
}

func NewIncCounterCall(ctx wasmlib.ScFuncContext) *IncCounterCall {
	return &IncCounterCall{Func: wasmlib.NewScFunc(HScName, HFuncIncCounter)}
}

type InitCall struct {
	Func *wasmlib.ScFunc
}

func NewInitCall(ctx wasmlib.ScFuncContext) *InitCall {
	return &InitCall{Func: wasmlib.NewScFunc(HScName, HFuncInit)}
}

type PassTypesFullCall struct {
	Func   *wasmlib.ScFunc
	Params MutablePassTypesFullParams
}

func NewPassTypesFullCall(ctx wasmlib.ScFuncContext) *PassTypesFullCall {
	f := &PassTypesFullCall{Func: wasmlib.NewScFunc(HScName, HFuncPassTypesFull)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type RunRecursionCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableRunRecursionParams
	Results ImmutableRunRecursionResults
}

func NewRunRecursionCall(ctx wasmlib.ScFuncContext) *RunRecursionCall {
	f := &RunRecursionCall{Func: wasmlib.NewScFunc(HScName, HFuncRunRecursion)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

type SendToAddressCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSendToAddressParams
}

func NewSendToAddressCall(ctx wasmlib.ScFuncContext) *SendToAddressCall {
	f := &SendToAddressCall{Func: wasmlib.NewScFunc(HScName, HFuncSendToAddress)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type SetIntCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetIntParams
}

func NewSetIntCall(ctx wasmlib.ScFuncContext) *SetIntCall {
	f := &SetIntCall{Func: wasmlib.NewScFunc(HScName, HFuncSetInt)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type TestCallPanicFullEPCall struct {
	Func *wasmlib.ScFunc
}

func NewTestCallPanicFullEPCall(ctx wasmlib.ScFuncContext) *TestCallPanicFullEPCall {
	return &TestCallPanicFullEPCall{Func: wasmlib.NewScFunc(HScName, HFuncTestCallPanicFullEP)}
}

type TestCallPanicViewEPFromFullCall struct {
	Func *wasmlib.ScFunc
}

func NewTestCallPanicViewEPFromFullCall(ctx wasmlib.ScFuncContext) *TestCallPanicViewEPFromFullCall {
	return &TestCallPanicViewEPFromFullCall{Func: wasmlib.NewScFunc(HScName, HFuncTestCallPanicViewEPFromFull)}
}

type TestChainOwnerIDFullCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableTestChainOwnerIDFullResults
}

func NewTestChainOwnerIDFullCall(ctx wasmlib.ScFuncContext) *TestChainOwnerIDFullCall {
	f := &TestChainOwnerIDFullCall{Func: wasmlib.NewScFunc(HScName, HFuncTestChainOwnerIDFull)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

type TestEventLogDeployCall struct {
	Func *wasmlib.ScFunc
}

func NewTestEventLogDeployCall(ctx wasmlib.ScFuncContext) *TestEventLogDeployCall {
	return &TestEventLogDeployCall{Func: wasmlib.NewScFunc(HScName, HFuncTestEventLogDeploy)}
}

type TestEventLogEventDataCall struct {
	Func *wasmlib.ScFunc
}

func NewTestEventLogEventDataCall(ctx wasmlib.ScFuncContext) *TestEventLogEventDataCall {
	return &TestEventLogEventDataCall{Func: wasmlib.NewScFunc(HScName, HFuncTestEventLogEventData)}
}

type TestEventLogGenericDataCall struct {
	Func   *wasmlib.ScFunc
	Params MutableTestEventLogGenericDataParams
}

func NewTestEventLogGenericDataCall(ctx wasmlib.ScFuncContext) *TestEventLogGenericDataCall {
	f := &TestEventLogGenericDataCall{Func: wasmlib.NewScFunc(HScName, HFuncTestEventLogGenericData)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type TestPanicFullEPCall struct {
	Func *wasmlib.ScFunc
}

func NewTestPanicFullEPCall(ctx wasmlib.ScFuncContext) *TestPanicFullEPCall {
	return &TestPanicFullEPCall{Func: wasmlib.NewScFunc(HScName, HFuncTestPanicFullEP)}
}

type WithdrawToChainCall struct {
	Func   *wasmlib.ScFunc
	Params MutableWithdrawToChainParams
}

func NewWithdrawToChainCall(ctx wasmlib.ScFuncContext) *WithdrawToChainCall {
	f := &WithdrawToChainCall{Func: wasmlib.NewScFunc(HScName, HFuncWithdrawToChain)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type CheckContextFromViewEPCall struct {
	Func   *wasmlib.ScView
	Params MutableCheckContextFromViewEPParams
}

func NewCheckContextFromViewEPCall(ctx wasmlib.ScFuncContext) *CheckContextFromViewEPCall {
	f := &CheckContextFromViewEPCall{Func: wasmlib.NewScView(HScName, HViewCheckContextFromViewEP)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func NewCheckContextFromViewEPCallFromView(ctx wasmlib.ScViewContext) *CheckContextFromViewEPCall {
	return NewCheckContextFromViewEPCall(wasmlib.ScFuncContext{})
}

type FibonacciCall struct {
	Func    *wasmlib.ScView
	Params  MutableFibonacciParams
	Results ImmutableFibonacciResults
}

func NewFibonacciCall(ctx wasmlib.ScFuncContext) *FibonacciCall {
	f := &FibonacciCall{Func: wasmlib.NewScView(HScName, HViewFibonacci)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewFibonacciCallFromView(ctx wasmlib.ScViewContext) *FibonacciCall {
	return NewFibonacciCall(wasmlib.ScFuncContext{})
}

type GetCounterCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetCounterResults
}

func NewGetCounterCall(ctx wasmlib.ScFuncContext) *GetCounterCall {
	f := &GetCounterCall{Func: wasmlib.NewScView(HScName, HViewGetCounter)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func NewGetCounterCallFromView(ctx wasmlib.ScViewContext) *GetCounterCall {
	return NewGetCounterCall(wasmlib.ScFuncContext{})
}

type GetIntCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetIntParams
	Results ImmutableGetIntResults
}

func NewGetIntCall(ctx wasmlib.ScFuncContext) *GetIntCall {
	f := &GetIntCall{Func: wasmlib.NewScView(HScName, HViewGetInt)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewGetIntCallFromView(ctx wasmlib.ScViewContext) *GetIntCall {
	return NewGetIntCall(wasmlib.ScFuncContext{})
}

type JustViewCall struct {
	Func *wasmlib.ScView
}

func NewJustViewCall(ctx wasmlib.ScFuncContext) *JustViewCall {
	return &JustViewCall{Func: wasmlib.NewScView(HScName, HViewJustView)}
}

func NewJustViewCallFromView(ctx wasmlib.ScViewContext) *JustViewCall {
	return NewJustViewCall(wasmlib.ScFuncContext{})
}

type PassTypesViewCall struct {
	Func   *wasmlib.ScView
	Params MutablePassTypesViewParams
}

func NewPassTypesViewCall(ctx wasmlib.ScFuncContext) *PassTypesViewCall {
	f := &PassTypesViewCall{Func: wasmlib.NewScView(HScName, HViewPassTypesView)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func NewPassTypesViewCallFromView(ctx wasmlib.ScViewContext) *PassTypesViewCall {
	return NewPassTypesViewCall(wasmlib.ScFuncContext{})
}

type TestCallPanicViewEPFromViewCall struct {
	Func *wasmlib.ScView
}

func NewTestCallPanicViewEPFromViewCall(ctx wasmlib.ScFuncContext) *TestCallPanicViewEPFromViewCall {
	return &TestCallPanicViewEPFromViewCall{Func: wasmlib.NewScView(HScName, HViewTestCallPanicViewEPFromView)}
}

func NewTestCallPanicViewEPFromViewCallFromView(ctx wasmlib.ScViewContext) *TestCallPanicViewEPFromViewCall {
	return NewTestCallPanicViewEPFromViewCall(wasmlib.ScFuncContext{})
}

type TestChainOwnerIDViewCall struct {
	Func    *wasmlib.ScView
	Results ImmutableTestChainOwnerIDViewResults
}

func NewTestChainOwnerIDViewCall(ctx wasmlib.ScFuncContext) *TestChainOwnerIDViewCall {
	f := &TestChainOwnerIDViewCall{Func: wasmlib.NewScView(HScName, HViewTestChainOwnerIDView)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func NewTestChainOwnerIDViewCallFromView(ctx wasmlib.ScViewContext) *TestChainOwnerIDViewCall {
	return NewTestChainOwnerIDViewCall(wasmlib.ScFuncContext{})
}

type TestPanicViewEPCall struct {
	Func *wasmlib.ScView
}

func NewTestPanicViewEPCall(ctx wasmlib.ScFuncContext) *TestPanicViewEPCall {
	return &TestPanicViewEPCall{Func: wasmlib.NewScView(HScName, HViewTestPanicViewEP)}
}

func NewTestPanicViewEPCallFromView(ctx wasmlib.ScViewContext) *TestPanicViewEPCall {
	return NewTestPanicViewEPCall(wasmlib.ScFuncContext{})
}

type TestSandboxCallCall struct {
	Func    *wasmlib.ScView
	Results ImmutableTestSandboxCallResults
}

func NewTestSandboxCallCall(ctx wasmlib.ScFuncContext) *TestSandboxCallCall {
	f := &TestSandboxCallCall{Func: wasmlib.NewScView(HScName, HViewTestSandboxCall)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func NewTestSandboxCallCallFromView(ctx wasmlib.ScViewContext) *TestSandboxCallCall {
	return NewTestSandboxCallCall(wasmlib.ScFuncContext{})
}
