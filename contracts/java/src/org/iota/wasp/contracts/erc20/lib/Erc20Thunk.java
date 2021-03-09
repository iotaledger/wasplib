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
        var p = ctx.Params();
        var params = new FuncApproveParams();
        params.Amount = p.GetInt64(Consts.ParamAmount);
        params.Delegation = p.GetAgentId(Consts.ParamDelegation);
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        ctx.Require(params.Delegation.Exists(), "missing mandatory delegation");
        ctx.Log("erc20.funcApprove");
        Erc20.funcApprove(ctx, params);
        ctx.Log("erc20.funcApprove ok");
    }

    private static void funcInitThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncInitParams();
        params.Creator = p.GetAgentId(Consts.ParamCreator);
        params.Supply = p.GetInt64(Consts.ParamSupply);
        ctx.Require(params.Creator.Exists(), "missing mandatory creator");
        ctx.Require(params.Supply.Exists(), "missing mandatory supply");
        ctx.Log("erc20.funcInit");
        Erc20.funcInit(ctx, params);
        ctx.Log("erc20.funcInit ok");
    }

    private static void funcTransferThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncTransferParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Amount = p.GetInt64(Consts.ParamAmount);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        ctx.Log("erc20.funcTransfer");
        Erc20.funcTransfer(ctx, params);
        ctx.Log("erc20.funcTransfer ok");
    }

    private static void funcTransferFromThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncTransferFromParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Amount = p.GetInt64(Consts.ParamAmount);
        params.Recipient = p.GetAgentId(Consts.ParamRecipient);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Amount.Exists(), "missing mandatory amount");
        ctx.Require(params.Recipient.Exists(), "missing mandatory recipient");
        ctx.Log("erc20.funcTransferFrom");
        Erc20.funcTransferFrom(ctx, params);
        ctx.Log("erc20.funcTransferFrom ok");
    }

    private static void viewAllowanceThunk(ScViewContext ctx) {
        var p = ctx.Params();
        var params = new ViewAllowanceParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        params.Delegation = p.GetAgentId(Consts.ParamDelegation);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Require(params.Delegation.Exists(), "missing mandatory delegation");
        ctx.Log("erc20.viewAllowance");
        Erc20.viewAllowance(ctx, params);
        ctx.Log("erc20.viewAllowance ok");
    }

    private static void viewBalanceOfThunk(ScViewContext ctx) {
        var p = ctx.Params();
        var params = new ViewBalanceOfParams();
        params.Account = p.GetAgentId(Consts.ParamAccount);
        ctx.Require(params.Account.Exists(), "missing mandatory account");
        ctx.Log("erc20.viewBalanceOf");
        Erc20.viewBalanceOf(ctx, params);
        ctx.Log("erc20.viewBalanceOf ok");
    }

    private static void viewTotalSupplyThunk(ScViewContext ctx) {
        var params = new ViewTotalSupplyParams();
        ctx.Log("erc20.viewTotalSupply");
        Erc20.viewTotalSupply(ctx, params);
        ctx.Log("erc20.viewTotalSupply ok");
    }
}