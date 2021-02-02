package types

import (
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"regexp"
	"strings"
)

var rustTypes = map[string]string{
	"address":     "ScAddress",
	"agent":       "ScAgent",
	"chain_id":    "ScChainId",
	"color":       "ScColor",
	"contract_id": "ScContractId",
	"hash":        "ScHash",
	"hname":       "Hname",
	"int":         "i64",
	"string":      "String",
}

func GenerateRustSchema(path string, contract string, gen *Generator) error {
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

	fmt.Fprintf(file, "pub const SC_NAME: &str = \"%s\";\n", gen.schema.Name)
	hName := coretypes.Hn(gen.schema.Name)
	fmt.Fprintf(file, "pub const SC_HNAME: Hname = Hname(0x%s);\n", hName.String())

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Params) {
		value := gen.schema.Params[name]
		fmt.Fprintf(file, "pub const PARAM_%s: &str = \"%s\";\n", snakecase(name), value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Vars) {
		value := gen.schema.Vars[name]
		fmt.Fprintf(file, "pub const VAR_%s: &str = \"%s\";\n", snakecase(name), value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Funcs) {
		value := gen.schema.Funcs[name]
		fmt.Fprintf(file, "pub const FUNC_%s: &str = \"%s\";\n", snakecase(name), value)
	}
	for _, name := range sorted(gen.schema.Views) {
		value := gen.schema.Views[name]
		fmt.Fprintf(file, "pub const VIEW_%s: &str = \"%s\";\n", snakecase(name), value)
	}

	fmt.Fprintln(file)
	for _, name := range sorted(gen.schema.Funcs) {
		value := gen.schema.Funcs[name]
		hName = coretypes.Hn(value)
		fmt.Fprintf(file, "pub const HFUNC_%s: Hname = Hname(0x%s);\n", snakecase(name), hName.String())
	}
	for _, name := range sorted(gen.schema.Views) {
		value := gen.schema.Views[name]
		hName = coretypes.Hn(value)
		fmt.Fprintf(file, "pub const HVIEW_%s: Hname = Hname(0x%s);\n", snakecase(name), hName.String())
	}

	fmt.Fprintf(file, "\n#[no_mangle]\n")
	fmt.Fprintf(file, "fn on_load() {\n")
	fmt.Fprintf(file, "    let exports = ScExports::new();\n")
	for _, name := range sorted(gen.schema.Funcs) {
		name = snakecase(name)
		funcName := strings.ToLower(name)
		fmt.Fprintf(file, "    exports.add_call(FUNC_%s, func_%s);\n", name, funcName)
	}
	for _, name := range sorted(gen.schema.Views) {
		name = snakecase(name)
		funcName := strings.ToLower(name)
		fmt.Fprintf(file, "    exports.add_view(VIEW_%s, view_%s);\n", name, funcName)
	}
	fmt.Fprintf(file, "}\n")

	return nil
}

func GenerateRustTypes(path string, contract string, gen *Generator) error {
	types := gen.schema.Types
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
	for _, structName := range gen.keys {
		gen.SplitComments(structName, rustTypes)
		spaces := strings.Repeat(" ", gen.maxName+gen.maxType)
		fmt.Fprintf(file, "\npub struct %s {\n", structName)
		fmt.Fprintf(file, "    //@formatter:off\n")
		for _, fld := range types[structName] {
			for name, _ := range fld {
				rustType := gen.types[name]
				comment := gen.comments[name]
				if comment != "" {
					comment = spaces[:gen.maxType-len(rustType)] + comment
				}
				rustType = spaces[:gen.maxName-len(name)] + rustType
				fmt.Fprintf(file, "    pub %s: %s,%s\n", name, rustType, comment)
			}
		}
		fmt.Fprintf(file, "    //@formatter:on\n")
		fmt.Fprintf(file, "}\n")
	}

	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	//  write encoder and decoder for structs
	for _, structName := range gen.keys {
		funcName := "code" + structName
		funcName = matchAllCap.ReplaceAllString(funcName, "${1}_${2}")
		funcName = strings.ToLower(funcName)
		fmt.Fprintf(file, "\npub fn en%s(o: &%s) -> Vec<u8> {\n", funcName, structName)
		fmt.Fprintf(file, "    let mut encode = BytesEncoder::new();\n")
		for _, fld := range types[structName] {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				ref := "&"
				if typeName == "int" {
					ref = ""
				}
				fmt.Fprintf(file, "    encode.%s(%so.%s);\n", typeName, ref, name)
			}
		}
		fmt.Fprintf(file, "    return encode.data();\n}\n")

		fmt.Fprintf(file, "\npub fn de%s(bytes: &[u8]) -> %s {\n", funcName, structName)
		fmt.Fprintf(file, "    let mut decode = BytesDecoder::new(bytes);\n    %s {\n", structName)
		for _, fld := range types[structName] {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				fmt.Fprintf(file, "        %s: decode.%s(),\n", name, typeName)
			}
		}
		fmt.Fprintf(file, "    }\n}\n")
	}

	//TODO write on_types function

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
