// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var goReplacements = []string{
	"pub fn ", "func ",
	"fn ", "func ",
	"Hname::new", "wasmlib.NewHname",
	"None", "nil",
	"ScColor::Iota", "wasmlib.IOTA",
	"ScColor::Mint", "wasmlib.MINT",
	"ScExports::new", "wasmlib.NewScExports",
	"ScExports::nothing", "wasmlib.Nothing",
	"ScMutableMap::new", "wasmlib.NewScMutableMap",
	"ScTransfers::new", "wasmlib.NewScTransfer",
	"String::new()", "\"\"",
	"(&", "(",
	".Post(PostRequestParams", ".Post(&PostRequestParams",
	"PostRequestParams", "wasmlib.PostRequestParams",
	", &", ", ",
	": &Sc", " *wasmlib.Sc",
	": i64", " int64",
	": &str", " string",
	"0_i64", "int64(0)",
	"\".ToString()", "\"",
	" &\"", " \"",
	" + &", " + ",
	" unsafe ", " ",
	".Value().String()", ".String()",
	".ToString()", ".String()",
	" onLoad()", " OnLoad()",
	"decode", "Decode",
	"encode", "Encode",
	"#[noMangle]", "",
	"mod types", "",
	"use types::*", "",
}

var javaReplacements = []string{
	"pub fn ", "public static void ",
	"fn ", "public static void ",
	"ScExports::new", "new ScExports",
	"ScTransfers::new", "new ScTransfers",
	"::Iota", ".IOTA",
	"::Mint", ".MINT",
	"String::new()", "\"\"",
	"(&", "(",
	", &", ", ",
	"};", "}",
	"0_i64", "0",
	"+ &\"", "+ \"",
	"\".ToString()", "\"",
	".Value().String()", ".toString()",
	".ToString()", ".toString()",
	".Equals(", ".equals(",
	"#[noMangle]", "",
	"mod types;", "",
	"use types::*;", "",
}

var matchCodec = regexp.MustCompile("(encode|decode)(\\w+)")
var matchComment = regexp.MustCompile("^\\s*//")
var matchConst = regexp.MustCompile("[^a-zA-Z_][A-Z][A-Z_0-9]+")
var matchConstInt = regexp.MustCompile("const ([A-Z])([A-Z_0-9]+): \\w+ = ([0-9]+)")
var matchConstStr = regexp.MustCompile("const (PARAM|VAR|KEY)([A-Z_0-9]+): &str = (\"[^\"]+\")")
var matchExtraBraces = regexp.MustCompile("\\((\\([^)]+\\))\\)")
var matchFieldName = regexp.MustCompile("\\.[a-z][a-z_]+")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_]+\\(")
var matchIf = regexp.MustCompile("if (.+) {")
var matchInitializer = regexp.MustCompile("(\\w+): (.+),$")
var matchInitializerHeader = regexp.MustCompile("(\\w+) :?= &?(\\w+) {")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &str)? =")
var matchParam = regexp.MustCompile("(\\(|, ?)(\\w+): &?(\\w+)")
var matchSome = regexp.MustCompile("Some\\(([^)]+)\\)")
var matchToString = regexp.MustCompile("\\+ &([^ ]+)\\.ToString\\(\\)")
var matchVarName = regexp.MustCompile("[^a-zA-Z_][a-z][a-z_]+")

var lastInit string

