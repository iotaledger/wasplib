package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) { run2(t, testGetSet) }
func testGetSet(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.SetInt(ctx)
	f.Params.Name().SetValue("ppp")
	f.Params.IntValue().SetValue(314)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	v := testcore.ScFuncs.GetInt(ctx)
	v.Params.Name().SetValue("ppp")
	v.Func.Call()
	require.NoError(t, ctx.Err)
	value := v.Results.Values().GetInt64("ppp")
	require.True(t, value.Exists())
	require.EqualValues(t, 314, value.Value())
}

func TestCallRecursive(t *testing.T) { run2(t, testCallRecursive) }
func testCallRecursive(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.CallOnChain(ctx)
	f.Params.IntValue().SetValue(31)
	f.Params.HnameContract().SetValue(testcore.HScName)
	f.Params.HnameEP().SetValue(testcore.HFuncRunRecursion)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	v := testcore.ScFuncs.GetCounter(ctx)
	v.Func.Call()
	require.NoError(t, ctx.Err)
	counter := v.Results.Counter()
	require.True(t, counter.Exists())
	require.EqualValues(t, 32, counter.Value())
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

	f := testcore.ScFuncs.Fibonacci(ctx)
	f.Params.IntValue().SetValue(n)
	f.Func.Call()
	require.NoError(t, ctx.Err)
	result := f.Results.IntValue()
	require.True(t, result.Exists())
	require.EqualValues(t, fibo(n), result.Value())
}

func TestCallFibonacciIndirect(t *testing.T) { run2(t, testCallFibonacciIndirect) }
func testCallFibonacciIndirect(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.CallOnChain(ctx)
	f.Params.IntValue().SetValue(n)
	f.Params.HnameContract().SetValue(testcore.HScName)
	f.Params.HnameEP().SetValue(testcore.HViewFibonacci)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
	result := f.Results.IntValue()
	require.True(t, result.Exists())
	require.EqualValues(t, fibo(n), result.Value())

	v := testcore.ScFuncs.GetCounter(ctx)
	v.Func.Call()
	require.NoError(t, ctx.Err)
	counter := v.Results.Counter()
	require.True(t, counter.Exists())
	require.EqualValues(t, 1, counter.Value())
}
