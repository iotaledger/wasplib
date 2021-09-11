// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_param_types(ctx: &ScFuncContext, f: &ParamTypesContext) {
    if f.params.address().exists() {
        ctx.require(f.params.address().value() == ctx.account_id().address(), "mismatch: Address");
    }
    if f.params.agent_id().exists() {
        ctx.require(f.params.agent_id().value() == ctx.account_id(), "mismatch: AgentID");
    }
    if f.params.bytes().exists() {
        let byte_data = "these are bytes".as_bytes();
        ctx.require(f.params.bytes().value() == byte_data, "mismatch: Bytes");
    }
    if f.params.chain_id().exists() {
        ctx.require(f.params.chain_id().value() == ctx.chain_id(), "mismatch: ChainID");
    }
    if f.params.color().exists() {
        let color = ScColor::from_bytes("RedGreenBlueYellowCyanBlackWhite".as_bytes());
        ctx.require(f.params.color().value() == color, "mismatch: Color");
    }
    if f.params.hash().exists() {
        let hash = ScHash::from_bytes("0123456789abcdeffedcba9876543210".as_bytes());
        ctx.require(f.params.hash().value() == hash, "mismatch: Hash");
    }
    if f.params.hname().exists() {
        ctx.require(f.params.hname().value() == ctx.account_id().hname(), "mismatch: Hname");
    }
    if f.params.int16().exists() {
        ctx.require(f.params.int16().value() == 12345, "mismatch: Int16");
    }
    if f.params.int32().exists() {
        ctx.require(f.params.int32().value() == 1234567890, "mismatch: Int32");
    }
    if f.params.int64().exists() {
        ctx.require(f.params.int64().value() == 1234567890123456789, "mismatch: Int64");
    }
    if f.params.request_id().exists() {
        let request_id = ScRequestID::from_bytes("abcdefghijklmnopqrstuvwxyz123456\x00\x00".as_bytes());
        ctx.require(f.params.request_id().value() == request_id, "mismatch: RequestID");
    }
    if f.params.string().exists() {
        ctx.require(f.params.string().value() == "this is a string", "mismatch: String");
    }
}

pub fn view_block_record(ctx: &ScViewContext, f: &BlockRecordContext) {
    let records = corecontracts::coreblocklog::ScFuncs::get_request_receipts_for_block(ctx);
    records.params.block_index().set_value(f.params.block_index().value());
    records.func.call();
    let record_index = f.params.record_index().value();
    ctx.require(record_index < records.results.request_record().length(), "invalid recordIndex");
    f.results.record().set_value(&records.results.request_record().get_bytes(record_index).value());
}

pub fn view_block_records(ctx: &ScViewContext, f: &BlockRecordsContext) {
    let records = corecontracts::coreblocklog::ScFuncs::get_request_receipts_for_block(ctx);
    records.params.block_index().set_value(f.params.block_index().value());
    records.func.call();
    f.results.count().set_value(records.results.request_record().length());
}

pub fn func_array_clear(_ctx: &ScFuncContext, f: &ArrayClearContext) {
    let name = f.params.name().value();
    let array = f.state.arrays().get_string_array(&name);
    array.clear();
}

pub fn func_array_create(_ctx: &ScFuncContext, f: &ArrayCreateContext) {
    let name = f.params.name().value();
    let array = f.state.arrays().get_string_array(&name);
    array.clear();
}

pub fn func_array_set(_ctx: &ScFuncContext, f: &ArraySetContext) {
    let name = f.params.name().value();
    let array = f.state.arrays().get_string_array(&name);
    let index = f.params.index().value();
    let value = f.params.value().value();
    array.get_string(index).set_value(&value);
}

pub fn view_array_length(_ctx: &ScViewContext, f: &ArrayLengthContext) {
    let name = f.params.name().value();
    let array = f.state.arrays().get_string_array(&name);
    let length = array.length();
    f.results.length().set_value(length);
}

pub fn view_array_value(_ctx: &ScViewContext, f: &ArrayValueContext) {
    let name = f.params.name().value();
    let array = f.state.arrays().get_string_array(&name);
    let index = f.params.index().value();
    let value = array.get_string(index).value();
    f.results.value().set_value(&value);
}

pub fn view_iota_balance(ctx: &ScViewContext, f: &IotaBalanceContext) {
    f.results.iotas().set_value(ctx.balances().balance(&ScColor::IOTA));
}
