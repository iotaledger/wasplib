// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.testwasmlib.lib;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;

public class Consts {
    public static final String ScName = "testwasmlib";
    public static final String ScDescription = "Exercise all aspects of WasmLib";
    public static final ScHname HScName = new ScHname(0x89703a45);

    public static final Key ParamAddress = new Key("address");
    public static final Key ParamAgentId = new Key("agentId");
    public static final Key ParamBytes = new Key("bytes");
    public static final Key ParamChainId = new Key("chainId");
    public static final Key ParamColor = new Key("color");
    public static final Key ParamHash = new Key("hash");
    public static final Key ParamHname = new Key("hname");
    public static final Key ParamInt64 = new Key("int64");
    public static final Key ParamRequestId = new Key("requestId");
    public static final Key ParamString = new Key("string");

    public static final String FuncParamTypes = "paramTypes";

    public static final ScHname HFuncParamTypes = new ScHname(0x6921c4cd);
}
