// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var auctioneer *ed25519.KeyPair
var auctioneerAddr  ledgerstate.Address
var tokenColor ledgerstate.Color

func setupTest(t *testing.T) *solo.Chain {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)

	// set up auctioneer account and mint some tokens to auction off
	auctioneer,auctioneerAddr = chain.Env.NewKeyPairWithFunds()
	newColor, err := chain.Env.MintTokens(auctioneer, 10)
	require.NoError(t, err)
	chain.Env.AssertAddressBalance(auctioneerAddr, ledgerstate.ColorIOTA, solo.Saldo-10)
	chain.Env.AssertAddressBalance(auctioneerAddr, newColor, 10)
	tokenColor = newColor

	// start auction
	req := solo.NewCallParams(ScName, FuncStartAuction,
		ParamColor, tokenColor,
		ParamMinimumBid, 500,
		ParamDescription, "Cool tokens for sale!",
	).WithTransfers(map[ledgerstate.Color]uint64{
		ledgerstate.ColorIOTA: 25, // deposit, must be >=minimum*margin
		tokenColor:        10, // the tokens to auction
	})
	_, err = chain.PostRequestSync(req, auctioneer)
	require.NoError(t, err)
	return chain
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestFaStartAuction(t *testing.T) {
	chain := setupTest(t)

	// note 1 iota should be stuck in the delayed finalize_auction
	chain.AssertAccountBalance(chain.ContractAgentID(ScName), ledgerstate.ColorIOTA, 25-1)
	chain.AssertAccountBalance(chain.ContractAgentID(ScName), tokenColor, 10)

	// auctioneer sent 25 deposit + 10 tokenColor + used 1 for request
	chain.Env.AssertAddressBalance(auctioneerAddr, ledgerstate.ColorIOTA, solo.Saldo-35)
	// 1 used for request was sent back to auctioneer's account on chain
	account := coretypes.NewAgentID(auctioneerAddr, 0)
	chain.AssertAccountBalance(account, ledgerstate.ColorIOTA, 0)

	// remove delayed finalize_auction from backlog
	chain.Env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()
}

func TestFaAuctionInfo(t *testing.T) {
	chain := setupTest(t)

	res, err := chain.CallView(
		ScName, ViewGetInfo,
		ParamColor, tokenColor,
	)
	require.NoError(t, err)
	account := coretypes.NewAgentID(auctioneerAddr, 0)
	requireAgent(t, res, VarCreator, *account)
	requireInt64(t, res, VarBidders, 0)

	// remove delayed finalize_auction from backlog
	chain.Env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()
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
	).WithIotas(100)
	_, err := chain.PostRequestSync(req, auctioneer)
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

	bidder, bidderAddr := chain.Env.NewKeyPairWithFunds()
	req := solo.NewCallParams(ScName, FuncPlaceBid,
		ParamColor, tokenColor,
	).WithIotas(500)
	_, err := chain.PostRequestSync(req, bidder)
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
	requireAgent(t, res, VarHighestBidder, *coretypes.NewAgentID(bidderAddr, 0))
}

func requireAgent(t *testing.T, res dict.Dict, key string, expected coretypes.AgentID) {
	actual, exists, err := codec.DecodeAgentID(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireInt64(t *testing.T, res dict.Dict, key string, expected int64) {
	actual, exists, err := codec.DecodeInt64(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireString(t *testing.T, res dict.Dict, key string, expected string) {
	actual, exists, err := codec.DecodeString(res.MustGet(kv.Key(key)))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}
