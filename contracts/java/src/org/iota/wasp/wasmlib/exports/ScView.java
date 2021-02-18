package org.iota.wasp.wasmlib.exports;

import org.iota.wasp.wasmlib.context.ScViewContext;

public interface ScView {
	void call(ScViewContext ctx);
}
