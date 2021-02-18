// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

var LocalStateMustIncrement = false

func funcCallIncrement(ctx wasmlib.ScFuncContext, params *FuncCallIncrementParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.CallSelf(HFuncCallIncrement, nil, nil)
	}
}

func funcCallIncrementRecurse5x(ctx wasmlib.ScFuncContext, params *FuncCallIncrementRecurse5xParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ctx.CallSelf(HFuncCallIncrementRecurse5x, nil, nil)
	}
}

func funcIncrement(ctx wasmlib.ScFuncContext, params *FuncIncrementParams) {
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

func funcInit(ctx wasmlib.ScFuncContext, params *FuncInitParams) {
    if params.Counter.Exists() {
	counter := params.Counter.Value()
        ctx.State().GetInt(VarCounter).SetValue(counter)
	}
}

func funcLocalStateInternalCall(ctx wasmlib.ScFuncContext, params *FuncLocalStateInternalCallParams) {
    {
        LocalStateMustIncrement = false
    }
    par := FuncWhenMustIncrementParams{}
    funcWhenMustIncrement(ctx, &par)
    {
        LocalStateMustIncrement = true
    }
    funcWhenMustIncrement(ctx, &par)
    funcWhenMustIncrement(ctx, &par)
    // counter ends up as 2
}

func funcLocalStatePost(ctx wasmlib.ScFuncContext, params *FuncLocalStatePostParams) {
    {
        LocalStateMustIncrement = false
    }
	request := &wasmlib.PostRequestParams{
		ContractId: ctx.ContractId(),
		Function:   HFuncWhenMustIncrement,
		Params:     nil,
		Transfer:   nil,
		Delay:      0,
	}
	ctx.Post(request)
    {
        LocalStateMustIncrement = true
    }
    ctx.Post(request)
    ctx.Post(request)
    // counter ends up as 0
}

func funcLocalStateSandboxCall(ctx wasmlib.ScFuncContext, params *FuncLocalStateSandboxCallParams) {
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

func funcPostIncrement(ctx wasmlib.ScFuncContext, params *FuncPostIncrementParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.Post(&wasmlib.PostRequestParams{
			ContractId: ctx.ContractId(),
			Function:   HFuncPostIncrement,
			Params:     nil,
			Transfer:   nil,
			Delay:      0,
		})
	}
}

func funcRepeatMany(ctx wasmlib.ScFuncContext, params *FuncRepeatManyParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := ctx.State().GetInt(VarNumRepeats)
	repeats := params.NumRepeats.Value()
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
	ctx.Post(&wasmlib.PostRequestParams{
		ContractId: ctx.ContractId(),
		Function:   HFuncRepeatMany,
		Params:     nil,
		Transfer:   nil,
		Delay:      0,
	})
}

func funcWhenMustIncrement(ctx wasmlib.ScFuncContext, params *FuncWhenMustIncrementParams) {
	ctx.Log("when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
func viewGetCounter(ctx wasmlib.ScViewContext, params *ViewGetCounterParams) {
	counter := ctx.State().GetInt(VarCounter)
	if counter.Exists() {
		ctx.Results().GetInt(VarCounter).SetValue(counter.Value())
	}
}
