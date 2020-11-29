// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.mutable.ScMutableInt;

public class Increment {
	private static boolean localStateMustIncrement = false;

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("init", Increment::init);
		exports.AddCall("increment", Increment::increment);
		exports.AddCall("incrementCallIncrement", Increment::incrementCallIncrement);
		exports.AddCall("incrementCallIncrementRecurse5x", Increment::incrementCallIncrementRecurse5x);
		exports.AddCall("incrementPostIncrement", Increment::incrementPostIncrement);
		exports.AddView("incrementViewCounter", Increment::incrementViewCounter);
		exports.AddCall("incrementRepeatMany", Increment::incrementRepeatMany);
		exports.AddCall("incrementWhenMustIncrement", Increment::incrementWhenMustIncrement);
		exports.AddCall("incrementLocalStateInternalCall", Increment::incrementLocalStateInternalCall);
		exports.AddCall("incrementLocalStateSandboxCall", Increment::incrementLocalStateSandboxCall);
		exports.AddCall("incrementLocalStatePost", Increment::incrementLocalStatePost);
		exports.AddCall("nothing", ScExports::nothing);
		exports.AddCall("test", Increment::test);
	}

	public static void init(ScCallContext sc) {
		long counter = sc.Request().Params().GetInt("counter").Value();
		if (counter == 0) {
			return;
		}
		sc.State().GetInt("counter").SetValue(counter);
	}

	public static void increment(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt("counter");
		counter.SetValue(counter.Value() + 1);
	}

	public static void incrementCallIncrement(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			sc.CallSelf("incrementCallIncrement").Call();
		}
	}

	public static void incrementCallIncrementRecurse5x(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value < 5) {
			sc.CallSelf("incrementCallIncrementRecurse5x").Call();
		}
	}

	public static void incrementPostIncrement(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			sc.PostSelf("incrementPostIncrement").Post(0);
		}
	}

	public static void incrementViewCounter(ScViewContext sc) {
		long counter = sc.State().GetInt("counter").Value();
		sc.Results().GetInt("counter").SetValue(counter);
	}

	public static void incrementRepeatMany(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt("counter");
		long value = counter.Value();
		counter.SetValue(value + 1);
		ScMutableInt stateRepeats = sc.State().GetInt("numRepeats");
		long repeats = sc.Request().Params().GetInt("numRepeats").Value();
		if (repeats == 0) {
			repeats = stateRepeats.Value();
			if (repeats == 0) {
				return;
			}
		}
		stateRepeats.SetValue(repeats - 1);
		sc.PostSelf("incrementRepeatMany").Post(0);
	}

	public static void incrementWhenMustIncrement(ScCallContext sc) {
		sc.Log("incrementWhenMustIncrement called");
		if (localStateMustIncrement) {
			ScMutableInt counter = sc.State().GetInt("counter");
			counter.SetValue(counter.Value() + 1);
		}
	}

	public static void incrementLocalStateInternalCall(ScCallContext sc) {
		incrementWhenMustIncrement(sc);
		localStateMustIncrement = true;
		incrementWhenMustIncrement(sc);
		incrementWhenMustIncrement(sc);
		// counter ends up as 2
	}

	public static void incrementLocalStateSandboxCall(ScCallContext sc) {
		sc.CallSelf("incrementWhenMustIncrement").Call();
		localStateMustIncrement = true;
		sc.CallSelf("incrementWhenMustIncrement").Call();
		sc.CallSelf("incrementWhenMustIncrement").Call();
		// counter ends up as 0
	}

	public static void incrementLocalStatePost(ScCallContext sc) {
		sc.PostSelf("incrementWhenMustIncrement").Post(0);
		localStateMustIncrement = true;
		sc.PostSelf("incrementWhenMustIncrement").Post(0);
		sc.PostSelf("incrementWhenMustIncrement").Post(0);
		// counter ends up as 0
	}

	public static void test(ScCallContext sc) {
		int keyId = Host.GetKeyId("timestamp");
		Host.SetInt(1, keyId, 123456789);
		long timestamp = Host.GetInt(1, keyId);
		Host.SetInt(1, keyId, timestamp);

		int keyId2 = Host.GetKeyId("string");
		Host.SetString(1, keyId2, "Test");
		String s1 = Host.GetString(1, keyId2);
		Host.SetString(1, keyId2, "Bleep");
		String s2 = Host.GetString(1, keyId2);
		Host.SetString(1, keyId2, "Klunky");
		String s3 = Host.GetString(1, keyId2);
		Host.SetString(1, keyId2, s1);
		Host.SetString(1, keyId2, s2);
		Host.SetString(1, keyId2, s3);
	}
}
