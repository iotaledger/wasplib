// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.exports;

import org.iota.wasplib.client.context.ScFuncContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableStringArray;

import java.util.ArrayList;

public class ScExports {
	private static final ArrayList<ScFunc> funcs = new ArrayList<>();
	private static final ArrayList<ScView> views = new ArrayList<>();

	ScMutableStringArray exports;

	public ScExports() {
		exports = Host.root.GetStringArray(Key.Exports);
		// tell host what our highest predefined key is
		// this helps detect missing or extra keys
		exports.GetString(Key.Zzzzzzz.GetId()).SetValue("Java:KEY_ZZZZZZZ");
	}

	//export on_call_entrypoint
	static void scCallEntrypoint(int index) {
		if ((index & 0x8000) != 0) {
			views.get(index & 0x7fff).call(new ScViewContext());
			return;
		}
		funcs.get(index).call(new ScFuncContext());
	}

	public static void nothing(ScFuncContext sc) {
		sc.Log("Doing nothing as requested. Oh, wait...");
	}

	public void AddFunc(String name, ScFunc f) {
		int index = funcs.size();
		funcs.add(f);
		exports.GetString(index).SetValue(name);
	}

	public void AddView(String name, ScView f) {
		int index = views.size();
		views.add(f);
		exports.GetString(index | 0x8000).SetValue(name);
	}
}
