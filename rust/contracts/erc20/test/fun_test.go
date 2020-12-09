package erc20

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/vm/alone"
	"github.com/stretchr/testify/require"
	"testing"
)

const supply = int64(1337)

var (
	creator        signaturescheme.SignatureScheme
	creatorAgentID coretypes.AgentID
)

func deployErc20(t *testing.T) *alone.Env {
	e := alone.New(t, false, false)
	creator = e.NewSigScheme()
	creatorAgentID = coretypes.NewAgentIDFromAddress(creator.Address())
	err := e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, supply,
		PARAM_CREATOR, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 4, len(rec))

	res, err := e.CallView(erc20name, "totalSupply")
	require.NoError(t, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(PARAM_SUPPLY))
	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, sup, supply)

	checkErc20Balance(e, creatorAgentID, supply)
	return e
}

func checkErc20Balance(e *alone.Env, account coretypes.AgentID, amount int64) {
	res, err := e.CallView(erc20name, "balanceOf", PARAM_ACCOUNT, account)
	require.NoError(e.T, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(PARAM_AMOUNT))
	require.NoError(e.T, err)
	require.True(e.T, ok)
	require.EqualValues(e.T, sup, amount)
}

func checkErc20Allowance(e *alone.Env, account coretypes.AgentID, delegation coretypes.AgentID, amount int64) {
	res, err := e.CallView(erc20name, "allowance", PARAM_ACCOUNT, account, PARAM_DELEGATION, delegation)
	require.NoError(e.T, err)
	del, ok, err := codec.DecodeInt64(res.MustGet(PARAM_AMOUNT))
	require.NoError(e.T, err)
	require.True(e.T, ok)
	require.EqualValues(e.T, del, amount)
}

func TestInitial(t *testing.T) {
	_ = deployErc20(t)
}

func TestTransferOk1(t *testing.T) {
	e := deployErc20(t)

	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(42)

	req := alone.NewCall(erc20name, "transfer", PARAM_ACCOUNT, userAgentID, PARAM_AMOUNT, amount)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Balance(e, creatorAgentID, supply-amount)
	checkErc20Balance(e, userAgentID, amount)
}

func TestTransferOk2(t *testing.T) {
	e := deployErc20(t)

	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(42)

	req := alone.NewCall(erc20name, "transfer", PARAM_ACCOUNT, userAgentID, PARAM_AMOUNT, amount)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Balance(e, creatorAgentID, supply-amount)
	checkErc20Balance(e, userAgentID, amount)

	req = alone.NewCall(erc20name, "transfer", PARAM_ACCOUNT, creatorAgentID, PARAM_AMOUNT, amount)
	_, err = e.PostRequest(req, user)
	require.NoError(t, err)

	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)
}

func TestTransferNotEnoughFunds1(t *testing.T) {
	e := deployErc20(t)

	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(1338)

	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)

	req := alone.NewCall(erc20name, "transfer", PARAM_ACCOUNT, userAgentID, PARAM_AMOUNT, amount)
	_, err := e.PostRequest(req, creator)
	require.Error(t, err)

	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)
}

func TestTransferNotEnoughFunds2(t *testing.T) {
	e := deployErc20(t)

	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(1338)

	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)

	req := alone.NewCall(erc20name, "transfer", PARAM_ACCOUNT, creatorAgentID, PARAM_AMOUNT, amount)
	_, err := e.PostRequest(req, user)
	require.Error(t, err)

	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)
}

func TestNoAllowance(t *testing.T) {
	e := deployErc20(t)
	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	checkErc20Allowance(e, creatorAgentID, userAgentID, 0)
}

func TestApprove(t *testing.T) {
	e := deployErc20(t)
	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := alone.NewCall(erc20name, "approve", PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 100)
	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)
}

func TestTransferFromOk1(t *testing.T) {
	e := deployErc20(t)
	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := alone.NewCall(erc20name, "approve", PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 100)
	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)

	req = alone.NewCall(erc20name, "transferFrom",
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 50,
	)
	_, err = e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 50)
	checkErc20Balance(e, creatorAgentID, supply-50)
	checkErc20Balance(e, userAgentID, 50)
}

func TestTransferFromOk2(t *testing.T) {
	e := deployErc20(t)
	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := alone.NewCall(erc20name, "approve", PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 100)
	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)

	req = alone.NewCall(erc20name, "transferFrom",
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 100,
	)
	_, err = e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 0)
	checkErc20Balance(e, creatorAgentID, supply-100)
	checkErc20Balance(e, userAgentID, 100)
}

func TestTransferFromFail(t *testing.T) {
	e := deployErc20(t)
	user := e.NewSigScheme()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := alone.NewCall(erc20name, "approve", PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := e.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 100)
	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)

	req = alone.NewCall(erc20name, "transferFrom",
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 101,
	)
	_, err = e.PostRequest(req, creator)
	require.Error(t, err)

	checkErc20Allowance(e, creatorAgentID, userAgentID, 100)
	checkErc20Balance(e, creatorAgentID, supply)
	checkErc20Balance(e, userAgentID, 0)
}
