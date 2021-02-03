// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
	"github.com/mr-tron/base58"
	"os"
	"sort"
)

type JsonDataModel struct {
	Contract  map[string]interface{} `json:"contract"`
	Balances  map[string]interface{} `json:"balances"`
	Timestamp int64                  `json:"timestamp"`
	Caller    string                 `json:"caller"`
	Chain     string                 `json:"chain"`
	Creator   string                 `json:"creator"`
	Function  string                 `json:"function"`
	Id        string                 `json:"id"`
	Incoming  map[string]interface{} `json:"incoming"`
	Params    map[string]interface{} `json:"params"`
	State     map[string]interface{} `json:"state"`
	Logs      map[string]interface{} `json:"logs"`
	Results   map[string]interface{} `json:"results"`
	Calls     []interface{}          `json:"calls"`
	Posts     []interface{}          `json:"posts"`
	Views     []interface{}          `json:"views"`
	Transfers []interface{}          `json:"transfers"`
	Utility   map[string]interface{} `json:"utility"`
}

type JsonFieldType struct {
	FieldName string `json:"field"`
	TypeName  string `json:"type"`
}

type JsonTest struct {
	JsonDataModel
	Name               string           `json:"name"`
	Setup              string           `json:"setup"`
	Flags              string           `json:"flags"`
	AdditionalRequests []*JsonDataModel `json:"additionalRequests"`
	Expect             *JsonDataModel   `json:"expect"`
}

type JsonTests struct {
	host   *SimpleWasmHost
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
	t.ClearObjectData(wasmhost.KeyBalances, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyIncoming, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyParams, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyState, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyLogs, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyResults, client.TYPE_MAP)
	t.ClearObjectData(wasmhost.KeyMaps, client.TYPE_MAP|client.TYPE_ARRAY)
	t.ClearObjectData(wasmhost.KeyTransfers, client.TYPE_MAP|client.TYPE_ARRAY)
}

func (t *JsonTests) ClearObjectData(keyId int32, typeId int32) {
	object := t.FindSubObject(nil, keyId, typeId)
	object.SetBytes(wasmhost.KeyLength, client.TYPE_INT, IntToBytes(0))
}

func (t *JsonTests) CompareArrayData(keyId int32, array []interface{}) bool {
	arrayObject := t.FindSubObject(nil, keyId, client.TYPE_MAP|client.TYPE_ARRAY)
	length := BytesToInt(arrayObject.GetBytes(wasmhost.KeyLength, client.TYPE_INT))
	if length != int64(len(array)) {
		key := t.host.GetKeyStringFromId(keyId)
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	for i := range array {
		mapObject := t.FindIndexedMap(arrayObject, i)
		if !t.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
			return false
		}
	}
	return true
}

func (t *JsonTests) CompareData(jsonTest *JsonTest) bool {
	expectData := jsonTest.Expect
	return t.CompareMapData(wasmhost.KeyBalances, expectData.Balances) &&
		t.CompareMapData(wasmhost.KeyState, expectData.State) &&
		t.CompareMapData(wasmhost.KeyLogs, expectData.Logs) &&
		t.CompareMapData(wasmhost.KeyResults, expectData.Results) &&
		t.CompareArrayData(wasmhost.KeyTransfers, expectData.Transfers)
}

func (t *JsonTests) CompareMapData(keyId int32, values map[string]interface{}) bool {
	mapObject := t.FindSubObject(nil, keyId, client.TYPE_MAP)
	return t.CompareSubMapData(mapObject, values)
}

