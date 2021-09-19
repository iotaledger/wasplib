package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) { run2(t, testGetSet) }
func testGetSet(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncSetInt.Name,
		sbtestsc.ParamIntParamName, "ppp",
		sbtestsc.ParamIntParamValue, 314,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetInt.Name,
		sbtestsc.ParamIntParamName, "ppp")
	require.NoError(t, err)

	retInt, exists, err := codec.DecodeInt64(ret.MustGet("ppp"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 314, retInt)
}

func TestCallRecursive(t *testing.T) { run2(t, testCallRecursive) }
func testCallRecursive(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCallOnChain.Name,
		sbtestsc.ParamIntParamValue, 31,
		sbtestsc.ParamHnameContract, cID.Hname(),
		sbtestsc.ParamHnameEP, sbtestsc.FuncRunRecursion.Hname(),
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)

	r, exists, err := codec.DecodeInt64(ret.MustGet(sbtestsc.VarCounter))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 32, r)
}

const n = 10

func fibo(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func TestCallFibonacci(t *testing.T) { run2(t, testCallFibonacci) }
func testCallFibonacci(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	fib := testcore.ScFuncs.Fibonacci(ctx)
	fib.Params.IntValue().SetValue(n)
	fib.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, fib.Results.IntValue().Exists())
	require.EqualValues(t, fibo(n), fib.Results.IntValue().Value())
}

func TestCallFibonacciIndirect(t *testing.T) { run2(t, testCallFibonacciIndirect) }
func testCallFibonacciIndirect(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	fib := testcore.ScFuncs.CallOnChain(ctx)
	fib.Params.IntValue().SetValue(n)
	fib.Params.HnameContract().SetValue(testcore.HScName)
	fib.Params.HnameEP().SetValue(testcore.HViewFibonacci)
	fib.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
	require.True(t, fib.Results.IntValue().Exists())
	require.EqualValues(t, fibo(n), fib.Results.IntValue().Value())

	gc := testcore.ScFuncs.GetCounter(ctx)
	gc.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, gc.Results.Counter().Exists())
	require.EqualValues(t, 1, gc.Results.Counter().Value())
}
