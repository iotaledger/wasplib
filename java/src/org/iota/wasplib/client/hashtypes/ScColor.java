package org.iota.wasplib.client.hashtypes;

public class ScColor {
	public static final ScColor IOTA = new ScColor("iota");
	public static final ScColor MINT = new ScColor("new");

	final String color;

	public ScColor(String bytes) {
		color = bytes;
//		if (bytes.length != color.length) {
//			throw new RuntimeException("color should be 32 bytes");
//		}
//		System.arraycopy(bytes, 0, color, 0, color.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScColor other = (ScColor) o;
		return color.equals(other.color);
	}

	@Override
	public int hashCode() {
		return color.hashCode();
	}

	public String toBytes() {
		return color;
	}

	@Override
	public String toString() {
		return color;
	}
}
