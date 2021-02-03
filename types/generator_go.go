package types

import (
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"strings"
)

var goTypes = map[string]string{
	"address":     "*client.ScAddress",
	"agent":       "*client.ScAgentId",
	"chain_id":    "*client.ScChainId",
	"color":       "*client.ScColor",
	"contract_id": "*client.ScContractId",
	"hash":        "*client.ScHash",
	"hname":       "client.ScHname",
	"int":         "int64",
	"string":      "string",
}

//TODO check for clashing Hnames

func GenerateGoSchema(path string, contract string, gen *Generator) error {
	file, err := os.Create(path + "schema.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\npackage %s\n\n", contract)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n\n")

	fmt.Fprintf(file, "const ScName = \"%s\"\n", gen.schema.Name)
	hName := coretypes.Hn(gen.schema.Name)
	fmt.Fprintf(file, "const ScHname = client.ScHname(0x%s)\n", hName.String())

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Params) {
		value := gen.schema.Params[name]
		fmt.Fprintf(file, "const Param%s = client.Key(\"%s\")\n", name, value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Vars) {
		value := gen.schema.Vars[name]
		fmt.Fprintf(file, "const Var%s = client.Key(\"%s\")\n", name, value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Funcs) {
		value := gen.schema.Funcs[name]
		fmt.Fprintf(file, "const Func%s = \"%s\"\n", name, value)
	}
	for _, name := range sorted(gen.schema.Views) {
		value := gen.schema.Views[name]
		fmt.Fprintf(file, "const View%s = \"%s\"\n", name, value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Funcs) {
		value := gen.schema.Funcs[name]
		hName = coretypes.Hn(value)
		fmt.Fprintf(file, "const HFunc%s = client.ScHname(0x%s)\n", name, hName.String())
	}
	for _, name := range sorted(gen.schema.Views) {
		value := gen.schema.Views[name]
		hName = coretypes.Hn(value)
		fmt.Fprintf(file, "const HView%s = client.ScHname(0x%s)\n", name, hName.String())
	}

	fmt.Fprintf(file, "\nfunc OnLoad() {\n")
	fmt.Fprintf(file, "    exports := client.NewScExports()\n")
	for _, name := range sorted(gen.schema.Funcs) {
		fmt.Fprintf(file, "    exports.AddCall(Func%s, func%s)\n", name, name)
	}
	for _, name := range sorted(gen.schema.Views) {
		fmt.Fprintf(file, "    exports.AddView(View%s, view%s)\n", name, name)
	}
	fmt.Fprintf(file, "}\n")

	return nil
}

func GenerateGoTypes(path string, contract string, gen *Generator) error {
	types := gen.schema.Types
	if len(types) == 0 {
		return nil
	}

	file, err := os.Create(path + "types.go")
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

func GenerateGoCoreContractsSchema() error {
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
		fmt.Fprintf(file, "\nconst Core%s = ScHname(0x%s)\n", nContract, hContract.String())
		for _, nFunc := range sorted(schema.Funcs) {
			funcName := schema.Funcs[nFunc]
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "const Core%s%s = ScHname(0x%s)\n", nContract, nFunc, hFunc.String())
		}
		for _, nFunc := range sorted(schema.Views) {
			funcName := schema.Views[nFunc]
			hFunc := coretypes.Hn(funcName)
			fmt.Fprintf(file, "const Core%s%s = ScHname(0x%s)\n", nContract, nFunc, hFunc.String())
		}
	}
	return nil
}
