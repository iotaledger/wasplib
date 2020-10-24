package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScTransfer {
	ScMutableMap transfer;

	ScTransfer(ScMutableMap transfer) {
		this.transfer = transfer;
	}

	public void Address(ScAddress address) {
		transfer.GetAddress("address").SetValue(address);
	}

	public void Amount(long amount) {
		transfer.GetInt("amount").SetValue(amount);
	}

	public void Color(ScColor color) {
		transfer.GetColor("color").SetValue(color);
	}
}
