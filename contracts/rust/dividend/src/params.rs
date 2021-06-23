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
pub struct ImmutableInitParams {
    pub(crate) id: i32,
}

impl ImmutableInitParams {
    pub fn owner(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_OWNER))
    }
}

#[derive(Clone, Copy)]
pub struct MutableInitParams {
    pub(crate) id: i32,
}

impl MutableInitParams {
    pub fn owner(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_OWNER))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableMemberParams {
    pub(crate) id: i32,
}

impl ImmutableMemberParams {
    pub fn address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }

    pub fn factor(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_FACTOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableMemberParams {
    pub(crate) id: i32,
}

impl MutableMemberParams {
    pub fn address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }

    pub fn factor(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_FACTOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSetOwnerParams {
    pub(crate) id: i32,
}

impl ImmutableSetOwnerParams {
    pub fn owner(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_OWNER))
    }
}

#[derive(Clone, Copy)]
pub struct MutableSetOwnerParams {
    pub(crate) id: i32,
}

impl MutableSetOwnerParams {
    pub fn owner(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_OWNER))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetFactorParams {
    pub(crate) id: i32,
}

impl ImmutableGetFactorParams {
    pub fn address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetFactorParams {
    pub(crate) id: i32,
}

impl MutableGetFactorParams {
    pub fn address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }
}
