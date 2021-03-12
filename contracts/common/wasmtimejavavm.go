// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"errors"
	"github.com/bytecodealliance/wasmtime-go"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"strconv"
)

type WasmTimeJavaVM struct {
	wasmhost.WasmVmBase
	instance   *wasmtime.Instance
	isJavaWasm bool
	linker     *wasmtime.Linker
	memory     *wasmtime.Memory
	module     *wasmtime.Module
	store      *wasmtime.Store
}

var javaTypes = []string{
	"ii:i",
	"i:i",
	"iiiii:i",
	"iiii:i",
	"if:i",
	"iiiii",
	"iii",
	"ii",
	"iii:i",
	"iiiiii",
	"iiiiii:i",
	"iiii",
	"i",
	"iiiiiii:i",
	"i:f",
	"iifi",
	"iif",
	"if",
	"iff:i",
	"iiff",
	"iiffi",
	"iiiiiiii",
	"iif:i",
	"iiiiiiii:i",
	":i",
}

var javaImports = []string{
	"stringutf16", "isBigEndian", "1",
	"vm", "newLambdaStaticInvocationStringMethodTypeMethodHandleObject", "2",
	"vm", "newLambdaConstructorInvocationMethodTypeMethodHandleObject", "3",
	"vm", "newLambdaInterfaceInvocationMethodTypeMethodHandleObject", "3",
	"vm", "newLambdaVirtualInvocationMethodTypeMethodHandleObject", "3",
	"vm", "newLambdaSpecialInvocationMethodTypeMethodHandleObject", "3",
	"system", "nanoTime", "1",
	"float", "floatToRawIntBitsFLOAT", "4",
	"fileinputstream", "open0String", "0",
	"fileinputstream", "readBytesINTL1BYTEINTINT", "2",
	"fileinputstream", "read0INT", "0",
	"fileinputstream", "available0INT", "0",
	"fileoutputstream", "writeBytesINTL1BYTEINTINT", "5",
	"fileoutputstream", "writeIntINTINT", "6",
	"fileoutputstream", "close0INT", "7",
	"unixfilesystem", "getBooleanAttributes0String", "0",
	"double", "doubleToRawLongBitsDOUBLE", "4",
	"WasmLib", "javaGetObjectId", "3",
	"WasmLib", "javaGetKeyId", "8",
	"WasmLib", "javaSetBytes", "9",
	"WasmLib", "javaGetBytes", "10",
	"memorymanager", "isUsedAsCallbackINT", "0",
	"memorymanager", "logAllocationDetailsINTINTINT", "11",
}

func NewWasmTimeJavaVM() *WasmTimeJavaVM {
	vm := &WasmTimeJavaVM{}
	vm.store = wasmtime.NewStore(wasmtime.NewEngine())
	vm.linker = wasmtime.NewLinker(vm.store)
	return vm
}

