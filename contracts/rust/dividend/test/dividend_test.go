// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/dividend"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *common.SoloContext {
	chain := common.StartChainAndDeployWasmContractByName(t, dividend.ScName)
	return common.NewSoloContext(dividend.ScName, dividend.OnLoad, chain)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	_, err := ctx.Chain.FindContract(dividend.ScName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	ctx := setupTest(t)
	_, memberAddr := ctx.Chain.Env.NewKeyPair()

	member := dividend.ScFuncs.Member(ctx)
	member.Params.Address().SetValue(ctx.ScAddress(memberAddr))
	member.Params.Factor().SetValue(100)
	member.Func.TransferIotas(1).Post()

	require.NoError(t, ctx.Err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	ctx := setupTest(t)

	member := dividend.ScFuncs.Member(ctx)
	member.Params.Factor().SetValue(100)
	member.Func.TransferIotas(1).Post()

	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	ctx := setupTest(t)
	_, memberAddr := ctx.Chain.Env.NewKeyPair()

	member := dividend.ScFuncs.Member(ctx)
	member.Params.Address().SetValue(ctx.ScAddress(memberAddr))
	member.Func.TransferIotas(1).Post()

	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory factor"))
}
