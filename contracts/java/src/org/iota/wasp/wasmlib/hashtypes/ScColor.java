// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

import java.util.*;

public class ScColor implements MapKey {
	public static final ScColor IOTA = new ScColor(new byte[32]);
	public static final ScColor MINT = new ScColor(new byte[32]);

	final byte[] id = new byte[32];

	public ScColor(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("invalid color id length");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScColor other = (ScColor) o;
		return Arrays.equals(id, other.id);
	}

	@Override
	public int KeyId() {
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
		return ScUtility.base58Encode(id);
	}

	static {
		Arrays.fill(IOTA.id, (byte) 0x00);
		Arrays.fill(MINT.id, (byte) 0xff);
	}
}
