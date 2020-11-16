// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import "fmt"

type HostExports struct {
	HostArray
}

func NewHostExports(host *SimpleWasmHost, keyId int32) *HostExports {
	return &HostExports{HostArray: *NewHostArray(host, keyId, OBJTYPE_STRING)}
}

func (a *HostExports) SetString(keyId int32, value string) {
	fmt.Printf("%s = %d\n", value, keyId)
	a.host.SetExport(keyId, value)
}
