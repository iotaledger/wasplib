// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.inccounter;

public class Inccounter {

static mut LocalStateMustIncrement: bool = false;

public static void funcCallIncrement(ScFuncContext ctx, FuncCallIncrementParams params) {
    counter = ctx.State().GetInt(VarCounter);
    value = counter.Value();
    counter.SetValue(value + 1);
    if (value == 0) {
        ctx.CallSelf(HFuncCallIncrement, null, null);
    }
}

public static void funcCallIncrementRecurse5x(ScFuncContext ctx, FuncCallIncrementRecurse5xParams params) {
    counter = ctx.State().GetInt(VarCounter);
    value = counter.Value();
    counter.SetValue(value + 1);
    if (value < 5) {
        ctx.CallSelf(HFuncCallIncrementRecurse5x, null, null);
    }
}

public static void funcIncrement(ScFuncContext ctx, FuncIncrementParams params) {
    counter = ctx.State().GetInt(VarCounter);
    counter.SetValue(counter.Value() + 1);
}

public static void funcInit(ScFuncContext ctx, FuncInitParams params) {
    if (params.Counter.Exists()) {
        counter = params.Counter.Value();
        ctx.State().GetInt(VarCounter).SetValue(counter);
    }
}

public static void funcLocalStateInternalCall(ScFuncContext ctx, FuncLocalStateInternalCallParams params) {
    unsafe {
        LocalStateMustIncrement = false;
    }
    par = FuncWhenMustIncrementParams{}
    funcWhenMustIncrement(ctx, par);
    unsafe {
        LocalStateMustIncrement = true;
    }
    funcWhenMustIncrement(ctx, par);
    funcWhenMustIncrement(ctx, par);
    // counter ends up as 2
}

public static void funcLocalStatePost(ScFuncContext ctx, FuncLocalStatePostParams params) {
    unsafe {
        LocalStateMustIncrement = false;
    }
    PostRequestParams request = new PostRequestParams();
         {
        request.ContractId = ctx.ContractId();
        request.Function = HFuncWhenMustIncrement;
        request.Params = null;
        request.Transfer = null;
        request.Delay = 0;
    }
    ctx.Post(request);
    unsafe {
        LocalStateMustIncrement = true;
    }
    ctx.Post(request);
    ctx.Post(request);
    // counter ends up as 0
}

public static void funcLocalStateSandboxCall(ScFuncContext ctx, FuncLocalStateSandboxCallParams params) {
    unsafe {
        LocalStateMustIncrement = false;
    }
    ctx.CallSelf(HFuncWhenMustIncrement, null, null);
    unsafe {
        LocalStateMustIncrement = true;
    }
    ctx.CallSelf(HFuncWhenMustIncrement, null, null);
    ctx.CallSelf(HFuncWhenMustIncrement, null, null);
    // counter ends up as 0
}

public static void funcPostIncrement(ScFuncContext ctx, FuncPostIncrementParams params) {
    counter = ctx.State().GetInt(VarCounter);
    value = counter.Value();
    counter.SetValue(value + 1);
    if (value == 0) {
        ctx.Post(PostRequestParams {
            request.ContractId = ctx.ContractId();
            request.Function = HFuncPostIncrement;
            request.Params = null;
            request.Transfer = null;
            request.Delay = 0;
        });
    }
}

public static void funcRepeatMany(ScFuncContext ctx, FuncRepeatManyParams params) {
    counter = ctx.State().GetInt(VarCounter);
    value = counter.Value();
    counter.SetValue(value + 1);
    stateRepeats = ctx.State().GetInt(VarNumRepeats);
    repeats = params.NumRepeats.Value();
    if (repeats == 0) {
        repeats = stateRepeats.Value();
        if (repeats == 0) {
            return;
        }
    }
    stateRepeats.SetValue(repeats - 1);
    ctx.Post(PostRequestParams {
        request.ContractId = ctx.ContractId();
        request.Function = HFuncRepeatMany;
        request.Params = null;
        request.Transfer = null;
        request.Delay = 0;
    });
}

public static void funcWhenMustIncrement(ScFuncContext ctx, FuncWhenMustIncrementParams params) {
    ctx.Log("when_must_increment called");
    unsafe {
        if (!LocalStateMustIncrement) {
            return;
        }
    }
    counter = ctx.State().GetInt(VarCounter);
    counter.SetValue(counter.Value() + 1);
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
public static void viewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
    counter = ctx.State().GetInt(VarCounter);
    if (counter.Exists()) {
        ctx.Results().GetInt(VarCounter).SetValue(counter.Value());
    }
}
}
