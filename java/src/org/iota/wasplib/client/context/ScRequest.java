package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScRequest {
	ScImmutableMap request;

	ScRequest(ScImmutableMap request) {
		this.request = request;
	}

	public String Address() {
		return request.GetString("address").Value();
	}

	public long Balance(ScColor color) {
		return request.GetMap("balance").GetInt(color.toBytes()).Value();
	}

	public ScColors Colors() {
		return new ScColors(request.GetStringArray("colors"));
	}

	public ScColor MintedColor() {
		return new ScColor(request.GetString("hash").Value());
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
