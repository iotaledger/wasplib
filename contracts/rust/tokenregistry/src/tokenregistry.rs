// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::types::*;

pub fn func_mint_supply(ctx: &ScFuncContext, f: &MintSupplyContext) {
    let minted = ctx.minted();
    let minted_colors = minted.colors();
    ctx.require(minted_colors.length() == 1, "need single minted color");
    let minted_color = minted_colors.get_color(0).value();
    let current_token = f.state.registry().get_token(&minted_color);
    if current_token.exists() {
        // should never happen, because transaction id is unique
        ctx.panic("TokenRegistry: registry for color already exists");
    }
    let mut token = Token {
        supply: minted.balance(&minted_color),
        minted_by: ctx.caller(),
        owner: ctx.caller(),
        created: ctx.timestamp(),
        updated: ctx.timestamp(),
        description: f.params.description().value(),
        user_defined: f.params.user_defined().value(),
    };
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    current_token.set_value(&token);
    let color_list = f.state.color_list();
    color_list.get_color(color_list.length()).set_value(&minted_color);
}

pub fn func_transfer_ownership(_ctx: &ScFuncContext, _f: &TransferOwnershipContext) {
    // TODO
}

pub fn func_update_metadata(_ctx: &ScFuncContext, _f: &UpdateMetadataContext) {
    // TODO
}

pub fn view_get_info(_ctx: &ScViewContext, _f: &GetInfoContext) {
    // TODO
}
