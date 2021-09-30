package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestSpawn(t *testing.T) {
	ctx := setupTest(t, false)

	f := testcore.ScFuncs.Spawn(ctx)
	f.Params.ProgHash().SetValue(ctx.Convertor.ScHash(sbtestsc.Contract.ProgramHash))
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	ctxSpawn := wasmsolo.NewSoloContextForCore(t, ctx.Chain, testcore.ScName+"_spawned", testcore.OnLoad)
	require.NoError(t, ctxSpawn.Err)
	v := testcore.ScFuncs.GetCounter(ctxSpawn)
	v.Func.Call()
	require.EqualValues(t, 5, v.Results.Counter().Value())

	_, _, recs := ctx.Chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+2, len(recs))
}
