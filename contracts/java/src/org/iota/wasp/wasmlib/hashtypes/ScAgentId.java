// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

import java.util.*;

public class ScAgentId implements MapKey {
	final byte[] id = new byte[37];

	public ScAgentId(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("invalid agent id length");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	public ScAddress Address() {
		return new ScAddress(Arrays.copyOf(id, 33));
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAgentId other = (ScAgentId) o;
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

	public boolean IsAddress() {
		return Address().AsAgentId().equals(this);
	}

	public byte[] toBytes() {
		return id;
	}

	public String toString() {
		return ScUtility.base58Encode(id);
	}
}
