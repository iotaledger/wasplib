package main

import (
	"github.com/iotaledger/wasplib/client"
)

type TokenInfo struct {
	supply      int64
	mintedBy    string
	owner       string
	created     int64
	updated     int64
	description string
	userDefined string
}

func main() {
}

//export mintSupply
func mintSupply() {
	ctx := client.NewScContext()
	request := ctx.Request()
	color := request.Hash()
	state := ctx.State()
	registry := state.GetMap("tr")
	if registry.GetString(color).Value() != "" {
		ctx.Log("TokenRegistry: Color already exists")
		return
	}
	reqParams := request.Params()
	token := &TokenInfo{}
	token.supply = request.Balance(color)
	token.mintedBy = request.Address()
	token.owner = request.Address()
	token.created = request.Timestamp()
	token.updated = request.Timestamp()
	token.description = reqParams.GetString("dscr").Value()
	token.userDefined = reqParams.GetString("ud").Value()
	if token.supply <= 0 {
		ctx.Log("TokenRegistry: Insufficient supply")
		return
	}
	if token.description == "" {
		token.description += "no dscr"
	}
	data := encodeTokenInfo(token)
	registry.GetBytes(color).SetValue(data)
	colors := state.GetString("lc")
	list := colors.Value()
	if list != "" {
		list += ","
	}
	list += color
	colors.SetValue(list)
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
