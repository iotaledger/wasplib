// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;

const KEY_COLOR_LIST: &str = "color_list";
const KEY_DESCRIPTION: &str = "description";
const KEY_REGISTRY: &str = "registry";
const KEY_USER_DEFINED: &str = "user_defined";

struct TokenInfo {
    supply: i64,
    minted_by: ScAgent,
    owner: ScAgent,
    created: i64,
    updated: i64,
    description: String,
    user_defined: String,
}

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("mint_supply", mint_supply);
    exports.add_call("update_metadata", update_metadata);
    exports.add_call("transfer_ownership", transfer_ownership);
}

fn mint_supply(sc: &ScCallContext) {
    let minted = sc.incoming().minted();
    if minted == ScColor::MINT {
        sc.log("TokenRegistry: No newly minted tokens found");
        return;
    }
    let state = sc.state();
    let registry = state.get_map(KEY_REGISTRY).get_bytes(&minted);
    if registry.exists() {
        sc.log("TokenRegistry: Color already exists");
        return;
    }
    let params = sc.params();
    let mut token = TokenInfo {
        supply: sc.incoming().balance(&minted),
        minted_by: sc.caller(),
        owner: sc.caller(),
        created: sc.timestamp(),
        updated: sc.timestamp(),
        description: params.get_string(KEY_DESCRIPTION).value(),
        user_defined: params.get_string(KEY_USER_DEFINED).value(),
    };
    if token.supply <= 0 {
        sc.log("TokenRegistry: Insufficient supply");
        return;
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    let data = encode_token_info(&token);
    registry.set_value(&data);
    let colors = state.get_color_array(KEY_COLOR_LIST);
    colors.get_color(colors.length()).set_value(&minted);
}

fn update_metadata(_sc: &ScCallContext) {
    //TODO
}

fn transfer_ownership(_sc: &ScCallContext) {
    //TODO
}

fn decode_token_info(bytes: &[u8]) -> TokenInfo {
    let mut decoder = BytesDecoder::new(bytes);
    TokenInfo {
        supply: decoder.int(),
        minted_by: decoder.agent(),
        owner: decoder.agent(),
        created: decoder.int(),
        updated: decoder.int(),
        description: decoder.string(),
        user_defined: decoder.string(),
    }
}

fn encode_token_info(token: &TokenInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(token.supply);
    encoder.agent(&token.minted_by);
    encoder.agent(&token.owner);
    encoder.int(token.created);
    encoder.int(token.updated);
    encoder.string(&token.description);
    encoder.string(&token.user_defined);
    encoder.data()
}
