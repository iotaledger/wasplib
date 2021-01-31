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

const ScName = "fairauction"

var WasmFile = wasmhost.WasmPath("fairauction_bg.wasm")

var bidderId [4]coretypes.AgentID
var bidderWallet [4]signaturescheme.SignatureScheme
var contractAgentID coretypes.AgentID
var contractID coretypes.ContractID
var creatorId coretypes.AgentID
var creatorWallet signaturescheme.SignatureScheme
var tokenColor balance.Color

func setupFaTest(t *testing.T) *govm.TestEnv {
	te := govm.NewTestEnv(t, ScName)

	contractID = coretypes.NewContractID(te.Chain.ChainID, coretypes.Hn(ScName))
	contractAgentID = coretypes.NewAgentIDFromContractID(contractID)

	creatorWallet = te.Env.NewSignatureSchemeWithFunds()
	creatorId = coretypes.NewAgentIDFromAddress(creatorWallet.Address())

	// set up 4 potential bidders
	for i := 0; i < 4; i++ {
		bidderWallet[i] = te.Env.NewSignatureSchemeWithFunds()
		bidderId[i] = coretypes.NewAgentIDFromAddress(bidderWallet[i].Address())
	}

	var err error
	tokenColor, err = te.Env.MintTokens(creatorWallet, 10)
	require.NoError(t, err)

	te.Env.AssertAddressBalance(creatorWallet.Address(), balance.ColorIOTA, 1337-10)
	te.Env.AssertAddressBalance(creatorWallet.Address(), tokenColor, 10)

	te.NewCallParams("start_auction",
		"color", tokenColor,
		"minimum", 500,
		"description", "Cool tokens for sale!").
		WithTransfers(map[balance.Color]int64{
			balance.ColorIOTA: 25, // deposit, must be >=minimum*margin
			tokenColor:        10,
		}).Post(0, creatorWallet)
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
	te.Chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 25-1)
	te.Chain.AssertAccountBalance(contractAgentID, tokenColor, 10)

	// creator sent 25 deposit + 10 tokenColor + used 1 for request
	te.Env.AssertAddressBalance(creatorWallet.Address(), balance.ColorIOTA, 1337-35-1)

	//TODO: seems silly to force creator to withdraw this 1 iota from chain account?
	// also look at how to send this back/retrieve it when creator was SC on other chain

	// 1 used for request was sent back to account on chain
	te.Chain.AssertAccountBalance(creatorId, balance.ColorIOTA, 1)
}

func TestFaAuctionInfo(t *testing.T) {
	te := setupFaTest(t)

	res := te.CallView("get_info", "color", tokenColor)
	requireAgent(t, res, "creator", creatorId)
	requireInt64(t, res, "bidders", 0)
}

func TestFaNoBids(t *testing.T) {
	te := setupFaTest(t)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView( "get_info", "color", tokenColor)
	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBidTooLow(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams( "place_bid", "color", tokenColor).
		PostFail(100, creatorWallet)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView( "get_info", "color", tokenColor)
	requireInt64(t, res, "highest_bid", -1)
	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBid(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams( "place_bid", "color", tokenColor).
		Post(500, bidderWallet[0])

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView( "get_info", "color", tokenColor)
	requireInt64(t, res, "bidders", 1)
	requireInt64(t, res, "highest_bid", 500)
	requireAgent(t, res, "highest_bidder", bidderId[0])
}

func TestFaClientAccess(t *testing.T) {
	te := setupFaTest(t)

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView( "get_info", "color", tokenColor)
	requireInt64(t, res, "bidders", 0)

	results := govm.GetClientMap(t, wasmhost.KeyResults, res)
	require.EqualValues(t, 0, results.GetInt(fairauction.KeyBidders).Value())
}

func TestFaClientFullAccess(t *testing.T) {
	te := setupFaTest(t)

	te.NewCallParams( "place_bid", "color", tokenColor).
		Post(500, bidderWallet[0])

	// wait for finalize_auction
	te.Env.AdvanceClockBy(61 * time.Minute)
	te.WaitForEmptyBacklog()

	res := te.CallView( "copy_all_state")
	state := govm.GetClientMap(t, wasmhost.KeyResults, res)
	auctions := state.GetMap(fairauction.KeyAuctions)
	color := client.NewScColor(tokenColor[:])
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(fairauction.KeyInfo)
	require.True(t, currentInfo.Exists())
	auction := fairauction.DecodeAuctionInfo(currentInfo.Value())
	require.EqualValues(t, 500, auction.HighestBid)
	require.EqualValues(t, bidderId[0][:], auction.HighestBidder.Bytes())
}
