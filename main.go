package main

import (
	"encoding/json"
	"fmt"
	"github.com/iotaledger/wasplib/host"
	"github.com/iotaledger/wasplib/jsontest"
	"os"
)

func main() {
	fmt.Println("Hello, WaspLib!")

	file, err := os.Open("increment.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	testData := &jsontest.JsonTest{}
	err = json.NewDecoder(file).Decode(&testData)
	if err != nil {
		panic(err)
	}

	ctx := host.NewHostImpl()
	err = ctx.LoadWasm("wasm/increment_bg.wasm")
	if err != nil {
		panic(err)
	}

	for name, t := range testData.Tests {
		ctx.RunTest(name, t, testData)
	}
}
