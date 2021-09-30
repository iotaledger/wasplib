package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/stretchr/testify/require"
)

func TestTypesFull(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := setupTest(t, w)

		f := testcore.ScFuncs.PassTypesFull(ctx)
		f.Params.Address().SetValue(ctx.ChainID().Address())
		f.Params.AgentID().SetValue(ctx.Originator().ScAgentID())
		f.Params.ChainID().SetValue(ctx.ChainID())
		f.Params.ContractID().SetValue(ctx.AccountID())
		f.Params.Hash().SetValue(ctx.Convertor.ScHash(hashing.HashStrings("Hash")))
		f.Params.Hname().SetValue(wasmlib.NewScHname("Hname"))
		f.Params.HnameZero().SetValue(wasmlib.ScHname(0))
		f.Params.Int64().SetValue(42)
		f.Params.Int64Zero().SetValue(0)
		f.Params.String().SetValue("string")
		f.Params.StringZero().SetValue("")
		f.Func.TransferIotas(1).Post()
		require.NoError(t, ctx.Err)
	})
}

func TestTypesView(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := setupTest(t, w)

		v := testcore.ScFuncs.PassTypesView(ctx)
		v.Params.Address().SetValue(ctx.ChainID().Address())
		v.Params.AgentID().SetValue(ctx.Originator().ScAgentID())
		v.Params.ChainID().SetValue(ctx.ChainID())
		v.Params.ContractID().SetValue(ctx.AccountID())
		v.Params.Hash().SetValue(ctx.Convertor.ScHash(hashing.HashStrings("Hash")))
		v.Params.Hname().SetValue(wasmlib.NewScHname("Hname"))
		v.Params.HnameZero().SetValue(wasmlib.ScHname(0))
		v.Params.Int64().SetValue(42)
		v.Params.Int64Zero().SetValue(0)
		v.Params.String().SetValue("string")
		v.Params.StringZero().SetValue("")
		v.Func.Call()
		require.NoError(t, ctx.Err)
	})
}
