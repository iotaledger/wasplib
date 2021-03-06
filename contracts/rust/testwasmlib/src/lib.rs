// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use testwasmlib::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod params;
mod results;
mod state;
mod testwasmlib;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_PARAM_TYPES, func_param_types_thunk);
    exports.add_view(VIEW_BLOCK_RECORD, view_block_record_thunk);
    exports.add_view(VIEW_BLOCK_RECORDS, view_block_records_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct ParamTypesContext {
    params: ImmutableParamTypesParams,
    state:  MutableTestWasmLibState,
}

fn func_param_types_thunk(ctx: &ScFuncContext) {
    ctx.log("testwasmlib.funcParamTypes");
    let f = ParamTypesContext {
        params: ImmutableParamTypesParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestWasmLibState {
            id: OBJ_ID_STATE,
        },
    };
    func_param_types(ctx, &f);
    ctx.log("testwasmlib.funcParamTypes ok");
}

pub struct BlockRecordContext {
    params:  ImmutableBlockRecordParams,
    results: MutableBlockRecordResults,
    state:   ImmutableTestWasmLibState,
}

fn view_block_record_thunk(ctx: &ScViewContext) {
    ctx.log("testwasmlib.viewBlockRecord");
    let f = BlockRecordContext {
        params: ImmutableBlockRecordParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableBlockRecordResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestWasmLibState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.block_index().exists(), "missing mandatory blockIndex");
    ctx.require(f.params.record_index().exists(), "missing mandatory recordIndex");
    view_block_record(ctx, &f);
    ctx.log("testwasmlib.viewBlockRecord ok");
}

pub struct BlockRecordsContext {
    params:  ImmutableBlockRecordsParams,
    results: MutableBlockRecordsResults,
    state:   ImmutableTestWasmLibState,
}

fn view_block_records_thunk(ctx: &ScViewContext) {
    ctx.log("testwasmlib.viewBlockRecords");
    let f = BlockRecordsContext {
        params: ImmutableBlockRecordsParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableBlockRecordsResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestWasmLibState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.block_index().exists(), "missing mandatory blockIndex");
    view_block_records(ctx, &f);
    ctx.log("testwasmlib.viewBlockRecords ok");
}

// @formatter:on
