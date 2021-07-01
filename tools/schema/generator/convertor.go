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

var matchCodec = regexp.MustCompile("(encode|decode)(\\w+)")
var matchComment = regexp.MustCompile("^\\s*// ")
var matchConst = regexp.MustCompile("[^a-zA-Z_]H?[A-Z][A-Z_0-9]+")
var matchConstInt = regexp.MustCompile("const ([A-Z])([A-Z_0-9]+): \\w+ = ([0-9]+)")
var matchConstStr = regexp.MustCompile("const (PARAM|VAR|KEY)([A-Z_0-9]+): &str = (\"[^\"]+\")")
var matchConstStr2 = regexp.MustCompile("const ([A-Z_0-9]+): &str = (\"[^\"]+\")")
var matchCore = regexp.MustCompile("([^a-zA-Z_])Core([A-Z])")
var matchExtraBraces = regexp.MustCompile("\\((\\([^)]+\\))\\)")
var matchFieldName = regexp.MustCompile("\\.[a-z][a-z_0-9]+")
var matchForLoop = regexp.MustCompile("for (\\w+) in ([0-9+])\\.\\.(\\w+)")
var matchFromBytes = regexp.MustCompile("(\\w+)::from_bytes")
var matchFuncCall = regexp.MustCompile("\\.[a-z][a-z_0-9]+\\(")
var matchIf = regexp.MustCompile("if (.+) {")
var matchInitializer = regexp.MustCompile("(\\w+): (.+),$")
var matchInitializerHeader = regexp.MustCompile("(\\w+) :?= &?(\\w+) {")
var matchLet = regexp.MustCompile("let (mut )?(\\w+)(: &?\\w+)? =")
var matchParam = regexp.MustCompile("(\\(|, ?)(\\w+): &?(\\w+)")
var matchSome = regexp.MustCompile("Some\\(([^)]+)\\)")
var matchToString = regexp.MustCompile("\\+ &([^ ]+)\\.ToString\\(\\)")
var matchVarName = regexp.MustCompile("[^a-zA-Z_][a-z][a-z_0-9]+")

var lastInit string

func replaceConst(m string) string {
	// "[^a-zA-Z_][A-Z][A-Z_]+"
	// replace Rust upper snake case to Go public camel case
	m = strings.ToLower(m)
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceFieldName(m string) string {
	// "\\.[a-z][a-z_0-9]+"
	// replace Rust lower snake case to Go public camel case
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceFuncCall(m string) string {
	// "\\.[a-z][a-z_0-9]+\\("
	// replace Rust lower snake case to Go public camel case
	return replaceVarName(strings.ToUpper(m[:2]) + m[2:])
}

func replaceInitializer(m string) string {
	// "(\\w+): (.+),$"
	// replace Rust lower case with Go upper case
	return strings.ToUpper(m[:1]) + m[1:]
}

func replaceVarName(m string) string {
	// "[^a-zA-Z_][a-z][a-z_0-9]+"
	// replace Rust lower snake case to Go camel case
	index := strings.Index(m, "_")
	for index > 0 && index < len(m)-1 {
		m = m[:index] + strings.ToUpper(m[index+1:index+2]) + m[index+2:]
		index = strings.Index(m, "_")
	}
	return m
}

var matchContract = regexp.MustCompile("^rust.(\\w+).src.(\\w+).rs")

func RustConvertor(convertLine func(string, string) string, outPath string) error {
	return filepath.Walk("rust", func(path string, info os.FileInfo, err error) error {
		return walker(convertLine, outPath, path, info, err)
	})
}

func walker(convertLine func(string, string) string, outPath string, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !matchContract.MatchString(path) {
		return nil
	}
	matches := matchContract.FindStringSubmatch(path)
	if len(matches) != 3 || matches[1] != matches[2] {
		return nil
	}
	contract := matches[1]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	outFile := strings.Replace(outPath, "$c", contract, -1)
	outFile = strings.Replace(outFile, "$C", capitalize(contract), -1)
	os.MkdirAll(outFile[:strings.LastIndex(outFile, "/")], 0755)
	out, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer out.Close()
	scanner := bufio.NewScanner(file)
	emptyLines := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := convertLine(text, contract)
		if line == "" {
			if emptyLines != 0 || text != "" {
				// remove empty line
				continue
			}
			emptyLines++
		} else {
			emptyLines = 0
		}
		fmt.Fprintln(out, line)
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	line := convertLine("\u001A", contract)
	if line != "" {
		fmt.Fprintln(out, line)
	}
	return nil
}
