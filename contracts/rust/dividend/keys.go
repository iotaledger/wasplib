// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	IdxParamAddress     = 0
	IdxParamFactor      = 1
	IdxParamOwner       = 2
	IdxResultFactor     = 3
	IdxStateFactor      = 4
	IdxStateMemberList  = 5
	IdxStateMembers     = 6
	IdxStateOwner       = 7
	IdxStateTotalFactor = 8
)

const keyMapLen = 9

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAddress,
	ParamFactor,
	ParamOwner,
	ResultFactor,
	StateFactor,
	StateMemberList,
	StateMembers,
	StateOwner,
	StateTotalFactor,
}

var idxMap [keyMapLen]wasmlib.Key32
