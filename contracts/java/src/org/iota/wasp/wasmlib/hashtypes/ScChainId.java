// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.context.ScUtility;
import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.keys.MapKey;

import java.util.Arrays;

public class ScChainId implements MapKey {
	final byte[] id = new byte[33];

	public ScChainId(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("invalid chain id length");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScChainId other = (ScChainId) o;
		return Arrays.equals(id, other.id);
	}

	@Override
	public int GetId() {
		return Host.GetKeyIdFromBytes(id);
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
