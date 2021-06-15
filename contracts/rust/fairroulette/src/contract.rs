// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct FairRouletteFunc {
    sc: ScContractFunc,
}

impl FairRouletteFunc {
    pub fn new(ctx: &ScFuncContext) -> FairRouletteFunc {
        FairRouletteFunc { sc: ScContractFunc::new(ctx, HSC_NAME) }
    }

    pub fn delay(&mut self, seconds: i32) -> &mut FairRouletteFunc {
        self.sc.delay(seconds);
        self
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut FairRouletteFunc {
        self.sc.of_contract(contract);
        self
    }

    pub fn post(&mut self) -> &mut FairRouletteFunc {
        self.sc.post();
        self
    }

    pub fn post_to_chain(&mut self, chain_id: ScChainId) -> &mut FairRouletteFunc {
        self.sc.post_to_chain(chain_id);
        self
    }

    pub fn lock_bets(&mut self, transfer: ScTransfers) {
        self.sc.run(HFUNC_LOCK_BETS, 0, Some(transfer));
    }

    pub fn pay_winners(&mut self, transfer: ScTransfers) {
        self.sc.run(HFUNC_PAY_WINNERS, 0, Some(transfer));
    }

    pub fn place_bet(&mut self, params: MutableFuncPlaceBetParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_PLACE_BET, params.id, Some(transfer));
    }

    pub fn play_period(&mut self, params: MutableFuncPlayPeriodParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_PLAY_PERIOD, params.id, Some(transfer));
    }

    pub fn last_winning_number(&mut self) -> ImmutableViewLastWinningNumberResults {
        self.sc.run(HVIEW_LAST_WINNING_NUMBER, 0, None);
        ImmutableViewLastWinningNumberResults { id: self.sc.result_map_id() }
    }
}

pub struct FairRouletteView {
    sc: ScContractView,
}

impl FairRouletteView {
    pub fn new(ctx: &ScViewContext) -> FairRouletteView {
        FairRouletteView { sc: ScContractView::new(ctx, HSC_NAME) }
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut FairRouletteView {
        self.sc.of_contract(contract);
        self
    }

    pub fn last_winning_number(&mut self) -> ImmutableViewLastWinningNumberResults {
        self.sc.run(HVIEW_LAST_WINNING_NUMBER, 0);
        ImmutableViewLastWinningNumberResults { id: self.sc.result_map_id() }
    }
}
