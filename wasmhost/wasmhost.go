// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"fmt"
	"github.com/mr-tron/base58"
)

const (
	OBJTYPE_ARRAY int32 = 0x20

	OBJTYPE_ADDRESS int32 = 1
	OBJTYPE_AGENT   int32 = 2
	OBJTYPE_BYTES   int32 = 3
	OBJTYPE_COLOR   int32 = 4
	OBJTYPE_INT     int32 = 5
	OBJTYPE_MAP     int32 = 6
	OBJTYPE_STRING  int32 = 7
)

const KeyFromString int32 = 0x4000

type HostObject interface {
	Exists(keyId int32) bool
	GetBytes(keyId int32) []byte
	GetInt(keyId int32) int64
	GetObjectId(keyId int32, typeId int32) int32
	GetString(keyId int32) string
	GetTypeId(keyId int32) int32
	SetBytes(keyId int32, value []byte)
	SetInt(keyId int32, value int64)
	SetString(keyId int32, value string)
}

type LogInterface interface {
	Log(logLevel int32, text string)
}

// implements client.ScHost interface
type WasmHost struct {
	vm            WasmVM
	codeToFunc    map[uint32]string
	error         string
	funcToCode    map[string]uint32
	funcToIndex   map[string]int32
	keyIdToKey    [][]byte
	keyIdToKeyMap [][]byte
	keyToKeyId    map[string]int32
	logger        LogInterface
	objIdToObj    []HostObject
	useBase58Keys bool
}

func (host *WasmHost) Init(null HostObject, root HostObject, logger LogInterface) {
	host.codeToFunc = make(map[uint32]string)
	host.error = ""
	host.funcToCode = make(map[string]uint32)
	host.funcToIndex = make(map[string]int32)
	host.logger = logger
	host.objIdToObj = nil
	host.keyIdToKey = [][]byte{[]byte("<null>")}
	host.keyToKeyId = make(map[string]int32)
	host.keyIdToKeyMap = make([][]byte, len(keyMap)+1)
	for k, v := range keyMap {
		host.keyIdToKeyMap[-v] = []byte(k)
	}
	host.TrackObject(null)
	host.TrackObject(root)
}

func (host *WasmHost) InitVM(vm WasmVM) error {
	return vm.LinkHost(host)
}

func (host *WasmHost) Exists(objId int32, keyId int32) bool {
	return host.FindObject(objId).Exists(keyId)
}

func (host *WasmHost) FindObject(objId int32) HostObject {
	if objId < 0 || objId >= int32(len(host.objIdToObj)) {
		host.SetError("Invalid objId")
		objId = 0
	}
	return host.objIdToObj[objId]
}

func (host *WasmHost) FindSubObject(obj HostObject, keyId int32, typeId int32) HostObject {
	if obj == nil {
		// use root object
		obj = host.FindObject(1)
	}
	return host.FindObject(obj.GetObjectId(keyId, typeId))
}

func (host *WasmHost) GetBytes(objId int32, keyId int32) []byte {
	if host.HasError() {
		return nil
	}
	obj := host.FindObject(objId)
	if !obj.Exists(keyId) {
		host.Trace("GetBytes o%d k%d missing key", objId, keyId)
		return nil
	}
	value := obj.GetBytes(keyId)
	host.Trace("GetBytes o%d k%d = '%s'", objId, keyId, base58.Encode(value))
	return value
}

func (host *WasmHost) GetInt(objId int32, keyId int32) int64 {
	host.TraceAll("GetInt(o%d,k%d)", objId, keyId)
	if keyId == KeyError && objId == 1 {
		if host.HasError() {
			return 1
		}
		return 0
	}
	if host.HasError() {
		return 0
	}
	value := host.FindObject(objId).GetInt(keyId)
	host.Trace("GetInt o%d k%d = %d", objId, keyId, value)
	return value
}

func (host *WasmHost) GetKeyIdFromBytes(bytes []byte) int32 {
	encoded := base58.Encode(bytes)
	if host.useBase58Keys {
		// transform byte slice key into base58 string
		// now all keys are byte slices from strings
		bytes = []byte(encoded)
	}

	keyId := host.getKeyId(bytes, false)
	host.Trace("GetKeyIdFromBytes '%s'=k%d", encoded, keyId)
	return keyId
}

func (host *WasmHost) GetKeyFromId(keyId int32) []byte {
	host.TraceAll("GetKeyFromId(k%d)", keyId)
	key := host.getKeyFromId(keyId)
	if (keyId & KeyFromString) == 0 {
		// originally a byte slice key
		host.Trace("GetKeyFromId k%d='%s'", keyId, base58.Encode(key))
		return key
	}
	// originally a string key
	host.Trace("GetKeyFromId k%d='%s'", keyId, string(key))
	return key
}

func (host *WasmHost) getKeyFromId(keyId int32) []byte {
	// find predefined key
	if keyId < 0 {
		return host.keyIdToKeyMap[-keyId]
	}

	// find user-defined key
	return host.keyIdToKey[keyId & ^KeyFromString]
}

