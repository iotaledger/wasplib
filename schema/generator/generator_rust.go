package generator

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
)

var rustTypes = StringMap{
	"Address":    "ScAddress",
	"AgentId":    "ScAgentId",
	"ChainId":    "ScChainId",
	"Color":      "ScColor",
	"ContractId": "ScContractId",
	"Hash":       "ScHash",
	"Hname":      "ScHname",
	"Int":        "i64",
	"String":     "String",
}

func (s *Schema) GenerateRust() error {
	err := os.MkdirAll("src", 0755)
	if err != nil { return err }
	err = os.Chdir("src")
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = s.GenerateRustLib()
	if err != nil {
		return  err
	}
	err = s.GenerateRustSchema()
	if err != nil {
		return  err
	}
	return s.GenerateRustTypes()
}

func (s *Schema) GenerateRustLib() error {
	file, err := os.Create("lib.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "use %s::*;\n", s.Name)
	fmt.Fprintf(file, "use schema::*;\n")
	fmt.Fprintf(file, "use wasplib::client::*;\n\n")

	fmt.Fprintf(file, "mod %s;\n", s.Name)
	fmt.Fprintf(file, "mod schema;\n")
	if len(s.Types) != 0 {
		fmt.Fprintf(file, "mod types;\n")
	}

	fmt.Fprintf(file, "\n#[no_mangle]\n")
	fmt.Fprintf(file, "fn on_load() {\n")
	fmt.Fprintf(file, "    let exports = ScExports::new();\n")
	for _, funcDef := range s.Funcs {
		name := snake(funcDef.Name)
		fmt.Fprintf(file, "    exports.add_call(FUNC_%s, func_%s);\n", upper(name), name)
	}
	for _, viewDef := range s.Views {
		name := snake(viewDef.Name)
		fmt.Fprintf(file, "    exports.add_view(VIEW_%s, view_%s);\n", upper(name), name)
	}
	fmt.Fprintf(file, "}\n")

	//TODO generate parameter checks

	return nil
}

func (s *Schema) GenerateRustSchema() error {
	file, err := os.Create("schema.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "#![allow(dead_code)]\n\n")
	fmt.Fprintf(file, "use wasplib::client::*;\n\n")

	fmt.Fprintf(file, "pub const SC_NAME: &str = \"%s\";\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "pub const SC_DESCRIPTION: &str =  \"%s\";\n", s.Description)
	}
	hName := coretypes.Hn(s.Name)
	fmt.Fprintf(file, "pub const SC_HNAME: ScHname = ScHname(0x%s);\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			param := s.Params[name]
			name = upper(snake(name))
			fmt.Fprintf(file, "pub const PARAM_%s: &str = \"%s\";\n", name, param.Alias)
		}
	}

	if len(s.Vars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.Vars {
			name := upper(snake(field.Name))
			fmt.Fprintf(file, "pub const VAR_%s: &str = \"%s\";\n", name, field.Alias)
		}
	}

	if len(s.Funcs)+len(s.Views) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := upper(snake(funcDef.Name))
			fmt.Fprintf(file, "pub const FUNC_%s: &str = \"%s\";\n", name, funcDef.Name)
		}
		for _, viewDef := range s.Views {
			name := upper(snake(viewDef.Name))
			fmt.Fprintf(file, "pub const VIEW_%s: &str = \"%s\";\n", name, viewDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := upper(snake(funcDef.Name))
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "pub const HFUNC_%s: ScHname = ScHname(0x%s);\n", name, hName.String())
		}
		for _, viewDef := range s.Views {
			name := upper(snake(viewDef.Name))
			hName = coretypes.Hn(viewDef.Name)
			fmt.Fprintf(file, "pub const HVIEW_%s: ScHname = ScHname(0x%s);\n", name, hName.String())
		}
	}
	return nil
}

func (s *Schema) GenerateRustTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	file, err := os.Create("types.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "use wasplib::client::*;\n")

	// write structs
	for _, typeDef := range s.Types {
		fmt.Fprintf(file, "\n//@formatter:off\n")
		fmt.Fprintf(file, "pub struct %s {\n", typeDef.Name)
		nameLen := 0
		typeLen := 0
		for _, field := range typeDef.Fields {
			fldName := snake(field.Name)
			if nameLen < len(fldName) { nameLen = len(fldName) }
			rustType := rustTypes[field.Type]
			if typeLen < len(rustType) { typeLen = len(rustType) }
		}
		for _, field := range typeDef.Fields {
			fldName := pad(snake(field.Name) + ":", nameLen+1)
			rfldType := pad(rustTypes[field.Type] + ",", typeLen+1)
			fmt.Fprintf(file, "    pub %s %s%s\n", fldName, rfldType, field.Comment)
		}
		fmt.Fprintf(file, "}\n")
		fmt.Fprintf(file, "//@formatter:on\n")
	}

	// write encoder and decoder for structs
	for _, typeDef := range s.Types {
		funcName := lower(snake(typeDef.Name))
		fmt.Fprintf(file, "\npub fn encode_%s(o: &%s) -> Vec<u8> {\n", funcName, typeDef.Name)
		fmt.Fprintf(file, "    let mut encode = BytesEncoder::new();\n")
		for _, field := range typeDef.Fields {
			name := snake(field.Name)
			ref := "&"
			if field.Type == "Int" || field.Type == "Hname" {
				ref = ""
			}
			fmt.Fprintf(file, "    encode.%s(%so.%s);\n", lower(field.Type), ref, name)
		}
		fmt.Fprintf(file, "    return encode.data();\n}\n")

		fmt.Fprintf(file, "\npub fn decode_%s(bytes: &[u8]) -> %s {\n", funcName, typeDef.Name)
		fmt.Fprintf(file, "    let mut decode = BytesDecoder::new(bytes);\n    %s {\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			name := snake(field.Name)
			fmt.Fprintf(file, "        %s: decode.%s(),\n", name, lower(field.Type))
		}
		fmt.Fprintf(file, "    }\n}\n")
	}

	return nil
}
