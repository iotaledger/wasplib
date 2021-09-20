// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"strconv"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type WasmTimeJavaVM struct {
	wasmhost.WasmVMBase
	instance   *wasmtime.Instance
	interrupt  *wasmtime.InterruptHandle
	isJavaWasm bool
	linker     *wasmtime.Linker
	memory     *wasmtime.Memory
	module     *wasmtime.Module
	store      *wasmtime.Store
}

var _ wasmhost.WasmVM = &WasmTimeJavaVM{}

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
	"WasmLib", "javaGetObjectID", "3",
	"WasmLib", "javaGetKeyID", "8",
	"WasmLib", "javaSetBytes", "9",
	"WasmLib", "javaGetBytes", "10",
	"memorymanager", "isUsedAsCallbackINT", "0",
	"memorymanager", "logAllocationDetailsINTINTINT", "11",
}

func NewWasmTimeJavaVM() *WasmTimeJavaVM {
	vm := &WasmTimeJavaVM{}
	config := wasmtime.NewConfig()
	config.SetInterruptable(true)
	vm.store = wasmtime.NewStore(wasmtime.NewEngineWithConfig(config))
	vm.interrupt, _ = vm.store.InterruptHandle()
	vm.linker = wasmtime.NewLinker(vm.store)
	return vm
}

func (vm *WasmTimeJavaVM) Interrupt() {
	vm.interrupt.Interrupt()
}

func (vm *WasmTimeJavaVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	_ = vm.WasmVMBase.LinkHost(impl, host)

	err := vm.linker.DefineFunc("WasmLib", "hostGetBytes",
		func(objID, keyID int32, typeID, stringRef, size int32) int32 {
			return vm.HostGetBytes(objID, keyID, typeID, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostGetKeyID",
		func(keyRef, size int32) int32 {
			return vm.HostGetKeyID(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostGetObjectID",
		func(objID, keyID, typeID int32) int32 {
			return vm.HostGetObjectID(objID, keyID, typeID)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "hostSetBytes",
		func(objID, keyID int32, typeID, stringRef, size int32) {
			vm.HostSetBytes(objID, keyID, typeID, stringRef, size)
		})
	if err != nil {
		return err
	}

	// TinyGo Wasm versions uses this one to write panic message to console
	fdWrite := func(fd, iovs, size, written int32) int32 {
		return vm.HostFdWrite(fd, iovs, size, written)
	}
	err = vm.linker.DefineFunc("wasi_unstable", "fd_write", fdWrite)
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasi_snapshot_preview1", "fd_write", fdWrite)
	if err != nil {
		return err
	}

	return vm.linkHostJava()
}

func (vm *WasmTimeJavaVM) linkHostJava() error {
	// java versions of host functions have one extra dummy parameter
	err := vm.linker.DefineFunc("WasmLib", "javaGetBytes",
		func(dummy, objID int32, keyID, typeID, stringRef, size int32) int32 {
			return vm.HostGetBytes(objID, keyID, typeID, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaGetKeyID",
		func(dummy, keyRef, size int32) int32 {
			return vm.HostGetKeyID(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaGetObjectID",
		func(dummy, objID, keyID, typeID int32) int32 {
			return vm.HostGetObjectID(objID, keyID, typeID)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("WasmLib", "javaSetBytes",
		func(dummy, objID int32, keyID, typeID, stringRef, size int32) {
			vm.HostSetBytes(objID, keyID, typeID, stringRef, size)
		})
	if err != nil {
		return err
	}

	return vm.linkHostJavaSymbols()
}

func (vm *WasmTimeJavaVM) linkHostJavaSymbols() error {
	for i := 0; i < len(javaImports); i += 3 {
		module := javaImports[i]
		if module == "WasmLib" {
			// should have already been defined
			continue
		}
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
		err := vm.linker.Define(module, name, wasmtime.NewFunc(vm.store, funcType,
			func(caller *wasmtime.Caller, vals []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
				panic("java called " + module + "." + name)
			}))
		if err != nil {
			return err
		}
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
		err = vm.RunFunction("bootstrap")
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
		return vm.Run(func() (err error) {
			// insert dummy zero first argument
			_, err = export.Func().Call(0)
			return
		})
	}
	return vm.Run(func() (err error) {
		_, err = export.Func().Call(args...)
		return
	})
}

func (vm *WasmTimeJavaVM) RunScFunction(index int32) error {
	export := vm.instance.GetExport("on_call")
	if export == nil {
		return errors.New("unknown export function: 'on_call'")
	}
	if vm.isJavaWasm {
		frame := vm.PreCall()
		err := vm.Run(func() (err error) {
			// insert dummy zero first argument
			_, err = export.Func().Call(0)
			return
		})
		vm.PostCall(frame)
		return err
	}

	frame := vm.PreCall()
	err := vm.Run(func() (err error) {
		_, err = export.Func().Call(index)
		return
	})
	vm.PostCall(frame)
	return err
}

func (vm *WasmTimeJavaVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}

func (vm *WasmTimeJavaVM) VMGetBytes(offset, size int32) []byte {
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

func (vm *WasmTimeJavaVM) VMSetBytes(offset, size int32, bytes []byte) int32 {
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
