package org.iota.wasplib;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScExports;
import org.iota.wasplib.client.mutable.ScMutableInt;

public class Increment {
	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.Add("increment");
		exports.Add("incrementRepeat1");
		exports.Add("incrementRepeatMany");
		exports.Add("test");
		exports.Add("nothing");
	}

	//export test
	public static void test() {
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
