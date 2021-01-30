// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/binary"
	"strconv"
)

type Hname uint32

func NewHname(name string) Hname {
	return ScCallContext{}.Utility().Hname(name)
}

func NewHnameFromBytes(bytes []byte) Hname {
	return Hname(binary.LittleEndian.Uint32(bytes))
}

func (hn Hname) Bytes() []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(hn))
	return bytes
}

func (hn Hname) Equals(other Hname) bool {
	return hn == other
}

func (hn Hname) KeyId() Key32 {
	return GetKeyIdFromBytes(hn.Bytes())
}

func (hn Hname) String() string {
	return strconv.FormatInt(int64(hn), 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAddress struct {
	id [33]byte
}

func NewScAddress(bytes []byte) *ScAddress {
	o := &ScAddress{}
	if len(bytes) != len(o.id) {
		logPanic("invalid address id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScAddress) AsAgent() *ScAgent {
	a := &ScAgent{}
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

type ScAgent struct {
	id [37]byte
}

func NewScAgent(bytes []byte) *ScAgent {
	o := &ScAgent{}
	if len(bytes) != len(o.id) {
		logPanic("invalid agent id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScAgent) Address() *ScAddress {
	a := &ScAddress{}
	copy(a.id[:], o.id[:])
	return a
}

func (o *ScAgent) Bytes() []byte {
	return o.id[:]
}

func (o *ScAgent) Equals(other *ScAgent) bool {
	return o.id == other.id
}

func (o *ScAgent) KeyId() Key32 {
	return GetKeyIdFromBytes(o.Bytes())
}

func (o *ScAgent) IsAddress() bool {
	return o.Address().AsAgent().Equals(o)
}

func (o *ScAgent) String() string {
	return base58Encode(o.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScChainId struct {
	id [33]byte
}

func NewScChainId(bytes []byte) *ScChainId {
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

func NewScColor(bytes []byte) *ScColor {
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

func NewScContractId(bytes []byte) *ScContractId {
	o := &ScContractId{}
	if len(bytes) != len(o.id) {
		logPanic("invalid contract id length")
	}
	copy(o.id[:], bytes)
	return o
}

func (o *ScContractId) AsAgent() *ScAgent {
	a := &ScAgent{}
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

func (o *ScContractId) Hname() Hname {
	return NewHnameFromBytes(o.id[33:])
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

func NewScHash(bytes []byte) *ScHash {
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

func logPanic(text string) {
	ScBaseContext{}.Panic(text)
}
