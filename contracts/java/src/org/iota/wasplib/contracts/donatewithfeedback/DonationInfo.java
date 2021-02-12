// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.donatewithfeedback;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.hashtypes.ScAgent;

public class DonationInfo {
	//@formatter:off
	public long    Amount;
	public ScAgent Donator;
	public String  Error;
	public String  Feedback;
	public long    Timestamp;
	//@formatter:on

	public static byte[] encode(DonationInfo o) {
		return new BytesEncoder().
				Int(o.Amount).
				Agent(o.Donator).
				String(o.Error).
				String(o.Feedback).
				Int(o.Timestamp).
				Data();
	}

	public static DonationInfo decode(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		DonationInfo data = new DonationInfo();
		data.Amount = decode.Int();
		data.Donator = decode.Agent();
		data.Error = decode.String();
		data.Feedback = decode.String();
		data.Timestamp = decode.Int();
		return data;
	}
}
