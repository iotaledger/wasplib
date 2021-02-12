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
	KeyDeploy          = Key32(-12)
	KeyEvent           = Key32(-13)
	KeyExports         = Key32(-14)
	KeyHashBlake2b     = Key32(-15)
	KeyHashSha3        = Key32(-16)
	KeyHname           = Key32(-17)
	KeyIncoming        = Key32(-18)
	KeyLength          = Key32(-19)
	KeyLog             = Key32(-20)
	KeyMaps            = Key32(-21)
	KeyName            = Key32(-22)
	KeyPanic           = Key32(-23)
	KeyParams          = Key32(-24)
	KeyPost            = Key32(-25)
	KeyRandom          = Key32(-26)
	KeyResults         = Key32(-27)
	KeyReturn          = Key32(-28)
	KeyState           = Key32(-29)
	KeyTimestamp       = Key32(-30)
	KeyTrace           = Key32(-31)
	KeyTransfers       = Key32(-32)
	KeyUtility         = Key32(-33)
	KeyValid           = Key32(-34)
	KeyValidBls        = Key32(-35)
	KeyValidEd25519    = Key32(-36)
	KeyZzzzzzz         = Key32(-37)
)
