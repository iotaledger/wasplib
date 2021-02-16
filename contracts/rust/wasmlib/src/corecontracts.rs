// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::hashtypes::*;

pub const CORE_ACCOUNTS: ScHname = ScHname(0x3c4b5e02);
pub const CORE_ACCOUNTS_FUNC_DEPOSIT: ScHname = ScHname(0xbdc9102d);
pub const CORE_ACCOUNTS_FUNC_WITHDRAW_TO_ADDRESS: ScHname = ScHname(0x26608cb5);
pub const CORE_ACCOUNTS_FUNC_WITHDRAW_TO_CHAIN: ScHname = ScHname(0x437bc026);
pub const CORE_ACCOUNTS_VIEW_ACCOUNTS: ScHname = ScHname(0x3c4b5e02);
pub const CORE_ACCOUNTS_VIEW_BALANCE: ScHname = ScHname(0x84168cb4);
pub const CORE_ACCOUNTS_VIEW_TOTAL_ASSETS: ScHname = ScHname(0xfab0f8d2);

pub const CORE_ACCOUNTS_PARAM_AGENT_ID: &str = "a";

pub const CORE_BLOB: ScHname = ScHname(0xfd91bc63);
pub const CORE_BLOB_FUNC_STORE_BLOB: ScHname = ScHname(0xddd4c281);
pub const CORE_BLOB_VIEW_GET_BLOB_FIELD: ScHname = ScHname(0x1f448130);
pub const CORE_BLOB_VIEW_GET_BLOB_INFO: ScHname = ScHname(0xfde4ab46);
pub const CORE_BLOB_VIEW_LIST_BLOBS: ScHname = ScHname(0x62ca7990);

pub const CORE_BLOB_PARAM_FIELD: &str = "field";
pub const CORE_BLOB_PARAM_HASH: &str = "hash";

pub const CORE_EVENTLOG: ScHname = ScHname(0x661aa7d8);
pub const CORE_EVENTLOG_VIEW_GET_NUM_RECORDS: ScHname = ScHname(0x2f4b4a8c);
pub const CORE_EVENTLOG_VIEW_GET_RECORDS: ScHname = ScHname(0xd01a8085);

pub const CORE_EVENTLOG_PARAM_CONTRACT_HNAME: &str = "contractHname";
pub const CORE_EVENTLOG_PARAM_FROM_TS: &str = "fromTs";
pub const CORE_EVENTLOG_PARAM_MAX_LAST_RECORDS: &str = "maxLastRecords";
pub const CORE_EVENTLOG_PARAM_TO_TS: &str = "toTs";

pub const CORE_ROOT: ScHname = ScHname(0xcebf5908);
pub const CORE_ROOT_FUNC_CLAIM_CHAIN_OWNERSHIP: ScHname = ScHname(0x03ff0fc0);
pub const CORE_ROOT_FUNC_DELEGATE_CHAIN_OWNERSHIP: ScHname = ScHname(0x93ecb6ad);
pub const CORE_ROOT_FUNC_DEPLOY_CONTRACT: ScHname = ScHname(0x28232c27);
pub const CORE_ROOT_FUNC_GRANT_DEPLOY_PERMISSION: ScHname = ScHname(0xf440263a);
pub const CORE_ROOT_FUNC_REVOKE_DEPLOY_PERMISSION: ScHname = ScHname(0x850744f1);
pub const CORE_ROOT_FUNC_SET_CONTRACT_FEE: ScHname = ScHname(0x8421a42b);
pub const CORE_ROOT_FUNC_SET_DEFAULT_FEE: ScHname = ScHname(0x3310ecd0);
pub const CORE_ROOT_VIEW_FIND_CONTRACT: ScHname = ScHname(0xc145ca00);
pub const CORE_ROOT_VIEW_GET_CHAIN_INFO: ScHname = ScHname(0x434477e2);
pub const CORE_ROOT_VIEW_GET_FEE_INFO: ScHname = ScHname(0x9fe54b48);

pub const CORE_ROOT_PARAM_CHAIN_OWNER: &str = "$$owner$$";
pub const CORE_ROOT_PARAM_DEPLOYER: &str = "$$deployer$$";
pub const CORE_ROOT_PARAM_DESCRIPTION: &str = "$$description$$";
pub const CORE_ROOT_PARAM_HNAME: &str = "$$hname$$";
pub const CORE_ROOT_PARAM_NAME: &str = "$$name$$";
pub const CORE_ROOT_PARAM_OWNER_FEE: &str = "$$ownerfee$$";
pub const CORE_ROOT_PARAM_PROGRAM_HASH: &str = "$$proghash$$";
pub const CORE_ROOT_PARAM_VALIDATOR_FEE: &str = "$$validatorfee$$";
