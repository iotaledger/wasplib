// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairauction.lib;

import org.iota.wasp.wasmlib.immutable.*;

//@formatter:off
public class FuncStartAuctionParams {
    public ScImmutableColor   Color;       // color of the tokens being auctioned
    public ScImmutableString  Description; // description of the tokens being auctioned
    public ScImmutableInt64   Duration;    // duration of auction in minutes
    public ScImmutableInt64   MinimumBid;  // minimum required amount for any bid
}
//@formatter:on