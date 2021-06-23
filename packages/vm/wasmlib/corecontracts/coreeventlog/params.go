// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreeventlog

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableGetNumRecordsParams struct {
	id int32
}

func (s ImmutableGetNumRecordsParams) ContractHname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamContractHname.KeyId())
}

type MutableGetNumRecordsParams struct {
	id int32
}

func (s MutableGetNumRecordsParams) ContractHname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamContractHname.KeyId())
}

type ImmutableGetRecordsParams struct {
	id int32
}

func (s ImmutableGetRecordsParams) ContractHname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamContractHname.KeyId())
}

func (s ImmutableGetRecordsParams) FromTs() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamFromTs.KeyId())
}

func (s ImmutableGetRecordsParams) MaxLastRecords() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamMaxLastRecords.KeyId())
}

func (s ImmutableGetRecordsParams) ToTs() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamToTs.KeyId())
}

type MutableGetRecordsParams struct {
	id int32
}

func (s MutableGetRecordsParams) ContractHname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamContractHname.KeyId())
}

func (s MutableGetRecordsParams) FromTs() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamFromTs.KeyId())
}

func (s MutableGetRecordsParams) MaxLastRecords() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamMaxLastRecords.KeyId())
}

func (s MutableGetRecordsParams) ToTs() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamToTs.KeyId())
}
