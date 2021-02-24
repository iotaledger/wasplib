// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.dividend.types;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class Member {
	//@formatter:off
	public ScAddress Address; // address of dividend recipient
	public long      Factor;  // relative division factor
	//@formatter:on

	public Member() {
	}

	public Member(byte[] bytes) {
		BytesDecoder decode = new BytesDecoder(bytes);
		Address = decode.Address();
		Factor = decode.Int();
	}

	public byte[] toBytes() {
		return new BytesEncoder().
				Address(Address).
				Int(Factor).
				Data();
	}
}
