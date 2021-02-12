// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasmtest

package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/contracts/rust/donatewithfeedback"
	"github.com/iotaledger/wasp/contracts/testenv"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupDwfTest(t *testing.T) *testenv.TestEnv {
	te := testenv.NewTestEnv(t, donatewithfeedback.ScName)
	return te
}

func TestDwfDeploy(t *testing.T) {
	te := setupDwfTest(t)
	ret := te.CallView(donatewithfeedback.ViewDonations)
	results := te.Results(ret)
	max := results.GetInt(donatewithfeedback.VarMaxDonation)
	require.EqualValues(t, 0, max.Value())
	tot := results.GetInt(donatewithfeedback.VarTotalDonation)
	require.EqualValues(t, 0, tot.Value())
}

func TestDonateOnce(t *testing.T) {
	te := setupDwfTest(t)
	te.NewCallParams(donatewithfeedback.FuncDonate,
		donatewithfeedback.ParamFeedback, "Nice work!").
		Post(42, te.Wallet(0))
	ret := te.CallView(donatewithfeedback.ViewDonations)
	results := te.Results(ret)
	max := results.GetInt(donatewithfeedback.VarMaxDonation)
	require.EqualValues(t, 42, max.Value())
	tot := results.GetInt(donatewithfeedback.VarTotalDonation)
	require.EqualValues(t, 42, tot.Value())

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	te.Env.AssertAddressBalance(te.Wallet(0).Address(), balance.ColorIOTA, 1337-42-1)
	// 42 iota transferred to contract
	te.Chain.AssertAccountBalance(te.ContractAccount, balance.ColorIOTA, 42)
	// returned 1 used for transaction to wallet account
	te.Chain.AssertAccountBalance(te.Agent(0), balance.ColorIOTA, 1)
}

func TestDonateTwice(t *testing.T) {
	te := setupDwfTest(t)
	te.NewCallParams(donatewithfeedback.FuncDonate,
		donatewithfeedback.ParamFeedback, "Nice work!").
		Post(42, te.Wallet(0))
	te.NewCallParams(donatewithfeedback.FuncDonate,
		donatewithfeedback.ParamFeedback, "Exactly what I needed!").
		Post(69, te.Wallet(1))
	ret := te.CallView(donatewithfeedback.ViewDonations)
	results := te.Results(ret)
	max := results.GetInt(donatewithfeedback.VarMaxDonation)
	require.EqualValues(t, 69, max.Value())
	tot := results.GetInt(donatewithfeedback.VarTotalDonation)
	require.EqualValues(t, 42+69, tot.Value())

	// 42 iota transferred from wallet to contract plus 1 used for transaction
	te.Env.AssertAddressBalance(te.Wallet(0).Address(), balance.ColorIOTA, 1337-42-1)
	// 69 iota transferred from wallet to contract plus 1 used for transaction
	te.Env.AssertAddressBalance(te.Wallet(1).Address(), balance.ColorIOTA, 1337-69-1)
	// 42+69 iota transferred to contract
	te.Chain.AssertAccountBalance(te.ContractAccount, balance.ColorIOTA, 42+69)
	// returned 1 used for transaction to wallet accounts
	te.Chain.AssertAccountBalance(te.Agent(0), balance.ColorIOTA, 1)
	te.Chain.AssertAccountBalance(te.Agent(1), balance.ColorIOTA, 1)
}
