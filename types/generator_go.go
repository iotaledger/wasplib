package types

import (
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"strings"
)

var goTypes = StringMap{
	"Address":     "*client.ScAddress",
	"Agent":       "*client.ScAgentId",
	"Chain_id":    "*client.ScChainId",
	"Color":       "*client.ScColor",
	"Contract_id": "*client.ScContractId",
	"Hash":        "*client.ScHash",
	"Hname":       "client.ScHname",
	"Int":         "int64",
	"String":      "string",
}

//TODO check for clashing Hnames

func GenerateGoSchema(path string, contract string, schema *Schema) error {
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

	fmt.Fprintf(file, "const ScName = \"%s\"\n", schema.Name)
	hName := coretypes.Hn(schema.Name)
	fmt.Fprintf(file, "const ScHname = client.ScHname(0x%s)\n", hName.String())

	fmt.Fprintln(file)
	params := make(StringMap)
	for _, funcDef := range schema.Funcs {
		for fldName := range funcDef {
			if !strings.HasPrefix(fldName, "#") {
				params[fldName] = fldName
			}
		}
	}
	for _, name := range sortedKeys(params) {
		fmt.Fprintf(file, "const Param%s = client.Key(\"%s\")\n", capitalize(name), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedKeys(schema.Vars) {
		fmt.Fprintf(file, "const Var%s = client.Key(\"%s\")\n", capitalize(name), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedMaps(schema.Funcs) {
		fmt.Fprintf(file, "const Func%s = \"%s\"\n", capitalize(name), name)
	}
	for _, name := range sortedMaps(schema.Views) {
		fmt.Fprintf(file, "const View%s = \"%s\"\n", capitalize(name), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedMaps(schema.Funcs) {
		hName = coretypes.Hn(name)
		fmt.Fprintf(file, "const HFunc%s = client.ScHname(0x%s)\n", capitalize(name), hName.String())
	}
	for _, name := range sortedMaps(schema.Views) {
		hName = coretypes.Hn(name)
		fmt.Fprintf(file, "const HView%s = client.ScHname(0x%s)\n", capitalize(name), hName.String())
	}

	fmt.Fprintf(file, "\nfunc OnLoad() {\n")
	fmt.Fprintf(file, "    exports := client.NewScExports()\n")
	for _, name := range sortedMaps(schema.Funcs) {
		name = capitalize(name)
		fmt.Fprintf(file, "    exports.AddCall(Func%s, func%s)\n", name, name)
	}
	for _, name := range sortedMaps(schema.Views) {
		name = capitalize(name)
		fmt.Fprintf(file, "    exports.AddView(View%s, view%s)\n", name, name)
	}
	fmt.Fprintf(file, "}\n")

	return nil
}

func GenerateGoTypes(path string, contract string, types StringMapMap) error {
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
	sortedTypes := sortedMaps(types)
	for _, typeName := range sortedTypes {
		fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
		fldDef := types[typeName]
		nameLen := 0
		typeLen := 0
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			goType, ok := goTypes[fldType]
			if !ok {
				return fmt.Errorf("invalid type name: %s", fldType)
			}
			if nameLen < len(fldName) {
				nameLen = len(fldName)
			}
			if typeLen < len(goType) {
				typeLen = len(goType)
			}
		}
		for _, fldName := range sortedKeys(fldDef) {
			fldType, comment := splitComment(fldDef[fldName])
			goType := pad(goTypes[fldType], typeLen)
			fldName = pad(capitalize(fldName), nameLen)
			fmt.Fprintf(file, "\t%s %s%s\n", fldName, goType, comment)
		}
		fmt.Fprintf(file, "}\n")
	}

	// write encoder and decoder for structs
	for _, typeName := range sortedTypes {
		fmt.Fprintf(file, "\nfunc Encode%s(o *%s) []byte {\n", typeName, typeName)
		fmt.Fprintf(file, "\treturn client.NewBytesEncoder().\n")
		fldDef := types[typeName]
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			fmt.Fprintf(file, "\t\t%s(o.%s).\n", fldType, capitalize(fldName))
		}
		fmt.Fprintf(file, "\t\tData()\n}\n")

		fmt.Fprintf(file, "\nfunc Decode%s(bytes []byte) *%s {\n", typeName, typeName)
		fmt.Fprintf(file, "\tdecode := client.NewBytesDecoder(bytes)\n\tdata := &%s{}\n", typeName)
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			fmt.Fprintf(file, "\tdata.%s = decode.%s()\n", capitalize(fldName), fldType)
		}
		fmt.Fprintf(file, "\treturn data\n}\n")
	}

	return nil
}

func GenerateGoCoreContractsSchema() error {
	coreSchemas, err := LoadCoreSchemas()
	if err != nil {
		return err
	}
	if coreSchemas == nil {
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

	for _, schema := range coreSchemas {
		scName := capitalize(schema.Name)
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\nconst Core%s = ScHname(0x%s)\n", scName, scHname.String())
		for _, funcName := range sortedMaps(schema.Funcs) {
			funcHname := coretypes.Hn(funcName)
			funcName = capitalize(funcName)
			fmt.Fprintf(file, "const Core%sFunc%s = ScHname(0x%s)\n", scName, funcName, funcHname.String())
		}
		for _, funcName := range sortedMaps(schema.Views) {
			funcHname := coretypes.Hn(funcName)
			funcName = capitalize(funcName)
			fmt.Fprintf(file, "const Core%sView%s = ScHname(0x%s)\n", scName, funcName, funcHname.String())
		}
	}
	return nil
}
