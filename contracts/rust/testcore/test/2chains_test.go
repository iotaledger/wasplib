package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func Test2Chains(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		core.PrintWellKnownHnames()

		chain1 := wasmsolo.StartChain(t, "chain1")
		chain1.CheckAccountLedger()

		chain2 := wasmsolo.StartChain(t, "chain2", chain1.Env)
		chain2.CheckAccountLedger()

		user := wasmsolo.NewSoloAgent(chain1.Env)
		require.EqualValues(t, solo.Saldo, user.Balance())

		ctx1 := setupTestForChain(t, w, chain1, nil)
		require.NoError(t, ctx1.Err)
		ctx2 := setupTestForChain(t, w, chain2, nil)
		require.NoError(t, ctx2.Err)

		require.EqualValues(t, 0, ctx1.Balance(user))
		require.EqualValues(t, 0, ctx1.Balance(ctx1.Agent()))
		require.EqualValues(t, 0, ctx1.Balance(ctx2.Agent()))
		chainAccountBalances(ctx1, w, 2, 2)

		require.EqualValues(t, 0, ctx2.Balance(user))
		require.EqualValues(t, 0, ctx2.Balance(ctx1.Agent()))
		require.EqualValues(t, 0, ctx2.Balance(ctx2.Agent()))
		chainAccountBalances(ctx2, w, 2, 2)

		deposit(t, ctx1, user, ctx2.Agent(), 42)
		require.EqualValues(t, solo.Saldo-42, user.Balance())

		require.EqualValues(t, 0, ctx1.Balance(user))
		require.EqualValues(t, 0, ctx1.Balance(ctx1.Agent()))
		require.EqualValues(t, 0+42, ctx1.Balance(ctx2.Agent()))
		chainAccountBalances(ctx1, w, 2, 2+42)

		require.EqualValues(t, 0, ctx2.Balance(user))
		require.EqualValues(t, 0, ctx2.Balance(ctx1.Agent()))
		require.EqualValues(t, 0, ctx2.Balance(ctx2.Agent()))
		chainAccountBalances(ctx2, w, 2, 2)

		f := testcore.ScFuncs.WithdrawToChain(ctx2.Sign(user))
		f.Params.ChainID().SetValue(ctx1.ChainID())
		f.Func.TransferIotas(1).Post()
		require.NoError(t, ctx2.Err)

		require.True(t, ctx1.WaitForPendingRequests(1))
		require.True(t, ctx2.WaitForPendingRequests(1))

		require.EqualValues(t, solo.Saldo-42-1, user.Balance())

		t.Logf("dump chain1 accounts:\n%s", ctx1.Chain.DumpAccounts())
		require.EqualValues(t, 0, ctx1.Balance(user))
		require.EqualValues(t, 0, ctx1.Balance(ctx1.Agent()))
		require.EqualValues(t, 0+42-42, ctx1.Balance(ctx2.Agent()))
		chainAccountBalances(ctx1, w, 2, 2+42-42)

		t.Logf("dump chain2 accounts:\n%s", ctx2.Chain.DumpAccounts())
		require.EqualValues(t, 0, ctx2.Balance(user))
		require.EqualValues(t, 0, ctx2.Balance(ctx1.Agent()))
		require.EqualValues(t, 1+42, ctx2.Balance(ctx2.Agent()))
		chainAccountBalances(ctx2, w, 2, 2+1+42)
	})
}
