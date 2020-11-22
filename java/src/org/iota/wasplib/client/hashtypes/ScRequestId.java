// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;

import java.util.Arrays;

public class ScRequestId {
	final byte[] id = new byte[34];

	public ScRequestId(byte[] bytes) {
		if (bytes.length != id.length) {
			throw new RuntimeException("request id should be 34 bytes");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScRequestId other = (ScRequestId) o;
		return Arrays.equals(id, other.id);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(id);
	}

	public byte[] toBytes() {
		return id;
	}

	@Override
	public String toString() {
		return ScUtility.Base58String(id);
	}
}
