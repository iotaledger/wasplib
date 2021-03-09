// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.tokenregistry.lib;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;

public class Consts {
    public static final String ScName = "tokenregistry";
    public static final ScHname HScName = new ScHname(0xe1ba0c78);

    public static final Key ParamColor = new Key("color");
    public static final Key ParamDescription = new Key("description");
    public static final Key ParamUserDefined = new Key("userDefined");

    public static final Key VarColorList = new Key("colorList");
    public static final Key VarRegistry = new Key("registry");

    public static final String FuncMintSupply = "mintSupply";
    public static final String FuncTransferOwnership = "transferOwnership";
    public static final String FuncUpdateMetadata = "updateMetadata";
    public static final String ViewGetInfo = "getInfo";

    public static final ScHname HFuncMintSupply = new ScHname(0x564349a7);
    public static final ScHname HFuncTransferOwnership = new ScHname(0xbb9eb5af);
    public static final ScHname HFuncUpdateMetadata = new ScHname(0xa26b23b6);
    public static final ScHname HViewGetInfo = new ScHname(0xcfedba5f);
}