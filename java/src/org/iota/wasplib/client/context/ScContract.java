package org.iota.wasplib.client.context;

import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScContract {
	ScImmutableMap contract;

	ScContract(ScImmutableMap contract) {
		this.contract = contract;
	}

	public String Address() {
		return contract.GetString("address").Value();
	}

	public String Color() {
		return contract.GetString("color").Value();
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

	public String Owner() {
		return contract.GetString("owner").Value();
	}
}
