// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ArrayOfImmutableBet struct {
	objID int32
}

func (a ArrayOfImmutableBet) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableBet) GetBet(index int32) ImmutableBet {
	return ImmutableBet{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type ImmutableFairRouletteState struct {
	id int32
}

func (s ImmutableFairRouletteState) Bets() ArrayOfImmutableBet {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateBets], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableBet{objID: arrID}
}

func (s ImmutableFairRouletteState) LastWinningNumber() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxStateLastWinningNumber])
}

func (s ImmutableFairRouletteState) LockedBets() ArrayOfImmutableBet {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateLockedBets], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableBet{objID: arrID}
}

func (s ImmutableFairRouletteState) PlayPeriod() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxStatePlayPeriod])
}

type ArrayOfMutableBet struct {
	objID int32
}

func (a ArrayOfMutableBet) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableBet) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableBet) GetBet(index int32) MutableBet {
	return MutableBet{objID: a.objID, keyID: wasmlib.Key32(index)}
}

type MutableFairRouletteState struct {
	id int32
}

func (s MutableFairRouletteState) Bets() ArrayOfMutableBet {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateBets], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableBet{objID: arrID}
}

func (s MutableFairRouletteState) LastWinningNumber() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxStateLastWinningNumber])
}

func (s MutableFairRouletteState) LockedBets() ArrayOfMutableBet {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateLockedBets], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableBet{objID: arrID}
}

func (s MutableFairRouletteState) PlayPeriod() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxStatePlayPeriod])
}
