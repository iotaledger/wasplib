// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.tokenregistry.types;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class Token {
    //@formatter:off
    public long      Created;     // creation timestamp
    public String    Description; // description what minted token represents
    public ScAgentId MintedBy;    // original minter
    public ScAgentId Owner;       // current owner
    public long      Supply;      // amount of tokens originally minted
    public long      Updated;     // last update timestamp
    public String    UserDefined; // any user defined text
    //@formatter:on

    public Token() {
    }

    public Token(byte[] bytes) {
        BytesDecoder decode = new BytesDecoder(bytes);
        Created = decode.Int64();
        Description = decode.String();
        MintedBy = decode.AgentId();
        Owner = decode.AgentId();
        Supply = decode.Int64();
        Updated = decode.Int64();
        UserDefined = decode.String();
    }

    public byte[] toBytes() {
        return new BytesEncoder().
                Int64(Created).
                String(Description).
                AgentId(MintedBy).
                AgentId(Owner).
                Int64(Supply).
                Int64(Updated).
                String(UserDefined).
                Data();
    }
}
