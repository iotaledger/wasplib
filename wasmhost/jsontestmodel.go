package wasmhost

type JsonType struct {
	FieldName string `json:"field"`
	TypeName  string `json:"type"`
}

type JsonTests struct {
	Types  map[string][]*JsonType    `json:"types"`
	Setups map[string]*JsonDataModel `json:"setups"`
	Tests  []*JsonTest               `json:"tests"`
}

type JsonDataModel struct {
	Contract       map[string]interface{} `json:"contract"`
	Account        map[string]interface{} `json:"account"`
	Request        map[string]interface{} `json:"request"`
	State          map[string]interface{} `json:"state"`
	Utility        map[string]interface{} `json:"utility"`
	Logs           map[string]interface{} `json:"logs"`
	PostedRequests []interface{}          `json:"postedRequests"`
	Transfers      []interface{}          `json:"transfers"`
}

type JsonTest struct {
	JsonDataModel
	Name               string         `json:"name"`
	Setup              string         `json:"setup"`
	Flags              string         `json:"flags"`
	AdditionalRequests []interface{}  `json:"additionalRequests"`
	Expect             *JsonDataModel `json:"expect"`
}
