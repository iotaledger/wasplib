// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.erc20.lib;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;

public class Consts {
    public static final String ScName = "erc20";
    public static final String ScDescription = "ERC-20 PoC for IOTA Smart Contracts";
    public static final ScHname HScName = new ScHname(0x200e3733);

    public static final Key ParamAccount = new Key("ac");
    public static final Key ParamAmount = new Key("am");
    public static final Key ParamCreator = new Key("c");
    public static final Key ParamDelegation = new Key("d");
    public static final Key ParamRecipient = new Key("r");
    public static final Key ParamSupply = new Key("s");

    public static final Key VarBalances = new Key("b");
    public static final Key VarSupply = new Key("s");

    public static final String FuncApprove = "approve";
    public static final String FuncInit = "init";
    public static final String FuncTransfer = "transfer";
    public static final String FuncTransferFrom = "transferFrom";
    public static final String ViewAllowance = "allowance";
    public static final String ViewBalanceOf = "balanceOf";
    public static final String ViewTotalSupply = "totalSupply";

    public static final ScHname HFuncApprove = new ScHname(0xa0661268);
    public static final ScHname HFuncInit = new ScHname(0x1f44d644);
    public static final ScHname HFuncTransfer = new ScHname(0xa15da184);
    public static final ScHname HFuncTransferFrom = new ScHname(0xd5e0a602);
    public static final ScHname HViewAllowance = new ScHname(0x5e16006a);
    public static final ScHname HViewBalanceOf = new ScHname(0x67ef8df4);
    public static final ScHname HViewTotalSupply = new ScHname(0x9505e6ca);
}
