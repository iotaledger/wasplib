package generator

import (
	"fmt"
	"github.com/iotaledger/wasplib/client"
	"regexp"
	"strings"
)

var fldNameRegexp = regexp.MustCompile("^[a-z][a-zA-Z0-9]*$")
var fldAliasRegexp = regexp.MustCompile("^[a-zA-Z0-9_$#@*%\\-]+$")
var fldTypeRegexp = regexp.MustCompile("^[A-Z][a-zA-Z0-9]+$")

var FieldTypes = map[string]int32{
	"Address":    client.TYPE_ADDRESS,
	"AgentId":    client.TYPE_AGENT_ID,
	"Bytes":      client.TYPE_BYTES,
	"ChainId":    client.TYPE_CHAIN_ID,
	"Color":      client.TYPE_COLOR,
	"ContractId": client.TYPE_CONTRACT_ID,
	"Hash":       client.TYPE_HASH,
	"Hname":      client.TYPE_HNAME,
	"Int":        client.TYPE_INT,
	"String":     client.TYPE_STRING,
}

type Field struct {
	Alias    string // internal name alias, can be different from Name
	Comment  string
	Name     string // external name for this field
	Array    bool
	Optional bool
	Type     string
	TypeId   int32
}

func (f *Field) Compile(schema *Schema, fldName string, fldType string) error {
	fldName = strings.TrimSpace(fldName)
	f.Name = fldName
	f.Alias = fldName
	index := strings.Index(fldName, "=")
	if index >= 0 {
		f.Name = strings.TrimSpace(fldName[:index])
		f.Alias = strings.TrimSpace(fldName[index+1:])
	}
	if !fldNameRegexp.MatchString(f.Name) {
		return fmt.Errorf("invalid field name: %s", f.Name)
	}
	if !fldAliasRegexp.MatchString(f.Alias) {
		return fmt.Errorf("invalid field alias: %s", f.Alias)
	}

	fldType = strings.TrimSpace(fldType)
	if strings.HasPrefix(fldType, "?") {
		f.Optional = true
		fldType = strings.TrimSpace(fldType[1:])
	}
	if strings.HasPrefix(fldType, "[]") {
		f.Array = true
		fldType = strings.TrimSpace(fldType[2:])
	}
	index = strings.Index(fldType, "//")
	if index >= 0 {
		f.Comment = " " + fldType[index:]
		fldType = strings.TrimSpace(fldType[:index])
	}
	f.Type = fldType
	if !fldTypeRegexp.MatchString(f.Type) {
		return fmt.Errorf("invalid field type: %s", f.Type)
	}
	typeId, ok := FieldTypes[f.Type]
	if ok {
		f.TypeId = typeId
		return nil
	}
	for _,typeDef := range schema.Types {
		if f.Type == typeDef.Name {
			return nil
		}
	}
	return fmt.Errorf("invalid field type: %s", f.Type)
}
