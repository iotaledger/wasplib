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

// @formatter:off
const (
	KeyAgent       = Key32(-1)
	KeyBalances    = Key32(-2)
	KeyBase58      = Key32(-3)
	KeyCaller      = Key32(-4)
	KeyCalls       = Key32(-5)
	KeyChain       = Key32(-6)
	KeyChainOwner  = Key32(-7)
	KeyColor       = Key32(-8)
	KeyContract    = Key32(-9)
	KeyCreator     = Key32(-10)
	KeyData        = Key32(-11)
	KeyDelay       = Key32(-12)
	KeyDeploys     = Key32(-13)
	KeyDescription = Key32(-14)
	KeyEvent       = Key32(-15)
	KeyExports     = Key32(-16)
	KeyFunction    = Key32(-17)
	KeyHash        = Key32(-18)
	KeyId          = Key32(-19)
	KeyIncoming    = Key32(-20)
	KeyLength      = Key32(-21)
	KeyLog         = Key32(-22)
	KeyLogs        = Key32(-23)
	KeyMaps        = Key32(-24)
	KeyName        = Key32(-25)
	KeyPanic       = Key32(-26)
	KeyParams      = Key32(-27)
	KeyRandom      = Key32(-28)
	KeyResults     = Key32(-29)
	KeyState       = Key32(-30)
	KeyTimestamp   = Key32(-31)
	KeyTrace       = Key32(-32)
	KeyTransfers   = Key32(-33)
	KeyUtility     = Key32(-34)
	KeyZzzzzzz     = Key32(-99)
)
// @formatter:on
