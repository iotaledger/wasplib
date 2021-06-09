package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestPanicFull(t *testing.T) { run2(t, testPanicFull) }
func testPanicFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgFullPanic))

	recStr := chain.GetLogRecordsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, 4+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, 1, strings.Count(str, "panic in VM"))
}

func TestPanicViewCall(t *testing.T) { run2(t, testPanicViewCall) }
func testPanicViewCall(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(ScName, sbtestsc.FuncPanicViewEP)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))

	recStr := chain.GetLogRecordsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, 3+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, 0, strings.Count(str, "panic in VM"))
}

func TestCallPanicFull(t *testing.T) { run2(t, testCallPanicFull) }
func testCallPanicFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCallPanicFullEP).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgFullPanic))

	recStr := chain.GetLogRecordsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, 4+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, 1, strings.Count(str, "panic in VM"))
}

func TestCallPanicViewFromFull(t *testing.T) { run2(t, testCallPanicViewFromFull) }
func testCallPanicViewFromFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCallPanicViewEPFromFull).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))

	recStr := chain.GetLogRecordsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, 4+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, 1, strings.Count(str, "panic in VM"))
}

func TestCallPanicViewFromView(t *testing.T) { run2(t, testCallPanicViewFromView) }
func testCallPanicViewFromView(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(ScName, sbtestsc.FuncCallPanicViewEPFromView)
	require.Error(t, err)
	require.EqualValues(t, 1, strings.Count(err.Error(), sbtestsc.MsgViewPanic))

	recStr := chain.GetLogRecordsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, 3+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, 0, strings.Count(str, "panic in VM"))
}
