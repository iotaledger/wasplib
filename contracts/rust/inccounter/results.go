// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableViewGetCounterResults struct {
	id int32
}

func (s ImmutableViewGetCounterResults) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultCounter])
}

type MutableViewGetCounterResults struct {
	id int32
}

func (s MutableViewGetCounterResults) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultCounter])
}
