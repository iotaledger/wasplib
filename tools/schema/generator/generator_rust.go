// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/iotaledger/wasp/packages/coretypes"
)

const (
	allowDeadCode      = "#![allow(dead_code)]\n"
	allowUnusedImports = "#![allow(unused_imports)]\n"
	useConsts          = "use crate::consts::*;\n"
	useCrate           = "use crate::*;\n"
	useKeys            = "use crate::keys::*;\n"
	useParams          = "use crate::params::*;\n"
	useResults         = "use crate::results::*;\n"
	useState           = "use crate::state::*;\n"
	useStdPtr          = "use std::ptr;\n"
	useSubtypes        = "use crate::subtypes::*;\n"
	useTypes           = "use crate::types::*;\n"
	useWasmLib         = "use wasmlib::*;\n"
	useWasmLibHost     = "use wasmlib::host::*;\n"
)

var rustFuncRegexp = regexp.MustCompile("^pub fn (\\w+).+$")

var rustTypes = StringMap{
	"Address":   "ScAddress",
	"AgentId":   "ScAgentId",
	"ChainId":   "ScChainId",
	"Color":     "ScColor",
	"Hash":      "ScHash",
	"Hname":     "ScHname",
	"Int16":     "i16",
	"Int32":     "i32",
	"Int64":     "i64",
	"RequestId": "ScRequestId",
	"String":    "String",
}

var rustKeyTypes = StringMap{
	"Address":   "&ScAddress",
	"AgentId":   "&ScAgentId",
	"ChainId":   "&ScChainId",
	"Color":     "&ScColor",
	"Hash":      "&ScHash",
	"Hname":     "&ScHname",
	"Int16":     "??TODO",
	"Int32":     "i32",
	"Int64":     "??TODO",
	"RequestId": "&ScRequestId",
	"String":    "&str",
}

var rustKeys = StringMap{
	"Address":   "key",
	"AgentId":   "key",
	"ChainId":   "key",
	"Color":     "key",
	"Hash":      "key",
	"Hname":     "key",
	"Int16":     "??TODO",
	"Int32":     "Key32(int32)",
	"Int64":     "??TODO",
	"RequestId": "key",
	"String":    "key",
}

var rustTypeIds = StringMap{
	"Address":   "TYPE_ADDRESS",
	"AgentId":   "TYPE_AGENT_ID",
	"ChainId":   "TYPE_CHAIN_ID",
	"Color":     "TYPE_COLOR",
	"Hash":      "TYPE_HASH",
	"Hname":     "TYPE_HNAME",
	"Int16":     "TYPE_INT16",
	"Int32":     "TYPE_INT32",
	"Int64":     "TYPE_INT64",
	"RequestId": "TYPE_REQUEST_ID",
	"String":    "TYPE_STRING",
}

func (s *Schema) GenerateRust() error {
	s.NewTypes = make(map[string]bool)

	if !s.CoreContracts {
		err := os.MkdirAll("src", 0755)
		if err != nil {
			return err
		}
		err = os.Chdir("src")
		if err != nil {
			return err
		}
		defer os.Chdir("..")
	}

	err := s.generateRustConsts()
	if err != nil {
		return err
	}
	err = s.generateRustTypes()
	if err != nil {
		return err
	}
	err = s.generateRustSubtypes()
	if err != nil {
		return err
	}
	err = s.generateRustParams()
	if err != nil {
		return err
	}
	err = s.generateRustResults()
	if err != nil {
		return err
	}
	err = s.generateRustContract()
	if err != nil {
		return err
	}

	if !s.CoreContracts {
		err = s.generateRustKeys()
		if err != nil {
			return err
		}
		err = s.generateRustState()
		if err != nil {
			return err
		}
		err = s.generateRustLib()
		if err != nil {
			return err
		}
		err = s.generateRustFuncs()
		if err != nil {
			return err
		}

		// rust-specific stuff
		return s.generateRustCargo()
	}

	return nil
}

