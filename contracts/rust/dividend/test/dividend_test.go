// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/dividend"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func dividendMember(ctx *wasmsolo.SoloContext, agent *wasmsolo.SoloAgent, factor int64) {
	member := dividend.ScFuncs.Member(ctx)
	member.Params.Address().SetValue(agent.ScAddress())
	member.Params.Factor().SetValue(factor)
	member.Func.TransferIotas(1).Post()
}

func dividendDivide(ctx *wasmsolo.SoloContext, amount int64) {
	divide := dividend.ScFuncs.Divide(ctx)
	divide.Func.TransferIotas(amount).Post()
}

func dividendGetFactor(ctx *wasmsolo.SoloContext, member3 *wasmsolo.SoloAgent) int64 {
	getFactor := dividend.ScFuncs.GetFactor(ctx)
	getFactor.Params.Address().SetValue(member3.ScAddress())
	getFactor.Func.Call()
	value := getFactor.Results.Factor().Value()
	return value
}

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
	require.NoError(t, ctx.ContractExists(dividend.ScName))
}

func TestAddMemberOk(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	dividendMember(ctx, member1, 100)
	require.NoError(t, ctx.Err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member := dividend.ScFuncs.Member(ctx)
	member.Params.Factor().SetValue(100)
	member.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	member := dividend.ScFuncs.Member(ctx)
	member.Params.Address().SetValue(member1.ScAddress())
	member.Func.TransferIotas(1).Post()
	require.Error(t, ctx.Err)
	require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory factor"))
}

func TestDivide1Member(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	dividendMember(ctx, member1, 100)
	require.NoError(t, ctx.Err)

	require.EqualValues(t, 1, ctx.Balance(nil))

	dividendDivide(ctx, 99)
	require.NoError(t, ctx.Err)

	// 99 from divide() + 1 from the member() call
	require.EqualValues(t, solo.Saldo+100, member1.Balance())
	require.EqualValues(t, 0, ctx.Balance(nil))
}

func TestDivide2Members(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	dividendMember(ctx, member1, 25)
	require.NoError(t, ctx.Err)

	member2 := ctx.NewSoloAgent()
	dividendMember(ctx, member2, 75)
	require.NoError(t, ctx.Err)

	require.EqualValues(t, 2, ctx.Balance(nil))

	dividendDivide(ctx, 98)
	require.NoError(t, ctx.Err)

	// 98 from divide() + 2 from the member() calls
	require.EqualValues(t, solo.Saldo+25, member1.Balance())
	require.EqualValues(t, solo.Saldo+75, member2.Balance())
	require.EqualValues(t, 0, ctx.Balance(nil))
}

func TestDivide3Members(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	dividendMember(ctx, member1, 25)
	require.NoError(t, ctx.Err)

	member2 := ctx.NewSoloAgent()
	dividendMember(ctx, member2, 50)
	require.NoError(t, ctx.Err)

	member3 := ctx.NewSoloAgent()
	dividendMember(ctx, member3, 75)
	require.NoError(t, ctx.Err)

	require.EqualValues(t, 3, ctx.Balance(nil))

	dividendDivide(ctx, 97)
	require.NoError(t, ctx.Err)

	// 97 from divide() + 3 from the member() calls
	require.EqualValues(t, solo.Saldo+16, member1.Balance())
	require.EqualValues(t, solo.Saldo+33, member2.Balance())
	require.EqualValues(t, solo.Saldo+50, member3.Balance())
	// 1 remaining due to fractions
	require.EqualValues(t, 1, ctx.Balance(nil))

	dividendDivide(ctx, 100)
	require.NoError(t, ctx.Err)

	// 100 from divide() + 1 remaining
	require.EqualValues(t, solo.Saldo+16+16, member1.Balance())
	require.EqualValues(t, solo.Saldo+33+33, member2.Balance())
	require.EqualValues(t, solo.Saldo+50+50, member3.Balance())
	// now we have 2 remaining due to fractions
	require.EqualValues(t, 2, ctx.Balance(nil))

	dividendDivide(ctx, 100)
	require.NoError(t, ctx.Err)

	// 100 from divide() + 2 remaining
	require.EqualValues(t, solo.Saldo+16+16+17, member1.Balance())
	require.EqualValues(t, solo.Saldo+33+33+34, member2.Balance())
	require.EqualValues(t, solo.Saldo+50+50+51, member3.Balance())
	// managed to give every one an exact integer amount, so no remainder
	require.EqualValues(t, 0, ctx.Balance(nil))
}

func TestGetFactor(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)

	member1 := ctx.NewSoloAgent()
	dividendMember(ctx, member1, 25)
	require.NoError(t, ctx.Err)

	member2 := ctx.NewSoloAgent()
	dividendMember(ctx, member2, 50)
	require.NoError(t, ctx.Err)

	member3 := ctx.NewSoloAgent()
	dividendMember(ctx, member3, 75)
	require.NoError(t, ctx.Err)

	require.EqualValues(t, 3, ctx.Balance(nil))

	value := dividendGetFactor(ctx, member3)
	require.NoError(t, ctx.Err)
	require.EqualValues(t, 75, value)

	value = dividendGetFactor(ctx, member2)
	require.NoError(t, ctx.Err)
	require.EqualValues(t, 50, value)

	value = dividendGetFactor(ctx, member1)
	require.NoError(t, ctx.Err)
	require.EqualValues(t, 25, value)
}
