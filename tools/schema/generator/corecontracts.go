// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"os"

	"github.com/iotaledger/wasp/packages/coretypes"
)

func GenerateGoCoreContractsSchema(coreSchemas []*Schema) error {
	file, err := os.Create("../corecontracts.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package wasmlib\n")

	for _, schema := range coreSchemas {
		scName := capitalize(schema.Name)
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\nconst Core%s = ScHname(0x%s)\n", scName, scHname.String())
		for _, funcDef := range schema.Funcs {
			funcHname := coretypes.Hn(funcDef.Name)
			funcName := capitalize(funcDef.FullName)
			fmt.Fprintf(file, "const Core%s%s = ScHname(0x%s)\n", scName, funcName, funcHname.String())
		}

		if len(schema.Params) != 0 {
			fmt.Fprintln(file)
			for _, param := range schema.Params {
				name := capitalize(param.Name)
				fmt.Fprintf(file, "const Core%sParam%s = Key(\"%s\")\n", scName, name, param.Alias)
			}
		}
	}
	return nil
}

func GenerateJavaCoreContractsSchema(coreSchemas []*Schema) error {
	file, err := os.Create("../Core.java")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "package org.iota.wasp.wasmlib.keys;\n\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.hashtypes.*;\n\n")
	fmt.Fprintf(file, "public class Core {\n")

	for _, schema := range coreSchemas {
		scName := capitalize(schema.Name)
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\n    public static final ScHname %s = new ScHname(0x%s);\n", scName, scHname.String())
		for _, funcDef := range schema.Funcs {
			funcHname := coretypes.Hn(funcDef.Name)
			funcName := capitalize(funcDef.FullName)
			fmt.Fprintf(file, "    public static final ScHname %s%s = new ScHname(0x%s);\n", scName, funcName, funcHname.String())
		}

		if len(schema.Params) != 0 {
			fmt.Fprintln(file)
			for _, param := range schema.Params {
				name := capitalize(param.Name)
				fmt.Fprintf(file, "    public static final Key %sParam%s = new Key(\"%s\");\n", scName, name, param.Alias)
			}
		}
	}

	fmt.Fprintf(file, "}\n")
	return nil
}

func GenerateRustCoreContractsSchema(coreSchemas []*Schema) error {
	file, err := os.Create("../corecontracts.rs")
	if err != nil {
		return err
	}
	defer file.Close()

	// write file header
	fmt.Fprintln(file, copyright(true))
	fmt.Fprintf(file, "use crate::hashtypes::*;\n")

	for _, schema := range coreSchemas {
		scName := upper(snake(schema.Name))
		scHname := coretypes.Hn(schema.Name)
		fmt.Fprintf(file, "\npub const CORE_%s: ScHname = ScHname(0x%s);\n", scName, scHname.String())
		for _, funcDef := range schema.Funcs {
			funcHname := coretypes.Hn(funcDef.Name)
			funcName := upper(snake(funcDef.FullName))
			fmt.Fprintf(file, "pub const CORE_%s_%s: ScHname = ScHname(0x%s);\n", scName, funcName, funcHname.String())
		}

		if len(schema.Params) != 0 {
			fmt.Fprintln(file)
			for _, param := range schema.Params {
				name := upper(snake(param.Name))
				fmt.Fprintf(file, "pub const CORE_%s_PARAM_%s: &str = \"%s\";\n", scName, name, param.Alias)
			}
		}
	}
	return nil
}
