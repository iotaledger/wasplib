// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScPostedRequest {
	ScMutableMap request;

	ScPostedRequest(ScMutableMap request) {
		this.request = request;
	}

	public void Code(long code) {
		request.GetInt("code").SetValue(code);
	}

	public void Contract(ScAgent contract) {
		request.GetAgent("contract").SetValue(contract);
	}

	public void Delay(long delay) {
		request.GetInt("delay").SetValue(delay);
	}

	public void Function(String function) {
		request.GetString("function").SetValue(function);
	}

	public ScMutableMap Params() {
		return request.GetMap("params");
	}
}
