// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddCall(FuncDivide, funcDivideThunk)
	exports.AddCall(FuncMember, funcMemberThunk)
}

type FuncDivideParams struct {
}

func funcDivideThunk(ctx *wasmlib.ScCallContext) {
	params := &FuncDivideParams{
	}
	funcDivide(ctx, params)
}

type FuncMemberParams struct {
	Address wasmlib.ScImmutableAddress
	Factor  wasmlib.ScImmutableInt
}

func funcMemberThunk(ctx *wasmlib.ScCallContext) {
	p := ctx.Params()
	params := &FuncMemberParams{
		Address: p.GetAddress(ParamAddress),
		Factor:  p.GetInt(ParamFactor),
	}
	ctx.Require(params.Address.Exists(), "missing mandatory address")
	ctx.Require(params.Factor.Exists(), "missing mandatory factor")
	funcMember(ctx, params)
}
