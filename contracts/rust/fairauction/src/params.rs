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
pub struct ImmutableFinalizeAuctionParams {
    pub(crate) id: i32,
}

impl ImmutableFinalizeAuctionParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFinalizeAuctionParams {
    pub(crate) id: i32,
}

impl MutableFinalizeAuctionParams {
    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutablePlaceBidParams {
    pub(crate) id: i32,
}

impl ImmutablePlaceBidParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutablePlaceBidParams {
    pub(crate) id: i32,
}

impl MutablePlaceBidParams {
    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSetOwnerMarginParams {
    pub(crate) id: i32,
}

impl ImmutableSetOwnerMarginParams {
    pub fn owner_margin(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_OWNER_MARGIN))
    }
}

#[derive(Clone, Copy)]
pub struct MutableSetOwnerMarginParams {
    pub(crate) id: i32,
}

impl MutableSetOwnerMarginParams {
    pub fn owner_margin(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_OWNER_MARGIN))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableStartAuctionParams {
    pub(crate) id: i32,
}

impl ImmutableStartAuctionParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }

    pub fn description(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
    }

    pub fn duration(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, idx_map(IDX_PARAM_DURATION))
    }

    pub fn minimum_bid(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_MINIMUM_BID))
    }
}

#[derive(Clone, Copy)]
pub struct MutableStartAuctionParams {
    pub(crate) id: i32,
}

impl MutableStartAuctionParams {
    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }

    pub fn description(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
    }

    pub fn duration(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, idx_map(IDX_PARAM_DURATION))
    }

    pub fn minimum_bid(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_MINIMUM_BID))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetInfoParams {
    pub(crate) id: i32,
}

impl ImmutableGetInfoParams {
    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetInfoParams {
    pub(crate) id: i32,
}

impl MutableGetInfoParams {
    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
    }
}
