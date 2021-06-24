// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreeventlog

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	ScName        = "eventlog"
	ScDescription = "Core event log contract"
	HScName       = wasmlib.ScHname(0x661aa7d8)
)

const (
	ParamContractHname  = wasmlib.Key("contractHname")
	ParamFromTs         = wasmlib.Key("fromTs")
	ParamMaxLastRecords = wasmlib.Key("maxLastRecords")
	ParamToTs           = wasmlib.Key("toTs")
)

const (
	ResultNumRecords = wasmlib.Key("numRecords")
	ResultRecords    = wasmlib.Key("records")
)

const (
	ViewGetNumRecords = "getNumRecords"
	ViewGetRecords    = "getRecords"
)

const (
	HViewGetNumRecords = wasmlib.ScHname(0x2f4b4a8c)
	HViewGetRecords    = wasmlib.ScHname(0xd01a8085)
)
