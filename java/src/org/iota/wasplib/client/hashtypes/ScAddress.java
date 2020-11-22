// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;

import java.util.Arrays;

public class ScAddress {
	final byte[] address = new byte[33];

	public ScAddress(byte[] bytes) {
		if (bytes == null || bytes.length != address.length) {
			throw new RuntimeException("address should be 33 bytes");
		}
		System.arraycopy(bytes, 0, address, 0, address.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAddress other = (ScAddress) o;
		return Arrays.equals(address, other.address);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(address);
	}

	public byte[] toBytes() {
		return address;
	}

	public String toString() {
		return ScUtility.Base58String(address);
	}
}
