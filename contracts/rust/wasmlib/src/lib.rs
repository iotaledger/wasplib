// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

pub use bytes::*;
pub use context::*;
pub use corecontracts::*;
pub use exports::ScExports;
pub use hashtypes::*;
pub use immutable::*;
pub use keys::*;
pub use mutable::*;

mod bytes;
mod context;
mod corecontracts;
mod exports;
mod hashtypes;
pub mod host;
mod immutable;
pub mod keys;
mod mutable;

mod utils;

// When the `wee_alloc` feature is enabled,
// use `wee_alloc` as the global allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;
