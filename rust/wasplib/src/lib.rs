// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

pub mod client;
mod utils;

// When the `wee_alloc` feature is enabled,
// use `wee_alloc` as the global allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;
