// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// base contract objects

use crate::context::*;
use crate::hashtypes::*;
use crate::immutable::*;
use crate::mutable::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScContractFunc {
    ctx: ScFuncContext,
    chain_id: ScChainId,
    contract: ScHname,
    delay: i32,
    post: bool,
    results: ScImmutableMap,
}

impl ScContractFunc {
    pub fn new(ctx: &ScFuncContext, contract: ScHname) -> ScContractFunc {
        ScContractFunc {
            ctx: ctx.clone(),
            chain_id: ctx.chain_id(),
            contract: contract,
            delay: 0,
            post: false,
            results: ScImmutableMap { obj_id: 0 },
        }
    }

    pub fn delay(&mut self, seconds: i32) -> &ScContractFunc {
        self.delay = seconds;
        self
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &ScContractFunc {
        self.contract = contract;
        self
    }

    pub fn post(&mut self) -> &ScContractFunc {
        self.post = true;
        self
    }

    pub fn post_to_chain(&mut self, chain_id: ScChainId) -> &ScContractFunc {
        self.post = true;
        self.chain_id = chain_id;
        self
    }

    pub fn result_map_id(&self) -> i32 {
        let map_id = self.results.map_id();
        self.ctx.require(map_id != 0, "Cannot get results from asynchronous post");
        map_id
    }

    pub fn run(&mut self, function: ScHname, params_id: i32, transfer: Option<ScTransfers>) {
        let params = ScMutableMap { obj_id: params_id };
        if self.post {
            self.ctx.require(transfer.is_some(), "Cannot post to view");
            self.ctx.post(&self.chain_id, self.contract, function, Some(params), transfer.unwrap(), self.delay);
            return;
        }

        self.results = self.ctx.call(self.contract, function, Some(params), transfer);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScContractView {
    ctx: ScViewContext,
    contract: ScHname,
    results: ScImmutableMap,
}

impl ScContractView {
    pub fn new(ctx: &ScViewContext, contract: ScHname) -> ScContractView {
        ScContractView {
            ctx: ctx.clone(),
            contract: contract,
            results: ScImmutableMap { obj_id: 0 },
        }
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &ScContractView {
        self.contract = contract;
        self
    }

    pub fn result_map_id(&self) -> i32 {
        self.results.map_id()
    }

    pub fn run(&mut self, function: ScHname, params_id: i32) {
        let params = ScMutableMap { obj_id: params_id };
        self.results = self.ctx.call(self.contract, function, Some(params));
    }
}
