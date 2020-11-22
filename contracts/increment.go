// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/iotaledger/wasplib/client"
)

func main() {
}

//export onLoad
func onLoadIncrement() {
	exports := client.NewScExports()
	exports.AddCall("increment", increment)
	exports.AddCall("incrementRepeat1", incrementRepeat1)
	exports.AddCall("incrementRepeatMany", incrementRepeatMany)
	exports.AddCall("test", test)
	exports.AddCall("nothing", client.Nothing)
	exports.AddCall("init", onInitIncrement)
}

func onInitIncrement(sc *client.ScCallContext) {
	counter := sc.Request().Params().GetInt("counter").Value()
	if counter == 0 {
		return
	}
	sc.State().GetInt("counter").SetValue(counter)
}

func increment(sc *client.ScCallContext) {
	counter := sc.State().GetInt("counter")
	counter.SetValue(counter.Value() + 1)
}

func incrementRepeat1(sc *client.ScCallContext) {
	counter := sc.State().GetInt("counter")
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		sc.PostSelf("increment", 0)
	}
}

func incrementRepeatMany(sc *client.ScCallContext) {
	counter := sc.State().GetInt("counter")
	value := counter.Value()
	counter.SetValue(value + 1)
	stateRepeats := sc.State().GetInt("numRepeats")
	repeats := sc.Request().Params().GetInt("numRepeats").Value()
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
	sc.PostSelf("incrementRepeatMany", 0)
}

func test(sc *client.ScCallContext) {
	keyId := client.GetKeyId("timestamp")
	client.SetInt(1, keyId, 123456789)
	timestamp := client.GetInt(1, keyId)
	client.SetInt(1, keyId, timestamp)

	keyId2 := client.GetKeyId("string")
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
