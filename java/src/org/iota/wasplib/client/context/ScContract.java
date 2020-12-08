// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;

public class ScContract {
	ScImmutableMap contract;

	ScContract(ScImmutableMap contract) {
		this.contract = contract;
	}

	public ScAddress Chain() {
		return contract.GetAddress(Key.Chain).Value();
	}

	public String Description() {
		return contract.GetString(Key.Description).Value();
	}

	public ScAgent Id() {
		return contract.GetAgent(Key.Id).Value();
	}

	public String Name() {
		return contract.GetString(Key.Name).Value();
	}

	public ScAgent Owner() {
		return contract.GetAgent(Key.Owner).Value();
	}
}
