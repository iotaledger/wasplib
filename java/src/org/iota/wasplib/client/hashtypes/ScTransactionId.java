package org.iota.wasplib.client.hashtypes;

public class ScTransactionId {
	final String id;

	public ScTransactionId(String bytes) {
		id = bytes;
//		if (bytes.length != id.length) {
//			throw new RuntimeException("transaction id should be 32 bytes");
//		}
//		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScTransactionId other = (ScTransactionId) o;
		return id.equals(other.id);
	}

	@Override
	public int hashCode() {
		return id.hashCode();
	}

	public String toBytes() {
		return id;
	}

	@Override
	public String toString() {
		return id;
	}
}
