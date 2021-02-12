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

var rustFuncRegexp = regexp.MustCompile("^pub fn (\\w+).+$")

var rustTypes = StringMap{
	"Address":    "ScAddress",
	"AgentId":    "ScAgentId",
	"ChainId":    "ScChainId",
	"Color":      "ScColor",
	"ContractId": "ScContractId",
	"Hash":       "ScHash",
	"Hname":      "ScHname",
	"Int":        "i64",
	"String":     "String",
}

func (s *Schema) GenerateRust() error {
	err := os.MkdirAll("src", 0755)
	if err != nil {
		return err
	}
	err = os.Chdir("src")
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = s.GenerateRustLib()
	if err != nil {
		return err
	}
	err = s.GenerateRustSchema()
	if err != nil {
		return err
	}
	err = s.GenerateRustTypes()
	if err != nil {
		return err
	}
	err = s.GenerateRustFuncs()
	if err != nil {
		return err
	}
	return s.GenerateRustCargo()
}

func (s *Schema) GenerateRustCargo() error {
	file, err := os.Open("../Cargo.toml")
	if err == nil {
		// already exists
		file.Close()
		return nil
	}

	file, err = os.Create("../Cargo.toml")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "[package]\n")
	fmt.Fprintf(file, "name = \"%s\"\n", s.Name)
	fmt.Fprintf(file, "description = \"%s\"\n", s.Description)
	fmt.Fprintf(file, "license = \"Apache-2.0\"\n")
	fmt.Fprintf(file, "version = \"0.1.0\"\n")
	fmt.Fprintf(file, "authors = [\"Eric Hop <eric@iota.org>\"]\n")
	fmt.Fprintf(file, "edition = \"2018\"\n")
	fmt.Fprintf(file, "repository = \"https://%s\"\n", ModuleName)
	fmt.Fprintf(file, "\n[lib]\n")
	fmt.Fprintf(file, "crate-type = [\"cdylib\", \"rlib\"]\n")
	fmt.Fprintf(file, "\n[features]\n")
	fmt.Fprintf(file, "default = [\"console_error_panic_hook\"]\n")
	fmt.Fprintf(file, "\n[dependencies]\n")
	fmt.Fprintf(file, "wasmlib = { git = \"https://github.com/iotaledger/wasp\", branch = \"develop\" }\n")
	fmt.Fprintf(file, "console_error_panic_hook = { version = \"0.1.6\", optional = true }\n")
	fmt.Fprintf(file, "wee_alloc = { version = \"0.4.5\", optional = true }\n")
	fmt.Fprintf(file, "\n[dev-dependencies]\n")
	fmt.Fprintf(file, "wasm-bindgen-test = \"0.3.13\"\n")

	return nil
}

