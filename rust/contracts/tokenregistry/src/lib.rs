// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use types::*;
use wasplib::client::*;

mod types;

const KEY_COLOR_LIST: &str = "color_list";
const KEY_DESCRIPTION: &str = "description";
const KEY_REGISTRY: &str = "registry";
const KEY_USER_DEFINED: &str = "user_defined";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("mint_supply", mint_supply);
    exports.add_call("update_metadata", update_metadata);
    exports.add_call("transfer_ownership", transfer_ownership);
}

fn mint_supply(ctx: &ScCallContext) {
    let minted = ctx.incoming().minted();
    if minted.equals(&ScColor::MINT) {
        ctx.panic("TokenRegistry: No newly minted tokens found");
    }
    let state = ctx.state();
    let registry = state.get_map(KEY_REGISTRY).get_bytes(&minted);
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
        description: params.get_string(KEY_DESCRIPTION).value(),
        user_defined: params.get_string(KEY_USER_DEFINED).value(),
    };
    if token.supply <= 0 {
        ctx.panic("TokenRegistry: Insufficient supply");
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    registry.set_value(&encode_token_info(&token));
    let colors = state.get_color_array(KEY_COLOR_LIST);
    colors.get_color(colors.length()).set_value(&minted);
}

fn update_metadata(_sc: &ScCallContext) {
    //TODO
}

fn transfer_ownership(_sc: &ScCallContext) {
    //TODO
}
