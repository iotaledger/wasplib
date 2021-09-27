// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableStringArray = ArrayOfImmutableString

type ArrayOfImmutableString struct {
	objID int32
}

func (a ArrayOfImmutableString) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableString) GetString(index int32) wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(a.objID, wasmlib.Key32(index))
}

type MutableStringArray = ArrayOfMutableString

type ArrayOfMutableString struct {
	objID int32
}

func (a ArrayOfMutableString) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableString) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableString) GetString(index int32) wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(a.objID, wasmlib.Key32(index))
}
