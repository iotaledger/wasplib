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

const importCoreTypes = "import \"github.com/iotaledger/wasp/packages/coretypes\"\n"
const importWasmLib = "import \"github.com/iotaledger/wasplib/packages/vm/wasmlib\"\n"
const importWasmClient = "import \"github.com/iotaledger/wasplib/packages/vm/wasmclient\"\n"

var goFuncRegexp = regexp.MustCompile("^func (\\w+).+$")

var goTypes = StringMap{
	"Address":   "wasmlib.ScAddress",
	"AgentId":   "wasmlib.ScAgentId",
	"ChainId":   "wasmlib.ScChainId",
	"Color":     "wasmlib.ScColor",
	"Hash":      "wasmlib.ScHash",
	"Hname":     "wasmlib.ScHname",
	"Int16":     "int16",
	"Int32":     "int32",
	"Int64":     "int64",
	"RequestId": "wasmlib.ScRequestId",
	"String":    "string",
}

var goTypeIds = StringMap{
	"Address":   "wasmlib.TYPE_ADDRESS",
	"AgentId":   "wasmlib.TYPE_AGENT_ID",
	"ChainId":   "wasmlib.TYPE_CHAIN_ID",
	"Color":     "wasmlib.TYPE_COLOR",
	"Hash":      "wasmlib.TYPE_HASH",
	"Hname":     "wasmlib.TYPE_HNAME",
	"Int16":     "wasmlib.TYPE_INT16",
	"Int32":     "wasmlib.TYPE_INT32",
	"Int64":     "wasmlib.TYPE_INT64",
	"RequestId": "wasmlib.TYPE_REQUEST_ID",
	"String":    "wasmlib.TYPE_STRING",
}

func (s *Schema) GenerateGo() error {
	s.NewTypes = make(map[string]bool)

	err := s.generateGoConsts(false)
	if err != nil {
		return err
	}
	err = s.generateGoTypes()
	if err != nil {
		return err
	}
	err = s.generateGoSubtypes()
	if err != nil {
		return err
	}
	err = s.generateGoParams()
	if err != nil {
		return err
	}
	err = s.generateGoResults()
	if err != nil {
		return err
	}
	err = s.generateGoContract()
	if err != nil {
		return err
	}

	if !s.CoreContracts {
		err = s.generateGoKeys()
		if err != nil {
			return err
		}
		err = s.generateGoState()
		if err != nil {
			return err
		}
		err = s.generateGoLib()
		if err != nil {
			return err
		}
		err = s.generateGoFuncs()
		if err != nil {
			return err
		}

		// go-specific stuff
		return s.generateGoWasmMain()
	}

	return nil
}

func (s *Schema) generateGoConsts(test bool) error {
	file, err := os.Create("consts.go")
	if err != nil {
		return err
	}
	defer file.Close()

	packageName := "test"
	importTypes := importCoreTypes
	if !test {
		packageName = s.Name
		importTypes = importWasmLib
	}
	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", packageName)
	fmt.Fprint(file, importTypes)

	scName := s.Name
	if s.CoreContracts {
		// remove 'core' prefix
		scName = scName[4:]
	}
	s.appendConst("ScName", "\""+scName+"\"")
	if s.Description != "" {
		s.appendConst("ScDescription", "\""+s.Description+"\"")
	}
	hName := coretypes.Hn(scName)
	hNameType := "wasmlib.ScHname"
	if test {
		hNameType = "coretypes.Hname"
	}
	s.appendConst("HScName", hNameType+"(0x"+hName.String()+")")
	s.flushGoConsts(file)

	s.generateGoConstsFields(file, test, s.Params, "Param")
	s.generateGoConstsFields(file, test, s.Results, "Result")
	s.generateGoConstsFields(file, test, s.StateVars, "State")

	if len(s.Funcs) != 0 {
		for _, f := range s.Funcs {
			constName := capitalize(f.FuncName)
			s.appendConst(constName, "\""+f.String+"\"")
		}
		s.flushGoConsts(file)

		for _, f := range s.Funcs {
			constHname := "H" + capitalize(f.FuncName)
			hName = coretypes.Hn(f.String)
			s.appendConst(constHname, hNameType+"(0x"+hName.String()+")")
		}
		s.flushGoConsts(file)
	}

	return nil
}

