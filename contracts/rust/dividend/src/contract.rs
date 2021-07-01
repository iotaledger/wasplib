// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct DivideCall {
    pub func: ScFunc,
}

impl DivideCall {
    pub fn new(_ctx: &ScFuncContext) -> DivideCall {
        DivideCall {
            func: ScFunc::new(HSC_NAME, HFUNC_DIVIDE),
        }
    }
}

pub struct InitCall {
    pub func:   ScFunc,
    pub params: MutableInitParams,
}

impl InitCall {
    pub fn new(_ctx: &ScFuncContext) -> InitCall {
        let mut f = InitCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_INIT),
            params: MutableInitParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct MemberCall {
    pub func:   ScFunc,
    pub params: MutableMemberParams,
}

impl MemberCall {
    pub fn new(_ctx: &ScFuncContext) -> MemberCall {
        let mut f = MemberCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_MEMBER),
            params: MutableMemberParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct SetOwnerCall {
    pub func:   ScFunc,
    pub params: MutableSetOwnerParams,
}

impl SetOwnerCall {
    pub fn new(_ctx: &ScFuncContext) -> SetOwnerCall {
        let mut f = SetOwnerCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_OWNER),
            params: MutableSetOwnerParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct GetFactorCall {
    pub func:    ScView,
    pub params:  MutableGetFactorParams,
    pub results: ImmutableGetFactorResults,
}

impl GetFactorCall {
    pub fn new(_ctx: &ScFuncContext) -> GetFactorCall {
        let mut f = GetFactorCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_FACTOR),
            params:  MutableGetFactorParams { id: 0 },
            results: ImmutableGetFactorResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetFactorCall {
        GetFactorCall::new(&ScFuncContext {})
    }
}

//@formatter:on
