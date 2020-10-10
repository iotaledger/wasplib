package wasmhost

import "fmt"

type HostExports struct {
	HostArray
}

func NewHostExports(host *SimpleWasmHost) *HostExports {
	return &HostExports{HostArray: *NewHostArray(host, OBJTYPE_STRING)}
}

func (a *HostExports) SetString(keyId int32, value string) {
	fmt.Printf("%s = %d\n", value, keyId)
}
