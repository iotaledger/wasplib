package wasmhost

import (
	"encoding/binary"
	"fmt"
	"github.com/bytecodealliance/wasmtime-go"
	"strings"
)

const (
	OBJTYPE_BYTES        int32 = 0
	OBJTYPE_BYTES_ARRAY  int32 = 1
	OBJTYPE_INT          int32 = 2
	OBJTYPE_INT_ARRAY    int32 = 3
	OBJTYPE_MAP          int32 = 4
	OBJTYPE_MAP_ARRAY    int32 = 5
	OBJTYPE_STRING       int32 = 6
	OBJTYPE_STRING_ARRAY int32 = 7
)

const (
	KeyError       = int32(-1)
	KeyLength      = KeyError - 1
	KeyLog         = KeyLength - 1
	KeyTrace       = KeyLog - 1
	KeyTraceHost   = KeyTrace - 1
	KeyWarning     = KeyTraceHost - 1
	KeyUserDefined = KeyWarning - 1
)

type HostObject interface {
	GetBytes(keyId int32) []byte
	GetInt(keyId int32) int64
	GetObjectId(keyId int32, typeId int32) int32
	GetString(keyId int32) string
	SetBytes(keyId int32, value []byte)
	SetInt(keyId int32, value int64)
	SetString(keyId int32, value string)
}

type LogInterface interface {
	Log(logLevel int32, text string)
}

var baseKeyMap = map[string]int32{
	"error":     KeyError,
	"length":    KeyLength,
	"log":       KeyLog,
	"trace":     KeyTrace,
	"traceHost": KeyTraceHost,
	"warning":   KeyWarning,
}

type WasmHost struct {
	error         string
	instance      *wasmtime.Instance
	keyIdToKey    []string
	keyIdToKeyMap []string
	keyMapToKeyId *map[string]int32
	keyToKeyId    map[string]int32
	linker        *wasmtime.Linker
	logger        LogInterface
	memory        *wasmtime.Memory
	memoryCopy    []byte
	memoryDirty   bool
	memoryNonZero int
	module        *wasmtime.Module
	objIdToObj    []HostObject
	store         *wasmtime.Store
}

func (host *WasmHost) Init(root HostObject, keyMap *map[string]int32, logger LogInterface) {
	if keyMap == nil {
		keyMap = &baseKeyMap
	}
	elements := len(*keyMap) + 1
	host.error = ""
	host.logger = logger
	host.keyIdToKey = []string{"<null>"}
	host.keyMapToKeyId = keyMap
	host.keyToKeyId = make(map[string]int32)
	host.keyIdToKeyMap = make([]string, elements, elements)
	for k, v := range *keyMap {
		host.keyIdToKeyMap[-v] = k
	}
	host.TrackObject(NewNullObject(host))
	host.TrackObject(root)
	host.initWasm()
}

func (host *WasmHost) initWasm() error {
	var externals = map[string]interface{}{
		"wasplib.hostGetBytes": func(objId int32, keyId int32, stringRef int32, size int32) int32 {
			if objId >= 0 {
				return host.wasmSetBytes(stringRef, size, host.GetBytes(objId, keyId))
			}
			return host.wasmSetBytes(stringRef, size, []byte(host.GetString(-objId, keyId)))
		},
		"wasplib.hostGetInt": func(objId int32, keyId int32) int64 {
			return host.GetInt(objId, keyId)
		},
		"wasplib.hostGetIntRef": func(objId int32, keyId int32, intRef int32) {
			host.wasmSetInt(intRef, host.GetInt(objId, keyId))
		},
		"wasplib.hostGetKeyId": func(keyRef int32, size int32) int32 {
			return host.GetKeyId(string(host.wasmGetBytes(keyRef, size)))
		},
		"wasplib.hostGetObjectId": func(objId int32, keyId int32, typeId int32) int32 {
			return host.GetObjectId(objId, keyId, typeId)
		},
		"wasplib.hostSetBytes": func(objId int32, keyId int32, stringRef int32, size int32) {
			if objId >= 0 {
				host.SetBytes(objId, keyId, host.wasmGetBytes(stringRef, size))
				return
			}
			host.SetString(-objId, keyId, string(host.wasmGetBytes(stringRef, size)))
		},
		"wasplib.hostSetInt": func(objId int32, keyId int32, value int64) {
			host.SetInt(objId, keyId, value)
		},
		"wasplib.hostSetIntRef": func(objId int32, keyId int32, intRef int32) {
			host.SetInt(objId, keyId, host.wasmGetInt(intRef))
		},
		//TODO: go implementation uses this one to write panic message
		"wasi_unstable.fd_write": func(fd int32, iovs int32, size int32, written int32) int32 {
			return host.fdWrite(fd, iovs, size, written)
		},
	}

	host.store = wasmtime.NewStore(wasmtime.NewEngine())
	host.linker = wasmtime.NewLinker(host.store)
	for name, function := range externals {
		names := strings.Split(name, ".")
		err := host.linker.DefineFunc(names[0], names[1], function)
		if err != nil {
			return err
		}
	}
	return nil
}

func (host *WasmHost) fdWrite(fd int32, iovs int32, size int32, written int32) int32 {
	// very basic implementation that expects fd to be stdout and iovs to be only one element
	ptr := host.memory.UnsafeData()
	txt := binary.LittleEndian.Uint32(ptr[iovs : iovs+4])
	siz := binary.LittleEndian.Uint32(ptr[iovs+4 : iovs+8])
	fmt.Print(string(ptr[txt : txt+siz]))
	binary.LittleEndian.PutUint32(ptr[written:written+4], siz)
	return int32(siz)
}

