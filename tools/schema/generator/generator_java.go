// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bufio"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"strings"
)

const generateJavaThunk = true

var javaTypes = StringMap{
	"Address":    "ScAddress",
	"AgentId":    "ScAgentId",
	"ChainId":    "ScChainId",
	"Color":      "ScColor",
	"ContractId": "ScContractId",
	"Hash":       "ScHash",
	"Hname":      "Hname",
	"Int":        "long",
	"String":     "String",
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
	defer os.Chdir(currentPath)

	err = os.MkdirAll("test", 0755)
	if err != nil {
		return err
	}
	//err = os.MkdirAll("wasmmain", 0755)
	//if err != nil {
	//	return err
	//}

	//err = s.GenerateJavaWasmMain()
	//if err != nil {
	//	return err
	//}
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
	//err = s.GenerateJavaFuncs()
	//if err != nil {
	//	return err
	//}
	return s.GenerateJavaTests()
}

func (s *Schema) GenerateJavaFunc(file *os.File, funcDef *FuncDef) error {
	funcName := funcDef.FullName
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\nfunc %s(ctx wasmlib.Sc%sContext, params *%sParams) {\n", funcName, funcKind, capitalize(funcName))
	fmt.Fprintf(file, "\tctx.Log(\"calling %s\")\n", funcDef.Name)
	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaFuncs() error {
	scFileName := s.Name + ".go"
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
	for _, funcDef := range s.Funcs {
		if existing[funcDef.FullName] == "" {
			err = s.GenerateJavaFunc(file, funcDef)
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
		matches := goFuncRegexp.FindStringSubmatch(line)
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
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintln(file, importWasmLib)

	for _, funcDef := range s.Funcs {
		err = s.GenerateJavaFunc(file, funcDef)
		if err != nil {
			return err
		}
	}
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
	fmt.Fprintf(file, "import org.iota.wasp.contracts.%s.*;\n", s.Name)
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.context.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.exports.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.immutable.*;\n\n")

	thunk := ""
	if generateJavaThunk {
		thunk = "Thunk"
	}

	fmt.Fprintf(file, "public class %sThunk {\n", s.FullName)

	fmt.Fprintf(file, "\tpublic static void onLoad() {\n")
	fmt.Fprintf(file, "\t\tScExports exports = new ScExports();\n")
	for _, funcDef := range s.Funcs {
		name := funcDef.FullName
		kind := capitalize(funcDef.FullName[:4])
		fmt.Fprintf(file, "\t\texports.Add%s(\"%s\", %s%s::%s%s);\n", kind, funcDef.Name, s.FullName, thunk, name, thunk)
	}
	fmt.Fprintf(file, "\t}\n")

	if generateJavaThunk {
		// generate parameter structs and thunks to set up and check parameters
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			params, err := os.Create("lib/" + name + "Params.java")
			if err != nil {
				return err
			}
			defer params.Close()
			s.GenerateJavaThunk(file, params, funcDef)
		}
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

	fmt.Fprintf(file, "\tpublic static final String ScName = \"%s\";\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "\tpublic static final String ScDescription = \"%s\";\n", s.Description)
	}
	hName := coretypes.Hn(s.Name)
	fmt.Fprintf(file, "\tpublic static final ScHname HScName = new ScHname(0x%s);\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			param := s.Params[name]
			name = capitalize(param.Name)
			fmt.Fprintf(file, "\tpublic static final Key Param%s = new Key(\"%s\");\n", name, param.Alias)
		}
	}

	if len(s.Vars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.Vars {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "\tpublic static final Key Var%s = new Key(\"%s\");\n", name, field.Alias)
		}
	}

	if len(s.Funcs) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			fmt.Fprintf(file, "\tpublic static final String %s = \"%s\";\n", name, funcDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "\tpublic static final ScHname H%s = new ScHname(0x%s);\n", name, hName.String())
		}
	}

	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateJavaTests() error {
	//TODO
	return nil
}

