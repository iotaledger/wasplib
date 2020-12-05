// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

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
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("mintSupply", mintSupply);
    exports.add_call("updateMetadata", updateMetadata);
    exports.add_call("transferOwnership", transferOwnership);
}

fn mintSupply(sc: &ScCallContext) {
    let minted = sc.incoming().minted();
    if minted == ScColor::MINT {
        sc.log("TokenRegistry: No newly minted tokens found");
        return;
    }
    let state = sc.state();
    let registry = state.get_key_map("registry").get_bytes(minted.to_bytes());
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
        description: params.get_string("dscr").value(),
        user_defined: params.get_string("ud").value(),
    };
    if token.supply <= 0 {
        sc.log("TokenRegistry: Insufficient supply");
        return;
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    let data = encodeTokenInfo(&token);
    registry.set_value(&data);
    let colors = state.get_color_array("colorList");
    colors.get_color(colors.length()).set_value(&minted);
}

fn updateMetadata(_sc: &ScCallContext) {
    //TODO
}

fn transferOwnership(_sc: &ScCallContext) {
    //TODO
}

fn decodeTokenInfo(bytes: &[u8]) -> TokenInfo {
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

fn encodeTokenInfo(token: &TokenInfo) -> Vec<u8> {
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
