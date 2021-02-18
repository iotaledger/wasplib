// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.bytes;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.Host;

import java.nio.charset.StandardCharsets;
import java.util.Arrays;

public class BytesDecoder {
	byte[] data;

	public BytesDecoder(byte[] data) {
		this.data = data;
	}

	public ScAddress Address() {
		return new ScAddress(Bytes());
	}

	public ScAgentId AgentId() {
		return new ScAgentId(Bytes());
	}

	public byte[] Bytes() {
		int size = (int) Int();
		if (data.length < size) {
			Host.panic("Cannot decode bytes");
		}
		byte[] value = Arrays.copyOfRange(data, 0, size);
		data = Arrays.copyOfRange(data, size, data.length);
		return value;
	}

	public ScChainId ChainId() {
		return new ScChainId(Bytes());
	}

	public ScColor Color() {
		return new ScColor(Bytes());
	}

	public ScContractId ContractId() {
		return new ScContractId(Bytes());
	}

	public ScHash Hash() {
		return new ScHash(Bytes());
	}

	public Hname Hname() {
		return new Hname(Bytes());
	}

	public long Int() {
		long val = 0;
		int s = 0;
		for (; ; ) {
			byte b = data[0];
			data = Arrays.copyOfRange(data, 1, data.length);
			val |= ((long) (b & 0x7f)) << s;
			if (b >= 0) {
				if (((byte) (val >> s) & 0x7f) != (b & 0x7f)) {
					Host.panic("integer too large");
					return 0;
				}
				// extend int7 sign to int8
				if ((b & 0x40) != 0) {
					b |= 0x80;
				}
				// extend int8 sign to int64
				return val | ((long) b << s);
			}
			s += 7;
			if (s >= 64) {
				Host.panic("integer representation too long");
				return 0;
			}
		}
	}

	public String String() {
		return new String(Bytes(), StandardCharsets.UTF_8);
	}
}
