// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coregovernance::*;
use crate::host::*;

pub struct ArrayOfImmutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableBytes {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetAllowedStateControllerAddressesResults {
    pub(crate) id: i32,
}

impl ImmutableGetAllowedStateControllerAddressesResults {
    pub fn allowed_state_controller_addresses(&self) -> ArrayOfImmutableBytes {
        let arr_id = get_object_id(self.id, RESULT_ALLOWED_STATE_CONTROLLER_ADDRESSES.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBytes { obj_id: arr_id }
    }
}

pub struct ArrayOfMutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableBytes {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetAllowedStateControllerAddressesResults {
    pub(crate) id: i32,
}

impl MutableGetAllowedStateControllerAddressesResults {
    pub fn allowed_state_controller_addresses(&self) -> ArrayOfMutableBytes {
        let arr_id = get_object_id(self.id, RESULT_ALLOWED_STATE_CONTROLLER_ADDRESSES.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBytes { obj_id: arr_id }
    }
}
