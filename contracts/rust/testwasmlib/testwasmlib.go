// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testwasmlib

import (
	"bytes"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcParamTypes(ctx wasmlib.ScFuncContext, f *FuncParamTypesContext) {
	if f.Params.Address.Exists() {
		ctx.Require(f.Params.Address.Value() == ctx.AccountId().Address(), "mismatch: Address")
	}
	if f.Params.AgentId.Exists() {
		ctx.Require(f.Params.AgentId.Value() == ctx.AccountId(), "mismatch: AgentId")
	}
	if f.Params.Bytes.Exists() {
		byteData := []byte("these are bytes")
		ctx.Require(bytes.Equal(f.Params.Bytes.Value(), byteData), "mismatch: Bytes")
	}
	if f.Params.ChainId.Exists() {
		ctx.Require(f.Params.ChainId.Value() == ctx.ChainId(), "mismatch: ChainId")
	}
	if f.Params.Color.Exists() {
		color := wasmlib.NewScColorFromBytes([]byte("RedGreenBlueYellowCyanBlackWhite"))
		ctx.Require(f.Params.Color.Value() == color, "mismatch: Color")
	}
	if f.Params.Hash.Exists() {
		hash := wasmlib.NewScHashFromBytes([]byte("0123456789abcdeffedcba9876543210"))
		ctx.Require(f.Params.Hash.Value() == hash, "mismatch: Hash")
	}
	if f.Params.Hname.Exists() {
		ctx.Require(f.Params.Hname.Value() == ctx.AccountId().Hname(), "mismatch: Hname")
	}
	if f.Params.Int64.Exists() {
		ctx.Require(f.Params.Int64.Value() == 1234567890123456789, "mismatch: Int64")
	}
	if f.Params.RequestId.Exists() {
		requestId := wasmlib.NewScRequestIdFromBytes([]byte("abcdefghijklmnopqrstuvwxyz123456\x00\x00"))
		ctx.Require(f.Params.RequestId.Value() == requestId, "mismatch: RequestId")
	}
	if f.Params.String.Exists() {
		ctx.Require(f.Params.String.Value() == "this is a string", "mismatch: String")
	}
}
