// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.helloworld.lib;

import org.iota.wasp.wasmlib.hashtypes.ScHname;
import org.iota.wasp.wasmlib.keys.Key;

public class Consts {
	public static final String ScName = "helloworld";
	public static final String ScDescription = "The ubiquitous hello world demo";
	public static final ScHname HScName = new ScHname(0x0683223c);

	public static final Key VarHelloWorld = new Key("helloWorld");

	public static final String FuncHelloWorld = "helloWorld";
	public static final String ViewGetHelloWorld = "getHelloWorld";

	public static final ScHname HFuncHelloWorld = new ScHname(0x9d042e65);
	public static final ScHname HViewGetHelloWorld = new ScHname(0x210439ce);
}
