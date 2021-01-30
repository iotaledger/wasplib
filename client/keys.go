// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

type MapKey interface {
	KeyId() Key32
}

type Key string

func (key Key) KeyId() Key32 {
	return GetKeyIdFromString(string(key))
}

type Key32 int32

func (key Key32) KeyId() Key32 {
	return key
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

const (
	KeyAddress      = Key32(-1)
	KeyAgent        = Key32(-2)
	KeyBalances     = Key32(-3)
	KeyBase58Bytes  = Key32(-4)
	KeyBase58String = Key32(-5)
	KeyCall         = Key32(-6)
	KeyCaller       = Key32(-7)
	KeyChain        = Key32(-8)
	KeyChainOwner   = Key32(-9)
	KeyColor        = Key32(-10)
	KeyContract     = Key32(-11)
	KeyCreator      = Key32(-12)
	KeyData         = Key32(-13)
	KeyDelay        = Key32(-14)
	KeyDeploy       = Key32(-15)
	KeyDescription  = Key32(-16)
	KeyEvent        = Key32(-17)
	KeyExports      = Key32(-18)
	KeyFunction     = Key32(-19)
	KeyHash         = Key32(-20)
	KeyHname        = Key32(-21)
	KeyId           = Key32(-22)
	KeyIncoming     = Key32(-23)
	KeyLength       = Key32(-24)
	KeyLog          = Key32(-25)
	KeyLogs         = Key32(-26)
	KeyMaps         = Key32(-27)
	KeyName         = Key32(-28)
	KeyPanic        = Key32(-29)
	KeyParams       = Key32(-30)
	KeyPost         = Key32(-31)
	KeyRandom       = Key32(-32)
	KeyResults      = Key32(-33)
	KeyReturn       = Key32(-34)
	KeyState        = Key32(-35)
	KeyTimestamp    = Key32(-36)
	KeyTrace        = Key32(-37)
	KeyTransfers    = Key32(-38)
	KeyUtility      = Key32(-39)
	KeyZzzzzzz      = Key32(-99)
)
