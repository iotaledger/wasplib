// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iotaledger/wasp/packages/iscp"
)

var javaFuncRegexp = regexp.MustCompile(`public static void (\w+).+$`)

var javaTypes = StringMap{
	"Address":   "ScAddress",
	"AgentID":   "ScAgentID",
	"ChainID":   "ScChainID",
	"Color":     "ScColor",
	"Hash":      "ScHash",
	"Hname":     "Hname",
	"Int16":     "short",
	"Int32":     "int",
	"Int64":     "long",
	"RequestID": "ScRequestID",
	"String":    "String",
}

func (s *Schema) GenerateJava() error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	javaPath := "../../java/src/org/iota/wasp/contracts/" + s.Name
	err = os.MkdirAll(javaPath, 0755)
	if err != nil {
		return err
	}
	err = os.Chdir(javaPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Chdir(currentPath)
	}()

	err = os.MkdirAll("test", 0755)
	if err != nil {
		return err
	}
	// err = os.MkdirAll("wasmmain", 0755)
	// if err != nil {
	// 	return err
	// }

	// err = s.GenerateJavaWasmMain()
	// if err != nil {
	// 	return err
	// }
	err = s.GenerateJavaLib()
	if err != nil {
		return err
	}
	err = s.GenerateJavaConsts()
	if err != nil {
		return err
	}
	err = s.GenerateJavaTypes()
	if err != nil {
		return err
	}
	// err = s.GenerateJavaFuncs()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *Schema) GenerateJavaFunc(file *os.File, f *FuncDef) error {
	funcName := f.FuncName
	funcKind := capitalize(f.FuncName[:4])
	fmt.Fprintf(file, "\npublic static void %s(Sc%sContext ctx, %sParams params) {\n", funcName, funcKind, capitalize(funcName))
	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaFuncs() error {
	scFileName := s.Name + ".java"
	file, err := os.Open(scFileName)
	if err != nil {
		return s.GenerateJavaFuncsNew(scFileName)
	}
	lines, existing, err := s.GenerateJavaFuncScanner(file)
	if err != nil {
		return err
	}

	// save old one from overwrite
	scOriginal := s.Name + ".bak"
	err = os.Rename(scFileName, scOriginal)
	if err != nil {
		return err
	}
	file, err = os.Create(scFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// make copy of file
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	// append any new funcs
	for _, f := range s.Funcs {
		if existing[f.FuncName] == "" {
			err = s.GenerateJavaFunc(file, f)
			if err != nil {
				return err
			}
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) GenerateJavaFuncScanner(file *os.File) ([]string, StringMap, error) {
	defer file.Close()
	existing := make(StringMap)
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := javaFuncRegexp.FindStringSubmatch(line)
		if matches != nil {
			existing[matches[1]] = line
		}
		lines = append(lines, line)
	}
	err := scanner.Err()
	if err != nil {
		return nil, nil, err
	}
	return lines, existing, nil
}

func (s *Schema) GenerateJavaFuncsNew(scFileName string) error {
	file, err := os.Create(scFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(false))
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package org.iota.wasp.contracts.%s;\n\n", s.Name)
	fmt.Fprintf(file, "import org.iota.wasp.contracts.%s.lib.*;\n", s.Name)
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.context.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.hashtypes.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.immutable.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.mutable.*;\n\n")

	fmt.Fprintf(file, "public class %s {\n", s.FullName)
	for _, f := range s.Funcs {
		err = s.GenerateJavaFunc(file, f)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaLib() error {
	err := os.MkdirAll("lib", 0755)
	if err != nil {
		return err
	}
	file, err := os.Create("lib/" + s.FullName + "Thunk.java")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package org.iota.wasp.contracts.%s.lib;\n\n", s.Name)
	fmt.Fprintf(file, "import de.mirkosertic.bytecoder.api.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.contracts.%s.*;\n", s.Name)
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.context.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.exports.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.immutable.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.keys.*;\n\n")

	fmt.Fprintf(file, "public class %sThunk {\n", s.FullName)
	fmt.Fprintf(file, "    public static void main(String[] args) {\n")
	fmt.Fprintf(file, "    }\n\n")

	fmt.Fprintf(file, "    @Export(\"on_load\")\n")
	fmt.Fprintf(file, "    public static void onLoad() {\n")
	fmt.Fprintf(file, "        ScExports exports = new ScExports();\n")
	for _, f := range s.Funcs {
		name := capitalize(f.FuncName)
		kind := capitalize(f.FuncName[:4])
		fmt.Fprintf(file, "        exports.Add%s(Consts.%s, %sThunk::%sThunk);\n", kind, name, s.FullName, f.FuncName)
	}
	fmt.Fprintf(file, "    }\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, f := range s.Funcs {
		name := capitalize(f.FuncName)
		params, err := os.Create("lib/" + name + "Params.java")
		if err != nil {
			return err
		}
		defer params.Close()
		s.GenerateJavaThunk(file, params, f)
	}

	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaConsts() error {
	file, err := os.Create("lib/Consts.java")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package org.iota.wasp.contracts.%s.lib;\n\n", s.Name)
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.hashtypes.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.keys.*;\n")
	fmt.Fprintf(file, "\npublic class Consts {\n")

	fmt.Fprintf(file, "    public static final String ScName = \"%s\";\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "    public static final String ScDescription = \"%s\";\n", s.Description)
	}
	hName := iscp.Hn(s.Name)
	fmt.Fprintf(file, "    public static final ScHname HScName = new ScHname(0x%s);\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, param := range s.Params {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "    public static final Key Param%s = new Key(\"%s\");\n", name, param.Alias)
		}
	}

	if len(s.StateVars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.StateVars {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "    public static final Key Var%s = new Key(\"%s\");\n", name, field.Alias)
		}
	}

	if len(s.Funcs) != 0 {
		fmt.Fprintln(file)
		for _, f := range s.Funcs {
			name := capitalize(f.FuncName)
			fmt.Fprintf(file, "    public static final String %s = \"%s\";\n", name, f.String)
		}

		fmt.Fprintln(file)
		for _, f := range s.Funcs {
			name := capitalize(f.FuncName)
			hName = iscp.Hn(f.String)
			fmt.Fprintf(file, "    public static final ScHname H%s = new ScHname(0x%s);\n", name, hName.String())
		}
	}

	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaThunk(file, params *os.File, f *FuncDef) {
	// calculate padding
	nameLen, typeLen := calculatePadding(f.Params, javaTypes, false)

	funcName := capitalize(f.FuncName)
	funcKind := capitalize(f.FuncName[:4])

	fmt.Fprintln(params, copyright(true))
	fmt.Fprintf(params, "package org.iota.wasp.contracts.%s.lib;\n", s.Name)
	if len(f.Params) != 0 {
		fmt.Fprintf(params, "\nimport org.iota.wasp.wasmlib.immutable.*;\n")
	}
	if len(f.Params) > 1 {
		fmt.Fprintf(params, "\n// @formatter:off")
	}
	fmt.Fprintf(params, "\npublic class %sParams {\n", funcName)
	for _, param := range f.Params {
		fldName := capitalize(param.Name) + ";"
		if param.Comment != "" {
			fldName = pad(fldName, nameLen+1)
		}
		fldType := pad(param.Type, typeLen)
		fmt.Fprintf(params, "    public ScImmutable%s %s%s\n", fldType, fldName, param.Comment)
	}
	fmt.Fprintf(params, "}\n")
	if len(f.Params) > 1 {
		fmt.Fprintf(params, "// @formatter:on\n")
	}

	fmt.Fprintf(file, "\n    private static void %sThunk(Sc%sContext ctx) {\n", f.FuncName, funcKind)
	fmt.Fprintf(file, "        ctx.Log(\"%s.%s\");\n", s.Name, f.FuncName)

	if f.Access != "" {
		s.generateJavaThunkAccessCheck(file, f)
	}

	if len(f.Params) != 0 {
		fmt.Fprintf(file, "        var p = ctx.Params();\n")
	}
	fmt.Fprintf(file, "        var params = new %sParams();\n", funcName)
	for _, param := range f.Params {
		name := capitalize(param.Name)
		fmt.Fprintf(file, "        params.%s = p.Get%s(Consts.Param%s);\n", name, param.Type, name)
	}
	for _, param := range f.Params {
		if !param.Optional {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "        ctx.Require(params.%s.Exists(), \"missing mandatory %s\");\n", name, param.Name)
		}
	}
	fmt.Fprintf(file, "        %s.%s(ctx, params);\n", s.FullName, f.FuncName)
	fmt.Fprintf(file, "        ctx.Log(\"%s.%s ok\");\n", s.Name, f.FuncName)
	fmt.Fprintf(file, "    }\n")
}

func (s *Schema) generateJavaThunkAccessCheck(file *os.File, f *FuncDef) {
	grant := f.Access
	index := strings.Index(grant, "//")
	if index >= 0 {
		fmt.Fprintf(file, "        %s\n", grant[index:])
		grant = strings.TrimSpace(grant[:index])
	}
	switch grant {
	case AccessSelf:
		grant = "ctx.AccountID()"
	case AccessChain:
		grant = "ctx.ChainOwnerID()"
	case AccessCreator:
		grant = "ctx.ContractCreator()"
	default:
		fmt.Fprintf(file, "        var access = ctx.State().GetAgentID(new Key(\"%s\"));\n", grant)
		fmt.Fprintf(file, "        ctx.Require(access.Exists(), \"access not set: %s\");\n", grant)
		grant = "access.Value()"
	}
	fmt.Fprintf(file, "        ctx.Require(ctx.Caller().equals(%s), \"no permission\");\n\n", grant)
}

func (s *Schema) GenerateJavaTypes() error {
	if len(s.Structs) == 0 {
		return nil
	}

	err := os.MkdirAll("types", 0755)
	if err != nil {
		return err
	}

	// write structs
	for _, typeDef := range s.Structs {
		err = typeDef.GenerateJavaType(s.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Schema) GenerateJavaWasmMain() error {
	file, err := os.Create("wasmmain/" + s.Name + ".go")
	if err != nil {
		return err
	}
	defer file.Close()

	importname := ModuleName + strings.Replace(ModuleCwd[len(ModulePath):], "\\", "/", -1)
	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprint(file, "// +build wasm\n\n")
	fmt.Fprint(file, "package main\n\n")
	fmt.Fprint(file, importWasmClient)
	fmt.Fprintf(file, "import \"%s\"\n\n", importname)

	fmt.Fprintf(file, "func main() {\n")
	fmt.Fprintf(file, "}\n\n")

	fmt.Fprintf(file, "//export on_load\n")
	fmt.Fprintf(file, "func OnLoad() {\n")
	fmt.Fprintf(file, "    wasmclient.ConnectWasmHost()\n")
	fmt.Fprintf(file, "    %s.OnLoad()\n", s.Name)
	fmt.Fprintf(file, "}\n")

	return nil
}
