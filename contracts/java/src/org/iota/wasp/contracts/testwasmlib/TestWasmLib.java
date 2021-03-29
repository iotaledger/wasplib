// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.testwasmlib;

import org.iota.wasp.contracts.testwasmlib.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;

import java.nio.charset.*;
import java.sql.*;
import java.util.*;

public class TestWasmLib {

    public static void funcParamTypes(ScFuncContext ctx, FuncParamTypesParams params) {
        if (params.Address.Exists()) {
            ctx.Require(params.Address.Value().equals(ctx.AccountId().Address()), "mismatch: Address");
        }
        if (params.AgentId.Exists()) {
            ctx.Require(params.AgentId.Value().equals(ctx.AccountId()), "mismatch: AgentId");
        }
        if (params.Bytes.Exists()) {
            var bytes = "these are bytes".getBytes(StandardCharsets.UTF_8);
            ctx.Require(Arrays.equals(params.Bytes.Value(), bytes), "mismatch: Bytes");
        }
        if (params.ChainId.Exists()) {
            ctx.Require(params.ChainId.Value().equals(ctx.ChainId()), "mismatch: ChainId");
        }
        if (params.Color.Exists()) {
            var color = new ScColor("RedGreenBlueYellowCyanBlackWhite".getBytes(StandardCharsets.UTF_8));
            ctx.Require(params.Color.Value().equals(color), "mismatch: Color");
        }
        if (params.Hash.Exists()) {
            var hash = new ScHash("0123456789abcdeffedcba9876543210".getBytes(StandardCharsets.UTF_8));
            ctx.Require(params.Hash.Value().equals(hash), "mismatch: Hash");
        }
        if (params.Hname.Exists()) {
            ctx.Require(params.Hname.Value().equals(ctx.AccountId().Hname()), "mismatch: Hname");
        }
        if (params.Int64.Exists()) {
            ctx.Require(params.Int64.Value() == 1234567890123456789L, "mismatch: Int64");
        }
        if (params.RequestId.Exists()) {
            var requestId = new ScRequestId("abcdefghijklmnopqrstuvwxyz12345678".getBytes(StandardCharsets.UTF_8));
            ctx.Require(params.RequestId.Value().equals(requestId), "mismatch: RequestId");
        }
        if (params.String.Exists()) {
            ctx.Require(params.String.Value().equals("this is a string"), "mismatch: String");
        }
    }
}
