// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var auctioneer signaturescheme.SignatureScheme
var tokenColor balance.Color

func setupTest(t *testing.T) *solo.Chain {
	chain := common.DeployContract(t, ScName)

	// set up auctioneer account and mint some tokens to auction off
	auctioneer = chain.Env.NewSignatureSchemeWithFunds()
	newColor, err := chain.Env.MintTokens(auctioneer, 10)
	require.NoError(t, err)
	chain.Env.AssertAddressBalance(auctioneer.Address(), balance.ColorIOTA, solo.Supply-10)
	chain.Env.AssertAddressBalance(auctioneer.Address(), newColor, 10)
	tokenColor = newColor

	// start auction
	req := solo.NewCallParams(ScName, FuncStartAuction,
		ParamColor, tokenColor,
		ParamMinimumBid, 500,
		ParamDescription, "Cool tokens for sale!",
	).WithTransfers(map[balance.Color]int64{
		balance.ColorIOTA: 25, // deposit, must be >=minimum*margin
		tokenColor:        10, // the tokens to auction
	})
	_, err = chain.PostRequest(req, auctioneer)
	require.NoError(t, err)
	return chain
}

func TestDeploy(t *testing.T) {
	chain := common.DeployContract(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestFaStartAuction(t *testing.T) {
	chain := setupTest(t)

	// note 1 iota should be stuck in the delayed finalize_auction
	chain.AssertAccountBalance(common.ContractAccount, balance.ColorIOTA, 25-1)
	chain.AssertAccountBalance(common.ContractAccount, tokenColor, 10)

	// auctioneer sent 25 deposit + 10 tokenColor + used 1 for request
	chain.Env.AssertAddressBalance(auctioneer.Address(), balance.ColorIOTA, solo.Supply-35-1)
	// 1 used for request was sent back to auctioneer's account on chain
	account := coretypes.NewAgentIDFromSigScheme(auctioneer)
	chain.AssertAccountBalance(account, balance.ColorIOTA, 1)
}

func TestFaAuctionInfo(t *testing.T) {
	chain := setupTest(t)

	res, err := chain.CallView(
		ScName, ViewGetInfo,
		ParamColor, tokenColor,
	)
	require.NoError(t, err)
	account := coretypes.NewAgentIDFromSigScheme(auctioneer)
	requireAgent(t, res, VarCreator, account)
	requireInt64(t, res, VarBidders, 0)
}

func TestFaNoBids(t *testing.T) {
	chain := setupTest(t)

	// wait for finalize_auction
	chain.Env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(
		ScName, ViewGetInfo,
		ParamColor, tokenColor,
	)
	require.NoError(t, err)
	requireInt64(t, res, VarBidders, 0)
}

func TestFaOneBidTooLow(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncPlaceBid,
		ParamColor, tokenColor,
	).WithTransfer(balance.ColorIOTA, 100)
	_, err := chain.PostRequest(req, auctioneer)
	require.Error(t, err)

	// wait for finalize_auction
	chain.Env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(
		ScName, ViewGetInfo,
		ParamColor, tokenColor,
	)
	require.NoError(t, err)
	requireInt64(t, res, VarHighestBid, -1)
	requireInt64(t, res, VarBidders, 0)
}

func TestFaOneBid(t *testing.T) {
	chain := setupTest(t)

	bidder := chain.Env.NewSignatureSchemeWithFunds()
	req := solo.NewCallParams(ScName, FuncPlaceBid,
		ParamColor, tokenColor,
	).WithTransfer(balance.ColorIOTA, 500)
	_, err := chain.PostRequest(req, bidder)
	require.NoError(t, err)

	// wait for finalize_auction
	chain.Env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(
		ScName, ViewGetInfo,
		ParamColor, tokenColor,
	)
	require.NoError(t, err)
	requireInt64(t, res, VarBidders, 1)
	requireInt64(t, res, VarHighestBid, 500)
	requireAgent(t, res, VarHighestBidder, coretypes.NewAgentIDFromSigScheme(bidder))
}

func requireAgent(t *testing.T, res dict.Dict, key wasmlib.Key, expected coretypes.AgentID) {
	actual, exists, err := codec.DecodeAgentID(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireInt64(t *testing.T, res dict.Dict, key wasmlib.Key, expected int64) {
	actual, exists, err := codec.DecodeInt64(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireString(t *testing.T, res dict.Dict, key wasmlib.Key, expected string) {
	actual, exists, err := codec.DecodeString(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}
