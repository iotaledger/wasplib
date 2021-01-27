// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.MapKey;

import java.util.Arrays;

public class ScContractId implements MapKey {
	final byte[] id = new byte[37];

	public ScContractId(byte[] bytes) {
		if (bytes == null || bytes.length != id.length) {
			throw new RuntimeException("invalid contract id length");
		}
		System.arraycopy(bytes, 0, id, 0, id.length);
	}

	public ScAgent AsAgent() {
		return new ScAgent(Arrays.copyOf(id, 37));
	}

	public ScChainId ChainId() {
		return new ScChainId(Arrays.copyOf(id, 33));
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScContractId other = (ScContractId) o;
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

	public Hname Hname() {
		return new Hname(Arrays.copyOfRange(id, 33, 37));
	}

	public byte[] toBytes() {
		return id;
	}

	public String toString() {
		return ScUtility.Base58String(id);
	}
}
