// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_mint_supply(ctx: &ScFuncContext, params: &FuncMintSupplyParams) {
    let minted_supply = ctx.minted_supply();
    if minted_supply == 0 {
        ctx.panic("TokenRegistry: No newly minted tokens found");
    }
    let minted_color = ctx.minted_color();
    let state = ctx.state();
    let registry = state.get_map(VAR_REGISTRY).get_bytes(&minted_color);
    if registry.exists() {
        // should never happen, because transaction id is unique
        ctx.panic("TokenRegistry: registry for color already exists");
    }
    let mut token = Token {
        supply: minted_supply,
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
    ctx.log("tokenregistry.mintSupply ok");
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
