// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/iotaledger/wasplib/contracts/dividend"
	"github.com/iotaledger/wasplib/contracts/donatewithfeedback"
	"github.com/iotaledger/wasplib/contracts/fairauction"
	"github.com/iotaledger/wasplib/contracts/fairroulette"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/contracts/tokenregistry"
	"github.com/iotaledger/wasplib/wasmhost"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var scForGoVM = map[string]func(){
	"dividend":           dividend.OnLoad,
	"donatewithfeedback": donatewithfeedback.OnLoad,
	"fairauction":        fairauction.OnLoad,
	"fairroulette":       fairroulette.OnLoad,
	"inccounter":         inccounter.OnLoad,
	"tokenregistry":      tokenregistry.OnLoad,
}
var failed = 0
var passed = 0

func main() {
	fmt.Println("Hello, WaspLib!")
	execJsonTests()
}

func execJsonTest(path string) {
	execTest(path, "sc") // Go VM
	execTest(path, "bg") // Rust Wasm VM
	execTest(path, "go") // Go Wasm VM
}

func execJsonTests() {
	err := filepath.Walk("tests",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".json") {
				execJsonTest(path)
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	if failed != 0 {
		fmt.Printf("%d FAILED, ", failed)
	}
	fmt.Printf("%d PASSED\n", passed)
}

func execTest(path string, language string) {
	contract := path[6 : len(path)-5]
	host := setupVM(contract, language)

	jsonTests, err := wasmhost.NewJsonTests(path)
	if err != nil {
		panic(err)
	}

	for _, test := range jsonTests.Tests {
		if jsonTests.RunTest(&host.WasmHost, test) {
			fmt.Printf("PASS\n")
			passed++
			continue
		}
		if language == "sc" && strings.Contains(test.Flags, "failWhenSC") {
			fmt.Printf("PASS (fail)\n")
			passed++
			continue
		}
		failed++
	}
}

func setupVM(contract string, language string) *wasmhost.SimpleWasmHost {
	if language == "sc" {
		host, err := wasmhost.NewSimpleWasmHost(wasmhost.NewGoVM())
		if err != nil {
			panic(err)
		}
		onLoad, ok := scForGoVM[contract]
		if !ok {
			panic("Unknown contract: " + contract)
		}
		onLoad()
		return host
	}

	host, err := wasmhost.NewSimpleWasmHost(wasmhost.NewWasmTimeVM())
	if err != nil {
		panic(err)
	}
	wasmData, err := ioutil.ReadFile("wasm/" + contract + "_" + language + ".wasm")
	if err != nil {
		panic(err)
	}
	err = host.LoadWasm(wasmData)
	if err != nil {
		panic(err)
	}
	return host
}
