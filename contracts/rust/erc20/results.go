// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableViewAllowanceResults struct {
	id int32
}

func (s ImmutableViewAllowanceResults) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultAmount])
}

type MutableViewAllowanceResults struct {
	id int32
}

func (s MutableViewAllowanceResults) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultAmount])
}

type ImmutableViewBalanceOfResults struct {
	id int32
}

func (s ImmutableViewBalanceOfResults) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultAmount])
}

type MutableViewBalanceOfResults struct {
	id int32
}

func (s MutableViewBalanceOfResults) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultAmount])
}

type ImmutableViewTotalSupplyResults struct {
	id int32
}

func (s ImmutableViewTotalSupplyResults) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultSupply])
}

type MutableViewTotalSupplyResults struct {
	id int32
}

func (s MutableViewTotalSupplyResults) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultSupply])
}
