// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct ParamTypesCall {
    pub func:   ScFunc,
    pub params: MutableParamTypesParams,
}

impl ParamTypesCall {
    pub fn new(_ctx: &ScFuncContext) -> ParamTypesCall {
        let mut f = ParamTypesCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_PARAM_TYPES),
            params: MutableParamTypesParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct BlockRecordCall {
    pub func:    ScView,
    pub params:  MutableBlockRecordParams,
    pub results: ImmutableBlockRecordResults,
}

impl BlockRecordCall {
    pub fn new(_ctx: &ScFuncContext) -> BlockRecordCall {
        let mut f = BlockRecordCall {
            func:    ScView::new(HSC_NAME, HVIEW_BLOCK_RECORD),
            params:  MutableBlockRecordParams { id: 0 },
            results: ImmutableBlockRecordResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> BlockRecordCall {
        BlockRecordCall::new(&ScFuncContext {})
    }
}

pub struct BlockRecordsCall {
    pub func:    ScView,
    pub params:  MutableBlockRecordsParams,
    pub results: ImmutableBlockRecordsResults,
}

impl BlockRecordsCall {
    pub fn new(_ctx: &ScFuncContext) -> BlockRecordsCall {
        let mut f = BlockRecordsCall {
            func:    ScView::new(HSC_NAME, HVIEW_BLOCK_RECORDS),
            params:  MutableBlockRecordsParams { id: 0 },
            results: ImmutableBlockRecordsResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> BlockRecordsCall {
        BlockRecordsCall::new(&ScFuncContext {})
    }
}

//@formatter:on
