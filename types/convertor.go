package types

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
	"Hname::new", "client.NewHname",
	"Hname::Self", "client.Hname(0)",
	"&ScAddress::Null", "nil",
	"ScAddress::Null", "&client.ScAddress{}",
	"ScAgent::Null", "&client.ScAgent{}",
	"ScColor::Iota", "client.IOTA",
	"ScColor::Mint", "client.MINT",
	"ScExports::new", "client.NewScExports",
	"ScExports::nothing", "client.Nothing",
	"ScHash::Null", "&client.ScHash{}",
	"ScMutableMap::new", "client.NewScMutableMap",
	"ScMutableMap::None", "nil",
	"ScTransfers::new", "client.NewScTransfer",
	"&ScTransfers::None", "nil",
	"String::new()", "\"\"",
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
	"::Null", ".NULL",
	"::Iota", ".IOTA",
	"::Mint", ".MINT",
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

var matchCodec = regexp.MustCompile("(encode|decode)(\\w+)")
var matchComment = regexp.MustCompile("^\\s*//")
var matchConst = regexp.MustCompile("[^a-zA-Z_][A-Z][A-Z_]+")
var matchConstInt = regexp.MustCompile("const (\\w+): \\w+ = ([0-9]+)")
var matchConstStr = regexp.MustCompile("const (\\w+): &str = (\"\\w+\")")
var matchExtraBraces = regexp.MustCompile("\\((\\([^)]+\\))\\)")
var matchFieldName = regexp.MustCompile("\\.[a-z][a-z_]+")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_]+\\(")
var matchIf = regexp.MustCompile("if (.+) {")
var matchInitializer = regexp.MustCompile("(\\w+): (.+),$")
var matchInitializerHeader = regexp.MustCompile("(\\w+) :?= &?(\\w+) {")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &str)? =")
var matchParam = regexp.MustCompile("(\\(|, ?)(\\w+): &?(\\w+)")
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

func replaceVarName(m string) string {
	// "[^a-zA-Z_][a-z][a-z_]+"
	// replace Rust lower snake case to Go camel case
	index := strings.Index(m, "_")
	for index > 0 {
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
	line = matchConstInt.ReplaceAllString(line, "const $1 = $2")
	line = matchConstStr.ReplaceAllString(line, "const $1 = client.Key($2)")
	line = matchLet.ReplaceAllString(line, "$2 :=")
	line = matchForLoop.ReplaceAllString(line, "for $1 := int32($2); $1 < $3; $1++")
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchToString.ReplaceAllString(line, "+ $1.String()")
	line = matchInitializerHeader.ReplaceAllString(line, "$1 := &$2 {")

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

	for i := 0; i < len(goReplacements); i += 2 {
		line = strings.Replace(line, goReplacements[i], goReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	if strings.HasPrefix(line, "use wasplib::client::*") {
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
	line = matchConstStr.ReplaceAllString(line, "private static final Key $1 = new Key($2)")
	line = matchConstInt.ReplaceAllString(line, "private static final int $1 = $2")
	line = matchLet.ReplaceAllString(line, "$2 =")
	line = matchForLoop.ReplaceAllString(line, "for (int $1 = $2; $1 < $3; $1++)")
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchInitializer.ReplaceAllString(line, lastInit+".$1 = $2;")
	line = matchFieldName.ReplaceAllStringFunc(line, replaceFieldName)
	line = matchToString.ReplaceAllString(line, "+ $1")
	line = matchIf.ReplaceAllString(line, "if ($1) {")
	line = matchParam.ReplaceAllString(line, "$1$3 $2")
	line = matchCodec.ReplaceAllString(line, "$2.$1")
	initParts := matchInitializerHeader.FindStringSubmatch(line)
	if initParts != nil {
		lastInit = initParts[1]
	}
	line = matchInitializerHeader.ReplaceAllString(line, "$2 $1 = new $2();\n         {")

	for i := 0; i < len(javaReplacements); i += 2 {
		line = strings.Replace(line, javaReplacements[i], javaReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	if strings.HasPrefix(line, "use wasplib::client::*") {
		line = fmt.Sprintf("package org.iota.wasplib.contracts.%s;", contract)
	}

	return line
}
