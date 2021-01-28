package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSet(t *testing.T) { run2(t, testGetSet) }
func testGetSet(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(SandboxSCName, test_sandbox_sc.FuncSetInt,
		test_sandbox_sc.ParamIntParamName, "ppp",
		test_sandbox_sc.ParamIntParamValue, 314,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	ret, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncGetInt,
		test_sandbox_sc.ParamIntParamName, "ppp")
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

	req := solo.NewCallParams(SandboxSCName, test_sandbox_sc.FuncCallOnChain,
		test_sandbox_sc.ParamIntParamValue, 31,
		test_sandbox_sc.ParamHnameContract, cID.Hname(),
		test_sandbox_sc.ParamHnameEP, coretypes.Hn(test_sandbox_sc.FuncRunRecursion),
	)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	ret, err = chain.CallView(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncGetCounter)
	require.NoError(t, err)

	r, exists, err := codec.DecodeInt64(ret.MustGet(test_sandbox_sc.VarCounter))
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
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	ret, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncGetFibonacci,
		test_sandbox_sc.ParamIntParamValue, n,
	)
	require.NoError(t, err)
	val, exists, err := codec.DecodeInt64(ret.MustGet(test_sandbox_sc.ParamIntParamValue))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, fibo(n), val)
}

func TestCallFibonacciIndirect(t *testing.T) { run2(t, testCallFibonacciIndirect) }
func testCallFibonacciIndirect(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(SandboxSCName, test_sandbox_sc.FuncCallOnChain,
		test_sandbox_sc.ParamIntParamValue, n,
		test_sandbox_sc.ParamHnameContract, cID.Hname(),
		test_sandbox_sc.ParamHnameEP, coretypes.Hn(test_sandbox_sc.FuncGetFibonacci),
	)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	r, exists, err := codec.DecodeInt64(ret.MustGet(test_sandbox_sc.ParamIntParamValue))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, fibo(n), r)

	ret, err = chain.CallView(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncGetCounter)
	require.NoError(t, err)

	r, exists, err = codec.DecodeInt64(ret.MustGet(test_sandbox_sc.VarCounter))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 1, r)
}
