package wasmhost

import (
	"fmt"
	"github.com/mr-tron/base58"
	"sort"
	"strings"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	WasmHost
	ExportsId   int32
	TransfersId int32
	JsonTests   *JsonTests
}

func NewSimpleWasmHost() (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	host.useBase58Keys = true
	err := host.Init(NewNullObject(host), NewHostMap(host, 0), nil, host)
	if err != nil {
		return nil, err
	}
	host.ExportsId = host.GetKeyId("exports")
	host.TransfersId = host.GetKeyId("transfers")
	return host, nil
}

func (host *SimpleWasmHost) ClearData() {
	host.ClearObjectData(OBJTYPE_MAP, "contract")
	host.ClearObjectData(OBJTYPE_MAP, "account")
	host.ClearObjectData(OBJTYPE_MAP, "request")
	host.ClearObjectData(OBJTYPE_MAP, "state")
	host.ClearObjectData(OBJTYPE_MAP, "logs")
	host.ClearObjectData(OBJTYPE_MAP_ARRAY, "postedRequests")
	host.ClearObjectData(OBJTYPE_MAP_ARRAY, "transfers")
}

func (host *SimpleWasmHost) ClearObjectData(typeId int32, key string) {
	object := host.FindSubObject(nil, key, typeId)
	object.SetInt(KeyLength, 0)
}

func (host *SimpleWasmHost) CompareArrayData(key string, array []interface{}) bool {
	arrayObject := host.FindSubObject(nil, key, OBJTYPE_MAP_ARRAY)
	if arrayObject.GetInt(KeyLength) != int64(len(array)) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	for i := range array {
		mapObject := host.FindObject(arrayObject.GetObjectId(int32(i), OBJTYPE_MAP))
		if !host.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
			fmt.Printf("      map %s\n", key)
			return false
		}
	}
	return true
}

func (host *SimpleWasmHost) CompareData(jsonTest *JsonTest) bool {
	expectData := jsonTest.Expect
	return host.CompareMapData("account", expectData.Account) &&
		host.CompareMapData("state", expectData.State) &&
		host.CompareMapData("logs", expectData.Logs) &&
		host.CompareArrayData("postedRequests", expectData.PostedRequests) &&
		host.CompareArrayData("transfers", expectData.Transfers)
}

func (host *SimpleWasmHost) CompareMapData(key string, values map[string]interface{}) bool {
	mapObject := host.FindSubObject(nil, key, OBJTYPE_MAP)
	if !host.CompareSubMapData(mapObject, values) {
		fmt.Printf("      map %s\n", key)
		return false
	}
	return true
}

