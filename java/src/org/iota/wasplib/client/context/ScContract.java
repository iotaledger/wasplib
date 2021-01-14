// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;

public class ScContract {
	ScContract() {
	}

	public ScAddress Chain() {
		return Host.root.GetAddress(Key.Chain).Value();
	}

	public ScAgent ChainOwner() {
		return Host.root.GetAgent(Key.ChainOwner).Value();
	}

	public ScAgent Creator() {
		return Host.root.GetAgent(Key.Creator).Value();
	}

	public ScAgent Id() {
		return Host.root.GetAgent(Key.Id).Value();
	}
}
