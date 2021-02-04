// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import "github.com/iotaledger/wasplib/client"

const ScName = "helloworld"
const ScDescription = "The ubiquitous hello world demo"
const ScHname = client.ScHname(0x0683223c)

const VarHelloWorld = client.Key("helloWorld")

const FuncHelloWorld = "helloWorld"
const ViewGetHelloWorld = "getHelloWorld"

const HFuncHelloWorld = client.ScHname(0x9d042e65)
const HViewGetHelloWorld = client.ScHname(0x210439ce)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncHelloWorld, funcHelloWorld)
    exports.AddView(ViewGetHelloWorld, viewGetHelloWorld)
}
