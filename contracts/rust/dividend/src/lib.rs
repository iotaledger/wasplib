// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use dividend::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::state::*;

mod consts;
mod keys;
mod state;
mod dividend;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DIVIDE, func_divide_thunk);
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_MEMBER, func_member_thunk);
    exports.add_func(FUNC_SET_OWNER, func_set_owner_thunk);
    exports.add_view(VIEW_GET_FACTOR, view_get_factor_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct FuncDivideContext {
    state: DividendFuncState,
}

fn func_divide_thunk(ctx: &ScFuncContext) {
    ctx.log("dividend.funcDivide");
    let f = FuncDivideContext {
        state: DividendFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_divide(ctx, &f);
    ctx.log("dividend.funcDivide ok");
}

pub struct FuncInitParams {
    pub owner: ScImmutableAgentId, // optional owner, defaults to contract creator
}

pub struct FuncInitContext {
    params: FuncInitParams,
    state:  DividendFuncState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
    ctx.log("dividend.funcInit");
    let p = ctx.params().map_id();
    let f = FuncInitContext {
        params: FuncInitParams {
            owner: ScImmutableAgentId::new(p, idx_map(IDX_PARAM_OWNER)),
        },
        state: DividendFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_init(ctx, &f);
    ctx.log("dividend.funcInit ok");
}

pub struct FuncMemberParams {
    pub address: ScImmutableAddress, // address of dividend recipient
    pub factor:  ScImmutableInt64,   // relative division factor
}

pub struct FuncMemberContext {
    params: FuncMemberParams,
    state:  DividendFuncState,
}

fn func_member_thunk(ctx: &ScFuncContext) {
    ctx.log("dividend.funcMember");
    // only defined owner can add members
    let access = ctx.state().get_agent_id("owner");
    ctx.require(access.exists(), "access not set: owner");
    ctx.require(ctx.caller() == access.value(), "no permission");

    let p = ctx.params().map_id();
    let f = FuncMemberContext {
        params: FuncMemberParams {
            address: ScImmutableAddress::new(p, idx_map(IDX_PARAM_ADDRESS)),
            factor:  ScImmutableInt64::new(p, idx_map(IDX_PARAM_FACTOR)),
        },
        state: DividendFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.address.exists(), "missing mandatory address");
    ctx.require(f.params.factor.exists(), "missing mandatory factor");
    func_member(ctx, &f);
    ctx.log("dividend.funcMember ok");
}

pub struct FuncSetOwnerParams {
    pub owner: ScImmutableAgentId, // new owner of smart contract
}

pub struct FuncSetOwnerContext {
    params: FuncSetOwnerParams,
    state:  DividendFuncState,
}

fn func_set_owner_thunk(ctx: &ScFuncContext) {
    ctx.log("dividend.funcSetOwner");
    // only defined owner can change owner
    let access = ctx.state().get_agent_id("owner");
    ctx.require(access.exists(), "access not set: owner");
    ctx.require(ctx.caller() == access.value(), "no permission");

    let p = ctx.params().map_id();
    let f = FuncSetOwnerContext {
        params: FuncSetOwnerParams {
            owner: ScImmutableAgentId::new(p, idx_map(IDX_PARAM_OWNER)),
        },
        state: DividendFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.owner.exists(), "missing mandatory owner");
    func_set_owner(ctx, &f);
    ctx.log("dividend.funcSetOwner ok");
}

pub struct ViewGetFactorParams {
    pub address: ScImmutableAddress, // address of dividend recipient
}

pub struct ViewGetFactorResults {
    pub factor: ScMutableInt64, // relative division factor
}

pub struct ViewGetFactorContext {
    params:  ViewGetFactorParams,
    results: ViewGetFactorResults,
    state:   DividendViewState,
}

fn view_get_factor_thunk(ctx: &ScViewContext) {
    ctx.log("dividend.viewGetFactor");
    let p = ctx.params().map_id();
    let r = ctx.results().map_id();
    let f = ViewGetFactorContext {
        params: ViewGetFactorParams {
            address: ScImmutableAddress::new(p, idx_map(IDX_PARAM_ADDRESS)),
        },
        results: ViewGetFactorResults {
            factor: ScMutableInt64::new(r, idx_map(IDX_RESULT_FACTOR)),
        },
        state: DividendViewState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.address.exists(), "missing mandatory address");
    view_get_factor(ctx, &f);
    ctx.log("dividend.viewGetFactor ok");
}

//@formatter:on