func (s *Schema) generateGoConstsFields(file *os.File, test bool, fields []*Field, prefix string) {
	if len(fields) != 0 {
		for _, field := range fields {
			name := prefix + capitalize(field.Name)
			value := "\"" + field.Alias + "\""
			if !test {
				value = "wasmlib.Key(" + value + ")"
			}
			s.appendConst(name, value)
		}
		s.flushGoConsts(file)
	}
}

func (s *Schema) generateGoContract() error {
	file, err := os.Create("contract.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	for _, f := range s.Funcs {
		fmt.Fprintf(file, "\ntype %sCall struct {\n", f.Type)
		fmt.Fprintf(file, "\tFunc wasmlib.Sc%s\n", f.Kind)
		paramsId := "nil"
		if len(f.Params) != 0 {
			paramsId = "&f.Params.id"
			fmt.Fprintf(file, "\tParams Mutable%sParams\n", f.Type)
		}
		resultsId := "nil"
		if len(f.Results) != 0 {
			resultsId = "&f.Results.id"
			fmt.Fprintf(file, "\tResults Immutable%sResults\n", f.Type)
		}
		fmt.Fprintf(file, "}\n")

		fmt.Fprintf(file, "\nfunc New%sCall(ctx wasmlib.ScFuncContext) *%sCall {\n", f.Type, f.Type)
		fmt.Fprintf(file, "\tf := &%sCall{}\n", f.Type)
		fmt.Fprintf(file, "\tf.Func.Init(HScName, H%s%s, %s, %s)\n", f.Kind, f.Type, paramsId, resultsId)
		fmt.Fprintf(file, "\treturn f\n")
		fmt.Fprintf(file, "}\n")

		if f.Kind == "View" {
			fmt.Fprintf(file, "\nfunc New%sCallFromView(ctx wasmlib.ScViewContext) *%sCall {\n", f.Type, f.Type)
			fmt.Fprintf(file, "\tf := &%sCall{}\n", f.Type)
			fmt.Fprintf(file, "\tf.Func.Init(HScName, H%s%s, %s, %s)\n", f.Kind, f.Type, paramsId, resultsId)
			fmt.Fprintf(file, "\treturn f\n")
			fmt.Fprintf(file, "}\n")
		}
	}

	return nil
}

