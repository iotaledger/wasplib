// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/iotaledger/wasp/packages/iscp"
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
	useSubtypes        = "use crate::typedefs::*;\n"
	useTypes           = "use crate::types::*;\n"
	useWasmLib         = "use wasmlib::*;\n"
	useWasmLibHost     = "use wasmlib::host::*;\n"
)

var rustFuncRegexp = regexp.MustCompile(`^pub fn (\w+).+$`)

var rustTypes = StringMap{
	"Address":   "ScAddress",
	"AgentID":   "ScAgentID",
	"ChainID":   "ScChainID",
	"Color":     "ScColor",
	"Hash":      "ScHash",
	"Hname":     "ScHname",
	"Int16":     "i16",
	"Int32":     "i32",
	"Int64":     "i64",
	"RequestID": "ScRequestID",
	"String":    "String",
}

var rustKeyTypes = StringMap{
	"Address":   "&ScAddress",
	"AgentID":   "&ScAgentID",
	"ChainID":   "&ScChainID",
	"Color":     "&ScColor",
	"Hash":      "&ScHash",
	"Hname":     "&ScHname",
	"Int16":     "??TODO",
	"Int32":     "i32",
	"Int64":     "??TODO",
	"RequestID": "&ScRequestID",
	"String":    "&str",
}

var rustKeys = StringMap{
	"Address":   "key",
	"AgentID":   "key",
	"ChainID":   "key",
	"Color":     "key",
	"Hash":      "key",
	"Hname":     "key",
	"Int16":     "??TODO",
	"Int32":     "Key32(int32)",
	"Int64":     "??TODO",
	"RequestID": "key",
	"String":    "key",
}

var rustTypeIds = StringMap{
	"Address":   "TYPE_ADDRESS",
	"AgentID":   "TYPE_AGENT_ID",
	"ChainID":   "TYPE_CHAIN_ID",
	"Color":     "TYPE_COLOR",
	"Hash":      "TYPE_HASH",
	"Hname":     "TYPE_HNAME",
	"Int16":     "TYPE_INT16",
	"Int32":     "TYPE_INT32",
	"Int64":     "TYPE_INT64",
	"RequestID": "TYPE_REQUEST_ID",
	"String":    "TYPE_STRING",
}

const (
	rustTypeBytes = "TYPE_BYTES"
	rustTypeMap   = "TYPE_MAP"
)

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
		defer func() {
			_ = os.Chdir("..")
		}()
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
	hName := iscp.Hn(scName)
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
			hName = iscp.Hn(f.String)
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
			if field.Alias == AliasThis {
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
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprintln(file, useStdPtr)
	fmt.Fprint(file, s.crateOrWasmLib(true, false))
	if !s.CoreContracts {
		fmt.Fprint(file, "\n"+useConsts)
		fmt.Fprint(file, useParams)
		fmt.Fprint(file, useResults)
	}

	for _, f := range s.Funcs {
		nameLen := f.nameLen(4) + 1
		kind := f.Kind
		if f.Type == InitFunc {
			kind = f.Type + f.Kind
		}
		fmt.Fprintf(file, "\npub struct %sCall {\n", f.Type)
		fmt.Fprintf(file, "    pub %s Sc%s,\n", pad("func:", nameLen), kind)
		if len(f.Params) != 0 {
			fmt.Fprintf(file, "    pub %s Mutable%sParams,\n", pad("params:", nameLen), f.Type)
		}
		if len(f.Results) != 0 {
			fmt.Fprintf(file, "    pub results: Immutable%sResults,\n", f.Type)
		}
		fmt.Fprintf(file, "}\n")
	}

	s.generateRustContractFuncs(file)
	formatter(file, true)
	return nil
}

