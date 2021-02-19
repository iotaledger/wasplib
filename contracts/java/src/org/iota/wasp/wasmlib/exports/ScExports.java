// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.exports;

import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.util.*;

public class ScExports {
	private static final ArrayList<ScFunc> funcs = new ArrayList<>();
	private static final ArrayList<ScView> views = new ArrayList<>();

	ScMutableStringArray exports;

	public ScExports() {
		exports = Host.root.GetStringArray(Key.Exports);
		// tell host what our highest predefined key is
		// this helps detect missing or extra keys
		exports.GetString(Key.Zzzzzzz.KeyId()).SetValue("Java:KEY_ZZZZZZZ");
	}

	//export on_call_entrypoint
	static void scCallEntrypoint(int index) {
		if ((index & 0x8000) != 0) {
			views.get(index & 0x7fff).call(new ScViewContext());
			return;
		}
		funcs.get(index).call(new ScFuncContext());
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
