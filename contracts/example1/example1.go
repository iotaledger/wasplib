// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package example1

import "github.com/iotaledger/wasplib/client"

const KeyParamString = client.Key("paramString")
const KeyStoredString = client.Key("storedString")

func OnLoad() {
	// declare entry points of the smart contract
	exports := client.NewScExports()
	exports.AddCall("storeString", storeString)
	exports.AddView("getString", getString)
}

// storeString entry point
func storeString(ctx *client.ScCallContext) {
	// take parameter paramString
	par := ctx.Params().GetString(KeyParamString)
	if !par.Exists() {
		ctx.Panic("string parameter not found") // panic if parameter does not exist
	}
	// store the string in "storedString" variable
	ctx.State().GetString(KeyStoredString).SetValue(par.Value())
	// log the text
	msg := "Message stored: " + par.Value()
	ctx.Log(msg)
}

// getString view
func getString(ctx *client.ScViewContext) {
	// take the stored string
	s := ctx.State().GetString(KeyStoredString).Value()
	// return the string value in the result dictionary
	ctx.Results().GetString(KeyParamString).SetValue(s)
}
