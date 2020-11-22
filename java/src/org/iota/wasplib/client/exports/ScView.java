package org.iota.wasplib.client.exports;

import org.iota.wasplib.client.context.ScViewContext;

public interface ScView {
	void call(ScViewContext sc);
}
