package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestChainOwnerIDView(t *testing.T) { run2(t, testChainOwnerIDView) }
func testChainOwnerIDView(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestChainOwnerIDView(ctx)
	f.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, ctx.Originator().ScAgentID(), f.Results.ChainOwnerID().Value())
}

func TestChainOwnerIDFull(t *testing.T) { run2(t, testChainOwnerIDFull) }
func testChainOwnerIDFull(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestChainOwnerIDFull(ctx)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, ctx.Originator().ScAgentID(), f.Results.ChainOwnerID().Value())
}

func TestSandboxCall(t *testing.T) { run2(t, testSandboxCall) }
func testSandboxCall(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.TestSandboxCall(ctx)
	f.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "'solo' testing chain", f.Results.SandboxCall().Value())
}
