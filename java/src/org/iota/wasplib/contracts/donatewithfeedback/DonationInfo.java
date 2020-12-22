// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.donatewithfeedback;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class DonationInfo {
	public long amount;
	public ScAgent donator;
	public String error;
	public String feedback;
	public long timestamp;

	public static byte[] encode(DonationInfo o) {
		return new BytesEncoder().
				Int(o.amount).
				Agent(o.donator).
				String(o.error).
				String(o.feedback).
				Int(o.timestamp).
				Data();
	}

	public static DonationInfo decode(byte[] bytes) {
		BytesDecoder d = new BytesDecoder(bytes);
		DonationInfo data = new DonationInfo();
		data.amount = d.Int();
		data.donator = d.Agent();
		data.error = d.String();
		data.feedback = d.String();
		data.timestamp = d.Int();
		return data;
	}
}
