package types

import (
	"encoding/json"
	"errors"
	"os"
)

type JsonType map[string]string
type JsonTypes map[string][]JsonType

type Schema struct {
	Funcs  map[string]string `json:"funcs"`
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
	Types  JsonTypes         `json:"types"`
	Vars   map[string]string `json:"vars"`
	Views  map[string]string `json:"views"`
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

func LoadCoreSchema() ([]*Schema, error) {
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
