package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScContext;

import java.util.Arrays;

public class ScAgent {
	final byte[] id = new byte[37];

	public ScAgent(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("agent id should be 37 bytes");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAgent other = (ScAgent) o;
		return Arrays.equals(id, other.id);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(id);
	}

	public byte[] toBytes() {
		return id;
	}

	public String toString() {
		return new ScContext().Utility().Base58Encode(id);
	}
}
