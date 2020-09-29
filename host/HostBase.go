package host

import (
	"encoding/binary"
	"fmt"
	"github.com/bytecodealliance/wasmtime-go"
	"github.com/iotaledger/wasplib/host/interfaces"
	"github.com/iotaledger/wasplib/host/interfaces/level"
	"strings"
)

var baseKeyMap = map[string]int32{
	"error":     interfaces.KeyError,
	"length":    interfaces.KeyLength,
	"log":       interfaces.KeyLog,
	"trace":     interfaces.KeyTrace,
	"traceHost": interfaces.KeyTraceHost,
}

type LogInterface interface {
	Log(logLevel int, text string)
}

type HostBase struct {
	error         string
	instance      *wasmtime.Instance
	linker        *wasmtime.Linker
	logger        LogInterface
	memory        *wasmtime.Memory
	memoryDirty   bool
	memoryNonZero int
	memoryCopy    []byte
	module        *wasmtime.Module
	store         *wasmtime.Store
	tracker       *Tracker
}

func (h *HostBase) Init(logger LogInterface, root interfaces.HostObject, keyMap *map[string]int32) {
	if keyMap == nil {
		keyMap = &baseKeyMap
	}
	h.error = ""
	h.logger = logger
	h.tracker = NewTracker(keyMap)
	h.AddObject(NewNullObject(h))
	h.AddObject(root)
	h.initWasmTime()
}

func (h *HostBase) initWasmTime() error {
	var externals = map[string]interface{}{
		"wasplib.hostGetBytes": func(objId int32, keyId int32, stringRef int32, size int32) int32 {
			if objId >= 0 {
				value := h.GetBytes(objId, keyId)
				return h.wasmSetBytes(stringRef, size, value)
			}
			value := []byte(h.GetString(-objId, keyId))
			return h.wasmSetBytes(stringRef, size, value)
		},
		"wasplib.hostGetInt": func(objId int32, keyId int32) int64 {
			return h.GetInt(objId, keyId)
		},
		"wasplib.hostGetIntRef": func(objId int32, keyId int32, intRef int32) {
			value := h.GetInt(objId, keyId)
			h.wasmSetInt(intRef, value)
		},
		"wasplib.hostGetKeyId": func(keyRef int32, size int32) int32 {
			key := h.wasmGetBytes(keyRef, size)
			return h.GetKeyId(string(key))
		},
		"wasplib.hostGetObjectId": func(objId int32, keyId int32, typeId int32) int32 {
			return h.GetObjectId(objId, keyId, typeId)
		},
		"wasplib.hostSetBytes": func(objId int32, keyId int32, stringRef int32, size int32) {
			if objId >= 0 {
				value := h.wasmGetBytes(stringRef, size)
				h.SetBytes(objId, keyId, value)
				return
			}
			value := string(h.wasmGetBytes(stringRef, size))
			h.SetString(-objId, keyId, value)
		},
		"wasplib.hostSetInt": func(objId int32, keyId int32, value int64) {
			h.SetInt(objId, keyId, value)
		},
		"wasplib.hostSetIntRef": func(objId int32, keyId int32, intRef int32) {
			value := h.wasmGetInt(intRef)
			h.SetInt(objId, keyId, value)
		},
		//TODO: go implementation uses this one to write panic message
		"wasi_unstable.fd_write": func(fd int32, iovs int32, size int32, written int32) int32 {
			return h.FdWrite(fd, iovs, size, written)
		},
	}

	h.store = wasmtime.NewStore(wasmtime.NewEngine())
	h.linker = wasmtime.NewLinker(h.store)
	for name, function := range externals {
		names := strings.Split(name, ".")
		err := h.linker.DefineFunc(names[0], names[1], function)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HostBase) LoadWasm(wasmFile string) error {
	var err error
	h.module, err = wasmtime.NewModuleFromFile(h.store.Engine, wasmFile)
	if err != nil {
		return err
	}
	//TODO: Does this instantiate fresh memory instance or only link externals?
	//      Same question for start() function. We need a *clean* instance!
	h.instance, err = h.linker.Instantiate(h.module)
	if err != nil {
		return err
	}

	// find initialized data range in memory
	h.memory = h.instance.GetExport("memory").Memory()
	ptr := h.memory.UnsafeData()
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
	h.memoryNonZero = len(ptr)
	if ptr[firstNonZero] != 0 {
		h.memoryNonZero = firstNonZero
		size := lastNonZero + 1 - firstNonZero
		h.memoryCopy = make([]byte, size, size)
		copy(h.memoryCopy, ptr[h.memoryNonZero:])
	}
	return nil
}

func (h *HostBase) Log(logLevel int, text string) {
	if logLevel >= level.MSG {
		fmt.Println(text)
	}
}

func (h *HostBase) AddObject(obj interfaces.HostObject) int32 {
	return h.tracker.AddObject(obj)
}

func (h *HostBase) FdWrite(fd int32, iovs int32, size int32, written int32) int32 {
	ptr := h.memory.UnsafeData()
	txt := binary.LittleEndian.Uint32(ptr[iovs : iovs+4])
	siz := binary.LittleEndian.Uint32(ptr[iovs+4 : iovs+8])
	fmt.Print(string(ptr[txt : txt+siz]))
	binary.LittleEndian.PutUint32(ptr[written:written+4], siz)
	return int32(siz)
}

func (h *HostBase) GetBytes(objId int32, keyId int32) []byte {
	if h.HasError() {
		return []byte(nil)
	}
	value := h.GetObject(objId).GetBytes(keyId)
	h.Logf("GetBytes o%d k%d = '%v'", objId, keyId, value)
	return value
}

func (h *HostBase) GetInt(objId int32, keyId int32) int64 {
	if keyId == interfaces.KeyError && objId == 1 {
		if h.HasError() {
			return 1
		}
		return 0
	}
	if h.HasError() {
		return 0
	}
	value := h.GetObject(objId).GetInt(keyId)
	h.Logf("GetInt o%d k%d = %d", objId, keyId, value)
	return value
}

func (h *HostBase) GetKey(keyId int32) string {
	key := h.tracker.GetKey(keyId)
	h.Logf("GetKey k%d='%s'", keyId, key)
	return key
}

func (h *HostBase) GetKeyId(key string) int32 {
	keyId := h.tracker.GetKeyId(key)
	h.Logf("GetKeyId '%s'=k%d", key, keyId)
	return keyId
}

func (h *HostBase) GetObject(objId int32) interfaces.HostObject {
	o := h.tracker.GetObject(objId)
	if o == nil {
		h.SetError("Invalid objId")
		return NewNullObject(h)
	}
	return o
}

func (h *HostBase) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	if h.HasError() {
		return 0
	}
	subId := h.GetObject(objId).GetObjectId(keyId, typeId)
	h.Logf("GetObjectId o%d k%d t%d = o%d", objId, keyId, typeId, subId)
	return subId
}

