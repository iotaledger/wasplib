// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bufio"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"regexp"
	"strings"
)

const importWasmLib = "import \"github.com/iotaledger/wasp/packages/vm/wasmlib\"\n"
const importWasmClient = "import \"github.com/iotaledger/wasp/packages/vm/wasmclient\"\n"

var goFuncRegexp = regexp.MustCompile("^func (\\w+).+$")

var goTypes = StringMap{
	"Address":    "*wasmlib.ScAddress",
	"AgentId":    "*wasmlib.ScAgentId",
	"ChainId":    "*wasmlib.ScChainId",
	"Color":      "*wasmlib.ScColor",
	"ContractId": "*wasmlib.ScContractId",
	"Hash":       "*wasmlib.ScHash",
	"Hname":      "wasmlib.ScHname",
	"Int":        "int64",
	"String":     "string",
}

func (s *Schema) GenerateGo() error {
	err := os.MkdirAll("test", 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll("wasmmain", 0755)
	if err != nil {
		return err
	}

	err = s.GenerateGoWasmMain()
	if err != nil {
		return err
	}
	err = s.GenerateGoOnLoad()
	if err != nil {
		return err
	}
	err = s.GenerateGoSchema()
	if err != nil {
		return err
	}
	err = s.GenerateGoTypes()
	if err != nil {
		return err
	}
	err = s.GenerateGoFuncs()
	if err != nil {
		return err
	}
	return s.GenerateGoTests()
}

func (s *Schema) GenerateGoFunc(file *os.File, funcDef *FuncDef) error {
	funcName := funcDef.FullName
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\nfunc %s(ctx *wasmlib.Sc%sContext, params *%sParams) {\n", funcName, funcKind, capitalize(funcName))
	fmt.Fprintf(file, "\tctx.Log(\"calling %s\")\n", funcDef.Name)
	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateGoFuncs() error {
	scFileName := s.Name + ".go"
	file, err := os.Open(scFileName)
	if err != nil {
		return s.GenerateGoFuncsNew(scFileName)
	}
	lines, existing, err := s.GenerateGoFuncScanner(file)
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
			err = s.GenerateGoFunc(file, funcDef)
			if err != nil {
				return err
			}
		}
	}

	// append any new views
	for _, viewDef := range s.Views {
		if existing[viewDef.FullName] == "" {
			err = s.GenerateGoFunc(file, viewDef)
			if err != nil {
				return err
			}
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) GenerateGoFuncScanner(file *os.File) ([]string, StringMap, error) {
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

func (s *Schema) GenerateGoFuncsNew(scFileName string) error {
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
		err = s.GenerateGoFunc(file, funcDef)
		if err != nil {
			return err
		}
	}
	for _, viewDef := range s.Views {
		err = s.GenerateGoFunc(file, viewDef)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Schema) GenerateGoOnLoad() error {
	file, err := os.Create("onload.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintln(file, importWasmLib)

	fmt.Fprintf(file, "func OnLoad() {\n")
	fmt.Fprintf(file, "\texports := wasmlib.NewScExports()\n")
	for _, funcDef := range s.Funcs {
		name := capitalize(funcDef.FullName)
		fmt.Fprintf(file, "\texports.AddFunc(%s, %sThunk)\n", name, funcDef.FullName)
	}
	for _, viewDef := range s.Views {
		name := capitalize(viewDef.FullName)
		fmt.Fprintf(file, "\texports.AddView(%s, %sThunk)\n", name, viewDef.FullName)
	}
	fmt.Fprintf(file, "}\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, funcDef := range s.Funcs {
		s.GenerateGoThunk(file, funcDef)
	}
	for _, viewDef := range s.Views {
		s.GenerateGoThunk(file, viewDef)
	}
	return nil
}

func (s *Schema) GenerateGoSchema() error {
	file, err := os.Create("schema.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintln(file, importWasmLib)

	fmt.Fprintf(file, "const ScName = \"%s\"\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "const ScDescription = \"%s\"\n", s.Description)
	}
	hName := coretypes.Hn(s.Name)
	fmt.Fprintf(file, "const ScHname = wasmlib.ScHname(0x%s)\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			param := s.Params[name]
			name = capitalize(param.Name)
			fmt.Fprintf(file, "const Param%s = wasmlib.Key(\"%s\")\n", name, param.Alias)
		}
	}

	if len(s.Vars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.Vars {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "const Var%s = wasmlib.Key(\"%s\")\n", name, field.Alias)
		}
	}

	if len(s.Funcs)+len(s.Views) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			fmt.Fprintf(file, "const %s = \"%s\"\n", name, funcDef.Name)
		}
		for _, viewDef := range s.Views {
			name := capitalize(viewDef.FullName)
			fmt.Fprintf(file, "const %s = \"%s\"\n", name, viewDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "const H%s = wasmlib.ScHname(0x%s)\n", name, hName.String())
		}
		for _, viewDef := range s.Views {
			name := capitalize(viewDef.FullName)
			hName = coretypes.Hn(viewDef.Name)
			fmt.Fprintf(file, "const H%s = wasmlib.ScHname(0x%s)\n", name, hName.String())
		}
	}

	return nil
}

func (s *Schema) GenerateGoTests() error {
	//TODO
	return nil
}

func (s *Schema) GenerateGoThunk(file *os.File, funcDef *FuncDef) {
	// calculate padding
	nameLen := 0
	typeLen := 0
	for _, param := range funcDef.Params {
		fldName := param.Name
		if nameLen < len(fldName) {
			nameLen = len(fldName)
		}
		fldType := param.Type
		if typeLen < len(fldType) {
			typeLen = len(fldType)
		}
	}

	funcName := capitalize(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\ntype %sParams struct {\n", funcName)
	for _, param := range funcDef.Params {
		fldName := pad(capitalize(param.Name), nameLen)
		fldType := param.Type
		if param.Comment != "" {
			fldType = pad(fldType, typeLen)
		}
		fmt.Fprintf(file, "\t%s wasmlib.ScImmutable%s%s\n", fldName, fldType, param.Comment)
	}
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\nfunc %sThunk(ctx *wasmlib.Sc%sContext) {\n", funcDef.FullName, funcKind)
	grant := funcDef.Annotations["#grant"]
	if grant != "" {
		index := strings.Index(grant, "//")
		if index >= 0 {
			fmt.Fprintf(file, "\t%s\n", grant[index:])
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
			fmt.Fprintf(file, "\tgrantee := ctx.State().GetAgentId(wasmlib.Key(\"%s\"))\n", grant)
			fmt.Fprintf(file, "\tctx.Require(grantee.Exists(), \"grantee not set: %s\")\n", grant)
			grant = fmt.Sprintf("grantee.Value()")
		}
		fmt.Fprintf(file, "\tctx.Require(ctx.From(%s), \"no permission\")\n\n", grant)
	}
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "\tp := ctx.Params()\n")
	}
	fmt.Fprintf(file, "\tparams := &%sParams{\n", funcName)
	for _, param := range funcDef.Params {
		name := capitalize(param.Name)
		field := pad(name+":", nameLen+1)
		fmt.Fprintf(file, "\t\t%s p.Get%s(Param%s),\n", field, param.Type, name)
	}
	fmt.Fprintf(file, "\t}\n")
	for _, param := range funcDef.Params {
		if !param.Optional {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "\tctx.Require(params.%s.Exists(), \"missing mandatory %s\")\n", name, param.Name)
		}
	}
	fmt.Fprintf(file, "\t%s(ctx, params)\n", funcDef.FullName)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) GenerateGoTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	file, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	// write structs
	for _, typeDef := range s.Types {
		// calculate padding
		nameLen := 0
		typeLen := 0
		for _, field := range typeDef.Fields {
			fldName := field.Name
			if nameLen < len(fldName) {
				nameLen = len(fldName)
			}
			fldType := goTypes[field.Type]
			if typeLen < len(fldType) {
				typeLen = len(fldType)
			}
		}

		fmt.Fprintf(file, "\ntype %s struct {\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			fldName := pad(capitalize(field.Name), nameLen)
			fldType := goTypes[field.Type]
			if field.Comment != "" {
				fldType = pad(fldType, typeLen)
			}
			fmt.Fprintf(file, "\t%s %s%s\n", fldName, fldType, field.Comment)
		}
		fmt.Fprintf(file, "}\n")

		// write encoder and decoder for struct
		fmt.Fprintf(file, "\nfunc New%sFromBytes(bytes []byte) *%s {\n", typeDef.Name, typeDef.Name)
		fmt.Fprintf(file, "\tdecode := wasmlib.NewBytesDecoder(bytes)\n\tdata := &%s{}\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "\tdata.%s = decode.%s()\n", name, field.Type)
		}
		fmt.Fprintf(file, "\treturn data\n}\n")

		fmt.Fprintf(file, "\nfunc (o *%s) Bytes() []byte {\n", typeDef.Name)
		fmt.Fprintf(file, "\treturn wasmlib.NewBytesEncoder().\n")
		for _, field := range typeDef.Fields {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "\t\t%s(o.%s).\n", field.Type, name)
		}
		fmt.Fprintf(file, "\t\tData()\n}\n")
	}

	return nil
}

func (s *Schema) GenerateGoWasmMain() error {
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
