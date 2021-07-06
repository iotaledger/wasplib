// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/donatewithfeedback"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, donatewithfeedback.ScName)
}

func TestDeploy(t *testing.T) {
	chain := setupTest(t)
	_, err := chain.FindContract(donatewithfeedback.ScName)
	require.NoError(t, err)
}

func TestStateAfterDeploy(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(donatewithfeedback.ScName, donatewithfeedback.OnLoad, chain, nil)

	donationInfo := donatewithfeedback.NewDonationInfoCall(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 0, donationInfo.Results.Count().Value())
	require.EqualValues(t, 0, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 0, donationInfo.Results.TotalDonation().Value())
}

func TestDonateOnce(t *testing.T) {
	chain := setupTest(t)

	donator1, donator1Addr := chain.Env.NewKeyPairWithFunds()
	ctx := common.NewSoloContext(donatewithfeedback.ScName, donatewithfeedback.OnLoad, chain, donator1)
	donate := donatewithfeedback.NewDonateCall(ctx)
	donate.Params.Feedback().SetValue("Nice work!")
	donate.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	donationInfo := donatewithfeedback.NewDonationInfoCall(ctx)
	donationInfo.Func.Call()

	require.EqualValues(t, 1, donationInfo.Results.Count().Value())
	require.EqualValues(t, 42, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42, donationInfo.Results.TotalDonation().Value())

	// 42 iota transferred from wallet to contract
	chain.Env.AssertAddressBalance(donator1Addr, ledgerstate.ColorIOTA, solo.Saldo-42)
	// 42 iota transferred to contract
	chain.AssertAccountBalance(chain.ContractAgentID(donatewithfeedback.ScName), ledgerstate.ColorIOTA, 42)
	// returned 1 used for transaction to wallet account
	account1 := coretypes.NewAgentID(donator1Addr, 0)
	chain.AssertAccountBalance(account1, ledgerstate.ColorIOTA, 0)
}

func TestDonateTwice(t *testing.T) {
	chain := setupTest(t)

	donator1, donator1Addr := chain.Env.NewKeyPairWithFunds()
	ctx1 := common.NewSoloContext(donatewithfeedback.ScName, donatewithfeedback.OnLoad, chain, donator1)
	donate1 := donatewithfeedback.NewDonateCall(ctx1)
	donate1.Params.Feedback().SetValue("Nice work!")
	donate1.Func.TransferIotas(42).Post()
	require.NoError(t, ctx1.Err)

	donator2, donator2Addr := chain.Env.NewKeyPairWithFunds()
	ctx2 := common.NewSoloContext(donatewithfeedback.ScName, donatewithfeedback.OnLoad, chain, donator2)
	donate2 := donatewithfeedback.NewDonateCall(ctx2)
	donate2.Params.Feedback().SetValue("Exactly what I needed!")
	donate2.Func.TransferIotas(69).Post()
	require.NoError(t, ctx2.Err)

	donationInfo := donatewithfeedback.NewDonationInfoCall(ctx1)
	donationInfo.Func.Call()

	require.EqualValues(t, 2, donationInfo.Results.Count().Value())
	require.EqualValues(t, 69, donationInfo.Results.MaxDonation().Value())
	require.EqualValues(t, 42+69, donationInfo.Results.TotalDonation().Value())

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator1Addr, ledgerstate.ColorIOTA, solo.Saldo-42)
	// 69 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator2Addr, ledgerstate.ColorIOTA, solo.Saldo-69)
	// 42+69 iota transferred to contract
	chain.AssertAccountBalance(chain.ContractAgentID(donatewithfeedback.ScName), ledgerstate.ColorIOTA, 42+69)
	// returned 1 used for transaction to wallet accounts
	account1 := coretypes.NewAgentID(donator1Addr, 0)
	chain.AssertAccountBalance(account1, ledgerstate.ColorIOTA, 0)
	account2 := coretypes.NewAgentID(donator2Addr, 0)
	chain.AssertAccountBalance(account2, ledgerstate.ColorIOTA, 0)
}
