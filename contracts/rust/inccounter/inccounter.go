// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

var LocalStateMustIncrement = false

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Counter().Exists() {
		counter := f.Params.Counter().Value()
		f.State.Counter().SetValue(counter)
	}
}

func funcCallIncrement(ctx wasmlib.ScFuncContext, f *CallIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ScFuncs.CallIncrement(ctx).Func.Call()
	}
}

func funcCallIncrementRecurse5x(ctx wasmlib.ScFuncContext, f *CallIncrementRecurse5xContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ScFuncs.CallIncrementRecurse5x(ctx).Func.Call()
	}
}

//nolint:unparam
func funcEndlessLoop(ctx wasmlib.ScFuncContext, f *EndlessLoopContext) {
	//nolint:staticcheck
	for {
		// intentional endless loop to see if Wasm VM can be interrupted
	}
}

func funcIncrement(ctx wasmlib.ScFuncContext, f *IncrementContext) {
	counter := f.State.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcLocalStateInternalCall(ctx wasmlib.ScFuncContext, f *LocalStateInternalCallContext) {
	LocalStateMustIncrement = false
	whenMustIncrementState(ctx, f.State)
	LocalStateMustIncrement = true
	whenMustIncrementState(ctx, f.State)
	whenMustIncrementState(ctx, f.State)
	// counter ends up as 2
}

//nolint:unparam
func funcLocalStatePost(ctx wasmlib.ScFuncContext, f *LocalStatePostContext) {
	LocalStateMustIncrement = false
	// prevent multiple identical posts, need a dummy param to differentiate them
	localStatePost(ctx, 1)
	LocalStateMustIncrement = true
	localStatePost(ctx, 2)
	localStatePost(ctx, 3)
	// counter ends up as 0
}

//nolint:unparam
func funcLocalStateSandboxCall(ctx wasmlib.ScFuncContext, f *LocalStateSandboxCallContext) {
	LocalStateMustIncrement = false
	ScFuncs.WhenMustIncrement(ctx).Func.Call()
	LocalStateMustIncrement = true
	ScFuncs.WhenMustIncrement(ctx).Func.Call()
	ScFuncs.WhenMustIncrement(ctx).Func.Call()
	// counter ends up as 0
}

func funcPostIncrement(ctx wasmlib.ScFuncContext, f *PostIncrementContext) {
	counter := f.State.Counter()
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ScFuncs.PostIncrement(ctx).Func.TransferIotas(1).Post()
	}
}

func funcRepeatMany(ctx wasmlib.ScFuncContext, f *RepeatManyContext) {
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
	ScFuncs.RepeatMany(ctx).Func.TransferIotas(1).Post()
}

func funcWhenMustIncrement(ctx wasmlib.ScFuncContext, f *WhenMustIncrementContext) {
	whenMustIncrementState(ctx, f.State)
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result

func viewGetCounter(ctx wasmlib.ScViewContext, f *GetCounterContext) {
	counter := f.State.Counter()
	if counter.Exists() {
		f.Results.Counter().SetValue(counter.Value())
	}
}

//nolint:unparam
func funcTestLeb128(ctx wasmlib.ScFuncContext, f *TestLeb128Context) {
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
	f := ScFuncs.WhenMustIncrement(ctx)
	f.Params.Dummy().SetValue(nr)
	f.Func.TransferIotas(1).Post()
}

func whenMustIncrementState(ctx wasmlib.ScFuncContext, state MutableIncCounterState) {
	ctx.Log("when_must_increment called")
	if !LocalStateMustIncrement {
		return
	}
	counter := state.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcIncrementWithDelay(ctx wasmlib.ScFuncContext, f *IncrementWithDelayContext) {
	delay := f.Params.Delay().Value()
	inc := ScFuncs.CallIncrement(ctx)
	inc.Func.Delay(delay).TransferIotas(1).Post()
}
