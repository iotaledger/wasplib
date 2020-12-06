// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

var (
	calls []func(sc *ScCallContext)
	views []func(sc *ScViewContext)
)

//export sc_call_entrypoint
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
	root := ScMutableMap{objId: 1}
	return ScExports{exports: root.GetStringArray(Key("exports"))}
}

func (ctx ScExports) AddCall(name string, f func(sc *ScCallContext)) {
	index := int32(len(calls))
	calls = append(calls, f)
	ctx.exports.GetString(index).SetValue(name)
}

func (ctx ScExports) AddView(name string, f func(sc *ScViewContext)) {
	index := int32(len(views))
	views = append(views, f)
	ctx.exports.GetString(index | 0x8000).SetValue(name)
}

func Nothing(sc *ScCallContext) {
	sc.Log("Doing nothing as requested. Oh, wait...")
}
