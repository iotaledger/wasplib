package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScContract {
	ScImmutableMap contract;

	ScContract(ScImmutableMap contract) {
		this.contract = contract;
	}

	public ScAddress Address() {
		return contract.GetAddress("address").Value();
	}

	public ScColor Color() {
		return contract.GetColor("color").Value();
	}

	public String Description() {
		return contract.GetString("description").Value();
	}

	public String Id() {
		return contract.GetString("id").Value();
	}

	public String Name() {
		return contract.GetString("name").Value();
	}

	public ScAddress Owner() {
		return contract.GetAddress("owner").Value();
	}
}
