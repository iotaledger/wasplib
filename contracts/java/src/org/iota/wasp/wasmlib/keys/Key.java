package org.iota.wasp.wasmlib.keys;

import org.iota.wasp.wasmlib.host.*;

public class Key implements MapKey {
    // @formatter:off
	public static final Key Address         = new Key(-1);
	public static final Key Balances        = new Key(-2);
	public static final Key Base58Bytes     = new Key(-3);
	public static final Key Base58String    = new Key(-4);
	public static final Key BlsAddress      = new Key(-5);
	public static final Key BlsAggregate    = new Key(-6);
	public static final Key BlsValid        = new Key(-7);
	public static final Key Call            = new Key(-8);
	public static final Key Caller          = new Key(-9);
	public static final Key ChainOwnerId    = new Key(-10);
	public static final Key Color           = new Key(-11);
	public static final Key ContractCreator = new Key(-12);
	public static final Key ContractId      = new Key(-13);
	public static final Key Deploy          = new Key(-14);
	public static final Key Ed25519Address  = new Key(-15);
	public static final Key Ed25519Valid    = new Key(-16);
	public static final Key Event           = new Key(-17);
	public static final Key Exports         = new Key(-18);
	public static final Key HashBlake2b     = new Key(-19);
	public static final Key HashSha3        = new Key(-20);
	public static final Key Hname           = new Key(-21);
	public static final Key Incoming        = new Key(-22);
	public static final Key Length          = new Key(-23);
	public static final Key Log             = new Key(-24);
	public static final Key Maps            = new Key(-25);
	public static final Key Minted          = new Key(-26);
	public static final Key Name            = new Key(-27);
	public static final Key Panic           = new Key(-28);
	public static final Key Params          = new Key(-29);
	public static final Key Post            = new Key(-30);
	public static final Key Random          = new Key(-31);
	public static final Key RequestId       = new Key(-32);
	public static final Key Results         = new Key(-33);
	public static final Key Return          = new Key(-34);
	public static final Key State           = new Key(-35);
	public static final Key Timestamp       = new Key(-36);
	public static final Key Trace           = new Key(-37);
	public static final Key Transfers       = new Key(-38);
	public static final Key Utility         = new Key(-39);
	public static final Key Valid           = new Key(-40);
	public static final Key Zzzzzzz         = new Key(-41);
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
    public int KeyId() {
        return key == null ? keyId : Host.GetKeyIdFromString(key);
    }
}
