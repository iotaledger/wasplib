package generator

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
)

var goTypes = StringMap{
	"Address":    "*client.ScAddress",
	"AgentId":    "*client.ScAgentId",
	"ChainId":    "*client.ScChainId",
	"Color":      "*client.ScColor",
	"ContractId": "*client.ScContractId",
	"Hash":       "*client.ScHash",
	"Hname":      "client.ScHname",
	"Int":        "int64",
	"String":     "string",
}

//TODO check for clashing Hnames

func (s *Schema) GenerateGo() error {
	err := os.MkdirAll("test/" + s.Name + "/wasm", 0755)
	if err != nil { return err }
	err = os.Chdir("test/" + s.Name)
	if err != nil {
		return err
	}
	defer os.Chdir("../..")

	err = s.GenerateGoWasmMain()
	if err != nil {
		return err
	}
	err = s.GenerateGoOnLoad()
	if err != nil {
		return err
	}
	err = s.GenerateGoSchema()
	if err != nil {
		return err
	}
	return s.GenerateGoTypes()
}

func (s *Schema) GenerateGoOnLoad() error {
	file, err := os.Create("onload.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n\n")

	fmt.Fprintf(file, "\nfunc OnLoad() {\n")
	fmt.Fprintf(file, "    exports := client.NewScExports()\n")
	for _, funcDef := range s.Funcs {
		name := capitalize(funcDef.Name)
		fmt.Fprintf(file, "    exports.AddCall(Func%s, func%s)\n", name, name)
	}
	for _, viewDef := range s.Views {
		name := capitalize(viewDef.Name)
		fmt.Fprintf(file, "    exports.AddView(View%s, view%s)\n", name, name)
	}
	fmt.Fprintf(file, "}\n")

	//TODO generate parameter checks

	return nil
}

func (s *Schema) GenerateGoSchema() error {
	file, err := os.Create("schema.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n\n")

	fmt.Fprintf(file, "const ScName = \"%s\"\n", s.Name)
	if s.Description != "" {
		fmt.Fprintf(file, "const ScDescription = \"%s\"\n", s.Description)
	}
	hName := coretypes.Hn(s.Name)
	fmt.Fprintf(file, "const ScHname = client.ScHname(0x%s)\n", hName.String())

	if len(s.Params) != 0 {
		fmt.Fprintln(file)
		for _, name := range sortedFields(s.Params) {
			param := s.Params[name]
			name = capitalize(param.Name)
			fmt.Fprintf(file, "const Param%s = client.Key(\"%s\")\n", name, param.Alias)
		}
	}

	if len(s.Vars) != 0 {
		fmt.Fprintln(file)
		for _, field := range s.Vars {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "const Var%s = client.Key(\"%s\")\n", name, field.Alias)
		}
	}

	if len(s.Funcs)+len(s.Views) != 0 {
		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.Name)
			fmt.Fprintf(file, "const Func%s = \"%s\"\n", name, funcDef.Name)
		}
		for _, viewDef := range s.Views {
			name := capitalize(viewDef.Name)
			fmt.Fprintf(file, "const View%s = \"%s\"\n", name, viewDef.Name)
		}

		fmt.Fprintln(file)
		for _, funcDef := range s.Funcs {
			name := capitalize(funcDef.Name)
			hName = coretypes.Hn(funcDef.Name)
			fmt.Fprintf(file, "const HFunc%s = client.ScHname(0x%s)\n", name, hName.String())
		}
		for _, viewDef := range s.Views {
			name := capitalize(viewDef.Name)
			hName = coretypes.Hn(viewDef.Name)
			fmt.Fprintf(file, "const HView%s = client.ScHname(0x%s)\n", name, hName.String())
		}
	}

	return nil
}

func (s *Schema) GenerateGoTypes() error {
	if len(s.Types) == 0 {
		return nil
	}

	file, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "package %s\n\n", s.Name)
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client\"\n")

	// write structs
	for _, typeDef := range s.Types {
		fmt.Fprintf(file, "\ntype %s struct {\n", typeDef.Name)
		nameLen := 0
		typeLen := 0
		for _, field := range typeDef.Fields {
			if nameLen < len(field.Name) {
				nameLen = len(field.Name)
			}
			goType := goTypes[field.Type]
			if typeLen < len(goType) {
				typeLen = len(goType)
			}
		}
		for _, field := range typeDef.Fields {
			fldName := pad(capitalize(field.Name), nameLen)
			fldType := pad(goTypes[field.Type], typeLen)
			fmt.Fprintf(file, "\t%s %s%s\n", fldName, fldType, field.Comment)
		}
		fmt.Fprintf(file, "}\n")
	}

	// write encoder and decoder for structs
	for _, typeDef := range s.Types {
		fmt.Fprintf(file, "\nfunc Encode%s(o *%s) []byte {\n", typeDef.Name, typeDef.Name)
		fmt.Fprintf(file, "\treturn client.NewBytesEncoder().\n")
		for _, field := range typeDef.Fields {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "\t\t%s(o.%s).\n", field.Type, name)
		}
		fmt.Fprintf(file, "\t\tData()\n}\n")

		fmt.Fprintf(file, "\nfunc Decode%s(bytes []byte) *%s {\n", typeDef.Name, typeDef.Name)
		fmt.Fprintf(file, "\tdecode := client.NewBytesDecoder(bytes)\n\tdata := &%s{}\n", typeDef.Name)
		for _, field := range typeDef.Fields {
			name := capitalize(field.Name)
			fmt.Fprintf(file, "\tdata.%s = decode.%s()\n", name, field.Type)
		}
		fmt.Fprintf(file, "\treturn data\n}\n")
	}

	return nil
}

func (s *Schema) GenerateGoWasmMain() error {
	file, err := os.Create("wasm/" + s.Name + ".go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "// +build wasm\n\n")
	fmt.Fprintf(file, "package main\n\n")
	fmt.Fprintf(file, "import \"github.com/iotaledger/wasplib/client/wasm\"\n")
	fmt.Fprintf(file, "import \"github.com/iotaledger/eric/" + s.Name +"/test/" + s.Name +"\"\n\n")

	fmt.Fprintf(file, "func main() {\n")
	fmt.Fprintf(file, "}\n\n")

	fmt.Fprintf(file, "//export on_load\n")
	fmt.Fprintf(file, "func %sOnLoad() {\n", s.Name)
	fmt.Fprintf(file, "\twasmclient.ConnectWasmHost()\n")
	fmt.Fprintf(file, "\t%s.OnLoad()\n", s.Name)
	fmt.Fprintf(file, "}\n")

	return nil
}
