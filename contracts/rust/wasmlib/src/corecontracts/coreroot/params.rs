// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coreroot::*;
use crate::host::*;

#[derive(Clone, Copy)]
pub struct ImmutableDelegateChainOwnershipParams {
    pub(crate) id: i32,
}

impl ImmutableDelegateChainOwnershipParams {
    pub fn chain_owner(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, PARAM_CHAIN_OWNER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableDelegateChainOwnershipParams {
    pub(crate) id: i32,
}

impl MutableDelegateChainOwnershipParams {
    pub fn chain_owner(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, PARAM_CHAIN_OWNER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableDeployContractParams {
    pub(crate) id: i32,
}

impl ImmutableDeployContractParams {
    pub fn description(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, PARAM_DESCRIPTION.get_key_id())
    }

    pub fn name(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, PARAM_NAME.get_key_id())
    }

    pub fn program_hash(&self) -> ScImmutableHash {
        ScImmutableHash::new(self.id, PARAM_PROGRAM_HASH.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableDeployContractParams {
    pub(crate) id: i32,
}

impl MutableDeployContractParams {
    pub fn description(&self) -> ScMutableString {
        ScMutableString::new(self.id, PARAM_DESCRIPTION.get_key_id())
    }

    pub fn name(&self) -> ScMutableString {
        ScMutableString::new(self.id, PARAM_NAME.get_key_id())
    }

    pub fn program_hash(&self) -> ScMutableHash {
        ScMutableHash::new(self.id, PARAM_PROGRAM_HASH.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGrantDeployPermissionParams {
    pub(crate) id: i32,
}

impl ImmutableGrantDeployPermissionParams {
    pub fn deployer(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, PARAM_DEPLOYER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGrantDeployPermissionParams {
    pub(crate) id: i32,
}

impl MutableGrantDeployPermissionParams {
    pub fn deployer(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, PARAM_DEPLOYER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableRevokeDeployPermissionParams {
    pub(crate) id: i32,
}

impl ImmutableRevokeDeployPermissionParams {
    pub fn deployer(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.id, PARAM_DEPLOYER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableRevokeDeployPermissionParams {
    pub(crate) id: i32,
}

impl MutableRevokeDeployPermissionParams {
    pub fn deployer(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.id, PARAM_DEPLOYER.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSetContractFeeParams {
    pub(crate) id: i32,
}

impl ImmutableSetContractFeeParams {
    pub fn hname(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }

    pub fn owner_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, PARAM_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, PARAM_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableSetContractFeeParams {
    pub(crate) id: i32,
}

impl MutableSetContractFeeParams {
    pub fn hname(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }

    pub fn owner_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, PARAM_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, PARAM_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableSetDefaultFeeParams {
    pub(crate) id: i32,
}

impl ImmutableSetDefaultFeeParams {
    pub fn owner_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, PARAM_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, PARAM_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableSetDefaultFeeParams {
    pub(crate) id: i32,
}

impl MutableSetDefaultFeeParams {
    pub fn owner_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, PARAM_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, PARAM_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableFindContractParams {
    pub(crate) id: i32,
}

impl ImmutableFindContractParams {
    pub fn hname(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableFindContractParams {
    pub(crate) id: i32,
}

impl MutableFindContractParams {
    pub fn hname(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetFeeInfoParams {
    pub(crate) id: i32,
}

impl ImmutableGetFeeInfoParams {
    pub fn hname(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetFeeInfoParams {
    pub(crate) id: i32,
}

impl MutableGetFeeInfoParams {
    pub fn hname(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, PARAM_HNAME.get_key_id())
    }
}
