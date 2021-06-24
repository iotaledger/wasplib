// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type DonateCall struct {
	Func   wasmlib.ScFunc
	Params MutableDonateParams
}

func NewDonateCall(ctx wasmlib.ScFuncContext) *DonateCall {
	f := &DonateCall{}
	f.Func.Init(HScName, HFuncDonate, &f.Params.id, nil)
	return f
}

type WithdrawCall struct {
	Func   wasmlib.ScFunc
	Params MutableWithdrawParams
}

func NewWithdrawCall(ctx wasmlib.ScFuncContext) *WithdrawCall {
	f := &WithdrawCall{}
	f.Func.Init(HScName, HFuncWithdraw, &f.Params.id, nil)
	return f
}

type DonationCall struct {
	Func    wasmlib.ScView
	Params  MutableDonationParams
	Results ImmutableDonationResults
}

func NewDonationCall(ctx wasmlib.ScFuncContext) *DonationCall {
	f := &DonationCall{}
	f.Func.Init(HScName, HViewDonation, &f.Params.id, &f.Results.id)
	return f
}

func NewDonationCallFromView(ctx wasmlib.ScViewContext) *DonationCall {
	f := &DonationCall{}
	f.Func.Init(HScName, HViewDonation, &f.Params.id, &f.Results.id)
	return f
}

type DonationInfoCall struct {
	Func    wasmlib.ScView
	Results ImmutableDonationInfoResults
}

func NewDonationInfoCall(ctx wasmlib.ScFuncContext) *DonationInfoCall {
	f := &DonationInfoCall{}
	f.Func.Init(HScName, HViewDonationInfo, nil, &f.Results.id)
	return f
}

func NewDonationInfoCallFromView(ctx wasmlib.ScViewContext) *DonationInfoCall {
	f := &DonationInfoCall{}
	f.Func.Init(HScName, HViewDonationInfo, nil, &f.Results.id)
	return f
}
