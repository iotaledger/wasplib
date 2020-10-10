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

	host := &wasmhost.SimpleWasmHost{}
	err = host.Init(wasmhost.NewNullObject(host), wasmhost.NewHostMap(host), nil, host)
	if err != nil {
		panic(err)
	}
	host.ExportsId = host.GetKeyId("exports")
	err = host.LoadWasmFile("wasm/increment_bg.wasm")
	if err != nil {
		panic(err)
	}
	err = host.RunWasmFunction("onLoad")
	if err != nil {
		panic(err)
	}

	for name, jsonModel := range jsonTests.Tests {
		host.RunTest(name, jsonModel, jsonTests)
	}
}
