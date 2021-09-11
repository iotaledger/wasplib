// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct InitCall {
    pub func:   ScFunc,
    pub params: MutableInitParams,
}

pub struct SetOwnerCall {
    pub func:   ScFunc,
    pub params: MutableSetOwnerParams,
}

pub struct GetOwnerCall {
    pub func:    ScView,
    pub results: ImmutableGetOwnerResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn init(_ctx: & dyn ScFuncCallContext) -> InitCall {
        let mut f = InitCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_INIT),
            params: MutableInitParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn set_owner(_ctx: & dyn ScFuncCallContext) -> SetOwnerCall {
        let mut f = SetOwnerCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_OWNER),
            params: MutableSetOwnerParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn get_owner(_ctx: & dyn ScViewCallContext) -> GetOwnerCall {
        let mut f = GetOwnerCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_OWNER),
            results: ImmutableGetOwnerResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
}

// @formatter:on
