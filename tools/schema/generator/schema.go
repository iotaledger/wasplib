// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type (
	FieldMap    map[string]*Field
	FieldMapMap map[string]FieldMap
)

type (
	StringMap    map[string]string
	StringMapMap map[string]StringMap
)

type FuncDesc struct {
	Access  string    `json:"access,omitempty"`
	Params  StringMap `json:"params,omitempty"`
	Results StringMap `json:"results,omitempty"`
}
type FuncDescMap map[string]*FuncDesc

type JSONSchema struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Structs     StringMapMap `json:"structs"`
	Typedefs    StringMap    `json:"typedefs"`
	State       StringMap    `json:"state"`
	Funcs       FuncDescMap  `json:"funcs"`
	Views       FuncDescMap  `json:"views"`
}

type FuncDef struct {
	Access   string
	Kind     string
	FuncName string
	String   string
	Params   []*Field
	Results  []*Field
	Type     string
}

func (f *FuncDef) nameLen(smallest int) int {
	if len(f.Results) != 0 {
		return 7
	}
	if len(f.Params) != 0 {
		return 6
	}
	return smallest
}

type Schema struct {
	Name          string
	FullName      string
	Description   string
	KeyID         int
	ConstLen      int
	ConstNames    []string
	ConstValues   []string
	CoreContracts bool
	Funcs         []*FuncDef
	NewTypes      map[string]bool
	Params        []*Field
	Results       []*Field
	StateVars     []*Field
	Structs       []*Struct
	Typedefs      []*Field
	Views         []*FuncDef
}

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) Compile(jsonSchema *JSONSchema) error {
	s.FullName = strings.TrimSpace(jsonSchema.Name)
	if s.FullName == "" {
		return fmt.Errorf("missing contract name")
	}
	s.Name = lower(s.FullName)
	s.Description = strings.TrimSpace(jsonSchema.Description)

	err := s.compileTypes(jsonSchema)
	if err != nil {
		return err
	}
	err = s.compileSubtypes(jsonSchema)
	if err != nil {
		return err
	}
	params := make(FieldMap)
	results := make(FieldMap)
	err = s.compileFuncs(jsonSchema, &params, &results, false)
	if err != nil {
		return err
	}
	err = s.compileFuncs(jsonSchema, &params, &results, true)
	if err != nil {
		return err
	}
	for _, name := range sortedFields(params) {
		s.Params = append(s.Params, params[name])
	}
	for _, name := range sortedFields(results) {
		s.Results = append(s.Results, results[name])
	}
	return s.compileStateVars(jsonSchema)
}

func (s *Schema) CompileField(fldName, fldType string) (*Field, error) {
	field := &Field{}
	err := field.Compile(s, fldName, fldType)
	if err != nil {
		return nil, err
	}
	return field, nil
}

func (s *Schema) compileFuncs(jsonSchema *JSONSchema, params, results *FieldMap, views bool) (err error) {
	// TODO check for clashing Hnames

	kind := "func"
	jsonFuncs := jsonSchema.Funcs
	if views {
		kind = "view"
		jsonFuncs = jsonSchema.Views
	}
	for _, funcName := range sortedFuncDescs(jsonFuncs) {
		if views && jsonSchema.Funcs[funcName] != nil {
			return fmt.Errorf("duplicate func/view name")
		}
		funcDesc := jsonFuncs[funcName]
		f := &FuncDef{}
		f.String = funcName
		f.Kind = capitalize(kind)
		f.Type = capitalize(funcName)
		f.FuncName = kind + f.Type
		f.Access = funcDesc.Access
		f.Params, err = s.compileFuncFields(funcDesc.Params, params, "param")
		if err != nil {
			return err
		}
		f.Results, err = s.compileFuncFields(funcDesc.Results, results, "result")
		if err != nil {
			return err
		}
		s.Funcs = append(s.Funcs, f)
	}
	return nil
}

