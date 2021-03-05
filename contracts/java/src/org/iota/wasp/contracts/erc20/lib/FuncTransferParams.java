// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.erc20.lib;

import org.iota.wasp.wasmlib.immutable.*;

public class FuncTransferParams {
    public ScImmutableAgentId Account; // target account
    public ScImmutableInt64 Amount;  // amount of tokens to transfer
}
