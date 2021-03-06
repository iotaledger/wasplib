// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coreaccounts::*;
use crate::host::*;

#[derive(Clone, Copy)]
pub struct ImmutableDepositParams {
    pub(crate) id: i32,
}

impl ImmutableDepositParams {
    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, PARAM_AGENT_ID.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableDepositParams {
    pub(crate) id: i32,
}

impl MutableDepositParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, PARAM_AGENT_ID.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableBalanceParams {
    pub(crate) id: i32,
}

impl ImmutableBalanceParams {
    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, PARAM_AGENT_ID.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableBalanceParams {
    pub(crate) id: i32,
}

impl MutableBalanceParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, PARAM_AGENT_ID.get_key_id())
    }
}
