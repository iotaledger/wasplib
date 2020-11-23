// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mr-tron/base58"
	"os"
	"sort"
	"strings"
)

type JsonDataModel struct {
	Contract  map[string]interface{} `json:"contract"`
	Account   map[string]interface{} `json:"account"`
	Request   map[string]interface{} `json:"request"`
	State     map[string]interface{} `json:"state"`
	Utility   map[string]interface{} `json:"utility"`
	Logs      map[string]interface{} `json:"logs"`
	Calls     []interface{}          `json:"calls"`
	Transfers []interface{}          `json:"transfers"`
}

type JsonFieldType struct {
	FieldName string `json:"field"`
	TypeName  string `json:"type"`
}

type JsonTest struct {
	JsonDataModel
	Name               string         `json:"name"`
	Setup              string         `json:"setup"`
	Flags              string         `json:"flags"`
	AdditionalRequests []interface{}  `json:"additionalRequests"`
	Expect             *JsonDataModel `json:"expect"`
}

type JsonTests struct {
	host   *WasmHost
	Types  map[string][]*JsonFieldType `json:"types"`
	Setups map[string]*JsonDataModel   `json:"setups"`
	Tests  []*JsonTest                 `json:"tests"`
}

func NewJsonTests(pathName string) (*JsonTests, error) {
	file, err := os.Open(pathName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonTests := &JsonTests{}
	err = json.NewDecoder(file).Decode(&jsonTests)
	if err != nil {
		return nil, errors.New("JSON error: " + err.Error())
	}
	return jsonTests, nil
}

func (t *JsonTests) ClearData() {
	t.ClearObjectData("contract", OBJTYPE_MAP)
	t.ClearObjectData("account", OBJTYPE_MAP)
	t.ClearObjectData("request", OBJTYPE_MAP)
	t.ClearObjectData("state", OBJTYPE_MAP)
	t.ClearObjectData("logs", OBJTYPE_MAP)
	t.ClearObjectData("calls", OBJTYPE_MAP_ARRAY)
	t.ClearObjectData("transfers", OBJTYPE_MAP_ARRAY)
}

func (t *JsonTests) ClearObjectData(key string, typeId int32) {
	object := t.FindSubObject(nil, key, typeId)
	object.SetInt(KeyLength, 0)
}

func (t *JsonTests) CompareArrayData(key string, array []interface{}) bool {
	arrayObject := t.FindSubObject(nil, key, OBJTYPE_MAP_ARRAY)
	if arrayObject.GetInt(KeyLength) != int64(len(array)) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	for i := range array {
		mapObject := t.FindIndexedMap(arrayObject, i)
		if !t.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
			fmt.Printf("      map %s\n", key)
			return false
		}
	}
	return true
}

func (t *JsonTests) CompareData(jsonTest *JsonTest) bool {
	expectData := jsonTest.Expect
	return t.CompareMapData("account", expectData.Account) &&
		t.CompareMapData("state", expectData.State) &&
		t.CompareMapData("logs", expectData.Logs) &&
		t.CompareArrayData("calls", expectData.Calls) &&
		t.CompareArrayData("transfers", expectData.Transfers)
}

func (t *JsonTests) CompareMapData(key string, values map[string]interface{}) bool {
	mapObject := t.FindSubObject(nil, key, OBJTYPE_MAP)
	if !t.CompareSubMapData(mapObject, values) {
		fmt.Printf("      map %s\n", key)
		return false
	}
	return true
}

