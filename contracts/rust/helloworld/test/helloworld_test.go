// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/helloworld"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, helloworld.ScName)
}

func TestDeploy(t *testing.T) {
	chain := setupTest(t)
	_, err := chain.FindContract(helloworld.ScName)
	require.NoError(t, err)
}

func TestFuncHelloWorld(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(helloworld.ScName, helloworld.OnLoad, chain, nil)

	helloWorld := helloworld.NewHelloWorldCall(ctx)
	helloWorld.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
}

func TestViewGetHelloWorld(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(helloworld.ScName, helloworld.OnLoad, chain, nil)

	getHelloWorld := helloworld.NewGetHelloWorldCall(ctx)
	getHelloWorld.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "Hello, world!", getHelloWorld.Results.HelloWorld().Value())
}
