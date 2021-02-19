// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

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
	KeyBlsAddress      = Key32(-5)
	KeyBlsAggregate    = Key32(-6)
	KeyBlsValid        = Key32(-7)
	KeyCall            = Key32(-8)
	KeyCaller          = Key32(-9)
	KeyChainOwnerId    = Key32(-10)
	KeyColor           = Key32(-11)
	KeyContractCreator = Key32(-12)
	KeyContractId      = Key32(-13)
	KeyDeploy          = Key32(-14)
	KeyEd25519Address  = Key32(-15)
	KeyEd25519Valid    = Key32(-16)
	KeyEvent           = Key32(-17)
	KeyExports         = Key32(-18)
	KeyHashBlake2b     = Key32(-19)
	KeyHashSha3        = Key32(-20)
	KeyHname           = Key32(-21)
	KeyIncoming        = Key32(-22)
	KeyLength          = Key32(-23)
	KeyLog             = Key32(-24)
	KeyMaps            = Key32(-25)
	KeyName            = Key32(-26)
	KeyPanic           = Key32(-27)
	KeyParams          = Key32(-28)
	KeyPost            = Key32(-29)
	KeyRandom          = Key32(-30)
	KeyResults         = Key32(-31)
	KeyReturn          = Key32(-32)
	KeyState           = Key32(-33)
	KeyTimestamp       = Key32(-34)
	KeyTrace           = Key32(-35)
	KeyTransfers       = Key32(-36)
	KeyUtility         = Key32(-37)
	KeyValid           = Key32(-38)
	KeyZzzzzzz         = Key32(-39)
)
