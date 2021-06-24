// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]

use std::ptr;
use crate::*;
use crate::corecontracts::coreeventlog::*;

pub struct GetNumRecordsCall {
    pub func: ScView,
    pub params: MutableGetNumRecordsParams,
    pub results: ImmutableGetNumRecordsResults,
}

impl GetNumRecordsCall {
    pub fn new(_ctx: &ScFuncContext) -> GetNumRecordsCall {
        let mut f = GetNumRecordsCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_NUM_RECORDS),
            params: MutableGetNumRecordsParams { id: 0 },
            results: ImmutableGetNumRecordsResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetNumRecordsCall {
        GetNumRecordsCall::new(&ScFuncContext{})
    }
}

pub struct GetRecordsCall {
    pub func: ScView,
    pub params: MutableGetRecordsParams,
}

impl GetRecordsCall {
    pub fn new(_ctx: &ScFuncContext) -> GetRecordsCall {
        let mut f = GetRecordsCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_RECORDS),
            params: MutableGetRecordsParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetRecordsCall {
        GetRecordsCall::new(&ScFuncContext{})
    }
}
