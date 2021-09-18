package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasplib/contracts/rust/erc20"
	common2 "github.com/iotaledger/wasplib/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

var (
	chain   *solo.Chain
	creator *common2.SoloAgent
)

func setupTest(t *testing.T) {
	chain = common2.StartChain(t, "chain1")
	creator = common2.NewSoloAgent(chain.Env)
}

func setupErc20(t *testing.T) *common2.SoloContext {
	setupTest(t)
	init := erc20.ScFuncs.Init(nil)
	init.Params.Supply().SetValue(solo.Saldo)
	init.Params.Creator().SetValue(creator.ScAgentID())
	ctx := common2.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad, init.Func)
	require.NoError(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	totalSupply := erc20.ScFuncs.TotalSupply(ctx)
	totalSupply.Func.Call()
	require.NoError(t, ctx.Err)
	supply := totalSupply.Results.Supply()
	require.True(t, supply.Exists())
	require.EqualValues(t, solo.Saldo, supply.Value())

	checkErc20Balance(ctx, creator, solo.Saldo)
	return ctx
}

func checkErc20Balance(ctx *common2.SoloContext, account *common2.SoloAgent, amount uint64) {
	t := chain.Env.T
	balanceOf := erc20.ScFuncs.BalanceOf(ctx)
	balanceOf.Params.Account().SetValue(account.ScAgentID())
	balanceOf.Func.Call()
	require.NoError(t, ctx.Err)
	balance := balanceOf.Results.Amount()
	require.True(t, balance.Exists())
	require.EqualValues(t, amount, balance.Value())
}

func checkErc20Allowance(ctx *common2.SoloContext, account, delegation *common2.SoloAgent, amount uint64) {
	t := chain.Env.T
	allowance := erc20.ScFuncs.Allowance(ctx)
	allowance.Params.Account().SetValue(account.ScAgentID())
	allowance.Params.Delegation().SetValue(delegation.ScAgentID())
	allowance.Func.Call()
	require.NoError(t, ctx.Err)
	balance := allowance.Results.Amount()
	require.True(t, balance.Exists())
	require.EqualValues(t, amount, balance.Value())
}

func approve(ctx *common2.SoloContext, from, to *common2.SoloAgent, amount uint64) error {
	appr := erc20.ScFuncs.Approve(ctx.Sign(from))
	appr.Params.Delegation().SetValue(to.ScAgentID())
	appr.Params.Amount().SetValue(int64(amount))
	appr.Func.TransferIotas(1).Post()
	return ctx.Err
}

func transfer(ctx *common2.SoloContext, from, to *common2.SoloAgent, amount uint64) error {
	tx := erc20.ScFuncs.Transfer(ctx.Sign(from))
	tx.Params.Account().SetValue(to.ScAgentID())
	tx.Params.Amount().SetValue(int64(amount))
	tx.Func.TransferIotas(1).Post()
	return ctx.Err
}

func transferFrom(ctx *common2.SoloContext, from, to *common2.SoloAgent, amount uint64) error {
	tx := erc20.ScFuncs.TransferFrom(ctx.Sign(from))
	tx.Params.Account().SetValue(from.ScAgentID())
	tx.Params.Recipient().SetValue(to.ScAgentID())
	tx.Params.Amount().SetValue(int64(amount))
	tx.Func.TransferIotas(1).Post()
	return ctx.Err
}

func TestInitial(t *testing.T) {
	_ = setupErc20(t)
}

func TestTransferOk1(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, transfer(ctx, creator, user, 42))
	checkErc20Balance(ctx, creator, solo.Saldo-42)
	checkErc20Balance(ctx, user, 42)
}

func TestTransferOk2(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, transfer(ctx, creator, user, 42))
	checkErc20Balance(ctx, creator, solo.Saldo-42)
	checkErc20Balance(ctx, user, 42)

	require.NoError(t, transfer(ctx, user, creator, 42))
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)
}

func TestTransferNotEnoughFunds1(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)

	require.Error(t, transfer(ctx, creator, user, solo.Saldo+1))

	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)
}

func TestTransferNotEnoughFunds2(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)

	require.Error(t, transfer(ctx, user, creator, 1))

	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)
}

func TestNoAllowance(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()
	checkErc20Allowance(ctx, creator, user, 0)
}

func TestApprove(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, approve(ctx, creator, user, 100))

	checkErc20Allowance(ctx, creator, user, 100)
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)
}

func TestTransferFromOk1(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, approve(ctx, creator, user, 100))

	checkErc20Allowance(ctx, creator, user, 100)
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)

	require.NoError(t, transferFrom(ctx, creator, user, 50))

	checkErc20Allowance(ctx, creator, user, 50)
	checkErc20Balance(ctx, creator, solo.Saldo-50)
	checkErc20Balance(ctx, user, 50)
}

func TestTransferFromOk2(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, approve(ctx, creator, user, 100))

	checkErc20Allowance(ctx, creator, user, 100)
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)

	require.NoError(t, transferFrom(ctx, creator, user, 100))

	checkErc20Allowance(ctx, creator, user, 0)
	checkErc20Balance(ctx, creator, solo.Saldo-100)
	checkErc20Balance(ctx, user, 100)
}

func TestTransferFromFail(t *testing.T) {
	ctx := setupErc20(t)
	user := ctx.NewSoloAgent()

	require.NoError(t, approve(ctx, creator, user, 100))

	checkErc20Allowance(ctx, creator, user, 100)
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)

	require.Error(t, transferFrom(ctx, creator, user, 101))

	checkErc20Allowance(ctx, creator, user, 100)
	checkErc20Balance(ctx, creator, solo.Saldo)
	checkErc20Balance(ctx, user, 0)
}
