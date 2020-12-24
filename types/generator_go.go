package types

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var goTypes = map[string]string{
	"int":     "int64",
	"address": "*client.ScAddress",
	"agent":   "*client.ScAgent",
	"color":   "*client.ScColor",
	"string":  "string",
}

func GenerateGoTypes(path string) error {
	gen := &Generator{}
	err := gen.LoadTypes(path)
	if err != nil {
		return err
	}

	var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wtypes.json")
	contract := matchContract.ReplaceAllString(path, "$1")

	file, err := os.Create(path[:len(path)-len(".json")] + ".go")
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
		for _, fld := range gen.jsonTypes[structName] {
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
		fmt.Fprintf(file, "\nfunc en%s(o *%s) []byte {\n", funcName, structName)
		fmt.Fprintf(file, "\treturn client.NewBytesEncoder().\n")
		for _, fld := range gen.jsonTypes[structName] {
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

		fmt.Fprintf(file, "\nfunc de%s(bytes []byte) *%s {\n", funcName, structName)
		fmt.Fprintf(file, "\tdecode := client.NewBytesDecoder(bytes)\n\tdata := &%s{}\n", structName)
		for _, fld := range gen.jsonTypes[structName] {
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
