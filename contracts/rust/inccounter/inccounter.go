// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

var LocalStateMustIncrement = false

func funcInit(ctx wasmlib.ScFuncContext, f *FuncInitContext) {
	if f.Params.Counter().Exists() {
		counter := f.Params.Counter().Value()
		f.State.Counter().SetValue(counter)
	}
}

func funcCallIncrement(ctx wasmlib.ScFuncContext, f *FuncCallIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		NewIncCounterFunc(ctx).CallIncrement(wasmlib.NoScTransfers())
	}
}

func funcCallIncrementRecurse5x(ctx wasmlib.ScFuncContext, f *FuncCallIncrementRecurse5xContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		NewIncCounterFunc(ctx).CallIncrementRecurse5x(wasmlib.NoScTransfers())
	}
}

func funcEndlessLoop(ctx wasmlib.ScFuncContext, f *FuncEndlessLoopContext) {
	for {
	}
}

func funcIncrement(ctx wasmlib.ScFuncContext, f *FuncIncrementContext) {
	counter := f.State.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcLocalStateInternalCall(ctx wasmlib.ScFuncContext, f *FuncLocalStateInternalCallContext) {
	{
		LocalStateMustIncrement = false
	}
	whenMustIncrementState(ctx, f.State)
	{
		LocalStateMustIncrement = true
	}
	whenMustIncrementState(ctx, f.State)
	whenMustIncrementState(ctx, f.State)
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

func funcLocalStateSandboxCall(ctx wasmlib.ScFuncContext, f *FuncLocalStateSandboxCallContext) {
	{
		LocalStateMustIncrement = false
	}
	params := NewMutableFuncWhenMustIncrementParams()
	none := wasmlib.NoScTransfers()
	sc := NewIncCounterFunc(ctx)
	sc.WhenMustIncrement(params, none)
	{
		LocalStateMustIncrement = true
	}
	sc.WhenMustIncrement(params, none)
	sc.WhenMustIncrement(params, none)
	// counter ends up as 0
}

func funcPostIncrement(ctx wasmlib.ScFuncContext, f *FuncPostIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		transfer := wasmlib.NewScTransferIotas(1)
		NewIncCounterFunc(ctx).Post().PostIncrement(transfer)
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
	NewIncCounterFunc(ctx).Post().RepeatMany(NewMutableFuncRepeatManyParams(), transfer)
}

func funcWhenMustIncrement(ctx wasmlib.ScFuncContext, f *FuncWhenMustIncrementContext) {
	whenMustIncrementState(ctx, f.State)
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
	leb128Save(ctx, "v-1", -1)
	leb128Save(ctx, "v-2", -2)
	leb128Save(ctx, "v-126", -126)
	leb128Save(ctx, "v-127", -127)
	leb128Save(ctx, "v-128", -128)
	leb128Save(ctx, "v-129", -129)
	leb128Save(ctx, "v0", 0)
	leb128Save(ctx, "v+1", 1)
	leb128Save(ctx, "v+2", 2)
	leb128Save(ctx, "v+126", 126)
	leb128Save(ctx, "v+127", 127)
	leb128Save(ctx, "v+128", 128)
	leb128Save(ctx, "v+129", 129)
}

func leb128Save(ctx wasmlib.ScFuncContext, name string, value int64) {
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

func localStatePost(ctx wasmlib.ScFuncContext, nr int64) {
	// note: we add a dummy parameter here to prevent "duplicate outputs not allowed" error
	params := NewMutableFuncWhenMustIncrementParams()
	params.Dummy().SetValue(nr)
	transfer := wasmlib.NewScTransferIotas(1)
	NewIncCounterFunc(ctx).Post().WhenMustIncrement(params, transfer)
}

func whenMustIncrementState(ctx wasmlib.ScFuncContext, state MutableIncCounterState) {
	ctx.Log("when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := state.Counter()
	counter.SetValue(counter.Value() + 1)
}
