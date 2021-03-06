// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.dividend.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.dividend.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class DividendThunk {
    public static void main(String[] args) {
        onLoad();
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc("divide", DividendThunk::funcDivideThunk);
        exports.AddFunc("member", DividendThunk::funcMemberThunk);
        exports.AddView("getFactor", DividendThunk::viewGetFactorThunk);
    }

    private static void funcDivideThunk(ScFuncContext ctx) {
        FuncDivideParams params = new FuncDivideParams();
        Dividend.funcDivide(ctx, params);
    }

    private static void funcMemberThunk(ScFuncContext ctx) {
        // only creator can add members
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        ScImmutableMap p = ctx.Params();
        FuncMemberParams params = new FuncMemberParams();
        params.Address = p.GetAddress(Consts.ParamAddress);
        params.Factor = p.GetInt64(Consts.ParamFactor);
        ctx.Require(params.Address.Exists(), "missing mandatory address");
        ctx.Require(params.Factor.Exists(), "missing mandatory factor");
        Dividend.funcMember(ctx, params);
    }

    private static void viewGetFactorThunk(ScViewContext ctx) {
        ScImmutableMap p = ctx.Params();
        ViewGetFactorParams params = new ViewGetFactorParams();
        params.Address = p.GetAddress(Consts.ParamAddress);
        ctx.Require(params.Address.Exists(), "missing mandatory address");
        Dividend.viewGetFactor(ctx, params);
    }
}
