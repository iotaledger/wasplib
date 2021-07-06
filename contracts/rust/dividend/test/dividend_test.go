// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/dividend"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, dividend.ScName)
}

func TestDeploy(t *testing.T) {
	chain := setupTest(t)
	_, err := chain.FindContract(dividend.ScName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(dividend.ScName, dividend.OnLoad, chain, nil)
	_, memberAddr := chain.Env.NewKeyPair()

	newMember := dividend.NewMemberCall(ctx)
	newMember.Params.Address().SetValue(ctx.ScAddress(memberAddr))
	newMember.Params.Factor().SetValue(100)
	newMember.Func.TransferIotas(1).Post()

	require.NoError(t, ctx.Err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(dividend.ScName, dividend.OnLoad, chain, nil)

	newMember := dividend.NewMemberCall(ctx)
	newMember.Params.Factor().SetValue(100)
	newMember.Func.TransferIotas(1).Post()

	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(dividend.ScName, dividend.OnLoad, chain, nil)
	_, memberAddr := chain.Env.NewKeyPair()

	newMember := dividend.NewMemberCall(ctx)
	newMember.Params.Address().SetValue(ctx.ScAddress(memberAddr))
	newMember.Func.TransferIotas(1).Post()

	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory factor"))
}
