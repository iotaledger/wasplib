// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasmtest

package test

import (
	"github.com/iotaledger/wasp/contracts/rust/helloworld"
	"github.com/iotaledger/wasp/contracts/testenv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeployHelloWorld(t *testing.T) {
	te := testenv.NewTestEnv(t, helloworld.ScName)
	_, err := te.Chain.FindContract(helloworld.ScName)
	require.NoError(t, err)
}

func TestHelloWorld(t *testing.T) {
	te := testenv.NewTestEnv(t, helloworld.ScName)
	_ = te.NewCallParams(helloworld.FuncHelloWorld).Post(0)
}

func TestGetHelloWorld(t *testing.T) {
	te := testenv.NewTestEnv(t, helloworld.ScName)
	res := te.CallView(helloworld.ViewGetHelloWorld)
	result := te.Results(res)
	hw := result.GetString(helloworld.VarHelloWorld)
	require.EqualValues(t, "Hello, world!", hw.Value())
}
