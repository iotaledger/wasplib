// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/fairauction"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/stretchr/testify/require"
)

var (
	auctioneer     *ed25519.KeyPair
	auctioneerAddr ledgerstate.Address
	tokenColor     ledgerstate.Color
)

func setupTest(t *testing.T) *common.SoloContext {
	chain := common.StartChainAndDeployWasmContractByName(t, fairauction.ScName)

	// set up auctioneer account and mint some tokens to auction off
	auctioneer, auctioneerAddr = chain.Env.NewKeyPairWithFunds()
	newColor, err := chain.Env.MintTokens(auctioneer, 10)
	require.NoError(t, err)
	tokenColor = newColor
	chain.Env.AssertAddressBalance(auctioneerAddr, ledgerstate.ColorIOTA, solo.Saldo-10)
	chain.Env.AssertAddressBalance(auctioneerAddr, tokenColor, 10)

	ctx := common.NewSoloContext(fairauction.ScName, fairauction.OnLoad, chain)
	auctionColor := ctx.ScColor(tokenColor)

	startAuction := fairauction.NewStartAuctionCall(ctx.SignWith(auctioneer))
	startAuction.Params.Color().SetValue(auctionColor)
	startAuction.Params.MinimumBid().SetValue(500)
	startAuction.Params.Description().SetValue("Cool tokens for sale!")
	transfer := ctx.Transfer()
	transfer.Set(wasmlib.IOTA, 25) // deposit, must be >=minimum*margin
	transfer.Set(auctionColor, 10) // the tokens to auction
	startAuction.Func.Transfer(transfer).Post()
	require.NoError(t, ctx.Err)
	return ctx
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, fairauction.ScName)
	_, err := chain.FindContract(fairauction.ScName)
	require.NoError(t, err)
}

func TestFaStartAuction(t *testing.T) {
	ctx := setupTest(t)
	chain := ctx.Chain

	// note 1 iota should be stuck in the delayed finalize_auction
	chain.AssertAccountBalance(chain.ContractAgentID(fairauction.ScName), ledgerstate.ColorIOTA, 25-1)
	chain.AssertAccountBalance(chain.ContractAgentID(fairauction.ScName), tokenColor, 10)

	// auctioneer sent 25 deposit + 10 tokenColor + used 1 for request
	chain.Env.AssertAddressBalance(auctioneerAddr, ledgerstate.ColorIOTA, solo.Saldo-35)
	// 1 used for request was sent back to auctioneer's account on chain
	account := coretypes.NewAgentID(auctioneerAddr, 0)
	chain.AssertAccountBalance(account, ledgerstate.ColorIOTA, 0)

	// remove delayed finalize_auction from backlog
	chain.Env.AdvanceClockBy(61 * time.Minute)
	require.True(t, ctx.WaitForRequestsThrough(5))
}

func TestFaAuctionInfo(t *testing.T) {
	ctx := setupTest(t)

	getInfo := fairauction.NewGetInfoCall(ctx)
	getInfo.Params.Color().SetValue(ctx.ScColor(tokenColor))
	getInfo.Func.Call()

	require.NoError(t, ctx.Err)
	require.EqualValues(t, ctx.ScAddress(auctioneerAddr).AsAgentID(), getInfo.Results.Creator().Value())
	require.EqualValues(t, 0, getInfo.Results.Bidders().Value())

	// remove delayed finalize_auction from backlog
	ctx.Chain.Env.AdvanceClockBy(61 * time.Minute)
	require.True(t, ctx.WaitForRequestsThrough(5))
}

func TestFaNoBids(t *testing.T) {
	ctx := setupTest(t)

	// wait for finalize_auction
	ctx.Chain.Env.AdvanceClockBy(61 * time.Minute)
	require.True(t, ctx.WaitForRequestsThrough(5))

	getInfo := fairauction.NewGetInfoCall(ctx)
	getInfo.Params.Color().SetValue(ctx.ScColor(tokenColor))
	getInfo.Func.Call()

	require.NoError(t, ctx.Err)
	require.EqualValues(t, 0, getInfo.Results.Bidders().Value())
}

func TestFaOneBidTooLow(t *testing.T) {
	ctx := setupTest(t)
	chain := ctx.Chain

	bidder, _ := chain.Env.NewKeyPairWithFunds()
	placeBid := fairauction.NewPlaceBidCall(ctx.SignWith(bidder))
	placeBid.Params.Color().SetValue(ctx.ScColor(tokenColor))
	placeBid.Func.TransferIotas(100).Post()
	require.Error(t, ctx.Err)

	// wait for finalize_auction
	chain.Env.AdvanceClockBy(61 * time.Minute)
	require.True(t, ctx.WaitForRequestsThrough(6))

	getInfo := fairauction.NewGetInfoCall(ctx)
	getInfo.Params.Color().SetValue(ctx.ScColor(tokenColor))
	getInfo.Func.Call()

	require.NoError(t, ctx.Err)
	require.EqualValues(t, 0, getInfo.Results.Bidders().Value())
	require.EqualValues(t, -1, getInfo.Results.HighestBid().Value())
}

func TestFaOneBid(t *testing.T) {
	ctx := setupTest(t)
	chain := ctx.Chain

	bidder, _ := chain.Env.NewKeyPairWithFunds()
	placeBid := fairauction.NewPlaceBidCall(ctx.SignWith(bidder))
	placeBid.Params.Color().SetValue(ctx.ScColor(tokenColor))
	placeBid.Func.TransferIotas(500).Post()
	require.NoError(t, ctx.Err)

	// wait for finalize_auction
	chain.Env.AdvanceClockBy(61 * time.Minute)
	require.True(t, ctx.WaitForRequestsThrough(6))

	getInfo := fairauction.NewGetInfoCall(ctx)
	getInfo.Params.Color().SetValue(ctx.ScColor(tokenColor))
	getInfo.Func.Call()

	require.NoError(t, ctx.Err)
	require.EqualValues(t, 1, getInfo.Results.Bidders().Value())
	require.EqualValues(t, 500, getInfo.Results.HighestBid().Value())
	require.Equal(t, ctx.Address().AsAgentID(), getInfo.Results.HighestBidder().Value())
}
