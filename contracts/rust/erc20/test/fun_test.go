package test

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	creator        *ed25519.KeyPair
	creatorAddr    ledgerstate.Address
	creatorAgentID *coretypes.AgentID
)

func deployErc20(t *testing.T) *solo.Chain {
	chain := common.StartChain(t, ScName)
	creator, creatorAddr = chain.Env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamSupply, solo.Saldo,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	res, err := chain.CallView(ScName, ViewTotalSupply)
	require.NoError(t, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(ParamSupply))
	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, sup, solo.Saldo)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	return chain
}

func checkErc20Balance(e *solo.Chain, account *coretypes.AgentID, amount uint64) {
	res, err := e.CallView(ScName, ViewBalanceOf,
		ParamAccount, account,
	)
	require.NoError(e.Env.T, err)
	sup, ok, err := codec.DecodeInt64(res.MustGet(ParamAmount))
	require.NoError(e.Env.T, err)
	require.True(e.Env.T, ok)
	require.EqualValues(e.Env.T, sup, amount)
}

func checkErc20Allowance(e *solo.Chain, account *coretypes.AgentID, delegation *coretypes.AgentID, amount int64) {
	res, err := e.CallView(ScName, ViewAllowance,
		ParamAccount, account,
		ParamDelegation, delegation,
	)
	require.NoError(e.Env.T, err)
	del, ok, err := codec.DecodeInt64(res.MustGet(ParamAmount))
	require.NoError(e.Env.T, err)
	require.True(e.Env.T, ok)
	require.EqualValues(e.Env.T, del, amount)
}

func TestInitial(t *testing.T) {
	_ = deployErc20(t)
}

func TestTransferOk1(t *testing.T) {
	chain := deployErc20(t)

	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	amount := uint64(42)

	req := solo.NewCallParams(ScName, FuncTransfer,
		ParamAccount, userAgentID,
		ParamAmount, amount,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo-amount)
	checkErc20Balance(chain, userAgentID, amount)
}

func TestTransferOk2(t *testing.T) {
	chain := deployErc20(t)

	user, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	amount := uint64(42)

	req := solo.NewCallParams(ScName, FuncTransfer,
		ParamAccount, userAgentID,
		ParamAmount, amount,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo-amount)
	checkErc20Balance(chain, userAgentID, amount)

	req = solo.NewCallParams(ScName, FuncTransfer,
		ParamAccount, creatorAgentID,
		ParamAmount, amount,
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, user)
	require.NoError(t, err)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferNotEnoughFunds1(t *testing.T) {
	chain := deployErc20(t)

	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	amount := int64(solo.Saldo + 1)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)

	req := solo.NewCallParams(ScName, FuncTransfer,
		ParamAccount, userAgentID,
		ParamAmount, amount,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.Error(t, err)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferNotEnoughFunds2(t *testing.T) {
	chain := deployErc20(t)

	user, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	amount := int64(1338)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)

	req := solo.NewCallParams(ScName, FuncTransfer,
		ParamAccount, creatorAgentID,
		ParamAmount, amount,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, user)
	require.Error(t, err)

	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestNoAllowance(t *testing.T) {
	chain := deployErc20(t)
	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	checkErc20Allowance(chain, creatorAgentID, userAgentID, 0)
}

func TestApprove(t *testing.T) {
	chain := deployErc20(t)
	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)

	req := solo.NewCallParams(ScName, FuncApprove,
		ParamDelegation, userAgentID,
		ParamAmount, 100,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)
}

func TestTransferFromOk1(t *testing.T) {
	chain := deployErc20(t)
	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)

	req := solo.NewCallParams(ScName, FuncApprove,
		ParamDelegation, userAgentID,
		ParamAmount, 100,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(ScName, FuncTransferFrom,
		ParamAccount, creatorAgentID,
		ParamRecipient, userAgentID,
		ParamAmount, 50,
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 50)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo-50)
	checkErc20Balance(chain, userAgentID, 50)
}

func TestTransferFromOk2(t *testing.T) {
	chain := deployErc20(t)
	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)

	req := solo.NewCallParams(ScName, FuncApprove,
		ParamDelegation, userAgentID,
		ParamAmount, 100,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(ScName, FuncTransferFrom,
		ParamAccount, creatorAgentID,
		ParamRecipient, userAgentID,
		ParamAmount, 100,
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 0)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo-100)
	checkErc20Balance(chain, userAgentID, 100)
}

func TestTransferFromFail(t *testing.T) {
	chain := deployErc20(t)
	_, userAddr := chain.Env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)

	req := solo.NewCallParams(ScName, FuncApprove,
		ParamDelegation, userAgentID,
		ParamAmount, 100,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, creator)
	require.NoError(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)

	req = solo.NewCallParams(ScName, FuncTransferFrom,
		ParamAccount, creatorAgentID,
		ParamRecipient, userAgentID,
		ParamAmount, 101,
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, creator)
	require.Error(t, err)

	checkErc20Allowance(chain, creatorAgentID, userAgentID, 100)
	checkErc20Balance(chain, creatorAgentID, solo.Saldo)
	checkErc20Balance(chain, userAgentID, 0)
}
