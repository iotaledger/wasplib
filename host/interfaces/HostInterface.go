package interfaces

const (
	KeyError       = int32(-1)
	KeyLength      = KeyError - 1
	KeyLog         = KeyLength - 1
	KeyTrace       = KeyLog - 1
	KeyTraceHost   = KeyTrace - 1
	KeyUserDefined = KeyTraceHost - 1
)

type HostInterface interface {
	GetBytes(objId int32, keyId int32) []byte
	GetInt(objId int32, keyId int32) int64
	GetKeyId(key string) int32
	GetObjectId(objId int32, keyId int32, typeId int32) int32
	SetBytes(objId int32, keyId int32, value []byte)
	SetInt(objId int32, keyId int32, value int64)
}
