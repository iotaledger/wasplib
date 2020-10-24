package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.ScType;
import org.iota.wasplib.client.immutable.ScImmutableKeyMap;

public class ScMutableKeyMap {
	int objId;

	public ScMutableKeyMap(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableAddress GetAddress(byte[] key) {
		return new ScMutableAddress(objId, Host.GetKey(key));
	}

	public ScMutableAddressArray GetAddressArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_BYTES_ARRAY);
		return new ScMutableAddressArray(arrId);
	}

	public ScMutableBytes GetBytes(byte[] key) {
		return new ScMutableBytes(objId, Host.GetKey(key));
	}

	public ScMutableBytesArray GetBytesArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_BYTES_ARRAY);
		return new ScMutableBytesArray(arrId);
	}

	public ScMutableColor GetColor(byte[] key) {
		return new ScMutableColor(objId, Host.GetKey(key));
	}

	public ScMutableColorArray GetColorArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_BYTES_ARRAY);
		return new ScMutableColorArray(arrId);
	}

	public ScMutableInt GetInt(byte[] key) {
		return new ScMutableInt(objId, Host.GetKey(key));
	}

	public ScMutableIntArray GetIntArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_INT_ARRAY);
		return new ScMutableIntArray(arrId);
	}

	public ScMutableKeyMap GetKeyMap(byte[] key) {
		int mapId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_MAP);
		return new ScMutableKeyMap(mapId);
	}

	public ScMutableMap GetMap(byte[] key) {
		int mapId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_MAP);
		return new ScMutableMap(mapId);
	}

	public ScMutableMapArray GetMapArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_MAP_ARRAY);
		return new ScMutableMapArray(arrId);
	}

	public ScMutableString GetString(byte[] key) {
		return new ScMutableString(objId, Host.GetKey(key));
	}

	public ScMutableStringArray GetStringArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.OBJTYPE_STRING_ARRAY);
		return new ScMutableStringArray(arrId);
	}

	public ScImmutableKeyMap Immutable() {
		return new ScImmutableKeyMap(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
