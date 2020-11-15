package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.hashtypes.ScRequestId;
import org.iota.wasplib.client.hashtypes.ScTxHash;
import org.iota.wasplib.client.immutable.ScImmutableColorArray;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScRequest {
	ScImmutableMap request;

	ScRequest(ScImmutableMap request) {
		this.request = request;
	}

	public long Balance(ScColor color) {
		return request.GetKeyMap("balance").GetInt(color.toBytes()).Value();
	}

	public ScImmutableColorArray Colors() {
		return request.GetColorArray("colors");
	}

	public Boolean From(ScAgent originator) {
		return Sender().equals(originator);
	}

	public ScRequestId Id() {
		return request.GetRequestId("id").Value();
	}

	public ScColor MintedColor() {
		return request.GetColor("hash").Value();
	}

	public ScImmutableMap Params() {
		return request.GetMap("params");
	}

	public ScAgent Sender() {
		return request.GetAgent("sender").Value();
	}

	public long Timestamp() {
		return request.GetInt("timestamp").Value();
	}

	public ScTxHash TxHash() {
		return request.GetTxHash("hash").Value();
	}
}
