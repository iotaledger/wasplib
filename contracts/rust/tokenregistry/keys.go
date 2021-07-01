// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	IdxParamColor       = 0
	IdxParamDescription = 1
	IdxParamUserDefined = 2
	IdxStateColorList   = 3
	IdxStateRegistry    = 4
)

const keyMapLen = 5

var keyMap = [keyMapLen]wasmlib.Key{
	ParamColor,
	ParamDescription,
	ParamUserDefined,
	StateColorList,
	StateRegistry,
}

var idxMap [keyMapLen]wasmlib.Key32