func (s *Schema) generateGoFuncs() error {
	scFileName := s.Name + ".go"
	file, err := os.Open(scFileName)
	if err != nil {
		// generate initial code file
		return s.generateGoFuncsNew(scFileName)
	}

	// append missing function signatures to existing code file

	lines, existing, err := s.scanExistingCode(file, goFuncRegexp)
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
			s.generateGoFuncSignature(file, f)
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) generateGoFuncSignature(file *os.File, f *FuncDef) {
	fmt.Fprintf(file, "\nfunc %s(ctx wasmlib.Sc%sContext, f *%sContext) {\n", f.FuncName, f.Kind, f.Type)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoFuncsNew(scFileName string) error {
	file, err := os.Create(scFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(false))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintln(file, importWasmLib)

	for _, f := range s.Funcs {
		s.generateGoFuncSignature(file, f)
	}
	return nil
}

func (s *Schema) generateGoKeys() error {
	file, err := os.Create("keys.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	s.KeyId = 0
	s.generateGoKeysIndexes(file, s.Params, "Param")
	s.generateGoKeysIndexes(file, s.Results, "Result")
	s.generateGoKeysIndexes(file, s.StateVars, "State")
	s.flushGoConsts(file)

	size := len(s.Params) + len(s.Results) + len(s.StateVars)
	fmt.Fprintf(file, "\nconst keyMapLen = %d\n", size)
	fmt.Fprintf(file, "\nvar keyMap = [keyMapLen]wasmlib.Key{\n")
	s.generateGoKeysArray(file, s.Params, "Param")
	s.generateGoKeysArray(file, s.Results, "Result")
	s.generateGoKeysArray(file, s.StateVars, "State")
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\nvar idxMap [keyMapLen]wasmlib.Key32\n")
	return nil
}

func (s *Schema) generateGoKeysArray(file *os.File, fields []*Field, prefix string) {
	for _, field := range fields {
		name := prefix + capitalize(field.Name)
		fmt.Fprintf(file, "\t%s,\n", name)
		s.KeyId++
	}
}

func (s *Schema) generateGoKeysIndexes(file *os.File, fields []*Field, prefix string) {
	for _, field := range fields {
		name := "Idx" + prefix + capitalize(field.Name)
		field.KeyId = s.KeyId
		value := strconv.Itoa(field.KeyId)
		s.KeyId++
		s.appendConst(name, value)
	}
}

func (s *Schema) generateGoLib() error {
	file, err := os.Create("lib.go")
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
	for _, f := range s.Funcs {
		constName := capitalize(f.FuncName)
		fmt.Fprintf(file, "\texports.Add%s(%s, %sThunk)\n", f.Kind, constName, f.FuncName)
	}

	fmt.Fprintf(file, "\n\tfor i, key := range keyMap {\n")
	fmt.Fprintf(file, "\t\tidxMap[i] = key.KeyId()\n")
	fmt.Fprintf(file, "\t}\n")

	fmt.Fprintf(file, "}\n")

	// generate parameter structs and thunks to set up and check parameters
	for _, f := range s.Funcs {
		s.generateGoThunk(file, f)
	}
	return nil
}

func (s *Schema) generateGoProxy(file *os.File, field *Field, mutability string) {
	if field.Array {
		proxyType := mutability + field.Type
		arrayType := "ArrayOf" + proxyType
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fmt.Fprintf(file, "\ntype %s%s = %s\n", mutability, field.Name, arrayType)
		}
		if s.NewTypes[arrayType] {
			// already generated this array
			return
		}
		s.NewTypes[arrayType] = true
		fmt.Fprintf(file, "\ntype %s struct {\n", arrayType)
		fmt.Fprintf(file, "\tobjId int32\n")
		fmt.Fprintf(file, "}\n")

		if mutability == "Mutable" {
			fmt.Fprintf(file, "\nfunc (a %s) Clear() {\n", arrayType)
			fmt.Fprintf(file, "\twasmlib.Clear(a.objId)\n")
			fmt.Fprintf(file, "}\n")
		}

		fmt.Fprintf(file, "\nfunc (a %s) Length() int32 {\n", arrayType)
		fmt.Fprintf(file, "\treturn wasmlib.GetLength(a.objId)\n")
		fmt.Fprintf(file, "}\n")

		if field.TypeId == 0 {
			for _, subtype := range s.Subtypes {
				if subtype.Name == field.Type {
					varType := "wasmlib.TYPE_MAP"
					if subtype.Array {
						varType = goTypeIds[subtype.Type]
						if len(varType) == 0 {
							varType = "wasmlib.TYPE_BYTES"
						}
						varType = "wasmlib.TYPE_ARRAY|" + varType
					}
					fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) %s {\n", arrayType, field.Type, proxyType)
					fmt.Fprintf(file, "\tsubId := wasmlib.GetObjectId(m.objId, wasmlib.Key32(index), %s)\n", varType)
					fmt.Fprintf(file, "\treturn %s{objId: subId}\n", proxyType)
					fmt.Fprintf(file, "}\n")
					return
				}
			}
			fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) %s {\n", arrayType, field.Type, proxyType)
			fmt.Fprintf(file, "\treturn %s{objId: a.objId, keyId: wasmlib.Key32(index)}\n", proxyType)
			fmt.Fprintf(file, "}\n")
			return
		}

		fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) wasmlib.Sc%s {\n", arrayType, field.Type, proxyType)
		fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(a.objId, wasmlib.Key32(index))\n", proxyType)
		fmt.Fprintf(file, "}\n")
		return
	}

	if len(field.MapKey) != 0 {
		proxyType := mutability + field.Type
		mapType := "Map" + field.MapKey + "To" + proxyType
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fmt.Fprintf(file, "\ntype %s%s = %s\n", mutability, field.Name, mapType)
		}
		if s.NewTypes[mapType] {
			// already generated this map
			return
		}
		s.NewTypes[mapType] = true
		fmt.Fprintf(file, "\ntype %s struct {\n", mapType)
		fmt.Fprintf(file, "\tobjId int32\n")
		fmt.Fprintf(file, "}\n")

		if mutability == "Mutable" {
			fmt.Fprintf(file, "\nfunc (m %s) Clear() {\n", mapType)
			fmt.Fprintf(file, "\twasmlib.Clear(m.objId)\n")
			fmt.Fprintf(file, "}\n")
		}

		if field.TypeId == 0 {
			for _, subtype := range s.Subtypes {
				if subtype.Name == field.Type {
					varType := "wasmlib.TYPE_MAP"
					if subtype.Array {
						varType = goTypeIds[subtype.Type]
						if len(varType) == 0 {
							varType = "wasmlib.TYPE_BYTES"
						}
						varType = "wasmlib.TYPE_ARRAY|" + varType
					}
					fmt.Fprintf(file, "\nfunc (m %s) Get%s(key wasmlib.Sc%s) %s {\n", mapType, field.Type, field.MapKey, proxyType)
					fmt.Fprintf(file, "\tsubId := wasmlib.GetObjectId(m.objId, key.KeyId(), %s)\n", varType)
					fmt.Fprintf(file, "\treturn %s{objId: subId}\n", proxyType)
					fmt.Fprintf(file, "}\n")
					return
				}
			}
			fmt.Fprintf(file, "\nfunc (m %s) Get%s(key wasmlib.Sc%s) %s {\n", mapType, field.Type, field.MapKey, proxyType)
			fmt.Fprintf(file, "\treturn %s{objId: m.objId, keyId: key.KeyId()}\n", proxyType)
			fmt.Fprintf(file, "}\n")
			return
		}

		fmt.Fprintf(file, "\nfunc (m %s) Get%s(key wasmlib.Sc%s) wasmlib.Sc%s {\n", mapType, field.Type, field.MapKey, proxyType)
		fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(m.objId, key.KeyId())\n", proxyType)
		fmt.Fprintf(file, "}\n")
	}
}

