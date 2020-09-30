package main

import (
	"encoding/json"
	"fmt"
	"github.com/iotaledger/wasplib/wasmhost"
	"os"
)

func main() {
	fmt.Println("Hello, WaspLib!")

	file, err := os.Open("increment.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jsonTests := &wasmhost.JsonTests{}
	err = json.NewDecoder(file).Decode(&jsonTests)
	if err != nil {
		panic(err)
	}

	host := wasmhost.NewHostImpl()
	err = host.LoadWasm("wasm/increment_bg.wasm")
	if err != nil {
		panic(err)
	}

	for name, jsonModel := range jsonTests.Tests {
		host.RunTest(name, jsonModel, jsonTests)
	}
}
