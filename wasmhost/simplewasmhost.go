package wasmhost

import (
	"fmt"
	"sort"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	WasmHost
	ExportsId int32
}

func NewSimpleWasmHost() (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	host.useBase58Keys = true
	err := host.Init(NewNullObject(host), NewHostMap(host), nil, host)
	if err != nil {
		return nil, err
	}
	host.ExportsId = host.GetKeyId("exports")
	return host, nil
}

func (host *SimpleWasmHost) ClearData() {
	host.ClearObjectData(OBJTYPE_MAP, "contract")
	host.ClearObjectData(OBJTYPE_MAP, "account")
	host.ClearObjectData(OBJTYPE_MAP, "request")
	host.ClearObjectData(OBJTYPE_MAP, "state")
	host.ClearObjectData(OBJTYPE_MAP_ARRAY, "logs")
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

func (host *SimpleWasmHost) CompareData(expectData *JsonDataModel) bool {
	return host.CompareMapData("state", expectData.State) &&
		host.CompareArrayData("logs", expectData.Logs) &&
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

func (host *SimpleWasmHost) CompareSubMapData(mapObject HostObject, values map[string]interface{}) bool {
	for _, key := range SortedKeys(values) {
		field := values[key]
		switch t := field.(type) {
		case string:
			value := mapObject.GetString(host.GetKeyId(key))
			expect := field.(string)
			if value != expect {
				fmt.Printf("FAIL: string %s, expected '%s', got '%s'\n", key, expect, value)
				return false
			}
		case float64:
			value := mapObject.GetInt(host.GetKeyId(key))
			expect := int64(field.(float64))
			if value != expect {
				fmt.Printf("FAIL: int %s, expected %d, got %d\n", key, expect, value)
				return false
			}
		case map[string]interface{}:
			subMapObject := host.FindSubObject(mapObject, key, OBJTYPE_MAP)
			if !host.CompareSubMapData(subMapObject, field.(map[string]interface{})) {
				fmt.Printf("      map %s\n", key)
				return false
			}
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
	host.Trace("GetKeyId '%s'=k%d", key, keyId)
	return keyId
}

func (host *SimpleWasmHost) LoadData(jsonData *JsonDataModel) {
	host.LoadMapData("contract", jsonData.Contract)
	host.LoadMapData("account", jsonData.Account)
	host.LoadMapData("request", jsonData.Request)
	host.LoadMapData("state", jsonData.State)
}

func (host *SimpleWasmHost) LoadMapData(key string, values map[string]interface{}) {
	mapObject := host.FindSubObject(nil, key, OBJTYPE_MAP)
	host.LoadSubMapData(mapObject, values)
}

func (host *SimpleWasmHost) LoadSubArrayData(arrayObject HostObject, values []interface{}) {
	for key, field := range values {
		switch t := field.(type) {
		case string:
			arrayObject.SetString(int32(key), field.(string))
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
	for _, key := range SortedKeys(values) {
		field := values[key]
		switch t := field.(type) {
		case string:
			mapObject.SetString(host.GetKeyId(key), field.(string))
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
		fmt.Println(text)
	case KeyLog:
		fmt.Println(text)
	case KeyWarning:
		fmt.Println(text)
	case KeyError:
		fmt.Println(text)
	}
}

func (host *SimpleWasmHost) RunTest(name string, jsonData *JsonDataModel, jsonTests *JsonTests) {
	fmt.Printf("Test: %s\n", name)
	if jsonData.Expect == nil {
		fmt.Printf("FAIL: Missing expect model data\n")
		return
	}
	host.ClearData()
	if jsonData.Setup != "" {
		setupData, ok := jsonTests.Setups[jsonData.Setup]
		if !ok {
			fmt.Printf("FAIL: Missing setup: %s\n", jsonData.Setup)
			return
		}
		host.LoadData(setupData)
	}
	host.LoadData(jsonData)
	function, ok := jsonData.Request["function"]
	if !ok {
		fmt.Printf("FAIL: Missing request.function\n")
		return
	}
	err := host.RunFunction(function.(string))
	if err != nil {
		fmt.Printf("FAIL: Function %v: %v\n", function, err)
		return
	}

	scAddress := host.FindSubObject(nil, "contract", OBJTYPE_MAP).GetString(host.GetKeyId("address"))
	request := host.FindSubObject(nil, "request", OBJTYPE_MAP)
	reqParams := host.FindSubObject(request, "params", OBJTYPE_MAP)
	postedRequests := host.FindSubObject(nil, "postedRequests", OBJTYPE_MAP_ARRAY)
	i := int64(0)
	expectedPostedRequests := int64(len(jsonData.Expect.PostedRequests))
	for i < postedRequests.GetInt(KeyLength) {
		postedRequest := host.FindObject(postedRequests.GetObjectId(int32(i), OBJTYPE_MAP))
		contractAddress := postedRequest.GetString(host.GetKeyId("contract"))
		if contractAddress != scAddress {
			fmt.Printf("FAIL: Expected contract address: %s\n", contractAddress)
			return
		}
		function := postedRequest.GetString(host.GetKeyId("function"))
		if i >= expectedPostedRequests {
			fmt.Printf("FAIL: Unexpected posted request: %s\n", function)
			return
		}

		request.SetString(host.GetKeyId("address"), scAddress)
		request.SetString(host.GetKeyId("function"), function)
		reqParams.SetInt(KeyLength, 0)
		params := host.FindObject(postedRequest.GetObjectId(host.GetKeyId("params"), OBJTYPE_MAP))
		params.(*HostMap).CopyDataTo(reqParams)
		err = host.RunFunction(function)
		if err != nil {
			fmt.Printf("FAIL: Request function %s: %v\n", function, err)
			return
		}
		i++
	}

	// now compare the expected json data model to the actual host data model
	if host.CompareData(jsonData.Expect) {
		fmt.Printf("PASS\n")
	}
}

func SortedKeys(values map[string]interface{}) []string {
	keys := make([]string, len(values))
	index := 0
	for key, _ := range values {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	return keys
}
