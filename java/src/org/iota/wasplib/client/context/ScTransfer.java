package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScTransfer {
	ScMutableMap transfer;

	ScTransfer(ScMutableMap transfer) {
		this.transfer = transfer;
	}

	public void Address(String address) {
		transfer.GetString("address").SetValue(address);
	}

	public void Amount(long amount) {
		transfer.GetInt("amount").SetValue(amount);
	}

	public void Color(String color) {
		transfer.GetString("color").SetValue(color);
	}
}
