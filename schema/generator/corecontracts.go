package generator

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"os"
)

func GenerateGoCoreContractsSchema(coreSchemas []*Schema) error {
	file, err := os.Create("../client/corecontracts.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "package client\n")

	for _, schema := range coreSchemas {
		scName := capitalize(schema.Name)
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\nconst Core%s = ScHname(0x%s)\n", scName, scHname.String())
		for _, funcDef := range schema.Funcs {
			funcHname := coretypes.Hn(funcDef.Name)
			funcName := capitalize(funcDef.Name)
			fmt.Fprintf(file, "const Core%sFunc%s = ScHname(0x%s)\n", scName, funcName, funcHname.String())
		}
		for _, viewDef := range schema.Views {
			viewHname := coretypes.Hn(viewDef.Name)
			viewName := capitalize(viewDef.Name)
			fmt.Fprintf(file, "const Core%sView%s = ScHname(0x%s)\n", scName, viewName, viewHname.String())
		}

		if len(schema.Params) != 0 {
			fmt.Fprintln(file)
			for _, name := range sortedFields(schema.Params) {
				param := schema.Params[name]
				name = capitalize(param.Name)
				fmt.Fprintf(file, "const Core%sParam%s = Key(\"%s\")\n", scName, name, param.Alias)
			}
		}
	}
	return nil
}

func GenerateRustCoreContractsSchema(coreSchemas []*Schema) error {
	file, err := os.Create("../rust/wasplib/src/client/corecontracts.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright())
	fmt.Fprintf(file, "use super::hashtypes::*;\n")

	for _, schema := range coreSchemas {
		scName := upper(snake(schema.Name))
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\npub const CORE_%s: ScHname = ScHname(0x%s);\n", scName, scHname.String())
		for _, funcDef := range schema.Funcs {
			funcHname := coretypes.Hn(funcDef.Name)
			funcName := upper(snake(funcDef.Name))
			fmt.Fprintf(file, "pub const CORE_%s_FUNC_%s: ScHname = ScHname(0x%s);\n", scName, funcName, funcHname.String())
		}
		for _, viewDef := range schema.Views {
			viewHname := coretypes.Hn(viewDef.Name)
			viewName := upper(snake(viewDef.Name))
			fmt.Fprintf(file, "pub const CORE_%s_VIEW_%s: ScHname = ScHname(0x%s);\n", scName, viewName, viewHname.String())
		}

		if len(schema.Params) != 0 {
			fmt.Fprintln(file)
			for _, name := range sortedFields(schema.Params) {
				param := schema.Params[name]
				name = upper(snake(name))
				fmt.Fprintf(file, "pub const CORE_%s_PARAM_%s: &str = \"%s\";\n", scName, name, param.Alias)
			}
		}
	}
	return nil
}

