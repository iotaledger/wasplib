// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/iotaledger/wasplib/tools/schema/generator"
)

var flagCore = flag.Bool("core", false, "generate core contract interface")

var flagInit = flag.String("init", "", "generate new schema.json for smart contract named <string>")

var (
	disabledFlag = false
	flagGo       = flag.Bool("go", false, "generate Go code")
	flagJava     = &disabledFlag // flag.Bool("java", false, "generate Java code <outdated>")
	flagRust     = flag.Bool("rust", false, "generate Rust code <default>")
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
		if *flagInit != "" {
			fmt.Println("schema.json already exists")
			return
		}
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
	info, err := file.Stat()
	if err != nil {
		return err
	}
	schemaTime := info.ModTime()

	schema, err := loadSchema(file)
	if err != nil {
		return err
	}

	schema.CoreContracts = *flagCore
	if *flagGo {
		info, err := os.Stat("consts.go")
		if err == nil && info.ModTime().After(schemaTime) {
			fmt.Println("skipping Go code generation")
		} else {
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
	}
	if *flagJava {
		fmt.Println("generating Java code")
		err = schema.GenerateJava()
		if err != nil {
			return err
		}
	}
	if *flagRust {
		info, err := os.Stat("src/consts.rs")
		if err == nil && info.ModTime().After(schemaTime) {
			fmt.Println("skipping Rust code generation")
		} else {
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
	}
	return nil
}

func generateSchemaNew() error {
	name := *flagInit
	fmt.Println("initializing " + name)

	subfolder := strings.ToLower(name)
	err := os.Mkdir(subfolder, 0755)
	if err != nil {
		return err
	}
	err = os.Chdir(subfolder)
	if err != nil {
		return err
	}

	file, err := os.Create("schema.json")
	if err != nil {
		return err
	}
	defer file.Close()

	jsonSchema := &generator.JSONSchema{}
	jsonSchema.Name = name
	jsonSchema.Description = name + " description"
	jsonSchema.Structs = make(generator.StringMapMap)
	jsonSchema.Typedefs = make(generator.StringMap)
	jsonSchema.State = make(generator.StringMap)
	jsonSchema.State["owner"] = "AgentID // current owner of this smart contract"
	jsonSchema.Funcs = make(generator.FuncDescMap)
	jsonSchema.Views = make(generator.FuncDescMap)

	funcInit := &generator.FuncDesc{}
	funcInit.Params = make(generator.StringMap)
	funcInit.Params["owner"] = "?AgentID // optional owner of this smart contract"
	jsonSchema.Funcs["init"] = funcInit

	funcSetOwner := &generator.FuncDesc{}
	funcSetOwner.Access = "owner // current owner of this smart contract"
	funcSetOwner.Params = make(generator.StringMap)
	funcSetOwner.Params["owner"] = "AgentID // new owner of this smart contract"
	jsonSchema.Funcs["setOwner"] = funcSetOwner

	viewGetOwner := &generator.FuncDesc{}
	viewGetOwner.Results = make(generator.StringMap)
	viewGetOwner.Results["owner"] = "AgentID // current owner of this smart contract"
	jsonSchema.Views["getOwner"] = viewGetOwner

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
