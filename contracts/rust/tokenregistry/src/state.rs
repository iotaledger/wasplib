// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::types::*;

pub struct ArrayOfMutableColor {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableColor {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_color(&self, index: i32) -> ScMutableColor {
        ScMutableColor::new(self.obj_id, Key32(index))
    }
}

pub struct MapColorToMutableToken {
    pub(crate) obj_id: i32,
}

impl MapColorToMutableToken {
    pub fn clear(&self) {
        clear(self.obj_id)
    }

    pub fn get_token(&self, key: &ScColor) -> MutableToken {
        MutableToken { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

pub struct TokenRegistryFuncState {
    pub(crate) state_id: i32,
}

impl TokenRegistryFuncState {
    pub fn color_list(&self) -> ArrayOfMutableColor {
        let arr_id = get_object_id(self.state_id, VAR_COLOR_LIST.get_key_id(), TYPE_ARRAY | TYPE_COLOR);
        ArrayOfMutableColor { obj_id: arr_id }
    }

    pub fn registry(&self) -> MapColorToMutableToken {
        let map_id = get_object_id(self.state_id, VAR_REGISTRY.get_key_id(), TYPE_MAP);
        MapColorToMutableToken { obj_id: map_id }
    }
}

pub struct ArrayOfImmutableColor {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableColor {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_color(&self, index: i32) -> ScImmutableColor {
        ScImmutableColor::new(self.obj_id, Key32(index))
    }
}

pub struct MapColorToImmutableToken {
    pub(crate) obj_id: i32,
}

impl MapColorToImmutableToken {
    pub fn get_token(&self, key: &ScColor) -> ImmutableToken {
        ImmutableToken { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

pub struct TokenRegistryViewState {
    pub(crate) state_id: i32,
}

impl TokenRegistryViewState {
    pub fn color_list(&self) -> ArrayOfImmutableColor {
        let arr_id = get_object_id(self.state_id, VAR_COLOR_LIST.get_key_id(), TYPE_ARRAY | TYPE_COLOR);
        ArrayOfImmutableColor { obj_id: arr_id }
    }

    pub fn registry(&self) -> MapColorToImmutableToken {
        let map_id = get_object_id(self.state_id, VAR_REGISTRY.get_key_id(), TYPE_MAP);
        MapColorToImmutableToken { obj_id: map_id }
    }
}
