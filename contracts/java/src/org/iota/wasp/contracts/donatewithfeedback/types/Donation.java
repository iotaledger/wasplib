// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.donatewithfeedback.types;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class Donation {
    //@formatter:off
    public long      Amount;    // amount donated
    public ScAgentId Donator;   // who donated
    public String    Error;     // error to be reported to donator if anything goes wrong
    public String    Feedback;  // the feedback for the person donated to
    public long      Timestamp; // when the donation took place
    //@formatter:on

    public Donation() {
    }

    public Donation(byte[] bytes) {
        BytesDecoder decode = new BytesDecoder(bytes);
        Amount = decode.Int64();
        Donator = decode.AgentId();
        Error = decode.String();
        Feedback = decode.String();
        Timestamp = decode.Int64();
        decode.Close();
    }

    public byte[] toBytes() {
        return new BytesEncoder().
                Int64(Amount).
                AgentId(Donator).
                String(Error).
                String(Feedback).
                Int64(Timestamp).
                Data();
    }
}
