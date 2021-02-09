// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package example1

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

// storeString entry point stores a string provided as parameters
// in the state as a value of the key 'storedString'
// panics if parameter is not provided
func funcStoreString(ctx *wasmlib.ScFuncContext, params *FuncStoreStringParams) {
	// take parameter paramString
	paramString := params.String.Value()

	// store the string in "storedString" variable
	ctx.State().GetString(VarString).SetValue(paramString)
	// log the text
	msg := "Message stored: " + paramString
	ctx.Log(msg)
}

// withdraw_iota sends all iotas contained in the contract's account
// to the caller's L1 address.
// Panics of the caller is not an address
// Panics if the address is not the creator of the contract is the caller
// The caller will be address only if request is sent from the wallet on the L1, not a smart contract
func funcWithdrawIota(ctx *wasmlib.ScFuncContext, params *FuncWithdrawIotaParams) {
	caller := ctx.Caller()
	ctx.Require(caller.IsAddress(), "caller must be an address")

	bal := ctx.Balances().Balance(wasmlib.IOTA)
	if bal > 0 {
		ctx.TransferToAddress(caller.Address(), wasmlib.NewScTransfer(wasmlib.IOTA, bal))
	}
}

// getString view returns the string value of the key 'storedString'
// The call return result as a key/value dictionary.
// the returned value in the result is under key 'paramString'
func viewGetString(ctx *wasmlib.ScViewContext, params *ViewGetStringParams) {
	// take the stored string
	s := ctx.State().GetString(VarString).Value()
	// return the string value in the result dictionary
	ctx.Results().GetString(ParamString).SetValue(s)
}