func (host *SimpleWasmHost) CompareSubArrayData(mapObject HostObject, key string, array []interface{}) bool {
	if len(array) == 0 {
		return true
	}
	keyId := host.GetKeyId(key)
	if !mapObject.Exists(keyId) {
		fmt.Printf("FAIL: missing array %s\n", key)
		return false
	}
	elem := array[0]
	typeId := mapObject.(*HostMap).GetTypeId(keyId)
	arrayObject := host.FindSubObject(mapObject, key, typeId)
	if arrayObject.(*HostArray).GetLength() != int32(len(array)) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	switch t := elem.(type) {
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
				mapObject := host.FindObject(arrayObject.GetObjectId(int32(i), OBJTYPE_MAP))
				if !host.CompareSubMapData(mapObject, array[i].(map[string]interface{})) {
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
			expect, ok := host.makeSerializedObject(key, elem)
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
		panic(fmt.Sprintf("Invalid type: %T", t))
	}
}

func (host *SimpleWasmHost) CompareSubMapData(mapObject HostObject, values map[string]interface{}) bool {
	for _, k := range SortedKeys(values) {
		field := values[k]
		key := process(k)
		keyId := host.GetKeyId(key)
		switch t := field.(type) {
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
			typeId := mapObject.(*HostMap).GetTypeId(keyId)
			if typeId == OBJTYPE_MAP {
				subMapObject := host.FindSubObject(mapObject, key, OBJTYPE_MAP)
				if !host.CompareSubMapData(subMapObject, field.(map[string]interface{})) {
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
			expect, ok := host.makeSerializedObject(key, field)
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
			host.CompareSubArrayData(mapObject, key, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", t))
		}
	}
	return true
}

func (host *SimpleWasmHost) FindSubObject(obj HostObject, key string, typeId int32) HostObject {
	if obj == nil {
		// use root object
		obj = host.FindObject(1)
	}
	return host.FindObject(obj.GetObjectId(host.GetKeyId(key), typeId))
}

func (host *SimpleWasmHost) GetKeyId(key string) int32 {
	keyId := host.WasmHost.getKeyId([]byte(key))
	host.Trace("GetKeyId('%s')=k%d", key, keyId)
	return keyId
}

func (host *SimpleWasmHost) LoadData(jsonData *JsonDataModel) {
	host.LoadMapData("contract", jsonData.Contract)
	host.LoadMapData("account", jsonData.Account)
	host.LoadMapData("request", jsonData.Request)
	host.LoadMapData("state", jsonData.State)
	host.LoadMapData("utility", jsonData.Utility)
}

func (host *SimpleWasmHost) LoadMapData(key string, values map[string]interface{}) {
	mapObject := host.FindSubObject(nil, key, OBJTYPE_MAP)
	host.LoadSubMapData(mapObject, values)
}

func (host *SimpleWasmHost) LoadSubArrayData(arrayObject HostObject, values []interface{}) {
	for key, field := range values {
		switch t := field.(type) {
		case string:
			arrayObject.SetString(int32(key), process(field.(string)))
		//case float64:
		//	mapObject.SetInt(host.GetKeyId(key), int64(field.(float64)))
		//case map[string]interface{}:
		//	subMapObject := host.FindSubObject(mapObject, key, OBJTYPE_MAP)
		//	host.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		//case []interface{}:
		//	subMapObject := host.FindSubObject(mapObject, key, OBJTYPE_STRING_ARRAY)
		//	host.LoadSubArrayData(subMapObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", t))
		}
	}
}

func (host *SimpleWasmHost) LoadSubMapData(mapObject HostObject, values map[string]interface{}) {
	for _, k := range SortedKeys(values) {
		field := values[k]
		key := process(k)
		switch t := field.(type) {
		case string:
			mapObject.SetString(host.GetKeyId(key), process(field.(string)))
		case float64:
			mapObject.SetInt(host.GetKeyId(key), int64(field.(float64)))
		case map[string]interface{}:
			subMapObject := host.FindSubObject(mapObject, key, OBJTYPE_MAP)
			host.LoadSubMapData(subMapObject, field.(map[string]interface{}))
		case []interface{}:
			subArrayObject := host.FindSubObject(mapObject, key, OBJTYPE_STRING_ARRAY)
			host.LoadSubArrayData(subArrayObject, field.([]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", t))
		}
	}
}

func (host *WasmHost) Log(logLevel int32, text string) {
	switch logLevel {
	case KeyTraceHost:
		//fmt.Println(text)
	case KeyTrace:
		//fmt.Println(text)
	case KeyLog:
		fmt.Println(text)
	case KeyWarning:
		fmt.Println(text)
	case KeyError:
		fmt.Println(text)
	}
}

func (host *SimpleWasmHost) makeSerializedObject(key string, field interface{}) (string, bool) {
	object := field.(map[string]interface{})
	if len(object) != 1 {
		fmt.Printf("FAIL: bytes array %s: object type not found\n", key)
	}
	encoder := NewBytesEncoder()
	// only 1 object
	for typeName, value := range object {
		if !host.makeSubObject(encoder, key, typeName, value) {
			return "", false
		}
	}
	return base58.Encode(encoder.Data()), true
}

func (host *SimpleWasmHost) makeSubObject(encoder *BytesEncoder, key string, typeName string, value interface{}) bool {
	fieldDefs, ok := host.JsonTests.Types[typeName]
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
		case "Address", "Bytes", "Color", "RequestId", "TxHash":
			bytes, _ := base58.Decode(process(value.(string)))
			encoder.Bytes(bytes)
		case "Int":
			encoder.Int(int64(value.(float64)))
		case "String":
			encoder.String(value.(string))
		default:
			_, ok = host.JsonTests.Types[typeName]
			if ok {
				enc := NewBytesEncoder()
				if !host.makeSubObject(enc, key, typeName, value) {
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
					if !host.makeSubObject(enc, key, typeName, value) {
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
	case '@': // 33-byte address
		size = 33
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

func (host *SimpleWasmHost) RunTest(test *JsonTest) bool {
	fmt.Printf("Test: %s\n", test.Name)
	if test.Expect == nil {
		fmt.Printf("FAIL: Missing expect model data\n")
		return false
	}
	host.ClearData()
	if test.Setup != "" {
		setupData, ok := host.JsonTests.Setups[test.Setup]
		if !ok {
			fmt.Printf("FAIL: Missing setup: %s\n", test.Setup)
			return false
		}
		host.LoadData(setupData)
	}
	host.LoadData(&test.JsonDataModel)
	request := host.FindSubObject(nil, "request", OBJTYPE_MAP).(*HostMap)
	if !host.runRequest(request, test.Request) {
		return false
	}
	for _, req := range test.AdditionalRequests {
		jsonRequest := req.(map[string]interface{})
		request.SetInt(KeyLength, 0)
		host.LoadSubMapData(request, jsonRequest)
		if !host.runRequest(request, jsonRequest) {
			return false
		}
	}

	scAddress := host.FindSubObject(nil, "contract", OBJTYPE_MAP).GetString(host.GetKeyId("address"))
	reqParams := host.FindSubObject(request, "params", OBJTYPE_MAP)
	postedRequests := host.FindSubObject(nil, "postedRequests", OBJTYPE_MAP_ARRAY)

	expectedPostedRequests := int64(len(test.Expect.PostedRequests))
	for i := int64(0); i < expectedPostedRequests && i < postedRequests.GetInt(KeyLength); i++ {
		postedRequest := host.FindObject(postedRequests.GetObjectId(int32(i), OBJTYPE_MAP))
		delay := postedRequest.GetInt(host.GetKeyId("delay"))
		if delay != 0 && !strings.Contains(test.Flags, "nodelay") {
			// only process posted requests when they have no delay
			// unless overridden by the nodelay flag
			// those are the only ones that will be incorporated in the final state
			continue
		}

		contractAddress := postedRequest.GetString(host.GetKeyId("contract"))
		if contractAddress != scAddress {
			// only process posted requests when they are for the current contract
			// those are the only ones that will be incorporated in the final state
			continue
		}

		function := postedRequest.GetString(host.GetKeyId("function"))
		request.SetString(host.GetKeyId("address"), scAddress)
		request.SetString(host.GetKeyId("function"), function)
		reqParams.SetInt(KeyLength, 0)
		params := host.FindObject(postedRequest.GetObjectId(host.GetKeyId("params"), OBJTYPE_MAP))
		params.(*HostMap).CopyDataTo(reqParams)
		fmt.Printf("    Run function: %v\n", function)
		err := host.RunFunction(function)
		if err != nil {
			fmt.Printf("FAIL: Request function %s: %v\n", function, err)
			return false
		}
	}

	// now compare the expected json data model to the actual host data model
	return host.CompareData(test)
}

func (host *SimpleWasmHost) runRequest(request *HostMap, req map[string]interface{}) bool {
	function, ok := req["function"]
	if !ok {
		fmt.Printf("FAIL: Missing request.function\n")
		return false
	}
	if request.Exists(host.GetKeyId("balance")) {
		reqColors := host.FindSubObject(request, "colors", OBJTYPE_STRING_ARRAY).(*HostArray)
		reqBalance := host.FindSubObject(request, "balance", OBJTYPE_MAP).(*HostMap)
		account := host.FindSubObject(nil, "account", OBJTYPE_MAP).(*HostMap)
		accBalance := host.FindSubObject(account, "balance", OBJTYPE_MAP).(*HostMap)
		for i := reqColors.GetLength() - 1; i >= 0; i-- {
			color := reqColors.GetBytes(i)
			colorKeyId := host.GetKeyId(base58.Encode(color))
			accBalance.SetInt(colorKeyId, accBalance.GetInt(colorKeyId)+reqBalance.GetInt(colorKeyId))
		}
	}

	fmt.Printf("    Run function: %v\n", function)
	err := host.RunFunction(function.(string))
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