func (t *JsonTests) CompareSubArrayData(mapObject HostObject, key string, array []interface{}) bool {
	if len(array) == 0 {
		return true
	}
	keyId := t.GetKeyId(key)
	if !mapObject.Exists(keyId) {
		fmt.Printf("FAIL: missing array %s\n", key)
		return false
	}
	elem := array[0]
	typeId := mapObject.GetTypeId(keyId)
	arrayObject := t.FindSubObject(mapObject, key, typeId)
	if int(arrayObject.GetInt(KeyLength)) != len(array) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	switch ty := elem.(type) {
	case string:
		if typeId != OBJTYPE_BYTES_ARRAY && typeId != OBJTYPE_STRING_ARRAY {
			fmt.Printf("FAIL: not a bytes or string array: %s\n", key)
			return false
		}
		for i, elem := range array {
			value := arrayObject.GetString(int32(i))
			expect := process(elem.(string))
			if value != expect {
				fmt.Printf("FAIL: string array %s[%d], expected '%s', got '%s'\n", key, i, expect, value)
				return false
			}
		}
		return true
	case float64:
		if typeId != OBJTYPE_INT_ARRAY {
			fmt.Printf("FAIL: not an int array: %s\n", key)
			return false
		}
		for i, elem := range array {
			value := arrayObject.GetInt(int32(i))
			expect := int64(elem.(float64))
			if value != expect {
				fmt.Printf("FAIL: int array %s[%d], expected '%d', got '%d'\n", key, i, expect, value)
				return false
			}
		}
		return true
	case map[string]interface{}:
		if typeId == OBJTYPE_MAP_ARRAY {
			for i := range array {
				mapObject := t.FindIndexedMap(arrayObject, i)
				if !t.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
					fmt.Printf("      map %s\n", key)
					return false
				}
			}
			return true
		}

		if typeId != OBJTYPE_BYTES_ARRAY {
			fmt.Printf("FAIL: not a bytes array: %s\n", key)
			return false
		}
		for i, elem := range array {
			value := arrayObject.GetString(int32(i))
			expect, ok := t.makeSerializedObject(key, elem)
			if !ok {
				return false
			}
			if value != expect {
				fmt.Printf("FAIL: string array %s[%d],\n    expected '%s',\n    got      '%s'\n", key, i, expect, value)
				decVal, _ := base58.Decode(value)
				expVal, _ := base58.Decode(expect)
				fmt.Printf("    %v\n    %v\n", decVal, expVal)
				return false
			}
		}
		return true

	default:
		panic(fmt.Sprintf("Invalid type: %T", ty))
	}
}

