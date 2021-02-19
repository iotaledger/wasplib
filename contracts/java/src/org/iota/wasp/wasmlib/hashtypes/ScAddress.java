// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

import java.util.*;

public class ScAddress implements MapKey {
	final byte[] id = new byte[33];

	public ScAddress(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("invalid address id length");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	public ScAgentId AsAgentId() {
		return new ScAgentId(Arrays.copyOf(id, 37));
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAddress other = (ScAddress) o;
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

	public String toString() {
		return ScUtility.base58Encode(id);
	}
}
