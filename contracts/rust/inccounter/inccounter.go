// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

var LocalStateMustIncrement = false

func funcCallIncrement(ctx *wasmlib.ScFuncContext, params *FuncCallIncrementParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.CallSelf(HFuncCallIncrement, nil, nil)
	}
}

func funcCallIncrementRecurse5x(ctx *wasmlib.ScFuncContext, params *FuncCallIncrementRecurse5xParams) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ctx.CallSelf(HFuncCallIncrementRecurse5x, nil, nil)
	}
}

func funcIncrement(ctx *wasmlib.ScFuncContext, params *FuncIncrementParams) {
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

func funcInit(ctx *wasmlib.ScFuncContext, params *FuncInitParams) {
	counter := params.Counter.Value()
	if counter == 0 {
		return
	}
	ctx.State().GetInt(VarCounter).SetValue(counter)
}

func funcLocalStateInternalCall(ctx *wasmlib.ScFuncContext, params *FuncLocalStateInternalCallParams) {
	LocalStateMustIncrement = false
	par := &FuncWhenMustIncrementParams{}
	funcWhenMustIncrement(ctx, par)
	LocalStateMustIncrement = true
	funcWhenMustIncrement(ctx, par)
	funcWhenMustIncrement(ctx, par)
	// counter ends up as 2
}

func funcLocalStatePost(ctx *wasmlib.ScFuncContext, params *FuncLocalStatePostParams) {
	LocalStateMustIncrement = false
	request := &wasmlib.PostRequestParams{
		ContractId: ctx.ContractId(),
		Function:   HFuncWhenMustIncrement,
		Params:     nil,
		Transfer:   nil,
		Delay:      0,
	}
	ctx.Post(request)
	LocalStateMustIncrement = true
	ctx.Post(request)
	ctx.Post(request)
	// counter ends up as 0 (non-existent)
}

func funcLocalStateSandboxCall(ctx *wasmlib.ScFuncContext, params *FuncLocalStateSandboxCallParams) {
	LocalStateMustIncrement = false
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	LocalStateMustIncrement = true
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	// counter ends up as 0 (non-existent)
}

func funcPostIncrement(ctx *wasmlib.ScFuncContext, params *FuncPostIncrementParams) {
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

func funcRepeatMany(ctx *wasmlib.ScFuncContext, params *FuncRepeatManyParams) {
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

func funcResultsTest(ctx *wasmlib.ScFuncContext, params *FuncResultsTestParams) {
	testMap(ctx.Results())
	checkMap(ctx.Results().Immutable())
	//ctx.CallSelf(HFuncResultsCheck, nil, nil)
}

func funcStateTest(ctx *wasmlib.ScFuncContext, params *FuncStateTestParams) {
	testMap(ctx.State())
	ctx.CallSelf(HViewStateCheck, nil, nil)
}

func funcWhenMustIncrement(ctx *wasmlib.ScFuncContext, params *FuncWhenMustIncrementParams) {
	ctx.Log("when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

func viewGetCounter(ctx *wasmlib.ScViewContext, params *ViewGetCounterParams) {
	counter := ctx.State().GetInt(VarCounter).Value()
	ctx.Results().GetInt(VarCounter).SetValue(counter)
}

func viewResultsCheck(ctx *wasmlib.ScViewContext, params *ViewResultsCheckParams) {
	checkMap(ctx.Results().Immutable())
}

func viewStateCheck(ctx *wasmlib.ScViewContext, params *ViewStateCheckParams) {
	checkMap(ctx.State())
}

func testMap(kvstore wasmlib.ScMutableMap) {
	int1 := kvstore.GetInt(VarInt1)
	check(int1.Value() == 0)
	int1.SetValue(1)

	string1 := kvstore.GetString(VarString1)
	check(string1.Value() == "")
	string1.SetValue("a")
	check(string1.Value() == "a")

	ia1 := kvstore.GetIntArray(VarIntArray1)
	int2 := ia1.GetInt(0)
	check(int2.Value() == 0)
	int2.SetValue(2)
	int3 := ia1.GetInt(1)
	check(int3.Value() == 0)
	int3.SetValue(3)

	sa1 := kvstore.GetStringArray(VarStringArray1)
	string2 := sa1.GetString(0)
	check(string2.Value() == "")
	string2.SetValue("bc")
	string3 := sa1.GetString(1)
	check(string3.Value() == "")
	string3.SetValue("def")
}

func checkMap(kvstore wasmlib.ScImmutableMap) {
	int1 := kvstore.GetInt(VarInt1)
	check(int1.Value() == 1)

	string1 := kvstore.GetString(VarString1)
	check(string1.Value() == "a")

	ia1 := kvstore.GetIntArray(VarIntArray1)
	int2 := ia1.GetInt(0)
	check(int2.Value() == 2)
	int3 := ia1.GetInt(1)
	check(int3.Value() == 3)

	sa1 := kvstore.GetStringArray(VarStringArray1)
	string2 := sa1.GetString(0)
	check(string2.Value() == "bc")
	string3 := sa1.GetString(1)
	check(string3.Value() == "def")
}

func check(condition bool) {
	if !condition {
		panic("Check failed!")
	}
}
