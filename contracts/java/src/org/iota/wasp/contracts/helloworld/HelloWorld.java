// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.helloworld;

import org.iota.wasp.contracts.helloworld.lib.*;
import org.iota.wasp.wasmlib.context.*;

public class HelloWorld {
	public static void FuncHelloWorld(ScFuncContext ctx, FuncHelloWorldParams params) {
		ctx.Log("Hello, world!");
	}

	public static void ViewGetHelloWorld(ScViewContext ctx, ViewGetHelloWorldParams params) {
		//TODO
	}
}
