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

func NewScAddressFromBytes(bytes []byte) *ScAddress {
	o := &ScAddress{}
	if len(bytes) != len(o.id) {
		logPanic("invalid address id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScAddress) AsAgentId() *ScAgentId {
	a := &ScAgentId{}
	// agent is address padded with zeroes
	copy(a.id[:], o.id[:])
	return a
}

func (o *ScAddress) Bytes() []byte {
	return o.id[:]
}

func (o *ScAddress) Equals(other *ScAddress) bool {
	return o.id == other.id
}

func (o *ScAddress) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScAddress) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAgentId struct {
	id [37]byte
}

func NewScAgentIdFromBytes(bytes []byte) *ScAgentId {
	o := &ScAgentId{}
	if len(bytes) != len(o.id) {
		logPanic("invalid agent id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScAgentId) Address() *ScAddress {
	a := &ScAddress{}
	copy(a.id[:], o.id[:])
	return a
}

func (o *ScAgentId) Bytes() []byte {
	return o.id[:]
}

func (o *ScAgentId) Equals(other *ScAgentId) bool {
	return o.id == other.id
}

func (o *ScAgentId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScAgentId) IsAddress() bool {
	return o.Address().AsAgentId().Equals(o)
}

func (o *ScAgentId) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScChainId struct {
	id [33]byte
}

func NewScChainIdFromBytes(bytes []byte) *ScChainId {
	o := &ScChainId{}
	if len(bytes) != len(o.id) {
		logPanic("invalid chain id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScChainId) Bytes() []byte {
	return o.id[:]
}

func (o *ScChainId) Equals(other *ScChainId) bool {
	return o.id == other.id
}

func (o *ScChainId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScChainId) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	id [32]byte
}

var IOTA = &ScColor{}
var MINT = &ScColor{}

func init() {
	for i := range MINT.id {
		MINT.id[i] = 0xff
	}
}

func NewScColorFromBytes(bytes []byte) *ScColor {
	o := &ScColor{}
	if len(bytes) != len(o.id) {
		logPanic("invalid color id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScColor) Bytes() []byte {
	return o.id[:]
}

func (o *ScColor) Equals(other *ScColor) bool {
	return o.id == other.id
}

func (o *ScColor) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScColor) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContractId struct {
	id [37]byte
}

func NewScContractId(chainId *ScChainId, hContract ScHname) *ScContractId {
	o := &ScContractId{}
	copy(o.id[:], chainId.Bytes())
	copy(o.id[33:], hContract.Bytes())
	return o
}

func NewScContractIdFromBytes(bytes []byte) *ScContractId {
	o := &ScContractId{}
	if len(bytes) != len(o.id) {
		logPanic("invalid contract id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScContractId) AsAgentId() *ScAgentId {
	a := &ScAgentId{}
	copy(a.id[:], o.id[:])
	return a
}

func (o *ScContractId) Bytes() []byte {
	return o.id[:]
}

func (o *ScContractId) ChainId() *ScChainId {
	c := &ScChainId{}
	copy(c.id[:], o.id[:])
	return c
}

func (o *ScContractId) Equals(other *ScContractId) bool {
	return o.id == other.id
}

func (o *ScContractId) Hname() ScHname {
	return NewScHnameFromBytes(o.id[33:])
}

func (o *ScContractId) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScContractId) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScHash struct {
	id [32]byte
}

func NewScHashFromBytes(bytes []byte) *ScHash {
	o := &ScHash{}
	if len(bytes) != len(o.id) {
		logPanic("invalid hash id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScHash) Bytes() []byte {
	return o.id[:]
}

func (o *ScHash) Equals(other *ScHash) bool {
	return o.id == other.id
}

func (o *ScHash) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScHash) String() string {
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

func (hn ScHname) Equals(other ScHname) bool {
	return hn == other
}

func (hn ScHname) KeyId() Key32 {
	return GetKeyIdFromBytes(hn.Bytes())
}

func (hn ScHname) String() string {
	return strconv.FormatInt(int64(hn), 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func logPanic(text string) {
	ScBaseContext{}.Panic(text)
}
