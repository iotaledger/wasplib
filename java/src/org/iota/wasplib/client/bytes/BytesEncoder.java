package org.iota.wasplib.client.bytes;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.hashtypes.ScRequestId;
import org.iota.wasplib.client.hashtypes.ScTxHash;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class BytesEncoder {
	ByteArrayOutputStream data;

	public BytesEncoder() {
		data = new ByteArrayOutputStream();
	}

	public BytesEncoder Address(ScAddress value) {
		return Bytes(value.toBytes());
	}

	public BytesEncoder Bytes(byte[] value) {
		Int(value.length);
		try {
			data.write(value);
		} catch (IOException e) {
		}
		return this;
	}

	public BytesEncoder Color(ScColor value) {
		return Bytes(value.toBytes());
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

	public BytesEncoder RequestId(ScRequestId value) {
		return Bytes(value.toBytes());
	}

	public BytesEncoder String(String value) {
		return Bytes(value.getBytes(StandardCharsets.UTF_8));
	}

	public BytesEncoder TxHash(ScTxHash value) {
		return Bytes(value.toBytes());
	}
}
