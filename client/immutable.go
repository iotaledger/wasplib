package client

type ScImmutableBytes struct {
	objId int32
	keyId int32
}

func (o ScImmutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableBytesArray struct {
	objId int32
}

func (o ScImmutableBytesArray) GetBytes(index int32) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: index}
}

func (o ScImmutableBytesArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt struct {
	objId int32
	keyId int32
}

func (o ScImmutableInt) Value() int64 {
	return GetInt(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableIntArray struct {
	objId int32
}

func (o ScImmutableIntArray) GetInt(index int32) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: index}
}

func (o ScImmutableIntArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableKeyMap struct {
	objId int32
}

func (o ScImmutableKeyMap) GetBytes(key []byte) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: GetKey(key)}
}

func (o ScImmutableKeyMap) GetBytesArray(key []byte) ScImmutableBytesArray {
	arrId := GetObjectId(o.objId, GetKey(key), OBJTYPE_BYTES_ARRAY)
	return ScImmutableBytesArray{objId: arrId}
}

func (o ScImmutableKeyMap) GetInt(key []byte) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: GetKey(key)}
}

func (o ScImmutableKeyMap) GetIntArray(key []byte) ScImmutableIntArray {
	arrId := GetObjectId(o.objId, GetKey(key), OBJTYPE_INT_ARRAY)
	return ScImmutableIntArray{objId: arrId}
}

func (o ScImmutableKeyMap) GetKeyMap(key []byte) ScImmutableKeyMap {
	mapId := GetObjectId(o.objId, GetKey(key), OBJTYPE_MAP)
	return ScImmutableKeyMap{objId: mapId}
}

func (o ScImmutableKeyMap) GetMap(key []byte) ScImmutableMap {
	mapId := GetObjectId(o.objId, GetKey(key), OBJTYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableKeyMap) GetMaprray(key []byte) ScImmutableMapArray {
	arrId := GetObjectId(o.objId, GetKey(key), OBJTYPE_MAP_ARRAY)
	return ScImmutableMapArray{objId: arrId}
}

func (o ScImmutableKeyMap) GetString(key []byte) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: GetKey(key)}
}

func (o ScImmutableKeyMap) GetStringArray(key []byte) ScImmutableStringArray {
	arrId := GetObjectId(o.objId, GetKey(key), OBJTYPE_STRING_ARRAY)
	return ScImmutableStringArray{objId: arrId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMap struct {
	objId int32
}

func (o ScImmutableMap) GetBytes(key string) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScImmutableMap) GetBytesArray(key string) ScImmutableBytesArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_BYTES_ARRAY)
	return ScImmutableBytesArray{objId: arrId}
}

func (o ScImmutableMap) GetInt(key string) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScImmutableMap) GetIntArray(key string) ScImmutableIntArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_INT_ARRAY)
	return ScImmutableIntArray{objId: arrId}
}

func (o ScImmutableMap) GetKeyMap(key string) ScImmutableKeyMap {
	mapId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_MAP)
	return ScImmutableKeyMap{objId: mapId}
}

func (o ScImmutableMap) GetMap(key string) ScImmutableMap {
	mapId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMap) GetMaprray(key string) ScImmutableMapArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_MAP_ARRAY)
	return ScImmutableMapArray{objId: arrId}
}

func (o ScImmutableMap) GetString(key string) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScImmutableMap) GetStringArray(key string) ScImmutableStringArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), OBJTYPE_STRING_ARRAY)
	return ScImmutableStringArray{objId: arrId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMapArray struct {
	objId int32
}

func (o ScImmutableMapArray) GetKeyMap(index int32) ScImmutableKeyMap {
	mapId := GetObjectId(o.objId, index, OBJTYPE_MAP)
	return ScImmutableKeyMap{objId: mapId}
}

func (o ScImmutableMapArray) GetMap(index int32) ScImmutableMap {
	mapId := GetObjectId(o.objId, index, OBJTYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMapArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableString struct {
	objId int32
	keyId int32
}

func (o ScImmutableString) Value() string {
	return GetString(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableStringArray struct {
	objId int32
}

func (o ScImmutableStringArray) GetString(index int32) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: index}
}

func (o ScImmutableStringArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}
