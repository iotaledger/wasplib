// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.inccounter;

import org.iota.wasp.contracts.inccounter.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class IncCounter {
	private static final Key KeyCounter = new Key("counter");
	private static final Key KeyNumRepeats = new Key("num_repeats");

	static boolean localStateMustIncrement = false;

	public static void FuncCallIncrement(ScFuncContext ctx, FuncCallIncrementParams params) {
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			ctx.CallSelf(new ScHname("increment_call_increment"), null, null);
		}
	}

	public static void FuncCallIncrementRecurse5x(ScFuncContext ctx, FuncCallIncrementRecurse5xParams params) {
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value < 5) {
			ctx.CallSelf(new ScHname("increment_call_increment_recurse5x"), null, null);
		}
	}

	public static void FuncIncrement(ScFuncContext ctx, FuncIncrementParams params) {
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		counter.SetValue(counter.Value() + 1);
	}

	public static void FuncInit(ScFuncContext ctx, FuncInitParams params) {
		long counter = ctx.Params().GetInt(KeyCounter).Value();
		if (counter == 0) {
			return;
		}
		ctx.State().GetInt(KeyCounter).SetValue(counter);
	}

	public static void FuncLocalStateInternalCall(ScFuncContext ctx, FuncLocalStateInternalCallParams params) {
		FuncWhenMustIncrementParams par = new FuncWhenMustIncrementParams();
		FuncWhenMustIncrement(ctx, par);
		{
			localStateMustIncrement = true;
		}
		FuncWhenMustIncrement(ctx, par);
		FuncWhenMustIncrement(ctx, par);
		// counter ends up as 2
	}

	public static void FuncLocalStatePost(ScFuncContext ctx, FuncLocalStatePostParams params) {
		PostRequestParams req = new PostRequestParams();
		req.ContractId = ctx.ContractId();
		req.Function = new ScHname("whenMustIncrement");
		ctx.Post(req);
		{
			localStateMustIncrement = true;
		}
		ctx.Post(req);
		ctx.Post(req);
		// counter ends up as 0
	}

	public static void FuncLocalStateSandboxCall(ScFuncContext ctx, FuncLocalStateSandboxCallParams params) {
		ctx.CallSelf(new ScHname("whenMustIncrement"), null, null);
		{
			localStateMustIncrement = true;
		}
		ctx.CallSelf(new ScHname("whenMustIncrement"), null, null);
		ctx.CallSelf(new ScHname("whenMustIncrement"), null, null);
		// counter ends up as 0
	}

	public static void FuncPostIncrement(ScFuncContext ctx, FuncPostIncrementParams params) {
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			PostRequestParams req = new PostRequestParams();
			req.ContractId = ctx.ContractId();
			req.Function = new ScHname("increment_post_increment");
			ctx.Post(req);
		}
	}

	public static void FuncRepeatMany(ScFuncContext ctx, FuncRepeatManyParams params) {
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		ScMutableInt stateRepeats = ctx.State().GetInt(KeyNumRepeats);
		long repeats = ctx.Params().GetInt(KeyNumRepeats).Value();
		if (repeats == 0) {
			repeats = stateRepeats.Value();
			if (repeats == 0) {
				return;
			}
		}
		stateRepeats.SetValue(repeats - 1);
		PostRequestParams req = new PostRequestParams();
		req.ContractId = ctx.ContractId();
		req.Function = new ScHname("increment_repeat_many");
		ctx.Post(req);
	}

	public static void FuncWhenMustIncrement(ScFuncContext ctx, FuncWhenMustIncrementParams params) {
		ctx.Log("increment_when_must_increment called");
		{
			if (!localStateMustIncrement) {
				return;
			}
		}
		ScMutableInt counter = ctx.State().GetInt(KeyCounter);
		counter.SetValue(counter.Value() + 1);
	}

	public static void ViewGetCounter(ScViewContext ctx, ViewGetCounterParams params) {
		long counter = ctx.State().GetInt(KeyCounter).Value();
		ctx.Results().GetInt(KeyCounter).SetValue(counter);
	}
}