func (vm *WasmTimeJavaVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	vm.WasmVmBase.LinkHost(impl, host)
	err := vm.linker.DefineFunc("WasmLib", "hostGetBytes",
		func(objId int32, keyId int32, typeId int32, stringRef int32, size int32) int32 {
			return vm.HostGetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostGetKeyId",
		func(keyRef int32, size int32) int32 {
			return vm.HostGetKeyId(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostGetObjectId",
		func(objId int32, keyId int32, typeId int32) int32 {
			return vm.HostGetObjectId(objId, keyId, typeId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostSetBytes",
		func(objId int32, keyId int32, typeId int32, stringRef int32, size int32) {
			vm.HostSetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}

	// go implementation uses this one to write panic message
	err = vm.linker.DefineFunc("wasi_unstable", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return vm.HostFdWrite(fd, iovs, size, written)
		})
	if err != nil {
		return err
	}

	// java versions of host functions have one extra dummy parameter
	err = vm.linker.DefineFunc("WasmLib", "javaGetBytes",
		func(dummy int32, objId int32, keyId int32, typeId int32, stringRef int32, size int32) int32 {
			return vm.HostGetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaGetKeyId",
		func(dummy int32, keyRef int32, size int32) int32 {
			return vm.HostGetKeyId(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaGetObjectId",
		func(dummy int32, objId int32, keyId int32, typeId int32) int32 {
			return vm.HostGetObjectId(objId, keyId, typeId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaSetBytes",
		func(dummy int32, objId int32, keyId int32, typeId int32, stringRef int32, size int32) {
			vm.HostSetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}

	for i := 0; i < len(javaImports); i += 3 {
		module := javaImports[i]
		name := javaImports[i+1]
		typeNr, _ := strconv.Atoi(javaImports[i+2])
		types := javaTypes[typeNr]
		params := make([]*wasmtime.ValType, 0)
		results := make([]*wasmtime.ValType, 0)
		for j := 0; j < len(types); j++ {
			switch types[j] {
			case 'i':
				params = append(params, wasmtime.NewValType(wasmtime.KindI32))
			case 'f':
				params = append(params, wasmtime.NewValType(wasmtime.KindF32))
			case ':':
				j++
				switch types[j] {
				case 'i':
					results = append(results, wasmtime.NewValType(wasmtime.KindI32))
				case 'f':
					results = append(results, wasmtime.NewValType(wasmtime.KindF32))
				}
			}
		}
		funcType := wasmtime.NewFuncType(params, results)
		vm.linker.Define(module, name, wasmtime.NewFunc(vm.store, funcType,
			func(caller *wasmtime.Caller, vals []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
				panic("java called " + module + "." + name)
			}))
	}
	return nil
}

func (vm *WasmTimeJavaVM) LoadWasm(wasmData []byte) error {
	var err error
	vm.module, err = wasmtime.NewModule(vm.store.Engine, wasmData)
	if err != nil {
		return err
	}
	vm.instance, err = vm.linker.Instantiate(vm.module)
	if err != nil {
		return err
	}
	err = vm.RunFunction("initMemory", 0)
	if err == nil {
		vm.RunFunction("bootstrap")
		if err != nil {
			return err
		}
		vm.isJavaWasm = true
	}
	memory := vm.instance.GetExport("memory")
	if memory == nil {
		return errors.New("no memory export")
	}
	vm.memory = memory.Memory()
	if vm.memory == nil {
		return errors.New("not a memory type")
	}
	return nil
}

func (vm *WasmTimeJavaVM) RunFunction(functionName string, args ...interface{}) error {
	export := vm.instance.GetExport(functionName)
	if export == nil {
		return errors.New("unknown export function: '" + functionName + "'")
	}
	if vm.isJavaWasm && functionName == "on_load" {
		// insert dummy zero first argument
		_, err := export.Func().Call(0)
		return err
	}
	_, err := export.Func().Call(args...)
	return err
}

func (vm *WasmTimeJavaVM) RunScFunction(index int32) error {
	export := vm.instance.GetExport("on_call")
	if export == nil {
		return errors.New("unknown export function: 'on_call'")
	}
	if vm.isJavaWasm {
		frame := vm.PreCall()
		// insert dummy zero first argument
		_, err := export.Func().Call(0, index)
		vm.PostCall(frame)
		return err
	}

	frame := vm.PreCall()
	_, err := export.Func().Call(index)
	vm.PostCall(frame)
	return err
}

func (vm *WasmTimeJavaVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}

func (vm *WasmTimeJavaVM) VmGetBytes(offset int32, size int32) []byte {
	ptr := vm.UnsafeMemory()
	bytes := make([]byte, size)
	if vm.isJavaWasm {
		offset += 20
		for i := int32(0); i < size; i++ {
			bytes[i] = ptr[offset]
			offset += 8
		}
		return bytes
	}
	copy(bytes, ptr[offset:offset+size])
	return bytes
}

func (vm *WasmTimeJavaVM) VmSetBytes(offset int32, size int32, bytes []byte) int32 {
	if size != 0 {
		ptr := vm.UnsafeMemory()
		if vm.isJavaWasm {
			offset += 20
			for i := int32(0); i < size; i++ {
				ptr[offset] = bytes[i]
				offset += 8
			}
			return int32(len(bytes))
		}
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}
