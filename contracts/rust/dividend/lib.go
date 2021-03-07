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
	exports.AddFunc(FuncMember, funcMemberThunk)
	exports.AddView(ViewGetFactor, viewGetFactorThunk)
}

type FuncDivideParams struct {
}

func funcDivideThunk(ctx wasmlib.ScFuncContext) {
	params := &FuncDivideParams{
	}
	funcDivide(ctx, params)
}

type FuncMemberParams struct {
	Address wasmlib.ScImmutableAddress // address of dividend recipient
	Factor  wasmlib.ScImmutableInt64   // relative division factor
}

func funcMemberThunk(ctx wasmlib.ScFuncContext) {
	// only creator can add members
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params()
	params := &FuncMemberParams{
		Address: p.GetAddress(ParamAddress),
		Factor:  p.GetInt64(ParamFactor),
	}
	ctx.Require(params.Address.Exists(), "missing mandatory address")
	ctx.Require(params.Factor.Exists(), "missing mandatory factor")
	funcMember(ctx, params)
}

type ViewGetFactorParams struct {
	Address wasmlib.ScImmutableAddress // address of dividend recipient
}

func viewGetFactorThunk(ctx wasmlib.ScViewContext) {
	p := ctx.Params()
	params := &ViewGetFactorParams{
		Address: p.GetAddress(ParamAddress),
	}
	ctx.Require(params.Address.Exists(), "missing mandatory address")
	viewGetFactor(ctx, params)
}
