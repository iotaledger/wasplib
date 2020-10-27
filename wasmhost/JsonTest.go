package wasmhost

type JsonTests struct {
	Types  map[string][]string       `json:"types"`
	Setups map[string]*JsonDataModel `json:"setups"`
	Tests  map[string]*JsonDataModel `json:"tests"`
}

type JsonDataModel struct {
	Setup          string                 `json:"setup"`
	Flags          string                 `json:"flags"`
	Contract       map[string]interface{} `json:"contract"`
	Account        map[string]interface{} `json:"account"`
	Request        map[string]interface{} `json:"request"`
	State          map[string]interface{} `json:"state"`
	Utility        map[string]interface{} `json:"utility"`
	Logs           []interface{}          `json:"logs"`
	PostedRequests []interface{}          `json:"postedRequests"`
	Transfers      []interface{}          `json:"transfers"`
	Expect         *JsonDataModel         `json:"expect"`
}
