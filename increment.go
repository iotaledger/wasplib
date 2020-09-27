package main

import (
	"github.com/iotaledger/wasplib/client"
)

func main() {
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

//export increment
func increment() {
	ctx := client.NewScContext()
	counter := ctx.State().GetInt("counter")
	counter.SetValue(counter.Value() + 1)
}

//export incrementRepeat1
func incrementRepeat1() {
	ctx := client.NewScContext()
	counter := ctx.State().GetInt("counter")
	value := counter.Value()
	counter.SetValue(value + 1)
	if value == 0 {
		ctx.Event("", "increment", 5)
	}
}

//export incrementRepeatMany
func incrementRepeatMany() {
	ctx := client.NewScContext()
	counter := ctx.State().GetInt("counter")
	value := counter.Value()
	counter.SetValue(value + 1)
	repeats := ctx.Request().Params().GetInt("numrepeats").Value()
	stateRepeats := ctx.State().GetInt("numrepeats")
	if repeats == 0 {
		repeats = stateRepeats.Value()
		if repeats == 0 {
			return
		}
	}
	stateRepeats.SetValue(repeats - 1)
	ctx.Event("", "incrementRepeatMany", 3)
}
