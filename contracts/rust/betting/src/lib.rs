// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use betting::*;
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
mod betting;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_SET_OWNER, func_set_owner_thunk);
    exports.add_view(VIEW_GET_OWNER, view_get_owner_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct InitContext {
    params: ImmutableInitParams,
    state:  MutableBettingState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
    ctx.log("betting.funcInit");
    let f = InitContext {
        params: ImmutableInitParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableBettingState {
            id: OBJ_ID_STATE,
        },
    };
    func_init(ctx, &f);
    ctx.log("betting.funcInit ok");
}

pub struct SetOwnerContext {
    params: ImmutableSetOwnerParams,
    state:  MutableBettingState,
}

fn func_set_owner_thunk(ctx: &ScFuncContext) {
    ctx.log("betting.funcSetOwner");
    // current owner of this smart contract
    let access = ctx.state().get_agent_id("owner");
    ctx.require(access.exists(), "access not set: owner");
    ctx.require(ctx.caller() == access.value(), "no permission");

    let f = SetOwnerContext {
        params: ImmutableSetOwnerParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableBettingState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.owner().exists(), "missing mandatory owner");
    func_set_owner(ctx, &f);
    ctx.log("betting.funcSetOwner ok");
}

pub struct GetOwnerContext {
    results: MutableGetOwnerResults,
    state:   ImmutableBettingState,
}

fn view_get_owner_thunk(ctx: &ScViewContext) {
    ctx.log("betting.viewGetOwner");
    let f = GetOwnerContext {
        results: MutableGetOwnerResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableBettingState {
            id: OBJ_ID_STATE,
        },
    };
    view_get_owner(ctx, &f);
    ctx.log("betting.viewGetOwner ok");
}

// @formatter:on
