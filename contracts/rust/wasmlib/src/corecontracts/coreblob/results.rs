// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coreblob::*;
use crate::host::*;

#[derive(Clone, Copy)]
pub struct ImmutableFuncStoreBlobResults {
    pub(crate) id: i32,
}

impl ImmutableFuncStoreBlobResults {
    pub fn hash(&self) -> ScImmutableHash {
        ScImmutableHash::new(self.id, RESULT_HASH.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableFuncStoreBlobResults {
    pub(crate) id: i32,
}

impl MutableFuncStoreBlobResults {
    pub fn hash(&self) -> ScMutableHash {
        ScMutableHash::new(self.id, RESULT_HASH.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableViewGetBlobFieldResults {
    pub(crate) id: i32,
}

impl ImmutableViewGetBlobFieldResults {
    pub fn bytes(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.id, RESULT_BYTES.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableViewGetBlobFieldResults {
    pub(crate) id: i32,
}

impl MutableViewGetBlobFieldResults {
    pub fn bytes(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.id, RESULT_BYTES.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableViewGetBlobInfoResults {
    pub(crate) id: i32,
}

#[derive(Clone, Copy)]
pub struct MutableViewGetBlobInfoResults {
    pub(crate) id: i32,
}

#[derive(Clone, Copy)]
pub struct ImmutableViewListBlobsResults {
    pub(crate) id: i32,
}

#[derive(Clone, Copy)]
pub struct MutableViewListBlobsResults {
    pub(crate) id: i32,
}
