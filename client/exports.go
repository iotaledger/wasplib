// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

var (
	calls []func(ctx *ScCallContext)
	views []func(ctx *ScViewContext)
)

//export on_call_entrypoint
func ScCallEntrypoint(index int32) {
	if (index & 0x8000) != 0 {
		views[index&0x7fff](&ScViewContext{})
		return
	}
	calls[index](&ScCallContext{})
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScExports struct {
	exports ScMutableStringArray
}

func NewScExports() ScExports {
	exports := Root.GetStringArray(KeyExports)
	// tell host what our highest predefined key is
	// this helps detect missing or extra keys
	exports.GetString(int32(KeyZzzzzzz)).SetValue("Go:KeyZzzzzzz")
	return ScExports{exports: exports}
}

func (ctx ScExports) AddCall(name string, f func(ctx *ScCallContext)) {
	index := int32(len(calls))
	calls = append(calls, f)
	ctx.exports.GetString(index).SetValue(name)
}

func (ctx ScExports) AddView(name string, f func(ctx *ScViewContext)) {
	index := int32(len(views))
	views = append(views, f)
	ctx.exports.GetString(index | 0x8000).SetValue(name)
}

func Nothing(ctx *ScCallContext) {
	ctx.Log("Doing nothing as requested. Oh, wait...")
}
