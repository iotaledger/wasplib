// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

var (
	funcs []func(ctx ScFuncContext)
	views []func(ctx ScViewContext)
)

//export on_call
func OnCall(index int32) {
	ctx := ScFuncContext{}
	ctx.Require(GetObjectID(OBJ_ID_ROOT, KeyState, TYPE_MAP) == OBJ_ID_STATE, "object id mismatch")
	ctx.Require(GetObjectID(OBJ_ID_ROOT, KeyParams, TYPE_MAP) == OBJ_ID_PARAMS, "object id mismatch")
	ctx.Require(GetObjectID(OBJ_ID_ROOT, KeyResults, TYPE_MAP) == OBJ_ID_RESULTS, "object id mismatch")

	if (index & 0x8000) != 0 {
		views[index&0x7fff](ScViewContext{})
		return
	}
	funcs[index](ctx)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScExports struct {
	exports ScMutableStringArray
}

func NewScExports() ScExports {
	funcs = nil
	views = nil
	exports := Root.GetStringArray(KeyExports)
	// tell host what our highest predefined key is
	// this helps detect missing or extra keys
	exports.GetString(int32(KeyZzzzzzz)).SetValue("Go:KeyZzzzzzz")
	return ScExports{exports: exports}
}

func (ctx ScExports) AddFunc(name string, f func(ctx ScFuncContext)) {
	index := int32(len(funcs))
	funcs = append(funcs, f)
	ctx.exports.GetString(index).SetValue(name)
}

func (ctx ScExports) AddView(name string, f func(ctx ScViewContext)) {
	index := int32(len(views))
	views = append(views, f)
	ctx.exports.GetString(index | 0x8000).SetValue(name)
}