func (host *WasmHost) FindObject(objId int32) HostObject {
	if objId < 0 || objId >= int32(len(host.objIdToObj)) {
		host.SetError("Invalid objId")
		return NewNullObject(host)
	}
	return host.objIdToObj[objId]
}

func (host *WasmHost) GetBytes(objId int32, keyId int32) []byte {
	if host.HasError() {
		return []byte(nil)
	}
	value := host.FindObject(objId).GetBytes(keyId)
	host.Trace("GetBytes o%d k%d = '%v'", objId, keyId, value)
	return value
}

func (host *WasmHost) GetInt(objId int32, keyId int32) int64 {
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

func (host *WasmHost) GetKey(keyId int32) string {
	key := host.getKey(keyId)
	host.Trace("GetKey k%d='%s'", keyId, key)
	return key
}

func (host *WasmHost) getKey(keyId int32) string {
	// find predefined key
	if keyId < 0 {
		return host.keyIdToKeyMap[-keyId]
	}

	// find user-defined key
	if keyId < int32(len(host.keyIdToKey)) {
		return host.keyIdToKey[keyId]
	}

	// unknown key
	return ""
}

func (host *WasmHost) GetKeyId(key string) int32 {
	keyId := host.getKeyId(key)
	host.Trace("GetKeyId '%s'=k%d", key, keyId)
	return keyId
}

func (host *WasmHost) getKeyId(key string) int32 {
	// first check predefined key map
	keyId, ok := (*host.keyMapToKeyId)[key]
	if ok {
		return keyId
	}

	// check additional user-defined keys
	keyId, ok = host.keyToKeyId[key]
	if ok {
		return keyId
	}

	// unknown key, add it to user-defined key map
	keyId = int32(len(host.keyIdToKey))
	host.keyToKeyId[key] = keyId
	host.keyIdToKey = append(host.keyIdToKey, key)
	return keyId
}

func (host *WasmHost) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	if host.HasError() {
		return 0
	}
	subId := host.FindObject(objId).GetObjectId(keyId, typeId)
	host.Trace("GetObjectId o%d k%d t%d = o%d", objId, keyId, typeId, subId)
	return subId
}

func (host *WasmHost) GetString(objId int32, keyId int32) string {
	if keyId == KeyError && objId == 1 {
		return host.error
	}
	if host.HasError() {
		return ""
	}
	value := host.FindObject(objId).GetString(keyId)
	host.Trace("GetString o%d k%d = '%s'", objId, keyId, value)
	return value
}

func (host *WasmHost) HasError() bool {
	return host.error != ""
}

func (host *WasmHost) LoadWasm(wasmFile string) error {
	var err error
	host.module, err = wasmtime.NewModuleFromFile(host.store.Engine, wasmFile)
	if err != nil {
		return err
	}
	//TODO: Does this instantiate fresh memory instance or only link externals?
	//      Same question for start() function. We need a *clean* instance!
	host.instance, err = host.linker.Instantiate(host.module)
	if err != nil {
		return err
	}

	// find initialized data range in memory
	host.memory = host.instance.GetExport("memory").Memory()
	ptr := host.memory.UnsafeData()
	firstNonZero := 0
	lastNonZero := 0
	for i, b := range ptr {
		if b != 0 {
			if firstNonZero == 0 {
				firstNonZero = i
			}
			lastNonZero = i
		}
	}

	// save copy of initialized data range
	host.memoryNonZero = len(ptr)
	if ptr[firstNonZero] != 0 {
		host.memoryNonZero = firstNonZero
		size := lastNonZero + 1 - firstNonZero
		host.memoryCopy = make([]byte, size, size)
		copy(host.memoryCopy, ptr[host.memoryNonZero:])
	}
	return nil
}

func (host *WasmHost) RunWasmFunction(functionName string) error {
	if host.memoryDirty {
		// clear memory and restore initialized data range
		ptr := host.memory.UnsafeData()
		size := len(ptr)
		copy(ptr, make([]byte, size, size))
		copy(ptr[host.memoryNonZero:], host.memoryCopy)
	}
	host.memoryDirty = true
	function := host.instance.GetExport(functionName).Func()
	_, err := function.Call()
	return err
}

func (host *WasmHost) SetBytes(objId int32, keyId int32, value []byte) {
	if host.HasError() {
		return
	}
	host.FindObject(objId).SetBytes(keyId, value)
	host.Trace("SetBytes o%d k%d v='%v'", objId, keyId, value)
}

func (host *WasmHost) SetError(text string) {
	host.Trace("SetError '%s'", text)
	if !host.HasError() {
		host.error = text
	}
}

func (host *WasmHost) SetInt(objId int32, keyId int32, value int64) {
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
		case KeyLog, KeyTrace, KeyTraceHost:
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

func (host *WasmHost) TrackObject(obj HostObject) int32 {
	objId := int32(len(host.objIdToObj))
	host.objIdToObj = append(host.objIdToObj, obj)
	return objId
}

func (host *WasmHost) wasmGetBytes(offset int32, size int32) []byte {
	ptr := host.memory.UnsafeData()
	bytes := make([]byte, size)
	copy(bytes, ptr[offset:offset+size])
	return bytes
}

func (host *WasmHost) wasmGetInt(offset int32) int64 {
	ptr := host.memory.UnsafeData()
	return int64(binary.LittleEndian.Uint64(ptr[offset : offset+8]))
}

func (host *WasmHost) wasmSetBytes(offset int32, size int32, value []byte) int32 {
	bytes := []byte(value)
	if size != 0 {
		ptr := host.memory.UnsafeData()
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}

func (host *WasmHost) wasmSetInt(offset int32, value int64) {
	ptr := host.memory.UnsafeData()
	binary.LittleEndian.PutUint64(ptr[offset:offset+8], uint64(value))
}
