package org.iota.wasplib;

import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableStringArray;

import java.util.ArrayList;

public class TokenRegistry {
	//export mintSupply
	public static void mintSupply() {
		ScContext ctx = new ScContext();
	}

	//export updateMetadata
	public static void updateMetadata() {
		ScContext ctx = new ScContext();
	}

	//export transferOwnership
	public static void transferOwnership() {
		ScContext ctx = new ScContext();
	}

	private static class TokenInfo {
		long supply;
		String mintedBy;
		String owner;
		long created;
		long updated;
		String description;
		String userDefined;
	}
}
