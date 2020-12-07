// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

var keyLength int32
var keyLog int32
var keyTrace int32

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type KeyId interface {
	GetId() int32
}

type Key string

func (key Key) GetId() int32 {
	return GetKeyId(string(key))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func KeyLength() int32 {
	if keyLength == 0 {
		keyLength = GetKeyId("length")
	}
	return keyLength
}

func KeyLog() int32 {
	if keyLog == 0 {
		keyLog = GetKeyId("log")
	}
	return keyLog
}

func KeyTrace() int32 {
	if keyTrace == 0 {
		keyTrace = GetKeyId("trace")
	}
	return keyTrace
}
