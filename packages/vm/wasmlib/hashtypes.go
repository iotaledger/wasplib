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

func (o ScAddress) AsAgentID() ScAgentID {
	a := ScAgentID{}
	// agent id is address padded with zeroes
	copy(a.id[:], o.id[:])
	return a
}

func (o ScAddress) Bytes() []byte {
	return o.id[:]
}

func (o ScAddress) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
}

func (o ScAddress) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAgentID struct {
	id [37]byte
}

func NewScAgentID(chainID ScChainID, hContract ScHname) ScAgentID {
	o := ScAgentID{}
	copy(o.id[:], chainID.Bytes())
	copy(o.id[33:], hContract.Bytes())
	return o
}

func NewScAgentIDFromBytes(bytes []byte) ScAgentID {
	o := ScAgentID{}
	if len(bytes) != len(o.id) {
		Panic("invalid agent id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScAgentID) Address() ScAddress {
	a := ScAddress{}
	copy(a.id[:], o.id[:])
	return a
}

func (o ScAgentID) Bytes() []byte {
	return o.id[:]
}

func (o ScAgentID) Hname() ScHname {
	return NewScHnameFromBytes(o.id[33:])
}

func (o ScAgentID) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
}

func (o ScAgentID) IsAddress() bool {
	return o.Address().AsAgentID() == o
}

func (o ScAgentID) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScChainID struct {
	id [33]byte
}

func NewScChainIDFromBytes(bytes []byte) ScChainID {
	o := ScChainID{}
	if len(bytes) != len(o.id) {
		Panic("invalid chain id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScChainID) Address() ScAddress {
	a := ScAddress{}
	copy(a.id[:], o.id[:])
	return a
}

func (o ScChainID) Bytes() []byte {
	return o.id[:]
}

func (o ScChainID) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
}

func (o ScChainID) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	id [32]byte
}

var (
	IOTA = ScColor{}
	MINT = ScColor{}
)

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

func NewScColorFromRequestID(requestID ScRequestID) ScColor {
	o := ScColor{}
	copy(o.id[:], requestID.Bytes())
	return o
}

func (o ScColor) Bytes() []byte {
	return o.id[:]
}

func (o ScColor) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
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

func (o ScHash) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
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

func (hn ScHname) KeyID() Key32 {
	return GetKeyIDFromBytes(hn.Bytes())
}

func (hn ScHname) String() string {
	return strconv.FormatInt(int64(hn), 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestID struct {
	id [34]byte
}

func NewScRequestIDFromBytes(bytes []byte) ScRequestID {
	o := ScRequestID{}
	if len(bytes) != len(o.id) {
		Panic("invalid request id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o ScRequestID) Bytes() []byte {
	return o.id[:]
}

func (o ScRequestID) KeyID() Key32 {
	return GetKeyIDFromBytes(o.Bytes())
}

func (o ScRequestID) String() string {
	return base58Encode(o.id[:])
}
