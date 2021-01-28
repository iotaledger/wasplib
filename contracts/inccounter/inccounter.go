// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import "github.com/iotaledger/wasplib/client"

const KeyCounter = client.Key("counter")
const KeyNumRepeats = client.Key("num_repeats")

var LocalStateMustIncrement = false

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddCall("increment", increment)
	exports.AddCall("increment_call_increment", incrementCallIncrement)
	exports.AddCall("increment_call_increment_recurse5x", incrementCallIncrementRecurse5x)
	exports.AddCall("increment_post_increment", incrementPostIncrement)
	exports.AddView("increment_view_counter", incrementViewCounter)
	exports.AddCall("increment_repeat_many", incrementRepeatMany)
	exports.AddCall("increment_when_must_increment", incrementWhenMustIncrement)
	exports.AddCall("increment_local_state_internal_call", incrementLocalStateInternalCall)
	exports.AddCall("increment_local_state_sandbox_call", incrementLocalStateSandboxCall)
	exports.AddCall("increment_local_state_post", incrementLocalStatePost)
	exports.AddCall("nothing", client.Nothing)
	exports.AddCall("test", test)
	exports.AddCall("state_test", stateTest)
	exports.AddView("state_check", stateCheck)
	exports.AddCall("results_test", resultsTest)
	exports.AddView("results_check", resultsCheck)
}

func onInit(ctx *client.ScCallContext) {
	counter := ctx.Params().GetInt(KeyCounter).Value()
	if counter == 0 {
		return
	}
	ctx.State().GetInt(KeyCounter).SetValue(counter)
}

func increment(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	counter.SetValue(counter.Value() + 1)
}

func incrementCallIncrement(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.Call(ctx.ContractId().Hname(), client.NewHname("increment_call_increment"), nil, nil)
	}
}

func incrementCallIncrementRecurse5x(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		ctx.Call(ctx.ContractId().Hname(), client.NewHname("increment_call_increment_recurse5x"), nil, nil)
	}
}

func incrementPostIncrement(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
        ctx.Post(&client.PostRequestParams {
            Contract: ctx.ContractId(),
            Function: client.NewHname("increment_post_increment"),
            Params: nil,
            Transfer: nil,
            Delay: 0,
        })
	}
}

func incrementViewCounter(ctx *client.ScViewContext) {
	counter := ctx.State().GetInt(KeyCounter).Value()
	ctx.Results().GetInt(KeyCounter).SetValue(counter)
}

func incrementRepeatMany(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := ctx.State().GetInt(KeyNumRepeats)
	repeats := ctx.Params().GetInt(KeyNumRepeats).Value()
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
    ctx.Post(&client.PostRequestParams {
        Contract: ctx.ContractId(),
        Function: client.NewHname("increment_repeat_many"),
        Params: nil,
        Transfer: nil,
        Delay: 0,
    })
}

func incrementWhenMustIncrement(ctx *client.ScCallContext) {
	ctx.Log("increment_when_must_increment called")
	{
		if !LocalStateMustIncrement {
			return
		}
	}
	counter := ctx.State().GetInt(KeyCounter)
	counter.SetValue(counter.Value() + 1)
}

func incrementLocalStateInternalCall(ctx *client.ScCallContext) {
	incrementWhenMustIncrement(ctx)
	{
		LocalStateMustIncrement = true
	}
	incrementWhenMustIncrement(ctx)
	incrementWhenMustIncrement(ctx)
	// counter ends up as 2
}

func incrementLocalStateSandboxCall(ctx *client.ScCallContext) {
	ctx.Call(ctx.ContractId().Hname(), client.NewHname("increment_when_must_increment"), nil, nil)
	{
		LocalStateMustIncrement = true
	}
	ctx.Call(ctx.ContractId().Hname(), client.NewHname("increment_when_must_increment"), nil, nil)
	ctx.Call(ctx.ContractId().Hname(), client.NewHname("increment_when_must_increment"), nil, nil)
	// counter ends up as 0
}

func incrementLocalStatePost(ctx *client.ScCallContext) {
    request := &client.PostRequestParams {
        Contract: ctx.ContractId(),
        Function: client.NewHname("increment_when_must_increment"),
        Params: nil,
        Transfer: nil,
        Delay: 0,
    }
    ctx.Post(request)
	{
		LocalStateMustIncrement = true
	}
    ctx.Post(request)
    ctx.Post(request)
	// counter ends up as 0
}

