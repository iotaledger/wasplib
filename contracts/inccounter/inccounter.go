// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyCounter    = client.Key("counter")
	keyNumRepeats = client.Key("numRepeats")
)

var localStateMustIncrement = false

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddCall("increment", increment)
	exports.AddCall("incrementCallIncrement", incrementCallIncrement)
	exports.AddCall("incrementCallIncrementRecurse5x", incrementCallIncrementRecurse5x)
	exports.AddCall("incrementPostIncrement", incrementPostIncrement)
	exports.AddView("incrementViewCounter", incrementViewCounter)
	exports.AddCall("incrementRepeatMany", incrementRepeatMany)
	exports.AddCall("incrementWhenMustIncrement", incrementWhenMustIncrement)
	exports.AddCall("incrementLocalStateInternalCall", incrementLocalStateInternalCall)
	exports.AddCall("incrementLocalStateSandboxCall", incrementLocalStateSandboxCall)
	exports.AddCall("incrementLocalStatePost", incrementLocalStatePost)
	exports.AddCall("nothing", client.Nothing)
	exports.AddCall("test", test)
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
		sc.Call("incrementCallIncrement").Call()
	}
}

func incrementCallIncrementRecurse5x(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value < 5 {
		sc.Call("incrementCallIncrementRecurse5x").Call()
	}
}

func incrementPostIncrement(sc *client.ScCallContext) {
	counter := sc.State().GetInt(keyCounter)
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		sc.Post("incrementPostIncrement").Post(0)
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
	sc.Post("incrementRepeatMany").Post(0)
}

func incrementWhenMustIncrement(sc *client.ScCallContext) {
	sc.Log("incrementWhenMustIncrement called")
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
	sc.Call("incrementWhenMustIncrement").Call()
	localStateMustIncrement = true
	sc.Call("incrementWhenMustIncrement").Call()
	sc.Call("incrementWhenMustIncrement").Call()
	// counter ends up as 0
}

func incrementLocalStatePost(sc *client.ScCallContext) {
	sc.Post("incrementWhenMustIncrement").Post(0)
	localStateMustIncrement = true
	sc.Post("incrementWhenMustIncrement").Post(0)
	sc.Post("incrementWhenMustIncrement").Post(0)
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
