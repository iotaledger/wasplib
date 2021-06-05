// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
	"regexp"
	"strings"
)

const generateGoThunk = true

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
	"Int64":     "wasmlib.TYPE_INT64",
	"RequestId": "wasmlib.TYPE_REQUEST_ID",
	"String":    "wasmlib.TYPE_STRING",
}

func (s *Schema) GenerateGo() error {
	s.NewTypes = make(map[string]bool)

	err := os.MkdirAll("test", 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll("wasmmain", 0755)
	if err != nil {
		return err
	}

	err = s.generateGoWasmMain()
	if err != nil {
		return err
	}
	err = s.generateGoLib()
	if err != nil {
		return err
	}
	err = s.generateGoConsts()
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
	err = s.generateGoState()
	if err != nil {
		return err
	}
	err = s.generateGoFuncs()
	if err != nil {
		return err
	}
	return s.generateGoTests()
}

func (s *Schema) generateGoConsts() error {
	file, err := os.Create("consts.go")
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
	fmt.Fprintf(file, "const HScName = wasmlib.ScHname(0x%s)\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			s.generateGoFieldConst(file, s.Params[name], "Param")
		}
	}

	if len(s.Results) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Results) {
			s.generateGoFieldConst(file, s.Results[name], "Result")
		}
	}

	if len(s.StateVars) != 0 {
		fmt.Fprintln(file)
		for _, stateVar := range s.StateVars {
			s.generateGoFieldConst(file, stateVar, "Var")
		}
	}

	if len(s.Funcs) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			fmt.Fprintf(file, "const %s = \"%s\"\n", name, funcDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.FullName)
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "const H%s = wasmlib.ScHname(0x%s)\n", name, hName.String())
		}
	}
	return nil
}

func (s *Schema) generateGoFieldConst(file *os.File, field *Field, prefix string) {
	name := prefix + capitalize(field.Name)
	fmt.Fprintf(file, "const %s = wasmlib.Key(\"%s\")\n", name, field.Alias)
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
	for _, funcDef := range s.Funcs {
		if existing[funcDef.FullName] == "" {
			s.generateGoFuncSignature(file, funcDef)
		}
	}

	return os.Remove(scOriginal)
}

func (s *Schema) generateGoFuncSignature(file *os.File, funcDef *FuncDef) {
	funcName := funcDef.FullName
	funcKind := capitalize(funcDef.FullName[:4])
	fmt.Fprintf(file, "\nfunc %s(ctx wasmlib.Sc%sContext, f *%sContext) {\n", funcName, funcKind, capitalize(funcName))
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

	for _, funcDef := range s.Funcs {
		s.generateGoFuncSignature(file, funcDef)
	}
	return nil
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

	thunk := ""
	if generateGoThunk {
		thunk = "Thunk"
	}

	fmt.Fprintf(file, "func OnLoad() {\n")
	fmt.Fprintf(file, "\texports := wasmlib.NewScExports()\n")
	for _, funcDef := range s.Funcs {
		name := capitalize(funcDef.FullName)
		kind := capitalize(funcDef.FullName[:4])
		fmt.Fprintf(file, "\texports.Add%s(%s, %s%s)\n", kind, name, funcDef.FullName, thunk)
	}
	fmt.Fprintf(file, "}\n")

	if generateGoThunk {
		// generate parameter structs and thunks to set up and check parameters
		for _, funcDef := range s.Funcs {
			s.generateGoThunk(file, funcDef)
		}
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
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, importWasmLib)

	s.generateGoStateStruct(file, "Func", "Mutable")
	s.generateGoStateStruct(file, "View", "Immutable")
	return nil
}

