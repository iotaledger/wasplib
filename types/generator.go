package types

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type Generator struct {
	schema *Schema
	keys      []string
	maxCamel  int
	maxName   int
	maxType   int
	camels    map[string]string
	types     map[string]string
	comments  map[string]string
}

var camelRegExp = regexp.MustCompile("_[a-z]")
var snakeRegExp = regexp.MustCompile("[a-z][A-Z]")

func camelcase(name string) string {
	name = camelRegExp.ReplaceAllStringFunc(name, func(sub string) string {
		return strings.ToUpper(sub[1:])
	})
	return strings.ToUpper(name[:1]) + name[1:]
}

func snakecase(name string) string {
	name = snakeRegExp.ReplaceAllStringFunc(name, func(sub string) string {
		return sub[:1] + "_" + sub[1:]
	})
	return strings.ToUpper(name)
}

func (gen *Generator) LoadTypes(path string) error {
	schema, err := LoadSchema(path)
	if err != nil {
		return err
	}
	gen.schema = schema

	gen.keys = make([]string, 0)
	for key := range schema.Types {
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
	types := gen.schema.Types
	for _, fld := range types[structName] {
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
