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

func GenerateRustTypes(path string) error {
	jsonTypes, err := LoadTypes(path)
	if err != nil {
		return err
	}

	keys := make([]string, 0)
	for key := range jsonTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

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
				if !ok {
					return errors.New("Invalid type: " + typeName)
				}
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

func GenerateGoTypes(path string) error {
	jsonTypes, err := LoadTypes(path)
	if err != nil {
		return err
	}

	keys := make([]string, 0)
	for key := range jsonTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

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
				if !ok {
					return errors.New("Invalid type: " + typeName)
				}
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

var javaTypes = map[string]string{
	"int":     "long",
	"address": "ScAddress",
	"agent":   "ScAgent",
	"color":   "ScColor",
	"string":  "String",
}

func GenerateJavaTypes(path string) error {
	jsonTypes, err := LoadTypes(path)
	if err != nil {
		return err
	}

	keys := make([]string, 0)
	for key := range jsonTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wtypes.json")
	contract := matchContract.ReplaceAllString(path, "$1")

	// write classes
	for _, structName := range keys {
		file, err := os.Create("../java/src/org/iota/wasplib/contracts/" + contract + "/" + structName + ".java")
		if err != nil {
			return err
		}
		defer file.Close()

		// write file header
		fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
		fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n\n")
		fmt.Fprintf(file, "package org.iota.wasplib.contracts.%s;\n\n", contract)
		fmt.Fprintf(file, "import org.iota.wasplib.client.bytes.BytesDecoder;\n")
		fmt.Fprintf(file, "import org.iota.wasplib.client.bytes.BytesEncoder;\n")
		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScAddress;\n")
		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScAgent;\n")
		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScColor;\n")

		fmt.Fprintf(file, "\npublic class %s{\n", structName)
		fields := jsonTypes[structName]
		for _, fld := range fields {
			for name, typeName := range fld {
				comment := ""
				index := strings.Index(typeName, "//")
				if index > 0 {
					comment = " // " + strings.TrimSpace(typeName[index+2:])
					typeName = strings.TrimSpace(typeName[:index])
				}
				javaType, ok := javaTypes[typeName]
				if !ok {
					return errors.New("Invalid type: " + typeName)
				}
				fmt.Fprintf(file, "    public %s %s;%s\n", javaType, camelcase(name), comment)
			}
		}
		//  write encoder and decoder for structs

		fmt.Fprintf(file, "\n    public static byte[] encode(%s o){\n", structName)
		fmt.Fprintf(file, "        return new BytesEncoder().\n")
		for _, fld := range fields {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "            %s(o.%s).\n", typeName, camelcase(name))
			}
		}
		fmt.Fprintf(file, "            Data();\n    }\n")

		fmt.Fprintf(file, "\n    public static %s decode(byte[] bytes) {\n", structName)
		fmt.Fprintf(file, "        BytesDecoder d = new BytesDecoder(bytes);\n        %s data = new %s();\n", structName, structName)
		for _, fld := range fields {
			for name, typeName := range fld {
				index := strings.Index(typeName, "//")
				if index > 0 {
					typeName = strings.TrimSpace(typeName[:index])
				}
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
				fmt.Fprintf(file, "        data.%s = d.%s();\n", camelcase(name), typeName)
			}
		}
		fmt.Fprintf(file, "        return data;\n    }\n")
		fmt.Fprintf(file, "}\n")
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
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateRustTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestGoTypes(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateGoTypes(path)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestJavaTypes(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\types.json") {
				return GenerateJavaTypes(path)
			}
			return nil
		})
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
				var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wsrc\\W.+")
				contract := matchContract.ReplaceAllString(path, "$1")
				return RustToGo(path, contract)
			}
			return nil
		})
	require.NoError(t, err)
}

func TestRustToJava(t *testing.T) {
	t.SkipNow()
	err := filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "\\lib.rs") {
				var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wsrc\\W.+")
				contract := matchContract.ReplaceAllString(path, "$1")
				return RustToJava(path, contract)
			}
			return nil
		})
	require.NoError(t, err)
}

func RustToGo(path string, contract string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	out, err := os.Create("../contracts/" + contract + "/lib.go")
	if err != nil {
		return err
	}
	defer out.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := RustToGoLine(text)
		if strings.HasPrefix(line, "use wasplib::client::*") {
			line = fmt.Sprintf("package %s\n\nimport \"github.com/iotaledger/wasplib/client\"", contract)
		}
		if line == "" && text != "" {
			continue
		}
		fmt.Fprintln(out, line)
	}
	return scanner.Err()
}

