package test

import (
	"fmt"
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/corecontracts/coreroot"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func run2(t *testing.T, test func(*testing.T, bool), skipWasm ...bool) {
	t.Run(fmt.Sprintf("run CORE version of %s", t.Name()), func(t *testing.T) {
		test(t, false)
	})
	if len(skipWasm) == 0 || !skipWasm[0] {
		t.Run(fmt.Sprintf("run WASM version of %s", t.Name()), func(t *testing.T) {
			test(t, true)
		})
	} else {
		t.Logf("skipped WASM version of '%s'", t.Name())
	}
}

func setupTest(t *testing.T, runWasm bool, addCreator ...bool) *wasmsolo.SoloContext {
	chain := wasmsolo.StartChain(t, "chain1")

	var creator *wasmsolo.SoloAgent
	if len(addCreator) != 0 && addCreator[0] {
		creator = wasmsolo.NewSoloAgent(chain.Env)

		ctxRoot := wasmsolo.NewSoloContextForCore(t, chain, coreroot.ScName, coreroot.OnLoad)
		grant := coreroot.ScFuncs.GrantDeployPermission(ctxRoot)
		grant.Params.Deployer().SetValue(creator.ScAgentID())
		grant.Func.TransferIotas(1).Post()
		require.NoError(t, ctxRoot.Err)
	}

	ctx := setupTestForChain(t, runWasm, chain, creator)
	require.NoError(t, ctx.Err)
	return ctx
}

func setupTestForChain(t *testing.T, runWasm bool, chain *solo.Chain, creator *wasmsolo.SoloAgent, init ...*wasmlib.ScInitFunc) *wasmsolo.SoloContext {
	if runWasm {
		return wasmsolo.NewSoloContextForChain(t, chain, creator, testcore.ScName, testcore.OnLoad, init...)
	}

	return wasmsolo.NewSoloContextForNative(t, chain, creator, testcore.ScName, testcore.OnLoad, sbtestsc.Processor, init...)
}

func TestSetup1(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := setupTest(t, w)
		require.EqualValues(t, ctx.Originator(), ctx.ContractCreator())
	})
}

func TestSetup2(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := setupTest(t, w, true)
		require.NotEqualValues(t, ctx.Originator(), ctx.ContractCreator())
	})
}
