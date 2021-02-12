// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

var javaTypes = StringMap{
	"Address":    "ScAddress",
	"AgentId":    "ScAgentId",
	"ChainId":    "ScChainId",
	"Color":      "ScColor",
	"ContractId": "ScContractId",
	"Hash":       "ScHash",
	"Hname":      "Hname",
	"Int":        "long",
	"String":     "String",
}

//func GenerateJavaTypes(path string) error {
//	gen := &Generator{}
//	err := gen.LoadTypes(path)
//	if err != nil {
//		return err
//	}
//
//	var matchContract = regexp.MustCompile(".+\\W(\\w+)\\Wtypes.json")
//	contract := matchContract.ReplaceAllString(path, "$1")
//
//	// write classes
//	for _, structName := range gen.keys {
//		gen.SplitComments(structName, javaTypes)
//		spaces := strings.Repeat(" ", gen.maxName+gen.maxType)
//
//		file, err := os.Create("../java/src/org/iota/wasplib/contracts/" + contract + "/" + structName + ".java")
//		if err != nil {
//			return err
//		}
//		defer file.Close()
//
//		// write file header
//		fmt.Fprintf(file, "// Copyright 2020 IOTA Stiftung\n")
//		fmt.Fprintf(file, "// SPDX-License-Identifier: Apache-2.0\n\n")
//		fmt.Fprintf(file, "package org.iota.wasplib.contracts.%s;\n\n", contract)
//		fmt.Fprintf(file, "import org.iota.wasplib.client.bytes.BytesDecoder;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.bytes.BytesEncoder;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.Hname;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScAddress;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScAgent;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScChainId;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScColor;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScContractId;\n")
//		fmt.Fprintf(file, "import org.iota.wasplib.client.hashtypes.ScHash;\n")
//
//		fmt.Fprintf(file, "\npublic class %s{\n", structName)
//
//		// write struct layout
//		fmt.Fprintf(file, "\t//@formatter:off\n")
//		types := gen.schema.Types
//		for _, fld := range types[structName] {
//			for name, _ := range fld {
//				camel := gen.camels[name]
//				comment := gen.comments[name]
//				javaType := gen.types[name]
//				if comment != "" {
//					comment = spaces[:gen.maxCamel-len(camel)] + comment
//				}
//				camel = spaces[:gen.maxType-len(javaType)] + camel
//				fmt.Fprintf(file, "\tpublic %s %s;%s\n", javaType, camel, comment)
//			}
//		}
//		fmt.Fprintf(file, "\t//@formatter:on\n")
//
//		// write encoder for struct
//		fmt.Fprintf(file, "\n\tpublic static byte[] encode(%s o){\n", structName)
//		fmt.Fprintf(file, "\t\treturn new BytesEncoder().\n")
//		for _, fld := range types[structName] {
//			for name, typeName := range fld {
//				index := strings.Index(typeName, "//")
//				if index > 0 {
//					typeName = strings.TrimSpace(typeName[:index])
//				}
//				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
//				fmt.Fprintf(file, "\t\t\t\t%s(o.%s).\n", typeName, camelcase(name))
//			}
//		}
//		fmt.Fprintf(file, "\t\t\t\tData();\n\t}\n")
//
//		// write decoder for struct
//		fmt.Fprintf(file, "\n\tpublic static %s decode(byte[] bytes) {\n", structName)
//		fmt.Fprintf(file, "\t\tBytesDecoder decode = new BytesDecoder(bytes);\n        %s data = new %s();\n", structName, structName)
//		for _, fld := range types[structName] {
//			for name, typeName := range fld {
//				index := strings.Index(typeName, "//")
//				if index > 0 {
//					typeName = strings.TrimSpace(typeName[:index])
//				}
//				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
//				fmt.Fprintf(file, "\t\tdata.%s = decode.%s();\n", camelcase(name), typeName)
//			}
//		}
//		fmt.Fprintf(file, "\t\treturn data;\n\t}\n")
//		fmt.Fprintf(file, "}\n")
//	}
//
//	return nil
//}
