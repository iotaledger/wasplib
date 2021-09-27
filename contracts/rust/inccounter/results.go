// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableGetCounterResults struct {
	id int32
}

func (s ImmutableGetCounterResults) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultCounter])
}

type MutableGetCounterResults struct {
	id int32
}

func (s MutableGetCounterResults) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultCounter])
}
