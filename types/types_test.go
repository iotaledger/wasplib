package types

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
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
				comment := ""
				index := strings.Index(typeName, "//")
				if index > 0 {
					comment = " // " + strings.TrimSpace(typeName[index+2:])
					typeName = strings.TrimSpace(typeName[:index])
				}
				rustType, ok := rustTypes[typeName]
				require.True(t, ok)
				fmt.Fprintf(file, "    pub %s: %s,%s\n", name, rustType, comment)
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
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
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
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
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
				comment := ""
				index := strings.Index(typeName, "//")
				if index > 0 {
					comment = " // " + strings.TrimSpace(typeName[index+2:])
					typeName = strings.TrimSpace(typeName[:index])
				}
				goType, ok := goTypes[typeName]
				require.True(t, ok)
				fmt.Fprintf(file, "    %s %s%s\n", camelcase(name), goType, comment)
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
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "        %s(o.%s).\n", typeName, camelcase(name))
			}
		}
		fmt.Fprintf(file, "        Data()\n}\n")

		fmt.Fprintf(file, "\nfunc de%s(bytes []byte) *%s {\n", funcName, structName)
		fmt.Fprintf(file, "    d := client.NewBytesDecoder(bytes)\n    data := &%s{}\n", structName)
		for _, fld := range fields {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
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
	t.SkipNow()
	contract := "dividend"
	path := "../rust/contracts/" + contract + "/src"
	jsonTypes, err := LoadTypes(path + "/types.json")
	require.NoError(t, err)
	require.NotNil(t, jsonTypes)
	err = GenerateRustTypes(t, jsonTypes, path+"/types.rs")
	require.NoError(t, err)
}

func TestGoTypes(t *testing.T) {
	t.SkipNow()
	contract := "dividend"
	path := "../contracts/" + contract
	jsonTypes, err := LoadTypes(path + "/types.json")
	require.NoError(t, err)
	require.NotNil(t, jsonTypes)
	err = GenerateGoTypes(t, jsonTypes, path, contract)
	require.NoError(t, err)
}

func TestRustToGo(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\lib.rs") {
				RustToGo(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func RustToGo(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	out, err := os.Create(path[:len(path) - len(".rs")] + ".go")
	if err != nil {
		return err
	}
	defer out.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := RustToGoLine(text)
		if line == "" && text != "" { continue }
		fmt.Fprintln(out, line)
	}
	return scanner.Err()
}

func replaceConst(m string) string {
	// replace Rust upper snake case to Go camel case
	return replaceVarName(strings.ToLower(m))
}

func replaceFuncCall(m string) string {
	// replace Rust . lower snake case ( to Go capitalized camel case
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceVarName(m string) string {
	if m[:1] == "\"" { return m }
	// replace Rust lower snake case to Go camel case
	index := strings.Index(m, "_")
	for index > 0 {
		m = m[:index] + strings.ToUpper(m[index+1:index+2]) + m[index+2:]
		index = strings.Index(m, "_")
	}
	return m
}

var replacements = []string{
	"pub fn ", "func ",
	"fn ", "func ",
	"ScExports::new", "client.NewScExports",
	"ScAgent::none", "&client.ScAgent{}",
	"ScColor::iota", "client.IOTA",
	"ScColor::mint", "client.MINT",
	"(&", "(",
	", &", ", ",
	": &Sc", " *client.Sc",
	": i64", " int64",
	"0_i64", "int64(0)",
	": &str", " string",
	"+ &\"", "+ \"",
	"\".ToString()", "\"",
	".Value().String()", ".String()",
	"#[noMangle]", "",
	"mod types", "",
	"use types::*", "",
	"func onLoad()", "func OnLoad()",
	"use wasplib::client::*", "import \"github.com/iotaledger/wasplib/client\"",
}

var matchComment = regexp.MustCompile("^\\s*//")
var matchConst = regexp.MustCompile("[A-Z][A-Z_]+")
var matchConstStr = regexp.MustCompile("(const \\w+): &str = (\"\\w+\")")
var matchConstInt = regexp.MustCompile("(const \\w+): \\w+ = ([0-9]+)")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &str)? =")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_]+\\(")
var matchVarName = regexp.MustCompile(".[a-z][a-z_]+")
var matchToString = regexp.MustCompile("\\+ &([^ ]+)\\.ToString\\(\\)")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")

func RustToGoLine(line string) string {
	if matchComment.MatchString(line) { return line }
	line = strings.Replace(line, ";", "", -1)
	line = matchConstStr.ReplaceAllString(line, "$1 = client.Key($2)")
	line = matchConstInt.ReplaceAllString(line, "$1 = $2")
	line = matchLet.ReplaceAllString(line, "$2 :=")
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchToString.ReplaceAllString(line, "+ $1.String()")
	line = matchForLoop.ReplaceAllString(line, "for $1 := int32($2); $1 < $3; $1++")

	for i := 0; i < len(replacements); i += 2 {
		line = strings.Replace(line, replacements[i], replacements[i + 1], -1)
	}

	return line
}
