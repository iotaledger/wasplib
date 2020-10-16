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

//export onLoad
func onLoadTokenRegistry() {
	exports := client.NewScExports()
	exports.Add("mintSupply")
	exports.Add("updateMetadata")
	exports.Add("transferOwnership")
}

//export mintSupply
func mintSupply() {
	sc := client.NewScContext()
	request := sc.Request()
	color := request.Hash()
	state := sc.State()
	registry := state.GetMap("tr")
	if registry.GetString(color).Value() != "" {
		sc.Log("TokenRegistry: Color already exists")
		return
	}
	params := request.Params()
	token := &TokenInfo{
		supply:      request.Balance(color),
		mintedBy:    request.Address(),
		owner:       request.Address(),
		created:     request.Timestamp(),
		updated:     request.Timestamp(),
		description: params.GetString("dscr").Value(),
		userDefined: params.GetString("ud").Value(),
	}
	if token.supply <= 0 {
		sc.Log("TokenRegistry: Insufficient supply")
		return
	}
	if token.description == "" {
		token.description += "no dscr"
	}
	bytes := encodeTokenInfo(token)
	registry.GetBytes(color).SetValue(bytes)
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
	//sc := client.NewScContext()
	//TODO
}

//export transferOwnership
func transferOwnership() {
	//sc := client.NewScContext()
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

func encodeTokenInfo(token *TokenInfo) []byte {
	return client.NewBytesEncoder().
		Int(token.supply).
		String(token.mintedBy).
		String(token.owner).
		Int(token.created).
		Int(token.updated).
		String(token.description).
		String(token.userDefined).
		Data()
}
