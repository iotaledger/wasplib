// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmproc

type ScExports struct {
	ScSandboxObject
}

func NewScExports(vm *wasmProcessor) *ScExports {
	o := &ScExports{}
	o.vm = vm
	return o
}

func (o *ScExports) SetBytes(keyId int32, typeId int32, bytes []byte) {
	o.vm.SetExport(keyId, string(bytes))
}