func (s *Schema) GenerateJavaThunk(file *os.File, params *os.File, funcDef *FuncDef) {
	// calculate padding
	nameLen, typeLen := calculatePadding(funcDef.Params, javaTypes, false)

	funcName := capitalize(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])

	fmt.Fprintln(params, copyright(true))
	fmt.Fprintf(params, "package org.iota.wasp.contracts.%s.lib;\n", s.Name)
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(params, "\nimport org.iota.wasp.wasmlib.immutable.*;\n")
	}
	fmt.Fprintf(params, "\npublic class %sParams {\n", funcName)
	for _, param := range funcDef.Params {
		fldName := capitalize(param.Name) + ";"
		if param.Comment != "" {
			fldName = pad(fldName, nameLen+1)
		}
		fldType := pad(param.Type, typeLen)
		fmt.Fprintf(params, "\tScImmutable%s %s%s\n", fldType, fldName, param.Comment)
	}
	fmt.Fprintf(params, "}\n")

	fmt.Fprintf(file, "\n\tprivate static void %sThunk(Sc%sContext ctx) {\n", funcDef.FullName, funcKind)
	grant := funcDef.Annotations["#grant"]
	if grant != "" {
		index := strings.Index(grant, "//")
		if index >= 0 {
			fmt.Fprintf(file, "\t\t%s\n", grant[index:])
			grant = strings.TrimSpace(grant[:index])
		}
		switch grant {
		case "self":
			grant = "ctx.ContractId().AsAgentId()"
		case "owner":
			grant = "ctx.ChainOwnerId()"
		case "creator":
			grant = "ctx.ContractCreator()"
		default:
			fmt.Fprintf(file, "\t\tScAgentId grantee := ctx.State().GetAgentId(new Key(\"%s\"));\n", grant)
			fmt.Fprintf(file, "\t\tctx.Require(grantee.Exists(), \"grantee not set: %s\");\n", grant)
			grant = fmt.Sprintf("grantee.Value()")
		}
		fmt.Fprintf(file, "\t\tctx.Require(ctx.Caller().equals(%s), \"no permission\");\n\n", grant)
	}
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "\t\tScImmutableMap p = ctx.Params();\n")
	}
	fmt.Fprintf(file, "\t\t%sParams params = new %sParams();\n", funcName, funcName)
	for _, param := range funcDef.Params {
		name := capitalize(param.Name)
		fmt.Fprintf(file, "\t\tparams.%s = p.Get%s(Consts.Param%s);\n", name, param.Type, name)
	}
	for _, param := range funcDef.Params {
		if !param.Optional {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "\t\tctx.Require(params.%s.Exists(), \"missing mandatory %s\");\n", name, param.Name)
		}
	}
	fmt.Fprintf(file, "\t\t%s.%s(ctx, params);\n", s.FullName, funcName)
	fmt.Fprintf(file, "\t}\n")
}

func (s *Schema) GenerateJavaTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	err := os.MkdirAll("types", 0755)
	if err != nil {
		return err
	}

	// write structs
	for _, typeDef := range s.Types {
		typeDef.GenerateJavaType(s.Name)
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
	fmt.Fprintf(file, "// +build wasm\n\n")
	fmt.Fprintf(file, "package main\n\n")
	fmt.Fprintf(file, importWasmClient)
	fmt.Fprintf(file, "import \"%s\"\n\n", importname)

	fmt.Fprintf(file, "func main() {\n")
	fmt.Fprintf(file, "}\n\n")

	fmt.Fprintf(file, "//export on_load\n")
	fmt.Fprintf(file, "func OnLoad() {\n")
	fmt.Fprintf(file, "\twasmclient.ConnectWasmHost()\n")
	fmt.Fprintf(file, "\t%s.OnLoad()\n", s.Name)
	fmt.Fprintf(file, "}\n")

	return nil
}
