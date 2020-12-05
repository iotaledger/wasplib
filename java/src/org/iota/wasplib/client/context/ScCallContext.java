// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableString;
import org.iota.wasplib.client.request.ScCallInfo;
import org.iota.wasplib.client.request.ScPostInfo;
import org.iota.wasplib.client.request.ScViewInfo;

public class ScCallContext {
	ScMutableMap root;

	public ScCallContext() {
		root = new ScMutableMap(1);
	}

	static ScMutableMap makeRequest(String key, String function) {
		ScMutableMap root = new ScMutableMap(1);
		ScMutableMapArray requests = root.GetMapArray(key);
		ScMutableMap request = requests.GetMap(requests.Length());
		request.GetString("function").SetValue(function);
		return request;
	}

	public ScBalances Balances() {
		return new ScBalances(root.GetKeyMap("balances").Immutable());
	}

	public ScAgent Caller() {
		return root.GetAgent("caller").Value();
	}

	public ScCallInfo Call(String function) {
		return new ScCallInfo(makeRequest("calls", function));
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap("contract").Immutable());
	}

	public ScMutableString Error() {
		return root.GetString("error");
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public ScBalances Incoming() {
		return new ScBalances(root.GetKeyMap("incoming").Immutable());
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScImmutableMap Params() {
		return root.GetMap("params").Immutable();
	}

	public ScPostInfo PostGlobal(ScAddress chain, String contract, String function) {
		ScMutableMapArray posts = root.GetMapArray("posts");
		ScMutableMap post = posts.GetMap(posts.Length());
		post.GetAddress("chain").SetValue(chain);
		post.GetString("contract").SetValue(contract);
		post.GetString("function").SetValue(function);
		return new ScPostInfo(post);
	}

	public ScPostInfo Post(String function) {
		return new ScPostInfo(makeRequest("posts", function));
	}

	public ScMutableMap Results() {
		return root.GetMap("results");
	}

	public ScMutableMap State() {
		return root.GetMap("state");
	}

	public long Timestamp() {
		return root.GetInt("timestamp").Value();
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

	public ScViewInfo View(String function) {
		return new ScViewInfo(makeRequest("views", function));
	}
}
