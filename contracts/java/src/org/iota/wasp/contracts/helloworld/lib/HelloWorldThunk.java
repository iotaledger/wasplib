// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.helloworld.lib;

import org.iota.wasp.contracts.helloworld.HelloWorld;
import org.iota.wasp.wasmlib.context.ScFuncContext;
import org.iota.wasp.wasmlib.context.ScViewContext;
import org.iota.wasp.wasmlib.exports.ScExports;

public class HelloWorldThunk {
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddFunc("helloWorld", HelloWorldThunk::funcHelloWorldThunk);
		exports.AddView("getHelloWorld", HelloWorldThunk::viewGetHelloWorldThunk);
	}

	private static void funcHelloWorldThunk(ScFuncContext ctx) {
		FuncHelloWorldParams params = new FuncHelloWorldParams();
		HelloWorld.FuncHelloWorld(ctx, params);
	}

	private static void viewGetHelloWorldThunk(ScViewContext ctx) {
		ViewGetHelloWorldParams params = new ViewGetHelloWorldParams();
		HelloWorld.ViewGetHelloWorld(ctx, params);
	}
}
