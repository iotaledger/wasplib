package types

import (
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"strings"
)

var rustTypes = StringMap{
	"Address":     "ScAddress",
	"Agent":       "ScAgentId",
	"Chain_id":    "ScChainId",
	"Color":       "ScColor",
	"Contract_id": "ScContractId",
	"Hash":        "ScHash",
	"Hname":       "ScHname",
	"Int":         "i64",
	"String":      "String",
}

func GenerateRustSchema(path string, contract string, schema *Schema) error {
	file, err := os.Create(path + "schema.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n\n")
	fmt.Fprintf(file, "#![allow(dead_code)]\n\n")
	fmt.Fprintf(file, "use wasplib::client::*;\n")
	fmt.Fprintf(file, "use super::*;\n\n")

	fmt.Fprintf(file, "pub const SC_NAME: &str = \"%s\";\n", schema.Name)
	hName := coretypes.Hn(schema.Name)
	fmt.Fprintf(file, "pub const SC_HNAME: ScHname = ScHname(0x%s);\n", hName.String())

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
		fmt.Fprintf(file, "pub const PARAM_%s: &str = \"%s\";\n", upper(snake(name)), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedKeys(schema.Vars) {
		fmt.Fprintf(file, "pub const VAR_%s: &str = \"%s\";\n", upper(snake(name)), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedMaps(schema.Funcs) {
		fmt.Fprintf(file, "pub const FUNC_%s: &str = \"%s\";\n", upper(snake(name)), name)
	}
	for _, name := range sortedMaps(schema.Views) {
		fmt.Fprintf(file, "pub const VIEW_%s: &str = \"%s\";\n", upper(snake(name)), name)
	}

	fmt.Fprintln(file)
	for _, name := range sortedMaps(schema.Funcs) {
		hName = coretypes.Hn(name)
		fmt.Fprintf(file, "pub const HFUNC_%s: ScHname = ScHname(0x%s);\n", upper(snake(name)), hName.String())
	}
	for _, name := range sortedMaps(schema.Views) {
		hName = coretypes.Hn(name)
		fmt.Fprintf(file, "pub const HVIEW_%s: ScHname = ScHname(0x%s);\n", upper(snake(name)), hName.String())
	}

	fmt.Fprintf(file, "\n#[no_mangle]\n")
	fmt.Fprintf(file, "fn on_load() {\n")
	fmt.Fprintf(file, "    let exports = ScExports::new();\n")
	for _, name := range sortedMaps(schema.Funcs) {
		name = snake(name)
		fmt.Fprintf(file, "    exports.add_call(FUNC_%s, func_%s);\n", upper(name), name)
	}
	for _, name := range sortedMaps(schema.Views) {
		name = snake(name)
		fmt.Fprintf(file, "    exports.add_view(VIEW_%s, view_%s);\n", upper(name), name)
	}
	fmt.Fprintf(file, "}\n")

	return nil
}

func GenerateRustTypes(path string, contract string, types StringMapMap) error {
	if len(types) == 0 {
		return nil
	}

	file, err := os.Create(path + "types.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n\n")
	fmt.Fprintf(file, "use wasplib::client::*;\n")

	// write structs
	sortedTypes := sortedMaps(types)
	for _, typeName := range sortedTypes {
		fmt.Fprintf(file, "\npub struct %s {\n", typeName)
		fmt.Fprintf(file, "    //@formatter:off\n")
		fldDef := types[typeName]
		nameLen := 0
		typeLen := 0
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			rustType, ok := rustTypes[fldType]
			if !ok {
				return fmt.Errorf("invalid type name: %s", fldType)
			}
			fldName = snake(fldName)
			if nameLen < len(fldName) { nameLen = len(fldName) }
			if typeLen < len(rustType) { typeLen = len(rustType) }
		}
		for _, fldName := range sortedKeys(fldDef) {
			fldType, comment := splitComment(fldDef[fldName])
			rustType := pad(rustTypes[fldType] + ",", typeLen+1)
			fldName = pad(snake(fldName) + ":", nameLen+1)
			fmt.Fprintf(file, "    pub %s %s%s\n", fldName, rustType, comment)
		}
		fmt.Fprintf(file, "    //@formatter:on\n")
		fmt.Fprintf(file, "}\n")
	}

	// write encoder and decoder for structs
	for _, typeName := range sortedTypes {
		funcName := lower(snake(typeName))
		fmt.Fprintf(file, "\npub fn encode_%s(o: &%s) -> Vec<u8> {\n", funcName, typeName)
		fmt.Fprintf(file, "    let mut encode = BytesEncoder::new();\n")
		fldDef := types[typeName]
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			ref := "&"
			if fldType == "Int" || fldType == "Hname" {
				ref = ""
			}
			fmt.Fprintf(file, "    encode.%s(%so.%s);\n", lower(fldType), ref, snake(fldName))
		}
		fmt.Fprintf(file, "    return encode.data();\n}\n")

		fmt.Fprintf(file, "\npub fn decode_%s(bytes: &[u8]) -> %s {\n", funcName, typeName)
		fmt.Fprintf(file, "    let mut decode = BytesDecoder::new(bytes);\n    %s {\n", typeName)
		for _, fldName := range sortedKeys(fldDef) {
			fldType, _ := splitComment(fldDef[fldName])
			fmt.Fprintf(file, "        %s: decode.%s(),\n", snake(fldName), lower(fldType))
		}
		fmt.Fprintf(file, "    }\n}\n")
	}

	return nil
}

func GenerateRustCoreContractsSchema() error {
	coreSchemas, err := LoadCoreSchemas()
	if err != nil {
		return err
	}
	if coreSchemas == nil {
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

	for _, schema := range coreSchemas {
		scName := upper(snake(schema.Name))
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\npub const CORE_%s: ScHname = ScHname(0x%s);\n", scName, scHname.String())
		for _, funcName := range sortedMaps(schema.Funcs) {
			funcHname := coretypes.Hn(funcName)
			funcName = upper(snake(funcName))
			fmt.Fprintf(file, "pub const CORE_%s_FUNC_%s: ScHname = ScHname(0x%s);\n", scName, funcName, funcHname.String())
		}
		for _, funcName := range sortedMaps(schema.Views) {
			funcHname := coretypes.Hn(funcName)
			funcName = upper(snake(funcName))
			fmt.Fprintf(file, "pub const CORE_%s_VIEW_%s: ScHname = ScHname(0x%s);\n", scName, funcName, funcHname.String())
		}
	}
	return nil
}
