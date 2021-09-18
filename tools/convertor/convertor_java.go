// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package convertor

import (
	"fmt"
	"regexp"
	"strings"
)

var javaReplacements = []string{
	"(&", "(",
	", &", ", ",
	"pub fn ", "public static void ",
	"fn ", "public static void ",
	"None", "null",
	"ScColor::Iota", "ScColor.IOTA",
	"ScColor::Mint", "ScColor.MINT",
	"ScAgentID::new", "new ScAgentID",
	"ScHname::new", "new ScHname",
	"ScMutableMap::new", "new ScMutableMap",
	"ScTransfers::new", "new ScTransfers",
	"ScTransfers::iotas", "ScTransfers.iotas",
	" str = \"", " = \"",
	"String::new()", "\"\"",

	"};", "}",
	"0_i64", "0",
	" as i64", "",
	"i64", "long",

	"_ctx", "ctx",
	"_params", "params",
	"Hview", "HView",
	"Hfunc", "HFunc",

	"+ &\"", "+ \"",
	" unsafe ", " ",
	"\".ToString()", "\"",
	".Value().String()", ".toString()",
	".ToString()", ".toString()",
	".ToBytes()", ".toBytes()",

	"#[noMangle]", "",
	"mod types;", "",
	"use crate::*;", "",
	"use crate::types::*;", "",
	"\u001A", "}",
}

func RustToJavaLine(line, contract string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = matchConstStr.ReplaceAllString(line, "private static final Key $1$2 = new Key($3)")
	line = matchConstInt.ReplaceAllString(line, "private static final int $1$2 = $3")
	line = matchLet.ReplaceAllString(line, "var $2 =")
	line = matchForLoop.ReplaceAllString(line, "for (var $1 = $2; $1 < $3; $1++)")
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchInitializer.ReplaceAllString(line, lastInit+".$1 = $2;")
	line = matchToString.ReplaceAllString(line, "+ $1")
	line = matchIf.ReplaceAllString(line, "if ($1) {")
	line = matchParam.ReplaceAllString(line, "$1$3 $2")
	initParts := matchInitializerHeader.FindStringSubmatch(line)
	if initParts != nil {
		lastInit = initParts[1]
	}
	line = matchInitializerHeader.ReplaceAllString(line, "$1 = new $2();\n         {")
	line = matchFromBytes.ReplaceAllString(line, "new $1")

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
		line = fmt.Sprintf("package org.iota.wasp.contracts.%s;\n\npublic class %s {", contract, capitalize(contract))
	}

	return line
}

func RustToJavaVarNames(line string) string {
	line = matchFieldName.ReplaceAllStringFunc(line, replaceFieldName)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchConst.ReplaceAllStringFunc(line, replaceJavaConst)
	return line
}

var matchJavaConst = regexp.MustCompile(`^[^a-zA-Z_](Param|Var|HFunc|HView)[A-Z]`)

func replaceJavaConst(m string) string {
	// "[^a-zA-Z_][A-Z][A-Z_]+"
	// replace Rust upper snake case to Java public camel case
	// with Consts. prefix
	m = strings.ToLower(m)
	m = replaceVarName(strings.ToUpper(m[:2]) + m[2:])
	if matchJavaConst.MatchString(m) {
		return m[:1] + "Consts." + m[1:]
	}
	return m
}
