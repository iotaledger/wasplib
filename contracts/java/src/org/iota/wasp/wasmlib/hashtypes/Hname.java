// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.keys.MapKey;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

public class Hname implements MapKey {
	int id;

	public Hname(int id) {
		this.id = id;
	}

	public Hname(String name) {
		ScMutableMap utility = Host.root.GetMap(Key.Utility);
		utility.GetString(Key.Name).SetValue(name);
		fromBytes(utility.GetBytes(Key.Name).Value());
	}

	public Hname(byte[] bytes) {
		fromBytes(bytes);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		Hname other = (Hname) o;
		return id == other.id;
	}

	private void fromBytes(byte[] bytes) {
		if (bytes == null || bytes.length != 4) {
			throw new RuntimeException("invalid hname length");
		}
		id = (bytes[0] & 0xff) | ((bytes[1] & 0xff) << 8) | ((bytes[2] & 0xff) << 16) | ((bytes[3] & 0xff) << 24);
	}

	@Override
	public int GetId() {
		return Host.GetKeyIdFromBytes(toBytes());
	}

	@Override
	public int hashCode() {
		return id;
	}

	public byte[] toBytes() {
		byte[] bytes = new byte[4];
		bytes[0] = (byte) id;
		bytes[1] = (byte) (id >> 8);
		bytes[2] = (byte) (id >> 16);
		bytes[3] = (byte) (id >> 24);
		return bytes;
	}

	@Override

	public String toString() {
		return "" + id;
	}
}
