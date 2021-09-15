// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"os"
)

type Struct struct {
	Name   string
	Fields []*Field
}

func (td *Struct) GenerateJavaType(contract string) error {
	file, err := os.Create("types/" + td.Name + ".java")
	if err != nil {
		return err
	}
	defer file.Close()

	// calculate padding
	nameLen, typeLen := calculatePadding(td.Fields, javaTypes, false)

	// write file header
	fmt.Fprint(file, copyright(true))
	fmt.Fprintf(file, "\npackage org.iota.wasp.contracts.%s.types;\n\n", contract)
	fmt.Fprint(file, "import org.iota.wasp.wasmlib.bytes.*;\n")
	fmt.Fprint(file, "import org.iota.wasp.wasmlib.hashtypes.*;\n\n")

	fmt.Fprintf(file, "public class %s {\n", td.Name)

	// write struct layout
	if len(td.Fields) > 1 {
		fmt.Fprint(file, "    // @formatter:off\n")
	}
	for _, field := range td.Fields {
		fldName := capitalize(field.Name) + ";"
		fldType := pad(javaTypes[field.Type], typeLen)
		if field.Comment != "" {
			fldName = pad(fldName, nameLen+1)
		}
		fmt.Fprintf(file, "    public %s %s%s\n", fldType, fldName, field.Comment)
	}
	if len(td.Fields) > 1 {
		fmt.Fprint(file, "    // @formatter:on\n")
	}

	// write default constructor
	fmt.Fprintf(file, "\n    public %s() {\n    }\n", td.Name)

	// write constructor from byte array
	fmt.Fprintf(file, "\n    public %s(byte[] bytes) {\n", td.Name)
	fmt.Fprintf(file, "        BytesDecoder decode = new BytesDecoder(bytes);\n")
	for _, field := range td.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "        %s = decode.%s();\n", name, field.Type)
	}
	fmt.Fprintf(file, "        decode.Close();\n")
	fmt.Fprintf(file, "    }\n")

	// write conversion to byte array
	fmt.Fprintf(file, "\n    public byte[] toBytes() {\n")
	fmt.Fprintf(file, "        return new BytesEncoder().\n")
	for _, field := range td.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "                %s(%s).\n", field.Type, name)
	}
	fmt.Fprintf(file, "                Data();\n    }\n")

	fmt.Fprintf(file, "}\n")
	return nil
}
