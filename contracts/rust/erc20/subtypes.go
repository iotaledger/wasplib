// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableAllowancesForAgent = MapAgentIDToImmutableInt64

type MapAgentIDToImmutableInt64 struct {
	objID int32
}

func (m MapAgentIDToImmutableInt64) GetInt64(key wasmlib.ScAgentID) wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(m.objID, key.KeyID())
}

type MutableAllowancesForAgent = MapAgentIDToMutableInt64

type MapAgentIDToMutableInt64 struct {
	objID int32
}

func (m MapAgentIDToMutableInt64) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapAgentIDToMutableInt64) GetInt64(key wasmlib.ScAgentID) wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(m.objID, key.KeyID())
}
