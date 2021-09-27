package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func verifyReceipts(t *testing.T, w bool, ctx *wasmsolo.SoloContext, requests, panics int) {
	recStr := ctx.Chain.GetRequestReceiptsForBlockRangeAsStrings(0, 0)
	str := strings.Join(recStr, "\n")
	t.Logf("\n%s", str)
	extra := 0
	if w {
		extra = 1
	}
	require.EqualValues(t, requests+extra, strings.Count(str, "req/tx"))
	require.EqualValues(t, panics, strings.Count(str, "panic in VM"))
}

func TestPanicFull(t *testing.T) { run2(t, testPanicFull) }
func testPanicFull(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestPanicFullEP(ctx)
	f.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	require.Contains(t, ctx.Err.Error(), sbtestsc.MsgFullPanic)

	verifyReceipts(t, w, ctx, 3, 1)
}

func TestPanicViewCall(t *testing.T) { run2(t, testPanicViewCall) }
func testPanicViewCall(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestPanicViewEP(ctx)
	f.Func.Call()
	require.Error(t, ctx.Err)
	require.Contains(t, ctx.Err.Error(), sbtestsc.MsgViewPanic)

	verifyReceipts(t, w, ctx, 2, 0)
}

func TestCallPanicFull(t *testing.T) { run2(t, testCallPanicFull) }
func testCallPanicFull(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestCallPanicFullEP(ctx)
	f.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	require.Contains(t, ctx.Err.Error(), sbtestsc.MsgFullPanic)

	verifyReceipts(t, w, ctx, 3, 1)
}

func TestCallPanicViewFromFull(t *testing.T) { run2(t, testCallPanicViewFromFull) }
func testCallPanicViewFromFull(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestCallPanicViewEPFromFull(ctx)
	f.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	require.Contains(t, ctx.Err.Error(), sbtestsc.MsgViewPanic)

	verifyReceipts(t, w, ctx, 3, 1)
}

func TestCallPanicViewFromView(t *testing.T) { run2(t, testCallPanicViewFromView) }
func testCallPanicViewFromView(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestCallPanicViewEPFromView(ctx)
	f.Func.Call()
	require.Error(t, ctx.Err)
	require.Contains(t, ctx.Err.Error(), sbtestsc.MsgViewPanic)

	verifyReceipts(t, w, ctx, 2, 0)
}
