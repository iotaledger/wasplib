// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type TestCoreFuncState struct {
	stateId int32
}

func (s TestCoreFuncState) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.stateId, VarCounter.KeyId())
}

func (s TestCoreFuncState) HnameEP() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.stateId, VarHnameEP.KeyId())
}

func (s TestCoreFuncState) MintedColor() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.stateId, VarMintedColor.KeyId())
}

func (s TestCoreFuncState) MintedSupply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.stateId, VarMintedSupply.KeyId())
}

type TestCoreViewState struct {
	stateId int32
}

func (s TestCoreViewState) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.stateId, VarCounter.KeyId())
}

func (s TestCoreViewState) HnameEP() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.stateId, VarHnameEP.KeyId())
}

func (s TestCoreViewState) MintedColor() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.stateId, VarMintedColor.KeyId())
}

func (s TestCoreViewState) MintedSupply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.stateId, VarMintedSupply.KeyId())
}
