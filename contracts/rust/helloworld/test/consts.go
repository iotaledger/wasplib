// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "helloworld"
const ScDescription = "The ubiquitous hello world demo"
const HScName = wasmlib.ScHname(0x0683223c)

const ResultHelloWorld = "helloWorld"

const VarDummy = "dummy"

const FuncHelloWorld = "helloWorld"
const ViewGetHelloWorld = "getHelloWorld"

const HFuncHelloWorld = wasmlib.ScHname(0x9d042e65)
const HViewGetHelloWorld = wasmlib.ScHname(0x210439ce)
