package types

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

var matchCodec = regexp.MustCompile("(encode|decode)(\\w+)")
var matchComment = regexp.MustCompile("^\\s*//")
var matchConst = regexp.MustCompile("[A-Z][A-Z_]+")
var matchConstInt = regexp.MustCompile("const (\\w+): \\w+ = ([0-9]+)")
var matchConstStr = regexp.MustCompile("const (\\w+): &str = (\"\\w+\")")
var matchExtraBraces = regexp.MustCompile("\\((\\([^)]+\\))\\)")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_]+\\(")
var matchIf = regexp.MustCompile("if (.+) {")
var matchInitializer = regexp.MustCompile("(\\w+): (.+),$")
var matchInitializerHeader = regexp.MustCompile("(\\w+) :?= &?(\\w+) {")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &str)? =")
var matchParam = regexp.MustCompile("(\\(|, ?)(\\w+): &?(\\w+)")
var matchToString = regexp.MustCompile("\\+ &([^ ]+)\\.ToString\\(\\)")
var matchVarName = regexp.MustCompile(".[a-z][a-z_]+")

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
		line := rustToGoLine(text)
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
		line := rustToJavaLine(text)
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

func rustToGoLine(line string) string {
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

func rustToJavaLine(line string) string {
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

