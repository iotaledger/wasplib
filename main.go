// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/iotaledger/wasplib/client"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/wasmhost"
	"io/ioutil"
)

// generate base58 strings
//var zero [34]byte
//// for i := range zero { zero[i] = 1 }
//for i := 0; i < 10; i++ {
//	zero[0] = byte(i)
//	fmt.Println(base58.Encode(zero[:]))
//}

//IOTA: 11111111111111111111111111111111
//MINT: JEKNVnkbo3jma5nREBBJCDoXFVeKkD56V3xKrvRmWxFG

//32:0: 11111111111111111111111111111111
//32:1: 4uQeVj5tqViQh7yWWGStvkEG1Zmhx6uasJtWCJziofM
//32:2: 8opHzTAnfzRpPEx21XtnrVTX28YQuCpAjcn1PczScKh
//32:3: CiDwVBFgWV9E5MvXWoLgnEgn2hK7rJikbvfWavzAQz3
//32:4: GcdayuLaLyrdmUu324nahyv33G5poQdLUEZ1nEytDeP
//32:5: LX3EUdRUBUa3TbsYXLEUdj9J3prXkWXvLYSWyYyc2Jj
//32:6: QRSsyMWN1yHT9ir42bgNZUNZ4PdEhcSWCrL2AryKpy5
//32:7: UKrXU5bFrTzrqqpZXs8GVDbp4xPweiM65ADXNAy3ddR
//32:8: YEGAxog9gxiGXxo538aAQxq55XAebpFfwU72ZUxmSHm
//32:9: c8fpTXm3XTRgE5maYQ24Li4L65wMYvAFomzXknxVEx7

//33:0: 111111111111111111111111111111111
//33:1: JEKNVnkbo3jma5nREBBJCDoXFVeKkD56V3xKrvRmWxFH
//33:2: bTdjzaWCb6UY9AZqTMMbPSc3VzHeVR9By6ueiqrY2uVZ
//33:3: tgx7VNFoP9DJiFMFgXXtafQZkUvyEdDHT9ryamHJYrjq
//33:4: 2BvGUzA1QBBx5HL8fuhiBmtD5zyaHyqHNwCpJSgi54oz7
//33:5: 2V9arUwkzyEgqrQv68stUy71cFUDcj3MURFmdJc8qamEP
//33:6: 2nNuDyjWbmHRcRVhWN44nAKp8VxrwUFRZuJixAXZc6iUf
//33:7: 35cDbUXGCZLANzaUvbEF5MYcekTWGDTVfPMgH2SzNcfiw
//33:8: 3NqXxyK1oMNu9ZfGLpQRNYmRAzx9axfZksQdbtNR98cyD
//33:9: 3g4rLU6mQ9Rdv8k3m3abfjzDhFSnuhsdrMTavkHqueaDV

//34:0: 1111111111111111111111111111111111
//34:1: 2K3n5t4wSaF5mj27Tw9vStXWLWyRjjiH5Cp3CFLpKVCr1d
//34:2: 3d6ZAm8st9VAYT3DvsJqtn41g2wrUURZ9Qd5PVgddyQh2F
//34:3: 4w9LFeCpKijFKB4LPoTmLfaX1YvHDD8qDcS7ak2SxTcY2s
//34:4: 6FC7LXGkmHyL5u5SrjcgnZ72M4thwwr7HpF9mzNGGwpP3V
//34:5: 7ZEtRQLhCsDQrd6ZKfmcESdXgas8ggZPN24ByEi5bS2E47
//34:6: 8sHfWHQdeSTVdM7fnbvXgLA326qZRRGfSDsEAV3tuvE54j
//34:7: ABLSbAUa61haQ58nFY5T8DgYMcozA9ywWRgGMjPiEQRv5M
//34:8: BVPDg3YWXawfAo9tiUENa7D3h8nQtthDadVJYyjXYtdm5y
//34:9: CoRzkvcSyABjwXB1BQPJ1zjZ2ekqddQVeqJLkE5LsNqc6b

func main() {
	fmt.Println("Hello, WaspLib!")

	//execGoHost()

	//err := filepath.Walk("../tests",
	//	func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		if strings.HasSuffix(path, ".json") {
	//			readerTest(path)
	//		}
	//		return nil
	//	})
	//if err != nil {
	//	log.Println(err)
	//}

	execJsonTest()
}

func execGoHost() {
	goHost, err := wasmhost.NewSimpleWasmHost()
	if err != nil {
		panic(err)
	}
	client.ConnectHost(goHost)
	inccounter.OnLoad()
	goHost.RunScFunction("incrementCallIncrement")
}

func execJsonTest() {
	//contract := "donatewithfeedback"
	//contract := "fairauction"
	//contract := "fairroulette"
	contract := "inccounter"
	//contract := "tokenregistry"
	language := "go" // "bg" = Rust, "go" = Go

	pathName := "tests/" + contract + ".json"
	jsonTests, err := wasmhost.NewJsonTests(pathName)
	if err != nil {
		panic(err)
	}

	host, err := wasmhost.NewSimpleWasmHost()
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

	failed := 0
	passed := 0
	for _, test := range jsonTests.Tests {
		if jsonTests.RunTest(&host.WasmHost, test) {
			fmt.Printf("PASS\n")
			passed++
			continue
		}
		failed++
	}
	if failed != 0 {
		fmt.Printf("%d FAILED, ", failed)
	}
	fmt.Printf("%d PASSED\n", passed)
}
