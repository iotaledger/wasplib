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
use crate::keys::*;

#[derive(Clone, Copy)]
pub struct ImmutableFuncMintSupplyParams {
    pub(crate) id: i32,
}

impl ImmutableFuncMintSupplyParams {
    pub fn description(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
    }

    pub fn user_defined(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_USER_DEFINED))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFuncMintSupplyParams {
    pub(crate) id: i32,
}

impl MutableFuncMintSupplyParams {
    pub fn new() -> MutableFuncMintSupplyParams {
        MutableFuncMintSupplyParams { id: ScMutableMap::new().map_id() }
    }

    pub fn description(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
    }

    pub fn user_defined(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_USER_DEFINED))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableFuncTransferOwnershipParams {
    pub(crate) id: i32,
}

impl ImmutableFuncTransferOwnershipParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFuncTransferOwnershipParams {
    pub(crate) id: i32,
}

impl MutableFuncTransferOwnershipParams {
    pub fn new() -> MutableFuncTransferOwnershipParams {
        MutableFuncTransferOwnershipParams { id: ScMutableMap::new().map_id() }
    }

    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableFuncUpdateMetadataParams {
    pub(crate) id: i32,
}

impl ImmutableFuncUpdateMetadataParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFuncUpdateMetadataParams {
    pub(crate) id: i32,
}

impl MutableFuncUpdateMetadataParams {
    pub fn new() -> MutableFuncUpdateMetadataParams {
        MutableFuncUpdateMetadataParams { id: ScMutableMap::new().map_id() }
    }

    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableViewGetInfoParams {
    pub(crate) id: i32,
}

impl ImmutableViewGetInfoParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableViewGetInfoParams {
    pub(crate) id: i32,
}

impl MutableViewGetInfoParams {
    pub fn new() -> MutableViewGetInfoParams {
        MutableViewGetInfoParams { id: ScMutableMap::new().map_id() }
    }

    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}
