// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyCounter    = client.Key("counter")
	keyNumRepeats = client.Key("num_repeats")
)

var localStateMustIncrement = false

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

func onInit(sc *client.ScCallContext) {
	counter := sc.Params().GetInt(keyCounter).Value()
	if counter == 0 {
		return
	}
	sc.State().GetInt(keyCounter).SetValue(counter)
}

func increment(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	counter.SetValue(counter.Value() + 1)
}

func incrementCallIncrement(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		sc.Call("increment_call_increment").Call()
	}
}

func incrementCallIncrementRecurse5x(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		sc.Call("increment_call_increment_recurse5x").Call()
	}
}

func incrementPostIncrement(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		sc.Post("increment_post_increment").Post(0)
	}
}

func incrementViewCounter(sc *client.ScViewContext) {
	counter := sc.State().GetInt(keyCounter).Value()
	sc.Results().GetInt(keyCounter).SetValue(counter)
}

func incrementRepeatMany(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := sc.State().GetInt(keyNumRepeats)
	repeats := sc.Params().GetInt(keyNumRepeats).Value()
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
	sc.Post("increment_repeat_many").Post(0)
}

func incrementWhenMustIncrement(sc *client.ScCallContext) {
	sc.Log("increment_when_must_increment called")
	if localStateMustIncrement {
		counter := sc.State().GetInt(keyCounter)
		counter.SetValue(counter.Value() + 1)
	}
}

func incrementLocalStateInternalCall(sc *client.ScCallContext) {
	incrementWhenMustIncrement(sc)
	localStateMustIncrement = true
	incrementWhenMustIncrement(sc)
	incrementWhenMustIncrement(sc)
	// counter ends up as 2
}

func incrementLocalStateSandboxCall(sc *client.ScCallContext) {
	sc.Call("increment_when_must_increment").Call()
	localStateMustIncrement = true
	sc.Call("increment_when_must_increment").Call()
	sc.Call("increment_when_must_increment").Call()
	// counter ends up as 0
}

func incrementLocalStatePost(sc *client.ScCallContext) {
	sc.Post("increment_when_must_increment").Post(0)
	localStateMustIncrement = true
	sc.Post("increment_when_must_increment").Post(0)
	sc.Post("increment_when_must_increment").Post(0)
	// counter ends up as 0
}

func test(sc *client.ScCallContext) {
	keyId := client.GetKeyIdFromString("timestamp")
	client.SetInt(1, keyId, 123456789)
	timestamp := client.GetInt(1, keyId)
	client.SetInt(1, keyId, timestamp)

	keyId2 := client.GetKeyIdFromString("string")
	client.SetString(1, keyId2, "Test")
	s1 := client.GetString(1, keyId2)
	client.SetString(1, keyId2, "Bleep")
	s2 := client.GetString(1, keyId2)
	client.SetString(1, keyId2, "Klunky")
	s3 := client.GetString(1, keyId2)
	client.SetString(1, keyId2, s1)
	client.SetString(1, keyId2, s2)
	client.SetString(1, keyId2, s3)
}

func resultsTest(sc *client.ScCallContext) {
	testKvstore(sc.Results())
	checkKvstore(sc.Results().Immutable())
	//sc.call("results_check")
}

func stateTest(sc *client.ScCallContext) {
	testKvstore(sc.State())
	sc.Call("state_check")
}

func resultsCheck(sc *client.ScViewContext) {
	checkKvstore(sc.Results().Immutable())
}

func stateCheck(sc *client.ScViewContext) {
	checkKvstore(sc.State())
}

func testKvstore(kvstore client.ScMutableMap) {
	int1 := kvstore.GetInt(client.Key("int1"))
	check(int1.Value() == 0)
	int1.SetValue(1)

	string1 := kvstore.GetString(client.Key("string1"))
	check(string1.Value() == "")
	string1.SetValue("a")

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

func checkKvstore(kvstore client.ScImmutableMap) {
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

func check(condition bool) {
	if !condition {
		panic("Check failed!")
	}
}
