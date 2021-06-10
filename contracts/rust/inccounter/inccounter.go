// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

var LocalStateMustIncrement = false

func funcCallIncrement(ctx wasmlib.ScFuncContext, f *FuncCallIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.CallSelf(HFuncCallIncrement, nil, nil)
	}
}

func funcCallIncrementRecurse5x(ctx wasmlib.ScFuncContext, f *FuncCallIncrementRecurse5xContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ctx.CallSelf(HFuncCallIncrementRecurse5x, nil, nil)
	}
}

func funcIncrement(ctx wasmlib.ScFuncContext, f *FuncIncrementContext) {
	counter := f.State.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcInit(ctx wasmlib.ScFuncContext, f *FuncInitContext) {
	if f.Params.Counter().Exists() {
		counter := f.Params.Counter().Value()
		f.State.Counter().SetValue(counter)
	}
}

func funcLocalStateInternalCall(ctx wasmlib.ScFuncContext, f *FuncLocalStateInternalCallContext) {
	{
		LocalStateMustIncrement = false
	}
	funcWhenMustIncrementState(ctx, f.State)
	{
		LocalStateMustIncrement = true
	}
	funcWhenMustIncrementState(ctx, f.State)
	funcWhenMustIncrementState(ctx, f.State)
	// counter ends up as 2
}

func funcLocalStatePost(ctx wasmlib.ScFuncContext, f *FuncLocalStatePostContext) {
	{
		LocalStateMustIncrement = false
	}
	// prevent multiple identical posts, need a dummy param to differentiate them
	localStatePost(ctx, 1)
	{
		LocalStateMustIncrement = true
	}
	localStatePost(ctx, 2)
	localStatePost(ctx, 3)
	// counter ends up as 0
}

func localStatePost(ctx wasmlib.ScFuncContext, nr int64) {
	//note: we add a dummy parameter here to prevent "duplicate outputs not allowed" error
	params := wasmlib.NewScMutableMap()
	params.GetInt64(StateCounter).SetValue(nr)
	transfer := wasmlib.NewScTransferIotas(1)
	ctx.PostSelf(HFuncWhenMustIncrement, params, transfer, 0)
}

func funcLocalStateSandboxCall(ctx wasmlib.ScFuncContext, f *FuncLocalStateSandboxCallContext) {
	{
		LocalStateMustIncrement = false
	}
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	{
		LocalStateMustIncrement = true
	}
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	// counter ends up as 0
}

func funcLoop(ctx wasmlib.ScFuncContext, f *FuncLoopContext) {
	for {
	}
}

func funcPostIncrement(ctx wasmlib.ScFuncContext, f *FuncPostIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		transfer := wasmlib.NewScTransferIotas(1)
		ctx.PostSelf(HFuncPostIncrement, nil, transfer, 0)
	}
}

func funcRepeatMany(ctx wasmlib.ScFuncContext, f *FuncRepeatManyContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := f.State.NumRepeats()
	repeats := f.Params.NumRepeats().Value()
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
	transfer := wasmlib.NewScTransferIotas(1)
	ctx.PostSelf(HFuncRepeatMany, nil, transfer, 0)
}

func funcWhenMustIncrement(ctx wasmlib.ScFuncContext, f *FuncWhenMustIncrementContext) {
	funcWhenMustIncrementState(ctx, f.State)
}

func funcWhenMustIncrementState(ctx wasmlib.ScFuncContext, state MutableIncCounterState) {
	ctx.Log("when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := state.Counter()
	counter.SetValue(counter.Value() + 1)
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
func viewGetCounter(ctx wasmlib.ScViewContext, f *ViewGetCounterContext) {
	counter := f.State.Counter()
	if counter.Exists() {
		f.Results.Counter().SetValue(counter.Value())
	}
}

func funcTestLeb128(ctx wasmlib.ScFuncContext, f *FuncTestLeb128Context) {
	save(ctx, "v-1", -1)
	save(ctx, "v-2", -2)
	save(ctx, "v-126", -126)
	save(ctx, "v-127", -127)
	save(ctx, "v-128", -128)
	save(ctx, "v-129", -129)
	save(ctx, "v0", 0)
	save(ctx, "v+1", 1)
	save(ctx, "v+2", 2)
	save(ctx, "v+126", 126)
	save(ctx, "v+127", 127)
	save(ctx, "v+128", 128)
	save(ctx, "v+129", 129)
}

func save(ctx wasmlib.ScFuncContext, name string, value int64) {
	encoder := wasmlib.NewBytesEncoder()
	encoder.Int64(value)
	spot := ctx.State().GetBytes(wasmlib.Key(name))
	spot.SetValue(encoder.Data())

	bytes := spot.Value()
	decoder := wasmlib.NewBytesDecoder(bytes)
	retrieved := decoder.Int64()
	if retrieved != value {
		ctx.Log(name + " in : " + ctx.Utility().String(value))
		ctx.Log(name + " out: " + ctx.Utility().String(retrieved))
	}
}
