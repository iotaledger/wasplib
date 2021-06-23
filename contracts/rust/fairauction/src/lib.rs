// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use fairauction::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod params;
mod results;
mod state;
mod subtypes;
mod types;
mod fairauction;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_FINALIZE_AUCTION, func_finalize_auction_thunk);
    exports.add_func(FUNC_PLACE_BID, func_place_bid_thunk);
    exports.add_func(FUNC_SET_OWNER_MARGIN, func_set_owner_margin_thunk);
    exports.add_func(FUNC_START_AUCTION, func_start_auction_thunk);
    exports.add_view(VIEW_GET_INFO, view_get_info_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct FinalizeAuctionContext {
    params: ImmutableFinalizeAuctionParams,
    state:  MutableFairAuctionState,
}

fn func_finalize_auction_thunk(ctx: &ScFuncContext) {
    ctx.log("fairauction.funcFinalizeAuction");
    // only SC itself can invoke this function
    ctx.require(ctx.caller() == ctx.account_id(), "no permission");

    let f = FinalizeAuctionContext {
        params: ImmutableFinalizeAuctionParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableFairAuctionState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.color().exists(), "missing mandatory color");
    func_finalize_auction(ctx, &f);
    ctx.log("fairauction.funcFinalizeAuction ok");
}

pub struct PlaceBidContext {
    params: ImmutablePlaceBidParams,
    state:  MutableFairAuctionState,
}

fn func_place_bid_thunk(ctx: &ScFuncContext) {
    ctx.log("fairauction.funcPlaceBid");
    let f = PlaceBidContext {
        params: ImmutablePlaceBidParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableFairAuctionState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.color().exists(), "missing mandatory color");
    func_place_bid(ctx, &f);
    ctx.log("fairauction.funcPlaceBid ok");
}

pub struct SetOwnerMarginContext {
    params: ImmutableSetOwnerMarginParams,
    state:  MutableFairAuctionState,
}

fn func_set_owner_margin_thunk(ctx: &ScFuncContext) {
    ctx.log("fairauction.funcSetOwnerMargin");
    // only SC creator can set owner margin
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let f = SetOwnerMarginContext {
        params: ImmutableSetOwnerMarginParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableFairAuctionState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.owner_margin().exists(), "missing mandatory ownerMargin");
    func_set_owner_margin(ctx, &f);
    ctx.log("fairauction.funcSetOwnerMargin ok");
}

pub struct StartAuctionContext {
    params: ImmutableStartAuctionParams,
    state:  MutableFairAuctionState,
}

fn func_start_auction_thunk(ctx: &ScFuncContext) {
    ctx.log("fairauction.funcStartAuction");
    let f = StartAuctionContext {
        params: ImmutableStartAuctionParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableFairAuctionState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.color().exists(), "missing mandatory color");
    ctx.require(f.params.minimum_bid().exists(), "missing mandatory minimumBid");
    func_start_auction(ctx, &f);
    ctx.log("fairauction.funcStartAuction ok");
}

pub struct GetInfoContext {
    params:  ImmutableGetInfoParams,
    results: MutableGetInfoResults,
    state:   ImmutableFairAuctionState,
}

fn view_get_info_thunk(ctx: &ScViewContext) {
    ctx.log("fairauction.viewGetInfo");
    let f = GetInfoContext {
        params: ImmutableGetInfoParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        results: MutableGetInfoResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: ImmutableFairAuctionState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.color().exists(), "missing mandatory color");
    view_get_info(ctx, &f);
    ctx.log("fairauction.viewGetInfo ok");
}

//@formatter:on
