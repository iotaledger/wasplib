package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableBytes;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScUtility {
	ScMutableMap utility;

	ScUtility(ScMutableMap utility) {
		this.utility = utility;
	}

	public byte[] Base58Decode(String value) {
		ScMutableString decode = utility.GetString("base58");
		ScMutableBytes encode = utility.GetBytes("base58");
		decode.SetValue(value);
		return encode.Value();
	}

	public String Base58Encode(byte[] value) {
		ScMutableString decode = utility.GetString("base58");
		ScMutableBytes encode = utility.GetBytes("base58");
		encode.SetValue(value);
		return decode.Value();
	}

	public String Hash(String value) {
		ScMutableString hash = utility.GetString("hash");
		hash.SetValue(value);
		return hash.Value();
	}
}
