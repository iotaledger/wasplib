#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::BytesDecoder;
use wasplib::client::BytesEncoder;
use wasplib::client::ScContext;

struct TokenInfo {
    supply: i64,
    minted_by: String,
    owner: String,
    created: i64,
    updated: i64,
    description: String,
    user_defined: String,
}

#[no_mangle]
pub fn mintSupply() {
    let ctx = ScContext::new();
    let request = ctx.request();
    let color = request.hash();
    let state = ctx.state();
    let registry = state.get_map("tr");
    if !registry.get_string(&color).value().is_empty() {
        ctx.log("TokenRegistry: Color already exists");
        return;
    }
    let reqParams = request.params();
    let mut token = TokenInfo {
        supply: request.balance(&color),
        minted_by: request.address(),
        owner: request.address(),
        created: request.timestamp(),
        updated: request.timestamp(),
        description: reqParams.get_string("dscr").value(),
        user_defined: reqParams.get_string("ud").value(),
    };
    if token.supply <= 0 {
        ctx.log("TokenRegistry: Insufficient supply");
        return;
    }
    if token.description.is_empty() {
        token.description += "no dscr";
    }
    let data = encodeTokenInfo(&token);
    registry.get_bytes(&color).set_value(&data);
    let colors = state.get_string("lc");
    let mut list = colors.value();
    if !list.is_empty() {
        list += ",";
    }
    list += &color;
    colors.set_value(&list)
}

#[no_mangle]
pub fn updateMetadata() {
    //let ctx = ScContext::new();
    //TODO
}

#[no_mangle]
pub fn transferOwnership() {
    //let ctx = ScContext::new();
    //TODO
}

fn decodeTokenInfo(data: &[u8]) -> TokenInfo {
    let mut decoder = BytesDecoder::new(data);
    TokenInfo {
        supply: decoder.int(),
        minted_by: decoder.string(),
        owner: decoder.string(),
        created: decoder.int(),
        updated: decoder.int(),
        description: decoder.string(),
        user_defined: decoder.string(),
    }
}

fn encodeTokenInfo(data: &TokenInfo) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(data.supply);
    encoder.string(&data.minted_by);
    encoder.string(&data.owner);
    encoder.int(data.created);
    encoder.int(data.updated);
    encoder.string(&data.description);
    encoder.string(&data.user_defined);
    encoder.data()
}
