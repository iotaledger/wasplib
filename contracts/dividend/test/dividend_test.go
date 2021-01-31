package test

import (
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	scName       = "dividend"

	paramAddress = "address"
	paramFactor  = "factor"
)

var WasmFile = wasmhost.WasmPath("dividend_bg.wasm")

func TestDeploy(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	_, err := te.Chain.FindContract(scName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	user1 := te.Env.NewSignatureSchemeWithFunds()
	_ = te.NewCallParams("member",
		paramAddress, user1.Address(),
		paramFactor, 100,
	).Post(0)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	_ = te.NewCallParams("member",
		paramFactor, 100,
	).PostFail(0)
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	user1 := te.Env.NewSignatureSchemeWithFunds()
	_ = te.NewCallParams("member",
		paramAddress, user1.Address(),
	).PostFail(0)
}
