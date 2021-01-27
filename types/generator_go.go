package types

import (
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"regexp"
	"sort"
	"strings"
)

var goTypes = map[string]string{
	"address":     "*client.ScAddress",
	"agent":       "*client.ScAgent",
	"chain_id":    "*client.ScChainId",
	"color":       "*client.ScColor",
	"contract_id": "*client.ScContractId",
	"hash":        "*client.ScHash",
	"hname":       "client.Hname",
	"int":         "int64",
	"string":      "string",
}

func GenerateGoTypes(path string) error {
	gen := &Generator{}
	err := gen.LoadTypes(path)
	if err != nil {
		return err
	}

	var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wschema.json")
	contract := matchContract.ReplaceAllString(path, "$1")

	file, err := os.Create(path[:len(path)-len("schema.json")] + "types.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\npackage %s\n\n", contract)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n")

	// write structs
	types := gen.schema.Types
	for _, structName := range gen.keys {
		gen.SplitComments(structName, goTypes)
		spaces := strings.Repeat(" ", gen.maxName+gen.maxType)
		fmt.Fprintf(file, "\ntype %s struct {\n", structName)
		for _, fld := range types[structName] {
			for name, _ := range fld {
				camel := gen.camels[name]
				comment := gen.comments[name]
				goType := gen.types[name]
				if comment != "" {
					comment = spaces[:gen.maxType-len(goType)] + comment
				}
				goType = spaces[:gen.maxCamel-len(camel)] + goType
				fmt.Fprintf(file, "\t%s %s%s\n", camel, goType, comment)
			}
		}
		fmt.Fprintf(file, "}\n")
	}

	//  write encoder and decoder for structs
	for _, structName := range gen.keys {
		funcName := "code" + structName
		fmt.Fprintf(file, "\nfunc En%s(o *%s) []byte {\n", funcName, structName)
		fmt.Fprintf(file, "\treturn client.NewBytesEncoder().\n")
		for _, fld := range types[structName] {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "\t\t%s(o.%s).\n", typeName, camelcase(name))
			}
		}
		fmt.Fprintf(file, "\t\tData()\n}\n")

		fmt.Fprintf(file, "\nfunc De%s(bytes []byte) *%s {\n", funcName, structName)
		fmt.Fprintf(file, "\tdecode := client.NewBytesDecoder(bytes)\n\tdata := &%s{}\n", structName)
		for _, fld := range types[structName] {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "\tdata.%s = decode.%s()\n", camelcase(name), typeName)
			}
		}
		fmt.Fprintf(file, "\treturn data\n}\n")
	}

	//TODO write on_types function

	return nil
}

func GenerateGoCoreSchema() error {
	core, err := LoadCoreSchema()
	if err != nil {
		return err
	}
	if core == nil {
		return errors.New("missing core schema")
	}

	file, err := os.Create("../client/corecontracts.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\npackage client\n")

	for _, schema := range core {
		nContract := camelcase(schema.Name)
		hContract := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\nconst Core%s = Hname(0x%s)\n", nContract, hContract.String())
		for _, nFunc := range sorted(schema.Funcs) {
			funcName := schema.Funcs[nFunc]
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "const Core%s%s = Hname(0x%s)\n", nContract, nFunc, hFunc.String())
		}
		for _, nFunc := range sorted(schema.Views) {
			funcName := schema.Views[nFunc]
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "const Core%s%s = Hname(0x%s)\n", nContract, nFunc, hFunc.String())
		}
	}
	return nil
}

func GenerateRustCoreSchema() error {
	core, err := LoadCoreSchema()
	if err != nil {
		return err
	}
	if core == nil {
		return errors.New("missing core schema")
	}

	file, err := os.Create("../rust/wasplib/src/client/corecontracts.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\nuse super::hashtypes::*;\n")

	for _, schema := range core {
		nContract := snakecase(schema.Name)
		hContract := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\npub const CORE_%s: Hname = Hname(0x%s);\n", nContract, hContract.String())
		for _, nFunc := range sorted(schema.Funcs) {
			funcName := schema.Funcs[nFunc]
			nFunc = snakecase(nFunc)
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "pub const CORE_%s_%s: Hname = Hname(0x%s);\n", nContract, nFunc, hFunc.String())
		}
		for _, nFunc := range sorted(schema.Views) {
			funcName := schema.Views[nFunc]
			nFunc = snakecase(nFunc)
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "pub const CORE_%s_%s: Hname = Hname(0x%s);\n", nContract, nFunc, hFunc.String())
		}
	}
	return nil
}

func sorted(dict map[string]string) []string {
	keys := make([]string, 0)
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys

}
