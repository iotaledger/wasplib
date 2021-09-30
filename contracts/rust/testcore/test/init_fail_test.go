package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestInitSuccess(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := setupTest(t, w)
		require.NoError(t, ctx.Err)
	})
}

func TestInitFail(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		init := testcore.ScFuncs.Init(nil)
		init.Params.Fail().SetValue(1)
		ctx := setupTestForChain(t, w, nil, nil, init.Func)
		require.Error(t, ctx.Err)
	})
}

func TestInitFailRepeat(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		init := testcore.ScFuncs.Init(nil)
		init.Params.Fail().SetValue(1)
		ctx := setupTestForChain(t, w, nil, nil, init.Func)
		require.Error(t, ctx.Err)

		_, _, rec := ctx.Chain.GetInfo()
		require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))

		ctxRetry := setupTestForChain(t, w, ctx.Chain, nil)
		require.NoError(t, ctxRetry.Err)

		_, _, rec = ctxRetry.Chain.GetInfo()
		require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))
	})
}
