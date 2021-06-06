// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncMintSupply, funcMintSupplyThunk)
	exports.AddFunc(FuncTransferOwnership, funcTransferOwnershipThunk)
	exports.AddFunc(FuncUpdateMetadata, funcUpdateMetadataThunk)
	exports.AddView(ViewGetInfo, viewGetInfoThunk)

	for i, key := range keyMap {
		idxMap[i] = wasmlib.GetKeyIdFromString(key)
	}
}

type FuncMintSupplyParams struct {
	Description wasmlib.ScImmutableString // description what minted token represents
	UserDefined wasmlib.ScImmutableString // any user defined text
}

type FuncMintSupplyContext struct {
	Params FuncMintSupplyParams
	State  TokenRegistryFuncState
}

func funcMintSupplyThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcMintSupply")
	p := ctx.Params().MapId()
	f := &FuncMintSupplyContext{
		Params: FuncMintSupplyParams{
			Description: wasmlib.NewScImmutableString(p, idxMap[IdxParamDescription]),
			UserDefined: wasmlib.NewScImmutableString(p, idxMap[IdxParamUserDefined]),
		},
		State: TokenRegistryFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState.KeyId(), wasmlib.TYPE_MAP),
		},
	}
	funcMintSupply(ctx, f)
	ctx.Log("tokenregistry.funcMintSupply ok")
}

type FuncTransferOwnershipParams struct {
	Color wasmlib.ScImmutableColor // color of token to transfer ownership of
}

type FuncTransferOwnershipContext struct {
	Params FuncTransferOwnershipParams
	State  TokenRegistryFuncState
}

func funcTransferOwnershipThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcTransferOwnership")
	//TODO the one who can transfer token ownership
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncTransferOwnershipContext{
		Params: FuncTransferOwnershipParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		State: TokenRegistryFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState.KeyId(), wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	funcTransferOwnership(ctx, f)
	ctx.Log("tokenregistry.funcTransferOwnership ok")
}

type FuncUpdateMetadataParams struct {
	Color wasmlib.ScImmutableColor // color of token to update metadata for
}

type FuncUpdateMetadataContext struct {
	Params FuncUpdateMetadataParams
	State  TokenRegistryFuncState
}

func funcUpdateMetadataThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("tokenregistry.funcUpdateMetadata")
	//TODO the one who can change the token info
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncUpdateMetadataContext{
		Params: FuncUpdateMetadataParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		State: TokenRegistryFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState.KeyId(), wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	funcUpdateMetadata(ctx, f)
	ctx.Log("tokenregistry.funcUpdateMetadata ok")
}

type ViewGetInfoParams struct {
	Color wasmlib.ScImmutableColor // color of token to view registry info of
}

type ViewGetInfoContext struct {
	Params ViewGetInfoParams
	State  TokenRegistryViewState
}

func viewGetInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("tokenregistry.viewGetInfo")
	p := ctx.Params().MapId()
	f := &ViewGetInfoContext{
		Params: ViewGetInfoParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		State: TokenRegistryViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState.KeyId(), wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	viewGetInfo(ctx, f)
	ctx.Log("tokenregistry.viewGetInfo ok")
}
