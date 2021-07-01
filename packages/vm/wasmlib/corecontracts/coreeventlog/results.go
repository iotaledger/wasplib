// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreeventlog

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableGetNumRecordsResults struct {
	id int32
}

func (s ImmutableGetNumRecordsResults) NumRecords() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ResultNumRecords.KeyID())
}

type MutableGetNumRecordsResults struct {
	id int32
}

func (s MutableGetNumRecordsResults) NumRecords() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ResultNumRecords.KeyID())
}

type ArrayOfImmutableBytes struct {
	objID int32
}

func (a ArrayOfImmutableBytes) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableBytes) GetBytes(index int32) wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(a.objID, wasmlib.Key32(index))
}

type ImmutableGetRecordsResults struct {
	id int32
}

func (s ImmutableGetRecordsResults) Records() ArrayOfImmutableBytes {
	arrID := wasmlib.GetObjectID(s.id, ResultRecords.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableBytes{objID: arrID}
}

type ArrayOfMutableBytes struct {
	objID int32
}

func (a ArrayOfMutableBytes) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableBytes) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableBytes) GetBytes(index int32) wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(a.objID, wasmlib.Key32(index))
}

type MutableGetRecordsResults struct {
	id int32
}

func (s MutableGetRecordsResults) Records() ArrayOfMutableBytes {
	arrID := wasmlib.GetObjectID(s.id, ResultRecords.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableBytes{objID: arrID}
}