func test(_sc *client.ScCallContext) {
	KeyId := client.GetKeyIdFromString("timestamp")
	client.SetInt(1, KeyId, 123456789)
	timestamp := client.GetInt(1, KeyId)
	client.SetInt(1, KeyId, timestamp)
	KeyId2 := client.GetKeyIdFromString("string")
	client.SetBytes(1, KeyId2, client.TYPE_STRING, []byte("Test"))
	s1 := client.GetBytes(1, KeyId2, client.TYPE_STRING)
	client.SetBytes(1, KeyId2, client.TYPE_STRING, []byte("Bleep"))
	s2 := client.GetBytes(1, KeyId2, client.TYPE_STRING)
	client.SetBytes(1, KeyId2, client.TYPE_STRING, []byte("Klunky"))
	s3 := client.GetBytes(1, KeyId2, client.TYPE_STRING)
	client.SetBytes(1, KeyId2, client.TYPE_STRING, s1)
	client.SetBytes(1, KeyId2, client.TYPE_STRING, s2)
	client.SetBytes(1, KeyId2, client.TYPE_STRING, s3)
}

func resultsTest(ctx *client.ScCallContext) {
	testMap(ctx.Results())
	checkMap(ctx.Results().Immutable())
	//ctx.Call(ctx.ContractId().Hname(), client.NewHname("results_check"), nil, nil)
}

func stateTest(ctx *client.ScCallContext) {
	testMap(ctx.State())
	ctx.Call(ctx.ContractId().Hname(), client.NewHname("state_check"), nil, nil)
}

func resultsCheck(ctx *client.ScViewContext) {
	checkMap(ctx.Results().Immutable())
}

func stateCheck(ctx *client.ScViewContext) {
	checkMap(ctx.State())
}

func testMap(kvstore client.ScMutableMap) {
	int1 := kvstore.GetInt(client.Key("int1"))
	check(int1.Value() == 0)
	int1.SetValue(1)

	string1 := kvstore.GetString(client.Key("string1"))
	check(string1.Value() == "")
	string1.SetValue("a")
	check(string1.Value() == "a")

	ia1 := kvstore.GetIntArray(client.Key("ia1"))
	int2 := ia1.GetInt(0)
	check(int2.Value() == 0)
	int2.SetValue(2)
	int3 := ia1.GetInt(1)
	check(int3.Value() == 0)
	int3.SetValue(3)

	sa1 := kvstore.GetStringArray(client.Key("sa1"))
	string2 := sa1.GetString(0)
	check(string2.Value() == "")
	string2.SetValue("bc")
	string3 := sa1.GetString(1)
	check(string3.Value() == "")
	string3.SetValue("def")
}

func checkMap(kvstore client.ScImmutableMap) {
	int1 := kvstore.GetInt(client.Key("int1"))
	check(int1.Value() == 1)

	string1 := kvstore.GetString(client.Key("string1"))
	check(string1.Value() == "a")

	ia1 := kvstore.GetIntArray(client.Key("ia1"))
	int2 := ia1.GetInt(0)
	check(int2.Value() == 2)
	int3 := ia1.GetInt(1)
	check(int3.Value() == 3)

	sa1 := kvstore.GetStringArray(client.Key("sa1"))
	string2 := sa1.GetString(0)
	check(string2.Value() == "bc")
	string3 := sa1.GetString(1)
	check(string3.Value() == "def")
}

func checkMapRev(kvstore client.ScImmutableMap) {
	sa1 := kvstore.GetStringArray(client.Key("sa1"))
	string3 := sa1.GetString(1)
	check(string3.Value() == "def")
	string2 := sa1.GetString(0)
	check(string2.Value() == "bc")

	ia1 := kvstore.GetIntArray(client.Key("ia1"))
	int3 := ia1.GetInt(1)
	check(int3.Value() == 3)
	int2 := ia1.GetInt(0)
	check(int2.Value() == 2)

	string1 := kvstore.GetString(client.Key("string1"))
	check(string1.Value() == "a")

	int1 := kvstore.GetInt(client.Key("int1"))
	check(int1.Value() == 1)
}

func check(condition bool) {
	if !condition {
		panic("Check failed!")
	}
}
