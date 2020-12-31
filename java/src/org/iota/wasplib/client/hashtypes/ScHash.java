// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.MapKey;

import java.util.Arrays;

public class ScHash implements MapKey {
	final byte[] hash = new byte[32];

	public ScHash(byte[] bytes) {
		if (bytes == null || bytes.length != hash.length) {
			throw new RuntimeException("hash should be 32 bytes");
		}
		System.arraycopy(bytes, 0, hash, 0, hash.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScHash other = (ScHash) o;
		return Arrays.equals(hash, other.hash);
	}

	@Override
	public int GetId() {
		return Host.GetKeyIdFromBytes(hash);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(hash);
	}

	public byte[] toBytes() {
		return hash;
	}

	@Override

	public String toString() {
		return ScUtility.Base58String(hash);
	}
}
