// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	IdxParamCounter    = 0
	IdxParamDummy      = 1
	IdxParamNumRepeats = 2
	IdxResultCounter   = 3
	IdxStateCounter    = 4
	IdxStateNumRepeats = 5
)

var keyMap = [6]wasmlib.Key{
	ParamCounter,
	ParamDummy,
	ParamNumRepeats,
	ResultCounter,
	StateCounter,
	StateNumRepeats,
}

var idxMap [6]wasmlib.Key32
