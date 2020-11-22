// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScCallInfo {
	ScMutableMap call;

	ScCallInfo(ScMutableMap call) {
		this.call = call;
	}

	void Contract(String contract) {
		call.GetString("contract").SetValue(contract);
	}

	void Delay(long delay) {
		call.GetInt("delay").SetValue(delay);
	}

	void Function(String function) {
		call.GetString("function").SetValue(function);
	}

	public ScMutableMap Params() {
		return call.GetMap("params");
	}
}
