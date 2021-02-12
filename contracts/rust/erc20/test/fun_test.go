package erc20

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/contracts/rust/erc20"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

const supply = int64(1337)

var (
	creator        signaturescheme.SignatureScheme
	creatorAgentID coretypes.AgentID
)

func deployErc20(t *testing.T) *solo.Chain {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	creator = env.NewSignatureSchemeWithFunds()
	creatorAgentID = coretypes.NewAgentIDFromAddress(creator.Address())
	err := chain.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, supply,
		PARAM_CREATOR, creatorAgentID,
	)
	require.NoError(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 5, len(rec))

	res, err := chain.CallView(erc20name, erc20.ViewTotalSupply)
	require.NoError(t, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(PARAM_SUPPLY))
	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, sup, supply)

	checkErc20Balance(chain, creatorAgentID, supply)
	return chain
}

func checkErc20Balance(e *solo.Chain, account coretypes.AgentID, amount int64) {
	res, err := e.CallView(erc20name, erc20.ViewBalanceOf, PARAM_ACCOUNT, account)
	require.NoError(e.Env.T, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(PARAM_AMOUNT))
	require.NoError(e.Env.T, err)
	require.True(e.Env.T, ok)
	require.EqualValues(e.Env.T, sup, amount)
}

func checkErc20Allowance(e *solo.Chain, account coretypes.AgentID, delegation coretypes.AgentID, amount int64) {
	res, err := e.CallView(erc20name, erc20.ViewAllowance, PARAM_ACCOUNT, account, PARAM_DELEGATION, delegation)
	require.NoError(e.Env.T, err)
	del, ok, err := codec.DecodeInt64(res.MustGet(PARAM_AMOUNT))
	require.NoError(e.Env.T, err)
	require.True(e.Env.T, ok)
	require.EqualValues(e.Env.T, del, amount)
}

func TestInitial(t *testing.T) {
	_ = deployErc20(t)
}

func TestTransferOk1(t *testing.T) {
	chain := deployErc20(t)

	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(42)

	req := solo.NewCallParams(erc20name, erc20.FuncTransfer,
		PARAM_ACCOUNT, userAgentID,
		PARAM_AMOUNT, amount,
	)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, supply-amount)
	checkErc20Balance(chain, userAgentID, amount)
}

func TestTransferOk2(t *testing.T) {
	chain := deployErc20(t)

	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(42)

	req := solo.NewCallParams(erc20name, erc20.FuncTransfer, PARAM_ACCOUNT, userAgentID, PARAM_AMOUNT, amount)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, supply-amount)
	checkErc20Balance(chain, userAgentID, amount)

	req = solo.NewCallParams(erc20name, erc20.FuncTransfer, PARAM_ACCOUNT, creatorAgentID, PARAM_AMOUNT, amount)
	_, err = chain.PostRequest(req, user)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferNotEnoughFunds1(t *testing.T) {
	chain := deployErc20(t)

	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(1338)

	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)

	req := solo.NewCallParams(erc20name, erc20.FuncTransfer, PARAM_ACCOUNT, userAgentID, PARAM_AMOUNT, amount)
	_, err := chain.PostRequest(req, creator)
	require.Error(t, err)

	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferNotEnoughFunds2(t *testing.T) {
	chain := deployErc20(t)

	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	amount := int64(1338)

	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)

	req := solo.NewCallParams(erc20name, erc20.FuncTransfer, PARAM_ACCOUNT, creatorAgentID, PARAM_AMOUNT, amount)
	_, err := chain.PostRequest(req, user)
	require.Error(t, err)

	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestNoAllowance(t *testing.T) {
	chain := deployErc20(t)
	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())
	checkErc20Allowance(chain, creatorAgentID, userAgentID, 0)
}

func TestApprove(t *testing.T) {
	chain := deployErc20(t)
	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := solo.NewCallParams(erc20name, erc20.FuncApprove, PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferFromOk1(t *testing.T) {
	chain := deployErc20(t)
	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := solo.NewCallParams(erc20name, erc20.FuncApprove, PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(erc20name, erc20.FuncTransferFrom,
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 50,
	)
	_, err = chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 50)
	checkErc20Balance(chain, creatorAgentID, supply-50)
	checkErc20Balance(chain, userAgentID, 50)
}

func TestTransferFromOk2(t *testing.T) {
	chain := deployErc20(t)
	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := solo.NewCallParams(erc20name, erc20.FuncApprove, PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(erc20name, erc20.FuncTransferFrom,
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 100,
	)
	_, err = chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 0)
	checkErc20Balance(chain, creatorAgentID, supply-100)
	checkErc20Balance(chain, userAgentID, 100)
}

func TestTransferFromFail(t *testing.T) {
	chain := deployErc20(t)
	user := chain.Env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(user.Address())

	req := solo.NewCallParams(erc20name, erc20.FuncApprove, PARAM_DELEGATION, userAgentID, PARAM_AMOUNT, 100)
	_, err := chain.PostRequest(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(erc20name, erc20.FuncTransferFrom,
		PARAM_ACCOUNT, creatorAgentID,
		PARAM_RECIPIENT, userAgentID,
		PARAM_AMOUNT, 101,
	)
	_, err = chain.PostRequest(req, creator)
	require.Error(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, supply)
	checkErc20Balance(chain, userAgentID, 0)
}
