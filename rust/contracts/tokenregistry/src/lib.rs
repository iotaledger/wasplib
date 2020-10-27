#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

struct TokenInfo {
    supply: i64,
    minted_by: ScAddress,
    owner: ScAddress,
    created: i64,
    updated: i64,
    description: String,
    user_defined: String,
}

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("mintSupply");
    exports.add("updateMetadata");
    exports.add("transferOwnership");
}

#[no_mangle]
pub fn mintSupply() {
    let sc = ScContext::new();
    let request = sc.request();
    let color = request.minted_color();
    let state = sc.state();
    let registry = state.get_key_map("registry").get_bytes(color.to_bytes());
    if registry.exists() {
        sc.log("TokenRegistry: Color already exists");
        return;
    }
    let params = request.params();
    let mut token = TokenInfo {
        supply: request.balance(&color),
        minted_by: request.address(),
        owner: request.address(),
        created: request.timestamp(),
        updated: request.timestamp(),
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
    colors.get_color(colors.length()).set_value(&color);
}

#[no_mangle]
pub fn updateMetadata() {
    //let sc = ScContext::new();
    //TODO
}

#[no_mangle]
pub fn transferOwnership() {
    //let sc = ScContext::new();
    //TODO
}

fn decodeTokenInfo(bytes: &[u8]) -> TokenInfo {
    let mut decoder = BytesDecoder::new(bytes);
    TokenInfo {
        supply: decoder.int(),
        minted_by: decoder.address(),
        owner: decoder.address(),
        created: decoder.int(),
        updated: decoder.int(),
        description: decoder.string(),
        user_defined: decoder.string(),
    }
}

fn encodeTokenInfo(token: &TokenInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(token.supply);
    encoder.address(&token.minted_by);
    encoder.address(&token.owner);
    encoder.int(token.created);
    encoder.int(token.updated);
    encoder.string(&token.description);
    encoder.string(&token.user_defined);
    encoder.data()
}
