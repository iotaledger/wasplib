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
pub struct ImmutableCallOnChainParams {
    pub(crate) id: i32,
}

impl ImmutableCallOnChainParams {
    pub fn hname_contract(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_CONTRACT))
    }

    pub fn hname_ep(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_EP))
    }

    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableCallOnChainParams {
    pub(crate) id: i32,
}

impl MutableCallOnChainParams {
    pub fn hname_contract(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_CONTRACT))
    }

    pub fn hname_ep(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_EP))
    }

    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableCheckContextFromFullEPParams {
    pub(crate) id: i32,
}

impl ImmutableCheckContextFromFullEPParams {
    pub fn agent_id(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_AGENT_ID))
    }

    pub fn caller(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_CALLER))
    }

    pub fn chain_id(&self) -> ScImmutableChainId {
        ScImmutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }

    pub fn chain_owner_id(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_CHAIN_OWNER_ID))
    }

    pub fn contract_creator(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_CONTRACT_CREATOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableCheckContextFromFullEPParams {
    pub(crate) id: i32,
}

impl MutableCheckContextFromFullEPParams {
    pub fn agent_id(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_AGENT_ID))
    }

    pub fn caller(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_CALLER))
    }

    pub fn chain_id(&self) -> ScMutableChainId {
        ScMutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }

    pub fn chain_owner_id(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_CHAIN_OWNER_ID))
    }

    pub fn contract_creator(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_CONTRACT_CREATOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutablePassTypesFullParams {
    pub(crate) id: i32,
}

impl ImmutablePassTypesFullParams {
    pub fn hash(&self) -> ScImmutableHash {
        ScImmutableHash::new(self.id, idx_map(IDX_PARAM_HASH))
    }

    pub fn hname(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME))
    }

    pub fn hname_zero(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_ZERO))
    }

    pub fn int64(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT64))
    }

    pub fn int64_zero(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT64_ZERO))
    }

    pub fn string(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_STRING))
    }

    pub fn string_zero(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_STRING_ZERO))
    }
}

#[derive(Clone, Copy)]
pub struct MutablePassTypesFullParams {
    pub(crate) id: i32,
}

impl MutablePassTypesFullParams {
    pub fn hash(&self) -> ScMutableHash {
        ScMutableHash::new(self.id, idx_map(IDX_PARAM_HASH))
    }

    pub fn hname(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME))
    }

    pub fn hname_zero(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_ZERO))
    }

    pub fn int64(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT64))
    }

    pub fn int64_zero(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT64_ZERO))
    }

    pub fn string(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_STRING))
    }

    pub fn string_zero(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_STRING_ZERO))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableRunRecursionParams {
    pub(crate) id: i32,
}

impl ImmutableRunRecursionParams {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableRunRecursionParams {
    pub(crate) id: i32,
}

impl MutableRunRecursionParams {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSendToAddressParams {
    pub(crate) id: i32,
}

impl ImmutableSendToAddressParams {
    pub fn address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }
}

#[derive(Clone, Copy)]
pub struct MutableSendToAddressParams {
    pub(crate) id: i32,
}

impl MutableSendToAddressParams {
    pub fn address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSetIntParams {
    pub(crate) id: i32,
}

impl ImmutableSetIntParams {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }

    pub fn name(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_NAME))
    }
}

#[derive(Clone, Copy)]
pub struct MutableSetIntParams {
    pub(crate) id: i32,
}

impl MutableSetIntParams {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }

    pub fn name(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_NAME))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableTestEventLogGenericDataParams {
    pub(crate) id: i32,
}

impl ImmutableTestEventLogGenericDataParams {
    pub fn counter(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_COUNTER))
    }
}

#[derive(Clone, Copy)]
pub struct MutableTestEventLogGenericDataParams {
    pub(crate) id: i32,
}

impl MutableTestEventLogGenericDataParams {
    pub fn counter(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_COUNTER))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableWithdrawToChainParams {
    pub(crate) id: i32,
}

impl ImmutableWithdrawToChainParams {
    pub fn chain_id(&self) -> ScImmutableChainId {
        ScImmutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }
}

#[derive(Clone, Copy)]
pub struct MutableWithdrawToChainParams {
    pub(crate) id: i32,
}

impl MutableWithdrawToChainParams {
    pub fn chain_id(&self) -> ScMutableChainId {
        ScMutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableCheckContextFromViewEPParams {
    pub(crate) id: i32,
}

impl ImmutableCheckContextFromViewEPParams {
    pub fn agent_id(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_AGENT_ID))
    }

    pub fn chain_id(&self) -> ScImmutableChainId {
        ScImmutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }

    pub fn chain_owner_id(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_CHAIN_OWNER_ID))
    }

    pub fn contract_creator(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, idx_map(IDX_PARAM_CONTRACT_CREATOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableCheckContextFromViewEPParams {
    pub(crate) id: i32,
}

impl MutableCheckContextFromViewEPParams {
    pub fn agent_id(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_AGENT_ID))
    }

    pub fn chain_id(&self) -> ScMutableChainId {
        ScMutableChainId::new(self.id, idx_map(IDX_PARAM_CHAIN_ID))
    }

    pub fn chain_owner_id(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_CHAIN_OWNER_ID))
    }

    pub fn contract_creator(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, idx_map(IDX_PARAM_CONTRACT_CREATOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableFibonacciParams {
    pub(crate) id: i32,
}

impl ImmutableFibonacciParams {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFibonacciParams {
    pub(crate) id: i32,
}

impl MutableFibonacciParams {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetIntParams {
    pub(crate) id: i32,
}

impl ImmutableGetIntParams {
    pub fn name(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_NAME))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetIntParams {
    pub(crate) id: i32,
}

impl MutableGetIntParams {
    pub fn name(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_NAME))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutablePassTypesViewParams {
    pub(crate) id: i32,
}

impl ImmutablePassTypesViewParams {
    pub fn hash(&self) -> ScImmutableHash {
        ScImmutableHash::new(self.id, idx_map(IDX_PARAM_HASH))
    }

    pub fn hname(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME))
    }

    pub fn hname_zero(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_ZERO))
    }

    pub fn int64(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT64))
    }

    pub fn int64_zero(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_INT64_ZERO))
    }

    pub fn string(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_STRING))
    }

    pub fn string_zero(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_PARAM_STRING_ZERO))
    }
}

#[derive(Clone, Copy)]
pub struct MutablePassTypesViewParams {
    pub(crate) id: i32,
}

impl MutablePassTypesViewParams {
    pub fn hash(&self) -> ScMutableHash {
        ScMutableHash::new(self.id, idx_map(IDX_PARAM_HASH))
    }

    pub fn hname(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME))
    }

    pub fn hname_zero(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_PARAM_HNAME_ZERO))
    }

    pub fn int64(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT64))
    }

    pub fn int64_zero(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_PARAM_INT64_ZERO))
    }

    pub fn string(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_STRING))
    }

    pub fn string_zero(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_PARAM_STRING_ZERO))
    }
}