func RustToJava(path string, contract string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	out, err := os.Create("../java/src/org/iota/wasplib/contracts/" + contract + "/lib.java")
	if err != nil {
		return err
	}
	defer out.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := RustToJavaLine(text)
		if strings.HasPrefix(line, "use wasplib::client::*") {
			line = fmt.Sprintf("package org.iota.wasplib.contracts.%s;", contract)
		}
		if line == "" && text != "" {
			continue
		}
		fmt.Fprintln(out, "    ", line)
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
	if m[:1] == "\"" {
		return m
	}
	// replace Rust lower snake case to Go camel case
	index := strings.Index(m, "_")
	for index > 0 {
		m = m[:index] + strings.ToUpper(m[index+1:index+2]) + m[index+2:]
		index = strings.Index(m, "_")
	}
	return m
}

var goReplacements = []string{
	"pub fn ", "func ",
	"fn ", "func ",
	"ScExports::new", "client.NewScExports",
	"ScExports::nothing", "client.Nothing",
	"ScAgent::none", "&client.ScAgent{}",
	"ScColor::iota", "client.IOTA",
	"ScColor::mint", "client.MINT",
	"(&", "(",
	", &", ", ",
	": &Sc", " *client.Sc",
	": i64", " int64",
	": &str", " string",
	"0_i64", "int64(0)",
	"+ &\"", "+ \"",
	" unsafe ", " ",
	"\".ToString()", "\"",
	".Value().String()", ".String()",
	".ToString()", ".String()",
	" onLoad()", " OnLoad()",
	"#[noMangle]", "",
	"mod types", "",
	"use types::*", "",
}

var matchComment = regexp.MustCompile("^\\s*//")
var matchConst = regexp.MustCompile("[A-Z][A-Z_]+")
var matchConstStr = regexp.MustCompile("const (\\w+): &str = (\"\\w+\")")
var matchConstInt = regexp.MustCompile("const (\\w+): \\w+ = ([0-9]+)")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &str)? =")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_]+\\(")
var matchVarName = regexp.MustCompile(".[a-z][a-z_]+")
var matchToString = regexp.MustCompile("\\+ &([^ ]+)\\.ToString\\(\\)")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")
var matchIf = regexp.MustCompile("if (.+) {")
var matchExtraBraces = regexp.MustCompile("\\((\\([^)]+\\))\\)")
var matchParam = regexp.MustCompile("(\\(|, ?)(\\w+): &?(\\w+)")
var matchInitializer = regexp.MustCompile("(\\w+): (.+),$")
var matchCodec = regexp.MustCompile("(encode|decode)(\\w+)")
var matchInitializerHeader = regexp.MustCompile("(\\w+) :?= &?(\\w+) {")

func RustToGoLine(line string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = strings.Replace(line, ";", "", -1)
	line = matchConstStr.ReplaceAllString(line, "const $1 = client.Key($2)")
	line = matchConstInt.ReplaceAllString(line, "const $1 = $2")
	line = matchLet.ReplaceAllString(line, "$2 :=")
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchToString.ReplaceAllString(line, "+ $1.String()")
	line = matchForLoop.ReplaceAllString(line, "for $1 := int32($2); $1 < $3; $1++")
	line = matchInitializerHeader.ReplaceAllString(line, "$1 := &$2 {")

	for i := 0; i < len(goReplacements); i += 2 {
		line = strings.Replace(line, goReplacements[i], goReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	return line
}

var javaReplacements = []string{
	"pub fn ", "public static void ",
	"fn ", "public static void ",
	"ScExports::new", "new ScExports",
	"ScAgent::none", "ScAgent.NONE",
	"ScColor::iota", "ScColor.IOTA",
	"ScColor::mint", "ScColor.MINT",
	"(&", "(",
	", &", ", ",
	"};", "}",
	"0_i64", "0",
	"+ &\"", "+ \"",
	"\".ToString()", "\"",
	".Value().String()", ".toString()",
	"#[noMangle]", "",
	"mod types;", "",
	"use types::*;", "",
}

func RustToJavaLine(line string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = matchConstStr.ReplaceAllString(line, "private static final Key $1 = new Key($2)")
	line = matchConstInt.ReplaceAllString(line, "private static final int $1 = $2")
	line = matchLet.ReplaceAllString(line, "$2 =")
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchToString.ReplaceAllString(line, "+ $1")
	line = matchForLoop.ReplaceAllString(line, "for (int $1 = $2; $1 < $3; $1++)")
	line = matchIf.ReplaceAllString(line, "if ($1) {")
	line = matchParam.ReplaceAllString(line, "$1$3 $2")
	line = matchInitializer.ReplaceAllString(line, "xxx.$1 = $2;")
	line = matchCodec.ReplaceAllString(line, "$2.$1")
	line = matchInitializerHeader.ReplaceAllString(line, "$2 $1 = new $2();\n         {")

	for i := 0; i < len(javaReplacements); i += 2 {
		line = strings.Replace(line, javaReplacements[i], javaReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	return line
}