func (s *Schema) generateGoState() error {
	file, err := os.Create("state.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n", s.Name)
	if len(s.StateVars) != 0 {
		fmt.Fprintf(file, "\n"+importWasmLib)
	}

	s.generateGoStruct(file, s.StateVars, "Immutable", s.FullName, "State")
	s.generateGoStruct(file, s.StateVars, "Mutable", s.FullName, "State")
	return nil
}

func (s *Schema) generateGoParams() error {
	file, err := os.Create("params.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n", s.Name)

	params := 0
	for _, f := range s.Funcs {
		params += len(f.Params)
	}
	if params != 0 {
		fmt.Fprintf(file, "\n"+importWasmLib)
	}

	for _, f := range s.Funcs {
		if len(f.Params) == 0 {
			continue
		}
		s.generateGoStruct(file, f.Params, "Immutable", f.Type, "Params")
		s.generateGoStruct(file, f.Params, "Mutable", f.Type, "Params")
	}

	return nil
}

func (s *Schema) generateGoResults() error {
	file, err := os.Create("results.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n", s.Name)

	results := 0
	for _, f := range s.Funcs {
		results += len(f.Results)
	}
	if results != 0 {
		fmt.Fprintf(file, "\n"+importWasmLib)
	}

	for _, f := range s.Funcs {
		if len(f.Results) == 0 {
			continue
		}
		s.generateGoStruct(file, f.Results, "Immutable", f.Type, "Results")
		s.generateGoStruct(file, f.Results, "Mutable", f.Type, "Results")
	}
	return nil
}

