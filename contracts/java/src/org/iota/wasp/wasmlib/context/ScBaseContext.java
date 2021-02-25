// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScBaseContext {
    protected ScBaseContext() {
    }

    public ScBalances Balances() {
        return new ScBalances(Host.root.GetMap(Key.Balances).Immutable());
    }

    public ScAgentId ChainOwnerId() {
        return Host.root.GetAgentId(Key.ChainOwnerId).Value();
    }

    public ScAgentId ContractCreator() {
        return Host.root.GetAgentId(Key.ContractCreator).Value();
    }

    public ScContractId ContractId() {
        return Host.root.GetContractId(Key.ContractId).Value();
    }

    public void Log(String text) {
        Host.root.GetString(Key.Log).SetValue(text);
    }

    public void Panic(String text) {
        Host.root.GetString(Key.Panic).SetValue(text);
    }

    public ScImmutableMap Params() {
        return Host.root.GetMap(Key.Params).Immutable();
    }

    // panics with specified message if specified condition is not satisfied
    public void Require(boolean cond, String msg) {
        if (!cond) {
            Panic(msg);
        }
    }

    public ScMutableMap Results() {
        return Host.root.GetMap(Key.Results);
    }

    public long Timestamp() {
        return Host.root.GetInt64(Key.Timestamp).Value();
    }

    public void Trace(String text) {
        Host.root.GetString(Key.Trace).SetValue(text);
    }

    public ScUtility Utility() {
        return new ScUtility(Host.root.GetMap(Key.Utility));
    }
}
