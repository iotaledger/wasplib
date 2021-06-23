// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutablePlaceBetParams struct {
	id int32
}

func (s ImmutablePlaceBetParams) Number() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamNumber])
}

type MutablePlaceBetParams struct {
	id int32
}

func (s MutablePlaceBetParams) Number() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamNumber])
}

type ImmutablePlayPeriodParams struct {
	id int32
}

func (s ImmutablePlayPeriodParams) PlayPeriod() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxParamPlayPeriod])
}

type MutablePlayPeriodParams struct {
	id int32
}

func (s MutablePlayPeriodParams) PlayPeriod() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxParamPlayPeriod])
}
