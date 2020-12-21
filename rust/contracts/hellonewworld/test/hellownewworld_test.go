package test

import (
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

const WASM_FILE = "../pkg/hellonewworld_bg.wasm"

func TestDeploy(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	_, err = chain.FindContract("hello_new_world_1")
	require.NoError(t, err)
}

func TestHelloOnce(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	req := solo.NewCall("hello_new_world_1", "hello")
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	// call the contract to extract value of the 'counter'. Must be equal 1
	res, err := chain.CallView("hello_new_world_1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 1, counter)
}

func TestHello5Times(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		req := solo.NewCall("hello_new_world_1", "hello")
		_, err = chain.PostRequest(req, nil)
		require.NoError(t, err)
	}
	// call the contract to extract value of the 'counter'. Must be equal 5
	res, err := chain.CallView("hello_new_world_1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 5, counter)
}
