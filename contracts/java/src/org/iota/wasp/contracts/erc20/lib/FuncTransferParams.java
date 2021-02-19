// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.erc20.lib;

import org.iota.wasp.wasmlib.immutable.ScImmutableAgentId;
import org.iota.wasp.wasmlib.immutable.ScImmutableInt;

public class FuncTransferParams {
	ScImmutableAgentId Account; // target account
	ScImmutableInt Amount;  // amount of tokens to transfer
}
