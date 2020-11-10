package wasmhost

import (
	"fmt"
	"hash/fnv"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	WasmHost
	ExportsId   int32
	TransfersId int32
}

func NewSimpleWasmHost() (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	host.useBase58Keys = true
	err := host.Init(NewNullObject(host), NewHostMap(host, 0), nil, host)
	if err != nil {
		return nil, err
	}
	host.ExportsId = host.GetKeyIdFromBytes([]byte("exports"))
	host.TransfersId = host.GetKeyIdFromBytes([]byte("transfers"))
	return host, nil
}

func (host *SimpleWasmHost) FindSubObject(obj HostObject, key string, typeId int32) HostObject {
	if obj == nil {
		// use root object
		obj = host.FindObject(1)
	}
	return host.FindObject(obj.GetObjectId(host.GetKeyId(key), typeId))
}

func (host *SimpleWasmHost) GetKeyId(key string) int32 {
	keyId := host.GetKeyIdFromBytes([]byte(key))
	host.Trace("GetKeyId('%s')=k%d", key, keyId)
	return keyId
}

func (host *SimpleWasmHost) SetExport(index int32, functionName string) {
	if index != int32(len(host.codeToFunc)+1) {
		host.SetError("SetExport: invalid index")
		return
	}
	_, ok := host.funcToCode[functionName]
	if ok {
		host.SetError("SetExport: duplicate function name")
		return
	}
	h := fnv.New32a()
	h.Write([]byte(functionName))
	hashedName := h.Sum32()
	_, ok = host.codeToFunc[hashedName]
	if ok {
		host.SetError("SetExport: duplicate hashed name")
		return
	}
	host.funcToCode[functionName] = hashedName
	host.codeToFunc[hashedName] = functionName
}

func (host *SimpleWasmHost) Log(logLevel int32, text string) {
	switch logLevel {
	case KeyTraceHost:
		//fmt.Println(text)
	case KeyTrace:
		//fmt.Println(text)
	case KeyLog:
		fmt.Println(text)
	case KeyWarning:
		fmt.Println(text)
	case KeyError:
		fmt.Println(text)
	}
}
