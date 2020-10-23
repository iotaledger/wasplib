package org.iota.wasplib.client.hashtypes;

public class ScRequestId {
	final String id;

	public ScRequestId(String bytes) {
		id = bytes;
//		if (bytes.length != id.length) {
//			throw new RuntimeException("request id should be 34 bytes");
//		}
//		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScRequestId other = (ScRequestId) o;
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

//	public ScTransactionId TransactionId() {
//		return new ScTransactionId(Arrays.copyOfRange(id, 0, 32));
//	}
}
