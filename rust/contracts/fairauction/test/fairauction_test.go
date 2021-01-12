package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/client"
	"github.com/iotaledger/wasplib/contracts/fairauction"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const SC_NAME = "fairauction"
const WASM_FILE = "../pkg/fairauction_bg.wasm"

var bidderId [4]coretypes.AgentID
var bidderWallet [4]signaturescheme.SignatureScheme
var chain *solo.Chain
var contractAgentID coretypes.AgentID
var contractID coretypes.ContractID
var creatorId coretypes.AgentID
var creatorWallet signaturescheme.SignatureScheme
var env *solo.Solo
var tokenColor balance.Color

func setupFaTest(t *testing.T) {
	wasmhost.HostTracing = true
	env = solo.New(t, true, false)
	chain = env.NewChain(nil, "ch1")
	contractID = coretypes.NewContractID(chain.ChainID, coretypes.Hn(SC_NAME))
	contractAgentID = coretypes.NewAgentIDFromContractID(contractID)

	err := chain.DeployWasmContract(nil, SC_NAME, WASM_FILE)
	require.NoError(t, err)

	creatorWallet = env.NewSignatureSchemeWithFunds()
	creatorId = coretypes.NewAgentIDFromAddress(creatorWallet.Address())

	// set up 4 potential bidders
	for i := 0; i < 4; i++ {
		bidderWallet[i] = env.NewSignatureSchemeWithFunds()
		bidderId[i] = coretypes.NewAgentIDFromAddress(bidderWallet[i].Address())
	}

	tokenColor, err = env.MintTokens(creatorWallet, 10)
	require.NoError(t, err)
	env.AssertAddressBalance(creatorWallet.Address(), balance.ColorIOTA, 1337-10)
	env.AssertAddressBalance(creatorWallet.Address(), tokenColor, 10)

	req := solo.NewCall(SC_NAME, "start_auction", "color", tokenColor,
		"minimum", 500, "description", "Cool tokens for sale!").
		WithTransfers(map[balance.Color]int64{
			balance.ColorIOTA: 25, // deposit, must be >=minimum*margin
			tokenColor:        10,
		})
	_, err = chain.PostRequest(req, creatorWallet)
	require.NoError(t, err)
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
	setupFaTest(t)

	// note 1 iota should be stuck in the delayed finalize_auction
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 25-1)
	chain.AssertAccountBalance(contractAgentID, tokenColor, 10)

	// creator sent 25 deposit + 10 tokenColor + used 1 for request
	env.AssertAddressBalance(creatorWallet.Address(), balance.ColorIOTA, 1337-35-1)

	//TODO: seems silly to force creator to withdraw this 1 iota from chain account?
    // also look at how to send this back/retrieve it when creator was SC on other chain

	// 1 used for request was sent back to account on chain
	chain.AssertAccountBalance(creatorId, balance.ColorIOTA, 1)
}

func TestFaAuctionInfo(t *testing.T) {
	setupFaTest(t)

	res, err := chain.CallView(SC_NAME, "get_info", "color", tokenColor)
	require.NoError(t, err)

	requireAgent(t, res, "creator", creatorId)
	requireInt64(t, res, "bidders", 0)
}

func TestFaNoBids(t *testing.T) {
	setupFaTest(t)

	// wait for finalize_auction
	env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(SC_NAME, "get_info", "color", tokenColor)
	require.NoError(t, err)

	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBidTooLow(t *testing.T) {
	setupFaTest(t)

	req := solo.NewCall(SC_NAME, "place_bid", "color", tokenColor).
		WithTransfer(balance.ColorIOTA, 100)
	_, err := chain.PostRequest(req, creatorWallet)
	require.Error(t, err)

	// wait for finalize_auction
	env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(SC_NAME, "get_info", "color", tokenColor)
	require.NoError(t, err)

	requireInt64(t, res, "highest_bid", -1)
	requireInt64(t, res, "bidders", 0)
}

func TestFaOneBid(t *testing.T) {
	setupFaTest(t)

	req := solo.NewCall(SC_NAME, "place_bid", "color", tokenColor).
		WithTransfer(balance.ColorIOTA, 500)
	_, err := chain.PostRequest(req, bidderWallet[0])
	require.NoError(t, err)

	// wait for finalize_auction
	env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(SC_NAME, "get_info", "color", tokenColor)
	require.NoError(t, err)

	requireInt64(t, res, "bidders", 1)
	requireInt64(t, res, "highest_bid", 500)
	requireAgent(t, res, "highest_bidder", bidderId[0])
}

func TestFaClientAccess(t *testing.T) {
	setupFaTest(t)

	// wait for finalize_auction
	env.AdvanceClockBy(61 * time.Minute)
	chain.WaitForEmptyBacklog()

	res, err := chain.CallView(SC_NAME, "get_info", "color", tokenColor)
	require.NoError(t, err)

	requireInt64(t, res, "bidders", 0)

	dict := getClientMap(t, wasmhost.KeyResults, res)
	require.EqualValues(t, 0, dict.GetInt(fairauction.KeyBidders).Value())
}

func getClientMap(t *testing.T, keyId int32, kvStore kv.KVStore) client.ScImmutableMap {
	logger := testutil.NewLogger(t, "04:05.000")
	host := &wasmhost.KvStoreHost{}
	null := wasmproc.NewNullObject(host)
	root := wasmproc.NewScDictFromKvStore(host, kvStore)
	host.Init(null, root, logger)
	root.InitObj(1, keyId, root)
	logger.Info("Direct access to %s", host.GetKeyStringFromId(keyId))
	client.ConnectHost(host)
	return client.Root.Immutable()
}
