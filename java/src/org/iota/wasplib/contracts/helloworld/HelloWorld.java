// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.helloworld;

import org.iota.wasplib.client.context.ScFuncContext;
import org.iota.wasplib.client.exports.ScExports;

public class HelloWorld {
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddFunc("hello_world", HelloWorld::helloWorld);
	}

	public static void helloWorld(ScFuncContext sc) {
		sc.Log("Hello, world!");
	}
}
