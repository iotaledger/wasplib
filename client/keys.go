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
	KeyAddress         = Key32(-1)
	KeyBalances        = Key32(-2)
	KeyBase58Bytes     = Key32(-3)
	KeyBase58String    = Key32(-4)
	KeyCall            = Key32(-5)
	KeyCaller          = Key32(-6)
	KeyChainOwnerId    = Key32(-7)
	KeyColor           = Key32(-8)
	KeyContractCreator = Key32(-9)
	KeyContractId      = Key32(-10)
	KeyData            = Key32(-11)
	KeyDeploy          = Key32(-12)
	KeyEvent           = Key32(-13)
	KeyExports         = Key32(-14)
	KeyHashBlake2b     = Key32(-15)
	KeyHashSha3        = Key32(-16)
	KeyHname           = Key32(-17)
	KeyIncoming        = Key32(-18)
	KeyLength          = Key32(-19)
	KeyLog             = Key32(-20)
	KeyLogs            = Key32(-21)
	KeyMaps            = Key32(-22)
	KeyName            = Key32(-23)
	KeyPanic           = Key32(-24)
	KeyParams          = Key32(-25)
	KeyPost            = Key32(-26)
	KeyRandom          = Key32(-27)
	KeyResults         = Key32(-28)
	KeyReturn          = Key32(-29)
	KeyState           = Key32(-30)
	KeyTimestamp       = Key32(-31)
	KeyTrace           = Key32(-32)
	KeyTransfers       = Key32(-33)
	KeyUtility         = Key32(-34)
	KeyValid           = Key32(-35)
	KeyValidEd25519    = Key32(-36)
	KeyZzzzzzz         = Key32(-95)
)
