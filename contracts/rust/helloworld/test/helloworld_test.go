// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/helloworld"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, helloworld.ScName, helloworld.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	require.NoError(t, ctx.ContractExists(helloworld.ScName))
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
