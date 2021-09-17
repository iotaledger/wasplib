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

const importCoreTypes = `import "github.com/iotaledger/wasp/packages/iscp"
`

const importWasmLib = `import "github.com/iotaledger/wasplib/packages/vm/wasmlib"
`

const importWasmClient = `import "github.com/iotaledger/wasplib/packages/vm/wasmclient"
`

var goFuncRegexp = regexp.MustCompile(`^func (\w+).+$`)

var goTypes = StringMap{
	"Address":   "wasmlib.ScAddress",
	"AgentID":   "wasmlib.ScAgentID",
	"ChainID":   "wasmlib.ScChainID",
	"Color":     "wasmlib.ScColor",
	"Hash":      "wasmlib.ScHash",
	"Hname":     "wasmlib.ScHname",
	"Int16":     "int16",
	"Int32":     "int32",
	"Int64":     "int64",
	"RequestID": "wasmlib.ScRequestID",
	"String":    "string",
}

var goKeys = StringMap{
	"Address":   "key",
	"AgentID":   "key",
	"ChainID":   "key",
	"Color":     "key",
	"Hash":      "key",
	"Hname":     "key",
	"Int16":     "??TODO",
	"Int32":     "wasmlib.Key32(int32)",
	"Int64":     "??TODO",
	"RequestID": "key",
	"String":    "wasmlib.Key(key)",
}

var goTypeIds = StringMap{
	"Address":   "wasmlib.TYPE_ADDRESS",
	"AgentID":   "wasmlib.TYPE_AGENT_ID",
	"ChainID":   "wasmlib.TYPE_CHAIN_ID",
	"Color":     "wasmlib.TYPE_COLOR",
	"Hash":      "wasmlib.TYPE_HASH",
	"Hname":     "wasmlib.TYPE_HNAME",
	"Int16":     "wasmlib.TYPE_INT16",
	"Int32":     "wasmlib.TYPE_INT32",
	"Int64":     "wasmlib.TYPE_INT64",
	"RequestID": "wasmlib.TYPE_REQUEST_ID",
	"String":    "wasmlib.TYPE_STRING",
}

const (
	goTypeBytes = "wasmlib.TYPE_BYTES"
	goTypeMap   = "wasmlib.TYPE_MAP"
)

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
	hName := iscp.Hn(scName)
	hNameType := "wasmlib.ScHname"
	if test {
		hNameType = "iscp.Hname"
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
			hName = iscp.Hn(f.String)
			s.appendConst(constHname, hNameType+"(0x"+hName.String()+")")
		}
		s.flushGoConsts(file)
	}

	return nil
}

