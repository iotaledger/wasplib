// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairauction.lib;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;

public class Consts {
    public static final String ScName = "fairauction";
    public static final ScHname HScName = new ScHname(0x1b5c43b1);

    public static final Key ParamColor = new Key("color");
    public static final Key ParamDescription = new Key("description");
    public static final Key ParamDuration = new Key("duration");
    public static final Key ParamMinimumBid = new Key("minimumBid");
    public static final Key ParamOwnerMargin = new Key("ownerMargin");

    public static final Key VarAuctions = new Key("auctions");
    public static final Key VarBidderList = new Key("bidderList");
    public static final Key VarBidders = new Key("bidders");
    public static final Key VarColor = new Key("color");
    public static final Key VarCreator = new Key("creator");
    public static final Key VarDeposit = new Key("deposit");
    public static final Key VarDescription = new Key("description");
    public static final Key VarDuration = new Key("duration");
    public static final Key VarHighestBid = new Key("highestBid");
    public static final Key VarHighestBidder = new Key("highestBidder");
    public static final Key VarInfo = new Key("info");
    public static final Key VarMinimumBid = new Key("minimumBid");
    public static final Key VarNumTokens = new Key("numTokens");
    public static final Key VarOwnerMargin = new Key("ownerMargin");
    public static final Key VarWhenStarted = new Key("whenStarted");

    public static final String FuncFinalizeAuction = "finalizeAuction";
    public static final String FuncPlaceBid = "placeBid";
    public static final String FuncSetOwnerMargin = "setOwnerMargin";
    public static final String FuncStartAuction = "startAuction";
    public static final String ViewGetInfo = "getInfo";

    public static final ScHname HFuncFinalizeAuction = new ScHname(0x8d534ddc);
    public static final ScHname HFuncPlaceBid = new ScHname(0x9bd72fa9);
    public static final ScHname HFuncSetOwnerMargin = new ScHname(0x1774461a);
    public static final ScHname HFuncStartAuction = new ScHname(0xd5b7bacb);
    public static final ScHname HViewGetInfo = new ScHname(0xcfedba5f);
}
