// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

pub use bytes::*;
pub use context::*;
pub use contract::*;
pub use exports::ScExports;
pub use hashtypes::*;
pub use immutable::*;
pub use keys::*;
pub use mutable::*;

mod bytes;
mod context;
mod contract;
pub mod corecontracts;
mod exports;
mod hashtypes;
pub mod host;
mod immutable;
pub mod keys;
mod mutable;

// When the `wee_alloc` feature is enabled,
// use `wee_alloc` as the global allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// When the `console_error_panic_hook` feature is enabled, we can call the
// `set_panic_hook` function at least once during initialization, and then
// we will get better error messages if our code ever panics.
//
// For more details see
// https://github.com/rustwasm/console_error_panic_hook#readme
pub fn set_panic_hook() {
    #[cfg(feature = "console_error_panic_hook")]
        console_error_panic_hook::set_once();
}
