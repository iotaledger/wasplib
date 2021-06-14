// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]

use crate::*;
use crate::corecontracts::coreaccounts::*;

pub struct CoreAccountsFunc {
    sc: ScContractFunc,
}

impl CoreAccountsFunc {
    pub fn new(ctx: &ScFuncContext) -> CoreAccountsFunc {
        CoreAccountsFunc { sc: ScContractFunc::new(ctx, HSC_NAME) }
    }

    pub fn delay(&mut self, seconds: i64) -> &mut CoreAccountsFunc {
        self.sc.delay(seconds);
        self
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut CoreAccountsFunc {
        self.sc.of_contract(contract);
        self
    }

    pub fn post(&mut self) -> &mut CoreAccountsFunc {
        self.sc.post();
        self
    }

    pub fn post_to_chain(&mut self, chain_id: ScChainId) -> &mut CoreAccountsFunc {
        self.sc.post_to_chain(chain_id);
        self
    }

    pub fn deposit(&mut self, params: MutableFuncDepositParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_DEPOSIT, params.id, Some(transfer));
    }

    pub fn withdraw(&mut self, transfer: ScTransfers) {
        self.sc.run(HFUNC_WITHDRAW, 0, Some(transfer));
    }

    pub fn accounts(&mut self) {
        self.sc.run(HVIEW_ACCOUNTS, 0, None);
    }

    pub fn balance(&mut self, params: MutableViewBalanceParams) {
        self.sc.run(HVIEW_BALANCE, params.id, None);
    }

    pub fn total_assets(&mut self) {
        self.sc.run(HVIEW_TOTAL_ASSETS, 0, None);
    }
}

pub struct CoreAccountsView {
    sc: ScContractView,
}

impl CoreAccountsView {
    pub fn new(ctx: &ScViewContext) -> CoreAccountsView {
        CoreAccountsView { sc: ScContractView::new(ctx, HSC_NAME) }
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut CoreAccountsView {
        self.sc.of_contract(contract);
        self
    }

    pub fn accounts(&mut self) {
        self.sc.run(HVIEW_ACCOUNTS, 0);
    }

    pub fn balance(&mut self, params: MutableViewBalanceParams) {
        self.sc.run(HVIEW_BALANCE, params.id);
    }

    pub fn total_assets(&mut self) {
        self.sc.run(HVIEW_TOTAL_ASSETS, 0);
    }
}