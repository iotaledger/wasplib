package test

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	ParamAddress = "address"
	ParamFactor  = "factor"
	ScName = "dividend"
)

var WasmFile = wasmhost.WasmPath("dividend_bg.wasm")

func TestDeploy(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, ScName, WasmFile)
	require.NoError(t, err)

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, ScName, WasmFile)
	require.NoError(t, err)

	user1 := glb.NewSignatureSchemeWithFunds()
	user1address := user1.Address()
	req := solo.NewCallParams(ScName, "member",
		ParamAddress, user1address,
		ParamFactor, 100,
	)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}

func TestAddMemberParamFail1(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, ScName, WasmFile)
	require.NoError(t, err)

	req := solo.NewCallParams(ScName, "member",
		ParamFactor, 100,
	)
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)
}

func TestAddMemberParamFail2(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, ScName, WasmFile)
	require.NoError(t, err)

	user1 := glb.NewSignatureSchemeWithFunds()
	user1address := user1.Address()
	req := solo.NewCallParams(ScName, "member",
		ParamAddress, user1address,
	)
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)
}