func (s *Schema) generateRustContractFuncs(file *os.File) {
	fmt.Fprint(file, "\npub struct ScFuncs {\n")
	fmt.Fprint(file, "}\n")
	fmt.Fprint(file, "\nimpl ScFuncs {\n")

	for _, f := range s.Funcs {
		nameLen := f.nameLen(4) + 1
		funcName := snake(f.FuncName)
		constName := upper(funcName)
		letMut := ""
		if len(f.Params) != 0 || len(f.Results) != 0 {
			letMut = "let mut f = "
		}
		kind := f.Kind
		if f.Type == InitFunc {
			kind = f.Type + f.Kind
		}
		fmt.Fprintf(file, "    pub fn %s(_ctx: & dyn Sc%sCallContext) -> %sCall {\n", funcName[5:], f.Kind, f.Type)
		fmt.Fprintf(file, "        %s%sCall {\n", letMut, f.Type)
		fmt.Fprintf(file, "            %s Sc%s::new(HSC_NAME, H%s),\n", pad("func:", nameLen), kind, constName)
		paramsID := "ptr::null_mut()"
		if len(f.Params) != 0 {
			paramsID = "&mut f.params.id"
			fmt.Fprintf(file, "            %s Mutable%sParams { id: 0 },\n", pad("params:", nameLen), f.Type)
		}
		resultsID := "ptr::null_mut()"
		if len(f.Results) != 0 {
			resultsID = "&mut f.results.id"
			fmt.Fprintf(file, "            results: Immutable%sResults { id: 0 },\n", f.Type)
		}
		fmt.Fprintf(file, "        }")
		if len(f.Params) != 0 || len(f.Results) != 0 {
			fmt.Fprintf(file, ";\n")
			fmt.Fprintf(file, "        f.func.set_ptrs(%s, %s);\n", paramsID, resultsID)
			fmt.Fprintf(file, "        f")
		}
		fmt.Fprintf(file, "\n    }\n")
	}
	fmt.Fprintf(file, "}\n")
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
	switch f.FuncName {
	case "funcInit":
		fmt.Fprintf(file, "\npub fn %s(ctx: &Sc%sContext, f: &%sContext) {\n", snake(f.FuncName), f.Kind, capitalize(f.Type))
		fmt.Fprintf(file, "    if f.params.owner().exists() {\n")
		fmt.Fprintf(file, "        f.state.owner().set_value(&f.params.owner().value());\n")
		fmt.Fprintf(file, "        return;\n")
		fmt.Fprintf(file, "    }\n")
		fmt.Fprintf(file, "    f.state.owner().set_value(&ctx.contract_creator());\n")
	case "funcSetOwner":
		fmt.Fprintf(file, "\npub fn %s(_ctx: &Sc%sContext, f: &%sContext) {\n", snake(f.FuncName), f.Kind, capitalize(f.Type))
		fmt.Fprintf(file, "    f.state.owner().set_value(&f.params.owner().value());\n")
	case "viewGetOwner":
		fmt.Fprintf(file, "\npub fn %s(_ctx: &Sc%sContext, f: &%sContext) {\n", snake(f.FuncName), f.Kind, capitalize(f.Type))
		fmt.Fprintf(file, "    f.results.owner().set_value(&f.state.owner().value());\n")
	default:
		fmt.Fprintf(file, "\npub fn %s(_ctx: &Sc%sContext, _f: &%sContext) {\n", snake(f.FuncName), f.Kind, capitalize(f.Type))
	}
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
	if len(s.Typedefs) != 0 {
		fmt.Fprint(file, useSubtypes)
	}
	if len(s.Structs) != 0 {
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

	s.KeyID = 0
	s.generateRustKeysIndexes(s.Params, "PARAM_")
	s.generateRustKeysIndexes(s.Results, "RESULT_")
	s.generateRustKeysIndexes(s.StateVars, "STATE_")
	s.flushRustConsts(file)

	size := s.KeyID
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
		if field.Alias == AliasThis {
			continue
		}
		name := prefix + upper(snake(field.Name))
		fmt.Fprintf(file, "    %s,\n", name)
		s.KeyID++
	}
}

