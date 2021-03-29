// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.testwasmlib.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.testwasmlib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class TestWasmLibThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc(Consts.FuncParamTypes, TestWasmLibThunk::funcParamTypesThunk);
    }

    private static void funcParamTypesThunk(ScFuncContext ctx) {
        ctx.Log("testwasmlib.funcParamTypes");
        var p = ctx.Params();
        var params = new FuncParamTypesParams();
        params.Address = p.GetAddress(Consts.ParamAddress);
        params.AgentId = p.GetAgentId(Consts.ParamAgentId);
        params.Bytes = p.GetBytes(Consts.ParamBytes);
        params.ChainId = p.GetChainId(Consts.ParamChainId);
        params.Color = p.GetColor(Consts.ParamColor);
        params.Hash = p.GetHash(Consts.ParamHash);
        params.Hname = p.GetHname(Consts.ParamHname);
        params.Int64 = p.GetInt64(Consts.ParamInt64);
        params.RequestId = p.GetRequestId(Consts.ParamRequestId);
        params.String = p.GetString(Consts.ParamString);
        TestWasmLib.funcParamTypes(ctx, params);
        ctx.Log("testwasmlib.funcParamTypes ok");
    }
}