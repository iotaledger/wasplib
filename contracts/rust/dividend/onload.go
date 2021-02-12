// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncDivide, funcDivideThunk)
	exports.AddFunc(FuncMember, funcMemberThunk)
}

type FuncDivideParams struct {
}

func funcDivideThunk(ctx *wasmlib.ScFuncContext) {
	params := &FuncDivideParams{
	}
	funcDivide(ctx, params)
}

type FuncMemberParams struct {
	Address wasmlib.ScImmutableAddress // address of dividend recipient
	Factor  wasmlib.ScImmutableInt     // relative division factor
}

func funcMemberThunk(ctx *wasmlib.ScFuncContext) {
	// only creator can add members
	ctx.Require(ctx.From(ctx.ContractCreator()), "no permission")

	p := ctx.Params()
	params := &FuncMemberParams{
		Address: p.GetAddress(ParamAddress),
		Factor:  p.GetInt(ParamFactor),
	}
	ctx.Require(params.Address.Exists(), "missing mandatory address")
	ctx.Require(params.Factor.Exists(), "missing mandatory factor")
	funcMember(ctx, params)
}
