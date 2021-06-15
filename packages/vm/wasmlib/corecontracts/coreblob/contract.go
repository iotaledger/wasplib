// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type CoreBlobFunc struct {
	sc wasmlib.ScContractFunc
}

func NewCoreBlobFunc(ctx wasmlib.ScFuncContext) *CoreBlobFunc {
	return &CoreBlobFunc{sc: wasmlib.NewScContractFunc(ctx, HScName)}
}

func (f *CoreBlobFunc) Delay(seconds int32) *CoreBlobFunc {
	f.sc.Delay(seconds)
	return f
}

func (f *CoreBlobFunc) OfContract(contract wasmlib.ScHname) *CoreBlobFunc {
	f.sc.OfContract(contract)
	return f
}

func (f *CoreBlobFunc) Post() *CoreBlobFunc {
	f.sc.Post()
	return f
}

func (f *CoreBlobFunc) PostToChain(chainId wasmlib.ScChainId) *CoreBlobFunc {
	f.sc.PostToChain(chainId)
	return f
}

func (f *CoreBlobFunc) StoreBlob(transfer wasmlib.ScTransfers) ImmutableFuncStoreBlobResults {
	f.sc.Run(HFuncStoreBlob, 0, &transfer)
	return ImmutableFuncStoreBlobResults{id: f.sc.ResultMapId()}
}

func (f *CoreBlobFunc) GetBlobField(params MutableViewGetBlobFieldParams) ImmutableViewGetBlobFieldResults {
	f.sc.Run(HViewGetBlobField, params.id, nil)
	return ImmutableViewGetBlobFieldResults{id: f.sc.ResultMapId()}
}

func (f *CoreBlobFunc) GetBlobInfo(params MutableViewGetBlobInfoParams) {
	f.sc.Run(HViewGetBlobInfo, params.id, nil)
}

func (f *CoreBlobFunc) ListBlobs() {
	f.sc.Run(HViewListBlobs, 0, nil)
}

type CoreBlobView struct {
	sc wasmlib.ScContractView
}

func NewCoreBlobView(ctx wasmlib.ScViewContext) *CoreBlobView {
	return &CoreBlobView{sc: wasmlib.NewScContractView(ctx, HScName)}
}

func (v *CoreBlobView) OfContract(contract wasmlib.ScHname) *CoreBlobView {
	v.sc.OfContract(contract)
	return v
}

func (v *CoreBlobView) GetBlobField(params MutableViewGetBlobFieldParams) ImmutableViewGetBlobFieldResults {
	v.sc.Run(HViewGetBlobField, params.id)
	return ImmutableViewGetBlobFieldResults{id: v.sc.ResultMapId()}
}

func (v *CoreBlobView) GetBlobInfo(params MutableViewGetBlobInfoParams) {
	v.sc.Run(HViewGetBlobInfo, params.id)
}

func (v *CoreBlobView) ListBlobs() {
	v.sc.Run(HViewListBlobs, 0)
}
