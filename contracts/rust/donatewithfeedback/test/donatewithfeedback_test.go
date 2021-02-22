// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestStateAfterDeploy(t *testing.T) {
	chain := setupTest(t)

	ret, err := chain.CallView(
		ScName, ViewDonations,
	)
	require.NoError(t, err)

	max, _, err := codec.DecodeInt64(ret[VarMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 0, max)

	tot, _, err := codec.DecodeInt64(ret[VarTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 0, tot)
}

func TestDonateOnce(t *testing.T) {
	chain := setupTest(t)

	donator1 := chain.Env.NewSignatureSchemeWithFunds()
	req := solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Nice work!",
	).WithTransfer(balance.ColorIOTA, 42)
	_, err := chain.PostRequestSync(req, donator1)
	require.NoError(t, err)

	ret, err := chain.CallView(
		ScName, ViewDonations,
	)
	require.NoError(t, err)

	max, _, err := codec.DecodeInt64(ret[VarMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42, max)

	tot, _, err := codec.DecodeInt64(ret[VarTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42, tot)

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator1.Address(), balance.ColorIOTA, solo.Saldo-42-1)
	// 42 iota transferred to contract
	chain.AssertAccountBalance(common.ContractAccount, balance.ColorIOTA, 42)
	// returned 1 used for transaction to wallet account
	account1 := coretypes.NewAgentIDFromSigScheme(donator1)
	chain.AssertAccountBalance(account1, balance.ColorIOTA, 1)
}

func TestDonateTwice(t *testing.T) {
	chain := setupTest(t)

	donator1 := chain.Env.NewSignatureSchemeWithFunds()
	req := solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Nice work!",
	).WithTransfer(balance.ColorIOTA, 42)
	_, err := chain.PostRequestSync(req, donator1)
	require.NoError(t, err)

	donator2 := chain.Env.NewSignatureSchemeWithFunds()
	req = solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Exactly what I needed!",
	).WithTransfer(balance.ColorIOTA, 69)
	_, err = chain.PostRequestSync(req, donator2)
	require.NoError(t, err)

	ret, err := chain.CallView(
		ScName, ViewDonations,
	)
	require.NoError(t, err)

	max, _, err := codec.DecodeInt64(ret[VarMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 69, max)

	tot, _, err := codec.DecodeInt64(ret[VarTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42+69, tot)

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator1.Address(), balance.ColorIOTA, solo.Saldo-42-1)
	// 69 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator2.Address(), balance.ColorIOTA, solo.Saldo-69-1)
	// 42+69 iota transferred to contract
	chain.AssertAccountBalance(common.ContractAccount, balance.ColorIOTA, 42+69)
	// returned 1 used for transaction to wallet accounts
	account1 := coretypes.NewAgentIDFromSigScheme(donator1)
	chain.AssertAccountBalance(account1, balance.ColorIOTA, 1)
	account2 := coretypes.NewAgentIDFromSigScheme(donator2)
	chain.AssertAccountBalance(account2, balance.ColorIOTA, 1)
}
