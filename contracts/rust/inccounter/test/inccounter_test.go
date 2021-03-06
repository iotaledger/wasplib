// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"
	"time"

	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/inccounter"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *common.SoloContext {
	chain := common.StartChainAndDeployWasmContractByName(t, inccounter.ScName)
	return common.NewSoloContext(inccounter.ScName, inccounter.OnLoad, chain)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	_, err := ctx.Chain.FindContract(inccounter.ScName)
	require.NoError(t, err)
}

func TestStateAfterDeploy(t *testing.T) {
	ctx := setupTest(t)

	checkStateCounter(t, ctx, nil)
}

func TestIncrementOnce(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.Increment(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 1)
}

func TestIncrementTwice(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.Increment(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	inccounter.ScFuncs.Increment(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 2)
}

func TestIncrementRepeatThrice(t *testing.T) {
	ctx := setupTest(t)

	repeatMany := inccounter.ScFuncs.RepeatMany(ctx)
	repeatMany.Params.NumRepeats().SetValue(3)
	repeatMany.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	require.True(t, ctx.WaitForRequestsThrough(7))

	checkStateCounter(t, ctx, 4)
}

func TestIncrementCallIncrement(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.CallIncrement(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 2)
}

func TestIncrementCallIncrementRecurse5x(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.CallIncrementRecurse5x(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 6)
}

func TestIncrementPostIncrement(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.PostIncrement(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	require.True(t, ctx.WaitForRequestsThrough(5))

	checkStateCounter(t, ctx, 2)
}

func TestIncrementLocalStateInternalCall(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.LocalStateInternalCall(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 2)
}

func TestIncrementLocalStateSandboxCall(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.LocalStateSandboxCall(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	if common.WasmRunner == 0 {
		// global var in wasm execution has no effect
		checkStateCounter(t, ctx, nil)
		return
	}

	checkStateCounter(t, ctx, 2)
}

func TestIncrementLocalStatePost(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.LocalStatePost(ctx).Func.TransferIotas(3).Post()
	require.NoError(t, ctx.Err)

	require.True(t, ctx.WaitForRequestsThrough(7))

	if common.WasmRunner == 0 {
		// global var in wasm execution has no effect
		checkStateCounter(t, ctx, nil)
		return
	}

	// when using WasmGoVM the 3 posts are run only after
	// the LocalStateMustIncrement has been set to true
	checkStateCounter(t, ctx, 3)
}

func TestLeb128(t *testing.T) {
	ctx := setupTest(t)

	inccounter.ScFuncs.TestLeb128(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	//res, err := chain.CallView(
	//	ScName, wasmproc.ViewCopyAllState,
	//)
	//require.NoError(t, err)
	//keys := make([]string, 0)
	//for key := range res {
	//	keys = append(keys, string(key))
	//}
	//sort.Strings(keys)
	//for _, key := range keys {
	//	fmt.Printf("%s: %v\n", key, res[kv.Key(key)])
	//}
}

func TestLoop(t *testing.T) {
	if common.WasmRunner != 0 {
		// no timeout possible with WasmGoVM
		// because goroutines cannot be killed
		t.SkipNow()
	}

	ctx := setupTest(t)

	wasmhost.WasmTimeout = 1 * time.Second
	inccounter.ScFuncs.EndlessLoop(ctx).Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	errText := ctx.Err.Error()
	require.True(t, strings.Contains(errText, "interrupt"))

	inccounter.ScFuncs.Increment(ctx).Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	checkStateCounter(t, ctx, 1)
}

func checkStateCounter(t *testing.T, ctx *common.SoloContext, expected interface{}) {
	getCounter := inccounter.ScFuncs.GetCounter(ctx)
	getCounter.Func.Call()
	require.NoError(t, ctx.Err)
	counter := getCounter.Results.Counter()
	if expected == nil {
		require.False(t, counter.Exists())
		return
	}
	require.True(t, counter.Exists())
	require.EqualValues(t, expected, counter.Value())
}
