package host

import "github.com/iotaledger/wasplib/host/interfaces"

type Tracker struct {
	keyMap        *map[string]int32
	keyToKeyId    map[string]int32
	keyIdToKey    []string
	keyIdToKeyMap []string
	objIdToObj    []interfaces.HostObject
}

func NewTracker(keyMap *map[string]int32) *Tracker {
	elements := len(*keyMap) + 1
	keyIdToKeyMap := make([]string, elements, elements)
	for k, v := range *keyMap {
		keyIdToKeyMap[-v] = k
	}
	return &Tracker{
		keyMap:        keyMap,
		keyToKeyId:    make(map[string]int32),
		keyIdToKey:    []string{"<null>"},
		keyIdToKeyMap: keyIdToKeyMap,
		objIdToObj:    []interfaces.HostObject{},
	}
}

func (t *Tracker) AddObject(obj interfaces.HostObject) int32 {
	objId := int32(len(t.objIdToObj))
	t.objIdToObj = append(t.objIdToObj, obj)
	return objId
}

func (t *Tracker) GetKey(keyId int32) string {
	if keyId < 0 {
		return t.keyIdToKeyMap[-keyId]
	}
	if keyId < int32(len(t.keyIdToKey)) {
		return t.keyIdToKey[keyId]
	}
	return ""
}

func (t *Tracker) GetKeyId(key string) int32 {
	keyId, ok := (*t.keyMap)[key]
	if ok {
		return keyId
	}
	keyId, ok = t.keyToKeyId[key]
	if ok {
		return keyId
	}
	keyId = int32(len(t.keyIdToKey))
	t.keyToKeyId[key] = keyId
	t.keyIdToKey = append(t.keyIdToKey, key)
	return keyId
}

func (t *Tracker) GetObject(objId int32) interfaces.HostObject {
	if objId < 0 || objId >= int32(len(t.objIdToObj)) {
		return nil
	}
	return t.objIdToObj[objId]
}
