// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/iotaledger/wasplib/tools/schema/generator"
)

var flagCore = flag.Bool("core", false, "generate core contract interface")

var flagInit = flag.String("init", "", "generate Go code")

var (
	flagGo   = flag.Bool("go", false, "generate Go code")
	flagJava = flag.Bool("java", false, "generate Java code")
	flagRust = flag.Bool("rust", false, "generate Rust code")
)

func main() {
	flag.Parse()
	if !(*flagGo || *flagJava) {
		*flagRust = true
	}

	err := generator.FindModulePath()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open("schema.json")
	if err == nil {
		defer file.Close()
		err = generateSchema(file)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *flagInit != "" {
		err = generateSchemaNew()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	flag.Usage()
}

func generateSchema(file *os.File) error {
	schema, err := loadSchema(file)
	if err != nil {
		return err
	}
	schema.CoreContracts = *flagCore
	if *flagGo {
		fmt.Println("generating Go code")
		err = schema.GenerateGo()
		if err != nil {
			return err
		}
		if !schema.CoreContracts {
			err = schema.GenerateGoTests()
			if err != nil {
				return err
			}
		}
	}
	if *flagJava {
		fmt.Println("generating Java code")
		err = schema.GenerateJava()
		if err != nil {
			return err
		}
	}
	if *flagRust {
		fmt.Println("generating Rust code")
		err = schema.GenerateRust()
		if err != nil {
			return err
		}
		if !schema.CoreContracts {
			err = schema.GenerateGoTests()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateSchemaNew() error {
	fmt.Println("generating schema.json")
	file, err := os.Create("schema.json")
	if err != nil {
		return err
	}
	defer file.Close()

	jsonSchema := &generator.JSONSchema{}
	jsonSchema.Name = *flagInit
	jsonSchema.Description = *flagInit + " description"
	jsonSchema.Types = make(generator.StringMapMap)
	jsonSchema.Subtypes = make(generator.StringMap)
	jsonSchema.State = make(generator.StringMap)
	jsonSchema.Funcs = make(generator.FuncDescMap)
	jsonSchema.Views = make(generator.FuncDescMap)
	funcInit := &generator.FuncDesc{}
	funcInit.Params = make(generator.StringMap)
	funcInit.Results = make(generator.StringMap)
	funcInit.Params["owner"] = "?AgentID // optional owner of this smart contract"
	jsonSchema.Funcs["init"] = funcInit

	b, err := json.Marshal(jsonSchema)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return err
	}

	_, err = out.WriteTo(file)
	return err
}

func loadSchema(file *os.File) (*generator.Schema, error) {
	fmt.Println("loading schema.json")
	jsonSchema := &generator.JSONSchema{}
	err := json.NewDecoder(file).Decode(jsonSchema)
	if err != nil {
		return nil, err
	}

	schema := generator.NewSchema()
	err = schema.Compile(jsonSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}
