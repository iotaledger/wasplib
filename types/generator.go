package types

import (
	"errors"
	"sort"
	"strings"
)

type Generator struct {
	jsonTypes JsonTypes
	keys      []string
	maxCamel   int
	maxName   int
	maxType   int
	camels     map[string]string
	types     map[string]string
	comments  map[string]string
}

func camelcase(name string) string {
	index := strings.Index(name, "_")
	for index > 0 {
		c := name[index+1 : index+2]
		name = name[:index] + strings.ToUpper(c) + name[index+2:]
		index = strings.Index(name, "_")
	}
	return name
}

func (gen *Generator) LoadTypes(path string) error {
	jsonTypes, err := LoadTypes(path)
	if err != nil {
		return err
	}
	gen.jsonTypes = jsonTypes

	gen.keys = make([]string, 0)
	for key := range jsonTypes {
		gen.keys = append(gen.keys, key)
	}
	sort.Strings(gen.keys)
	return nil
}

func (gen *Generator) SplitComments(structName string, myTypes map[string]string) error {
	gen.camels = make(map[string]string)
	gen.types = make(map[string]string)
	gen.comments = make(map[string]string)
	gen.maxCamel = 0
	gen.maxName = 0
	gen.maxType = 0
	for _, fld := range gen.jsonTypes[structName] {
		for name, typeName := range fld {
			comment := ""
			index := strings.Index(typeName, "//")
			if index > 0 {
				comment = " // " + strings.TrimSpace(typeName[index+2:])
				typeName = strings.TrimSpace(typeName[:index])
			}
			myType, ok := myTypes[typeName]
			if !ok {
				return errors.New("Invalid type: " + typeName)
			}
			camel := camelcase(name)
			gen.camels[name] = camel
			gen.types[name] = myType
			gen.comments[name] = comment
			if len(camel) > gen.maxCamel {
				gen.maxCamel = len(camel)
			}
			if len(name) > gen.maxName {
				gen.maxName = len(name)
			}
			if len(myType) > gen.maxType {
				gen.maxType = len(myType)
			}
		}
	}
	return nil
}
