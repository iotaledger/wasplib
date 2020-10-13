package wasmhost

import (
	"github.com/iotaledger/wart/wasm/consts/value"
	"github.com/iotaledger/wart/wasm/executors"
	"github.com/iotaledger/wart/wasm/sections"
)

type WartVM struct {
	runner *executors.WasmRunner
}

func NewWartVM() *WartVM {
	host := &WartVM{}
	host.runner = executors.NewWasmRunner()
	return host
}

func (vm *WartVM) LinkHost(host *WasmHost) error {
	m := executors.DefineModule("wasplib")
	lnk := executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("hostGetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			stringRef := ctx.Frame[ctx.SP+2].I32
			size := ctx.Frame[ctx.SP+3].I32
			if objId >= 0 {
				ctx.Frame[ctx.SP].I32 = host.vmSetBytes(stringRef, size, host.GetBytes(objId, keyId))
				return nil
			}
			ctx.Frame[ctx.SP].I32 = host.vmSetBytes(stringRef, size, []byte(host.GetString(-objId, keyId)))
			return nil
		})
	_ = lnk.DefineFunction("hostGetInt",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I64},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			ctx.Frame[ctx.SP].I64 = host.GetInt(objId, keyId)
			return nil
		})
	_ = lnk.DefineFunction("hostGetIntRef",
		[]value.DataType{value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			intRef := ctx.Frame[ctx.SP+2].I32
			host.vmSetInt(intRef, host.GetInt(objId, keyId))
			return nil
		})
	_ = lnk.DefineFunction("hostGetKeyId",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			keyRef := ctx.Frame[ctx.SP].I32
			size := ctx.Frame[ctx.SP+1].I32
			ctx.Frame[ctx.SP].I32 = host.GetKeyId(string(host.vmGetBytes(keyRef, size)))
			return nil
		})
	_ = lnk.DefineFunction("hostGetObjectId",
		[]value.DataType{value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			typeId := ctx.Frame[ctx.SP+2].I32
			ctx.Frame[ctx.SP].I32 = host.GetObjectId(objId, keyId, typeId)
			return nil
		})
	_ = lnk.DefineFunction("hostSetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			stringRef := ctx.Frame[ctx.SP+2].I32
			size := ctx.Frame[ctx.SP+3].I32
			if objId >= 0 {
				host.SetBytes(objId, keyId, host.vmGetBytes(stringRef, size))
				return nil
			}
			host.SetString(-objId, keyId, string(host.vmGetBytes(stringRef, size)))
			return nil
		})
	_ = lnk.DefineFunction("hostSetInt",
		[]value.DataType{value.I32, value.I32, value.I64},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			value := ctx.Frame[ctx.SP+2].I64
			host.SetInt(objId, keyId, value)
			return nil
		})
	_ = lnk.DefineFunction("hostSetIntRef",
		[]value.DataType{value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			intRef := ctx.Frame[ctx.SP+2].I32
			host.SetInt(objId, keyId, host.vmGetInt(intRef))
			return nil
		})

	// go implementation uses this one to write panic message
	m = executors.DefineModule("wasi_unstable")
	lnk = executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("fd_write",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			fd := ctx.Frame[ctx.SP].I32
			iovs := ctx.Frame[ctx.SP+1].I32
			size := ctx.Frame[ctx.SP+2].I32
			written := ctx.Frame[ctx.SP+3].I32
			ctx.Frame[ctx.SP].I32 = host.fdWrite(fd, iovs, size, written)
			return nil
		})
	return nil
}

func (vm *WartVM) LoadWasm(wasmData []byte) error {
	return vm.runner.Load(wasmData)
}

func (vm *WartVM) RunFunction(functionName string) error {
	return vm.runner.RunExport(functionName)
}

func (vm *WartVM) UnsafeMemory() []byte {
	return vm.runner.Memory()
}
