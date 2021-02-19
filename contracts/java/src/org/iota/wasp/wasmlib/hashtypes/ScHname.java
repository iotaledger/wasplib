// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScHname implements MapKey {
	int id;

	public ScHname(int id) {
		this.id = id;
	}

	public ScHname(String name) {
		ScMutableMap utility = Host.root.GetMap(Key.Utility);
		utility.GetString(Key.Name).SetValue(name);
		fromBytes(utility.GetBytes(Key.Name).Value());
	}

	public ScHname(byte[] bytes) {
		fromBytes(bytes);
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		ScHname other = (ScHname) o;
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
