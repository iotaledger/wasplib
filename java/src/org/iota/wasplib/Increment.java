package org.iota.wasplib;

import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.mutable.ScMutableInt;

public class Increment {
	//export increment
	public static void increment() {
		ScContext sc = new ScContext();
		ScMutableInt counter = sc.State().GetInt("counter");
		counter.SetValue(counter.Value() + 1);
	}

	//export incrementRepeat1
	public static void incrementRepeat1() {
		ScContext sc = new ScContext();
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			sc.Event("", "increment", 5);
		}
	}

	//export incrementRepeatMany
	public static void incrementRepeatMany() {
		ScContext sc = new ScContext();
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		long repeats = sc.Request().Params().GetInt("numRepeats").Value();
		ScMutableInt stateRepeats = sc.State().GetInt("numRepeats");
		if (repeats == 0) {
			repeats = stateRepeats.Value();
			if (repeats == 0) {
				return;
			}
		}
		stateRepeats.SetValue(repeats - 1);
		sc.Event("", "incrementRepeatMany", 3);
	}
}
