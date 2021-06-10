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

pub struct MutableFuncLockBetsParams {
    pub(crate) id: i32,
}

pub struct ImmutableFuncLockBetsParams {
    pub(crate) id: i32,
}

pub struct MutableFuncPayWinnersParams {
    pub(crate) id: i32,
}

pub struct ImmutableFuncPayWinnersParams {
    pub(crate) id: i32,
}

pub struct MutableFuncPlaceBetParams {
    pub(crate) id: i32,
}

impl MutableFuncPlaceBetParams {
    pub fn number(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_NUMBER))
    }
}

pub struct ImmutableFuncPlaceBetParams {
    pub(crate) id: i32,
}

impl ImmutableFuncPlaceBetParams {
    pub fn number(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_NUMBER))
    }
}

pub struct MutableFuncPlayPeriodParams {
    pub(crate) id: i32,
}

impl MutableFuncPlayPeriodParams {
    pub fn play_period(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_PLAY_PERIOD))
    }
}

pub struct ImmutableFuncPlayPeriodParams {
    pub(crate) id: i32,
}

impl ImmutableFuncPlayPeriodParams {
    pub fn play_period(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_PLAY_PERIOD))
    }
}

pub struct MutableViewLastWinningNumberParams {
    pub(crate) id: i32,
}

pub struct ImmutableViewLastWinningNumberParams {
    pub(crate) id: i32,
}
