// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_param_types(ctx: &ScFuncContext, f: &FuncParamTypesContext) {
    if f.params.address.exists() {
        ctx.require(f.params.address.value() == ctx.account_id().address(), "mismatch: Address");
    }
    if f.params.agent_id.exists() {
        ctx.require(f.params.agent_id.value() == ctx.account_id(), "mismatch: AgentId");
    }
    if f.params.bytes.exists() {
        let byte_data = "these are bytes".as_bytes();
        ctx.require(f.params.bytes.value() == byte_data, "mismatch: Bytes");
    }
    if f.params.chain_id.exists() {
        ctx.require(f.params.chain_id.value() == ctx.chain_id(), "mismatch: ChainId");
    }
    if f.params.color.exists() {
        let color = ScColor::from_bytes("RedGreenBlueYellowCyanBlackWhite".as_bytes());
        ctx.require(f.params.color.value() == color, "mismatch: Color");
    }
    if f.params.hash.exists() {
        let hash = ScHash::from_bytes("0123456789abcdeffedcba9876543210".as_bytes());
        ctx.require(f.params.hash.value() == hash, "mismatch: Hash");
    }
    if f.params.hname.exists() {
        ctx.require(f.params.hname.value() == ctx.account_id().hname(), "mismatch: Hname");
    }
    if f.params.int64.exists() {
        ctx.require(f.params.int64.value() == 1234567890123456789, "mismatch: Int64");
    }
    if f.params.request_id.exists() {
        let request_id = ScRequestId::from_bytes("abcdefghijklmnopqrstuvwxyz123456\x00\x00".as_bytes());
        ctx.require(f.params.request_id.value() == request_id, "mismatch: RequestId");
    }
    if f.params.string.exists() {
        ctx.require(f.params.string.value() == "this is a string", "mismatch: String");
    }
}