func (s *Schema) generateRustKeysIndexes(fields []*Field, prefix string) {
	for _, field := range fields {
		if field.Alias == AliasThis {
			continue
		}
		name := "IDX_" + prefix + upper(snake(field.Name))
		field.KeyID = s.KeyID
		value := "usize = " + strconv.Itoa(field.KeyID)
		s.KeyID++
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
	if len(s.Typedefs) != 0 {
		fmt.Fprintf(file, "mod typedefs;\n")
	}
	if len(s.Structs) != 0 {
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
		s.generateRustProxyArray(file, field, mutability)
		return
	}

	if field.MapKey != "" {
		s.generateRustProxyMap(file, field, mutability)
	}
}

func (s *Schema) generateRustProxyArray(file *os.File, field *Field, mutability string) {
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

	if mutability == PropMutable {
		fmt.Fprintf(file, "\n    pub fn clear(&self) {\n")
		fmt.Fprintf(file, "        clear(self.obj_id);\n")
		fmt.Fprintf(file, "    }\n")
	}

	fmt.Fprintf(file, "\n    pub fn length(&self) -> i32 {\n")
	fmt.Fprintf(file, "        get_length(self.obj_id)\n")
	fmt.Fprintf(file, "    }\n")

	if field.TypeID == 0 {
		s.generateRustProxyArrayNewType(file, field, proxyType)
		return
	}

	// array of predefined type
	fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> Sc%s {\n", snake(field.Type), proxyType)
	fmt.Fprintf(file, "        Sc%s::new(self.obj_id, Key32(index))\n", proxyType)
	fmt.Fprintf(file, "    }\n")
}

func (s *Schema) generateRustProxyArrayNewType(file *os.File, field *Field, proxyType string) {
	for _, subtype := range s.Typedefs {
		if subtype.Name != field.Type {
			continue
		}
		varType := rustTypeMap
		if subtype.Array {
			varType = rustTypeIds[subtype.Type]
			if varType == "" {
				varType = rustTypeBytes
			}
			varType = "TYPE_ARRAY | " + varType
		}
		fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> %s {\n", snake(field.Type), proxyType)
		fmt.Fprintf(file, "        let sub_id = get_object_id(self.obj_id, Key32(index), %s)\n", varType)
		fmt.Fprintf(file, "        %s { obj_id: sub_id }\n", proxyType)
		fmt.Fprintf(file, "    }\n")
		return
	}

	fmt.Fprintf(file, "\n    pub fn get_%s(&self, index: i32) -> %s {\n", snake(field.Type), proxyType)
	fmt.Fprintf(file, "        %s { obj_id: self.obj_id, key_id: Key32(index) }\n", proxyType)
	fmt.Fprintf(file, "    }\n")
}

func (s *Schema) generateRustProxyMap(file *os.File, field *Field, mutability string) {
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

	if mutability == PropMutable {
		fmt.Fprintf(file, "\n    pub fn clear(&self) {\n")
		fmt.Fprintf(file, "        clear(self.obj_id)\n")
		fmt.Fprintf(file, "    }\n")
	}

	if field.TypeID == 0 {
		s.generateRustProxyMapNewType(file, field, proxyType, keyType, keyValue)
		return
	}

	// map of predefined type
	fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> Sc%s {\n", snake(field.Type), keyType, proxyType)
	fmt.Fprintf(file, "        Sc%s::new(self.obj_id, %s.get_key_id())\n", proxyType, keyValue)
	fmt.Fprintf(file, "    }\n")
}

func (s *Schema) generateRustProxyMapNewType(file *os.File, field *Field, proxyType, keyType, keyValue string) {
	for _, subtype := range s.Typedefs {
		if subtype.Name != field.Type {
			continue
		}
		varType := rustTypeMap
		if subtype.Array {
			varType = rustTypeIds[subtype.Type]
			if varType == "" {
				varType = rustTypeBytes
			}
			varType = "TYPE_ARRAY | " + varType
		}
		fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> %s {\n", snake(field.Type), keyType, proxyType)
		fmt.Fprintf(file, "        let sub_id = get_object_id(self.obj_id, %s.get_key_id(), %s);\n", keyValue, varType)
		fmt.Fprintf(file, "        %s { obj_id: sub_id }\n", proxyType)
		fmt.Fprintf(file, "    }\n")
		return
	}

	fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: %s) -> %s {\n", snake(field.Type), keyType, proxyType)
	fmt.Fprintf(file, "        %s { obj_id: self.obj_id, key_id: %s.get_key_id() }\n", proxyType, keyValue)
	fmt.Fprintf(file, "    }\n")
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
	if len(s.Typedefs) != 0 {
		fmt.Fprint(file, useSubtypes)
	}
	if len(s.Structs) != 0 {
		fmt.Fprint(file, useTypes)
	}

	s.generateRustStruct(file, s.StateVars, PropImmutable, s.FullName, "State")
	s.generateRustStruct(file, s.StateVars, PropMutable, s.FullName, "State")
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
		params := make([]*Field, 0, len(f.Params))
		for _, param := range f.Params {
			if param.Alias != "@" {
				params = append(params, param)
			}
		}
		if len(params) == 0 {
			continue
		}
		s.generateRustStruct(file, params, PropImmutable, f.Type, "Params")
		s.generateRustStruct(file, params, PropMutable, f.Type, "Params")
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
		s.generateRustStruct(file, f.Results, PropImmutable, f.Type, "Results")
		s.generateRustStruct(file, f.Results, PropMutable, f.Type, "Results")
	}
	return nil
}

func (s *Schema) generateRustStruct(file *os.File, fields []*Field, mutability, typeName, kind string) {
	typeName = mutability + typeName + kind
	kind = strings.TrimSuffix(kind, "s")
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
		varID := "idx_map(IDX_" + kind + upper(varName) + ")"
		if s.CoreContracts {
			varID = kind + upper(varName) + ".get_key_id()"
		}
		varType := rustTypeIds[field.Type]
		if varType == "" {
			varType = rustTypeBytes
		}
		if field.Array {
			varType = "TYPE_ARRAY | " + varType
			arrayType := "ArrayOf" + mutability + field.Type
			fmt.Fprintf(file, "\n    pub fn %s(&self) -> %s {\n", varName, arrayType)
			fmt.Fprintf(file, "        let arr_id = get_object_id(self.id, %s, %s);\n", varID, varType)
			fmt.Fprintf(file, "        %s { obj_id: arr_id }\n", arrayType)
			fmt.Fprintf(file, "    }\n")
			continue
		}
		if field.MapKey != "" {
			varType = rustTypeMap
			mapType := "Map" + field.MapKey + "To" + mutability + field.Type
			fmt.Fprintf(file, "\n    pub fn %s(&self) -> %s {\n", varName, mapType)
			mapID := "self.id"
			if field.Alias != AliasThis {
				mapID = "map_id"
				fmt.Fprintf(file, "        let map_id = get_object_id(self.id, %s, %s);\n", varID, varType)
			}
			fmt.Fprintf(file, "        %s { obj_id: %s }\n", mapType, mapID)
			fmt.Fprintf(file, "    }\n")
			continue
		}

		proxyType := mutability + field.Type
		fmt.Fprintf(file, "\n    pub fn %s(&self) -> Sc%s {\n", varName, proxyType)
		fmt.Fprintf(file, "        Sc%s::new(self.id, %s)\n", proxyType, varID)
		fmt.Fprintf(file, "    }\n")
	}
}

func (s *Schema) generateRustSubtypes() error {
	if len(s.Typedefs) == 0 {
		return nil
	}

	file, err := os.Create("typedefs.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	formatter(file, false)
	fmt.Fprintln(file, allowDeadCode)
	fmt.Fprint(file, useWasmLib)
	fmt.Fprint(file, useWasmLibHost)
	if len(s.Structs) != 0 {
		fmt.Fprint(file, "\n", useTypes)
	}

	for _, subtype := range s.Typedefs {
		s.generateRustProxy(file, subtype, PropImmutable)
		s.generateRustProxy(file, subtype, PropMutable)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustThunk(file *os.File, f *FuncDef) {
	nameLen := f.nameLen(5) + 1
	fmt.Fprintf(file, "\npub struct %sContext {\n", f.Type)
	if len(f.Params) != 0 {
		fmt.Fprintf(file, "    %s Immutable%sParams,\n", pad("params:", nameLen), f.Type)
	}
	if len(f.Results) != 0 {
		fmt.Fprintf(file, "    results: Mutable%sResults,\n", f.Type)
	}
	mutability := PropMutable
	if f.Kind == KindView {
		mutability = PropImmutable
	}
	fmt.Fprintf(file, "    %s %s%sState,\n", pad("state:", nameLen), mutability, s.FullName)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfn %s_thunk(ctx: &Sc%sContext) {\n", snake(f.FuncName), f.Kind)
	fmt.Fprintf(file, "    ctx.log(\"%s.%s\");\n", s.Name, f.FuncName)

	if f.Access != "" {
		s.generateRustThunkAccessCheck(file, f)
	}

	fmt.Fprintf(file, "    let f = %sContext {\n", f.Type)

	if len(f.Params) != 0 {
		fmt.Fprintf(file, "        params: Immutable%sParams {\n", f.Type)
		fmt.Fprintf(file, "            id: OBJ_ID_PARAMS,\n")
		fmt.Fprintf(file, "        },\n")
	}

	if len(f.Results) != 0 {
		fmt.Fprintf(file, "        results: Mutable%sResults {\n", f.Type)
		fmt.Fprintf(file, "            id: OBJ_ID_RESULTS,\n")
		fmt.Fprintf(file, "        },\n")
	}

	fmt.Fprintf(file, "        state: %s%sState {\n", mutability, s.FullName)
	fmt.Fprintf(file, "            id: OBJ_ID_STATE,\n")
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

func (s *Schema) generateRustThunkAccessCheck(file *os.File, f *FuncDef) {
	grant := f.Access
	index := strings.Index(grant, "//")
	if index >= 0 {
		fmt.Fprintf(file, "    %s\n", grant[index:])
		grant = strings.TrimSpace(grant[:index])
	}
	switch grant {
	case AccessSelf:
		grant = "ctx.account_id()"
	case AccessChain:
		grant = "ctx.chain_owner_id()"
	case AccessCreator:
		grant = "ctx.contract_creator()"
	default:
		fmt.Fprintf(file, "    let access = ctx.state().get_agent_id(\"%s\");\n", grant)
		fmt.Fprintf(file, "    ctx.require(access.exists(), \"access not set: %s\");\n", grant)
		grant = "access.value()"
	}
	fmt.Fprintf(file, "    ctx.require(ctx.caller() == %s, \"no permission\");\n\n", grant)
}

func (s *Schema) generateRustTypes() error {
	if len(s.Structs) == 0 {
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
	for _, typeDef := range s.Structs {
		s.generateRustType(file, typeDef)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustType(file *os.File, typeDef *Struct) {
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

func (s *Schema) generateRustTypeProxy(file *os.File, typeDef *Struct, mutable bool) {
	typeName := PropImmutable + typeDef.Name
	if mutable {
		typeName = PropMutable + typeDef.Name
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
	s.flushConsts(func(name string, value string, padLen int) {
		fmt.Fprintf(file, "pub const %s %s;\n", pad(name+":", padLen+1), value)
	})
}
