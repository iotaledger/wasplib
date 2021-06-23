// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableLastWinningNumberResults struct {
	id int32
}

func (s ImmutableLastWinningNumberResults) LastWinningNumber() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultLastWinningNumber])
}

type MutableLastWinningNumberResults struct {
	id int32
}

func (s MutableLastWinningNumberResults) LastWinningNumber() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultLastWinningNumber])
}
