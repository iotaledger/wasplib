// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type MapAgentIdToMutableAllowancesForAgent struct {
	objId int32
}

func (m MapAgentIdToMutableAllowancesForAgent) Clear() {
	wasmlib.Clear(m.objId)
}

func (m MapAgentIdToMutableAllowancesForAgent) GetAllowancesForAgent(key wasmlib.ScAgentId) MutableAllowancesForAgent {
	subId := wasmlib.GetObjectId(m.objId, key.KeyId(), wasmlib.TYPE_MAP)
	return MutableAllowancesForAgent{objId: subId}
}

type MutableErc20State struct {
	id int32
}

func (s MutableErc20State) AllAllowances() MapAgentIdToMutableAllowancesForAgent {
	mapId := wasmlib.GetObjectId(s.id, idxMap[IdxStateAllAllowances], wasmlib.TYPE_MAP)
	return MapAgentIdToMutableAllowancesForAgent{objId: mapId}
}

func (s MutableErc20State) Balances() MapAgentIdToMutableInt64 {
	mapId := wasmlib.GetObjectId(s.id, idxMap[IdxStateBalances], wasmlib.TYPE_MAP)
	return MapAgentIdToMutableInt64{objId: mapId}
}

func (s MutableErc20State) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxStateSupply])
}

type MapAgentIdToImmutableAllowancesForAgent struct {
	objId int32
}

func (m MapAgentIdToImmutableAllowancesForAgent) GetAllowancesForAgent(key wasmlib.ScAgentId) ImmutableAllowancesForAgent {
	subId := wasmlib.GetObjectId(m.objId, key.KeyId(), wasmlib.TYPE_MAP)
	return ImmutableAllowancesForAgent{objId: subId}
}

type ImmutableErc20State struct {
	id int32
}

func (s ImmutableErc20State) AllAllowances() MapAgentIdToImmutableAllowancesForAgent {
	mapId := wasmlib.GetObjectId(s.id, idxMap[IdxStateAllAllowances], wasmlib.TYPE_MAP)
	return MapAgentIdToImmutableAllowancesForAgent{objId: mapId}
}

func (s ImmutableErc20State) Balances() MapAgentIdToImmutableInt64 {
	mapId := wasmlib.GetObjectId(s.id, idxMap[IdxStateBalances], wasmlib.TYPE_MAP)
	return MapAgentIdToImmutableInt64{objId: mapId}
}

func (s ImmutableErc20State) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxStateSupply])
}
