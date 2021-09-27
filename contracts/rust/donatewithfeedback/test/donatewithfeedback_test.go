// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/donatewithfeedback"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, donatewithfeedback.ScName, donatewithfeedback.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	require.NoError(t, ctx.ContractExists(donatewithfeedback.ScName))
}

func TestStateAfterDeploy(t *testing.T) {
	ctx := setupTest(t)

	donationInfo := donatewithfeedback.ScFuncs.DonationInfo(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 0, donationInfo.Results.Count().Value())
	require.EqualValues(t, 0, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 0, donationInfo.Results.TotalDonation().Value())
}

func TestDonateOnce(t *testing.T) {
	ctx := setupTest(t)

	donator1 := ctx.NewSoloAgent()
	donate := donatewithfeedback.ScFuncs.Donate(ctx.Sign(donator1))
	donate.Params.Feedback().SetValue("Nice work!")
	donate.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	donationInfo := donatewithfeedback.ScFuncs.DonationInfo(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 1, donationInfo.Results.Count().Value())
	require.EqualValues(t, 42, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42, donationInfo.Results.TotalDonation().Value())

	// 42 iota transferred from wallet to contract
	require.EqualValues(t, solo.Saldo-42, donator1.Balance())
	require.EqualValues(t, 42, ctx.Balance(nil))
}

func TestDonateTwice(t *testing.T) {
	ctx := setupTest(t)

	donator1 := ctx.NewSoloAgent()
	donate1 := donatewithfeedback.ScFuncs.Donate(ctx.Sign(donator1))
	donate1.Params.Feedback().SetValue("Nice work!")
	donate1.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	donator2 := ctx.NewSoloAgent()
	donate2 := donatewithfeedback.ScFuncs.Donate(ctx.Sign(donator2))
	donate2.Params.Feedback().SetValue("Exactly what I needed!")
	donate2.Func.TransferIotas(69).Post()
	require.NoError(t, ctx.Err)

	donationInfo := donatewithfeedback.ScFuncs.DonationInfo(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 2, donationInfo.Results.Count().Value())
	require.EqualValues(t, 69, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42+69, donationInfo.Results.TotalDonation().Value())

	require.EqualValues(t, solo.Saldo-42, donator1.Balance())
	require.EqualValues(t, solo.Saldo-69, donator2.Balance())
	require.EqualValues(t, 42+69, ctx.Balance(nil))
}
