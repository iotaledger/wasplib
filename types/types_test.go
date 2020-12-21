package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
)

type JsonType map[string]string

type JsonTypes map[string][]JsonType

func LoadTypes(path string) (JsonTypes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonTypes := make(JsonTypes)
	err = json.NewDecoder(file).Decode(&jsonTypes)
	if err != nil {
		return nil, errors.New("JSON error: " + err.Error())
	}
	return jsonTypes, nil
}

var rustTypes = map[string]string{
	"int":     "i64",
	"address": "ScAddress",
	"agent":   "ScAgent",
	"color":   "ScColor",
	"string":  "String",
}

func GenerateRustTypes(t *testing.T, jsonTypes JsonTypes, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	keys := make([]string, 0)
	for key := range jsonTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "use wasplib::client::*;\n")

	// write structs
	for _, structName := range keys {
		fmt.Fprintf(file, "\npub struct %s {\n", structName)
		fields := jsonTypes[structName]
		for _, fld := range fields {
			for name, typeName := range fld {
				rustType, ok := rustTypes[typeName]
				require.True(t, ok)
				fmt.Fprintf(file, "    pub %s: %s,\n", name, rustType)
			}
		}
		fmt.Fprintf(file, "}\n")
	}

	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	//  write encoder and decoder for structs
	for _, structName := range keys {
		funcName := "code" + structName
		funcName = matchAllCap.ReplaceAllString(funcName, "${1}_${2}")
		funcName = strings.ToLower(funcName)
		fields := jsonTypes[structName]
		fmt.Fprintf(file, "\npub fn en%s(o: &%s) -> Vec<u8> {\n", funcName, structName)
		fmt.Fprintf(file, "    let mut e = BytesEncoder::new();\n")
		for _, fld := range fields {
			for name, typeName := range fld {
				ref := "&"
				if typeName == "int" {
					ref = ""
				}
				fmt.Fprintf(file, "    e.%s(%so.%s);\n", typeName, ref, name)
			}
		}
		fmt.Fprintf(file, "    return e.data();\n}\n")

		fmt.Fprintf(file, "\npub fn de%s(bytes: &[u8]) -> %s {\n", funcName, structName)
		fmt.Fprintf(file, "    let mut d = BytesDecoder::new(bytes);\n    %s {\n", structName)
		for _, fld := range fields {
			for name, typeName := range fld {
				fmt.Fprintf(file, "        %s: d.%s(),\n", name, typeName)
			}
		}
		fmt.Fprintf(file, "    }\n}\n")
	}

	//TODO write on_types function

	return nil
}

var goTypes = map[string]string{
	"int":     "int64",
	"address": "*client.ScAddress",
	"agent":   "*client.ScAgent",
	"color":   "*client.ScColor",
	"string":  "string",
}

func GenerateGoTypes(t *testing.T, jsonTypes JsonTypes, path string, contract string) error {
	file, err := os.Create(path + "/types.go")
	if err != nil {
		return err
	}
	defer file.Close()
	keys := make([]string, 0)
	for key := range jsonTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// write file header
	fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
	fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n")
	fmt.Fprintf(file, "\npackage %s\n\n", contract)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n")

	// write structs
	for _, structName := range keys {
		fmt.Fprintf(file, "\ntype %s struct {\n", structName)
		fields := jsonTypes[structName]
		for _, fld := range fields {
			for name, typeName := range fld {
				goType, ok := goTypes[typeName]
				require.True(t, ok)
				fmt.Fprintf(file, "    %s %s\n", camelcase(name), goType)
			}
		}
		fmt.Fprintf(file, "}\n")
	}

	//  write encoder and decoder for structs
	for _, structName := range keys {
		funcName := "code" + structName
		fields := jsonTypes[structName]
		fmt.Fprintf(file, "\nfunc en%s(o *%s) []byte {\n", funcName, structName)
		fmt.Fprintf(file, "    return client.NewBytesEncoder().\n")
		for _, fld := range fields {
			for name, typeName := range fld {
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "        %s(o.%s).\n", typeName, camelcase(name))
			}
		}
		fmt.Fprintf(file, "        Data()\n}\n")

		fmt.Fprintf(file, "\nfunc de%s(bytes []byte) *%s {\n", funcName, structName)
		fmt.Fprintf(file, "    d := client.NewBytesDecoder(bytes)\n    data := &%s{}\n", structName)
		for _, fld := range fields {
			for name, typeName := range fld {
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "    data.%s = d.%s()\n", camelcase(name), typeName)
			}
		}
		fmt.Fprintf(file, "    return data\n}\n")
	}

	//TODO write on_types function

	return nil
}

func camelcase(name string) string {
	index := strings.Index(name, "_")
	for index > 0 {
		c := name[index+1 : index+2]
		name = name[:index] + strings.ToUpper(c) + name[index+2:]
		index = strings.Index(name, "_")
	}
	return name
}

func TestRustTypes(t *testing.T) {
	contract := "tokenregistry"
	path := "../rust/contracts/" + contract + "/src"
	jsonTypes, err := LoadTypes(path + "/types.json")
	require.NoError(t, err)
	require.NotNil(t, jsonTypes)
	err = GenerateRustTypes(t, jsonTypes, path+"/types.rs")
	require.NoError(t, err)
}

func TestGoTypes(t *testing.T) {
	contract := "donatewithfeedback"
	path := "../contracts/" + contract
	jsonTypes, err := LoadTypes(path + "/types.json")
	require.NoError(t, err)
	require.NotNil(t, jsonTypes)
	err = GenerateGoTypes(t, jsonTypes, path, contract)
	require.NoError(t, err)
}
