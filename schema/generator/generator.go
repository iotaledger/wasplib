package generator

import (
	"regexp"
	"sort"
	"strings"
)

type Generator struct {
	schema *JsonSchema
}

var camelRegExp = regexp.MustCompile("_[a-z]")
var snakeRegExp = regexp.MustCompile("[a-z0-9][A-Z]")

// convert lowercase snake case to camel case
func camel(name string) string {
	return camelRegExp.ReplaceAllStringFunc(name, func(sub string) string {
		return strings.ToUpper(sub[1:])
	})
}

// capitalize first letter
func capitalize(name string) string {
	return upper(name[:1]) + name[1:]
}

// convert to lower case
func lower(name string) string {
	return strings.ToLower(name)
}

// pad to specified size with spaces
func pad(name string, size int) string {
	for i := len(name); i < size; i++ {
		name += " "
	}
	return name
}

// convert camel case to lower case snake case
func snake(name string) string {
	return snakeRegExp.ReplaceAllStringFunc(name, func(sub string) string {
		return sub[:1] + "_" + lower(sub[1:])
	})
}

// convert to upper case
func upper(name string) string {
	return strings.ToUpper(name)
}

func sortedFields(dict FieldMap) []string {
	keys := make([]string, 0)
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedKeys(dict StringMap) []string {
	keys := make([]string, 0)
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedMaps(dict StringMapMap) []string {
	keys := make([]string, 0)
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
