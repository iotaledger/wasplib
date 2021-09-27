package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestBasicBlockContext1(t *testing.T) {
	ctx := setupTest(t, false)

	f := testcore.ScFuncs.TestBlockContext1(ctx)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
}

func TestBasicBlockContext2(t *testing.T) {
	ctx := setupTest(t, false)

	f := testcore.ScFuncs.TestBlockContext2(ctx)
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)

	v := testcore.ScFuncs.GetStringValue(ctx)
	v.Params.VarName().SetValue("atTheEndKey")
	v.Func.Call()
	require.NoError(t, ctx.Err)
	value := v.Results.Vars().GetString("atTheEndKey")
	require.True(t, value.Exists())
	require.EqualValues(t, "atTheEndValue", value.Value())
}
