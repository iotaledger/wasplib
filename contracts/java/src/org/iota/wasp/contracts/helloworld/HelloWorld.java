// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.helloworld;

import org.iota.wasp.contracts.helloworld.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class HelloWorld {

    public static void funcHelloWorld(ScFuncContext ctx, FuncHelloWorldParams params) {
        ctx.Log("Hello, world!");
    }

    public static void viewGetHelloWorld(ScViewContext ctx, ViewGetHelloWorldParams params) {
        ctx.Log("Get Hello world!");
        ctx.Results().GetString(Consts.VarHelloWorld).SetValue("Hello, world!");
    }
}
