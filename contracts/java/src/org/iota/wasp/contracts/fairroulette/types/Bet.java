// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.fairroulette.types;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class Bet {
    //@formatter:off
    public long      Amount;
    public ScAgentId Better;
    public long      Number;
    //@formatter:on

    public Bet() {
    }

    public Bet(byte[] bytes) {
        BytesDecoder decode = new BytesDecoder(bytes);
        Amount = decode.Int64();
        Better = decode.AgentId();
        Number = decode.Int64();
        decode.Close();
    }

    public byte[] toBytes() {
        return new BytesEncoder().
                Int64(Amount).
                AgentId(Better).
                Int64(Number).
                Data();
    }
}
