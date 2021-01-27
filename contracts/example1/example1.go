// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package example1

import "github.com/iotaledger/wasplib/client"

const ParamString = client.Key("paramString")
const VarString = client.Key("storedString")

func OnLoad() {
    // declare entry points of the smart contract
    exports := client.NewScExports()
    exports.AddCall("storeString", storeString)
    exports.AddView("getString", getString)
    exports.AddCall("withdraw_iota", withdrawIota)
}

// storeString entry point stores a string provided as parameters
// in the state as a value of the key 'storedString'
// panics if parameter is not provided
func storeString(ctx *client.ScCallContext) {
    // take parameter paramString
    par := ctx.Params().GetString(ParamString)
    // require parameter exists
    ctx.Require(par.Exists(), "string parameter not found")

    // store the string in "storedString" variable
    ctx.State().GetString(VarString).SetValue(par.Value())
    // log the text
    msg := "Message stored: " + par.Value()
    ctx.Log(msg)
}

// getString view returns the string value of the key 'storedString'
// The call return result as a key/value dictionary.
// the returned value in the result is under key 'paramString'
func getString(ctx *client.ScViewContext) {
    // take the stored string
    s := ctx.State().GetString(VarString).Value()
    // return the string value in the result dictionary
    ctx.Results().GetString(ParamString).SetValue(s)
}

// withdraw_iota sends all iotas contained in the contract's account
// to the caller's L1 address.
// Panics of the caller is not an address
// Panics if the address is not the creator of the contract is the caller
// The caller will be address only if request is sent from the wallet on the L1, not a smart contract
func withdrawIota(ctx *client.ScCallContext) {
    creator := ctx.ContractCreator()
    caller := ctx.Caller()

    ctx.Require(creator.Equals(caller), "not authorised")
    ctx.Require(caller.IsAddress(), "caller must be an address")

    bal := ctx.Balances().Balance(client.IOTA)
    if bal > 0 {
        ctx.TransferToAddress(caller.Address(), client.NewScTransfer(client.IOTA, bal))
    }
}
