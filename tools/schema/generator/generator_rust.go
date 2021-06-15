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

const allowDeadCode = "#![allow(dead_code)]\n"
const allowUnusedImports = "#![allow(unused_imports)]\n"
const useConsts = "use crate::consts::*;\n"
const useCrate = "use crate::*;\n"
const useKeys = "use crate::keys::*;\n"
const useParams = "use crate::params::*;\n"
const useResults = "use crate::results::*;\n"
const useState = "use crate::state::*;\n"
const useSubtypes = "use crate::subtypes::*;\n"
const useTypes = "use crate::types::*;\n"
const useWasmLib = "use wasmlib::*;\n"
const useWasmLibHost = "use wasmlib::host::*;\n"

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
	fmt.Fprintln(file, s.crateOrWasmLib(false, false))

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
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := upper(snake(funcDef.FullName))
			s.appendConst(name, "&str = \""+funcDef.Name+"\"")
		}
		s.flushRustConsts(file)

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := "H" + upper(snake(funcDef.FullName))
			hName = coretypes.Hn(funcDef.Name)
			s.appendConst(name, "ScHname = ScHname(0x"+hName.String()+")")
		}
		s.flushRustConsts(file)
	}

	formatter(file, true)
	return nil
}

func (s *Schema) generateRustConstsFields(file *os.File, fields []*Field, prefix string) {
	if len(fields) != 0 {
		fmt.Fprintln(file)
		for _, field := range fields {
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
	fmt.Fprint(file, s.crateOrWasmLib(true, false))
	if !s.CoreContracts {
		fmt.Fprint(file, "\n"+useConsts)
		fmt.Fprint(file, useParams)
		fmt.Fprint(file, useResults)
	}

	typeName := s.FullName + "Func"
	fmt.Fprintf(file, "\npub struct %s {\n", typeName)
	fmt.Fprintf(file, "    sc: ScContractFunc,\n")
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\nimpl %s {", typeName)

	fmt.Fprintf(file, "\n    pub fn new(ctx: &ScFuncContext) -> %s {\n", typeName)
	fmt.Fprintf(file, "        %s { sc: ScContractFunc::new(ctx, HSC_NAME) }\n", typeName)
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn delay(&mut self, seconds: i32) -> &mut %s {\n", typeName)
	fmt.Fprintf(file, "        self.sc.delay(seconds);\n")
	fmt.Fprintf(file, "        self\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn of_contract(&mut self, contract: ScHname) -> &mut %s {\n", typeName)
	fmt.Fprintf(file, "        self.sc.of_contract(contract);\n")
	fmt.Fprintf(file, "        self\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn post(&mut self) -> &mut %s {\n", typeName)
	fmt.Fprintf(file, "        self.sc.post();\n")
	fmt.Fprintf(file, "        self\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn post_to_chain(&mut self, chain_id: ScChainId) -> &mut %s {\n", typeName)
	fmt.Fprintf(file, "        self.sc.post_to_chain(chain_id);\n")
	fmt.Fprintf(file, "        self\n")
	fmt.Fprintf(file, "    }\n")

	for _, funcDef := range s.Funcs {
		funcName := capitalize(funcDef.FullName)
		shortName := snake(funcName[4:])
		params := ""
		params2 := "0"
		if len(funcDef.Params) != 0 {
			params = ", params: Mutable" + funcName + "Params"
			params2 = "params.id"
		}
		results := ""
		if len(funcDef.Results) != 0 {
			results = "-> Immutable" + funcName + "Results "
		}

		if strings.HasPrefix(funcName, "Func") {
			fmt.Fprintf(file, "\n    pub fn %s(&mut self%s, transfer: ScTransfers) %s{\n", shortName, params, results)
			fmt.Fprintf(file, "        self.sc.run(H%s, %s, Some(transfer));\n", upper(snake(funcName)), params2)
			if len(funcDef.Results) != 0 {
				fmt.Fprintf(file, "        %s{ id: self.sc.result_map_id() }\n", results[3:])
			}
			fmt.Fprintf(file, "    }\n")
			continue
		}

		fmt.Fprintf(file, "\n    pub fn %s(&mut self%s) %s{\n", shortName, params, results)
		fmt.Fprintf(file, "        self.sc.run(H%s, %s, None);\n", upper(snake(funcName)), params2)
		if len(funcDef.Results) != 0 {
			fmt.Fprintf(file, "        %s{ id: self.sc.result_map_id() }\n", results[3:])
		}
		fmt.Fprintf(file, "    }\n")
	}
	fmt.Fprintf(file, "}\n")

	typeName = s.FullName + "View"
	fmt.Fprintf(file, "\npub struct %s {\n", typeName)
	fmt.Fprintf(file, "    sc: ScContractView,\n")
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\nimpl %s {", typeName)

	fmt.Fprintf(file, "\n    pub fn new(ctx: &ScViewContext) -> %s {\n", typeName)
	fmt.Fprintf(file, "        %s { sc: ScContractView::new(ctx, HSC_NAME) }\n", typeName)
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "\n    pub fn of_contract(&mut self, contract: ScHname) -> &mut %s {\n", typeName)
	fmt.Fprintf(file, "        self.sc.of_contract(contract);\n")
	fmt.Fprintf(file, "        self\n")
	fmt.Fprintf(file, "    }\n")

	for _, funcDef := range s.Funcs {
		funcName := capitalize(funcDef.FullName)
		shortName := snake(funcName[4:])
		if strings.HasPrefix(funcName, "Func") {
			continue
		}

		params := ""
		params2 := "0"
		if len(funcDef.Params) != 0 {
			params = ", params: Mutable" + funcName + "Params"
			params2 = "params.id"
		}
		results := ""
		if len(funcDef.Results) != 0 {
			results = "-> Immutable" + funcName + "Results "
		}

		fmt.Fprintf(file, "\n    pub fn %s(&mut self%s) %s{\n", shortName, params, results)
		fmt.Fprintf(file, "        self.sc.run(H%s, %s);\n", upper(snake(funcName)), params2)
		if len(funcDef.Results) != 0 {
			fmt.Fprintf(file, "        %s{ id: self.sc.result_map_id() }\n", results[3:])
		}
		fmt.Fprintf(file, "    }\n")
	}
	fmt.Fprintf(file, "}\n")

	return nil
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
	for _, funcDef := range s.Funcs {
		name := snake(funcDef.FullName)
		if existing[name] == "" {
			s.generateRustFuncSignature(file, funcDef)
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) generateRustFuncSignature(file *os.File, funcDef *FuncDef) {
	funcName := snake(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\npub fn %s(_ctx: &Sc%sContext, _f: &%sContext) {\n", funcName, funcKind, capitalize(funcDef.FullName))
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

	for _, funcDef := range s.Funcs {
		s.generateRustFuncSignature(file, funcDef)
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
	fmt.Fprintln(file, useCrate)

	s.KeyId = 0
	s.generateRustKeysIndexes(file, s.Params, "PARAM_")
	s.generateRustKeysIndexes(file, s.Results, "RESULT_")
	s.generateRustKeysIndexes(file, s.StateVars, "STATE_")
	s.flushRustConsts(file)

	size := len(s.Params) + len(s.Results) + len(s.StateVars)
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
		name := prefix + upper(snake(field.Name))
		fmt.Fprintf(file, "    %s,\n", name)
		s.KeyId++
	}
}

func (s *Schema) generateRustKeysIndexes(file *os.File, fields []*Field, prefix string) {
	for _, field := range fields {
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
	for _, funcDef := range s.Funcs {
		name := snake(funcDef.FullName)
		kind := funcDef.FullName[:4]
		fmt.Fprintf(file, "    exports.add_%s(%s, %s_thunk);\n", kind, upper(name), name)
	}

	fmt.Fprintf(file, "\n    unsafe {\n")
	fmt.Fprintf(file, "        for i in 0..KEY_MAP_LEN {\n")
	fmt.Fprintf(file, "            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);\n")
	fmt.Fprintf(file, "        }\n")
	fmt.Fprintf(file, "    }\n")

	fmt.Fprintf(file, "}\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, funcDef := range s.Funcs {
		s.generateRustThunk(file, funcDef)
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
					fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: &Sc%s) -> %s {\n", snake(field.Type), field.MapKey, proxyType)
					fmt.Fprintf(file, "        let sub_id = get_object_id(self.obj_id, key.get_key_id(), %s);\n", varType)
					fmt.Fprintf(file, "        %s { obj_id: sub_id }\n", proxyType)
					fmt.Fprintf(file, "    }\n")
					return
				}
			}
			fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: &Sc%s) -> %s {\n", snake(field.Type), field.MapKey, proxyType)
			fmt.Fprintf(file, "        %s { obj_id: self.obj_id, key_id: key.get_key_id() }\n", proxyType)
			fmt.Fprintf(file, "    }\n")
			return
		}

		fmt.Fprintf(file, "\n    pub fn get_%s(&self, key: &Sc%s) -> Sc%s {\n", snake(field.Type), field.MapKey, proxyType)
		fmt.Fprintf(file, "        Sc%s::new(self.obj_id, key.get_key_id())\n", proxyType)
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

	s.generateRustStruct(file, s.StateVars, "Immutable", s.FullName, "State", false)
	s.generateRustStruct(file, s.StateVars, "Mutable", s.FullName, "State", false)
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
	fmt.Fprint(file, s.crateOrWasmLib(true, false))
	if !s.CoreContracts {
		fmt.Fprint(file, useWasmLibHost)
		fmt.Fprint(file, "\n"+useCrate)
		fmt.Fprint(file, useKeys)
	}

	for _, f := range s.Funcs {
		typeName := capitalize(f.FullName)
		s.generateRustStruct(file, f.Params, "Immutable", typeName, "Params", false)
		s.generateRustStruct(file, f.Params, "Mutable", typeName, "Params", true)
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
		fmt.Fprint(file, useWasmLibHost)
		fmt.Fprint(file, "\n"+useCrate)
		fmt.Fprint(file, useKeys)
	}

	for _, f := range s.Funcs {
		typeName := capitalize(f.FullName)
		s.generateRustStruct(file, f.Results, "Immutable", typeName, "Results", false)
		s.generateRustStruct(file, f.Results, "Mutable", typeName, "Results", false)
	}
	return nil
}

func (s *Schema) generateRustStruct(file *os.File, fields []*Field, mutability string, typeName string, kind string, new bool) {
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

	if len(fields) != 0 || new {
		fmt.Fprintf(file, "\nimpl %s {", typeName)
		defer fmt.Fprintf(file, "}\n")
	}

	if new {
		fmt.Fprintf(file, "\n    pub fn new() -> %s {\n", typeName)
		fmt.Fprintf(file, "        %s { id: ScMutableMap::new().map_id() }\n", typeName)
		fmt.Fprintf(file, "    }\n")
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
			fmt.Fprintf(file, "        let map_id = get_object_id(self.id, %s, %s);\n", varId, varType)
			fmt.Fprintf(file, "        %s { obj_id: map_id }\n", mapType)
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

func (s *Schema) generateRustThunk(file *os.File, funcDef *FuncDef) {
	funcName := capitalize(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])
	nameLen := 5
	if len(funcDef.Params) != 0 {
		nameLen = 6
	}
	if len(funcDef.Results) != 0 {
		nameLen = 7
	}

	fmt.Fprintf(file, "\npub struct %sContext {\n", funcName)
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "    %s Immutable%sParams,\n", pad("params:", nameLen+1), funcName)
	}
	if len(funcDef.Results) != 0 {
		fmt.Fprintf(file, "    results: Mutable%sResults,\n", funcName)
	}
	mutability := "Mutable"
	if funcKind == "View" {
		mutability = "Immutable"
	}
	fmt.Fprintf(file, "    %s %s%sState,\n", pad("state:", nameLen+1), mutability, s.FullName)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfn %s_thunk(ctx: &Sc%sContext) {\n", snake(funcDef.FullName), funcKind)
	fmt.Fprintf(file, "    ctx.log(\"%s.%s\");\n", s.Name, funcDef.FullName)
	grant := funcDef.Access
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

	fmt.Fprintf(file, "    let f = %sContext {\n", funcName)

	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "        params: Immutable%sParams {\n", funcName)
		fmt.Fprintf(file, "            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),\n")
		fmt.Fprintf(file, "        },\n")
	}
	if len(funcDef.Results) != 0 {
		fmt.Fprintf(file, "        results: Mutable%sResults {\n", funcName)
		fmt.Fprintf(file, "            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),\n")
		fmt.Fprintf(file, "        },\n")
	}

	fmt.Fprintf(file, "        state: %s%sState {\n", mutability, s.FullName)
	fmt.Fprintf(file, "            id: get_object_id(1, KEY_STATE, TYPE_MAP),\n")
	fmt.Fprintf(file, "        },\n")

	fmt.Fprintf(file, "    };\n")

	for _, param := range funcDef.Params {
		if !param.Optional {
			name := snake(param.Name)
			fmt.Fprintf(file, "    ctx.require(f.params.%s().exists(), \"missing mandatory %s\");\n", name, param.Name)
		}
	}

	fmt.Fprintf(file, "    %s(ctx, &f);\n", snake(funcDef.FullName))
	fmt.Fprintf(file, "    ctx.log(\"%s.%s ok\");\n", s.Name, funcDef.FullName)
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
	s.flushConsts(file, func(name string, value string, padLen int) {
		fmt.Fprintf(file, "pub const %s %s;\n", pad(name+":", padLen+1), value)
	})
}
