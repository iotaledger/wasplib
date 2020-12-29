package org.iota.wasplib.client.keys;

import org.iota.wasplib.client.host.Host;

public class Key implements MapKey {
	// @formatter:off
	public static final int KEY_AGENT       = -1;
	public static final int KEY_BALANCES    = KEY_AGENT       -1;
	public static final int KEY_BASE58      = KEY_BALANCES    -1;
	public static final int KEY_CALLER      = KEY_BASE58      -1;
	public static final int KEY_CALLS       = KEY_CALLER      -1;
	public static final int KEY_CHAIN       = KEY_CALLS       -1;
	public static final int KEY_CHAIN_OWNER = KEY_CHAIN       -1;
	public static final int KEY_COLOR       = KEY_CHAIN_OWNER -1;
	public static final int KEY_CONTRACT    = KEY_COLOR       -1;
	public static final int KEY_CREATOR     = KEY_CONTRACT    -1;
	public static final int KEY_DATA        = KEY_CREATOR     -1;
	public static final int KEY_DELAY       = KEY_DATA        -1;
	public static final int KEY_DESCRIPTION = KEY_DELAY       -1;
	public static final int KEY_EVENT       = KEY_DESCRIPTION -1;
	public static final int KEY_EXPORTS     = KEY_EVENT       -1;
	public static final int KEY_FUNCTION    = KEY_EXPORTS     -1;
	public static final int KEY_HASH        = KEY_FUNCTION    -1;
	public static final int KEY_ID          = KEY_HASH        -1;
	public static final int KEY_INCOMING    = KEY_ID          -1;
	public static final int KEY_LENGTH      = KEY_INCOMING    -1;
	public static final int KEY_LOG         = KEY_LENGTH      -1;
	public static final int KEY_LOGS        = KEY_LOG         -1;
	public static final int KEY_NAME        = KEY_LOGS        -1;
	public static final int KEY_PANIC       = KEY_NAME        -1;
	public static final int KEY_PARAMS      = KEY_PANIC       -1;
	public static final int KEY_POSTS       = KEY_PARAMS      -1;
	public static final int KEY_RANDOM      = KEY_POSTS       -1;
	public static final int KEY_RESULTS     = KEY_RANDOM      -1;
	public static final int KEY_STATE       = KEY_RESULTS     -1;
	public static final int KEY_TIMESTAMP   = KEY_STATE       -1;
	public static final int KEY_TRACE       = KEY_TIMESTAMP   -1;
	public static final int KEY_TRANSFERS   = KEY_TRACE       -1;
	public static final int KEY_UTILITY     = KEY_TRANSFERS   -1;
	public static final int KEY_VIEWS       = KEY_UTILITY     -1;
	public static final int KEY_ZZZZZZZ     = KEY_VIEWS       -1;

	public static final Key Agent       = new Key(KEY_AGENT);
	public static final Key Balances    = new Key(KEY_BALANCES);
	public static final Key Base58      = new Key(KEY_BASE58);
	public static final Key Caller      = new Key(KEY_CALLER);
	public static final Key Calls       = new Key(KEY_CALLS);
	public static final Key Chain       = new Key(KEY_CHAIN);
	public static final Key ChainOwner  = new Key(KEY_CHAIN_OWNER);
	public static final Key Color       = new Key(KEY_COLOR);
	public static final Key Contract    = new Key(KEY_CONTRACT);
	public static final Key Creator     = new Key(KEY_CREATOR);
	public static final Key Data        = new Key(KEY_DATA);
	public static final Key Delay       = new Key(KEY_DELAY);
	public static final Key Description = new Key(KEY_DESCRIPTION);
	public static final Key Event       = new Key(KEY_EVENT);
	public static final Key Exports     = new Key(KEY_EXPORTS);
	public static final Key Function    = new Key(KEY_FUNCTION);
	public static final Key Hash        = new Key(KEY_HASH);
	public static final Key Id          = new Key(KEY_ID);
	public static final Key Incoming    = new Key(KEY_INCOMING);
	public static final Key Length      = new Key(KEY_LENGTH);
	public static final Key Log         = new Key(KEY_LOG);
	public static final Key Logs        = new Key(KEY_LOGS);
	public static final Key Name        = new Key(KEY_NAME);
	public static final Key Panic       = new Key(KEY_PANIC);
	public static final Key Params      = new Key(KEY_PARAMS);
	public static final Key Posts       = new Key(KEY_POSTS);
	public static final Key Random      = new Key(KEY_RANDOM);
	public static final Key Results     = new Key(KEY_RESULTS);
	public static final Key State       = new Key(KEY_STATE);
	public static final Key Timestamp   = new Key(KEY_TIMESTAMP);
	public static final Key Trace       = new Key(KEY_TRACE);
	public static final Key Transfers   = new Key(KEY_TRANSFERS);
	public static final Key Utility     = new Key(KEY_UTILITY);
	public static final Key Views       = new Key(KEY_VIEWS);
	// @formatter:on

	int keyId;
	String key;

	public Key(String key) {
		this.key = key;
	}

	public Key(int keyId) {
		this.keyId = keyId;
	}

	@Override
	public int GetId() {
		return key == null ? keyId : Host.GetKeyIdFromString(key);
	}
}
