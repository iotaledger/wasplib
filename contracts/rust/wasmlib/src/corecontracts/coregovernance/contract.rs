// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]

use std::ptr;

use crate::*;
use crate::corecontracts::coregovernance::*;

pub struct AddAllowedStateControllerAddressCall {
    pub func: ScFunc,
    pub params: MutableAddAllowedStateControllerAddressParams,
}

impl AddAllowedStateControllerAddressCall {
    pub fn new(_ctx: &ScFuncContext) -> AddAllowedStateControllerAddressCall {
        let mut f = AddAllowedStateControllerAddressCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params: MutableAddAllowedStateControllerAddressParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct RemoveAllowedStateControllerAddressCall {
    pub func: ScFunc,
    pub params: MutableRemoveAllowedStateControllerAddressParams,
}

impl RemoveAllowedStateControllerAddressCall {
    pub fn new(_ctx: &ScFuncContext) -> RemoveAllowedStateControllerAddressCall {
        let mut f = RemoveAllowedStateControllerAddressCall {
            func: ScFunc::new(HSC_NAME, HFUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params: MutableRemoveAllowedStateControllerAddressParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct RotateStateControllerCall {
    pub func: ScFunc,
    pub params: MutableRotateStateControllerParams,
}

impl RotateStateControllerCall {
    pub fn new(_ctx: &ScFuncContext) -> RotateStateControllerCall {
        let mut f = RotateStateControllerCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ROTATE_STATE_CONTROLLER),
            params: MutableRotateStateControllerParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct GetAllowedStateControllerAddressesCall {
    pub func: ScView,
    pub results: ImmutableGetAllowedStateControllerAddressesResults,
}

impl GetAllowedStateControllerAddressesCall {
    pub fn new(_ctx: &ScFuncContext) -> GetAllowedStateControllerAddressesCall {
        let mut f = GetAllowedStateControllerAddressesCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES),
            results: ImmutableGetAllowedStateControllerAddressesResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetAllowedStateControllerAddressesCall {
        GetAllowedStateControllerAddressesCall::new(&ScFuncContext {})
    }
}
