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
	exports.Add("increment")
	exports.Add("incrementRepeat1")
	exports.Add("incrementRepeatMany")
	exports.Add("test")
	exports.Add("nothing")
	exports.Add("init")
}

//export init
func init() {
	sc := client.NewScContext()
	counter := sc.Request().Params().GetInt("counter").Value()
	if counter == 0 {
		return
	}
	sc.State().GetInt("counter").SetValue(counter)
}

//export increment
func increment() {
	sc := client.NewScContext()
	counter := sc.State().GetInt("counter")
	counter.SetValue(counter.Value() + 1)
}

//export incrementRepeat1
func incrementRepeat1() {
	sc := client.NewScContext()
	counter := sc.State().GetInt("counter")
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		sc.PostRequest(sc.Contract().Id(), "increment", 0)
	}
}

//export incrementRepeatMany
func incrementRepeatMany() {
	sc := client.NewScContext()
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
	sc.PostRequest(sc.Contract().Id(), "incrementRepeatMany", 0)
}

//export test
func test() {
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
