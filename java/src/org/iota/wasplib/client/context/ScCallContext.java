// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Key;
import org.iota.wasplib.client.KeyId;
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
	private static final ScMutableMap root = new ScMutableMap(1);

	public ScCallContext() {
	}

	static ScMutableMap makeRequest(KeyId key, String function) {
		ScMutableMap root = new ScMutableMap(1);
		ScMutableMapArray requests = root.GetMapArray(key);
		ScMutableMap request = requests.GetMap(requests.Length());
		request.GetString(new Key("function")).SetValue(function);
		return request;
	}

	public ScBalances Balances() {
		return new ScBalances(root.GetMap(new Key("balances")).Immutable());
	}

	public ScAgent Caller() {
		return root.GetAgent(new Key("caller")).Value();
	}

	public ScCallInfo Call(String function) {
		return new ScCallInfo(makeRequest(new Key("calls"), function));
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap(new Key("contract")).Immutable());
	}

	public ScMutableString Error() {
		return root.GetString(new Key("error"));
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public ScBalances Incoming() {
		return new ScBalances(root.GetMap(new Key("incoming")).Immutable());
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScImmutableMap Params() {
		return root.GetMap(new Key("params")).Immutable();
	}

	public ScPostInfo PostGlobal(ScAddress chain, String contract, String function) {
		ScMutableMapArray posts = root.GetMapArray(new Key("posts"));
		ScMutableMap post = posts.GetMap(posts.Length());
		post.GetAddress(new Key("chain")).SetValue(chain);
		post.GetString(new Key("contract")).SetValue(contract);
		post.GetString(new Key("function")).SetValue(function);
		return new ScPostInfo(post);
	}

	public ScPostInfo Post(String function) {
		return new ScPostInfo(makeRequest(new Key("posts"), function));
	}

	public ScMutableMap Results() {
		return root.GetMap(new Key("results"));
	}

	public ScMutableMap State() {
		return root.GetMap(new Key("state"));
	}

	public long Timestamp() {
		return root.GetInt(new Key("timestamp")).Value();
	}

	public ScLog TimestampedLog(KeyId key) {
		return new ScLog(root.GetMap(new Key("logs")).GetMapArray(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public void Transfer(ScAgent agent, ScColor color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray(new Key("transfers"));
		ScMutableMap transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgent(new Key("agent")).SetValue(agent);
		transfer.GetColor(new Key("color")).SetValue(color);
		transfer.GetInt(new Key("amount")).SetValue(amount);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap(new Key("utility")));
	}

	public ScViewInfo View(String function) {
		return new ScViewInfo(makeRequest(new Key("views"), function));
	}
}