func (s *Schema) generateRustCargo() error {
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

func (s *Schema) generateRustConsts() error {
	file, err := os.Create("consts.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprint(file, s.crateOrWasmLib(false, false))

	scName := s.Name
	if s.CoreContracts {
		// remove 'core' prefix
		scName = scName[4:]
	}
	s.appendConst("SC_NAME", "&str = \""+scName+"\"")
	if s.Description != "" {
		s.appendConst("SC_DESCRIPTION", "&str = \""+s.Description+"\"")
	}
	hName := coretypes.Hn(scName)
	s.appendConst("HSC_NAME", "ScHname = ScHname(0x"+hName.String()+")")
	s.flushRustConsts(file)

	s.generateRustConstsFields(file, s.Params, "PARAM_")
	s.generateRustConstsFields(file, s.Results, "RESULT_")
	s.generateRustConstsFields(file, s.StateVars, "STATE_")

	if len(s.Funcs) != 0 {
		for _, f := range s.Funcs {
			constName := upper(snake(f.FuncName))
			s.appendConst(constName, "&str = \""+f.String+"\"")
		}
		s.flushRustConsts(file)

		for _, f := range s.Funcs {
			constHname := "H" + upper(snake(f.FuncName))
			hName = coretypes.Hn(f.String)
			s.appendConst(constHname, "ScHname = ScHname(0x"+hName.String()+")")
		}
		s.flushRustConsts(file)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustConstsFields(file *os.File, fields []*Field, prefix string) {
	if len(fields) != 0 {
		for _, field := range fields {
			if field.Alias == "this" {
				continue
			}
			name := prefix + upper(snake(field.Name))
			value := "&str = \"" + field.Alias + "\""
			s.appendConst(name, value)
		}
		s.flushRustConsts(file)
	}
}

func (s *Schema) generateRustContract() error {
	file, err := os.Create("contract.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprintln(file, useStdPtr)
	fmt.Fprint(file, s.crateOrWasmLib(true, false))
	if !s.CoreContracts {
		fmt.Fprint(file, "\n"+useConsts)
		fmt.Fprint(file, useParams)
		fmt.Fprint(file, useResults)
	}

	for _, f := range s.Funcs {
		fmt.Fprintf(file, "\npub struct %sCall {\n", f.Type)
		fmt.Fprintf(file, "    pub func: Sc%s,\n", f.Kind)
		if len(f.Params) != 0 {
			fmt.Fprintf(file, "    pub params: Mutable%sParams,\n", f.Type)
		}
		if len(f.Results) != 0 {
			fmt.Fprintf(file, "    pub results: Immutable%sResults,\n", f.Type)
		}
		fmt.Fprintf(file, "}\n")

		fmt.Fprintf(file, "\nimpl %sCall {\n", f.Type)
		s.generateRustContractFunc(file, f, "new", "Func")
		if f.Kind == "View" {
			fmt.Fprintf(file, "\n    pub fn new_from_view(_ctx: &ScViewContext) -> %sCall {\n", f.Type)
			fmt.Fprintf(file, "        %sCall::new(&ScFuncContext {})\n", f.Type)
			fmt.Fprintf(file, "    }\n")
		}

		fmt.Fprintf(file, "}\n")
	}

	return nil
}

func (s *Schema) generateRustContractFunc(file *os.File, f *FuncDef, funcName string, funcKind string) {
	constName := upper(snake(f.FuncName))
	letMut := ""
	if len(f.Params) != 0 || len(f.Results) != 0 {
		letMut = "let mut f = "
	}
	fmt.Fprintf(file, "    pub fn %s(_ctx: &Sc%sContext) -> %sCall {\n", funcName, funcKind, f.Type)
	fmt.Fprintf(file, "        %s%sCall {\n", letMut, f.Type)
	fmt.Fprintf(file, "            func: Sc%s::new(HSC_NAME, H%s),\n", f.Kind, constName)
	paramsId := "ptr::null_mut()"
	if len(f.Params) != 0 {
		paramsId = "&mut f.params.id"
		fmt.Fprintf(file, "            params: Mutable%sParams { id: 0 },\n", f.Type)
	}
	resultsId := "ptr::null_mut()"
	if len(f.Results) != 0 {
		resultsId = "&mut f.results.id"
		fmt.Fprintf(file, "            results: Immutable%sResults { id: 0 },\n", f.Type)
	}
	fmt.Fprintf(file, "        }")
	if len(f.Params) != 0 || len(f.Results) != 0 {
		fmt.Fprintf(file, ";\n")
		fmt.Fprintf(file, "        f.func.set_ptrs(%s, %s);\n", paramsId, resultsId)
		fmt.Fprintf(file, "        f")
	}
	fmt.Fprintf(file, "\n    }\n")
}

func (s *Schema) generateRustFuncs() error {
	scFileName := s.Name + ".rs"
	file, err := os.Open(scFileName)
	if err != nil {
		// generate initial code file
		return s.generateRustFuncsNew(scFileName)
	}

	// append missing function signatures to existing code file

	lines, existing, err := s.scanExistingCode(file, rustFuncRegexp)
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
		name := snake(f.FuncName)
		if existing[name] == "" {
			s.generateRustFuncSignature(file, f)
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) generateRustFuncSignature(file *os.File, f *FuncDef) {
	fmt.Fprintf(file, "\npub fn %s(_ctx: &Sc%sContext, _f: &%sContext) {\n", snake(f.FuncName), f.Kind, capitalize(f.Type))
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateRustFuncsNew(scFileName string) error {
	file, err := os.Create(scFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(false))
	fmt.Fprintln(file, useWasmLib)

	fmt.Fprint(file, useCrate)
	if len(s.Subtypes) != 0 {
		fmt.Fprint(file, useSubtypes)
	}
	if len(s.Types) != 0 {
		fmt.Fprint(file, useTypes)
	}

	for _, f := range s.Funcs {
		s.generateRustFuncSignature(file, f)
	}
	return nil
}

func (s *Schema) generateRustKeys() error {
	file, err := os.Create("keys.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprintln(file, useWasmLib)
	fmt.Fprint(file, useCrate)

	s.KeyId = 0
	s.generateRustKeysIndexes(file, s.Params, "PARAM_")
	s.generateRustKeysIndexes(file, s.Results, "RESULT_")
	s.generateRustKeysIndexes(file, s.StateVars, "STATE_")
	s.flushRustConsts(file)

	size := s.KeyId
	fmt.Fprintf(file, "\npub const KEY_MAP_LEN: usize = %d;\n", size)
	fmt.Fprintf(file, "\npub const KEY_MAP: [&str; KEY_MAP_LEN] = [\n")
	s.generateRustKeysArray(file, s.Params, "PARAM_")
	s.generateRustKeysArray(file, s.Results, "RESULT_")
	s.generateRustKeysArray(file, s.StateVars, "STATE_")
	fmt.Fprintf(file, "];\n")

	fmt.Fprintf(file, "\npub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];\n")

	fmt.Fprintf(file, "\npub fn idx_map(idx: usize) -> Key32 {\n")
	fmt.Fprintf(file, "    unsafe {\n")
	fmt.Fprintf(file, "        IDX_MAP[idx]\n")
	fmt.Fprintf(file, "    }\n")
	fmt.Fprintf(file, "}\n")

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustKeysArray(file *os.File, fields []*Field, prefix string) {
	for _, field := range fields {
		if field.Alias == "this" {
			continue
		}
		name := prefix + upper(snake(field.Name))
		fmt.Fprintf(file, "    %s,\n", name)
		s.KeyId++
	}
}

func (s *Schema) generateRustKeysIndexes(file *os.File, fields []*Field, prefix string) {
	for _, field := range fields {
		if field.Alias == "this" {
			continue
		}
		name := "IDX_" + prefix + upper(snake(field.Name))
		field.KeyId = s.KeyId
		value := "usize = " + strconv.Itoa(field.KeyId)
		s.KeyId++
		s.appendConst(name, value)
	}
}

func (s *Schema) generateRustLib() error {
	file, err := os.Create("lib.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprintln(file, allowUnusedImports)
	fmt.Fprintf(file, "use %s::*;\n", s.Name)
	fmt.Fprint(file, useWasmLib)
	fmt.Fprintln(file, useWasmLibHost)
	fmt.Fprint(file, useConsts)
	fmt.Fprint(file, useKeys)
	fmt.Fprint(file, useParams)
	fmt.Fprint(file, useResults)
	fmt.Fprintln(file, useState)

	fmt.Fprintf(file, "mod consts;\n")
	fmt.Fprintf(file, "mod contract;\n")
	fmt.Fprintf(file, "mod keys;\n")
	fmt.Fprintf(file, "mod params;\n")
	fmt.Fprintf(file, "mod results;\n")
	fmt.Fprintf(file, "mod state;\n")
	if len(s.Subtypes) != 0 {
		fmt.Fprintf(file, "mod subtypes;\n")
	}
	if len(s.Types) != 0 {
		fmt.Fprintf(file, "mod types;\n")
	}
	fmt.Fprintf(file, "mod %s;\n", s.Name)

	fmt.Fprintf(file, "\n#[no_mangle]\n")
	fmt.Fprintf(file, "fn on_load() {\n")
	if len(s.Funcs) != 0 {
		fmt.Fprintf(file, "    let exports = ScExports::new();\n")
	}
	for _, f := range s.Funcs {
		name := snake(f.FuncName)
		fmt.Fprintf(file, "    exports.add_%s(%s, %s_thunk);\n", lower(f.Kind), upper(name), name)
	}

	fmt.Fprintf(file, "\n    unsafe {\n")
	fmt.Fprintf(file, "        for i in 0..KEY_MAP_LEN {\n")
	fmt.Fprintf(file, "            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);\n")
	fmt.Fprintf(file, "        }\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "}\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, f := range s.Funcs {
		s.generateRustThunk(file, f)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustProxy(file *os.File, field *Field, mutability string) {
	if field.Array {
		proxyType := mutability + field.Type
		arrayType := "ArrayOf" + proxyType
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fmt.Fprintf(file, "\npub type %s%s = %s;\n", mutability, field.Name, arrayType)
		}
		if s.NewTypes[arrayType] {
			// already generated this array
			return
		}
		s.NewTypes[arrayType] = true

		fmt.Fprintf(file, "\npub struct %s {\n", arrayType)
		fmt.Fprintf(file, "    pub(crate) obj_id: i32,\n")
		fmt.Fprintf(file, "}\n")

		fmt.Fprintf(file, "\nimpl %s {", arrayType)
		defer fmt.Fprintf(file, "}\n")

		if mutability == "Mutable" {
			fmt.Fprintf(file, "\n    pub fn clear(&self) {\n")
			fmt.Fprintf(file, "        clear(self.obj_id);\n")
			fmt.Fprintf(file, "    }\n")
		}

		fmt.Fprintf(file, "\n    pub fn length(&self) -> i32 {\n")
		fmt.Fprintf(file, "        get_length(self.obj_id)\n")
		fmt.Fprintf(file, "    }\n")

		if field.TypeId == 0 {
			for _, subtype := range s.Subtypes {
				if subtype.Name == field.Type {
					varType := "TYPE_MAP"
					if subtype.Array {
						varType = rustTypeIds[subtype.Type]
						if len(varType) == 0 {
							varType = "TYPE_BYTES"
						}
						varType = "TYPE_ARRAY | " + varType
					}
					fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> %s {\n", snake(field.Type), proxyType)
					fmt.Fprintf(file, "        let sub_id = get_object_id(self.obj_id, Key32(index), %s)\n", varType)
					fmt.Fprintf(file, "        %s { obj_id: sub_id }\n", proxyType)
					fmt.Fprintf(file, "    }\n")
					return
				}
			}
			fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> %s {\n", snake(field.Type), proxyType)
			fmt.Fprintf(file, "        %s { obj_id: self.obj_id, key_id: Key32(index) }\n", proxyType)
			fmt.Fprintf(file, "    }\n")
			return
		}

		fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> Sc%s {\n", snake(field.Type), proxyType)
		fmt.Fprintf(file, "        Sc%s::new(self.obj_id, Key32(index))\n", proxyType)
		fmt.Fprintf(file, "    }\n")
		return
	}

	if len(field.MapKey) != 0 {
		proxyType := mutability + field.Type
		mapType := "Map" + field.MapKey + "To" + proxyType
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fmt.Fprintf(file, "\npub type %s%s = %s;\n", mutability, field.Name, mapType)
		}
		if s.NewTypes[mapType] {
			// already generated this map
			return
		}
		s.NewTypes[mapType] = true

		keyType := rustKeyTypes[field.MapKey]
		keyValue := rustKeys[field.MapKey]

		fmt.Fprintf(file, "\npub struct %s {\n", mapType)
		fmt.Fprintf(file, "    pub(crate) obj_id: i32,\n")
		fmt.Fprintf(file, "}\n")

		fmt.Fprintf(file, "\nimpl %s {", mapType)
		defer fmt.Fprintf(file, "}\n")

		if mutability == "Mutable" {
			fmt.Fprintf(file, "\n    pub fn clear(&self) {\n")
			fmt.Fprintf(file, "        clear(self.obj_id)\n")
			fmt.Fprintf(file, "    }\n")
		}

		if field.TypeId == 0 {
			for _, subtype := range s.Subtypes {
				if subtype.Name == field.Type {
					varType := "TYPE_MAP"
					if subtype.Array {
						varType = rustTypeIds[subtype.Type]
						if len(varType) == 0 {
							varType = "TYPE_BYTES"
						}
						varType = "TYPE_ARRAY | " + varType
					}
					fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> %s {\n", snake(field.Type), keyType, proxyType)
					fmt.Fprintf(file, "        let sub_id = get_object_id(self.obj_id, %s.get_key_id(), %s);\n", keyValue, varType)
					fmt.Fprintf(file, "        %s { obj_id: sub_id }\n", proxyType)
					fmt.Fprintf(file, "    }\n")
					return
				}
			}
			fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> %s {\n", snake(field.Type), keyType, proxyType)
			fmt.Fprintf(file, "        %s { obj_id: self.obj_id, key_id: %s.get_key_id() }\n", proxyType, keyValue)
			fmt.Fprintf(file, "    }\n")
			return
		}

		fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> Sc%s {\n", snake(field.Type), keyType, proxyType)
		fmt.Fprintf(file, "        Sc%s::new(self.obj_id, %s.get_key_id())\n", proxyType, keyValue)
		fmt.Fprintf(file, "    }\n")
	}
}

func (s *Schema) generateRustState() error {
	file, err := os.Create("state.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprint(file, allowDeadCode)
	fmt.Fprintln(file, allowUnusedImports)
	fmt.Fprint(file, useWasmLib)
	fmt.Fprintln(file, useWasmLibHost)
	fmt.Fprint(file, useCrate)
	fmt.Fprint(file, useKeys)
	if len(s.Subtypes) != 0 {
		fmt.Fprint(file, useSubtypes)
	}
	if len(s.Types) != 0 {
		fmt.Fprint(file, useTypes)
	}

	s.generateRustStruct(file, s.StateVars, "Immutable", s.FullName, "State")
	s.generateRustStruct(file, s.StateVars, "Mutable", s.FullName, "State")
	return nil
}

func (s *Schema) generateRustParams() error {
	file, err := os.Create("params.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprint(file, allowDeadCode)
	fmt.Fprintln(file, allowUnusedImports)
	fmt.Fprint(file, s.crateOrWasmLib(true, true))
	if !s.CoreContracts {
		fmt.Fprint(file, "\n"+useCrate)
		fmt.Fprint(file, useKeys)
	}

	for _, f := range s.Funcs {
		if len(f.Params) == 0 {
			continue
		}
		s.generateRustStruct(file, f.Params, "Immutable", f.Type, "Params")
		s.generateRustStruct(file, f.Params, "Mutable", f.Type, "Params")
	}
	return nil
}

func (s *Schema) generateRustResults() error {
	file, err := os.Create("results.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprint(file, allowDeadCode)
	fmt.Fprintln(file, allowUnusedImports)
	fmt.Fprint(file, s.crateOrWasmLib(true, true))
	if !s.CoreContracts {
		fmt.Fprint(file, "\n"+useCrate)
		fmt.Fprint(file, useKeys)
	}

	for _, f := range s.Funcs {
		if len(f.Results) == 0 {
			continue
		}
		s.generateRustStruct(file, f.Results, "Immutable", f.Type, "Results")
		s.generateRustStruct(file, f.Results, "Mutable", f.Type, "Results")
	}
	return nil
}

func (s *Schema) generateRustStruct(file *os.File, fields []*Field, mutability string, typeName string, kind string) {
	typeName = mutability + typeName + kind
	if strings.HasSuffix(kind, "s") {
		kind = kind[0 : len(kind)-1]
	}
	kind = upper(kind) + "_"

	// first generate necessary array and map types
	for _, field := range fields {
		s.generateRustProxy(file, field, mutability)
	}

	fmt.Fprintf(file, "\n#[derive(Clone, Copy)]\n")
	fmt.Fprintf(file, "pub struct %s {\n", typeName)
	fmt.Fprintf(file, "    pub(crate) id: i32,\n")
	fmt.Fprintf(file, "}\n")

	if len(fields) != 0 {
		fmt.Fprintf(file, "\nimpl %s {", typeName)
		defer fmt.Fprintf(file, "}\n")
	}

	for _, field := range fields {
		varName := snake(field.Name)
		varId := "idx_map(IDX_" + kind + upper(varName) + ")"
		if s.CoreContracts {
			varId = kind + upper(varName) + ".get_key_id()"
		}
		varType := rustTypeIds[field.Type]
		if len(varType) == 0 {
			varType = "TYPE_BYTES"
		}
		if field.Array {
			varType = "TYPE_ARRAY | " + varType
			arrayType := "ArrayOf" + mutability + field.Type
			fmt.Fprintf(file, "\n    pub fn %s(&self) -> %s {\n", varName, arrayType)
			fmt.Fprintf(file, "        let arr_id = get_object_id(self.id, %s, %s);\n", varId, varType)
			fmt.Fprintf(file, "        %s { obj_id: arr_id }\n", arrayType)
			fmt.Fprintf(file, "    }\n")
			continue
		}
		if len(field.MapKey) != 0 {
			varType = "TYPE_MAP"
			mapType := "Map" + field.MapKey + "To" + mutability + field.Type
			fmt.Fprintf(file, "\n    pub fn %s(&self) -> %s {\n", varName, mapType)
			mapId := "self.id"
			if field.Alias != "this" {
				mapId = "map_id"
				fmt.Fprintf(file, "        let map_id = get_object_id(self.id, %s, %s);\n", varId, varType)
			}
			fmt.Fprintf(file, "        %s { obj_id: %s }\n", mapType, mapId)
			fmt.Fprintf(file, "    }\n")
			continue
		}

		proxyType := mutability + field.Type
		fmt.Fprintf(file, "\n    pub fn %s(&self) -> Sc%s {\n", varName, proxyType)
		fmt.Fprintf(file, "        Sc%s::new(self.id, %s)\n", proxyType, varId)
		fmt.Fprintf(file, "    }\n")
	}
}

func (s *Schema) generateRustSubtypes() error {
	if len(s.Subtypes) == 0 {
		return nil
	}

	file, err := os.Create("subtypes.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprint(file, useWasmLib)
	fmt.Fprint(file, useWasmLibHost)
	if len(s.Types) != 0 {
		fmt.Fprint(file, "\n", useTypes)
	}

	for _, subtype := range s.Subtypes {
		s.generateRustProxy(file, subtype, "Immutable")
		s.generateRustProxy(file, subtype, "Mutable")
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustThunk(file *os.File, f *FuncDef) {
	nameLen := 5
	if len(f.Params) != 0 {
		nameLen = 6
	}
	if len(f.Results) != 0 {
		nameLen = 7
	}

	fmt.Fprintf(file, "\npub struct %sContext {\n", f.Type)
	if len(f.Params) != 0 {
		fmt.Fprintf(file, "    %s Immutable%sParams,\n", pad("params:", nameLen+1), f.Type)
	}
	if len(f.Results) != 0 {
		fmt.Fprintf(file, "    results: Mutable%sResults,\n", f.Type)
	}
	mutability := "Mutable"
	if f.Kind == "View" {
		mutability = "Immutable"
	}
	fmt.Fprintf(file, "    %s %s%sState,\n", pad("state:", nameLen+1), mutability, s.FullName)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfn %s_thunk(ctx: &Sc%sContext) {\n", snake(f.FuncName), f.Kind)
	fmt.Fprintf(file, "    ctx.log(\"%s.%s\");\n", s.Name, f.FuncName)
	grant := f.Access
	if grant != "" {
		index := strings.Index(grant, "//")
		if index >= 0 {
			fmt.Fprintf(file, "    %s\n", grant[index:])
			grant = strings.TrimSpace(grant[:index])
		}
		switch grant {
		case "self":
			grant = "ctx.account_id()"
		case "chain":
			grant = "ctx.chain_owner_id()"
		case "creator":
			grant = "ctx.contract_creator()"
		default:
			fmt.Fprintf(file, "    let access = ctx.state().get_agent_id(\"%s\");\n", grant)
			fmt.Fprintf(file, "    ctx.require(access.exists(), \"access not set: %s\");\n", grant)
			grant = fmt.Sprintf("access.value()")
		}
		fmt.Fprintf(file, "    ctx.require(ctx.caller() == %s, \"no permission\");\n\n", grant)
	}

	fmt.Fprintf(file, "    let f = %sContext {\n", f.Type)

	if len(f.Params) != 0 {
		fmt.Fprintf(file, "        params: Immutable%sParams {\n", f.Type)
		fmt.Fprintf(file, "            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),\n")
		fmt.Fprintf(file, "        },\n")
	}
	if len(f.Results) != 0 {
		fmt.Fprintf(file, "        results: Mutable%sResults {\n", f.Type)
		fmt.Fprintf(file, "            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),\n")
		fmt.Fprintf(file, "        },\n")
	}

	fmt.Fprintf(file, "        state: %s%sState {\n", mutability, s.FullName)
	fmt.Fprintf(file, "            id: get_object_id(1, KEY_STATE, TYPE_MAP),\n")
	fmt.Fprintf(file, "        },\n")

	fmt.Fprintf(file, "    };\n")

	for _, param := range f.Params {
		if !param.Optional {
			name := snake(param.Name)
			fmt.Fprintf(file, "    ctx.require(f.params.%s().exists(), \"missing mandatory %s\");\n", name, param.Name)
		}
	}

	fmt.Fprintf(file, "    %s(ctx, &f);\n", snake(f.FuncName))
	fmt.Fprintf(file, "    ctx.log(\"%s.%s ok\");\n", s.Name, f.FuncName)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateRustTypes() error {
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
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprint(file, useWasmLib)
	fmt.Fprint(file, useWasmLibHost)

	// write structs
	for _, typeDef := range s.Types {
		s.generateRustType(file, typeDef)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustType(file *os.File, typeDef *TypeDef) {
	nameLen, typeLen := calculatePadding(typeDef.Fields, rustTypes, true)

	fmt.Fprintf(file, "\npub struct %s {\n", typeDef.Name)
	for _, field := range typeDef.Fields {
		fldName := pad(snake(field.Name)+":", nameLen+1)
		fldType := rustTypes[field.Type] + ","
		if field.Comment != "" {
			fldType = pad(fldType, typeLen+1)
		}
		fmt.Fprintf(file, "    pub %s %s%s\n", fldName, fldType, field.Comment)
	}
	fmt.Fprintf(file, "}\n")

	// write encoder and decoder for struct
	fmt.Fprintf(file, "\nimpl %s {", typeDef.Name)

	fmt.Fprintf(file, "\n    pub fn from_bytes(bytes: &[u8]) -> %s {\n", typeDef.Name)
	fmt.Fprintf(file, "        let mut decode = BytesDecoder::new(bytes);\n")
	fmt.Fprintf(file, "        %s {\n", typeDef.Name)
	for _, field := range typeDef.Fields {
		name := snake(field.Name)
		fmt.Fprintf(file, "            %s: decode.%s(),\n", name, snake(field.Type))
	}
	fmt.Fprintf(file, "        }\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn to_bytes(&self) -> Vec<u8> {\n")
	fmt.Fprintf(file, "        let mut encode = BytesEncoder::new();\n")
	for _, field := range typeDef.Fields {
		name := snake(field.Name)
		ref := "&"
		if field.Type == "Hname" || field.Type == "Int64" || field.Type == "Int32" || field.Type == "Int16" {
			ref = ""
		}
		fmt.Fprintf(file, "        encode.%s(%sself.%s);\n", snake(field.Type), ref, name)
	}
	fmt.Fprintf(file, "        return encode.data();\n")
	fmt.Fprintf(file, "    }\n")
	fmt.Fprintf(file, "}\n")

	s.generateRustTypeProxy(file, typeDef, false)
	s.generateRustTypeProxy(file, typeDef, true)
}

func (s *Schema) generateRustTypeProxy(file *os.File, typeDef *TypeDef, mutable bool) {
	typeName := "Immutable" + typeDef.Name
	if mutable {
		typeName = "Mutable" + typeDef.Name
	}

	fmt.Fprintf(file, "\npub struct %s {\n", typeName)
	fmt.Fprintf(file, "    pub(crate) obj_id: i32,\n")
	fmt.Fprintf(file, "    pub(crate) key_id: Key32,\n")
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nimpl %s {", typeName)

	fmt.Fprintf(file, "\n    pub fn exists(&self) -> bool {\n")
	fmt.Fprintf(file, "        exists(self.obj_id, self.key_id, TYPE_BYTES)\n")
	fmt.Fprintf(file, "    }\n")

	if mutable {
		fmt.Fprintf(file, "\n    pub fn set_value(&self, value: &%s) {\n", typeDef.Name)
		fmt.Fprintf(file, "        set_bytes(self.obj_id, self.key_id, TYPE_BYTES, &value.to_bytes());\n")
		fmt.Fprintf(file, "    }\n")
	}

	fmt.Fprintf(file, "\n    pub fn value(&self) -> %s {\n", typeDef.Name)
	fmt.Fprintf(file, "        %s::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_BYTES))\n", typeDef.Name)
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "}\n")
}

func (s *Schema) flushRustConsts(file *os.File) {
	if len(s.ConstNames) == 0 {
		return
	}

	fmt.Fprintln(file)
	s.flushConsts(file, func(name string, value string, padLen int) {
		fmt.Fprintf(file, "pub const %s %s;\n", pad(name+":", padLen+1), value)
	})
}