func (s *Schema) generateGoStateStruct(file *os.File, kind string, mutability string) {
	// first generate necessary array and map types
	for _, stateVar := range s.StateVars {
		s.generateGoProxy(file, stateVar, mutability)
	}

	x := s.FullName + kind + "State"
	fmt.Fprintf(file, "\ntype %s struct {\n", x)
	fmt.Fprintf(file, "\tstateId int32\n")
	fmt.Fprintf(file, "}\n")

	for _, stateVar := range s.StateVars {
		varName := capitalize(stateVar.Name)
		varType := goTypeIds[stateVar.Type]
		if len(varType) == 0 {
			varType = "wasmlib.TYPE_BYTES"
		}
		if stateVar.Array {
			varType = "wasmlib.TYPE_ARRAY|" + varType
			arrayType := "ArrayOf" + mutability + stateVar.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", x, varName, arrayType)
			fmt.Fprintf(file, "\tarrId := wasmlib.GetObjectId(s.stateId, Var%s.KeyId(), %s)\n", varName, varType)
			fmt.Fprintf(file, "\treturn %s{objId: arrId}\n", arrayType)
			fmt.Fprintf(file, "}\n")
			continue
		}
		if len(stateVar.MapKey) != 0 {
			varType = "wasmlib.TYPE_MAP"
			mapType := "Map" + stateVar.MapKey + "To" + mutability + stateVar.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", x, varName, mapType)
			fmt.Fprintf(file, "\tmapId := wasmlib.GetObjectId(s.stateId, Var%s.KeyId(), %s)\n", varName, varType)
			fmt.Fprintf(file, "\treturn %s{objId: mapId}\n", mapType)
			fmt.Fprintf(file, "}\n")
			continue
		}

		proxyType := mutability + stateVar.Type
		fmt.Fprintf(file, "\nfunc (s %s) %s() wasmlib.Sc%s {\n", x, varName, proxyType)
		fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(s.stateId, Var%s.KeyId())\n", proxyType, varName)
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

func (s *Schema) generateGoTests() error {
	//TODO
	return nil
}

func (s *Schema) generateGoThunk(file *os.File, funcDef *FuncDef) {
	funcName := capitalize(funcDef.FullName)
	funcKind := capitalize(funcDef.FullName[:4])
	nameLen := 5
	if len(funcDef.Params) != 0 {
		nameLen = 6
		s.generateGoThunkStruct(file, funcName, "Immutable", "Param", funcDef.Params)
	}
	if len(funcDef.Results) != 0 {
		nameLen = 7
		s.generateGoThunkStruct(file, funcName, "Mutable", "Result", funcDef.Results)
	}

	fmt.Fprintf(file, "\ntype %sContext struct {\n", funcName)
	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "\t%s %sParams\n", pad("Params", nameLen), funcName)
	}
	if len(funcDef.Results) != 0 {
		fmt.Fprintf(file, "\tResults %sResults\n", funcName)
	}
	fmt.Fprintf(file, "\t%s %s%sState\n", pad("State", nameLen), s.FullName, funcKind)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfunc %sThunk(ctx wasmlib.Sc%sContext) {\n", funcDef.FullName, funcKind)
	fmt.Fprintf(file, "\tctx.Log(\"%s.%s\")\n", s.Name, funcDef.FullName)
	grant := funcDef.Access
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

	if len(funcDef.Params) != 0 {
		fmt.Fprintf(file, "\tp := ctx.Params().MapId()\n")
	}
	if len(funcDef.Results) != 0 {
		fmt.Fprintf(file, "\tr := ctx.Results().MapId()\n")
	}

	fmt.Fprintf(file, "\tf := &%sContext{\n", funcName)

	if len(funcDef.Params) != 0 {
		s.generateGoThunkStructInit(file, funcName, "Immutable", "Param", funcDef.Params)
	}
	if len(funcDef.Results) != 0 {
		s.generateGoThunkStructInit(file, funcName, "Mutable", "Result", funcDef.Results)
	}

	fmt.Fprintf(file, "\t\tState: %s%sState{\n", s.FullName, funcKind)
	fmt.Fprintf(file, "\t\t\tstateId: wasmlib.GetObjectId(1, wasmlib.KeyState.KeyId(), wasmlib.TYPE_MAP),\n")
	fmt.Fprintf(file, "\t\t},\n")

	fmt.Fprintf(file, "\t}\n")

	for _, param := range funcDef.Params {
		if !param.Optional {
			name := capitalize(param.Name)
			fmt.Fprintf(file, "\tctx.Require(f.Params.%s.Exists(), \"missing mandatory %s\")\n", name, param.Name)
		}
	}

	fmt.Fprintf(file, "\t%s(ctx, f)\n", funcDef.FullName)
	fmt.Fprintf(file, "\tctx.Log(\"%s.%s ok\")\n", s.Name, funcDef.FullName)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoThunkStruct(file *os.File, funcName string, mutability string, kind string, fields []*Field) {
	nameLen, typeLen := calculatePadding(fields, nil, false)
	fmt.Fprintf(file, "\ntype %s%ss struct {\n", funcName, kind)
	for _, field := range fields {
		fldName := pad(capitalize(field.Name), nameLen)
		fldType := field.Type
		if field.Comment != "" {
			fldType = pad(fldType, typeLen)
		}
		fmt.Fprintf(file, "\t%s wasmlib.Sc%s%s%s\n", fldName, mutability, fldType, field.Comment)
	}
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoThunkStructInit(file *os.File, funcName string, mutability string, kind string, fields []*Field) {
	mapId := lower(kind[0:1])
	nameLen, _ := calculatePadding(fields, nil, false)
	fmt.Fprintf(file, "\t\t%ss: %s%ss{\n", kind, funcName, kind)
	for _, field := range fields {
		fldId := capitalize(field.Name)
		fldName := pad(fldId+":", nameLen+1)
		fmt.Fprintf(file, "\t\t\t%s wasmlib.NewSc%s%s(%s, %s%s.KeyId()),\n", fldName, mutability, field.Type, mapId, kind, fldId)
	}
	fmt.Fprintf(file, "\t\t},\n")
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
