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

type Erc20FuncState struct {
	stateId int32
}

func (s Erc20FuncState) AllAllowances() MapAgentIdToMutableAllowancesForAgent {
	mapId := wasmlib.GetObjectId(s.stateId, VarAllAllowances.KeyId(), wasmlib.TYPE_MAP)
	return MapAgentIdToMutableAllowancesForAgent{objId: mapId}
}

func (s Erc20FuncState) Balances() MapAgentIdToMutableInt64 {
	mapId := wasmlib.GetObjectId(s.stateId, VarBalances.KeyId(), wasmlib.TYPE_MAP)
	return MapAgentIdToMutableInt64{objId: mapId}
}

func (s Erc20FuncState) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.stateId, VarSupply.KeyId())
}

type MapAgentIdToImmutableAllowancesForAgent struct {
	objId int32
}

func (m MapAgentIdToImmutableAllowancesForAgent) GetAllowancesForAgent(key wasmlib.ScAgentId) ImmutableAllowancesForAgent {
	subId := wasmlib.GetObjectId(m.objId, key.KeyId(), wasmlib.TYPE_MAP)
	return ImmutableAllowancesForAgent{objId: subId}
}

type Erc20ViewState struct {
	stateId int32
}

func (s Erc20ViewState) AllAllowances() MapAgentIdToImmutableAllowancesForAgent {
	mapId := wasmlib.GetObjectId(s.stateId, VarAllAllowances.KeyId(), wasmlib.TYPE_MAP)
	return MapAgentIdToImmutableAllowancesForAgent{objId: mapId}
}

func (s Erc20ViewState) Balances() MapAgentIdToImmutableInt64 {
	mapId := wasmlib.GetObjectId(s.stateId, VarBalances.KeyId(), wasmlib.TYPE_MAP)
	return MapAgentIdToImmutableInt64{objId: mapId}
}

func (s Erc20ViewState) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.stateId, VarSupply.KeyId())
}
