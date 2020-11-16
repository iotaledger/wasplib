// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client;

public class Keys {
	private static int keyLength;
	private static int keyLog;
	private static int keyTrace;

	public static int KeyLength() {
		if (keyLength == 0) {
			keyLength = Host.GetKeyId("length");
		}
		return keyLength;
	}

	public static int KeyLog() {
		if (keyLog == 0) {
			keyLog = Host.GetKeyId("log");
		}
		return keyLog;
	}

	public static int KeyTrace() {
		if (keyTrace == 0) {
			keyTrace = Host.GetKeyId("trace");
		}
		return keyTrace;
	}
}
