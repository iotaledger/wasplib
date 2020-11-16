// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScContext;

import java.util.Arrays;

public class ScColor {
	public static final ScColor IOTA = new ScColor(new byte[32]);
	public static final ScColor MINT = new ScColor(new byte[32]);
	final byte[] color = new byte[32];

	public ScColor(byte[] bytes) {
		if (bytes == null || bytes.length != color.length) {
			throw new RuntimeException("color should be 32 bytes");
		}
		System.arraycopy(bytes, 0, color, 0, color.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScColor other = (ScColor) o;
		return Arrays.equals(color, other.color);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(color);
	}

	public byte[] toBytes() {
		return color;
	}

	@Override

	public String toString() {
		return new ScContext().Utility().Base58Encode(color);
	}

	static {
		Arrays.fill(IOTA.color, (byte) 0x00);
		Arrays.fill(MINT.color, (byte) 0xff);
	}
}
