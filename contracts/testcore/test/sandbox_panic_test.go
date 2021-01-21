package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPanicFull(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	req := solo.NewCall(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncPanicFullEP)
	_, err := chain.PostRequest(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), test_sandbox_sc.MsgFullPanic))
}

func TestPanicViewCall(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	_, err := chain.CallView(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncPanicViewEP)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), test_sandbox_sc.MsgViewPanic))
}

func TestCallPanicFull(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	req := solo.NewCall(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncCallPanicFullEP)
	_, err := chain.PostRequest(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), test_sandbox_sc.MsgFullPanic))
}

func TestCallPanicViewFromFull(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	req := solo.NewCall(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncCallPanicViewEPFromFull)
	_, err := chain.PostRequest(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), test_sandbox_sc.MsgViewPanic))
}

func TestCallPanicViewFromView(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	_, err := chain.CallView(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncCallPanicViewEPFromView)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), test_sandbox_sc.MsgViewPanic))
}
