package org.iota.wasp.wasmlib.exports;

import org.iota.wasp.wasmlib.context.ScFuncContext;

public interface ScFunc {
	void call(ScFuncContext ctx);
}
