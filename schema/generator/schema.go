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

	for _, funcName := range sortedMaps(jsonSchema.Funcs) {
		paramMap := jsonSchema.Funcs[funcName]
		funcDef := &FuncDef{}
		funcDef.Name = funcName
		funcDef.Annotations = make(StringMap)
		fieldNames := make(StringMap)
		fieldAliases := make(StringMap)
		for _, fldName := range sortedKeys(paramMap) {
			fldType := paramMap[fldName]
			if strings.HasPrefix(fldName, "#") {
				funcDef.Annotations[fldName] = fldType
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
		s.Funcs = append(s.Funcs, funcDef)
	}

	for _, viewName := range sortedMaps(jsonSchema.Views) {
		if _, ok := jsonSchema.Funcs[viewName]; ok {
			return fmt.Errorf("duplicate func/view name")
		}
		paramMap := jsonSchema.Views[viewName]
		viewDef := &FuncDef{}
		viewDef.Name = viewName
		viewDef.Annotations = make(StringMap)
		fieldNames := make(StringMap)
		fieldAliases := make(StringMap)
		for _, fldName := range sortedKeys(paramMap) {
			fldType := paramMap[fldName]
			if strings.HasPrefix(fldName, "#") {
				viewDef.Annotations[fldName] = fldType
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
			viewDef.Params = append(viewDef.Params, param)
		}
		s.Views = append(s.Views, viewDef)
	}

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

func (s *Schema) CompileField(fldName string, fldType string) (*Field, error) {
	field := &Field{}
	err := field.Compile(s, fldName, fldType)
	if err != nil {
		return nil, err
	}
	return field, nil
}
