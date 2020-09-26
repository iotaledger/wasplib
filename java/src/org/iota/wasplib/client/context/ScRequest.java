package org.iota.wasplib.client.context;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;

public class ScRequest {
	ScImmutableMap request;

	ScRequest(ScImmutableMap request) {
		this.request = request;
	}

	public String Address() {
		return request.GetString("address").Value();
	}

	public long Balance(String color) {
		return request.GetMap("balance").GetInt(color).Value();
	}

	public ScImmutableStringArray Colors() {
		return request.GetStringArray("colors");
	}

	public String Hash() {
		return request.GetString("hash").Value();
	}

	public String Id() {
		return request.GetString("id").Value();
	}

	public ScImmutableMap Params() {
		return request.GetMap("params");
	}

	public long Timestamp() {
		return request.GetInt("timestamp").Value();
	}
}
