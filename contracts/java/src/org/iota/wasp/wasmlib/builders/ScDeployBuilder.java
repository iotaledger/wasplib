// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScDeployBuilder {
	ScMutableMap deploy;

	public ScDeployBuilder(String name, String description) {
		ScMutableMapArray deploys = Host.root.GetMapArray(Key.Deploys);
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
