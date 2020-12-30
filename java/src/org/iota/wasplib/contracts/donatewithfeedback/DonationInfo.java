// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.donatewithfeedback;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class DonationInfo {
	//@formatter:off
	public long    amount;
	public ScAgent donator;
	public String  error;
	public String  feedback;
	public long    timestamp;
	//@formatter:on

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
		BytesDecoder decode = new BytesDecoder(bytes);
		DonationInfo data = new DonationInfo();
		data.amount = decode.Int();
		data.donator = decode.Agent();
		data.error = decode.String();
		data.feedback = decode.String();
		data.timestamp = decode.Int();
		return data;
	}
}
