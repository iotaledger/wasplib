// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableArrayLengthResults struct {
	id int32
}

func (s ImmutableArrayLengthResults) Length() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxResultLength])
}

type MutableArrayLengthResults struct {
	id int32
}

func (s MutableArrayLengthResults) Length() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxResultLength])
}

type ImmutableArrayValueResults struct {
	id int32
}

func (s ImmutableArrayValueResults) Value() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxResultValue])
}

type MutableArrayValueResults struct {
	id int32
}

func (s MutableArrayValueResults) Value() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxResultValue])
}

type ImmutableBlockRecordResults struct {
	id int32
}

func (s ImmutableBlockRecordResults) Record() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, idxMap[IdxResultRecord])
}

type MutableBlockRecordResults struct {
	id int32
}

func (s MutableBlockRecordResults) Record() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, idxMap[IdxResultRecord])
}

type ImmutableBlockRecordsResults struct {
	id int32
}

func (s ImmutableBlockRecordsResults) Count() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxResultCount])
}

type MutableBlockRecordsResults struct {
	id int32
}

func (s MutableBlockRecordsResults) Count() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxResultCount])
}

type ImmutableIotaBalanceResults struct {
	id int32
}

func (s ImmutableIotaBalanceResults) Iotas() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultIotas])
}

type MutableIotaBalanceResults struct {
	id int32
}

func (s MutableIotaBalanceResults) Iotas() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultIotas])
}
