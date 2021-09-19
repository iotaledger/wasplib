package test

import (
	"testing"

	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) { run2(t, testGetSet) }
func testGetSet(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	si := testcore.ScFuncs.SetInt(ctx)
	si.Params.Name().SetValue("ppp")
	si.Params.IntValue().SetValue(314)
	si.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	gi := testcore.ScFuncs.GetInt(ctx)
	gi.Params.Name().SetValue("ppp")
	gi.Func.Call()
	require.NoError(t, ctx.Err)
	ppp := gi.Results.Values().GetInt64("ppp")
	require.True(t, ppp.Exists())
	require.EqualValues(t, 314, ppp.Value())
}

func TestCallRecursive(t *testing.T) { run2(t, testCallRecursive) }
func testCallRecursive(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	coc := testcore.ScFuncs.CallOnChain(ctx)
	coc.Params.IntValue().SetValue(31)
	coc.Params.HnameContract().SetValue(testcore.HScName)
	coc.Params.HnameEP().SetValue(testcore.HFuncRunRecursion)
	coc.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	gc := testcore.ScFuncs.GetCounter(ctx)
	gc.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, gc.Results.Counter().Exists())
	require.EqualValues(t, 32, gc.Results.Counter().Value())
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
