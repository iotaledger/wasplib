package org.iota.wasplib.client.bytes;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class BytesEncoder {
	ByteArrayOutputStream data;

	public BytesEncoder() {
		data = new ByteArrayOutputStream();
	}

	public BytesEncoder Bytes(byte[] value) {
		Int(value.length);
		try {
			data.write(value);
		} catch (IOException e) {
		}
		return this;
	}

	public byte[] Data() {
		return data.toByteArray();
	}

	public BytesEncoder Int(long value) {
		for (; ; ) {
			byte b = (byte) value;
			byte s = (byte) (b & 0x40);
			value >>= 7;
			if ((value == 0 && s == 0) || (value == -1 && s != 0)) {
				data.write((byte) (b & 0x7f));
				return this;
			}
			data.write((byte) (b | 0x80));
		}
	}

	public BytesEncoder String(String value) {
		return Bytes(value.getBytes(StandardCharsets.UTF_8));
	}
}
