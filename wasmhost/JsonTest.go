package wasmhost

type JsonTests struct {
	Setups map[string]*JsonDataModel `json:"setups"`
	Tests  map[string]*JsonDataModel `json:"tests"`
}

type JsonDataModel struct {
	Setup     string                 `json:"setup"`
	Contract  map[string]interface{} `json:"contract"`
	Account   map[string]interface{} `json:"account"`
	Request   map[string]interface{} `json:"request"`
	State     map[string]interface{} `json:"state"`
	Logs      []interface{}          `json:"logs"`
	Events    []interface{}          `json:"events"`
	Transfers []interface{}          `json:"transfers"`
	Expect    *JsonDataModel         `json:"expect"`
}
