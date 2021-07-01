// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type DonateCall struct {
	Func   *wasmlib.ScFunc
	Params MutableDonateParams
}

func NewDonateCall(ctx wasmlib.ScHostContext) *DonateCall {
	f := &DonateCall{Func: wasmlib.NewScFunc(HScName, HFuncDonate)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type WithdrawCall struct {
	Func   *wasmlib.ScFunc
	Params MutableWithdrawParams
}

func NewWithdrawCall(ctx wasmlib.ScHostContext) *WithdrawCall {
	f := &WithdrawCall{Func: wasmlib.NewScFunc(HScName, HFuncWithdraw)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type DonationCall struct {
	Func    *wasmlib.ScView
	Params  MutableDonationParams
	Results ImmutableDonationResults
}

func NewDonationCall(ctx wasmlib.ScHostContext) *DonationCall {
	f := &DonationCall{Func: wasmlib.NewScView(HScName, HViewDonation)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewDonationCallFromView(ctx wasmlib.ScViewContext) *DonationCall {
	return NewDonationCall(wasmlib.ScFuncContext{})
}

type DonationInfoCall struct {
	Func    *wasmlib.ScView
	Results ImmutableDonationInfoResults
}

func NewDonationInfoCall(ctx wasmlib.ScHostContext) *DonationInfoCall {
	f := &DonationInfoCall{Func: wasmlib.NewScView(HScName, HViewDonationInfo)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func NewDonationInfoCallFromView(ctx wasmlib.ScViewContext) *DonationInfoCall {
	return NewDonationInfoCall(wasmlib.ScFuncContext{})
}
