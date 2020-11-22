// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;

import java.util.Arrays;

public class ScTxHash {
	final byte[] hash = new byte[32];

	public ScTxHash(byte[] bytes) {
		if (bytes.length != hash.length) {
			throw new RuntimeException("tx hash should be 32 bytes");
		}
		System.arraycopy(bytes, 0, hash, 0, hash.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScTxHash other = (ScTxHash) o;
		return Arrays.equals(hash, other.hash);
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
