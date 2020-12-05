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

func (a *ScAddress) AsAgent() *ScAgent {
	agent := &ScAgent{}
	// agent is address padded with zeroes
	copy(agent.id[:], a.address[:])
	return agent
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

func base58Encode(bytes []byte) string {
	return rootCallContext.Utility().Base58Encode(bytes)
}
