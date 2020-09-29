package host

import (
	"fmt"
	"github.com/iotaledger/wasplib/host/interfaces"
	"github.com/iotaledger/wasplib/host/interfaces/objtype"
	"github.com/iotaledger/wasplib/jsontest"
)

var EnableImmutableChecks = true

type HostImpl struct {
	HostBase
}

func NewHostImpl() *HostImpl {
	h := &HostImpl{}
	h.Init(h, NewHostMap(h), nil)
	return h
}

func (h *HostImpl) AddBalance(obj interfaces.HostObject, color string, amount int64) {
	colors := h.Object(obj, "colors", objtype.OBJTYPE_STRING_ARRAY)
	length := colors.GetInt(interfaces.KeyLength)
	colors.SetString(int32(length), color)
	colorId := h.GetKeyId(color)
	balance := h.Object(obj, "balance", objtype.OBJTYPE_MAP)
	balance.SetInt(colorId, amount)
}

func (h *HostImpl) ClearData() {
	h.ClearMapData("contract")
	h.ClearMapData("account")
	h.ClearMapData("request")
	h.ClearMapData("state")
	h.ClearArrayData("logs")
	h.ClearArrayData("events")
	h.ClearArrayData("transfers")
}

func (h *HostImpl) ClearArrayData(key string) {
	data := h.Object(nil, key, objtype.OBJTYPE_MAP_ARRAY)
	data.SetInt(interfaces.KeyLength, 0)
}

func (h *HostImpl) ClearMapData(key string) {
	data := h.Object(nil, key, objtype.OBJTYPE_MAP)
	data.SetInt(interfaces.KeyLength, 0)
}

func (h *HostImpl) CompareArrayData(key string, array []interface{}) bool {
	data := h.Object(nil, key, objtype.OBJTYPE_MAP_ARRAY)
	if data.GetInt(interfaces.KeyLength) != int64(len(array)) {
		fmt.Printf("FAIL: array %s length\n", key)
		return false
	}
	for i := range array {
		submap := h.GetObject(data.GetObjectId(int32(i), objtype.OBJTYPE_MAP))
		if !h.CompareSubMapData(submap, array[i].(map[string]interface{})) {
			fmt.Printf("      map %s\n", key)
			return false
		}
	}
	return true
}

func (h *HostImpl) CompareData(expect *jsontest.JsonModel) bool {
	return h.CompareMapData("state", expect.State) &&
		h.CompareArrayData("logs", expect.Logs) &&
		h.CompareArrayData("events", expect.Events) &&
		h.CompareArrayData("transfers", expect.Transfers)
}

func (h *HostImpl) CompareMapData(key string, values map[string]interface{}) bool {
	data := h.Object(nil, key, objtype.OBJTYPE_MAP)
	if !h.CompareSubMapData(data, values) {
		fmt.Printf("      map %s\n", key)
		return false
	}
	return true
}

func (h *HostImpl) CompareSubMapData(data interfaces.HostObject, values map[string]interface{}) bool {
	for k, v := range values {
		switch c := v.(type) {
		case string:
			got := data.GetString(h.GetKeyId(k))
			exp := v.(string)
			if got != exp {
				fmt.Printf("FAIL: string %s, expected '%s', got '%s'\n", k, exp, got)
				return false
			}
		case float64:
			got := data.GetInt(h.GetKeyId(k))
			exp := int64(v.(float64))
			if exp != got {
				fmt.Printf("FAIL: int %s, expected %d, got %d\n", k, exp, got)
				return false
			}
		case map[string]interface{}:
			submap := h.Object(data, k, objtype.OBJTYPE_MAP)
			if !h.CompareSubMapData(submap, v.(map[string]interface{})) {
				fmt.Printf("      map %s\n", k)
				return false
			}
		default:
			panic(fmt.Sprintf("Invalid type: %T", c))
		}
	}
	return true
}

func (h *HostImpl) LoadData(model *jsontest.JsonModel) {
	h.LoadMapData("contract", model.Contract)
	h.LoadMapData("account", model.Account)
	h.LoadMapData("request", model.Request)
	h.LoadMapData("state", model.State)
}

func (h *HostImpl) LoadMapData(key string, values map[string]interface{}) {
	data := h.Object(nil, key, objtype.OBJTYPE_MAP)
	h.LoadSubMapData(data, values)
}

func (h *HostImpl) LoadSubMapData(data interfaces.HostObject, values map[string]interface{}) {
	for k, v := range values {
		switch c := v.(type) {
		case string:
			data.SetString(h.GetKeyId(k), v.(string))
		case float64:
			data.SetInt(h.GetKeyId(k), int64(v.(float64)))
		case map[string]interface{}:
			submap := h.Object(data, k, objtype.OBJTYPE_MAP)
			h.LoadSubMapData(submap, v.(map[string]interface{}))
		default:
			panic(fmt.Sprintf("Invalid type: %T", c))
		}
	}
}

func (h *HostImpl) Object(obj interfaces.HostObject, key string, typeId int32) interfaces.HostObject {
	if obj == nil {
		// use root object
		obj = h.GetObject(1)
	}
	return h.GetObject(obj.GetObjectId(h.GetKeyId(key), typeId))
}

func (h *HostImpl) RunTest(name string, t *jsontest.JsonModel, testData *jsontest.JsonTest) {
	fmt.Printf("Test: %s\n", name)
	if t.Expect == nil {
		fmt.Printf("FAIL: Missing expect model data\n")
		return
	}
	h.ClearData()
	if t.Setup != "" {
		s, ok := testData.Setups[t.Setup]
		if !ok {
			fmt.Printf("FAIL: Missing setup: %s\n", t.Setup)
			return
		}
		h.LoadData(s)
	}
	h.LoadData(t)
	params, ok := t.Request["params"]
	if !ok {
		fmt.Printf("FAIL: Missing request.params\n")
		return
	}
	paramsMap := params.(map[string]interface{})
	fn, ok := paramsMap["fn"]
	if !ok {
		fmt.Printf("FAIL: Missing request.params.fn\n")
		return
	}
	err := h.RunWasmFunction(fn.(string))
	if err != nil {
		fmt.Printf("FAIL: Missing function: %v\n", fn)
		return
	}

	request := h.Object(nil, "request", objtype.OBJTYPE_MAP)
	reqParams := h.Object(request, "params", objtype.OBJTYPE_MAP)
	events := h.Object(nil, "events", objtype.OBJTYPE_MAP_ARRAY)
	i := int64(0)
	expectedEvents := int64(len(t.Expect.Events))
	for i < events.GetInt(interfaces.KeyLength) {
		event := h.GetObject(events.GetObjectId(int32(i), objtype.OBJTYPE_MAP))
		contract := event.GetString(h.GetKeyId("contract"))
		if contract != "" {
			fmt.Printf("FAIL: Expected empty contract name: %s\n", contract)
			return
		}
		function := event.GetString(h.GetKeyId("function"))
		if i >= expectedEvents {
			fmt.Printf("FAIL: Unexpected event function call: %s\n", function)
			return
		}
		reqParams.SetInt(interfaces.KeyLength, 0)
		reqParams.SetString(h.GetKeyId("fn"), function)
		params := h.GetObject(event.GetObjectId(h.GetKeyId("params"), objtype.OBJTYPE_MAP))
		params.(*HostMap).CopyDataTo(reqParams)
		err = h.RunWasmFunction(function)
		if err != nil {
			fmt.Printf("FAIL: Missing event function: %s\n", function)
			return
		}
		i++
	}

	// now compare the expect data model to the actual data model
	if h.CompareData(t.Expect) {
		fmt.Printf("PASS\n")
	}
}
