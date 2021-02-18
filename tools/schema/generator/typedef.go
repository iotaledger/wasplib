package generator

import (
	"fmt"
	"os"
)

type TypeDef struct {
	Name   string
	Fields []*Field
}

func (td *TypeDef) GenerateJavaType(contract string) error {
	file, err := os.Create("types/" + td.Name + ".java")
	if err != nil {
		return err
	}
	defer file.Close()

	// calculate padding
	nameLen, typeLen := calculatePadding(td.Fields, javaTypes, false)

	// write file header
	fmt.Fprintf(file, copyright(true))
	fmt.Fprintf(file, "\npackage org.iota.wasp.contracts.%s.types;\n\n", contract)
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.bytes.*;\n")
	fmt.Fprintf(file, "import org.iota.wasp.wasmlib.hashtypes.*;\n\n")

	fmt.Fprintf(file, "public class %s{\n", td.Name)

	// write struct layout
	fmt.Fprintf(file, "\t//@formatter:off\n")
	for _, field := range td.Fields {
		fldName := capitalize(field.Name) + ";"
		fldType := pad(javaTypes[field.Type], typeLen)
		if field.Comment != "" {
			fldName = pad(fldName, nameLen + 1)
		}
		fmt.Fprintf(file, "\tpublic %s %s%s\n", fldType, fldName, field.Comment)
	}
	fmt.Fprintf(file, "\t//@formatter:on\n")

	// write default constructor
	fmt.Fprintf(file, "\n\tpublic %s() {\n\t}\n", td.Name)

	// write constructor from byte array
	fmt.Fprintf(file, "\n\tpublic %s(byte[] bytes) {\n", td.Name)
	fmt.Fprintf(file, "\t\tBytesDecoder decode = new BytesDecoder(bytes);\n")
	for _, field := range td.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "\t\t%s = decode.%s();\n", name, field.Type)
	}
	fmt.Fprintf(file, "\t}\n")

	// write conversion to byte array
	fmt.Fprintf(file, "\n\tpublic byte[] toBytes(){\n")
	fmt.Fprintf(file, "\t\treturn new BytesEncoder().\n")
	for _, field := range td.Fields {
		name := capitalize(field.Name)
		fmt.Fprintf(file, "\t\t\t\t%s(%s).\n", field.Type, name)
	}
	fmt.Fprintf(file, "\t\t\t\tData();\n\t}\n")

	fmt.Fprintf(file, "}\n")
	return nil
}


