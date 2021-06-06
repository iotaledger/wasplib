// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncCallIncrement, funcCallIncrementThunk)
	exports.AddFunc(FuncCallIncrementRecurse5x, funcCallIncrementRecurse5xThunk)
	exports.AddFunc(FuncIncrement, funcIncrementThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncLocalStateInternalCall, funcLocalStateInternalCallThunk)
	exports.AddFunc(FuncLocalStatePost, funcLocalStatePostThunk)
	exports.AddFunc(FuncLocalStateSandboxCall, funcLocalStateSandboxCallThunk)
	exports.AddFunc(FuncLoop, funcLoopThunk)
	exports.AddFunc(FuncPostIncrement, funcPostIncrementThunk)
	exports.AddFunc(FuncRepeatMany, funcRepeatManyThunk)
	exports.AddFunc(FuncTestLeb128, funcTestLeb128Thunk)
	exports.AddFunc(FuncWhenMustIncrement, funcWhenMustIncrementThunk)
	exports.AddView(ViewGetCounter, viewGetCounterThunk)

	for i, key := range keyMap {
		idxMap[i] = wasmlib.GetKeyIdFromString(key)
	}
}

type FuncCallIncrementContext struct {
	State IncCounterFuncState
}

func funcCallIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcCallIncrement")
	f := &FuncCallIncrementContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcCallIncrement(ctx, f)
	ctx.Log("inccounter.funcCallIncrement ok")
}

type FuncCallIncrementRecurse5xContext struct {
	State IncCounterFuncState
}

func funcCallIncrementRecurse5xThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcCallIncrementRecurse5x")
	f := &FuncCallIncrementRecurse5xContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcCallIncrementRecurse5x(ctx, f)
	ctx.Log("inccounter.funcCallIncrementRecurse5x ok")
}

type FuncIncrementContext struct {
	State IncCounterFuncState
}

func funcIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcIncrement")
	f := &FuncIncrementContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcIncrement(ctx, f)
	ctx.Log("inccounter.funcIncrement ok")
}

type FuncInitParams struct {
	Counter wasmlib.ScImmutableInt64 // value to initialize state counter with
}

type FuncInitContext struct {
	Params FuncInitParams
	State  IncCounterFuncState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcInit")
	p := ctx.Params().MapId()
	f := &FuncInitContext{
		Params: FuncInitParams{
			Counter: wasmlib.NewScImmutableInt64(p, idxMap[IdxParamCounter]),
		},
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcInit(ctx, f)
	ctx.Log("inccounter.funcInit ok")
}

type FuncLocalStateInternalCallContext struct {
	State IncCounterFuncState
}

func funcLocalStateInternalCallThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStateInternalCall")
	f := &FuncLocalStateInternalCallContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcLocalStateInternalCall(ctx, f)
	ctx.Log("inccounter.funcLocalStateInternalCall ok")
}

type FuncLocalStatePostContext struct {
	State IncCounterFuncState
}

func funcLocalStatePostThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStatePost")
	f := &FuncLocalStatePostContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcLocalStatePost(ctx, f)
	ctx.Log("inccounter.funcLocalStatePost ok")
}

type FuncLocalStateSandboxCallContext struct {
	State IncCounterFuncState
}

func funcLocalStateSandboxCallThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLocalStateSandboxCall")
	f := &FuncLocalStateSandboxCallContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcLocalStateSandboxCall(ctx, f)
	ctx.Log("inccounter.funcLocalStateSandboxCall ok")
}

type FuncLoopContext struct {
	State IncCounterFuncState
}

func funcLoopThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcLoop")
	f := &FuncLoopContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcLoop(ctx, f)
	ctx.Log("inccounter.funcLoop ok")
}

type FuncPostIncrementContext struct {
	State IncCounterFuncState
}

func funcPostIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcPostIncrement")
	f := &FuncPostIncrementContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcPostIncrement(ctx, f)
	ctx.Log("inccounter.funcPostIncrement ok")
}

type FuncRepeatManyParams struct {
	NumRepeats wasmlib.ScImmutableInt64 // number of times to recursively call myself
}

type FuncRepeatManyContext struct {
	Params FuncRepeatManyParams
	State  IncCounterFuncState
}

func funcRepeatManyThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcRepeatMany")
	p := ctx.Params().MapId()
	f := &FuncRepeatManyContext{
		Params: FuncRepeatManyParams{
			NumRepeats: wasmlib.NewScImmutableInt64(p, idxMap[IdxParamNumRepeats]),
		},
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcRepeatMany(ctx, f)
	ctx.Log("inccounter.funcRepeatMany ok")
}

type FuncTestLeb128Context struct {
	State IncCounterFuncState
}

func funcTestLeb128Thunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcTestLeb128")
	f := &FuncTestLeb128Context{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcTestLeb128(ctx, f)
	ctx.Log("inccounter.funcTestLeb128 ok")
}

type FuncWhenMustIncrementContext struct {
	State IncCounterFuncState
}

func funcWhenMustIncrementThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("inccounter.funcWhenMustIncrement")
	f := &FuncWhenMustIncrementContext{
		State: IncCounterFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcWhenMustIncrement(ctx, f)
	ctx.Log("inccounter.funcWhenMustIncrement ok")
}

type ViewGetCounterResults struct {
	Counter wasmlib.ScMutableInt64
}

type ViewGetCounterContext struct {
	Results ViewGetCounterResults
	State   IncCounterViewState
}

func viewGetCounterThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("inccounter.viewGetCounter")
	r := ctx.Results().MapId()
	f := &ViewGetCounterContext{
		Results: ViewGetCounterResults{
			Counter: wasmlib.NewScMutableInt64(r, idxMap[IdxResultCounter]),
		},
		State: IncCounterViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	viewGetCounter(ctx, f)
	ctx.Log("inccounter.viewGetCounter ok")
}
