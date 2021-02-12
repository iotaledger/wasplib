// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"github.com/iotaledger/wasp/tools/schema/generator"
	"os"
)

func main() {
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

	fmt.Println("no schema file found")
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
	return nil
}

func generateSchema(file *os.File) error {
	schema, err := loadSchema(file)
	if err != nil {
		return err
	}
	err = schema.GenerateRust()
	if err != nil {
		return err
	}
	return schema.GenerateGo()
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
