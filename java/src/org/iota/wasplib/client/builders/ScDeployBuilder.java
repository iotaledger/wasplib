// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.builders;

import org.iota.wasplib.client.hashtypes.ScHash;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class ScDeployBuilder {
	ScMutableMap deploy;

	public ScDeployBuilder(String name, String description) {
		ScMutableMapArray deploys = ScRequestBuilder.root.GetMapArray(Key.Deploys);
		deploy = deploys.GetMap(deploys.Length());
		deploy.GetString(Key.Name).SetValue(name);
		deploy.GetString(Key.Description).SetValue(description);
	}

	public void Deploy(ScHash programHash) {
		deploy.GetHash(Key.Hash).SetValue(programHash);
	}

	public ScMutableMap Params() {
		return deploy.GetMap(Key.Params);
	}
}
