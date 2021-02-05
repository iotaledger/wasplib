package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/iotaledger/wasplib/schema/generator"
	"os"
)

var core = flag.Bool("core", false, "generate core contract schemas")

func main() {
	flag.Parse()
	if *core {
		err := generateCoreContractsSchema()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	err := generateSchema()
	if err != nil {
		fmt.Println(err)
	}
}

func generateCoreContractsSchema() error {
	coreSchemas, err := loadCoreSchemas()
	if err != nil {
		return err
	}
	err = generator.GenerateGoCoreContractsSchema(coreSchemas)
	if err != nil {
		return err
	}
	err = generator.GenerateRustCoreContractsSchema(coreSchemas)
	if err != nil {
		return err
	}
	return nil
}

func generateSchema() error {
	schema, err := loadSchema()
	if err != nil {
		return err
	}
	err = schema.GenerateGoSchema()
	if err != nil {
		return err
	}
	err = schema.GenerateRustSchema()
	if err != nil {
		return err
	}
	err = schema.GenerateGoTypes()
	if err != nil {
		return err
	}
	err = schema.GenerateRustTypes()
	if err != nil {
		return err
	}
	return nil
}

func loadSchema() (*generator.Schema, error) {
	fmt.Println("loading schema.json")
	file, err := os.Open("schema.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonSchema := &generator.JsonSchema{}
	err = json.NewDecoder(file).Decode(jsonSchema)
	if err != nil {
		return nil, err
	}

	schema := generator.NewSchema()
	err = schema.Compile(jsonSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func loadCoreSchemas() ([]*generator.Schema, error) {
	fmt.Println("loading corecontracts.json")
	file, err := os.Open("corecontracts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	coreJsonSchemas := make([]*generator.JsonSchema, 0)
	err = json.NewDecoder(file).Decode(&coreJsonSchemas)
	if err != nil {
		return nil, err
	}

	coreSchemas := make([]*generator.Schema, 0)
	for _, jsonSchema := range coreJsonSchemas {
		schema := generator.NewSchema()
		err = schema.Compile(jsonSchema)
		if err != nil {
			return nil, err
		}
		coreSchemas = append(coreSchemas, schema)
	}
	return coreSchemas, nil
}
