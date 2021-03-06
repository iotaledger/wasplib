// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.erc20.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.erc20.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class Erc20Thunk {
    public static void main(String[] args) {
        onLoad();
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc("approve", Erc20Thunk::funcApproveThunk);
        exports.AddFunc("init", Erc20Thunk::funcInitThunk);
        exports.AddFunc("transfer", Erc20Thunk::funcTransferThunk);
        exports.AddFunc("transferFrom", Erc20Thunk::funcTransferFromThunk);
        exports.AddView("allowance", Erc20Thunk::viewAllowanceThunk);
        exports.AddView("balanceOf", Erc20Thunk::viewBalanceOfThunk);
        exports.AddView("totalSupply", Erc20Thunk::viewTotalSupplyThunk);
    }

    private static void funcApproveThunk(ScFuncContext ctx) {
        ScImmutableMap p = ctx.Params();
        FuncApproveParams params = new FuncApproveParams();
        params.Amount = p.GetInt64(Consts.ParamAmount);
        params.Delegation = p.GetAgentId(Consts.ParamDelegation);
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        ctx.Require(params.Delegation.Exists(), "missing mandatory delegation");
        Erc20.funcApprove(ctx, params);
    }

    private static void funcInitThunk(ScFuncContext ctx) {
        ScImmutableMap p = ctx.Params();
        FuncInitParams params = new FuncInitParams();
        params.Creator = p.GetAgentId(Consts.ParamCreator);
        params.Supply = p.GetInt64(Consts.ParamSupply);
        ctx.Require(params.Creator.Exists(), "missing mandatory creator");
        ctx.Require(params.Supply.Exists(), "missing mandatory supply");
        Erc20.funcInit(ctx, params);
    }

    private static void funcTransferThunk(ScFuncContext ctx) {
        ScImmutableMap p = ctx.Params();
        FuncTransferParams params = new FuncTransferParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Amount = p.GetInt64(Consts.ParamAmount);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        Erc20.funcTransfer(ctx, params);
    }

    private static void funcTransferFromThunk(ScFuncContext ctx) {
        ScImmutableMap p = ctx.Params();
        FuncTransferFromParams params = new FuncTransferFromParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Amount = p.GetInt64(Consts.ParamAmount);
        params.Recipient = p.GetAgentId(Consts.ParamRecipient);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        ctx.Require(params.Recipient.Exists(), "missing mandatory recipient");
        Erc20.funcTransferFrom(ctx, params);
    }

    private static void viewAllowanceThunk(ScViewContext ctx) {
        ScImmutableMap p = ctx.Params();
        ViewAllowanceParams params = new ViewAllowanceParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Delegation = p.GetAgentId(Consts.ParamDelegation);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Delegation.Exists(), "missing mandatory delegation");
        Erc20.viewAllowance(ctx, params);
    }

    private static void viewBalanceOfThunk(ScViewContext ctx) {
        ScImmutableMap p = ctx.Params();
        ViewBalanceOfParams params = new ViewBalanceOfParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        Erc20.viewBalanceOf(ctx, params);
    }

    private static void viewTotalSupplyThunk(ScViewContext ctx) {
        ViewTotalSupplyParams params = new ViewTotalSupplyParams();
        Erc20.viewTotalSupply(ctx, params);
    }
}
