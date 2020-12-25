package test

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	WasmFile = "../pkg/dividend_bg.wasm"

	ParamAddress = "address"
	ParamFactor  = "factor"
)

func TestDeploy(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "dividend", WasmFile)
	require.NoError(t, err)

	_, err = chain.FindContract("dividend")
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "dividend", WasmFile)
	require.NoError(t, err)

	user1 := glb.NewSignatureSchemeWithFunds()
	user1address := user1.Address()
	req := solo.NewCall("dividend", "member",
		ParamAddress, user1address,
		ParamFactor, 100,
	)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}

func TestAddMembeParamFail1(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "dividend", WasmFile)
	require.NoError(t, err)

	req := solo.NewCall("dividend", "member",
		ParamFactor, 100,
	)
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)
}

func TestAddMemberParamFail2(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "dividend", WasmFile)
	require.NoError(t, err)

	user1 := glb.NewSignatureSchemeWithFunds()
	user1address := user1.Address()
	req := solo.NewCall("dividend", "member",
		ParamAddress, user1address,
	)
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)
}
