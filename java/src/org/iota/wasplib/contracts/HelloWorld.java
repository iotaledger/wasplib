// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;

public class HelloWorld {
	//export on_load
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("hello_world", HelloWorld::helloWorld);
	}

	public static void helloWorld(ScCallContext sc) {
		sc.Log("Hello, world!");
	}
}
