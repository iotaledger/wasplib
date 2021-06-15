// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreroot

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type CoreRootFunc struct {
	sc wasmlib.ScContractFunc
}

func NewCoreRootFunc(ctx wasmlib.ScFuncContext) *CoreRootFunc {
	return &CoreRootFunc{sc: wasmlib.NewScContractFunc(ctx, HScName)}
}

func (f *CoreRootFunc) Delay(seconds int32) *CoreRootFunc {
	f.sc.Delay(seconds)
	return f
}

func (f *CoreRootFunc) OfContract(contract wasmlib.ScHname) *CoreRootFunc {
	f.sc.OfContract(contract)
	return f
}

func (f *CoreRootFunc) Post() *CoreRootFunc {
	f.sc.Post()
	return f
}

func (f *CoreRootFunc) PostToChain(chainId wasmlib.ScChainId) *CoreRootFunc {
	f.sc.PostToChain(chainId)
	return f
}

func (f *CoreRootFunc) ClaimChainOwnership(transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncClaimChainOwnership, 0, &transfer)
}

func (f *CoreRootFunc) DelegateChainOwnership(params MutableFuncDelegateChainOwnershipParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncDelegateChainOwnership, params.id, &transfer)
}

func (f *CoreRootFunc) DeployContract(params MutableFuncDeployContractParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncDeployContract, params.id, &transfer)
}

func (f *CoreRootFunc) GrantDeployPermission(params MutableFuncGrantDeployPermissionParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncGrantDeployPermission, params.id, &transfer)
}

func (f *CoreRootFunc) RevokeDeployPermission(params MutableFuncRevokeDeployPermissionParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncRevokeDeployPermission, params.id, &transfer)
}

func (f *CoreRootFunc) SetContractFee(params MutableFuncSetContractFeeParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncSetContractFee, params.id, &transfer)
}

func (f *CoreRootFunc) SetDefaultFee(params MutableFuncSetDefaultFeeParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncSetDefaultFee, params.id, &transfer)
}

func (f *CoreRootFunc) FindContract(params MutableViewFindContractParams) ImmutableViewFindContractResults {
	f.sc.Run(HViewFindContract, params.id, nil)
	return ImmutableViewFindContractResults{id: f.sc.ResultMapId()}
}

func (f *CoreRootFunc) GetChainInfo() ImmutableViewGetChainInfoResults {
	f.sc.Run(HViewGetChainInfo, 0, nil)
	return ImmutableViewGetChainInfoResults{id: f.sc.ResultMapId()}
}

func (f *CoreRootFunc) GetFeeInfo(params MutableViewGetFeeInfoParams) ImmutableViewGetFeeInfoResults {
	f.sc.Run(HViewGetFeeInfo, params.id, nil)
	return ImmutableViewGetFeeInfoResults{id: f.sc.ResultMapId()}
}

type CoreRootView struct {
	sc wasmlib.ScContractView
}

func NewCoreRootView(ctx wasmlib.ScViewContext) *CoreRootView {
	return &CoreRootView{sc: wasmlib.NewScContractView(ctx, HScName)}
}

func (v *CoreRootView) OfContract(contract wasmlib.ScHname) *CoreRootView {
	v.sc.OfContract(contract)
	return v
}

func (v *CoreRootView) FindContract(params MutableViewFindContractParams) ImmutableViewFindContractResults {
	v.sc.Run(HViewFindContract, params.id)
	return ImmutableViewFindContractResults{id: v.sc.ResultMapId()}
}

func (v *CoreRootView) GetChainInfo() ImmutableViewGetChainInfoResults {
	v.sc.Run(HViewGetChainInfo, 0)
	return ImmutableViewGetChainInfoResults{id: v.sc.ResultMapId()}
}

func (v *CoreRootView) GetFeeInfo(params MutableViewGetFeeInfoParams) ImmutableViewGetFeeInfoResults {
	v.sc.Run(HViewGetFeeInfo, params.id)
	return ImmutableViewGetFeeInfoResults{id: v.sc.ResultMapId()}
}
