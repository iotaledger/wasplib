// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.MapKey;

import java.util.Arrays;

public class ScAgent implements MapKey {
	final byte[] id = new byte[37];

	public ScAgent(byte[] bytes) {
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
		ScAgent other = (ScAgent) o;
		return Arrays.equals(id, other.id);
	}

	@Override
	public int GetId() {
		return Host.GetKeyIdFromBytes(id);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(id);
	}

	public boolean IsAddress() {
		return Address().AsAgent().equals(this);
	}

	public byte[] toBytes() {
		return id;
	}

	public String toString() {
		return ScUtility.Base58String(id);
	}
}
