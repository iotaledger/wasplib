// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coregovernance

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableAddAllowedStateControllerAddressParams struct {
	id int32
}

func (s ImmutableAddAllowedStateControllerAddressParams) StateControllerAddress() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, ParamStateControllerAddress.KeyID())
}

type MutableAddAllowedStateControllerAddressParams struct {
	id int32
}

func (s MutableAddAllowedStateControllerAddressParams) StateControllerAddress() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, ParamStateControllerAddress.KeyID())
}

type ImmutableRemoveAllowedStateControllerAddressParams struct {
	id int32
}

func (s ImmutableRemoveAllowedStateControllerAddressParams) StateControllerAddress() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, ParamStateControllerAddress.KeyID())
}

type MutableRemoveAllowedStateControllerAddressParams struct {
	id int32
}

func (s MutableRemoveAllowedStateControllerAddressParams) StateControllerAddress() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, ParamStateControllerAddress.KeyID())
}

type ImmutableRotateStateControllerParams struct {
	id int32
}

func (s ImmutableRotateStateControllerParams) StateControllerAddress() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, ParamStateControllerAddress.KeyID())
}

type MutableRotateStateControllerParams struct {
	id int32
}

func (s MutableRotateStateControllerParams) StateControllerAddress() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, ParamStateControllerAddress.KeyID())
}