func (s *Schema) generateGoStruct(file *os.File, fields []*Field, mutability string, typeName string, kind string) {
	typeName = mutability + typeName + kind
	if strings.HasSuffix(kind, "s") {
		kind = kind[0 : len(kind)-1]
	}

	// first generate necessary array and map types
	for _, field := range fields {
		s.generateGoProxy(file, field, mutability)
	}

	fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
	fmt.Fprintf(file, "\tid int32\n")
	fmt.Fprintf(file, "}\n")

	for _, field := range fields {
		varName := capitalize(field.Name)
		varId := "idxMap[Idx" + kind + varName + "]"
		if s.CoreContracts {
			varId = kind + varName + ".KeyId()"
		}
		varType := goTypeIds[field.Type]
		if len(varType) == 0 {
			varType = "wasmlib.TYPE_BYTES"
		}
		if field.Array {
			varType = "wasmlib.TYPE_ARRAY|" + varType
			arrayType := "ArrayOf" + mutability + field.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", typeName, varName, arrayType)
			fmt.Fprintf(file, "\tarrId := wasmlib.GetObjectId(s.id, %s, %s)\n", varId, varType)
			fmt.Fprintf(file, "\treturn %s{objId: arrId}\n", arrayType)
			fmt.Fprintf(file, "}\n")
			continue
		}
		if len(field.MapKey) != 0 {
			varType = "wasmlib.TYPE_MAP"
			mapType := "Map" + field.MapKey + "To" + mutability + field.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", typeName, varName, mapType)
			fmt.Fprintf(file, "\tmapId := wasmlib.GetObjectId(s.id, %s, %s)\n", varId, varType)
			fmt.Fprintf(file, "\treturn %s{objId: mapId}\n", mapType)
			fmt.Fprintf(file, "}\n")
			continue
		}

		proxyType := mutability + field.Type
		fmt.Fprintf(file, "\nfunc (s %s) %s() wasmlib.Sc%s {\n", typeName, varName, proxyType)
		fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(s.id, %s)\n", proxyType, varId)
		fmt.Fprintf(file, "}\n")
	}
}

func (s *Schema) generateGoSubtypes() error {
	if len(s.Subtypes) == 0 {
		return nil
	}

	file, err := os.Create("subtypes.go")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	for _, subtype := range s.Subtypes {
		s.generateGoProxy(file, subtype, "Immutable")
		s.generateGoProxy(file, subtype, "Mutable")
	}

	return nil
}

func (s *Schema) GenerateGoTests() error {
	err := os.MkdirAll("test", 0755)
	if err != nil {
		return err
	}
	err = os.Chdir("test")
	if err != nil {
		return err
	}
	defer os.Chdir("..")
	//TODO <scname>_test.go
	s.generateGoConsts(true)
	return nil
}

func (s *Schema) generateGoThunk(file *os.File, f *FuncDef) {
	nameLen := 5
	if len(f.Params) != 0 {
		nameLen = 6
	}
	if len(f.Results) != 0 {
		nameLen = 7
	}

	fmt.Fprintf(file, "\ntype %sContext struct {\n", f.Type)
	if len(f.Params) != 0 {
		fmt.Fprintf(file, "\t%s Immutable%sParams\n", pad("Params", nameLen), f.Type)
	}
	if len(f.Results) != 0 {
		fmt.Fprintf(file, "\tResults Mutable%sResults\n", f.Type)
	}
	mutability := "Mutable"
	if f.Kind == "View" {
		mutability = "Immutable"
	}
	fmt.Fprintf(file, "\t%s %s%sState\n", pad("State", nameLen), mutability, s.FullName)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfunc %sThunk(ctx wasmlib.Sc%sContext) {\n", f.FuncName, f.Kind)
	fmt.Fprintf(file, "\tctx.Log(\"%s.%s\")\n", s.Name, f.FuncName)
	grant := f.Access
	if grant != "" {
		index := strings.Index(grant, "//")
		if index >= 0 {
			fmt.Fprintf(file, "\t%s\n", grant[index:])
			grant = strings.TrimSpace(grant[:index])
		}
		switch grant {
		case "self":
			grant = "ctx.AccountId()"
		case "chain":
			grant = "ctx.ChainOwnerId()"
		case "creator":
			grant = "ctx.ContractCreator()"
		default:
			fmt.Fprintf(file, "\taccess := ctx.State().GetAgentId(wasmlib.Key(\"%s\"))\n", grant)
			fmt.Fprintf(file, "\tctx.Require(access.Exists(), \"access not set: %s\")\n", grant)
			grant = fmt.Sprintf("access.Value()")
		}
		fmt.Fprintf(file, "\tctx.Require(ctx.Caller() == %s, \"no permission\")\n\n", grant)
	}

	fmt.Fprintf(file, "\tf := &%sContext{\n", f.Type)

	if len(f.Params) != 0 {
		fmt.Fprintf(file, "\t\tParams: Immutable%sParams{\n", f.Type)
		fmt.Fprintf(file, "\t\t\tid: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),\n")
		fmt.Fprintf(file, "\t\t},\n")
	}

	if len(f.Results) != 0 {
		fmt.Fprintf(file, "\t\tResults: Mutable%sResults{\n", f.Type)
		fmt.Fprintf(file, "\t\t\tid: wasmlib.GetObjectId(1, wasmlib.KeyResults, wasmlib.TYPE_MAP),\n")
		fmt.Fprintf(file, "\t\t},\n")
	}

	fmt.Fprintf(file, "\t\tState: %s%sState{\n", mutability, s.FullName)
	fmt.Fprintf(file, "\t\t\tid: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),\n")
	fmt.Fprintf(file, "\t\t},\n")

	fmt.Fprintf(file, "\t}\n")

	for _, param := range f.Params {
		if !param.Optional {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "\tctx.Require(f.Params.%s().Exists(), \"missing mandatory %s\")\n", name, param.Name)
		}
	}

	fmt.Fprintf(file, "\t%s(ctx, f)\n", f.FuncName)
	fmt.Fprintf(file, "\tctx.Log(\"%s.%s ok\")\n", s.Name, f.FuncName)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	file, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	for _, typeDef := range s.Types {
		s.generateGoType(file, typeDef)
	}

	return nil
}