func (t *JsonTests) CompareSubArrayData(mapObject VmObject, keyId int32, array []interface{}) bool {
	if len(array) == 0 {
		return true
	}
	if !mapObject.Exists(keyId, 0) {
		key := t.host.GetKeyStringFromId(keyId)
		return mapObject.Fail("missing array %s", key)
	}
	elem := array[0]
	typeId := mapObject.GetTypeId(keyId)
	arrayObject := t.FindSubObject(mapObject, keyId, typeId)
	if (typeId & client.TYPE_ARRAY) == 0 {
		return arrayObject.Fail("not an array")
	}
	length := BytesToInt(arrayObject.GetBytes(wasmhost.KeyLength, client.TYPE_INT))
	if length != int64(len(array)) {
		return arrayObject.Fail("length mismatch")
	}
	typeId &= ^client.TYPE_ARRAY
	switch ty := elem.(type) {
	case string:
		if typeId != client.TYPE_ADDRESS &&
			typeId != client.TYPE_AGENT_ID &&
			typeId != client.TYPE_BYTES &&
			typeId != client.TYPE_CHAIN_ID &&
			typeId != client.TYPE_COLOR &&
			typeId != client.TYPE_CONTRACT_ID &&
			typeId != client.TYPE_STRING {
			return arrayObject.Fail("not a bytes or string array")
		}
		for i, elem := range array {
			value := arrayObject.GetBytes(int32(i), typeId)
			expect := process(elem.(string))
			if string(value) != expect {
				return arrayObject.Fail("[%d]:\n    expected '%s'\n    got      '%s'", i, expect, value)
			}
		}
		return true
	case float64:
		if typeId != client.TYPE_INT {
			return arrayObject.Fail("not an int array")
		}
		for i, elem := range array {
			value := BytesToInt(arrayObject.GetBytes(int32(i), typeId))
			expect := int64(elem.(float64))
			if value != expect {
				return arrayObject.Fail("[%d]: expected '%d', got '%d'", i, expect, value)
			}
		}
		return true
	case map[string]interface{}:
		if typeId == client.TYPE_MAP {
			for i := range array {
				mapObject := t.FindIndexedMap(arrayObject, i)
				if !t.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
					return false
				}
			}
			return true
		}

		if typeId != client.TYPE_BYTES {
			return arrayObject.Fail("not a bytes array")
		}
		for i, elem := range array {
			value := string(arrayObject.GetBytes(int32(i), typeId))
			expect, ok := t.makeSerializedObject(keyId, elem)
			if !ok {
				return false
			}
			if value != expect {
				arrayObject.Fail("[%d]:\n    expected '%s'\n    got      '%s'", i, expect, value)
				expVal, _ := base58.Decode(expect)
				gotVal, _ := base58.Decode(value)
				fmt.Printf("    %v\n    %v\n", expVal, gotVal)
				return false
			}
		}
		return true

	default:
		panic(fmt.Sprintf("Invalid type: %T", ty))
	}
}

