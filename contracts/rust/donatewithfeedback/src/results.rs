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

pub struct MutableFuncDonateResults {
    pub(crate) id: i32,
}

pub struct ImmutableFuncDonateResults {
    pub(crate) id: i32,
}

pub struct MutableFuncWithdrawResults {
    pub(crate) id: i32,
}

pub struct ImmutableFuncWithdrawResults {
    pub(crate) id: i32,
}

pub struct MutableViewDonationResults {
    pub(crate) id: i32,
}

impl MutableViewDonationResults {
    pub fn amount(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_AMOUNT))
    }

    pub fn donator(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_RESULT_DONATOR))
    }

    pub fn error(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_RESULT_ERROR))
    }

    pub fn feedback(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_RESULT_FEEDBACK))
    }

    pub fn timestamp(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_TIMESTAMP))
    }
}

pub struct ImmutableViewDonationResults {
    pub(crate) id: i32,
}

impl ImmutableViewDonationResults {
    pub fn amount(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_AMOUNT))
    }

    pub fn donator(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_RESULT_DONATOR))
    }

    pub fn error(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_RESULT_ERROR))
    }

    pub fn feedback(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_RESULT_FEEDBACK))
    }

    pub fn timestamp(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_TIMESTAMP))
    }
}

pub struct MutableViewDonationInfoResults {
    pub(crate) id: i32,
}

impl MutableViewDonationInfoResults {
    pub fn count(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_COUNT))
    }

    pub fn max_donation(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_MAX_DONATION))
    }

    pub fn total_donation(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_TOTAL_DONATION))
    }
}

pub struct ImmutableViewDonationInfoResults {
    pub(crate) id: i32,
}

impl ImmutableViewDonationInfoResults {
    pub fn count(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_COUNT))
    }

    pub fn max_donation(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_MAX_DONATION))
    }

    pub fn total_donation(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_TOTAL_DONATION))
    }
}
