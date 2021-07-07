// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/helloworld"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *common.SoloContext {
	chain := common.StartChainAndDeployWasmContractByName(t, helloworld.ScName)
	return common.NewSoloContext(helloworld.ScName, helloworld.OnLoad, chain)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	_, err := ctx.Chain.FindContract(helloworld.ScName)
	require.NoError(t, err)
}

func TestFuncHelloWorld(t *testing.T) {
	ctx := setupTest(t)

	helloWorld := helloworld.ScFuncs.HelloWorld(ctx)
	helloWorld.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
}

func TestViewGetHelloWorld(t *testing.T) {
	ctx := setupTest(t)

	getHelloWorld := helloworld.ScFuncs.GetHelloWorld(ctx)
	getHelloWorld.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "Hello, world!", getHelloWorld.Results.HelloWorld().Value())
}
