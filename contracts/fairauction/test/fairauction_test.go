package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
	"github.com/iotaledger/wasplib/contracts/fairauction"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const scName = "fairauction"

var auctioneerId coretypes.AgentID
var auctioneerWallet signaturescheme.SignatureScheme
var contractAgentId coretypes.AgentID
var contractId coretypes.ContractID
var tokenColor balance.Color

func setupFaTest(t *testing.T) *govm.TestEnv {
	te := govm.NewTestEnv(t, scName)

	contractId = coretypes.NewContractID(te.Chain.ChainID, coretypes.Hn(scName))
	contractAgentId = coretypes.NewAgentIDFromContractID(contractId)

	auctioneerWallet = te.Wallet(0)
	auctioneerId = te.Agent(0)
	var err error
	tokenColor, err = te.Env.MintTokens(auctioneerWallet, 10)
	require.NoError(t, err)
	te.Env.AssertAddressBalance(auctioneerWallet.Address(), balance.ColorIOTA, 1337-10)
	te.Env.AssertAddressBalance(auctioneerWallet.Address(), tokenColor, 10)

	te.NewCallParams("start_auction",
		fairauction.ParamColor, tokenColor,
		fairauction.ParamMinimumBid, 500,
		fairauction.ParamDescription, "Cool tokens for sale!").
		WithTransfers(map[balance.Color]int64{
			balance.ColorIOTA: 25, // deposit, must be >=minimum*margin
			tokenColor:        10,
		}).Post(0, auctioneerWallet)
	return te
}

func requireAgent(t *testing.T, res dict.Dict, key kv.Key, expected coretypes.AgentID) {
	actual, exists, err := codec.DecodeAgentID(res.MustGet(key))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireInt64(t *testing.T, res dict.Dict, key kv.Key, expected int64) {
	actual, exists, err := codec.DecodeInt64(res.MustGet(key))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func requireString(t *testing.T, res dict.Dict, key kv.Key, expected string) {
	actual, exists, err := codec.DecodeString(res.MustGet(key))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, expected, actual)
}

func TestFaStartAuction(t *testing.T) {
	te := setupFaTest(t)

	// note 1 iota should be stuck in the delayed finalize_auction
	te.Chain.AssertAccountBalance(contractAgentId, balance.ColorIOTA, 25-1)
	te.Chain.AssertAccountBalance(contractAgentId, tokenColor, 10)

	// auctioneer sent 25 deposit + 10 tokenColor + used 1 for request
	te.Env.AssertAddressBalance(auctioneerWallet.Address(), balance.ColorIOTA, 1337-35-1)

	//TODO: it seems silly to force auctioneer to withdraw this 1 iota from chain account?
	// also look at how to send this back/retrieve it when auctioneer was SC on other chain

	// 1 used for request was sent back to auctioneer's account on chain
	te.Chain.AssertAccountBalance(auctioneerId, balance.ColorIOTA, 1)
}

func TestFaAuctionInfo(t *testing.T) {
	te := setupFaTest(t)

	res := te.CallView("get_info", fairauction.ParamColor, tokenColor)
	requireAgent(t, res, "creator", auctioneerId)
	requireInt64(t, res, "bidders", 0)
}

func TestFaNoBids(t *testing.T) {
	te := setupFaTest(t)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView("get_info", fairauction.ParamColor, tokenColor)
	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBidTooLow(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams("place_bid", fairauction.ParamColor, tokenColor).
		PostFail(100, auctioneerWallet)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView("get_info", fairauction.ParamColor, tokenColor)
	requireInt64(t, res, "highest_bid", -1)
	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBid(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams("place_bid", fairauction.ParamColor, tokenColor).
		Post(500, te.Wallet(1))

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView("get_info", fairauction.ParamColor, tokenColor)
	requireInt64(t, res, "bidders", 1)
	requireInt64(t, res, "highest_bid", 500)
	requireAgent(t, res, "highest_bidder", te.Agent(1))
}

func TestFaClientAccess(t *testing.T) {
	te := setupFaTest(t)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView("get_info", fairauction.ParamColor, tokenColor)
	requireInt64(t, res, "bidders", 0)

	results := te.GetClientMap(wasmhost.KeyResults, res)
	require.EqualValues(t, 0, results.GetInt(fairauction.VarBidders).Value())
}

func TestFaClientFullAccess(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams("place_bid", fairauction.ParamColor, tokenColor).
		Post(500, te.Wallet(1))

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	state := te.State()
	auctions := state.GetMap(fairauction.VarAuctions)
	color := client.NewScColor(tokenColor[:])
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(fairauction.VarInfo)
	require.True(t, currentInfo.Exists())
	auction := fairauction.DecodeAuctionInfo(currentInfo.Value())
	require.EqualValues(t, 500, auction.HighestBid)
	require.EqualValues(t, te.Agent(1).Bytes(), auction.HighestBidder.Bytes())
}
