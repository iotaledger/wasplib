// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::hashtypes::*;

pub const CORE_ACCOUNTS: Hname = Hname(0x3c4b5e02);
pub const FUNC_BALANCE: Hname = Hname(0x84168cb4);
pub const FUNC_DEPOSIT: Hname = Hname(0xbdc9102d);
pub const FUNC_TOTAL_ASSETS: Hname = Hname(0xfab0f8d2);
pub const VIEW_ACCOUNTS: Hname = Hname(0x3c4b5e02);
pub const VIEW_ALLOW: Hname = Hname(0xfebef3ae);
pub const VIEW_WITHDRAW_TO_ADDRESS: Hname = Hname(0x26608cb5);
pub const VIEW_WITHDRAW_TO_CHAIN: Hname = Hname(0x437bc026);

pub const CORE_BLOB: Hname = Hname(0xfd91bc63);
pub const FUNC_STORE_BLOB: Hname = Hname(0xddd4c281);
pub const VIEW_GET_BLOB_FIELD: Hname = Hname(0x1f448130);
pub const VIEW_GET_BLOB_INFO: Hname = Hname(0xfde4ab46);
pub const VIEW_LIST_BLOBS: Hname = Hname(0x62ca7990);

pub const CORE_EVENTLOG: Hname = Hname(0x661aa7d8);
pub const VIEW_GET_LOG_RECORDS: Hname = Hname(0xc7407442);
pub const VIEW_GET_NUM_RECORDS: Hname = Hname(0x2f4b4a8c);

pub const CORE_ROOT: Hname = Hname(0xcebf5908);
pub const FUNC_CLAIM_CHAIN_OWNERSHIP: Hname = Hname(0x03ff0fc0);
pub const FUNC_DELEGATE_CHAIN_OWNERSHIP: Hname = Hname(0x93ecb6ad);
pub const FUNC_DEPLOY_CONTRACT: Hname = Hname(0x28232c27);
pub const FUNC_GRANT_DEPLOY: Hname = Hname(0xf440263a);
pub const FUNC_REVOKE_DEPLOY: Hname = Hname(0x850744f1);
pub const FUNC_SET_CONTRACT_FEE: Hname = Hname(0x8421a42b);
pub const FUNC_SET_DEFAULT_FEE: Hname = Hname(0x3310ecd0);
pub const VIEW_FIND_CONTRACT: Hname = Hname(0xc145ca00);
pub const VIEW_GET_CHAIN_INFO: Hname = Hname(0x434477e2);
pub const VIEW_GET_FEE_INFO: Hname = Hname(0x9fe54b48);
