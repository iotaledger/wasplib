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
	chain := common.StartChainAndDeployWasmContractByName(t, dividend.ScName)
	_, err := chain.FindContract(dividend.ScName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	chain := setupTest(t)
	_, member1Addr := chain.Env.NewKeyPairWithFunds()
	ctx := common.NewSoloContext(chain, member1Addr)

	f := dividend.NewMemberCall(ctx)
	f.Params.Address().SetValue(ctx.Address())
	f.Params.Factor().SetValue(100)
	f.Func.TransferIotas(1).Post()
	// req := solo.NewCallParams(ScName, FuncMember,
	// 	ParamAddress, member1Addr,
	// 	ParamFactor, 100,
	// )
	// req.WithIotas(1)
	// _, err := chain.PostRequestSync(req, nil)
	require.NoError(t, ctx.Err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	chain := setupTest(t)
	ctx := common.NewSoloContext(chain, nil)

	f := dividend.NewMemberCall(ctx)
	f.Params.Factor().SetValue(100)
	f.Func.TransferIotas(1).Post()
	// req := solo.NewCallParams(ScName, FuncMember,
	// 	ParamFactor, 100,
	// )
	// req.WithIotas(1)
	// _, err := chain.PostRequestSync(req, nil)
	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	chain := setupTest(t)
	_, member1Addr := chain.Env.NewKeyPairWithFunds()
	ctx := common.NewSoloContext(chain, member1Addr)

	f := dividend.NewMemberCall(ctx)
	f.Params.Address().SetValue(ctx.Address())
	f.Params.Factor().SetValue(100)
	f.Func.TransferIotas(1).Post()
	// req := solo.NewCallParams(ScName, FuncMember,
	// 	ParamAddress, member1Addr,
	// )
	// req.WithIotas(1)
	// _, err := chain.PostRequestSync(req, nil)
	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory factor"))
}
