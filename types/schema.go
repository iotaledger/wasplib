package types

import (
	"encoding/json"
	"errors"
	"os"
)

type StringMap map[string]string
type StringMapMap map[string]StringMap

type Schema struct {
	Funcs StringMapMap `json:"funcs"`
	Name  string       `json:"name"`
	Types StringMapMap `json:"types"`
	Vars  StringMap    `json:"vars"`
	Views StringMapMap `json:"views"`
}

func LoadSchema(path string) (*Schema, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	schema := &Schema{}
	err = json.NewDecoder(file).Decode(schema)
	if err != nil {
		return nil, errors.New("JSON error: " + err.Error())
	}
	return schema, nil
}

func LoadCoreSchemas() ([]*Schema, error) {
	file, err := os.Open("corecontracts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	core := make([]*Schema, 0)
	err = json.NewDecoder(file).Decode(&core)
	if err != nil {
		return nil, errors.New("JSON error: " + err.Error())
	}
	return core, nil
}