func (s *Schema) GenerateRustFunc(file *os.File, funcDef *FuncDef) error {
	funcName := snake(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\npub fn %s(ctx: &Sc%sContext, params: &%sParams) {\n", funcName, funcKind, capitalize(funcDef.FullName))
	fmt.Fprintf(file, "    ctx.log(\"calling %s\");\n", funcDef.Name)
	fmt.Fprintf(file, "}\n")
	return nil
}

func (s *Schema) GenerateRustFuncs() error {
	scFileName := s.Name + ".rs"
	file, err := os.Open(scFileName)
	if err != nil {
		return s.GenerateRustFuncsNew(scFileName)
	}
	lines, existing, err := s.GenerateRustFuncScanner(file)
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
		name := snake(funcDef.FullName)
		if existing[name] == "" {
			err = s.GenerateRustFunc(file, funcDef)
			if err != nil {
				return err
			}
		}
	}

	// append any new views
	for _, viewDef := range s.Views {
		name := snake(viewDef.FullName)
		if existing[name] == "" {
			err = s.GenerateRustFunc(file, viewDef)
			if err != nil {
				return err
			}
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) GenerateRustFuncScanner(file *os.File) ([]string, StringMap, error) {
	defer file.Close()
	existing := make(StringMap)
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := rustFuncRegexp.FindStringSubmatch(line)
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

func (s *Schema) GenerateRustFuncsNew(scFileName string) error {
	file, err := os.Create(scFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(false))
	fmt.Fprintf(file, "use wasmlib::*;\n\n")

	fmt.Fprintf(file, "use crate::*;\n")
	if len(s.Types) != 0 {
		fmt.Fprintf(file, "use crate::types::*;\n")
	}

	for _, funcDef := range s.Funcs {
		err = s.GenerateRustFunc(file, funcDef)
		if err != nil {
			return err
		}
	}
	for _, viewDef := range s.Views {
		err = s.GenerateRustFunc(file, viewDef)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Schema) GenerateRustLib() error {
	file, err := os.Create("lib.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "use %s::*;\n", s.Name)
	fmt.Fprintf(file, "use schema::*;\n")
	fmt.Fprintf(file, "use wasmlib::*;\n\n")

	fmt.Fprintf(file, "mod %s;\n", s.Name)
	fmt.Fprintf(file, "mod schema;\n")
	if len(s.Types) != 0 {
		fmt.Fprintf(file, "mod types;\n")
	}

	fmt.Fprintf(file, "\n#[no_mangle]\n")
	fmt.Fprintf(file, "fn on_load() {\n")
	fmt.Fprintf(file, "    let exports = ScExports::new();\n")
	for _, funcDef := range s.Funcs {
		name := snake(funcDef.FullName)
		fmt.Fprintf(file, "    exports.add_func(%s, %s_thunk);\n", upper(name), name)
	}
	for _, viewDef := range s.Views {
		name := snake(viewDef.FullName)
		fmt.Fprintf(file, "    exports.add_view(%s, %s_thunk);\n", upper(name), name)
	}
	fmt.Fprintf(file, "}\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, funcDef := range s.Funcs {
		s.GenerateRustThunk(file, funcDef)
	}
	for _, viewDef := range s.Views {
		s.GenerateRustThunk(file, viewDef)
	}
	return nil
}

func (s *Schema) GenerateRustSchema() error {
	file, err := os.Create("schema.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "#![allow(dead_code)]\n\n")
	fmt.Fprintf(file, "use wasmlib::*;\n\n")

	fmt.Fprintf(file, "pub const SC_NAME: &str = \"%s\";\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "pub const SC_DESCRIPTION: &str = \"%s\";\n", s.Description)
	}
	hName := coretypes.Hn(s.Name)
	fmt.Fprintf(file, "pub const SC_HNAME: ScHname = ScHname(0x%s);\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			param := s.Params[name]
			name = upper(snake(name))
			fmt.Fprintf(file, "pub const PARAM_%s: &str = \"%s\";\n", name, param.Alias)
		}
	}

	if len(s.Vars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.Vars {
			name := upper(snake(field.Name))
			fmt.Fprintf(file, "pub const VAR_%s: &str = \"%s\";\n", name, field.Alias)
		}
	}

	if len(s.Funcs)+len(s.Views) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := upper(snake(funcDef.FullName))
			fmt.Fprintf(file, "pub const %s: &str = \"%s\";\n", name, funcDef.Name)
		}
		for _, viewDef := range s.Views {
			name := upper(snake(viewDef.FullName))
			fmt.Fprintf(file, "pub const %s: &str = \"%s\";\n", name, viewDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := upper(snake(funcDef.FullName))
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "pub const H%s: ScHname = ScHname(0x%s);\n", name, hName.String())
		}
		for _, viewDef := range s.Views {
			name := upper(snake(viewDef.FullName))
			hName = coretypes.Hn(viewDef.Name)
			fmt.Fprintf(file, "pub const H%s: ScHname = ScHname(0x%s);\n", name, hName.String())
		}
	}
	return nil
}

func (s *Schema) GenerateRustThunk(file *os.File, funcDef *FuncDef) {
	// calculate padding
	nameLen := 0
	typeLen := 0
	for _, param := range funcDef.Params {
		fldName := snake(param.Name)
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
	fmt.Fprintln(file)
	if len(funcDef.Params) > 1 {
		fmt.Fprintf(file, "//@formatter:off\n")
	}
	fmt.Fprintf(file, "pub struct %sParams {", funcName)
	if len(funcDef.Params) != 0 {
		fmt.Fprintln(file)
		for _, param := range funcDef.Params {
			fldName := pad(snake(param.Name) + ":", nameLen+1)
			fldType := param.Type + ","
			if param.Comment != "" {
				fldType= pad(fldType, typeLen+1)
			}
			fmt.Fprintf(file, "    pub %s ScImmutable%s%s\n", fldName, fldType, param.Comment)
		}
	}
	fmt.Fprintf(file, "}\n")
	if len(funcDef.Params) > 1 {
		fmt.Fprintf(file, "//@formatter:on\n")
	}

	fmt.Fprintf(file, "\nfn %s_thunk(ctx: &Sc%sContext) {\n", snake(funcDef.FullName), funcKind)
	grant := funcDef.Annotations["#grant"]
	if grant != "" {
		index := strings.Index(grant, "//")
		if index >= 0 {
			fmt.Fprintf(file, "    %s\n", grant[index:])
			grant = strings.TrimSpace(grant[:index])
		}
		switch grant {
		case "self":
			grant = "ctx.contract_id().as_agent_id()"
		case "owner":
			grant = "ctx.chain_owner_id()"
		case "creator":
			grant = "ctx.contract_creator()"
		default:
			fmt.Fprintf(file, "    let grantee = ctx.state().get_agent_id(\"%s\");\n", grant)
			fmt.Fprintf(file, "    ctx.require(grantee.exists(), \"grantee not set: %s\");\n", grant)
			grant = fmt.Sprintf("grantee.value()")
		}
		fmt.Fprintf(file, "    ctx.require(ctx.from(&%s), \"no permission\");\n\n", grant)
	}
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "    let p = ctx.params();\n")
	}
	fmt.Fprintf(file, "    let params = %sParams {", funcName)
	if len(funcDef.Params) != 0 {
		fmt.Fprintln(file)
		for _, param := range funcDef.Params {
			name := snake(param.Name)
			fmt.Fprintf(file, "        %s: p.get_%s(PARAM_%s),\n", name, snake(param.Type), upper(name))
		}
		fmt.Fprintf(file, "    ")
	}
	fmt.Fprintf(file, "};\n")
	for _, param := range funcDef.Params {
		if !param.Optional {
			name := snake(param.Name)
			fmt.Fprintf(file, "    ctx.require(params.%s.exists(), \"missing mandatory %s\");\n", name, param.Name)
		}
	}
	fmt.Fprintf(file, "    %s(ctx, &params);\n", snake(funcDef.FullName))
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) GenerateRustTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	file, err := os.Create("types.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "#![allow(dead_code)]\n\n")
	fmt.Fprintf(file, "use wasmlib::*;\n")

	// write structs
	for _, typeDef := range s.Types {
		// calculate padding
		nameLen := 0
		typeLen := 0
		for _, field := range typeDef.Fields {
			fldName := snake(field.Name)
			if nameLen < len(fldName) {
				nameLen = len(fldName)
			}
			fldType := rustTypes[field.Type]
			if typeLen < len(fldType) {
				typeLen = len(fldType)
			}
		}

		// write struct
		if len(typeDef.Fields) > 1 {
			fmt.Fprintf(file, "\n//@formatter:off\n")
		}
		fmt.Fprintf(file, "pub struct %s {\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			fldName := pad(snake(field.Name) + ":", nameLen+1)
			fldType := rustTypes[field.Type] + ","
			if field.Comment != "" {
				fldType= pad(fldType, typeLen+1)
			}
			fmt.Fprintf(file, "    pub %s %s%s\n", fldName, fldType, field.Comment)
		}
		fmt.Fprintf(file, "}\n")
		if len(typeDef.Fields) > 1 {
			fmt.Fprintf(file, "//@formatter:on\n")
		}

		// write encoder and decoder for struct
		fmt.Fprintf(file, "\nimpl %s {\n", typeDef.Name)

		fmt.Fprintf(file, "    pub fn from_bytes(bytes: &[u8]) -> %s {\n", typeDef.Name)
		fmt.Fprintf(file, "        let mut decode = BytesDecoder::new(bytes);\n")
		fmt.Fprintf(file, "        %s {\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			name := snake(field.Name)
			fmt.Fprintf(file, "            %s: decode.%s(),\n", name, snake(field.Type))
		}
		fmt.Fprintf(file, "        }\n")
		fmt.Fprintf(file, "    }\n\n")

		fmt.Fprintf(file, "    pub fn to_bytes(&self) -> Vec<u8> {\n")
		fmt.Fprintf(file, "        let mut encode = BytesEncoder::new();\n")
		for _, field := range typeDef.Fields {
			name := snake(field.Name)
			ref := "&"
			if field.Type == "Int" || field.Type == "Hname" {
				ref = ""
			}
			fmt.Fprintf(file, "        encode.%s(%sself.%s);\n", snake(field.Type), ref, name)
		}
		fmt.Fprintf(file, "        return encode.data();\n")
		fmt.Fprintf(file, "    }\n")
		fmt.Fprintf(file, "}\n")
	}

	return nil
}
