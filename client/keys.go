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
	KeyAggregateBls    = Key32(-2)
	KeyBalances        = Key32(-3)
	KeyBase58Bytes     = Key32(-4)
	KeyBase58String    = Key32(-5)
	KeyCall            = Key32(-6)
	KeyCaller          = Key32(-7)
	KeyChainOwnerId    = Key32(-8)
	KeyColor           = Key32(-9)
	KeyContractCreator = Key32(-10)
	KeyContractId      = Key32(-11)
	KeyData            = Key32(-12)
	KeyDeploy          = Key32(-13)
	KeyEvent           = Key32(-14)
	KeyExports         = Key32(-15)
	KeyHashBlake2b     = Key32(-16)
	KeyHashSha3        = Key32(-17)
	KeyHname           = Key32(-18)
	KeyIncoming        = Key32(-19)
	KeyLength          = Key32(-20)
	KeyLog             = Key32(-21)
	KeyLogs            = Key32(-22)
	KeyMaps            = Key32(-23)
	KeyName            = Key32(-24)
	KeyPanic           = Key32(-25)
	KeyParams          = Key32(-26)
	KeyPost            = Key32(-27)
	KeyRandom          = Key32(-28)
	KeyResults         = Key32(-29)
	KeyReturn          = Key32(-30)
	KeyState           = Key32(-31)
	KeyTimestamp       = Key32(-32)
	KeyTrace           = Key32(-33)
	KeyTransfers       = Key32(-34)
	KeyUtility         = Key32(-35)
	KeyValid           = Key32(-36)
	KeyValidBls        = Key32(-37)
	KeyValidEd25519    = Key32(-38)
	KeyZzzzzzz         = Key32(-96)
)
