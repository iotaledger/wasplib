// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "blocklog";
pub const SC_DESCRIPTION: &str = "Core block log contract";
pub const HSC_NAME:       ScHname = ScHname(0xf538ef2b);

pub const PARAM_BLOCK_INDEX: &str = "n";
pub const PARAM_REQUEST_ID:  &str = "u";

pub const RESULT_BLOCK_INDEX:       &str = "n";
pub const RESULT_BLOCK_INFO:        &str = "i";
pub const RESULT_REQUEST_INDEX:     &str = "r";
pub const RESULT_REQUEST_PROCESSED: &str = "p";
pub const RESULT_REQUEST_RECORD:    &str = "d";

pub const VIEW_GET_BLOCK_INFO:                    &str = "getBlockInfo";
pub const VIEW_GET_LATEST_BLOCK_INFO:             &str = "getLatestBlockInfo";
pub const VIEW_GET_REQUEST_I_DS_FOR_BLOCK:        &str = "getRequestIDsForBlock";
pub const VIEW_GET_REQUEST_LOG_RECORD:            &str = "getRequestLogRecord";
pub const VIEW_GET_REQUEST_LOG_RECORDS_FOR_BLOCK: &str = "getRequestLogRecordsForBlock";
pub const VIEW_IS_REQUEST_PROCESSED:              &str = "isRequestProcessed";

pub const HVIEW_GET_BLOCK_INFO:                    ScHname = ScHname(0xbe89f9b3);
pub const HVIEW_GET_LATEST_BLOCK_INFO:             ScHname = ScHname(0x084a1760);
pub const HVIEW_GET_REQUEST_I_DS_FOR_BLOCK:        ScHname = ScHname(0x5a20327a);
pub const HVIEW_GET_REQUEST_LOG_RECORD:            ScHname = ScHname(0x31e07e48);
pub const HVIEW_GET_REQUEST_LOG_RECORDS_FOR_BLOCK: ScHname = ScHname(0x7210e621);
pub const HVIEW_IS_REQUEST_PROCESSED:              ScHname = ScHname(0xd57d50a9);

//@formatter:on
