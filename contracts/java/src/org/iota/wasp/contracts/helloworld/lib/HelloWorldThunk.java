// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead


package org.iota.wasp.contracts.helloworld.lib;

import org.iota.wasp.contracts.helloworld.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class HelloWorldThunk {
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddFunc("helloWorld", HelloWorld::FuncHelloWorld);
		exports.AddView("getHelloWorld", HelloWorld::ViewGetHelloWorld);
	}
}