func (h *HostBase) GetString(objId int32, keyId int32) string {
	if keyId == interfaces.KeyError && objId == 1 {
		return h.error
	}
	if h.HasError() {
		return ""
	}
	value := h.GetObject(objId).GetString(keyId)
	h.Logf("GetString o%d k%d = '%s'", objId, keyId, value)
	return value
}

func (h *HostBase) HasError() bool {
	return h.error != ""
}

func (h *HostBase) Logf(format string, a ...interface{}) {
	h.logger.Log(level.TRACE, fmt.Sprintf(format, a...))
}

func (h *HostBase) RunWasmFunction(functionName string) error {
	if h.memoryDirty {
		// clear memory and restore initialized data range
		ptr := h.memory.UnsafeData()
		size := len(ptr)
		copy(ptr, make([]byte, size, size))
		copy(ptr[h.memoryNonZero:], h.memoryCopy)
	}
	h.memoryDirty = true
	function := h.instance.GetExport(functionName).Func()
	_, err := function.Call()
	return err
}

func (h *HostBase) SetBytes(objId int32, keyId int32, value []byte) {
	if h.HasError() {
		return
	}
	h.GetObject(objId).SetBytes(keyId, value)
	h.Logf("SetBytes o%d k%d v='%v'", objId, keyId, value)
}

func (h *HostBase) SetError(text string) {
	h.Logf("SetError '%s'", text)
	if !h.HasError() {
		h.error = text
	}
}

func (h *HostBase) SetInt(objId int32, keyId int32, value int64) {
	if h.HasError() {
		return
	}
	h.GetObject(objId).SetInt(keyId, value)
	h.Logf("SetInt o%d k%d v=%d", objId, keyId, value)
}

func (h *HostBase) SetString(objId int32, keyId int32, value string) {
	if objId == 1 {
		// intercept logging keys to prevent final logging of SetBytes itself
		switch keyId {
		case interfaces.KeyError:
			h.SetError(value)
			return
		case interfaces.KeyLog:
			h.logger.Log(level.MSG, value)
			return
		case interfaces.KeyTrace:
			h.logger.Log(level.TRACE, value)
			return
		case interfaces.KeyTraceHost:
			h.logger.Log(level.HOST, value)
			return
		}
	}
	if h.HasError() {
		return
	}
	h.GetObject(objId).SetString(keyId, value)
	h.Logf("SetString o%d k%d v='%s'", objId, keyId, value)
}

func (h *HostBase) wasmGetBytes(offset int32, size int32) []byte {
	ptr := h.memory.UnsafeData()
	bytes := make([]byte, size)
	copy(bytes, ptr[offset:offset+size])
	return bytes
}

func (h *HostBase) wasmGetInt(offset int32) int64 {
	ptr := h.memory.UnsafeData()
	return int64(binary.LittleEndian.Uint64(ptr[offset : offset+8]))
}

func (h *HostBase) wasmSetBytes(offset int32, size int32, value []byte) int32 {
	bytes := []byte(value)
	if size != 0 {
		ptr := h.memory.UnsafeData()
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}

func (h *HostBase) wasmSetInt(offset int32, value int64) {
	ptr := h.memory.UnsafeData()
	binary.LittleEndian.PutUint64(ptr[offset:offset+8], uint64(value))
}
