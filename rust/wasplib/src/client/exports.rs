// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::context::*;
use super::keys::*;
use super::mutable::*;

static mut CALLS: Vec<fn(&ScCallContext)> = vec![];
static mut VIEWS: Vec<fn(&ScViewContext)> = vec![];

#[no_mangle]
fn on_call_entrypoint(index: i32) {
    unsafe {
        if (index & 0x8000) != 0 {
            VIEWS[(index & 0x7fff) as usize](&ScViewContext {});
            return;
        }

        CALLS[index as usize](&ScCallContext {});
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScExports {
    exports: ScMutableStringArray,
}

impl ScExports {
    pub fn new() -> ScExports {
        let exports = ROOT.get_string_array(&KEY_EXPORTS);
        // tell host what our highest predefined key is
        // this helps detect missing or extra keys
        exports.get_string(KEY_ZZZZZZZ.0).set_value("Rust:KEY_ZZZZZZZ");
        ScExports { exports: exports }
    }

    pub fn add_call(&self, name: &str, f: fn(&ScCallContext)) {
        unsafe {
            let index = CALLS.len() as i32;
            CALLS.push(f);
            self.exports.get_string(index).set_value(name);
        }
    }

    pub fn add_view(&self, name: &str, f: fn(&ScViewContext)) {
        unsafe {
            let index = VIEWS.len() as i32;
            VIEWS.push(f);
            self.exports.get_string(index | 0x8000).set_value(name);
        }
    }

    pub fn nothing(ctx: &ScCallContext) {
        ctx.log("Doing nothing as requested. Oh, wait...");
    }
}

