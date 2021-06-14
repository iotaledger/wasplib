// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "blob";
pub const SC_DESCRIPTION: &str = "Core blob contract";
pub const HSC_NAME:       ScHname = ScHname(0xfd91bc63);

pub const PARAM_FIELD: &str = "field";
pub const PARAM_HASH:  &str = "hash";

pub const RESULT_BYTES: &str = "bytes";
pub const RESULT_HASH:  &str = "hash";

pub const FUNC_STORE_BLOB:     &str = "storeBlob";
pub const VIEW_GET_BLOB_FIELD: &str = "getBlobField";
pub const VIEW_GET_BLOB_INFO:  &str = "getBlobInfo";
pub const VIEW_LIST_BLOBS:     &str = "listBlobs";

pub const HFUNC_STORE_BLOB:     ScHname = ScHname(0xddd4c281);
pub const HVIEW_GET_BLOB_FIELD: ScHname = ScHname(0x1f448130);
pub const HVIEW_GET_BLOB_INFO:  ScHname = ScHname(0xfde4ab46);
pub const HVIEW_LIST_BLOBS:     ScHname = ScHname(0x62ca7990);

//@formatter:on
