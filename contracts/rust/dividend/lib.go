// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncDivide, funcDivideThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncMember, funcMemberThunk)
	exports.AddFunc(FuncSetOwner, funcSetOwnerThunk)
	exports.AddView(ViewGetFactor, viewGetFactorThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyId()
	}
}

type FuncDivideContext struct {
	State DividendFuncState
}

func funcDivideThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcDivide")
	f := &FuncDivideContext{
		State: DividendFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcDivide(ctx, f)
	ctx.Log("dividend.funcDivide ok")
}

type FuncInitParams struct {
	Owner wasmlib.ScImmutableAgentId // optional owner, defaults to contract creator
}

type FuncInitContext struct {
	Params FuncInitParams
	State  DividendFuncState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcInit")
	p := ctx.Params().MapId()
	f := &FuncInitContext{
		Params: FuncInitParams{
			Owner: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamOwner]),
		},
		State: DividendFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcInit(ctx, f)
	ctx.Log("dividend.funcInit ok")
}

type FuncMemberParams struct {
	Address wasmlib.ScImmutableAddress // address of dividend recipient
	Factor  wasmlib.ScImmutableInt64   // relative division factor
}

type FuncMemberContext struct {
	Params FuncMemberParams
	State  DividendFuncState
}

func funcMemberThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcMember")
	// only defined owner can add members
	access := ctx.State().GetAgentId(wasmlib.Key("owner"))
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncMemberContext{
		Params: FuncMemberParams{
			Address: wasmlib.NewScImmutableAddress(p, idxMap[IdxParamAddress]),
			Factor:  wasmlib.NewScImmutableInt64(p, idxMap[IdxParamFactor]),
		},
		State: DividendFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Address.Exists(), "missing mandatory address")
	ctx.Require(f.Params.Factor.Exists(), "missing mandatory factor")
	funcMember(ctx, f)
	ctx.Log("dividend.funcMember ok")
}

type FuncSetOwnerParams struct {
	Owner wasmlib.ScImmutableAgentId // new owner of smart contract
}

type FuncSetOwnerContext struct {
	Params FuncSetOwnerParams
	State  DividendFuncState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcSetOwner")
	// only defined owner can change owner
	access := ctx.State().GetAgentId(wasmlib.Key("owner"))
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncSetOwnerContext{
		Params: FuncSetOwnerParams{
			Owner: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamOwner]),
		},
		State: DividendFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Owner.Exists(), "missing mandatory owner")
	funcSetOwner(ctx, f)
	ctx.Log("dividend.funcSetOwner ok")
}

type ViewGetFactorParams struct {
	Address wasmlib.ScImmutableAddress // address of dividend recipient
}

type ViewGetFactorResults struct {
	Factor wasmlib.ScMutableInt64 // relative division factor
}

type ViewGetFactorContext struct {
	Params  ViewGetFactorParams
	Results ViewGetFactorResults
	State   DividendViewState
}

func viewGetFactorThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("dividend.viewGetFactor")
	p := ctx.Params().MapId()
	r := ctx.Results().MapId()
	f := &ViewGetFactorContext{
		Params: ViewGetFactorParams{
			Address: wasmlib.NewScImmutableAddress(p, idxMap[IdxParamAddress]),
		},
		Results: ViewGetFactorResults{
			Factor: wasmlib.NewScMutableInt64(r, idxMap[IdxResultFactor]),
		},
		State: DividendViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Address.Exists(), "missing mandatory address")
	viewGetFactor(ctx, f)
	ctx.Log("dividend.viewGetFactor ok")
}
