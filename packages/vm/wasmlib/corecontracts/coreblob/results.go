// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableFuncStoreBlobResults struct {
	id int32
}

func (s ImmutableFuncStoreBlobResults) Hash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, ResultHash.KeyId())
}

type MutableFuncStoreBlobResults struct {
	id int32
}

func (s MutableFuncStoreBlobResults) Hash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, ResultHash.KeyId())
}

type ImmutableViewGetBlobFieldResults struct {
	id int32
}

func (s ImmutableViewGetBlobFieldResults) Bytes() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, ResultBytes.KeyId())
}

type MutableViewGetBlobFieldResults struct {
	id int32
}

func (s MutableViewGetBlobFieldResults) Bytes() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, ResultBytes.KeyId())
}