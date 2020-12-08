// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.hashtypes;

import org.iota.wasplib.client.context.ScUtility;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.MapKey;

import java.util.Arrays;

public class ScAgent implements MapKey {
	final byte[] agent = new byte[37];

	public ScAgent(byte[] bytes) {
		if (bytes == null || bytes.length != agent.length) {
			throw new RuntimeException("agent id should be 37 bytes");
		}
		System.arraycopy(bytes, 0, agent, 0, agent.length);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScAgent other = (ScAgent) o;
		return Arrays.equals(agent, other.agent);
	}

	@Override
	public int GetId() {
		return Host.GetKey(agent);
	}

	@Override
	public int hashCode() {
		return Arrays.hashCode(agent);
	}

	public byte[] toBytes() {
		return agent;
	}

	public String toString() {
		return ScUtility.Base58String(agent);
	}
}
