package jsontest

type JsonTest struct {
	Setups map[string]*JsonModel `json:"setups"`
	Tests  map[string]*JsonModel `json:"tests"`
}

type JsonModel struct {
	Setup     string                 `json:"setup"`
	Contract  map[string]interface{} `json:"contract"`
	Account   map[string]interface{} `json:"account"`
	Request   map[string]interface{} `json:"request"`
	State     map[string]interface{} `json:"state"`
	Logs      []interface{}          `json:"logs"`
	Events    []interface{}          `json:"events"`
	Transfers []interface{}          `json:"transfers"`
	Expect    *JsonModel             `json:"expect"`
}
