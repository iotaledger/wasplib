// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testwasmlib

import (
	"bytes"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcParamTypes(ctx wasmlib.ScFuncContext, params *FuncParamTypesParams) {
	if params.Address.Exists() {
		ctx.Require(params.Address.Value() == ctx.AccountId().Address(), "mismatch: Address")
	}
	if params.AgentId.Exists() {
		ctx.Require(params.AgentId.Value() == ctx.AccountId(), "mismatch: AgentId")
	}
	if params.Bytes.Exists() {
		bytez := []byte("these are bytes")
		ctx.Require(bytes.Equal(params.Bytes.Value(), bytez), "mismatch: Bytes")
	}
	if params.ChainId.Exists() {
		ctx.Require(params.ChainId.Value() == ctx.ChainId(), "mismatch: ChainId")
	}
	if params.Color.Exists() {
		color := wasmlib.NewScColorFromBytes([]byte("RedGreenBlueYellowCyanBlackWhite"))
		ctx.Require(params.Color.Value() == color, "mismatch: Color")
	}
	if params.Hash.Exists() {
		hash :=  wasmlib.NewScHashFromBytes([]byte("0123456789abcdeffedcba9876543210"))
		ctx.Require(params.Hash.Value() == hash, "mismatch: Hash")
	}
	if params.Hname.Exists() {
		ctx.Require(params.Hname.Value() == ctx.AccountId().Hname(), "mismatch: Hname")
	}
	if params.Int64.Exists() {
		ctx.Require(params.Int64.Value() == 1234567890123456789, "mismatch: Int64")
	}
	if params.RequestId.Exists() {
		requestId :=  wasmlib.NewScRequestIdFromBytes([]byte("abcdefghijklmnopqrstuvwxyz12345678"))
		ctx.Require(params.RequestId.Value() == requestId, "mismatch: RequestId")
	}
	if params.String.Exists() {
		ctx.Require(params.String.Value() == "this is a string", "mismatch: String")
	}
}
