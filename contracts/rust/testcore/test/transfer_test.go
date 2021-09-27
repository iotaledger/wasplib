//nolint:dupl
package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/governance"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

// chainAccountBalances checks the balance of the chain account and the total
// balance of all accounts, taking any extra uploadWasm() into account
func chainAccountBalances(ctx *wasmsolo.SoloContext, w bool, chain, total uint64) {
	if w {
		// wasm setup takes 1 more iota than core setup due to uploadWasm()
		chain++
		total++
	}
	ctx.Chain.AssertCommonAccountIotas(chain)
	ctx.Chain.AssertTotalIotas(total)
}

// originatorBalanceReducedBy checks the balance of the originator address has
// reduced by the given amount, taking any extra uploadWasm() into account
func originatorBalanceReducedBy(ctx *wasmsolo.SoloContext, w bool, minus uint64) {
	if w {
		// wasm setup takes 1 more iota than core setup due to uploadWasm()
		minus++
	}
	ctx.Chain.Env.AssertAddressIotas(ctx.Chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-minus)
}

func TestDoNothing(t *testing.T) { run2(t, testDoNothing) }
func testDoNothing(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	nop := testcore.ScFuncs.DoNothing(ctx)
	nop.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
	require.EqualValues(t, 42, ctx.Balance(nil))
	require.EqualValues(t, 0, ctx.Balance(ctx.Originator()))
	originatorBalanceReducedBy(ctx, w, 2+42)
	chainAccountBalances(ctx, w, 2, 2+42)
}

func TestDoNothingUser(t *testing.T) { run2(t, testDoNothingUser) }
func testDoNothingUser(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	user := ctx.NewSoloAgent()
	nop := testcore.ScFuncs.DoNothing(ctx.Sign(user))
	nop.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
	require.EqualValues(t, solo.Saldo-42, user.Balance())
	require.EqualValues(t, 42, ctx.Balance(nil))

	require.EqualValues(t, 0, ctx.Balance(ctx.Originator()))
	require.EqualValues(t, 0, ctx.Balance(user))
	originatorBalanceReducedBy(ctx, w, 2)
	chainAccountBalances(ctx, w, 2, 2+42)
}

func TestWithdrawToAddress(t *testing.T) { run2(t, testWithdrawToAddress) }
func testWithdrawToAddress(t *testing.T, w bool) {
	ctx := setupTest(t, w)
	user := ctx.NewSoloAgent()

	nop := testcore.ScFuncs.DoNothing(ctx.Sign(user))
	nop.Func.TransferIotas(42).Post()
	require.NoError(t, ctx.Err)

	t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
	require.EqualValues(t, solo.Saldo-42, user.Balance())
	require.EqualValues(t, 42, ctx.Balance(nil))

	require.EqualValues(t, 0, ctx.Balance(ctx.Originator()))
	require.EqualValues(t, 0, ctx.Balance(user))
	originatorBalanceReducedBy(ctx, w, 2)
	chainAccountBalances(ctx, w, 2, 2+42)

	// send entire contract balance back to user
	// note that that includes the token that we transfer here
	xfer := testcore.ScFuncs.SendToAddress(ctx.Sign(ctx.Originator()))
	xfer.Params.Address().SetValue(user.ScAddress())
	xfer.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
	require.EqualValues(t, solo.Saldo-42+42+1, user.Balance())
	require.EqualValues(t, 0, ctx.Balance(nil))

	require.EqualValues(t, 0, ctx.Balance(ctx.Originator()))
	require.EqualValues(t, 0, ctx.Balance(user))
	originatorBalanceReducedBy(ctx, w, 2+1)
	chainAccountBalances(ctx, w, 2, 2)
}

func TestDoPanicUser(t *testing.T) { run2(t, testDoPanicUser) }
func testDoPanicUser(t *testing.T, w bool) {
	// ctx := setupTest(t, w)
	// user := setupDeployUser(t, ctx)

	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name).WithIotas(42)
	_, err := chain.PostRequestSync(req, user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)
}

func TestDoPanicUserFeeless(t *testing.T) { run2(t, testDoPanicUserFeeless) }
func testDoPanicUserFeeless(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name).WithIotas(42)
	_, err := chain.PostRequestSync(req, user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncWithdraw.Name).WithIotas(1)
	_, err = chain.PostRequestSync(req, user)
	require.NoError(t, err)

	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(4 + extraToken)
	chain.AssertTotalIotas(5 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo-1)
}

func TestDoPanicUserFee(t *testing.T) { run2(t, testDoPanicUserFee) }
func testDoPanicUserFee(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
		governance.ParamHname, cAID.Hname(),
		governance.ParamOwnerFee, 10,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(4 + extraToken)
	chain.AssertTotalIotas(5 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-1-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req = solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name).WithIotas(42)
	_, err = chain.PostRequestSync(req, user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(14 + extraToken)
	chain.AssertTotalIotas(15 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-1-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo-10)
}

func TestRequestToView(t *testing.T) { run2(t, testRequestToView) }
func testRequestToView(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	// sending request to the view entry point should return an error and invoke fallback for tokens
	req := solo.NewCallParams(ScName, sbtestsc.FuncJustView.Name).WithIotas(42)
	_, err := chain.PostRequestSync(req, user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertIotas(userAgentID, 0)
	chain.AssertIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertTotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)
}