func (s *Schema) generateGoConstsFields(file *os.File, test bool, fields []*Field, prefix string) {
	if len(fields) != 0 {
		for _, field := range fields {
			if field.Alias == AliasThis {
				continue
			}
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
	fmt.Fprint(file, importWasmLib)

	for _, f := range s.Funcs {
		nameLen := f.nameLen(4)
		kind := f.Kind
		if f.Type == InitFunc {
			kind = f.Type + f.Kind
		}
		fmt.Fprintf(file, "\ntype %sCall struct {\n", f.Type)
		fmt.Fprintf(file, "\t%s *wasmlib.Sc%s\n", pad(KindFunc, nameLen), kind)
		if len(f.Params) != 0 {
			fmt.Fprintf(file, "\t%s Mutable%sParams\n", pad("Params", nameLen), f.Type)
		}
		if len(f.Results) != 0 {
			fmt.Fprintf(file, "\tResults Immutable%sResults\n", f.Type)
		}
		fmt.Fprintf(file, "}\n")
	}

	s.generateGoContractFuncs(file)
	return nil
}

func (s *Schema) generateGoContractFuncs(file *os.File) {
	fmt.Fprint(file, "\ntype Funcs struct{}\n")
	fmt.Fprint(file, "\nvar ScFuncs Funcs\n")
	for _, f := range s.Funcs {
		assign := "return"
		paramsID := "nil"
		if len(f.Params) != 0 {
			assign = "f :="
			paramsID = "&f.Params.id"
		}
		resultsID := "nil"
		if len(f.Results) != 0 {
			assign = "f :="
			resultsID = "&f.Results.id"
		}
		kind := f.Kind
		keyMap := ""
		if f.Type == InitFunc {
			kind = f.Type + f.Kind
			keyMap = ", ctx, keyMap[:], idxMap[:]"
		}
		fmt.Fprintf(file, "\nfunc (sc Funcs) %s(ctx wasmlib.Sc%sCallContext) *%sCall {\n", f.Type, f.Kind, f.Type)
		fmt.Fprintf(file, "\t%s &%sCall{Func: wasmlib.NewSc%s(HScName, H%s%s%s)}\n", assign, f.Type, kind, f.Kind, f.Type, keyMap)
		if len(f.Params) != 0 || len(f.Results) != 0 {
			fmt.Fprintf(file, "\tf.Func.SetPtrs(%s, %s)\n", paramsID, resultsID)
			fmt.Fprintf(file, "\treturn f\n")
		}
		fmt.Fprintf(file, "}\n")
	}
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
	switch f.FuncName {
	case "funcInit":
		fmt.Fprintf(file, "    if f.Params.Owner().Exists() {\n")
		fmt.Fprintf(file, "        f.State.Owner().SetValue(f.Params.Owner().Value())\n")
		fmt.Fprintf(file, "        return\n")
		fmt.Fprintf(file, "    }\n")
		fmt.Fprintf(file, "    f.State.Owner().SetValue(ctx.ContractCreator())\n")
	case "funcSetOwner":
		fmt.Fprintf(file, "    f.State.Owner().SetValue(f.Params.Owner().Value())\n")
	case "viewGetOwner":
		fmt.Fprintf(file, "    f.Results.Owner().SetValue(f.State.Owner().Value())\n")
	default:
	}
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
	fmt.Fprint(file, importWasmLib)

	s.KeyID = 0
	s.generateGoKeysIndexes(s.Params, "Param")
	s.generateGoKeysIndexes(s.Results, "Result")
	s.generateGoKeysIndexes(s.StateVars, "State")
	s.flushGoConsts(file)

	size := s.KeyID
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
		if field.Alias == AliasThis {
			continue
		}
		name := prefix + capitalize(field.Name)
		fmt.Fprintf(file, "\t%s,\n", name)
		s.KeyID++
	}
}

func (s *Schema) generateGoKeysIndexes(fields []*Field, prefix string) {
	for _, field := range fields {
		if field.Alias == AliasThis {
			continue
		}
		name := "Idx" + prefix + capitalize(field.Name)
		field.KeyID = s.KeyID
		value := strconv.Itoa(field.KeyID)
		s.KeyID++
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
	fmt.Fprintln(file, "//nolint:dupl")
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintln(file, importWasmLib)

	fmt.Fprintf(file, "func OnLoad() {\n")
	fmt.Fprintf(file, "\texports := wasmlib.NewScExports()\n")
	for _, f := range s.Funcs {
		constName := capitalize(f.FuncName)
		fmt.Fprintf(file, "\texports.Add%s(%s, %sThunk)\n", f.Kind, constName, f.FuncName)
	}

	fmt.Fprintf(file, "\n\tfor i, key := range keyMap {\n")
	fmt.Fprintf(file, "\t\tidxMap[i] = key.KeyID()\n")
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
		s.generateGoProxyArray(file, field, mutability)
		return
	}

	if field.MapKey != "" {
		s.generateGoProxyMap(file, field, mutability)
	}
}

func (s *Schema) generateGoProxyArray(file *os.File, field *Field, mutability string) {
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
	fmt.Fprintf(file, "\tobjID int32\n")
	fmt.Fprintf(file, "}\n")

	if mutability == PropMutable {
		fmt.Fprintf(file, "\nfunc (a %s) Clear() {\n", arrayType)
		fmt.Fprintf(file, "\twasmlib.Clear(a.objID)\n")
		fmt.Fprintf(file, "}\n")
	}

	fmt.Fprintf(file, "\nfunc (a %s) Length() int32 {\n", arrayType)
	fmt.Fprintf(file, "\treturn wasmlib.GetLength(a.objID)\n")
	fmt.Fprintf(file, "}\n")

	if field.TypeID == 0 {
		s.generateGoProxyArrayNewType(file, field, proxyType, arrayType)
		return
	}

	// array of predefined type
	fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) wasmlib.Sc%s {\n", arrayType, field.Type, proxyType)
	fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(a.objID, wasmlib.Key32(index))\n", proxyType)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoProxyArrayNewType(file *os.File, field *Field, proxyType, arrayType string) {
	for _, subtype := range s.Typedefs {
		if subtype.Name != field.Type {
			continue
		}
		varType := goTypeMap
		if subtype.Array {
			varType = goTypeIds[subtype.Type]
			if varType == "" {
				varType = goTypeBytes
			}
			varType = "wasmlib.TYPE_ARRAY|" + varType
		}
		fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) %s {\n", arrayType, field.Type, proxyType)
		fmt.Fprintf(file, "\tsubID := wasmlib.GetObjectID(a.objID, wasmlib.Key32(index), %s)\n", varType)
		fmt.Fprintf(file, "\treturn %s{objID: subID}\n", proxyType)
		fmt.Fprintf(file, "}\n")
		return
	}

	fmt.Fprintf(file, "\nfunc (a %s) Get%s(index int32) %s {\n", arrayType, field.Type, proxyType)
	fmt.Fprintf(file, "\treturn %s{objID: a.objID, keyID: wasmlib.Key32(index)}\n", proxyType)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoProxyMap(file *os.File, field *Field, mutability string) {
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

	keyType := goTypes[field.MapKey]
	keyValue := goKeys[field.MapKey]

	fmt.Fprintf(file, "\ntype %s struct {\n", mapType)
	fmt.Fprintf(file, "\tobjID int32\n")
	fmt.Fprintf(file, "}\n")

	if mutability == PropMutable {
		fmt.Fprintf(file, "\nfunc (m %s) Clear() {\n", mapType)
		fmt.Fprintf(file, "\twasmlib.Clear(m.objID)\n")
		fmt.Fprintf(file, "}\n")
	}

	if field.TypeID == 0 {
		s.generateGoProxyMapNewType(file, field, proxyType, mapType, keyType, keyValue)
		return
	}

	// map of predefined type
	fmt.Fprintf(file, "\nfunc (m %s) Get%s(key %s) wasmlib.Sc%s {\n", mapType, field.Type, keyType, proxyType)
	fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(m.objID, %s.KeyID())\n", proxyType, keyValue)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoProxyMapNewType(file *os.File, field *Field, proxyType, mapType, keyType, keyValue string) {
	for _, subtype := range s.Typedefs {
		if subtype.Name != field.Type {
			continue
		}
		varType := goTypeMap
		if subtype.Array {
			varType = goTypeIds[subtype.Type]
			if varType == "" {
				varType = goTypeBytes
			}
			varType = "wasmlib.TYPE_ARRAY|" + varType
		}
		fmt.Fprintf(file, "\nfunc (m %s) Get%s(key %s) %s {\n", mapType, field.Type, keyType, proxyType)
		fmt.Fprintf(file, "\tsubID := wasmlib.GetObjectID(m.objID, %s.KeyID(), %s)\n", keyValue, varType)
		fmt.Fprintf(file, "\treturn %s{objID: subID}\n", proxyType)
		fmt.Fprintf(file, "}\n")
		return
	}

	fmt.Fprintf(file, "\nfunc (m %s) Get%s(key %s) %s {\n", mapType, field.Type, keyType, proxyType)
	fmt.Fprintf(file, "\treturn %s{objID: m.objID, keyID: %s.KeyID()}\n", proxyType, keyValue)
	fmt.Fprintf(file, "}\n")
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

	s.generateGoStruct(file, s.StateVars, PropImmutable, s.FullName, "State")
	s.generateGoStruct(file, s.StateVars, PropMutable, s.FullName, "State")
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
		s.generateGoStruct(file, f.Params, PropImmutable, f.Type, "Params")
		s.generateGoStruct(file, f.Params, PropMutable, f.Type, "Params")
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
		s.generateGoStruct(file, f.Results, PropImmutable, f.Type, "Results")
		s.generateGoStruct(file, f.Results, PropMutable, f.Type, "Results")
	}
	return nil
}

func (s *Schema) generateGoStruct(file *os.File, fields []*Field, mutability, typeName, kind string) {
	typeName = mutability + typeName + kind
	kind = strings.TrimSuffix(kind, "s")

	// first generate necessary array and map types
	for _, field := range fields {
		s.generateGoProxy(file, field, mutability)
	}

	fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
	fmt.Fprintf(file, "\tid int32\n")
	fmt.Fprintf(file, "}\n")

	for _, field := range fields {
		varName := capitalize(field.Name)
		varID := "idxMap[Idx" + kind + varName + "]"
		if s.CoreContracts {
			varID = kind + varName + ".KeyID()"
		}
		varType := goTypeIds[field.Type]
		if varType == "" {
			varType = goTypeBytes
		}
		if field.Array {
			varType = "wasmlib.TYPE_ARRAY|" + varType
			arrayType := "ArrayOf" + mutability + field.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", typeName, varName, arrayType)
			fmt.Fprintf(file, "\tarrID := wasmlib.GetObjectID(s.id, %s, %s)\n", varID, varType)
			fmt.Fprintf(file, "\treturn %s{objID: arrID}\n", arrayType)
			fmt.Fprintf(file, "}\n")
			continue
		}
		if field.MapKey != "" {
			varType = goTypeMap
			mapType := "Map" + field.MapKey + "To" + mutability + field.Type
			fmt.Fprintf(file, "\nfunc (s %s) %s() %s {\n", typeName, varName, mapType)
			mapID := "s.id"
			if field.Alias != AliasThis {
				mapID = "mapID"
				fmt.Fprintf(file, "\tmapID := wasmlib.GetObjectID(s.id, %s, %s)\n", varID, varType)
			}
			fmt.Fprintf(file, "\treturn %s{objID: %s}\n", mapType, mapID)
			fmt.Fprintf(file, "}\n")
			continue
		}

		proxyType := mutability + field.Type
		fmt.Fprintf(file, "\nfunc (s %s) %s() wasmlib.Sc%s {\n", typeName, varName, proxyType)
		fmt.Fprintf(file, "\treturn wasmlib.NewSc%s(s.id, %s)\n", proxyType, varID)
		fmt.Fprintf(file, "}\n")
	}
}

func (s *Schema) generateGoSubtypes() error {
	if len(s.Typedefs) == 0 {
		return nil
	}

	file, err := os.Create("typedefs.go")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprint(file, importWasmLib)

	for _, subtype := range s.Typedefs {
		s.generateGoProxy(file, subtype, PropImmutable)
		s.generateGoProxy(file, subtype, PropMutable)
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
	defer func() {
		_ = os.Chdir("..")
	}()
	// TODO <scname>_test.go
	return nil // s.generateGoConsts(true)
}

func (s *Schema) generateGoThunk(file *os.File, f *FuncDef) {
	nameLen := f.nameLen(5)
	fmt.Fprintf(file, "\ntype %sContext struct {\n", f.Type)
	if len(f.Params) != 0 {
		fmt.Fprintf(file, "\t%s Immutable%sParams\n", pad("Params", nameLen), f.Type)
	}
	if len(f.Results) != 0 {
		fmt.Fprintf(file, "\tResults Mutable%sResults\n", f.Type)
	}
	mutability := PropMutable
	if f.Kind == KindView {
		mutability = PropImmutable
	}
	fmt.Fprintf(file, "\t%s %s%sState\n", pad("State", nameLen), mutability, s.FullName)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfunc %sThunk(ctx wasmlib.Sc%sContext) {\n", f.FuncName, f.Kind)
	fmt.Fprintf(file, "\tctx.Log(\"%s.%s\")\n", s.Name, f.FuncName)

	if f.Access != "" {
		s.generateGoThunkAccessCheck(file, f)
	}

	fmt.Fprintf(file, "\tf := &%sContext{\n", f.Type)

	if len(f.Params) != 0 {
		fmt.Fprintf(file, "\t\tParams: Immutable%sParams{\n", f.Type)
		fmt.Fprintf(file, "\t\t\tid: wasmlib.OBJ_ID_PARAMS,\n")
		fmt.Fprintf(file, "\t\t},\n")
	}

	if len(f.Results) != 0 {
		fmt.Fprintf(file, "\t\tResults: Mutable%sResults{\n", f.Type)
		fmt.Fprintf(file, "\t\t\tid: wasmlib.OBJ_ID_RESULTS,\n")
		fmt.Fprintf(file, "\t\t},\n")
	}

	fmt.Fprintf(file, "\t\tState: %s%sState{\n", mutability, s.FullName)
	fmt.Fprintf(file, "\t\t\tid: wasmlib.OBJ_ID_STATE,\n")
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

func (s *Schema) generateGoThunkAccessCheck(file *os.File, f *FuncDef) {
	grant := f.Access
	index := strings.Index(grant, "//")
	if index >= 0 {
		fmt.Fprintf(file, "\t%s\n", grant[index:])
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
		fmt.Fprintf(file, "\taccess := ctx.State().GetAgentID(wasmlib.Key(\"%s\"))\n", grant)
		fmt.Fprintf(file, "\tctx.Require(access.Exists(), \"access not set: %s\")\n", grant)
		grant = "access.Value()"
	}
	fmt.Fprintf(file, "\tctx.Require(ctx.Caller() == %s, \"no permission\")\n\n", grant)
}

func (s *Schema) generateGoTypes() error {
	if len(s.Structs) == 0 {
		return nil
	}

	file, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprint(file, importWasmLib)

	for _, typeDef := range s.Structs {
		s.generateGoType(file, typeDef)
	}

	return nil
}

func (s *Schema) generateGoType(file *os.File, typeDef *Struct) {
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

func (s *Schema) generateGoTypeProxy(file *os.File, typeDef *Struct, mutable bool) {
	typeName := PropImmutable + typeDef.Name
	if mutable {
		typeName = PropMutable + typeDef.Name
	}

	fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
	fmt.Fprintf(file, "\tobjID int32\n")
	fmt.Fprintf(file, "\tkeyID wasmlib.Key32\n")
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\nfunc (o %s) Exists() bool {\n", typeName)
	fmt.Fprintf(file, "\treturn wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)\n")
	fmt.Fprintf(file, "}\n")

	if mutable {
		fmt.Fprintf(file, "\nfunc (o %s) SetValue(value *%s) {\n", typeName, typeDef.Name)
		fmt.Fprintf(file, "\twasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())\n")
		fmt.Fprintf(file, "}\n")
	}

	fmt.Fprintf(file, "\nfunc (o %s) Value() *%s {\n", typeName, typeDef.Name)
	fmt.Fprintf(file, "\treturn New%sFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))\n", typeDef.Name)
	fmt.Fprintf(file, "}\n")
}

func (s *Schema) generateGoWasmMain() error {
	err := os.MkdirAll("wasmmain", 0755)
	if err != nil {
		return err
	}

	file, err := os.Create("wasmmain/main.go")
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
		s.flushConsts(func(name string, value string, padLen int) {})
		return
	}

	fmt.Fprintf(file, "\nconst (\n")
	s.flushConsts(func(name string, value string, padLen int) {
		fmt.Fprintf(file, "\t%s = %s\n", pad(name, padLen), value)
	})
	fmt.Fprintf(file, ")\n")
}
