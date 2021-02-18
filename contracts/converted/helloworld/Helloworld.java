// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.helloworld;

public class Helloworld {

public static void funcHelloWorld(ScFuncContext ctx, FuncHelloWorldParams params) {
    ctx.Log("Hello, world!");
}

public static void viewGetHelloWorld(ScViewContext ctx, ViewGetHelloWorldParams params) {
    ctx.Log("Get Hello world!");
    ctx.Results().GetString(VarHelloWorld).SetValue("Hello, world!");
}
}
