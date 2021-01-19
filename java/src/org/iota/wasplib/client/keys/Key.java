package org.iota.wasplib.client.keys;

import org.iota.wasplib.client.host.Host;

public class Key implements MapKey {
	// @formatter:off
	public static final Key Agent       = new Key(-1);
	public static final Key Balances    = new Key(-2);
	public static final Key Base58      = new Key(-3);
	public static final Key Caller      = new Key(-4);
	public static final Key Calls       = new Key(-5);
	public static final Key Chain       = new Key(-6);
	public static final Key ChainOwner  = new Key(-7);
	public static final Key Color       = new Key(-8);
	public static final Key Contract    = new Key(-9);
	public static final Key Creator     = new Key(-10);
	public static final Key Data        = new Key(-11);
	public static final Key Delay       = new Key(-12);
	public static final Key Deploys     = new Key(-13);
	public static final Key Description = new Key(-14);
	public static final Key Event       = new Key(-15);
	public static final Key Exports     = new Key(-16);
	public static final Key Function    = new Key(-17);
	public static final Key Hash        = new Key(-18);
	public static final Key Id          = new Key(-19);
	public static final Key Incoming    = new Key(-20);
	public static final Key Length      = new Key(-21);
	public static final Key Log         = new Key(-22);
	public static final Key Logs        = new Key(-23);
	public static final Key Name        = new Key(-24);
	public static final Key Panic       = new Key(-25);
	public static final Key Params      = new Key(-26);
	public static final Key Posts       = new Key(-27);
	public static final Key Random      = new Key(-28);
	public static final Key Results     = new Key(-29);
	public static final Key State       = new Key(-30);
	public static final Key Timestamp   = new Key(-31);
	public static final Key Trace       = new Key(-32);
	public static final Key Transfers   = new Key(-33);
	public static final Key Utility     = new Key(-34);
	public static final Key Views       = new Key(-35);
	public static final Key Zzzzzzz     = new Key(-36);
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
