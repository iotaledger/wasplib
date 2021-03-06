// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.inccounter;

import org.iota.wasp.contracts.inccounter.lib.*;
import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class IncCounter {

    static boolean LocalStateMustIncrement = false;

    public static void funcCallIncrement(ScFuncContext ctx, FuncCallIncrementParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        var value = counter.Value();
        counter.SetValue(value + 1);
        if (value == 0) {
            ctx.CallSelf(Consts.HFuncCallIncrement, null, null);
        }
    }

    public static void funcCallIncrementRecurse5x(ScFuncContext ctx, FuncCallIncrementRecurse5xParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        var value = counter.Value();
        counter.SetValue(value + 1);
        if (value < 5) {
            ctx.CallSelf(Consts.HFuncCallIncrementRecurse5x, null, null);
        }
    }

    public static void funcIncrement(ScFuncContext ctx, FuncIncrementParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        counter.SetValue(counter.Value() + 1);
    }

    public static void funcInit(ScFuncContext ctx, FuncInitParams params) {
        if (params.Counter.Exists()) {
            var counter = params.Counter.Value();
            ctx.State().GetInt64(Consts.VarCounter).SetValue(counter);
        }
    }

    public static void funcLocalStateInternalCall(ScFuncContext ctx, FuncLocalStateInternalCallParams params) {
        {
            LocalStateMustIncrement = false;
        }
        var par = new FuncWhenMustIncrementParams();
        {
        }
        funcWhenMustIncrement(ctx, par);
        {
            LocalStateMustIncrement = true;
        }
        funcWhenMustIncrement(ctx, par);
        funcWhenMustIncrement(ctx, par);
        // counter ends up as 2
    }

    public static void funcLocalStatePost(ScFuncContext ctx, FuncLocalStatePostParams params) {
        {
            LocalStateMustIncrement = false;
        }
        // prevent multiple identical posts, need a dummy param to differentiate them
        localStatePost(ctx, 1);
        {
            LocalStateMustIncrement = true;
        }
        localStatePost(ctx, 2);
        localStatePost(ctx, 3);
        // counter ends up as 0
    }

    private static void localStatePost(ScFuncContext ctx, long nr) {
        var params = new ScMutableMap();
        params.GetInt64(Consts.VarInt1).SetValue(nr);
        var transfer = ScTransfers.iotas(1);
        ctx.PostSelf(Consts.HFuncWhenMustIncrement, params, transfer, 0);
    }

    public static void funcLocalStateSandboxCall(ScFuncContext ctx, FuncLocalStateSandboxCallParams params) {
        {
            LocalStateMustIncrement = false;
        }
        ctx.CallSelf(Consts.HFuncWhenMustIncrement, null, null);
        {
            LocalStateMustIncrement = true;
        }
        ctx.CallSelf(Consts.HFuncWhenMustIncrement, null, null);
        ctx.CallSelf(Consts.HFuncWhenMustIncrement, null, null);
        // counter ends up as 0
    }

    public static void funcLoop(ScFuncContext ctx, FuncLoopParams params) {
        while (true) {
        }
    }

    public static void funcPostIncrement(ScFuncContext ctx, FuncPostIncrementParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        var value = counter.Value();
        counter.SetValue(value + 1);
        if (value == 0) {
            var transfer = ScTransfers.iotas(1);
            ctx.PostSelf(Consts.HFuncPostIncrement, null, transfer, 0);
        }
    }

    public static void funcRepeatMany(ScFuncContext ctx, FuncRepeatManyParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        var value = counter.Value();
        counter.SetValue(value + 1);
        var stateRepeats = ctx.State().GetInt64(Consts.VarNumRepeats);
        var repeats = params.NumRepeats.Value();
        if (repeats == 0) {
            repeats = stateRepeats.Value();
            if (repeats == 0) {
                return;
            }
        }
        stateRepeats.SetValue(repeats - 1);
        var transfer = ScTransfers.iotas(1);
        ctx.PostSelf(Consts.HFuncRepeatMany, null, transfer, 0);
    }

    public static void funcWhenMustIncrement(ScFuncContext ctx, FuncWhenMustIncrementParams params) {
        ctx.Log("when_must_increment called");
        {
            if (!LocalStateMustIncrement) {
                return;
            }
        }
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        counter.SetValue(counter.Value() + 1);
    }

    // note that get_counter mirrors the state of the 'counter' state variable
    // which means that if the state variable was not present it also will not be present in the result
    public static void viewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
        var counter = ctx.State().GetInt64(Consts.VarCounter);
        if (counter.Exists()) {
            ctx.Results().GetInt64(Consts.VarCounter).SetValue(counter.Value());
        }
    }

    public static void funcTestLeb128(ScFuncContext ctx, FuncTestLeb128Params params) {
        save(ctx, "v-1", -1);
        save(ctx, "v-2", -2);
        save(ctx, "v-126", -126);
        save(ctx, "v-127", -127);
        save(ctx, "v-128", -128);
        save(ctx, "v-129", -129);
        save(ctx, "v0", 0);
        save(ctx, "v+1", 1);
        save(ctx, "v+2", 2);
        save(ctx, "v+126", 126);
        save(ctx, "v+127", 127);
        save(ctx, "v+128", 128);
        save(ctx, "v+129", 129);
    }

    public static void save(ScFuncContext ctx, String name, long value) {
        var encode = new BytesEncoder();
        encode.Int64(value);
        var spot = ctx.State().GetBytes(new Key(name));
        spot.SetValue(encode.Data());

        var bytes = spot.Value();
        var decode = new BytesDecoder(bytes);
        var retrieved = decode.Int64();
        if (retrieved != value) {
            ctx.Log((name.toString() + " in : " + value));
            ctx.Log((name.toString() + " out: " + retrieved));
        }
    }
}
