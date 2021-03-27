// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

import (
	"encoding/binary"
	"strconv"
)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAddress struct {
	id [33]byte
}

func NewScAddressFromBytes(bytes []byte) ScAddress {
	o := ScAddress{}
	if len(bytes) != len(o.id) {
		Panic("invalid address id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScAddress) AsAgentId() ScAgentId {
	a := ScAgentId{}
	// agent id is address padded with zeroes
	copy(a.id[:], o.id[:])
	return a
}

func (o ScAddress) Bytes() []byte {
	return o.id[:]
}

func (o ScAddress) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScAddress) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAgentId struct {
	id [37]byte
}

func NewScAgentId(chainId ScChainId, hContract ScHname) ScAgentId {
	o := ScAgentId{}
	copy(o.id[:], chainId.Bytes())
	copy(o.id[33:], hContract.Bytes())
	return o
}

func NewScAgentIdFromBytes(bytes []byte) ScAgentId {
	o := ScAgentId{}
	if len(bytes) != len(o.id) {
		Panic("invalid agent id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScAgentId) Address() ScAddress {
	a := ScAddress{}
	copy(a.id[:], o.id[:])
	return a
}

func (o ScAgentId) Bytes() []byte {
	return o.id[:]
}

func (o ScAgentId) Hname() ScHname {
	return NewScHnameFromBytes(o.id[33:])
}

func (o ScAgentId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScAgentId) IsAddress() bool {
	return o.Address().AsAgentId() == o
}

func (o ScAgentId) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScChainId struct {
	id [33]byte
}

func NewScChainIdFromBytes(bytes []byte) ScChainId {
	o := ScChainId{}
	if len(bytes) != len(o.id) {
		Panic("invalid chain id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScChainId) Bytes() []byte {
	return o.id[:]
}

func (o ScChainId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScChainId) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	id [32]byte
}

var IOTA = ScColor{}
var MINT = ScColor{}

func init() {
	for i := range MINT.id {
		MINT.id[i] = 0xff
	}
}

func NewScColorFromBytes(bytes []byte) ScColor {
	o := ScColor{}
	if len(bytes) != len(o.id) {
		Panic("invalid color id length")
	}
	copy(o.id[:], bytes)
	return o
}

func NewScColorFromRequestId(requestId ScRequestId) ScColor {
	o := ScColor{}
	copy(o.id[:], requestId.Bytes())
	return o
}

func (o ScColor) Bytes() []byte {
	return o.id[:]
}

func (o ScColor) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScColor) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScHash struct {
	id [32]byte
}

func NewScHashFromBytes(bytes []byte) ScHash {
	o := ScHash{}
	if len(bytes) != len(o.id) {
		Panic("invalid hash id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScHash) Bytes() []byte {
	return o.id[:]
}

func (o ScHash) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScHash) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScHname uint32

func NewScHname(name string) ScHname {
	return ScFuncContext{}.Utility().Hname(name)
}

func NewScHnameFromBytes(bytes []byte) ScHname {
	return ScHname(binary.LittleEndian.Uint32(bytes))
}

func (hn ScHname) Bytes() []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(hn))
	return bytes
}

func (hn ScHname) KeyId() Key32 {
	return GetKeyIdFromBytes(hn.Bytes())
}

func (hn ScHname) String() string {
	return strconv.FormatInt(int64(hn), 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestId struct {
	id [34]byte
}

func NewScRequestIdFromBytes(bytes []byte) ScRequestId {
	o := ScRequestId{}
	if len(bytes) != len(o.id) {
		Panic("invalid request id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScRequestId) Bytes() []byte {
	return o.id[:]
}

func (o ScRequestId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o ScRequestId) String() string {
	return base58Encode(o.id[:])
}
