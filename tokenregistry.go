package main

import (
	"github.com/iotaledger/wasplib/client"
)

func main() {
}

type TokenInfo struct {
	supply      int64
	mintedBy    string
	owner       string
	created     int64
	updated     int64
	description string
	userDefined string
}

//export mintSupply
func mintSupply() {
	//ctx := client.NewScContext()
	//TODO
}

//export updateMetadata
func updateMetadata() {
	//ctx := client.NewScContext()
	//TODO
}

//export transferOwnership
func transferOwnership() {
	//ctx := client.NewScContext()
	//TODO
}

func decodeTokenInfo(bytes []byte) *TokenInfo {
	decoder := client.NewBytesDecoder(bytes)
	data := &TokenInfo{}
	data.supply = decoder.Int()
	data.mintedBy = decoder.String()
	data.owner = decoder.String()
	data.created = decoder.Int()
	data.updated = decoder.Int()
	data.description = decoder.String()
	data.userDefined = decoder.String()
	return data
}

func encodeTokenInfo(data *TokenInfo) []byte {
	return client.NewBytesEncoder().
		Int(data.supply).
		String(data.mintedBy).
		String(data.owner).
		Int(data.created).
		Int(data.updated).
		String(data.description).
		String(data.userDefined).
		Data()
}