func (t *JsonTests) CompareSubMapData(mapObject HostObject, values map[string]interface{}) bool {
	for _, k := range SortedKeys(values) {
		field := values[k]
		key := process(k)
		keyId := t.GetKeyId(key)
		switch ty := field.(type) {
		case string:
			value := mapObject.GetString(keyId)
			expect := process(field.(string))
			if value != expect {
				fmt.Printf("FAIL: string %s, expected '%s', got '%s'\n", key, expect, value)
				return false
			}
		case float64:
			value := mapObject.GetInt(keyId)
			expect := int64(field.(float64))
			if value != expect {
				fmt.Printf("FAIL: int %s, expected %d, got %d\n", key, expect, value)
				return false
			}
		case map[string]interface{}:
			typeId := mapObject.GetTypeId(keyId)
			if typeId == OBJTYPE_MAP {
				subMapObject := t.FindSubObject(mapObject, key, OBJTYPE_MAP)
				if !t.CompareSubMapData(subMapObject, field.(map[string]interface{})) {
					fmt.Printf("      map %s\n", key)
					return false
				}
				return true
			}

			if typeId != OBJTYPE_STRING {
				fmt.Printf("FAIL: not a string field: %s\n", key)
				return false
			}

			value := mapObject.GetString(keyId)
			expect, ok := t.makeSerializedObject(key, field)
			if !ok {
				return false
			}
			if value != expect {
				fmt.Printf("FAIL: string %s,\n    expected '%s',\n    got      '%s'\n", key, expect, value)
				decVal, _ := base58.Decode(value)
				expVal, _ := base58.Decode(expect)
				fmt.Printf("    %v\n    %v\n", decVal, expVal)
				return false
			}

		case []interface{}:
			t.CompareSubArrayData(mapObject, key, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
	return true
}

func (t *JsonTests) FindIndexedMap(arrayObject HostObject, index int) HostObject {
	return t.host.FindObject(arrayObject.GetObjectId(int32(index), OBJTYPE_MAP))
}

func (t *JsonTests) FindSubObject(mapObject HostObject, key string, typeId int32) HostObject {
	if mapObject == nil {
		// use root object
		mapObject = t.host.FindObject(1)
	}
	return t.host.FindObject(mapObject.GetObjectId(t.GetKeyId(key), typeId))
}

func (t *JsonTests) GetKeyId(key string) int32 {
	keyId := t.host.GetKeyIdFromBytes([]byte(key))
	t.host.logger.Log(KeyTraceHost, fmt.Sprintf("t.GetKeyId('%s')=k%d", key, keyId))
	return keyId
}

func (t *JsonTests) LoadData(jsonData *JsonDataModel) {
	t.LoadMapData("contract", jsonData.Contract)
	t.LoadMapData("account", jsonData.Account)
	t.LoadMapData("request", jsonData.Request)
	t.LoadMapData("state", jsonData.State)
	t.LoadMapData("utility", jsonData.Utility)
}

func (t *JsonTests) LoadMapData(key string, values map[string]interface{}) {
	mapObject := t.FindSubObject(nil, key, OBJTYPE_MAP)
	t.LoadSubMapData(mapObject, values)
}

func (t *JsonTests) LoadSubArrayData(arrayObject HostObject, values []interface{}) {
	for key, field := range values {
		switch ty := field.(type) {
		case string:
			arrayObject.SetString(int32(key), process(field.(string)))
		//case float64:
		//	mapObject.SetInt(t.GetKeyId(key), int64(field.(float64)))
		//case map[string]interface{}:
		//	subMapObject := t.FindSubObject(mapObject, key, OBJTYPE_MAP)
		//	t.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		//case []interface{}:
		//	subMapObject := t.FindSubObject(mapObject, key, OBJTYPE_STRING_ARRAY)
		//	t.LoadSubArrayData(subMapObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
}

func (t *JsonTests) LoadSubMapData(mapObject HostObject, values map[string]interface{}) {
	for _, k := range SortedKeys(values) {
		field := values[k]
		key := process(k)
		switch ty := field.(type) {
		case string:
			mapObject.SetString(t.GetKeyId(key), process(field.(string)))
		case float64:
			mapObject.SetInt(t.GetKeyId(key), int64(field.(float64)))
		case map[string]interface{}:
			subMapObject := t.FindSubObject(mapObject, key, OBJTYPE_MAP)
			t.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		case []interface{}:
			subArrayObject := t.FindSubObject(mapObject, key, OBJTYPE_STRING_ARRAY)
			t.LoadSubArrayData(subArrayObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
}

func (t *JsonTests) makeSerializedObject(key string, field interface{}) (string, bool) {
	object := field.(map[string]interface{})
	if len(object) != 1 {
		fmt.Printf("FAIL: bytes array %s: object type not found\n", key)
	}
	encoder := NewBytesEncoder()
	// only 1 object
	for typeName, value := range object {
		if !t.makeSubObject(encoder, key, typeName, value) {
			return "", false
		}
	}
	return base58.Encode(encoder.Data()), true
}

func (t *JsonTests) makeSubObject(encoder *BytesEncoder, key string, typeName string, value interface{}) bool {
	fieldDefs, ok := t.Types[typeName]
	if !ok {
		fmt.Printf("FAIL: bytes array %s: object typedef for %s missing\n", key, typeName)
		return false
	}
	fieldValues := value.(map[string]interface{})
	if len(fieldValues) != len(fieldDefs) {
		fmt.Printf("FAIL: bytes array %s: object typedef for %s mismatch\n", key, typeName)
		return false
	}
	for _, def := range fieldDefs {
		value = fieldValues[def.FieldName]
		typeName = def.TypeName
		switch typeName {
		case "Address", "Agent", "Bytes", "Color", "RequestId", "TxHash":
			bytes, _ := base58.Decode(process(value.(string)))
			encoder.Bytes(bytes)
		case "Int":
			encoder.Int(int64(value.(float64)))
		case "String":
			encoder.String(value.(string))
		default:
			_, ok = t.Types[typeName]
			if ok {
				enc := NewBytesEncoder()
				if !t.makeSubObject(enc, key, typeName, value) {
					return false
				}
				encoder.Bytes(enc.Data())
				return true
			}
			if typeName[:2] == "[]" {
				typeName = typeName[2:]
				array := value.([]interface{})
				encoder.Int(int64(len(array)))
				for _, value = range array {
					enc := NewBytesEncoder()
					if !t.makeSubObject(enc, key, typeName, value) {
						return false
					}
					encoder.Bytes(enc.Data())
				}
				return true
			}
			panic("Unhandled type '" + typeName + "' of field in" + key)
		}
	}
	return true
}

func process(value string) string {
	if len(value) == 0 {
		return value
	}
	// preprocesses keys and values by replacing special named values
	size := 32
	switch value[0] {
	case '#': // 32-byte hash value
		if value == "#iota" {
			return processHash("", 32)
		}
	case '@': // 37-byte agent
		size = 37
	case '$': // 34-byte request id
		size = 34
	default:
		return value
	}
	return processHash(value[1:], size)
}

func processHash(value string, size int) string {
	hash := make([]byte, size)
	copy(hash, value)
	return base58.Encode(hash)
}

func (t *JsonTests) RunTest(host *WasmHost, test *JsonTest) bool {
	t.host = host
	fmt.Printf("Test: %s\n", test.Name)
	if test.Expect == nil {
		fmt.Printf("FAIL: Missing expect model data\n")
		return false
	}
	t.ClearData()
	if test.Setup != "" {
		setupData, ok := t.Setups[test.Setup]
		if !ok {
			fmt.Printf("FAIL: Missing setup: %s\n", test.Setup)
			return false
		}
		t.LoadData(setupData)
	}
	t.LoadData(&test.JsonDataModel)
	request := t.FindSubObject(nil, "request", OBJTYPE_MAP)
	if !t.runRequest(request, test.Request) {
		return false
	}
	for _, req := range test.AdditionalRequests {
		jsonRequest := req.(map[string]interface{})
		request.SetInt(KeyLength, 0)
		t.LoadSubMapData(request, jsonRequest)
		if !t.runRequest(request, jsonRequest) {
			return false
		}
	}

	scId := t.FindSubObject(nil, "contract", OBJTYPE_MAP).GetString(t.GetKeyId("id"))
	reqParams := t.FindSubObject(request, "params", OBJTYPE_MAP)
	calls := t.FindSubObject(nil, "calls", OBJTYPE_MAP_ARRAY)

	expectedCalls := len(test.Expect.Calls)
	for i := 0; i < expectedCalls && i < int(calls.GetInt(KeyLength)); i++ {
		call := t.FindIndexedMap(calls, i)
		delay := call.GetInt(t.GetKeyId("delay"))
		if delay != 0 && !strings.Contains(test.Flags, "nodelay") {
			// only process calls when they have no delay
			// unless overridden by the nodelay flag
			// those are the only ones that will be incorporated in the final state
			continue
		}

		contract := call.GetString(t.GetKeyId("contract"))
		if contract != "" && contract != scId {
			// only process calls when they are for the current contract
			// those are the only ones that will be incorporated in the final state
			continue
		}

		function := call.GetString(t.GetKeyId("function"))
		request.SetString(t.GetKeyId("sender"), scId)
		request.SetString(t.GetKeyId("function"), function)
		reqParams.SetInt(KeyLength, 0)
		params := t.FindSubObject(call, "params", OBJTYPE_MAP)
		//TODO how to iterate
		params.(*HostMap).CopyDataTo(reqParams)
		fmt.Printf("    Run function: %v\n", function)
		err := t.host.RunScFunction(function)
		if err != nil {
			fmt.Printf("FAIL: Request function %s: %v\n", function, err)
			return false
		}
	}

	// now compare the expected json data model to the actual host data model
	return t.CompareData(test)
}

func (t *JsonTests) runRequest(request HostObject, req map[string]interface{}) bool {
	function, ok := req["function"]
	if !ok {
		fmt.Printf("FAIL: Missing request.function\n")
		return false
	}
	if request.Exists(t.GetKeyId("balance")) {
		reqColors := t.FindSubObject(request, "colors", OBJTYPE_STRING_ARRAY)
		reqBalance := t.FindSubObject(request, "balance", OBJTYPE_MAP)
		account := t.FindSubObject(nil, "account", OBJTYPE_MAP)
		accBalance := t.FindSubObject(account, "balance", OBJTYPE_MAP)
		nrOfColors := int32(reqColors.GetInt(KeyLength))
		for i := nrOfColors - 1; i >= 0; i-- {
			color := reqColors.GetBytes(i)
			colorKeyId := t.GetKeyId(base58.Encode(color))
			accBalance.SetInt(colorKeyId, accBalance.GetInt(colorKeyId)+reqBalance.GetInt(colorKeyId))
		}
	}

	fmt.Printf("    Run function: %v\n", function)
	err := t.host.RunScFunction(function.(string))
	if err != nil {
		fmt.Printf("FAIL: Function %v: %v\n", function, err)
		return false
	}
	return true
}

func SortedKeys(values map[string]interface{}) []string {
	keys := make([]string, len(values))
	index := 0
	for key := range values {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	return keys
}