func (s *Schema) generateGoType(file *os.File, typeDef *TypeDef) {
	nameLen, typeLen := calculatePadding(typeDef.Fields, goTypes, false)

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
	fmt.Fprintf(file, "\tdecode := wasmlib.NewBytesDecoder(bytes)\n")
	fmt.Fprintf(file, "\tdata := &%s{}\n", typeDef.Name)
	for _, field := range typeDef.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "\tdata.%s = decode.%s()\n", name, field.Type)
	}
	fmt.Fprintf(file, "\tdecode.Close()\n")
	fmt.Fprintf(file, "\treturn data\n}\n")

	fmt.Fprintf(file, "\nfunc (o *%s) Bytes() []byte {\n", typeDef.Name)
	fmt.Fprintf(file, "\treturn wasmlib.NewBytesEncoder().\n")
	for _, field := range typeDef.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "\t\t%s(o.%s).\n", field.Type, name)
	}
	fmt.Fprintf(file, "\t\tData()\n}\n")

	s.generateGoTypeProxy(file, typeDef, false)
	s.generateGoTypeProxy(file, typeDef, true)
}

func (s *Schema) generateGoTypeProxy(file *os.File, typeDef *TypeDef, mutable bool) {
	typeName := "Immutable" + typeDef.Name
	if mutable {
		typeName = "Mutable" + typeDef.Name
	}

	fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
	fmt.Fprintf(file, "\tobjId int32\n")
	fmt.Fprintf(file, "\tkeyId wasmlib.Key32\n")
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfunc (o %s) Exists() bool {\n", typeName)
	fmt.Fprintf(file, "\treturn wasmlib.Exists(o.objId, o.keyId, wasmlib.TYPE_BYTES)\n")
	fmt.Fprintf(file, "}\n")

	if mutable {
		fmt.Fprintf(file, "\nfunc (o %s) SetValue(value *%s) {\n", typeName, typeDef.Name)
		fmt.Fprintf(file, "\twasmlib.SetBytes(o.objId, o.keyId, wasmlib.TYPE_BYTES, value.Bytes())\n")
		fmt.Fprintf(file, "}\n")
	}

	fmt.Fprintf(file, "\nfunc (o %s) Value() *%s {\n", typeName, typeDef.Name)
	fmt.Fprintf(file, "\treturn New%sFromBytes(wasmlib.GetBytes(o.objId, o.keyId, wasmlib.TYPE_BYTES))\n", typeDef.Name)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoWasmMain() error {
	err := os.MkdirAll("wasmmain", 0755)
	if err != nil {
		return err
	}

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

func (s *Schema) flushGoConsts(file *os.File) {
	if len(s.ConstNames) == 0 {
		return
	}

	if len(s.ConstNames) == 1 {
		name := s.ConstNames[0]
		value := s.ConstValues[0]
		fmt.Fprintf(file, "\nconst %s = %s\n", name, value)
		s.flushConsts(file, func(name string, value string, padLen int) {})
		return
	}

	fmt.Fprintf(file, "\nconst (\n")
	s.flushConsts(file, func(name string, value string, padLen int) {
		fmt.Fprintf(file, "\t%s = %s\n", pad(name, padLen), value)
	})
	fmt.Fprintf(file, ")\n")
}
