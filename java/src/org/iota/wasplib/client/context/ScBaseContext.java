// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.builders.ScViewBuilder;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScContractId;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScBaseContext {
	protected ScBaseContext() {
	}

	public ScBalances Balances() {
		return new ScBalances(Host.root.GetMap(Key.Balances).Immutable());
	}

	public ScAgent Caller() {
		return Host.root.GetAgent(Key.Caller).Value();
	}

	public ScAgent ChainOwner() {
		return Host.root.GetAgent(Key.ChainOwner).Value();
	}

	public ScAgent ContractCreator() {
		return Host.root.GetAgent(Key.Creator).Value();
	}

	public ScContractId ContractId() {
		return Host.root.GetContractId(Key.Id).Value();
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public void Log(String text) {
		Host.root.GetString(Key.Log).SetValue(text);
	}

	public void Panic(String text) {
		Host.root.GetString(Key.Panic).SetValue(text);
	}

	public ScImmutableMap Params() {
		return Host.root.GetMap(Key.Params).Immutable();
	}

	public ScMutableMap Results() {
		return Host.root.GetMap(Key.Results);
	}

	public long Timestamp() {
		return Host.root.GetInt(Key.Timestamp).Value();
	}

	public void Trace(String text) {
		Host.root.GetString(Key.Trace).SetValue(text);
	}

	public ScUtility Utility() {
		return new ScUtility(Host.root.GetMap(Key.Utility));
	}

	public ScViewBuilder View(String function) {
		return new ScViewBuilder(function);
	}
}
