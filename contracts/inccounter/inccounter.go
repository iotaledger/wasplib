// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

var LocalStateMustIncrement = false

func funcInit(ctx *wasmlib.ScCallContext) {
	counter := ctx.Params().GetInt(ParamCounter).Value()
	if counter == 0 {
		return
	}
	ctx.State().GetInt(VarCounter).SetValue(counter)
}

func funcIncrement(ctx *wasmlib.ScCallContext) {
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

func funcCallIncrement(ctx *wasmlib.ScCallContext) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.CallSelf(HFuncCallIncrement, nil, nil)
	}
}

func funcCallIncrementRecurse5x(ctx *wasmlib.ScCallContext) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ctx.CallSelf(HFuncCallIncrementRecurse5x, nil, nil)
	}
}

func funcPostIncrement(ctx *wasmlib.ScCallContext) {
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

func viewGetCounter(ctx *wasmlib.ScViewContext) {
	counter := ctx.State().GetInt(VarCounter).Value()
	ctx.Results().GetInt(VarCounter).SetValue(counter)
}

func funcRepeatMany(ctx *wasmlib.ScCallContext) {
	counter := ctx.State().GetInt(VarCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := ctx.State().GetInt(VarNumRepeats)
	repeats := ctx.Params().GetInt(ParamNumRepeats).Value()
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

func funcWhenMustIncrement(ctx *wasmlib.ScCallContext) {
	ctx.Log("when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := ctx.State().GetInt(VarCounter)
	counter.SetValue(counter.Value() + 1)
}

func funcLocalStateInternalCall(ctx *wasmlib.ScCallContext) {
	LocalStateMustIncrement = false
	funcWhenMustIncrement(ctx)
	LocalStateMustIncrement = true
	funcWhenMustIncrement(ctx)
	funcWhenMustIncrement(ctx)
	// counter ends up as 2
}

func funcLocalStateSandboxCall(ctx *wasmlib.ScCallContext) {
	LocalStateMustIncrement = false
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	LocalStateMustIncrement = true
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	ctx.CallSelf(HFuncWhenMustIncrement, nil, nil)
	// counter ends up as 0 (non-existent)
}

func funcLocalStatePost(ctx *wasmlib.ScCallContext) {
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

func funcResultsTest(ctx *wasmlib.ScCallContext) {
	testMap(ctx.Results())
	checkMap(ctx.Results().Immutable())
	//ctx.CallSelf(HFuncResultsCheck, nil, nil)
}

func funcStateTest(ctx *wasmlib.ScCallContext) {
	testMap(ctx.State())
	ctx.CallSelf(HViewStateCheck, nil, nil)
}

func viewResultsCheck(ctx *wasmlib.ScViewContext) {
	checkMap(ctx.Results().Immutable())
}

func viewStateCheck(ctx *wasmlib.ScViewContext) {
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
