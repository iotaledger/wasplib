// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "helloworld"
const ScDescription = "The ubiquitous hello world demo"
const ScHname = coretypes.Hname(0x0683223c)

const VarHelloWorld = "helloWorld"

const FuncHelloWorld = "helloWorld"
const ViewGetHelloWorld = "getHelloWorld"

const HFuncHelloWorld = coretypes.Hname(0x9d042e65)
const HViewGetHelloWorld = coretypes.Hname(0x210439ce)