func (t *JsonTests) CompareSubMapData(mapObject VmObject, values map[string]interface{}) bool {
	for _, key := range SortedKeys(values) {
		field := values[key]
		keyId := t.GetKeyId(key)
		switch ty := field.(type) {
		case string:
			value := string(mapObject.GetBytes(keyId, client.TYPE_STRING))
			expect := process(field.(string))
			if value != expect {
				return mapObject.Fail("%s: expected '%s', got '%s'", key, expect, value)
			}
		case float64:
			value := BytesToInt(mapObject.GetBytes(keyId, client.TYPE_INT))
			expect := int64(field.(float64))
			if value != expect {
				return mapObject.Fail("%s: expected %d, got %d", key, expect, value)
			}
		case map[string]interface{}:
			typeId := mapObject.GetTypeId(keyId)
			if typeId == client.TYPE_MAP {
				subMapObject := t.FindSubObject(mapObject, keyId, client.TYPE_MAP)
				return t.CompareSubMapData(subMapObject, field.(map[string]interface{}))
			}

			if typeId != client.TYPE_STRING {
				return mapObject.Fail("%s: not a string field", key)
			}

			value := string(mapObject.GetBytes(keyId, typeId))
			expect, ok := t.makeSerializedObject(keyId, field)
			if !ok {
				return false
			}
			if value != expect {
				mapObject.Fail("%s:\n    expected '%s'\n    got      '%s'", key, expect, value)
				expVal, _ := base58.Decode(expect)
				gotVal, _ := base58.Decode(value)
				fmt.Printf("    %v\n    %v\n", expVal, gotVal)
				return false
			}

		case []interface{}:
			return t.CompareSubArrayData(mapObject, keyId, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
	return true
}

func (t *JsonTests) Dump(test *JsonTest) {
	contractName := t.Setups["default"].Contract["name"].(string)
	folder := "dump/" + contractName
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(folder + "/" + test.Name + ".json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t.FindObject(1).(*HostMap).Dump(f)
}

func (t *JsonTests) FindIndexedMap(arrayObject VmObject, index int) VmObject {
	return t.FindObject(arrayObject.GetObjectId(int32(index), client.TYPE_MAP))
}

func (t *JsonTests) FindObject(objId int32) VmObject {
	return t.host.FindObject(objId).(VmObject)
}

func (t *JsonTests) FindSubObject(mapObject VmObject, keyId int32, typeId int32) VmObject {
	if mapObject == nil {
		// use root object
		mapObject = t.FindObject(1)
	}
	return t.FindObject(mapObject.GetObjectId(keyId, typeId))
}

func (t *JsonTests) GetKeyId(key string) int32 {
	keyValue := process(key)
	if keyValue != key {
		bytes, _ := base58.Decode(keyValue)
		return t.host.GetKeyIdFromBytes(bytes)
	}
	return t.host.GetKeyIdFromString(key)
}

func (t *JsonTests) LoadData(jsonData *JsonDataModel) {
	t.LoadMapData(wasmhost.KeyBalances, jsonData.Balances)
	t.LoadMapData(wasmhost.KeyIncoming, jsonData.Incoming)
	t.LoadMapData(wasmhost.KeyParams, jsonData.Params)
	t.LoadMapData(wasmhost.KeyState, jsonData.State)
	root := t.FindObject(1)
	if jsonData.Timestamp != 0 {
		root.SetBytes(wasmhost.KeyTimestamp, client.TYPE_INT, IntToBytes(jsonData.Timestamp))
	}
	if jsonData.Caller != "" {
		root.SetBytes(wasmhost.KeyCaller, client.TYPE_AGENT_ID, []byte(process(jsonData.Caller)))
	}
	//if jsonData.Chain != "" {
	//	root.SetBytes(wasmhost.KeyChain, client.TYPE_CHAIN_ID, []byte(process(jsonData.Chain)))
	//}
	if jsonData.Creator != "" {
		root.SetBytes(wasmhost.KeyContractCreator, client.TYPE_AGENT_ID, []byte(process(jsonData.Creator)))
	}
	if jsonData.Id != "" {
		root.SetBytes(wasmhost.KeyContractId, client.TYPE_CONTRACT_ID, []byte(process(jsonData.Id)))
	}
}

func (t *JsonTests) LoadMapData(keyId int32, values map[string]interface{}) {
	mapObject := t.FindSubObject(nil, keyId, client.TYPE_MAP)
	t.LoadSubMapData(mapObject, values)
}

func (t *JsonTests) LoadSubArrayData(arrayObject VmObject, values []interface{}) {
	for key, field := range values {
		switch ty := field.(type) {
		case string:
			arrayObject.SetBytes(int32(key), client.TYPE_STRING, []byte(process(field.(string))))
		//case float64:
		//	mapObject.SetInt(t.GetKeyId(key), int64(field.(float64)))
		//case map[string]interface{}:
		//	subMapObject := t.FindSubObject(mapObject, key, client.TYPE_MAP)
		//	t.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		//case []interface{}:
		//	subMapObject := t.FindSubObject(mapObject, key, client.TYPE_STRING_ARRAY)
		//	t.LoadSubArrayData(subMapObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
}

func (t *JsonTests) LoadSubMapData(mapObject VmObject, values map[string]interface{}) {
	for _, key := range SortedKeys(values) {
		field := values[key]
		keyId := t.GetKeyId(key)
		switch ty := field.(type) {
		case string:
			mapObject.SetBytes(keyId, client.TYPE_STRING, []byte(process(field.(string))))
		case float64:
			mapObject.SetBytes(keyId, client.TYPE_INT, IntToBytes(int64(field.(float64))))
		case map[string]interface{}:
			subMapObject := t.FindSubObject(mapObject, keyId, client.TYPE_MAP)
			t.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		case []interface{}:
			subArrayObject := t.FindSubObject(mapObject, keyId, client.TYPE_STRING|client.TYPE_ARRAY)
			t.LoadSubArrayData(subArrayObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", ty))
		}
	}
}

func (t *JsonTests) makeSerializedObject(keyId int32, field interface{}) (string, bool) {
	object := field.(map[string]interface{})
	if len(object) != 1 {
		key := t.host.GetKeyStringFromId(keyId)
		fmt.Printf("FAIL: bytes array %s: object type not found\n", key)
	}
	encoder := NewBytesEncoder()
	// only 1 object
	for typeName, value := range object {
		if !t.makeSubObject(encoder, keyId, typeName, value) {
			return "", false
		}
	}
	return base58.Encode(encoder.Data()), true
}

func (t *JsonTests) makeSubObject(encoder *BytesEncoder, keyId int32, typeName string, value interface{}) bool {
	fieldDefs, ok := t.Types[typeName]
	if !ok {
		key := t.host.GetKeyStringFromId(keyId)
		fmt.Printf("FAIL: bytes array %s: object typedef for %s missing\n", key, typeName)
		return false
	}
	fieldValues := value.(map[string]interface{})
	if len(fieldValues) != len(fieldDefs) {
		key := t.host.GetKeyStringFromId(keyId)
		fmt.Printf("FAIL: bytes array %s: object typedef for %s mismatch\n", key, typeName)
		return false
	}
	for _, def := range fieldDefs {
		value = fieldValues[def.FieldName]
		typeName = def.TypeName
		switch typeName {
		case "Address", "Agent", "Bytes", "Color":
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
				if !t.makeSubObject(enc, keyId, typeName, value) {
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
					if !t.makeSubObject(enc, keyId, typeName, value) {
						return false
					}
					encoder.Bytes(enc.Data())
				}
				return true
			}
			key := t.host.GetKeyStringFromId(keyId)
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
	case '@': // 37-byte agent
		size = 37
	case '$': // 33-byte address
		size = 33
	case '#': // 32-byte hash value
		if value == "#iota" {
			return base58.Encode(client.IOTA.Bytes())
		}
		if value == "#mint" {
			return base58.Encode(client.MINT.Bytes())
		}
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

func (t *JsonTests) runRequest(function string) (success bool) {
	incoming := t.FindSubObject(nil, wasmhost.KeyIncoming, client.TYPE_MAP).(*HostMap)
	balances := t.FindSubObject(nil, wasmhost.KeyBalances, client.TYPE_MAP).(*HostMap)
	mintKeyId := t.GetKeyId("#mint")
	for keyId := range incoming.fields {
		if keyId != mintKeyId {
			balances.SetInt(keyId, balances.GetInt(keyId)+incoming.GetInt(keyId))
		}
	}

	defer func() {
		if err := recover(); err != nil {
			// deliberate panic call by SC??
			success = t.host.panicked
			t.host.panicked = false
			if !success {
				fmt.Printf("FAIL: Function %s: %v\n", function, err)
			}
		}
	}()

	fmt.Printf("    Run function: %s\n", function)
	err := t.host.RunScFunction(function)
	if err != nil {
		fmt.Printf("FAIL: Function %s: %v\n", function, err)
		return false
	}
	return true
}

func (t *JsonTests) RunTest(host *SimpleWasmHost, test *JsonTest) bool {
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
	if !t.runRequest(test.Function) {
		return false
	}
	incoming := t.FindSubObject(nil, wasmhost.KeyIncoming, client.TYPE_MAP)
	params := t.FindSubObject(nil, wasmhost.KeyParams, client.TYPE_MAP)
	zero := IntToBytes(0)
	for _, jsonRequest := range test.AdditionalRequests {
		incoming.SetBytes(wasmhost.KeyLength, client.TYPE_INT, zero)
		params.SetBytes(wasmhost.KeyLength, client.TYPE_INT, zero)
		t.LoadData(jsonRequest)
		if !t.runRequest(jsonRequest.Function) {
			return false
		}
	}

	//root := t.FindObject(1)
	//scId := root.GetBytes(wasmhost.KeyId, client.TYPE_CONTRACT_ID)
	//calls := t.FindSubObject(nil, wasmhost.KeyCalls, client.TYPE_MAP|client.TYPE_ARRAY)
	//
	//expectedCalls := len(test.Expect.Calls)
	//for i := 0; i < expectedCalls && i < int(BytesToInt(calls.GetBytes(wasmhost.KeyLength, client.TYPE_INT))); i++ {
	//	post := t.FindIndexedMap(calls, i)
	//	delay := BytesToInt(post.GetBytes(wasmhost.KeyDelay, client.TYPE_INT))
	//	if delay != 0 && !strings.Contains(test.Flags, "nodelay") {
	//		// only process posts when they have no delay
	//		// unless overridden by the nodelay flag
	//		// those are the only ones that will be incorporated in the final state
	//		continue
	//	}
	//
	//	contract := post.GetBytes(wasmhost.KeyContract,  client.TYPE_CONTRACT_ID)
	//	if string(contract) != string(scId) {
	//		// only process posts when they are for the current contract
	//		// those are the only ones that will be incorporated in the final state
	//		continue
	//	}
	//
	//	root.SetBytes(wasmhost.KeyCaller, client.TYPE_AGENT_ID, scId)
	//	//TODO increment timestamp and pass post.transfers as incoming
	//	//TODO how do we pass incoming when we call instead of post?
	//	params.SetBytes(wasmhost.KeyLength, client.TYPE_INT, zero)
	//	postParams := t.FindSubObject(post, wasmhost.KeyParams, client.TYPE_MAP)
	//	//TODO how to iterate
	//	postParams.(*HostMap).CopyDataTo(params)
	//	function := string(post.GetBytes(wasmhost.KeyFunction, client.TYPE_STRING))
	//	fmt.Printf("    Run function: %s\n", function)
	//	err := t.host.RunScFunction(function)
	//	if err != nil {
	//		fmt.Printf("FAIL: Request function %s: %v\n", function, err)
	//		// dump even when failing so that we can examine why it failed
	//		t.Dump(test)
	//		return false
	//	}
	//}

	t.Dump(test)

	// now compare the expected json data model to the actual host data model
	return t.CompareData(test)
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
