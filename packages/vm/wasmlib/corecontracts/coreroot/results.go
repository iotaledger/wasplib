// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreroot

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableFindContractResults struct {
	id int32
}

func (s ImmutableFindContractResults) ContractFound() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, ResultContractFound.KeyID())
}

func (s ImmutableFindContractResults) ContractRecData() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, ResultContractRecData.KeyID())
}

type MutableFindContractResults struct {
	id int32
}

func (s MutableFindContractResults) ContractFound() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, ResultContractFound.KeyID())
}

func (s MutableFindContractResults) ContractRecData() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, ResultContractRecData.KeyID())
}

type MapHnameToImmutableBytes struct {
	objID int32
}

func (m MapHnameToImmutableBytes) GetBytes(key wasmlib.ScHname) wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(m.objID, key.KeyID())
}

type ImmutableGetContractRecordsResults struct {
	id int32
}

func (s ImmutableGetContractRecordsResults) ContractRegistry() MapHnameToImmutableBytes {
	mapID := wasmlib.GetObjectID(s.id, ResultContractRegistry.KeyID(), wasmlib.TYPE_MAP)
	return MapHnameToImmutableBytes{objID: mapID}
}

type MapHnameToMutableBytes struct {
	objID int32
}

func (m MapHnameToMutableBytes) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHnameToMutableBytes) GetBytes(key wasmlib.ScHname) wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(m.objID, key.KeyID())
}

type MutableGetContractRecordsResults struct {
	id int32
}

func (s MutableGetContractRecordsResults) ContractRegistry() MapHnameToMutableBytes {
	mapID := wasmlib.GetObjectID(s.id, ResultContractRegistry.KeyID(), wasmlib.TYPE_MAP)
	return MapHnameToMutableBytes{objID: mapID}
}
