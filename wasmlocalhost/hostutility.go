// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"encoding/binary"
	"fmt"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/mr-tron/base58"
)

type HostUtility struct {
	HostMap
	base58Decoded []byte
	base58Encoded string
	hash          []byte
	hname         coretypes.Hname
	nextRandom    int
	random        []byte
}

func NewHostUtility(host *SimpleWasmHost, keyId int32) *HostUtility {
	o := &HostUtility{HostMap: *NewHostMap(host, keyId)}
	// preset randomizer to generate sequence 1..8 before
	// continuing with proper hashed values
	o.random = make([]byte, 8*8)
	for i := 0; i < len(o.random); i += 8 {
		o.random[i] = byte(i + 1)
	}
	return o
}

func (o *HostUtility) Exists(keyId int32) bool {
	switch keyId {
	case wasmhost.KeyBase58:
	case wasmhost.KeyHash:
	case wasmhost.KeyName:
	case wasmhost.KeyRandom:
	default:
		return false
	}
	return true
}

func (o *HostUtility) GetBytes(keyId int32) []byte {
	switch keyId {
	case wasmhost.KeyName:
		return codec.EncodeHname(o.hname)
	case wasmhost.KeyBase58:
		return o.base58Decoded
	case wasmhost.KeyHash:
		return o.hash
	}
	o.invalidKey(keyId)
	return nil
}

func (o *HostUtility) GetInt(keyId int32) int64 {
	switch keyId {
	case wasmhost.KeyRandom:
		i := o.nextRandom
		if i+8 > len(o.random) {
			// not enough bytes left, generate more bytes
			h := hashing.HashData(o.random)
			o.random = h[:]
			i = 0
		}
		o.nextRandom = i + 8
		return int64(binary.LittleEndian.Uint64(o.random[i : i+8]))
	}
	o.invalidKey(keyId)
	return 0
}

func (o *HostUtility) GetString(keyId int32) string {
	switch keyId {
	case wasmhost.KeyBase58:
		return o.base58Encoded
	}
	o.invalidKey(keyId)
	return ""
}

func (o *HostUtility) GetTypeId(keyId int32) int32 {
	switch keyId {
	case wasmhost.KeyBase58:
		return wasmhost.OBJTYPE_BYTES
	case wasmhost.KeyHash:
		return wasmhost.OBJTYPE_BYTES //TODO OBJTYPE_HASH
	case wasmhost.KeyRandom:
		return wasmhost.OBJTYPE_INT
	}
	return 0
}

func (o *HostUtility) invalidKey(keyId int32) {
	panic(fmt.Sprintf("Invalid key: %d", keyId))
}

func (o *HostUtility) SetBytes(keyId int32, value []byte) {
	switch keyId {
	case wasmhost.KeyBase58:
		o.base58Encoded = base58.Encode(value)
	case wasmhost.KeyHash:
		h := hashing.HashData(value)
		o.hash = h[:]
	default:
		o.invalidKey(keyId)
	}
}

func (o *HostUtility) SetString(keyId int32, value string) {
	switch keyId {
	case wasmhost.KeyName:
		o.hname = coretypes.Hn(value)
	case wasmhost.KeyBase58:
		o.base58Decoded, _ = base58.Decode(value)
	default:
		o.invalidKey(keyId)
	}
}
