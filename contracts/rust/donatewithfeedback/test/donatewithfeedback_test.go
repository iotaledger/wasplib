// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/donatewithfeedback"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *common.SoloContext {
	chain := common.StartChainAndDeployWasmContractByName(t, donatewithfeedback.ScName)
	return common.NewSoloContext(donatewithfeedback.ScName, donatewithfeedback.OnLoad, chain)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	_, err := ctx.Chain.FindContract(donatewithfeedback.ScName)
	require.NoError(t, err)
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

	donator1, donator1Addr := ctx.Chain.Env.NewKeyPairWithFunds()
	donate := donatewithfeedback.ScFuncs.Donate(ctx.SignWith(donator1))
	donate.Params.Feedback().SetValue("Nice work!")
	donate.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	donationInfo := donatewithfeedback.ScFuncs.DonationInfo(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 1, donationInfo.Results.Count().Value())
	require.EqualValues(t, 42, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42, donationInfo.Results.TotalDonation().Value())

	// 42 iota transferred from wallet to contract
	ctx.Chain.Env.AssertAddressBalance(donator1Addr, colored.IOTA, solo.Saldo-42)
	// 42 iota transferred to contract
	ctx.Chain.AssertAccountBalance(ctx.Chain.ContractAgentID(donatewithfeedback.ScName), colored.IOTA, 42)
	// returned 1 used for transaction to wallet account
	account1 := iscp.NewAgentID(donator1Addr, 0)
	ctx.Chain.AssertAccountBalance(account1, colored.IOTA, 0)
}

func TestDonateTwice(t *testing.T) {
	ctx := setupTest(t)

	donator1, donator1Addr := ctx.Chain.Env.NewKeyPairWithFunds()
	donate1 := donatewithfeedback.ScFuncs.Donate(ctx.SignWith(donator1))
	donate1.Params.Feedback().SetValue("Nice work!")
	donate1.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	donator2, donator2Addr := ctx.Chain.Env.NewKeyPairWithFunds()
	donate2 := donatewithfeedback.ScFuncs.Donate(ctx.SignWith(donator2))
	donate2.Params.Feedback().SetValue("Exactly what I needed!")
	donate2.Func.TransferIotas(69).Post()
	require.NoError(t, ctx.Err)

	donationInfo := donatewithfeedback.ScFuncs.DonationInfo(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 2, donationInfo.Results.Count().Value())
	require.EqualValues(t, 69, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42+69, donationInfo.Results.TotalDonation().Value())

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	ctx.Chain.Env.AssertAddressBalance(donator1Addr, colored.IOTA, solo.Saldo-42)
	// 69 iota transferred from wallet to contract plus 1 used for transaction
	ctx.Chain.Env.AssertAddressBalance(donator2Addr, colored.IOTA, solo.Saldo-69)
	// 42+69 iota transferred to contract
	ctx.Chain.AssertAccountBalance(ctx.Chain.ContractAgentID(donatewithfeedback.ScName), colored.IOTA, 42+69)
	// returned 1 used for transaction to wallet accounts
	account1 := iscp.NewAgentID(donator1Addr, 0)
	ctx.Chain.AssertAccountBalance(account1, colored.IOTA, 0)
	account2 := iscp.NewAgentID(donator2Addr, 0)
	ctx.Chain.AssertAccountBalance(account2, colored.IOTA, 0)
}
