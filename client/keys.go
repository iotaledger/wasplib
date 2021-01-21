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
	KeyAddress     = Key32(-1)
	KeyAgent       = Key32(-2)
	KeyBalances    = Key32(-3)
	KeyBase58      = Key32(-4)
	KeyCaller      = Key32(-5)
	KeyCalls       = Key32(-6)
	KeyChain       = Key32(-7)
	KeyChainOwner  = Key32(-8)
	KeyColor       = Key32(-9)
	KeyContract    = Key32(-10)
	KeyCreator     = Key32(-11)
	KeyData        = Key32(-12)
	KeyDelay       = Key32(-13)
	KeyDeploys     = Key32(-14)
	KeyDescription = Key32(-15)
	KeyEvent       = Key32(-16)
	KeyExports     = Key32(-17)
	KeyFunction    = Key32(-18)
	KeyHash        = Key32(-19)
	KeyId          = Key32(-20)
	KeyIncoming    = Key32(-21)
	KeyLength      = Key32(-22)
	KeyLog         = Key32(-23)
	KeyLogs        = Key32(-24)
	KeyMaps        = Key32(-25)
	KeyName        = Key32(-26)
	KeyPanic       = Key32(-27)
	KeyParams      = Key32(-28)
	KeyRandom      = Key32(-29)
	KeyResults     = Key32(-30)
	KeyState       = Key32(-31)
	KeyTimestamp   = Key32(-32)
	KeyTrace       = Key32(-33)
	KeyTransfers   = Key32(-34)
	KeyUtility     = Key32(-35)
	KeyZzzzzzz     = Key32(-98)
)
// @formatter:on
