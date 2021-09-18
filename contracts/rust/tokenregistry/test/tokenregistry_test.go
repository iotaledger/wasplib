// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasplib/contracts/rust/tokenregistry"
	"github.com/iotaledger/wasplib/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContract(t, tokenregistry.ScName, tokenregistry.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	require.NoError(t, ctx.ContractExists(tokenregistry.ScName))
}
