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

var flagInit = flag.String("init", "", "generate Go code")

var flagGo = flag.Bool("go", false, "generate Go code")
var flagJava = flag.Bool("java", false, "generate Java code")
var flagRust = flag.Bool("rust", false, "generate Rust code")

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

	// tool is also used to (re-)generate the core contract
	// definitions inside the go and rust sections of wasmlib
	file, err = os.Open("corecontracts.json")
	if err == nil {
		defer file.Close()
		err = generateCoreContractsSchema(file)
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

func generateCoreContractsSchema(file *os.File) error {
	coreSchemas, err := loadCoreSchemas(file)
	if err != nil {
		return err
	}
	err = generator.GenerateRustCoreContractsSchema(coreSchemas)
	if err != nil {
		return err
	}
	err = generator.GenerateGoCoreContractsSchema(coreSchemas)
	if err != nil {
		return err
	}
	err = generator.GenerateJavaCoreContractsSchema(coreSchemas)
	if err != nil {
		return err
	}
	return nil
}

func generateSchema(file *os.File) error {
	schema, err := loadSchema(file)
	if err != nil {
		return err
	}
	if *flagGo {
		fmt.Println("generating Go code")
		err = schema.GenerateGo()
		if err != nil {
			return err
		}
		err = schema.GenerateGoTests()
		if err != nil {
			return err
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
		err = schema.GenerateGoTests()
		if err != nil {
			return err
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

	jsonSchema := &generator.JsonSchema{}
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
	funcInit.Params["owner"] = "?AgentId // optional owner of this smart contract"
	jsonSchema.Funcs["init"] = funcInit

	b, err := json.Marshal(jsonSchema)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	_, err = out.WriteTo(file)
	return err
}

func loadSchema(file *os.File) (*generator.Schema, error) {
	fmt.Println("loading schema.json")
	jsonSchema := &generator.JsonSchema{}
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

func loadCoreSchemas(file *os.File) ([]*generator.Schema, error) {
	fmt.Println("loading corecontracts.json")
	coreJsonSchemas := make([]*generator.JsonSchema, 0)
	err := json.NewDecoder(file).Decode(&coreJsonSchemas)
	if err != nil {
		return nil, err
	}

	coreSchemas := make([]*generator.Schema, 0)
	for _, jsonSchema := range coreJsonSchemas {
		schema := generator.NewSchema()
		err = schema.Compile(jsonSchema)
		if err != nil {
			return nil, err
		}
		coreSchemas = append(coreSchemas, schema)
	}
	return coreSchemas, nil
}
