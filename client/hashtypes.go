// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

type ScAddress struct {
	address [33]byte
}

func NewScAddress(bytes []byte) *ScAddress {
	if len(bytes) != 33 {
		panic("address should be 33 bytes")
	}
	a := &ScAddress{}
	copy(a.address[:], bytes)
	return a
}

func (a *ScAddress) Bytes() []byte {
	return a.address[:]
}

func (a *ScAddress) Equals(other *ScAddress) bool {
	return a.address == other.address
}

func (a *ScAddress) String() string {
	return base58Encode(a.address[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAgent struct {
	id [37]byte
}

func NewScAgent(bytes []byte) *ScAgent {
	if len(bytes) != 37 {
		panic("agent id should be 37 bytes")
	}
	a := &ScAgent{}
	copy(a.id[:], bytes)
	return a
}

func (a *ScAgent) Bytes() []byte {
	return a.id[:]
}

func (a *ScAgent) Equals(other *ScAgent) bool {
	return a.id == other.id
}

func (a *ScAgent) String() string {
	return base58Encode(a.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	color [32]byte
}

var IOTA = &ScColor{}
var MINT = &ScColor{}

func init() {
	for i := range MINT.color {
		MINT.color[i] = 0xff
	}
}

func NewScColor(bytes []byte) *ScColor {
	if len(bytes) != 32 {
		panic("color should be 32 bytes")
	}
	a := &ScColor{}
	copy(a.color[:], bytes)
	return a
}

func (c *ScColor) Bytes() []byte {
	return c.color[:]
}

func (c *ScColor) Equals(other *ScColor) bool {
	return c.color == other.color
}

func (c *ScColor) String() string {
	return base58Encode(c.color[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestId struct {
	id [34]byte
}

func NewScRequestId(bytes []byte) *ScRequestId {
	if len(bytes) != 34 {
		panic("request id should be 34 bytes")
	}
	t := &ScRequestId{}
	copy(t.id[:], bytes)
	return t
}

func (t *ScRequestId) Bytes() []byte {
	return t.id[:]
}

func (t *ScRequestId) Equals(other *ScRequestId) bool {
	return t.id == other.id
}

func (t *ScRequestId) String() string {
	return base58Encode(t.id[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTxHash struct {
	hash [32]byte
}

func NewScTxHash(bytes []byte) *ScTxHash {
	if len(bytes) != 32 {
		panic("tx hash should be 32 bytes")
	}
	t := &ScTxHash{}
	copy(t.hash[:], bytes)
	return t
}

func (t *ScTxHash) Bytes() []byte {
	return t.hash[:]
}

func (t *ScTxHash) Equals(other *ScTxHash) bool {
	return t.hash == other.hash
}

func (t *ScTxHash) String() string {
	return base58Encode(t.hash[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func base58Encode(bytes []byte) string {
	return NewScContext().Utility().Base58Encode(bytes)
}
