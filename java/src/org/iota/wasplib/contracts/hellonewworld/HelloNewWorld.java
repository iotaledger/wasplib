// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.hellonewworld;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableInt;

public class HelloNewWorld {
	private static final Key keyCounter = new Key("counter");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("hello", HelloNewWorld::hello);
		exports.AddView("getCounter", HelloNewWorld::getCounter);
	}

	// Function hello implement smart contract entry point "hello".
	// Function hello logs the message "Hello, new world!" with the counter and increments the counter
	public static void hello(ScCallContext ctx) {
		ScMutableInt counter = ctx.State().GetInt(keyCounter);
		String msg = "Hello, new world! #" + counter;
		ctx.Log(msg);  // todo info and debug levels, not events!
		counter.SetValue(counter.Value() + 1);
	}

	// Function get_counter implements smart contract VIEW entry point "getCounter".
	// It return counter value in the result dictionary with the key "counter"
	public static void getCounter(ScViewContext ctx) {
		long counter = ctx.State().GetInt(keyCounter).Value();
		ctx.Results().GetInt(keyCounter).SetValue(counter);
	}
}
