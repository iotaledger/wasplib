// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use crate::context::*;
use crate::keys::*;
use crate::mutable::*;

// note that we do not use the Wasm export symbol table on purpose
// because Wasm does not allow us to determine whether the symbols
// are view or func, or even if their interface handling is correct
// in fact, there are only 2 symbols the host will look for in the
// export table:
// on_load (defined by the SC code) and
// on_call_entrypoint (defined here as part of wasmlib)

static mut FUNCS: Vec<fn(&ScFuncContext)> = vec![];
static mut VIEWS: Vec<fn(&ScViewContext)> = vec![];

#[no_mangle]
// general entrypoint for the host to call any SC function
// the host will pass the index of the entrypoint that was
// defined by the on_load SC initializer function
fn on_call_entrypoint(index: i32) {
    unsafe {
        if (index & 0x8000) != 0 {
            // immutable view function, invoke with view context
            VIEWS[(index & 0x7fff) as usize](&ScViewContext {});
            return;
        }

        // mutable full function, invoke with func context
        FUNCS[index as usize](&ScFuncContext {});
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// context for on_load function to be able to tell host which
// views and funcs are available as entry points to the SC
pub struct ScExports {
    exports: ScMutableStringArray,
}

impl ScExports {
    // constructs the symbol export context for the on_load function
    pub fn new() -> ScExports {
        let exports = ROOT.get_string_array(&KEY_EXPORTS);
        // tell host what values our special predefined key is
        // this helps detect versioning problems between host
        // and client versions of wasmlib
        exports.get_string(KEY_ZZZZZZZ.0).set_value("Rust:KEY_ZZZZZZZ");
        ScExports { exports: exports }
    }

    // defines the external name of a mutable full function
    // and the entry point function associated with that name
    pub fn add_func(&self, name: &str, f: fn(&ScFuncContext)) {
        unsafe {
            let index = FUNCS.len() as i32;
            FUNCS.push(f);
            self.exports.get_string(index).set_value(name);
        }
    }

    // defines the external name of an immutable view function
    // and the entry point function associated with that name
    pub fn add_view(&self, name: &str, f: fn(&ScViewContext)) {
        unsafe {
            let index = VIEWS.len() as i32;
            VIEWS.push(f);
            self.exports.get_string(index | 0x8000).set_value(name);
        }
    }
}

