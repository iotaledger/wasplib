package types

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var rustTypes = map[string]string{
	"int":     "i64",
	"address": "ScAddress",
	"agent":   "ScAgent",
	"color":   "ScColor",
	"string":  "String",
}

func GenerateRustTypes(path string) error {
	gen := &Generator{}
	err := gen.LoadTypes(path)
	if err != nil {
		return err
	}

	file, err := os.Create(path[:len(path)-len(".json")] + ".rs")
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
		for _, fld := range gen.jsonTypes[structName] {
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
		for _, fld := range gen.jsonTypes[structName] {
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
		for _, fld := range gen.jsonTypes[structName] {
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
