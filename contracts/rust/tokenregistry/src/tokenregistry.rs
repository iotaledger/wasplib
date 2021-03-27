// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_mint_supply(ctx: &ScFuncContext, params: &FuncMintSupplyParams) {
    let minted = ctx.minted();
    let minted_colors = minted.colors();
    ctx.require(minted_colors.length() == 1, "need single minted color");
    let minted_color = minted_colors.get_color(0).value();
    let state = ctx.state();
    let registry = state.get_map(VAR_REGISTRY).get_bytes(&minted_color);
    if registry.exists() {
        // should never happen, because transaction id is unique
        ctx.panic("TokenRegistry: registry for color already exists");
    }
    let mut token = Token {
        supply: minted.balance(&minted_color),
        minted_by: ctx.caller(),
        owner: ctx.caller(),
        created: ctx.timestamp(),
        updated: ctx.timestamp(),
        description: params.description.value(),
        user_defined: params.user_defined.value(),
    };
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    registry.set_value(&token.to_bytes());
    let colors = state.get_color_array(VAR_COLOR_LIST);
    colors.get_color(colors.length()).set_value(&minted_color);
}

pub fn func_transfer_ownership(_ctx: &ScFuncContext, _params: &FuncTransferOwnershipParams) {
    //TODO
}

pub fn func_update_metadata(_ctx: &ScFuncContext, _params: &FuncUpdateMetadataParams) {
    //TODO
}

pub fn view_get_info(_ctx: &ScViewContext, _params: &ViewGetInfoParams) {
    //TODO
}
