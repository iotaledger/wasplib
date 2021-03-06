// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

public class ScImmutableMap {
    final int objId;

    public ScImmutableMap(int objId) {
        this.objId = objId;
    }

    public ScImmutableAddress GetAddress(MapKey key) {
        return new ScImmutableAddress(objId, key.KeyId());
    }

    public ScImmutableAddressArray GetAddressArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_ADDRESS | ScType.TYPE_ARRAY);
        return new ScImmutableAddressArray(arrId);
    }

    public ScImmutableAgentId GetAgentId(MapKey key) {
        return new ScImmutableAgentId(objId, key.KeyId());
    }

    public ScImmutableAgentIdArray GetAgentIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_AGENT_ID | ScType.TYPE_ARRAY);
        return new ScImmutableAgentIdArray(arrId);
    }

    public ScImmutableBytes GetBytes(MapKey key) {
        return new ScImmutableBytes(objId, key.KeyId());
    }

    public ScImmutableBytesArray GetBytesArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_BYTES | ScType.TYPE_ARRAY);
        return new ScImmutableBytesArray(arrId);
    }

    public ScImmutableChainId GetChainId(MapKey key) {
        return new ScImmutableChainId(objId, key.KeyId());
    }

    public ScImmutableChainIdArray GetChainIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_CHAIN_ID | ScType.TYPE_ARRAY);
        return new ScImmutableChainIdArray(arrId);
    }

    public ScImmutableColor GetColor(MapKey key) {
        return new ScImmutableColor(objId, key.KeyId());
    }

    public ScImmutableColorArray GetColorArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_COLOR | ScType.TYPE_ARRAY);
        return new ScImmutableColorArray(arrId);
    }

    public ScImmutableHash GetHash(MapKey key) {
        return new ScImmutableHash(objId, key.KeyId());
    }

    public ScImmutableHashArray GetHashArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_HASH | ScType.TYPE_ARRAY);
        return new ScImmutableHashArray(arrId);
    }

    public ScImmutableHname GetHname(MapKey key) {
        return new ScImmutableHname(objId, key.KeyId());
    }

    public ScImmutableHnameArray GetHnameArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_HNAME | ScType.TYPE_ARRAY);
        return new ScImmutableHnameArray(arrId);
    }

    public ScImmutableInt64 GetInt64(MapKey key) {
        return new ScImmutableInt64(objId, key.KeyId());
    }

    public ScImmutableInt64Array GetInt64Array(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_INT64 | ScType.TYPE_ARRAY);
        return new ScImmutableInt64Array(arrId);
    }

    public ScImmutableMap GetMap(MapKey key) {
        int mapId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_MAP);
        return new ScImmutableMap(mapId);
    }

    public ScImmutableMapArray GetMapArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_MAP | ScType.TYPE_ARRAY);
        return new ScImmutableMapArray(arrId);
    }

    public ScImmutableRequestId GetRequestId(MapKey key) {
        return new ScImmutableRequestId(objId, key.KeyId());
    }

    public ScImmutableRequestIdArray GetRequestIdArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_REQUEST_ID | ScType.TYPE_ARRAY);
        return new ScImmutableRequestIdArray(arrId);
    }

    public ScImmutableString GetString(MapKey key) {
        return new ScImmutableString(objId, key.KeyId());
    }

    public ScImmutableStringArray GetStringArray(MapKey key) {
        int arrId = Host.GetObjectId(objId, key.KeyId(), ScType.TYPE_STRING | ScType.TYPE_ARRAY);
        return new ScImmutableStringArray(arrId);
    }
}
