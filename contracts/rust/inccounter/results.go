// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type MutableFuncCallIncrementResults struct {
	id int32
}

type ImmutableFuncCallIncrementResults struct {
	id int32
}

type MutableFuncCallIncrementRecurse5xResults struct {
	id int32
}

type ImmutableFuncCallIncrementRecurse5xResults struct {
	id int32
}

type MutableFuncIncrementResults struct {
	id int32
}

type ImmutableFuncIncrementResults struct {
	id int32
}

type MutableFuncInitResults struct {
	id int32
}

type ImmutableFuncInitResults struct {
	id int32
}

type MutableFuncLocalStateInternalCallResults struct {
	id int32
}

type ImmutableFuncLocalStateInternalCallResults struct {
	id int32
}

type MutableFuncLocalStatePostResults struct {
	id int32
}

type ImmutableFuncLocalStatePostResults struct {
	id int32
}

type MutableFuncLocalStateSandboxCallResults struct {
	id int32
}

type ImmutableFuncLocalStateSandboxCallResults struct {
	id int32
}

type MutableFuncLoopResults struct {
	id int32
}

type ImmutableFuncLoopResults struct {
	id int32
}

type MutableFuncPostIncrementResults struct {
	id int32
}

type ImmutableFuncPostIncrementResults struct {
	id int32
}

type MutableFuncRepeatManyResults struct {
	id int32
}

type ImmutableFuncRepeatManyResults struct {
	id int32
}

type MutableFuncTestLeb128Results struct {
	id int32
}

type ImmutableFuncTestLeb128Results struct {
	id int32
}

type MutableFuncWhenMustIncrementResults struct {
	id int32
}

type ImmutableFuncWhenMustIncrementResults struct {
	id int32
}

type MutableViewGetCounterResults struct {
	id int32
}

func (s MutableViewGetCounterResults) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultCounter])
}

type ImmutableViewGetCounterResults struct {
	id int32
}

func (s ImmutableViewGetCounterResults) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultCounter])
}
