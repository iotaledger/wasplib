package org.iota.wasplib;

import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.mutable.ScMutableInt;

public class Increment {
	//export no_op
	public static void no_op() {
		ScContext ctx = new ScContext();
		ctx.Log("Doing nothing as requested. Oh, wait...");
	}

	//export increment
	public static void increment() {
		ScContext ctx = new ScContext();
		ctx.Log("Increment...");
		ScMutableInt counter = ctx.State().GetInt("counter");
		counter.SetValue(counter.Value() + 1);
	}

	//export incrementRepeat1
	public static void incrementRepeat1() {
		ScContext ctx = new ScContext();
		ctx.Log("incrementRepeat1...");
		ScMutableInt counter = ctx.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			ctx.Event("", "increment", 5);
		}
	}

	//export incrementRepeatMany
	public static void incrementRepeatMany() {
		ScContext ctx = new ScContext();
		ctx.Log("incrementRepeatMany...");
		ScMutableInt counter = ctx.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		long repeats = ctx.Request().Params().GetInt("numrepeats").Value();
		ScMutableInt stateRepeats = ctx.State().GetInt("numrepeats");
		if (repeats == 0) {
			repeats = stateRepeats.Value();
			if (repeats == 0) {
				return;
			}
		}
		stateRepeats.SetValue(repeats - 1);
		ctx.Event("", "incrementRepeatMany", 3);
	}
}
