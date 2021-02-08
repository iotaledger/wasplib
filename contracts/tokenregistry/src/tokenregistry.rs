// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_mint_supply(ctx: &ScCallContext, params: &FuncMintSupplyParams) {
    let minted = ctx.incoming().minted();
    if minted.equals(&ScColor::MINT) {
        ctx.panic("TokenRegistry: No newly minted tokens found");
    }
    let state = ctx.state();
    let registry = state.get_map(VAR_REGISTRY).get_bytes(&minted);
    if registry.exists() {
        ctx.panic("TokenRegistry: Color already exists");
    }
    let mut token = TokenInfo {
        supply: ctx.incoming().balance(&minted),
        minted_by: ctx.caller(),
        owner: ctx.caller(),
        created: ctx.timestamp(),
        updated: ctx.timestamp(),
        description: params.description.value(),
        user_defined: params.user_defined.value(),
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

pub fn func_transfer_ownership(_sc: &ScCallContext, _params: &FuncTransferOwnershipParams) {
    //TODO
}

pub fn func_update_metadata(_sc: &ScCallContext, _params: &FuncUpdateMetadataParams) {
    //TODO
}

pub fn view_get_info(_sc: &ScViewContext, _params: &ViewGetInfoParams) {
    //TODO
}
