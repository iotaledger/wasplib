package types

import (
	"encoding/json"
	"errors"
	"os"
)

type JsonType map[string]string
type JsonTypes map[string][]JsonType

func LoadTypes(path string) (JsonTypes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonTypes := make(JsonTypes)
	err = json.NewDecoder(file).Decode(&jsonTypes)
	if err != nil {
		return nil, errors.New("JSON error: " + err.Error())
	}
	return jsonTypes, nil
}