func (host *WasmHost) GetKeyIdFromString(key string) int32 {
	keyId := host.getKeyId([]byte(key), true)
	host.Trace("GetKeyIdFromString '%s'=k%d", key, keyId)
	return keyId
}

func (host *WasmHost) getKeyId(key []byte, fromString bool) int32 {
	// cannot use []byte as key in maps
	// so we will convert to (non-utf8) string
	// most will have started out as string anyway
	keyString := string(key)

	// first check predefined key map
	keyId, ok := keyMap[keyString]
	if ok {
		return keyId
	}

	// check additional user-defined keys
	keyId, ok = host.keyToKeyId[keyString]
	if ok {
		return keyId
	}

	// unknown key, add it to user-defined key map
	keyId = int32(len(host.keyIdToKey))
	if fromString {
		keyId |= KeyFromString
	}
	host.keyToKeyId[keyString] = keyId
	host.keyIdToKey = append(host.keyIdToKey, key)
	return keyId
}

func (host *WasmHost) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	host.TraceAll("GetObjectId(o%d,k%d,t%d)", objId, keyId, typeId)
	if host.HasError() {
		return 0
	}
	subId := host.FindObject(objId).GetObjectId(keyId, typeId)
	host.Trace("GetObjectId o%d k%d t%d = o%d", objId, keyId, typeId, subId)
	return subId
}

func (host *WasmHost) GetString(objId int32, keyId int32) string {
	value := host.getString(objId, keyId)
	if value == nil {
		return ""
	}
	return *value
}

func (host *WasmHost) getString(objId int32, keyId int32) *string {
	// get error string takes precedence over returning error code
	if keyId == KeyError && objId == 1 {
		host.Trace("GetString o%d k%d = '%s'", objId, keyId, host.error)
		return &host.error
	}
	if host.HasError() {
		return nil
	}
	obj := host.FindObject(objId)
	if !obj.Exists(keyId) {
		host.Trace("GetString o%d k%d missing key", objId, keyId)
		return nil
	}
	value := obj.GetString(keyId)
	host.Trace("GetString o%d k%d = '%s'", objId, keyId, value)
	return &value
}

func (host *WasmHost) HasError() bool {
	if host.error != "" {
		host.Trace("HasError")
		return true
	}
	return false
}

func (host *WasmHost) LoadWasm(wasmData []byte) error {
	err := host.vm.LoadWasm(wasmData)
	if err != nil {
		return err
	}
	err = host.vm.RunFunction("on_load")
	if err != nil {
		return err
	}
	host.vm.SaveMemory()
	return nil
}

func (host *WasmHost) RunFunction(functionName string) error {
	return host.vm.RunFunction(functionName)
}

func (host *WasmHost) RunScFunction(functionName string) error {
	index, ok := host.funcToIndex[functionName]
	if !ok {
		return errors.New("unknown SC function name: " + functionName)
	}
	err := host.vm.RunScFunction(index)
	if err == nil && host.error != "" {
		err = errors.New(host.error)
	}
	host.error = ""
	return err
}

func (host *WasmHost) SetBytes(objId int32, keyId int32, bytes []byte) {
	if host.HasError() {
		return
	}
	host.FindObject(objId).SetBytes(keyId, bytes)
	host.Trace("SetBytes o%d k%d v='%s'", objId, keyId, base58.Encode(bytes))

}

func (host *WasmHost) SetError(text string) {
	host.Trace("SetError '%s'", text)
	if !host.HasError() {
		host.error = text
	}
}

func (host *WasmHost) SetInt(objId int32, keyId int32, value int64) {
	host.TraceAll("SetInt(o%d,k%d)", objId, keyId)
	if host.HasError() {
		return
	}
	host.FindObject(objId).SetInt(keyId, value)
	host.Trace("SetInt o%d k%d v=%d", objId, keyId, value)
}

func (host *WasmHost) SetString(objId int32, keyId int32, value string) {
	if objId == 1 {
		// intercept logging keys to prevent final logging of SetBytes itself
		switch keyId {
		case KeyError:
			host.SetError(value)
			return
		case KeyLog, KeyPanic, KeyTrace, KeyTraceAll:
			host.logger.Log(keyId, value)
			return
		}
	}

	if host.HasError() {
		return
	}
	host.FindObject(objId).SetString(keyId, value)
	host.Trace("SetString o%d k%d v='%s'", objId, keyId, value)
}

func (host *WasmHost) Trace(format string, a ...interface{}) {
	host.logger.Log(KeyTrace, fmt.Sprintf(format, a...))
}

func (host *WasmHost) TraceAll(format string, a ...interface{}) {
	host.logger.Log(KeyTraceAll, fmt.Sprintf(format, a...))
}

func (host *WasmHost) TrackObject(obj HostObject) int32 {
	objId := int32(len(host.objIdToObj))
	host.objIdToObj = append(host.objIdToObj, obj)
	return objId
}
