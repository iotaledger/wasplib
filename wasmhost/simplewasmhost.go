package wasmhost

import (
	"fmt"
	"github.com/iotaledger/wasplib/jsontest"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	WasmHost
}

func NewHostImpl() *SimpleWasmHost {
	host := &SimpleWasmHost{}
	host.Init(NewHostMap(host), nil, host)
	return host
}

func (host *SimpleWasmHost) AddBalance(obj HostObject, color string, amount int64) {
	colors := host.Object(obj, "colors", OBJTYPE_STRING_ARRAY)
	length := colors.GetInt(KeyLength)
	colors.SetString(int32(length), color)
	colorId := host.GetKeyId(color)
	balance := host.Object(obj, "balance", OBJTYPE_MAP)
	balance.SetInt(colorId, amount)
}

func (host *SimpleWasmHost) ClearData() {
	host.ClearMapData("contract")
	host.ClearMapData("account")
	host.ClearMapData("request")
	host.ClearMapData("state")
	host.ClearArrayData("logs")
	host.ClearArrayData("events")
	host.ClearArrayData("transfers")
}

func (host *SimpleWasmHost) ClearArrayData(key string) {
	data := host.Object(nil, key, OBJTYPE_MAP_ARRAY)
	data.SetInt(KeyLength, 0)
}

func (host *SimpleWasmHost) ClearMapData(key string) {
	data := host.Object(nil, key, OBJTYPE_MAP)
	data.SetInt(KeyLength, 0)
}

func (host *SimpleWasmHost) CompareArrayData(key string, array []interface{}) bool {
	data := host.Object(nil, key, OBJTYPE_MAP_ARRAY)
	if data.GetInt(KeyLength) != int64(len(array)) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	for i := range array {
		submap := host.FindObject(data.GetObjectId(int32(i), OBJTYPE_MAP))
		if !host.CompareSubMapData(submap, array[i].(map[string]interface{})) {
			fmt.Printf("      map %s\n", key)
			return false
		}
	}
	return true
}

func (host *SimpleWasmHost) CompareData(expect *jsontest.JsonModel) bool {
	return host.CompareMapData("state", expect.State) &&
		host.CompareArrayData("logs", expect.Logs) &&
		host.CompareArrayData("events", expect.Events) &&
		host.CompareArrayData("transfers", expect.Transfers)
}

func (host *SimpleWasmHost) CompareMapData(key string, values map[string]interface{}) bool {
	data := host.Object(nil, key, OBJTYPE_MAP)
	if !host.CompareSubMapData(data, values) {
		fmt.Printf("      map %s\n", key)
		return false
	}
	return true
}

func (host *SimpleWasmHost) CompareSubMapData(data HostObject, values map[string]interface{}) bool {
	for k, v := range values {
		switch c := v.(type) {
		case string:
			got := data.GetString(host.GetKeyId(k))
			exp := v.(string)
			if got != exp {
				fmt.Printf("FAIL: string %s, expected '%s', got '%s'\n", k, exp, got)
				return false
			}
		case float64:
			got := data.GetInt(host.GetKeyId(k))
			exp := int64(v.(float64))
			if exp != got {
				fmt.Printf("FAIL: int %s, expected %d, got %d\n", k, exp, got)
				return false
			}
		case map[string]interface{}:
			submap := host.Object(data, k, OBJTYPE_MAP)
			if !host.CompareSubMapData(submap, v.(map[string]interface{})) {
				fmt.Printf("      map %s\n", k)
				return false
			}
		default:
			panic(fmt.Sprintf("Invalid type: %T", c))
		}
	}
	return true
}

func (host *SimpleWasmHost) LoadData(model *jsontest.JsonModel) {
	host.LoadMapData("contract", model.Contract)
	host.LoadMapData("account", model.Account)
	host.LoadMapData("request", model.Request)
	host.LoadMapData("state", model.State)
}

func (host *SimpleWasmHost) LoadMapData(key string, values map[string]interface{}) {
	data := host.Object(nil, key, OBJTYPE_MAP)
	host.LoadSubMapData(data, values)
}

func (host *SimpleWasmHost) LoadSubMapData(data HostObject, values map[string]interface{}) {
	for k, v := range values {
		switch c := v.(type) {
		case string:
			data.SetString(host.GetKeyId(k), v.(string))
		case float64:
			data.SetInt(host.GetKeyId(k), int64(v.(float64)))
		case map[string]interface{}:
			submap := host.Object(data, k, OBJTYPE_MAP)
			host.LoadSubMapData(submap, v.(map[string]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", c))
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

func (host *SimpleWasmHost) Object(obj HostObject, key string, typeId int32) HostObject {
	if obj == nil {
		// use root object
		obj = host.FindObject(1)
	}
	return host.FindObject(obj.GetObjectId(host.GetKeyId(key), typeId))
}

func (host *SimpleWasmHost) RunTest(name string, t *jsontest.JsonModel, testData *jsontest.JsonTest) {
	fmt.Printf("Test: %s\n", name)
	if t.Expect == nil {
		fmt.Printf("FAIL: Missing expect model data\n")
		return
	}
	host.ClearData()
	if t.Setup != "" {
		s, ok := testData.Setups[t.Setup]
		if !ok {
			fmt.Printf("FAIL: Missing setup: %s\n", t.Setup)
			return
		}
		host.LoadData(s)
	}
	host.LoadData(t)
	function, ok := t.Request["function"]
	if !ok {
		fmt.Printf("FAIL: Missing request.function\n")
		return
	}
	err := host.RunWasmFunction(function.(string))
	if err != nil {
		fmt.Printf("FAIL: Unknown function: %v\n", function)
		return
	}

	request := host.Object(nil, "request", OBJTYPE_MAP)
	reqParams := host.Object(request, "params", OBJTYPE_MAP)
	events := host.Object(nil, "events", OBJTYPE_MAP_ARRAY)
	i := int64(0)
	expectedEvents := int64(len(t.Expect.Events))
	for i < events.GetInt(KeyLength) {
		event := host.FindObject(events.GetObjectId(int32(i), OBJTYPE_MAP))
		contract := event.GetString(host.GetKeyId("contract"))
		if contract != "" {
			fmt.Printf("FAIL: Expected empty contract name: %s\n", contract)
			return
		}
		function := event.GetString(host.GetKeyId("function"))
		if i >= expectedEvents {
			fmt.Printf("FAIL: Unexpected event function call: %s\n", function)
			return
		}
		//TODO set request parameters for new request
		request.SetString(host.GetKeyId("function"), function)
		reqParams.SetInt(KeyLength, 0)
		params := host.FindObject(event.GetObjectId(host.GetKeyId("params"), OBJTYPE_MAP))
		params.(*HostMap).CopyDataTo(reqParams)
		err = host.RunWasmFunction(function)
		if err != nil {
			fmt.Printf("FAIL: Unknown event function: %s\n", function)
			return
		}
		i++
	}

	// now compare the expect data model to the actual data model
	if host.CompareData(t.Expect) {
		fmt.Printf("PASS\n")
	}
}
