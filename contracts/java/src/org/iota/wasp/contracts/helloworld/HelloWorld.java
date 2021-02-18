// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.helloworld;

import org.iota.wasp.wasmlib.context.ScFuncContext;
import org.iota.wasp.wasmlib.context.ScViewContext;
import org.iota.wasp.wasmlib.exports.ScExports;

public class HelloWorld {
	public static void FuncHelloWorld(ScFuncContext ctx) {
		ctx.Log("Hello, world!");
	}

	public static void ViewGetHelloWorld(ScViewContext ctx) {
		//TODO
	}
}
