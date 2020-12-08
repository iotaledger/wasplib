// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScUtility {
	ScMutableMap utility;

	ScUtility(ScMutableMap utility) {
		this.utility = utility;
	}

	public static String Base58String(byte[] bytes) {
		return new ScCallContext().Utility().Base58Encode(bytes);
	}

	public byte[] Base58Decode(String value) {
		ScMutableString decode = utility.GetString(Key.Base58);
		ScMutableBytes encode = utility.GetBytes(Key.Base58);
		decode.SetValue(value);
		return encode.Value();
	}

	public String Base58Encode(byte[] value) {
		ScMutableString decode = utility.GetString(Key.Base58);
		ScMutableBytes encode = utility.GetBytes(Key.Base58);
		encode.SetValue(value);
		return decode.Value();
	}

	public String Hash(String value) {
		ScMutableString hash = utility.GetString(Key.Hash);
		hash.SetValue(value);
		return hash.Value();
	}

	public long Random(long max) {
		long rnd = utility.GetInt(Key.Random).Value();
		return Long.remainderUnsigned(rnd, max);
	}

	public String String(long value) {
		return "" + value;
	}
}
