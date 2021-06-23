// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableFinalizeAuctionParams struct {
	id int32
}

func (s ImmutableFinalizeAuctionParams) Color() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxParamColor])
}

type MutableFinalizeAuctionParams struct {
	id int32
}

func (s MutableFinalizeAuctionParams) Color() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxParamColor])
}

type ImmutablePlaceBidParams struct {
	id int32
}

func (s ImmutablePlaceBidParams) Color() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxParamColor])
}

type MutablePlaceBidParams struct {
	id int32
}

func (s MutablePlaceBidParams) Color() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxParamColor])
}

type ImmutableSetOwnerMarginParams struct {
	id int32
}

func (s ImmutableSetOwnerMarginParams) OwnerMargin() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamOwnerMargin])
}

type MutableSetOwnerMarginParams struct {
	id int32
}

func (s MutableSetOwnerMarginParams) OwnerMargin() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamOwnerMargin])
}

type ImmutableStartAuctionParams struct {
	id int32
}

func (s ImmutableStartAuctionParams) Color() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxParamColor])
}

func (s ImmutableStartAuctionParams) Description() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamDescription])
}

func (s ImmutableStartAuctionParams) Duration() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxParamDuration])
}

func (s ImmutableStartAuctionParams) MinimumBid() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamMinimumBid])
}

type MutableStartAuctionParams struct {
	id int32
}

func (s MutableStartAuctionParams) Color() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxParamColor])
}

func (s MutableStartAuctionParams) Description() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamDescription])
}

func (s MutableStartAuctionParams) Duration() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxParamDuration])
}

func (s MutableStartAuctionParams) MinimumBid() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamMinimumBid])
}

type ImmutableGetInfoParams struct {
	id int32
}

func (s ImmutableGetInfoParams) Color() wasmlib.ScImmutableColor {
	return wasmlib.NewScImmutableColor(s.id, idxMap[IdxParamColor])
}

type MutableGetInfoParams struct {
	id int32
}

func (s MutableGetInfoParams) Color() wasmlib.ScMutableColor {
	return wasmlib.NewScMutableColor(s.id, idxMap[IdxParamColor])
}
