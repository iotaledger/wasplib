// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.inccounter;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.*;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.*;

public class IncCounter {
	private static final Key keyCounter = new Key("counter");
	private static final Key keyNumRepeats = new Key("num_repeats");

	static boolean localStateMustIncrement = false;

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("init", IncCounter::onInit);
		exports.AddCall("increment", IncCounter::increment);
		exports.AddCall("increment_call_increment", IncCounter::incrementCallIncrement);
		exports.AddCall("increment_call_increment_recurse5x", IncCounter::incrementCallIncrementRecurse5x);
		exports.AddCall("increment_post_increment", IncCounter::incrementPostIncrement);
		exports.AddView("increment_view_counter", IncCounter::incrementViewCounter);
		exports.AddCall("increment_repeat_many", IncCounter::incrementRepeatMany);
		exports.AddCall("increment_when_must_increment", IncCounter::incrementWhenMustIncrement);
		exports.AddCall("increment_local_state_internal_call", IncCounter::incrementLocalStateInternalCall);
		exports.AddCall("increment_local_state_sandbox_call", IncCounter::incrementLocalStateSandboxCall);
		exports.AddCall("increment_local_state_post", IncCounter::incrementLocalStatePost);
		exports.AddCall("nothing", ScExports::nothing);
		exports.AddCall("test", IncCounter::test);
		exports.AddCall("state_test", IncCounter::stateTest);
		exports.AddView("state_check", IncCounter::stateCheck);
		exports.AddCall("results_test", IncCounter::resultsTest);
		exports.AddView("results_check", IncCounter::resultsCheck);
	}

	public static void onInit(ScCallContext sc) {
		long counter = sc.Params().GetInt(keyCounter).Value();
		if (counter == 0) {
			return;
		}
		sc.State().GetInt(keyCounter).SetValue(counter);
	}

	public static void increment(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		counter.SetValue(counter.Value() + 1);
	}

	public static void incrementCallIncrement(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			sc.Call("increment_call_increment").Call();
		}
	}

	public static void incrementCallIncrementRecurse5x(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value < 5) {
			sc.Call("increment_call_increment_recurse5x").Call();
		}
	}

	public static void incrementPostIncrement(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		if (value == 0) {
			sc.Post("increment_post_increment").Post(0);
		}
	}

	public static void incrementViewCounter(ScViewContext sc) {
		long counter = sc.State().GetInt(keyCounter).Value();
		sc.Results().GetInt(keyCounter).SetValue(counter);
	}

	public static void incrementRepeatMany(ScCallContext sc) {
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		long value = counter.Value();
		counter.SetValue(value + 1);
		ScMutableInt stateRepeats = sc.State().GetInt(keyNumRepeats);
		long repeats = sc.Params().GetInt(keyNumRepeats).Value();
		if (repeats == 0) {
			repeats = stateRepeats.Value();
			if (repeats == 0) {
				return;
			}
		}
		stateRepeats.SetValue(repeats - 1);
		sc.Post("increment_repeat_many").Post(0);
	}

	public static void incrementWhenMustIncrement(ScCallContext sc) {
		sc.Log("increment_when_must_increment called");
		{
			if (!localStateMustIncrement) {
				return;
			}
		}
		ScMutableInt counter = sc.State().GetInt(keyCounter);
		counter.SetValue(counter.Value() + 1);
	}

	public static void incrementLocalStateInternalCall(ScCallContext sc) {
		incrementWhenMustIncrement(sc);
		{
			localStateMustIncrement = true;
		}
		incrementWhenMustIncrement(sc);
		incrementWhenMustIncrement(sc);
		// counter ends up as 2
	}

	public static void incrementLocalStateSandboxCall(ScCallContext sc) {
		sc.Call("increment_when_must_increment").Call();
		{
			localStateMustIncrement = true;
		}
		sc.Call("increment_when_must_increment").Call();
		sc.Call("increment_when_must_increment").Call();
		// counter ends up as 0
	}

	public static void incrementLocalStatePost(ScCallContext sc) {
		sc.Post("increment_when_must_increment").Post(0);
		{
			localStateMustIncrement = true;
		}
		sc.Post("increment_when_must_increment").Post(0);
		sc.Post("increment_when_must_increment").Post(0);
		// counter ends up as 0
	}

	public static void test(ScCallContext _sc) {
		int keyId = Host.GetKeyIdFromString("timestamp");
		Host.SetInt(1, keyId, 123456789);
		long timestamp = Host.GetInt(1, keyId);
		Host.SetInt(1, keyId, timestamp);
		int keyId2 = Host.GetKeyIdFromString("string");
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

	public static void resultsTest(ScCallContext sc) {
		testMap(sc.Results());
		checkMap(sc.Results().Immutable());
		//sc.call("results_check");
	}

	public static void stateTest(ScCallContext sc) {
		testMap(sc.State());
		sc.Call("state_check");
	}

	public static void resultsCheck(ScViewContext sc) {
		checkMap(sc.Results().Immutable());
	}

	public static void stateCheck(ScViewContext sc) {
		checkMap(sc.State());
	}

	public static void testMap(ScMutableMap kvstore) {
		ScMutableInt int1 = kvstore.GetInt(new Key("int1"));
		check(int1.Value() == 0);
		int1.SetValue(1);

		ScMutableString string1 = kvstore.GetString(new Key("string1"));
		check(string1.Value().equals(""));
		string1.SetValue("a");

		ScMutableIntArray ia1 = kvstore.GetIntArray(new Key("ia1"));
		ScMutableInt int2 = ia1.GetInt(0);
		check(int2.Value() == 0);
		int2.SetValue(2);
		ScMutableInt int3 = ia1.GetInt(1);
		check(int3.Value() == 0);
		int3.SetValue(3);

		ScMutableStringArray sa1 = kvstore.GetStringArray(new Key("sa1"));
		ScMutableString string2 = sa1.GetString(0);
		check(string2.Value().equals(""));
		string2.SetValue("bc");
		ScMutableString string3 = sa1.GetString(1);
		check(string3.Value().equals(""));
		string3.SetValue("def");
	}

	public static void checkMap(ScImmutableMap kvstore) {
		ScImmutableInt int1 = kvstore.GetInt(new Key("int1"));
		check(int1.Value() == 1);

		ScImmutableString string1 = kvstore.GetString(new Key("string1"));
		check(string1.Value().equals("a"));

		ScImmutableIntArray ia1 = kvstore.GetIntArray(new Key("ia1"));
		ScImmutableInt int2 = ia1.GetInt(0);
		check(int2.Value() == 2);
		ScImmutableInt int3 = ia1.GetInt(1);
		check(int3.Value() == 3);

		ScImmutableStringArray sa1 = kvstore.GetStringArray(new Key("sa1"));
		ScImmutableString string2 = sa1.GetString(0);
		check(string2.Value().equals("bc"));
		ScImmutableString string3 = sa1.GetString(1);
		check(string3.Value().equals("def"));
	}

	public static void checkMapRev(ScImmutableMap kvstore) {
		ScImmutableStringArray sa1 = kvstore.GetStringArray(new Key("sa1"));
		ScImmutableString string3 = sa1.GetString(1);
		check(string3.Value().equals("def"));
		ScImmutableString string2 = sa1.GetString(0);
		check(string2.Value().equals("bc"));

		ScImmutableIntArray ia1 = kvstore.GetIntArray(new Key("ia1"));
		ScImmutableInt int3 = ia1.GetInt(1);
		check(int3.Value() == 3);
		ScImmutableInt int2 = ia1.GetInt(0);
		check(int2.Value() == 2);

		ScImmutableString string1 = kvstore.GetString(new Key("string1"));
		check(string1.Value().equals("a"));

		ScImmutableInt int1 = kvstore.GetInt(new Key("int1"));
		check(int1.Value() == 1);
	}

	public static void check(boolean condition) {
		if (!condition) {
			Host.panic("Check failed!");
		}
	}
}
