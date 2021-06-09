// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
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
		ScName, ViewDonationInfo,
	)
	require.NoError(t, err)

	count, _, err := codec.DecodeInt64(ret[ResultCount])
	require.NoError(t, err)
	require.EqualValues(t, 0, count)

	max, _, err := codec.DecodeInt64(ret[ResultMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 0, max)

	tot, _, err := codec.DecodeInt64(ret[ResultTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 0, tot)
}

func TestDonateOnce(t *testing.T) {
	chain := setupTest(t)

	donator1, donator1Addr := chain.Env.NewKeyPairWithFunds()
	req := solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Nice work!",
	).WithIotas(42)
	_, err := chain.PostRequestSync(req, donator1)
	require.NoError(t, err)

	ret, err := chain.CallView(
		ScName, ViewDonationInfo,
	)
	require.NoError(t, err)

	count, _, err := codec.DecodeInt64(ret[ResultCount])
	require.NoError(t, err)
	require.EqualValues(t, 1, count)

	max, _, err := codec.DecodeInt64(ret[VarMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42, max)

	tot, _, err := codec.DecodeInt64(ret[VarTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42, tot)

	// 42 iota transferred from wallet to contract
	chain.Env.AssertAddressBalance(donator1Addr, ledgerstate.ColorIOTA, solo.Saldo-42)
	// 42 iota transferred to contract
	chain.AssertAccountBalance(chain.ContractAgentID(ScName), ledgerstate.ColorIOTA, 42)
	// returned 1 used for transaction to wallet account
	account1 := coretypes.NewAgentID(donator1Addr, 0)
	chain.AssertAccountBalance(account1, ledgerstate.ColorIOTA, 0)
}

func TestDonateTwice(t *testing.T) {
	chain := setupTest(t)

	donator1, donator1Addr := chain.Env.NewKeyPairWithFunds()
	req := solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Nice work!",
	).WithIotas(42)
	_, err := chain.PostRequestSync(req, donator1)
	require.NoError(t, err)

	donator2, donator2Addr := chain.Env.NewKeyPairWithFunds()
	req = solo.NewCallParams(ScName, FuncDonate,
		ParamFeedback, "Exactly what I needed!",
	).WithIotas(69)
	_, err = chain.PostRequestSync(req, donator2)
	require.NoError(t, err)

	ret, err := chain.CallView(
		ScName, ViewDonationInfo,
	)
	require.NoError(t, err)

	count, _, err := codec.DecodeInt64(ret[ResultCount])
	require.NoError(t, err)
	require.EqualValues(t, 2, count)

	max, _, err := codec.DecodeInt64(ret[VarMaxDonation])
	require.NoError(t, err)
	require.EqualValues(t, 69, max)

	tot, _, err := codec.DecodeInt64(ret[VarTotalDonation])
	require.NoError(t, err)
	require.EqualValues(t, 42+69, tot)

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator1Addr, ledgerstate.ColorIOTA, solo.Saldo-42)
	// 69 iota transferred from wallet to contract plus 1 used for transaction
	chain.Env.AssertAddressBalance(donator2Addr, ledgerstate.ColorIOTA, solo.Saldo-69)
	// 42+69 iota transferred to contract
	chain.AssertAccountBalance(chain.ContractAgentID(ScName), ledgerstate.ColorIOTA, 42+69)
	// returned 1 used for transaction to wallet accounts
	account1 := coretypes.NewAgentID(donator1Addr, 0)
	chain.AssertAccountBalance(account1, ledgerstate.ColorIOTA, 0)
	account2 := coretypes.NewAgentID(donator2Addr, 0)
	chain.AssertAccountBalance(account2, ledgerstate.ColorIOTA, 0)
}
