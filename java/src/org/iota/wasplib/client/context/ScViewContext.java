// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScViewContext {
	ScMutableMap root;

	public ScViewContext() {
		root = new ScMutableMap(1);
	}

	public ScAccount Account() {
		return new ScAccount(root.GetMap("account").Immutable());
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap("contract").Immutable());
	}

	public ScMutableString Error() {
		return root.GetString("error");
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScRequest Request() {
		return new ScRequest(root.GetMap("request").Immutable());
	}

	public ScMutableMap Results() {
		return root.GetMap("results");
	}

	public ScImmutableMap State() {
		return root.GetMap("state").Immutable();
	}

	public ScLog TimestampedLog(String key) {
		return new ScLog(root.GetMap("logs").GetMapArray(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap("utility"));
	}

	public ScViewInfo View(String contract, String function) {
		ScMutableMapArray views = root.GetMapArray("views");
		ScMutableMap view = views.GetMap(views.Length());
		view.GetString("contract").SetValue(contract);
		view.GetString("function").SetValue(function);
		return new ScViewInfo(view);
	}

	public ScViewInfo ViewSelf(String function) {
		return View("", function);
	}
}
