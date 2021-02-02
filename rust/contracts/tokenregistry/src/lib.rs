// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use schema::*;
use types::*;
use wasplib::client::*;

mod schema;
mod types;

fn func_mint_supply(ctx: &ScCallContext) {
    let minted = ctx.incoming().minted();
    if minted.equals(&ScColor::MINT) {
        ctx.panic("TokenRegistry: No newly minted tokens found");
    }
    let state = ctx.state();
    let registry = state.get_map(VAR_REGISTRY).get_bytes(&minted);
    if registry.exists() {
        ctx.panic("TokenRegistry: Color already exists");
    }
    let params = ctx.params();
    let mut token = TokenInfo {
        supply: ctx.incoming().balance(&minted),
        minted_by: ctx.caller(),
        owner: ctx.caller(),
        created: ctx.timestamp(),
        updated: ctx.timestamp(),
        description: params.get_string(PARAM_DESCRIPTION).value(),
        user_defined: params.get_string(PARAM_USER_DEFINED).value(),
    };
    if token.supply <= 0 {
        ctx.panic("TokenRegistry: Insufficient supply");
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    registry.set_value(&encode_token_info(&token));
    let colors = state.get_color_array(VAR_COLOR_LIST);
    colors.get_color(colors.length()).set_value(&minted);
}

fn func_update_metadata(_sc: &ScCallContext) {
    //TODO
}

fn func_transfer_ownership(_sc: &ScCallContext) {
    //TODO
}
