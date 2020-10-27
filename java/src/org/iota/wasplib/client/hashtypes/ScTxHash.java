package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScContext;

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
		return hash.equals(other.hash);
	}

	@Override
	public int hashCode() {
		return hash.hashCode();
	}

	public byte[] toBytes() {
		return hash;
	}

	@Override
	public String toString() {
		return new ScContext().Utility().Base58Encode(hash);
	}
}
