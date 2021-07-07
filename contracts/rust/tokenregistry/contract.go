// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type MintSupplyCall struct {
	Func   *wasmlib.ScFunc
	Params MutableMintSupplyParams
}

type TransferOwnershipCall struct {
	Func   *wasmlib.ScFunc
	Params MutableTransferOwnershipParams
}

type UpdateMetadataCall struct {
	Func   *wasmlib.ScFunc
	Params MutableUpdateMetadataParams
}

type GetInfoCall struct {
	Func   *wasmlib.ScView
	Params MutableGetInfoParams
}

type tokenregistryFuncs struct{}

var ScFuncs tokenregistryFuncs

func (sc tokenregistryFuncs) MintSupply(ctx wasmlib.ScFuncCallContext) *MintSupplyCall {
	f := &MintSupplyCall{Func: wasmlib.NewScFunc(HScName, HFuncMintSupply)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc tokenregistryFuncs) TransferOwnership(ctx wasmlib.ScFuncCallContext) *TransferOwnershipCall {
	f := &TransferOwnershipCall{Func: wasmlib.NewScFunc(HScName, HFuncTransferOwnership)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc tokenregistryFuncs) UpdateMetadata(ctx wasmlib.ScFuncCallContext) *UpdateMetadataCall {
	f := &UpdateMetadataCall{Func: wasmlib.NewScFunc(HScName, HFuncUpdateMetadata)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc tokenregistryFuncs) GetInfo(ctx wasmlib.ScViewCallContext) *GetInfoCall {
	f := &GetInfoCall{Func: wasmlib.NewScView(HScName, HViewGetInfo)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}
