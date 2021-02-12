// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"strings"
)

type FieldMap map[string]*Field
type FieldMapMap map[string]FieldMap

type StringMap map[string]string
type StringMapMap map[string]StringMap

type JsonSchema struct {
	Description string       `json:"description"`
	Funcs       StringMapMap `json:"funcs"`
	Name        string       `json:"name"`
	Types       StringMapMap `json:"types"`
	Vars        StringMap    `json:"vars"`
	Views       StringMapMap `json:"views"`
}

type FuncDef struct {
	Name        string
	FullName    string
	Annotations StringMap
	Params      []*Field
}

type TypeDef struct {
	Name   string
	Fields []*Field
}

type Schema struct {
	Description string
	Funcs       []*FuncDef
	Name        string
	Params      FieldMap
	Types       []*TypeDef
	Vars        []*Field
	Views       []*FuncDef
}

func NewSchema() *Schema {
	s := &Schema{}
	s.Params = make(FieldMap)
	return s
}

func (s *Schema) Compile(jsonSchema *JsonSchema) error {
	s.Name = strings.TrimSpace(jsonSchema.Name)
	if s.Name == "" {
		return fmt.Errorf("missing contract name")
	}
	s.Description = strings.TrimSpace(jsonSchema.Description)

	err := s.compileTypes(jsonSchema)
	if err != nil {
		return err
	}
	err = s.compileFuncs(jsonSchema, false)
	if err != nil {
		return err
	}
	err = s.compileFuncs(jsonSchema, true)
	if err != nil {
		return err
	}
	return s.compileVars(jsonSchema)
}

func (s *Schema) CompileField(fldName string, fldType string) (*Field, error) {
	field := &Field{}
	err := field.Compile(s, fldName, fldType)
	if err != nil {
		return nil, err
	}
	return field, nil
}

func (s *Schema) compileFuncs(jsonSchema *JsonSchema, views bool) error {
	//TODO check for clashing Hnames

	jsonFuncs := jsonSchema.Funcs
	if views {
		jsonFuncs = jsonSchema.Views
	}
	for _, funcName := range sortedMaps(jsonFuncs) {
		if views && jsonSchema.Funcs[funcName] != nil {
			return fmt.Errorf("duplicate func/view name")
		}
		paramMap := jsonFuncs[funcName]
		funcDef := &FuncDef{}
		funcDef.Name = funcName
		funcDef.FullName = "func" + capitalize(funcDef.Name)
		if views {
			funcDef.FullName = "view" + funcDef.FullName[4:]
		}
		funcDef.Annotations = make(StringMap)
		fieldNames := make(StringMap)
		fieldAliases := make(StringMap)
		for _, fldName := range sortedKeys(paramMap) {
			fldType := paramMap[fldName]
			if strings.HasPrefix(fldName, "#") {
				funcDef.Annotations[fldName] = strings.TrimSpace(fldType)
				continue
			}
			param, err := s.CompileField(fldName, fldType)
			if err != nil {
				return err
			}
			if _, ok := fieldNames[param.Name]; ok {
				return fmt.Errorf("duplicate param name")
			}
			fieldNames[param.Name] = param.Name
			if _, ok := fieldAliases[param.Alias]; ok {
				return fmt.Errorf("duplicate param alias")
			}
			fieldAliases[param.Alias] = param.Alias
			existing, ok := s.Params[param.Name]
			if !ok {
				s.Params[param.Name] = param
				existing = param
			}
			if existing.Alias != param.Alias {
				return fmt.Errorf("redefined param alias")
			}
			if existing.Type != param.Type {
				return fmt.Errorf("redefined param type")
			}
			funcDef.Params = append(funcDef.Params, param)
		}
		if views {
			s.Views = append(s.Views, funcDef)
		} else {
			s.Funcs = append(s.Funcs, funcDef)
		}
	}
	return nil
}

func (s *Schema) compileTypes(jsonSchema *JsonSchema) error {
	for _, typeName := range sortedMaps(jsonSchema.Types) {
		fieldMap := jsonSchema.Types[typeName]
		typeDef := &TypeDef{}
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
		s.Types = append(s.Types, typeDef)
	}
	return nil
}

func (s *Schema) compileVars(jsonSchema *JsonSchema) error {
	varNames := make(StringMap)
	varAliases := make(StringMap)
	for _, varName := range sortedKeys(jsonSchema.Vars) {
		varType := jsonSchema.Vars[varName]
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
		s.Vars = append(s.Vars, varDef)
	}
	return nil
}
