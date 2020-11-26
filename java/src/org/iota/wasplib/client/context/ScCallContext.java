// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScCallContext {
	ScMutableMap root;

	public ScCallContext() {
		root = new ScMutableMap(1);
	}

	public ScAccount Account() {
		return new ScAccount(root.GetMap("account").Immutable());
	}

	public ScCallInfo Call(String contract, String function) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScMutableMap call = calls.GetMap(calls.Length());
		call.GetString("contract").SetValue(contract);
		call.GetString("function").SetValue(function);
		return new ScCallInfo(call);
	}

	public ScCallInfo CallSelf(String function) {
		return Call("", function);
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

	public ScPostInfo PostGlobal(ScAddress chain, String contract, String function) {
		ScMutableMapArray posts = root.GetMapArray("posts");
		ScMutableMap post = posts.GetMap(posts.Length());
		post.GetAddress("chain").SetValue(chain);
		post.GetString("contract").SetValue(contract);
		post.GetString("function").SetValue(function);
		return new ScPostInfo(post);
	}

	public ScPostInfo PostLocal(String contract, String function) {
		ScMutableMapArray posts = root.GetMapArray("posts");
		ScMutableMap post = posts.GetMap(posts.Length());
		post.GetString("contract").SetValue(contract);
		post.GetString("function").SetValue(function);
		return new ScPostInfo(post);
	}

	public ScPostInfo PostSelf(String function) {
		return PostLocal("", function);
	}

	public ScRequest Request() {
		return new ScRequest(root.GetMap("request").Immutable());
	}

	public ScMutableMap Results() {
		return root.GetMap("results");
	}

	public ScMutableMap State() {
		return root.GetMap("state");
	}

	public ScLog TimestampedLog(String key) {
		return new ScLog(root.GetMap("logs").GetMapArray(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public void Transfer(ScAgent agent, ScColor color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray("transfers");
		ScMutableMap transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgent("agent").SetValue(agent);
		transfer.GetColor("color").SetValue(color);
		transfer.GetInt("amount").SetValue(amount);
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
