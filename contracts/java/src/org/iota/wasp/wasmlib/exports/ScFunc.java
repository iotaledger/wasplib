package org.iota.wasp.wasmlib.exports;

import org.iota.wasp.wasmlib.context.*;

public interface ScFunc {
    void call(ScFuncContext ctx);
}
