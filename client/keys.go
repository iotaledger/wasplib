// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

// @formatter:off
const (
	KeyAgent       = Key32(-1)
	KeyBalances    = KeyAgent       -1
	KeyBase58      = KeyBalances    -1
	KeyCaller      = KeyBase58      -1
	KeyCalls       = KeyCaller      -1
	KeyChain       = KeyCalls       -1
	KeyChainOwner  = KeyChain       -1
	KeyColor       = KeyChainOwner  -1
	KeyContract    = KeyColor       -1
	KeyCreator     = KeyContract    -1
	KeyData        = KeyCreator     -1
	KeyDelay       = KeyData        -1
	KeyDeploys     = KeyDelay       -1
	KeyDescription = KeyDeploys      -1
	KeyEvent       = KeyDescription -1
	KeyExports     = KeyEvent       -1
	KeyFunction    = KeyExports     -1
	KeyHash        = KeyFunction    -1
	KeyId          = KeyHash        -1
	KeyIncoming    = KeyId          -1
	KeyLength      = KeyIncoming    -1
	KeyLog         = KeyLength      -1
	KeyLogs        = KeyLog         -1
	KeyName        = KeyLogs        -1
	KeyPanic       = KeyName        -1
	KeyParams      = KeyPanic       -1
	KeyPosts       = KeyParams      -1
	KeyRandom      = KeyPosts       -1
	KeyResults     = KeyRandom      -1
	KeyState       = KeyResults     -1
	KeyTimestamp   = KeyState       -1
	KeyTrace       = KeyTimestamp   -1
	KeyTransfers   = KeyTrace       -1
	KeyUtility     = KeyTransfers   -1
	KeyViews       = KeyUtility     -1
	KeyZzzzzzz     = KeyViews       -1
)
// @formatter:on

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type MapKey interface {
	KeyId() int32
}

type Key string

func (key Key) KeyId() int32 {
	return GetKeyIdFromString(string(key))
}

type Key32 int32

func (key Key32) KeyId() int32 {
	return int32(key)
}
