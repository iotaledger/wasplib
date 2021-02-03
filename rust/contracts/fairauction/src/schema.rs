// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;
use super::*;

pub const SC_NAME: &str = "fairauction";
pub const SC_HNAME: ScHname = ScHname(0x1b5c43b1);

pub const PARAM_COLOR: &str = "color";
pub const PARAM_DESCRIPTION: &str = "description";
pub const PARAM_DURATION: &str = "duration";
pub const PARAM_MINIMUM_BID: &str = "minimum";
pub const PARAM_OWNER_MARGIN: &str = "owner_margin";

pub const VAR_AUCTIONS: &str = "auctions";
pub const VAR_BIDDER_LIST: &str = "bidder_list";
pub const VAR_BIDDERS: &str = "bidders";
pub const VAR_COLOR: &str = "color";
pub const VAR_CREATOR: &str = "creator";
pub const VAR_DEPOSIT: &str = "deposit";
pub const VAR_DESCRIPTION: &str = "description";
pub const VAR_DURATION: &str = "duration";
pub const VAR_HIGHEST_BID: &str = "highest_bid";
pub const VAR_HIGHEST_BIDDER: &str = "highest_bidder";
pub const VAR_INFO: &str = "info";
pub const VAR_MINIMUM_BID: &str = "minimum";
pub const VAR_NUM_TOKENS: &str = "num_tokens";
pub const VAR_OWNER_MARGIN: &str = "owner_margin";
pub const VAR_WHEN_STARTED: &str = "when_started";

pub const FUNC_FINALIZE_AUCTION: &str = "finalize_auction";
pub const FUNC_PLACE_BID: &str = "place_bid";
pub const FUNC_SET_OWNER_MARGIN: &str = "set_owner_margin";
pub const FUNC_START_AUCTION: &str = "start_auction";
pub const VIEW_GET_INFO: &str = "get_info";

pub const HFUNC_FINALIZE_AUCTION: ScHname = ScHname(0xb427dd28);
pub const HFUNC_PLACE_BID: ScHname = ScHname(0xf2cc1c44);
pub const HFUNC_SET_OWNER_MARGIN: ScHname = ScHname(0x65402dca);
pub const HFUNC_START_AUCTION: ScHname = ScHname(0x7ee53d08);
pub const HVIEW_GET_INFO: ScHname = ScHname(0x2b9d8867);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_FINALIZE_AUCTION, func_finalize_auction);
    exports.add_call(FUNC_PLACE_BID, func_place_bid);
    exports.add_call(FUNC_SET_OWNER_MARGIN, func_set_owner_margin);
    exports.add_call(FUNC_START_AUCTION, func_start_auction);
    exports.add_view(VIEW_GET_INFO, view_get_info);
}
