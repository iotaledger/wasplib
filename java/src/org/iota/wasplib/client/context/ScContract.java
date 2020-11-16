// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScContract {
	ScImmutableMap contract;

	ScContract(ScImmutableMap contract) {
		this.contract = contract;
	}

	public ScColor Color() {
		return contract.GetColor("color").Value();
	}

	public String Description() {
		return contract.GetString("description").Value();
	}

	public ScAgent Id() {
		return contract.GetAgent("id").Value();
	}

	public String Name() {
		return contract.GetString("name").Value();
	}

	public ScAgent Owner() {
		return contract.GetAgent("owner").Value();
	}
}
