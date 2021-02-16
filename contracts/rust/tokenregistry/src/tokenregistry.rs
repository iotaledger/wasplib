// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_mint_supply(ctx: &ScFuncContext) {
    let p = ctx.params();
    let param_description = p.get_string(PARAM_DESCRIPTION);
    let param_user_defined = p.get_string(PARAM_USER_DEFINED);

    let minted = ctx.incoming().minted();
    if minted == ScColor::MINT {
        ctx.panic("TokenRegistry: No newly minted tokens found");
    }
    let state = ctx.state();
    let registry = state.get_map(VAR_REGISTRY).get_bytes(&minted);
    if registry.exists() {
        ctx.panic("TokenRegistry: Color already exists");
    }
    let mut token = Token {
        supply: ctx.incoming().balance(&minted),
        minted_by: ctx.caller(),
        owner: ctx.caller(),
        created: ctx.timestamp(),
        updated: ctx.timestamp(),
        description: param_description.value(),
        user_defined: param_user_defined.value(),
    };
    if token.supply <= 0 {
        ctx.panic("TokenRegistry: Insufficient supply");
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    registry.set_value(&token.to_bytes());
    let colors = state.get_color_array(VAR_COLOR_LIST);
    colors.get_color(colors.length()).set_value(&minted);
}

pub fn func_transfer_ownership(ctx: &ScFuncContext) {
    //TODO the one who can transfer token ownership
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let param_color = p.get_color(PARAM_COLOR);
    ctx.require(param_color.exists(), "missing mandatory color");

    //TODO
}

pub fn func_update_metadata(ctx: &ScFuncContext) {
    //TODO the one who can change the token info
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let param_color = p.get_color(PARAM_COLOR);
    ctx.require(param_color.exists(), "missing mandatory color");

    //TODO
}

pub fn view_get_info(ctx: &ScViewContext) {
    let p = ctx.params();
    let param_color = p.get_color(PARAM_COLOR);
    ctx.require(param_color.exists(), "missing mandatory color");

    //TODO
}
