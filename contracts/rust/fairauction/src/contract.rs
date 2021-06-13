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

pub struct FairAuctionFunc {
    sc: ScContractFunc,
}

impl FairAuctionFunc {
    pub fn new(ctx: &ScFuncContext) -> FairAuctionFunc {
        FairAuctionFunc { sc: ScContractFunc::new(ctx, HSC_NAME) }
    }

    pub fn delay(&mut self, seconds: i64) -> &mut FairAuctionFunc {
        self.sc.delay(seconds);
        self
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut FairAuctionFunc {
        self.sc.of_contract(contract);
        self
    }

    pub fn post(&mut self) -> &mut FairAuctionFunc {
        self.sc.post();
        self
    }

    pub fn post_to_chain(&mut self, chain_id: ScChainId) -> &mut FairAuctionFunc {
        self.sc.post_to_chain(chain_id);
        self
    }

    pub fn finalize_auction(&mut self, params: MutableFuncFinalizeAuctionParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_FINALIZE_AUCTION, params.id, Some(transfer));
    }

    pub fn place_bid(&mut self, params: MutableFuncPlaceBidParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_PLACE_BID, params.id, Some(transfer));
    }

    pub fn set_owner_margin(&mut self, params: MutableFuncSetOwnerMarginParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_SET_OWNER_MARGIN, params.id, Some(transfer));
    }

    pub fn start_auction(&mut self, params: MutableFuncStartAuctionParams, transfer: ScTransfers) {
        self.sc.run(HFUNC_START_AUCTION, params.id, Some(transfer));
    }

    pub fn get_info(&mut self, params: MutableViewGetInfoParams) -> ImmutableViewGetInfoResults {
        self.sc.run(HVIEW_GET_INFO, params.id, None);
        ImmutableViewGetInfoResults { id: self.sc.result_map_id() }
    }
}

pub struct FairAuctionView {
    sc: ScContractView,
}

impl FairAuctionView {
    pub fn new(ctx: &ScViewContext) -> FairAuctionView {
        FairAuctionView { sc: ScContractView::new(ctx, HSC_NAME) }
    }

    pub fn of_contract(&mut self, contract: ScHname) -> &mut FairAuctionView {
        self.sc.of_contract(contract);
        self
    }

    pub fn get_info(&mut self, params: MutableViewGetInfoParams) -> ImmutableViewGetInfoResults {
        self.sc.run(HVIEW_GET_INFO, params.id);
        ImmutableViewGetInfoResults { id: self.sc.result_map_id() }
    }
}
