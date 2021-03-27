package generator

import (
	"fmt"
	"strings"
)

var goReplacements = []string{
	"(&", "(",
	", &", ", ",
	"pub fn ", "func ",
	"fn ", "func ",
	"None", "nil",
	"ScColor::Iota", "wasmlib.IOTA",
	"ScColor::Mint", "wasmlib.MINT",
	"ScHname::new", "wasmlib.NewScHname",
	"ScMutableMap::new", "wasmlib.NewScMutableMap",
	"ScTransfers::new", "wasmlib.NewScTransfer",
	"ScTransfers::iotas", "wasmlib.NewScTransferIotas",
	" str = \"", " = \"",
	"String::new()", "\"\"",

	".Post(PostRequestParams", ".Post(&PostRequestParams",
	"PostRequestParams", "wasmlib.PostRequestParams",
	": &Sc", " wasmlib.Sc",
	": i64", " int64",
	": &str", " string",
	"0_i64", "int64(0)",

	"_ctx", "ctx",
	"_params", "params",
	"Hview", "HView",
	"Hfunc", "HFunc",

	"\".ToString()", "\"",
	" &\"", " \"",
	" + &", " + ",
	" unsafe ", " ",
	".ToString()", ".String()",
	".ToBytes()", ".Bytes()",
	".Value().String()", ".String()",
	" onLoad()", " OnLoad()",
	"params: &", "params *",

	"#[noMangle]", "",
	"mod types", "",
	"use crate::*", "",
	"use crate::types::*", "",
	"\u001A", "",
}

func RustToGoLine(line string, contract string) string {
	if matchComment.MatchString(line) {
		return line
	}
	line = strings.Replace(line, ";", "", -1)
	line = matchConstInt.ReplaceAllString(line, "const $1$2 = $3")
	line = matchConstStr.ReplaceAllString(line, "const $1$2 = wasmlib.Key($3)")
	line = matchConstStr2.ReplaceAllString(line, "const $1 = $2")
	line = matchLet.ReplaceAllString(line, "$2 :=")
	line = matchForLoop.ReplaceAllString(line, "for $1 := int32($2); $1 < $3; $1++")
	line = matchFuncCall.ReplaceAllStringFunc(line, replaceFuncCall)
	line = matchToString.ReplaceAllString(line, "+ $1.String()")
	line = matchInitializerHeader.ReplaceAllString(line, "$1 := &$2 {")
	line = matchSome.ReplaceAllString(line, "$1")
	line = matchFromBytes.ReplaceAllString(line, "New${1}FromBytes")

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

	line = matchCore.ReplaceAllString(line, "${1}wasmlib.Core$2")
	line = matchInitializer.ReplaceAllStringFunc(line, replaceInitializer)

	for i := 0; i < len(goReplacements); i += 2 {
		line = strings.Replace(line, goReplacements[i], goReplacements[i+1], -1)
	}

	line = matchExtraBraces.ReplaceAllString(line, "$1")

	if strings.HasPrefix(line, "use wasmlib::*") {
		line = fmt.Sprintf("package %s\n\nimport (\n\t\"github.com/iotaledger/wasplib/packages/vm/wasmlib\"\n)", contract)
	}

	return line
}

func RustToGoVarNames(line string) string {
	line = matchFieldName.ReplaceAllStringFunc(line, replaceFieldName)
	line = matchVarName.ReplaceAllStringFunc(line, replaceVarName)
	line = matchConst.ReplaceAllStringFunc(line, replaceConst)
	return line
}
