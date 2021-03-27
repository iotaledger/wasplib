// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class ScMutableMap {
    final int objId;

    public ScMutableMap() {
        ScMutableMapArray maps = Host.root.GetMapArray(Key.Maps);
        objId = maps.GetMap(maps.Length()).objId;
    }

    public ScMutableMap(int objId) {
        this.objId = objId;
    }

    public void Clear() {
        Host.Clear(objId);
    }

    public ScMutableAddress GetAddress(MapKey key) {
        return new ScMutableAddress(objId, key.KeyId());
    }

    public ScMutableAddressArray GetAddressArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_ADDRESS | ScType.TYPE_ARRAY);
        return new ScMutableAddressArray(arrId);
    }

    public ScMutableAgentId GetAgentId(MapKey key) {
        return new ScMutableAgentId(objId, key.KeyId());
    }

    public ScMutableAgentIdArray GetAgentIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_AGENT_ID | ScType.TYPE_ARRAY);
        return new ScMutableAgentIdArray(arrId);
    }

    public ScMutableBytes GetBytes(MapKey key) {
        return new ScMutableBytes(objId, key.KeyId());
    }

    public ScMutableBytesArray GetBytesArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_BYTES | ScType.TYPE_ARRAY);
        return new ScMutableBytesArray(arrId);
    }

    public ScMutableChainId GetChainId(MapKey key) {
        return new ScMutableChainId(objId, key.KeyId());
    }

    public ScMutableChainIdArray GetChainIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_CHAIN_ID | ScType.TYPE_ARRAY);
        return new ScMutableChainIdArray(arrId);
    }

    public ScMutableColor GetColor(MapKey key) {
        return new ScMutableColor(objId, key.KeyId());
    }

    public ScMutableColorArray GetColorArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_COLOR | ScType.TYPE_ARRAY);
        return new ScMutableColorArray(arrId);
    }

    public ScMutableHash GetHash(MapKey key) {
        return new ScMutableHash(objId, key.KeyId());
    }

    public ScMutableHashArray GetHashArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_HASH | ScType.TYPE_ARRAY);
        return new ScMutableHashArray(arrId);
    }

    public ScMutableHname GetHname(MapKey key) {
        return new ScMutableHname(objId, key.KeyId());
    }

    public ScMutableHnameArray GetHnameArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_HNAME | ScType.TYPE_ARRAY);
        return new ScMutableHnameArray(arrId);
    }

    public ScMutableInt64 GetInt64(MapKey key) {
        return new ScMutableInt64(objId, key.KeyId());
    }

    public ScMutableInt64Array GetInt64Array(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_INT64 | ScType.TYPE_ARRAY);
        return new ScMutableInt64Array(arrId);
    }

    public ScMutableMap GetMap(MapKey key) {
        int mapId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_MAP);
        return new ScMutableMap(mapId);
    }

    public ScMutableMapArray GetMapArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_MAP | ScType.TYPE_ARRAY);
        return new ScMutableMapArray(arrId);
    }

    public ScMutableRequestId GetRequestId(MapKey key) {
        return new ScMutableRequestId(objId, key.KeyId());
    }

    public ScMutableRequestIdArray GetRequestIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_REQUEST_ID | ScType.TYPE_ARRAY);
        return new ScMutableRequestIdArray(arrId);
    }

    public ScMutableString GetString(MapKey key) {
        return new ScMutableString(objId, key.KeyId());
    }

    public ScMutableStringArray GetStringArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_STRING | ScType.TYPE_ARRAY);
        return new ScMutableStringArray(arrId);
    }

    public ScImmutableMap Immutable() {
        return new ScImmutableMap(objId);
    }

    public int mapId() {
        return objId;
    }
}