func (s *Schema) compileFuncFields(fieldMap StringMap, allFieldMap *FieldMap, what string) ([]*Field, error) {
	fields := make([]*Field, 0, len(fieldMap))
	fieldNames := make(StringMap)
	fieldAliases := make(StringMap)
	for _, fldName := range sortedKeys(fieldMap) {
		fldType := fieldMap[fldName]
		field, err := s.CompileField(fldName, fldType)
		if err != nil {
			return nil, err
		}
		if _, ok := fieldNames[field.Name]; ok {
			return nil, fmt.Errorf("duplicate %s name", what)
		}
		fieldNames[field.Name] = field.Name
		if _, ok := fieldAliases[field.Alias]; ok {
			return nil, fmt.Errorf("duplicate %s alias", what)
		}
		fieldAliases[field.Alias] = field.Alias
		existing, ok := (*allFieldMap)[field.Name]
		if !ok {
			(*allFieldMap)[field.Name] = field
			existing = field
		}
		if existing.Alias != field.Alias {
			return nil, fmt.Errorf("redefined %s alias: '%s' != '%s", what, existing.Alias, field.Alias)
		}
		if existing.Type != field.Type {
			return nil, fmt.Errorf("redefined %s type", what)
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func (s *Schema) compileStateVars(jsonSchema *JSONSchema) error {
	varNames := make(StringMap)
	varAliases := make(StringMap)
	for _, varName := range sortedKeys(jsonSchema.State) {
		varType := jsonSchema.State[varName]
		varDef, err := s.CompileField(varName, varType)
		if err != nil {
			return err
		}
		if _, ok := varNames[varDef.Name]; ok {
			return fmt.Errorf("duplicate var name")
		}
		varNames[varDef.Name] = varDef.Name
		if _, ok := varAliases[varDef.Alias]; ok {
			return fmt.Errorf("duplicate var alias")
		}
		varAliases[varDef.Alias] = varDef.Alias
		s.StateVars = append(s.StateVars, varDef)
	}
	return nil
}

func (s *Schema) compileSubtypes(jsonSchema *JSONSchema) error {
	varNames := make(StringMap)
	varAliases := make(StringMap)
	for _, varName := range sortedKeys(jsonSchema.Typedefs) {
		varType := jsonSchema.Typedefs[varName]
		varDef, err := s.CompileField(varName, varType)
		if err != nil {
			return err
		}
		if _, ok := varNames[varDef.Name]; ok {
			return fmt.Errorf("duplicate sybtype name")
		}
		varNames[varDef.Name] = varDef.Name
		if _, ok := varAliases[varDef.Alias]; ok {
			return fmt.Errorf("duplicate subtype alias")
		}
		varAliases[varDef.Alias] = varDef.Alias
		s.Typedefs = append(s.Typedefs, varDef)
	}
	return nil
}

func (s *Schema) compileTypes(jsonSchema *JSONSchema) error {
	for _, typeName := range sortedMaps(jsonSchema.Structs) {
		fieldMap := jsonSchema.Structs[typeName]
		typeDef := &Struct{}
		typeDef.Name = typeName
		fieldNames := make(StringMap)
		fieldAliases := make(StringMap)
		for _, fldName := range sortedKeys(fieldMap) {
			fldType := fieldMap[fldName]
			field, err := s.CompileField(fldName, fldType)
			if err != nil {
				return err
			}
			if field.Optional {
				return fmt.Errorf("type field cannot be optional")
			}
			if _, ok := fieldNames[field.Name]; ok {
				return fmt.Errorf("duplicate field name")
			}
			fieldNames[field.Name] = field.Name
			if _, ok := fieldAliases[field.Alias]; ok {
				return fmt.Errorf("duplicate field alias")
			}
			fieldAliases[field.Alias] = field.Alias
			typeDef.Fields = append(typeDef.Fields, field)
		}
		s.Structs = append(s.Structs, typeDef)
	}
	return nil
}

func (s *Schema) scanExistingCode(file *os.File, funcRegexp *regexp.Regexp) ([]string, StringMap, error) {
	defer file.Close()
	existing := make(StringMap)
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := funcRegexp.FindStringSubmatch(line)
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

func (s *Schema) appendConst(name, value string) {
	if s.ConstLen < len(name) {
		s.ConstLen = len(name)
	}
	s.ConstNames = append(s.ConstNames, name)
	s.ConstValues = append(s.ConstValues, value)
}

func (s *Schema) flushConsts(printer func(name string, value string, padLen int)) {
	for i, name := range s.ConstNames {
		printer(name, s.ConstValues[i], s.ConstLen)
	}
	s.ConstLen = 0
	s.ConstNames = nil
	s.ConstValues = nil
}

func (s *Schema) crateOrWasmLib(withContract, withHost bool) string {
	if s.CoreContracts {
		retVal := useCrate
		if withContract {
			retVal += "use crate::corecontracts::" + s.Name + "::*;\n"
		}
		if withHost {
			retVal += "use crate::host::*;\n"
		}
		return retVal
	}
	retVal := useWasmLib
	if withHost {
		retVal += useWasmLibHost
	}
	return retVal
}