func replaceConst(m string) string {
	// "[^a-zA-Z_][A-Z][A-Z_]+"
	// replace Rust upper snake case to Go public camel case
	m = strings.ToLower(m)
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceFieldName(m string) string {
	// "\\.[a-z][a-z_]+"
	// replace Rust lower snake case to Go public camel case
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceFuncCall(m string) string {
	// "\\.[a-z][a-z_]+\\("
	// replace Rust lower snake case to Go public camel case
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceInitializer(m string) string {
	// "(\\w+): (.+),$"
	// replace Rust lower case with Go upper case
	return strings.ToUpper(m[:1]) + m[1:]
}

func replaceVarName(m string) string {
	// "[^a-zA-Z_][a-z][a-z_]+"
	// replace Rust lower snake case to Go camel case
	index := strings.Index(m, "_")
	for index > 0 && index < len(m)-1 {
		m = m[:index] + strings.ToUpper(m[index+1:index+2]) + m[index+2:]
		index = strings.Index(m, "_")
	}
	return m
}

func RustConvertor(convertLine func(string, string) string, outPath string) error {
	return filepath.Walk("../rust/contracts",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !strings.HasSuffix(path, "\\lib.rs") {
				return nil
			}
			var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wsrc\\W.+")
			contract := matchContract.ReplaceAllString(path, "$1")
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			outFile := strings.Replace(outPath, "$1", contract, -1)
			os.MkdirAll(outFile[:strings.LastIndex(outFile, "/")], 0755)
			out, err := os.Create(outFile)
			if err != nil {
				return err
			}
			defer out.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text := scanner.Text()
				line := convertLine(text, contract)
				if line == "" && text != "" {
					continue
				}
				fmt.Fprintln(out, line)
			}
			return scanner.Err()
		})

}

func RustToGoLine(line string, contract string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = strings.Replace(line, ";", "", -1)
	line = matchConstInt.ReplaceAllString(line, "const $1$2 = $3")
	line = matchConstStr.ReplaceAllString(line, "const $1$2 = wasmlib.Key($3)")
	line = matchLet.ReplaceAllString(line, "$2 :=")
	line = matchForLoop.ReplaceAllString(line, "for $1 := int32($2); $1 < $3; $1++")
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchToString.ReplaceAllString(line, "+ $1.String()")
	line = matchInitializerHeader.ReplaceAllString(line, "$1 := &$2 {")
	line = matchSome.ReplaceAllString(line, "$1")

	lhs := strings.Index(line, "\"")
	if lhs < 0 {
		line = RustToGoVarNames(line)
	} else {
		rhs := strings.LastIndex(line, "\"")
		left := RustToGoVarNames(line[:lhs+1])
		mid := line[lhs+1 : rhs]
		right := RustToGoVarNames(line[rhs:])
		line = left + mid + right
	}

	line = matchInitializer.ReplaceAllStringFunc(line, replaceInitializer)

	for i := 0; i < len(goReplacements); i += 2 {
		line = strings.Replace(line, goReplacements[i], goReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	if strings.HasPrefix(line, "use wasmlib::*") {
		line = fmt.Sprintf("package %s\n\nimport \"github.com/iotaledger/wasplib/client\"", contract)
	}

	return line
}

func RustToGoVarNames(line string) string {
	line = matchFieldName.ReplaceAllStringFunc(line, replaceFieldName)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	return line
}

func RustToJavaLine(line string, contract string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = matchConstStr.ReplaceAllString(line, "private static final Key $1$2 = new Key($3)")
	line = matchConstInt.ReplaceAllString(line, "private static final int $1$2 = $3")
	line = matchLet.ReplaceAllString(line, "$2 =")
	line = matchForLoop.ReplaceAllString(line, "for (int $1 = $2; $1 < $3; $1++)")
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchInitializer.ReplaceAllString(line, lastInit+".$1 = $2;")
	line = matchToString.ReplaceAllString(line, "+ $1")
	line = matchIf.ReplaceAllString(line, "if ($1) {")
	line = matchParam.ReplaceAllString(line, "$1$3 $2")
	initParts := matchInitializerHeader.FindStringSubmatch(line)
	if initParts != nil {
		lastInit = initParts[1]
	}
	line = matchInitializerHeader.ReplaceAllString(line, "$2 $1 = new $2();\n         {")

	lhs := strings.Index(line, "\"")
	if lhs < 0 {
		line = RustToJavaVarNames(line)
	} else {
		rhs := strings.LastIndex(line, "\"")
		left := RustToJavaVarNames(line[:lhs+1])
		mid := line[lhs+1 : rhs]
		right := RustToJavaVarNames(line[rhs:])
		line = left + mid + right
	}

	line = matchCodec.ReplaceAllString(line, "$2.$1")

	for i := 0; i < len(javaReplacements); i += 2 {
		line = strings.Replace(line, javaReplacements[i], javaReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	if strings.HasPrefix(line, "use wasmlib::*") {
		line = fmt.Sprintf("package org.iota.wasplib.contracts.%s;", contract)
	}

	return line
}

func RustToJavaVarNames(line string) string {
	line = matchFieldName.ReplaceAllStringFunc(line, replaceFieldName)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	return line
}
