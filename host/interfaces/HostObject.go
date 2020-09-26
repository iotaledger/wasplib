package interfaces

type HostObject interface {
	GetBytes(keyId int32) []byte
	GetInt(keyId int32) int64
	GetObjectId(keyId int32, typeId int32) int32
	GetString(keyId int32) string
	SetBytes(keyId int32, value []byte)
	SetInt(keyId int32, value int64)
	SetString(keyId int32, value string)
}
