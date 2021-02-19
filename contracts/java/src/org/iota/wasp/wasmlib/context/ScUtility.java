// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScUtility {
	ScMutableMap utility;

	ScUtility(ScMutableMap utility) {
		this.utility = utility;
	}

	public static String Base58String(byte[] bytes) {
		return new ScFuncContext().Utility().Base58Encode(bytes);
	}

	public byte[] Base58Decode(String value) {
		ScMutableString decode = utility.GetString(Key.Base58String);
		ScMutableBytes encode = utility.GetBytes(Key.Base58Bytes);
		decode.SetValue(value);
		return encode.Value();
	}

	public String Base58Encode(byte[] value) {
		ScMutableString decode = utility.GetString(Key.Base58String);
		ScMutableBytes encode = utility.GetBytes(Key.Base58Bytes);
		encode.SetValue(value);
		return decode.Value();
	}

	public ScHash Hash(byte[] value) {
		ScMutableBytes hash = utility.GetBytes(Key.HashBlake2b);
		hash.SetValue(value);
		return new ScHash(hash.Value());
	}

	public long Random(long max) {
		long rnd = utility.GetInt(Key.Random).Value();
		return Long.remainderUnsigned(rnd, max);
	}
}
