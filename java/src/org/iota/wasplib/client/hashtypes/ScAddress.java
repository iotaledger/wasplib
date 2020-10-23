package org.iota.wasplib.client.hashtypes;

public class ScAddress {
	final String address;

	public ScAddress(String bytes) {
		address = bytes;
//		if (bytes.length != address.length) {
//			throw new RuntimeException("address should be 33 bytes");
//		}
//		System.arraycopy(bytes, 0, address, 0, address.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAddress other = (ScAddress) o;
		return address.equals(other.address);
	}

	@Override
	public int hashCode() {
		return address.hashCode();
	}

	public String toBytes() {
		return address;
	}

	public String toString() {
		return address;
	}
}
