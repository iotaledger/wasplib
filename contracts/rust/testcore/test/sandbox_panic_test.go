package test

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPanicFull(t *testing.T) { run2(t, testPanicFull) }
func testPanicFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(sbtestsc.Interface.Name, sbtestsc.FuncPanicFullEP)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgFullPanic))
}

func TestPanicViewCall(t *testing.T) { run2(t, testPanicViewCall) }
func testPanicViewCall(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(sbtestsc.Interface.Name, sbtestsc.FuncPanicViewEP)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))
}

func TestCallPanicFull(t *testing.T) { run2(t, testCallPanicFull) }
func testCallPanicFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(sbtestsc.Interface.Name, sbtestsc.FuncCallPanicFullEP)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgFullPanic))
}

func TestCallPanicViewFromFull(t *testing.T) { run2(t, testCallPanicViewFromFull) }
func testCallPanicViewFromFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(sbtestsc.Interface.Name, sbtestsc.FuncCallPanicViewEPFromFull)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))
}

func TestCallPanicViewFromView(t *testing.T) { run2(t, testCallPanicViewFromView) }
func testCallPanicViewFromView(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(sbtestsc.Interface.Name, sbtestsc.FuncCallPanicViewEPFromView)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))
}
