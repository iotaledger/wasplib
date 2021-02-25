// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScUtility {
    ScMutableMap utility;

    ScUtility(ScMutableMap utility) {
        this.utility = utility;
    }

    public static String base58Encode(byte[] bytes) {
        return new ScFuncContext().Utility().Base58Encode(bytes);
    }

    public byte[] Base58Decode(String value) {
        utility.GetString(Key.Base58String).SetValue(value);
        return utility.GetBytes(Key.Base58Bytes).Value();
    }

    public String Base58Encode(byte[] value) {
        utility.GetBytes(Key.Base58Bytes).SetValue(value);
        return utility.GetString(Key.Base58String).Value();
    }

    public ScAddress BlsAddressFromPubKey(byte[] pubKey) {
        utility.GetBytes(Key.BlsAddress).SetValue(pubKey);
        return utility.GetAddress(Key.Address).Value();
    }

    public byte[][] BlsAggregateSignatures(byte[][] pubKeys, byte[][] sigs) {
        BytesEncoder encode = new BytesEncoder();
        encode.Int64(pubKeys.length);
        for (int i = 0; i < pubKeys.length; i++) {
            encode.Bytes(pubKeys[i]);
        }
        encode.Int64(sigs.length);
        for (int i = 0; i < sigs.length; i++) {
            encode.Bytes(sigs[i]);
        }
        ScMutableBytes aggregator = utility.GetBytes(Key.BlsAggregate);
        aggregator.SetValue(encode.Data());
        BytesDecoder decode = new BytesDecoder(aggregator.Value());
        byte[][] ret = new byte[2][];
        ret[0] = decode.Bytes();
        ret[1] = decode.Bytes();
        return ret;
    }

    public boolean BlsValidSignature(byte[] data, byte[] pubKey, byte[] signature) {
        byte[] bytes = new BytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature).Data();
        utility.GetBytes(Key.BlsValid).SetValue(bytes);
        return utility.GetInt64(Key.Valid).Value() != 0;
    }

    public ScAddress Ed25519AddressFromPubKey(byte[] pubKey) {
        utility.GetBytes(Key.Ed25519Address).SetValue(pubKey);
        return utility.GetAddress(Key.Address).Value();
    }

    public boolean Ed25519ValidSignature(byte[] data, byte[] pubKey, byte[] signature) {
        byte[] bytes = new BytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature).Data();
        utility.GetBytes(Key.Ed25519Valid).SetValue(bytes);
        return utility.GetInt64(Key.Valid).Value() != 0;
    }

    public ScHash HashBlake2b(byte[] value) {
        ScMutableBytes hash = utility.GetBytes(Key.HashBlake2b);
        hash.SetValue(value);
        return new ScHash(hash.Value());
    }

    public ScHash HashSha3(byte[] value) {
        ScMutableBytes hash = utility.GetBytes(Key.HashSha3);
        hash.SetValue(value);
        return new ScHash(hash.Value());
    }

    public ScHname Hname(String value) {
        utility.GetString(Key.Name).SetValue(value);
        return utility.GetHname(Key.Hname).Value();
    }

    public long Random(long max) {
        long rnd = utility.GetInt64(Key.Random).Value();
        return Long.remainderUnsigned(rnd, max);
    }
}
