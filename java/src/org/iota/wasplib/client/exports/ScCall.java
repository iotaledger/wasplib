package org.iota.wasplib.client.exports;

import org.iota.wasplib.client.context.ScCallContext;

public interface ScCall {
	void call(ScCallContext sc);
}
