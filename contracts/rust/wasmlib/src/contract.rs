// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// base contract objects

use std::ptr;

use crate::bytes::*;
use crate::context::*;
use crate::hashtypes::*;
use crate::host::*;
use crate::keys::*;
use crate::mutable::*;

pub trait ScFuncCallContext {
    fn can_call_func(&self);
}

pub trait ScViewCallContext {
    fn can_call_view(&self);
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Clone, Copy)]
pub struct ScView {
    h_contract: ScHname,
    h_function: ScHname,
    params_id: *mut i32,
    results_id: *mut i32,
}

impl ScView {
    pub fn new(h_contract: ScHname, h_function: ScHname) -> ScView {
        ScView {
            h_contract: h_contract,
            h_function: h_function,
            params_id: ptr::null_mut(),
            results_id: ptr::null_mut(),
        }
    }

    pub fn set_ptrs(&mut self, params_id: *mut i32, results_id: *mut i32) {
        self.params_id = params_id;
        self.results_id = results_id;

        unsafe {
            if params_id != ptr::null_mut() {
                *params_id = ScMutableMap::new().map_id();
            }
        }
    }

    pub fn call(&self) {
        self.call_with_transfer(0);
    }

    fn call_with_transfer(&self, transfer_id: i32) {
        let mut encode = BytesEncoder::new();
        encode.hname(&self.h_contract);
        encode.hname(&self.h_function);
        encode.int32(self.id(self.params_id));
        encode.int32(transfer_id);
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());

        unsafe {
            if self.results_id != ptr::null_mut() {
                *self.results_id = get_object_id(1, KEY_RETURN, TYPE_MAP);
            }
        }
    }

    pub fn of_contract(&self, h_contract: ScHname) -> ScView {
        let mut ret = self.clone();
        ret.h_contract = h_contract;
        ret
    }

    fn id(&self, params_id: *mut i32) -> i32 {
        unsafe {
            if params_id == ptr::null_mut() {
                return 0;
            }
            *params_id
        }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Clone, Copy)]
pub struct ScFunc {
    view: ScView,
    delay: i32,
    transfer_id: i32,
}

impl ScFunc {
    pub fn new(h_contract: ScHname, h_function: ScHname) -> ScFunc {
        ScFunc {
            view: ScView::new(h_contract, h_function),
            delay: 0,
            transfer_id: 0,
        }
    }

    pub fn set_ptrs(&mut self, params_id: *mut i32, results_id: *mut i32) {
        self.view.set_ptrs(params_id, results_id);
    }

    pub fn call(&self) {
        self.view.call_with_transfer(self.transfer_id);
    }

    pub fn delay(&self, seconds: i32) -> ScFunc {
        let mut ret = self.clone();
        ret.delay = seconds;
        ret
    }

    pub fn of_contract(&self, h_contract: ScHname) -> ScFunc {
        let mut ret = self.clone();
        ret.view.h_contract = h_contract;
        ret
    }

    pub fn post(&self) {
        self.post_to_chain(ROOT.get_chain_id(&KEY_CHAIN_ID).value())
    }

    pub fn post_to_chain(&self, chain_id: ScChainID) {
        if self.transfer_id == 0 {
            panic("transfer is required for post")
        }

        let mut encode = BytesEncoder::new();
        encode.chain_id(&chain_id);
        encode.hname(&self.view.h_contract);
        encode.hname(&self.view.h_function);
        encode.int32(self.view.id(self.view.params_id));
        encode.int32(self.transfer_id);
        encode.int32(self.delay);
        ROOT.get_bytes(&KEY_POST).set_value(&encode.data());
    }

    pub fn transfer(&self, transfer: ScTransfers) -> ScFunc {
        let mut ret = self.clone();
        ret.transfer_id = transfer.transfers.obj_id;
        ret
    }

    pub fn transfer_iotas(&self, amount: i64) -> ScFunc {
        self.transfer(ScTransfers::iotas(amount))
    }
}
