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

    public ScAgentId AccountId() {
        return Host.root.GetAgentId(Key.AccountId).Value();
    }

    public ScBalances Balances() {
        return new ScBalances(Host.root.GetMap(Key.Balances).Immutable());
    }

    public ScChainId ChainId() {
        return Host.root.GetChainId(Key.ChainId).Value();
    }

    public ScAgentId ChainOwnerId() {
        return Host.root.GetAgentId(Key.ChainOwnerId).Value();
    }

    public ScHname Contract() {
        return Host.root.GetHname(Key.Contract).Value();
    }

    public ScAgentId ContractCreator() {
        return Host.root.GetAgentId(Key.ContractCreator).Value();
    }

    public void Log(String text) {
        Host.Log(text);
    }

    public void Panic(String text) {
        Host.Panic(text);
    }

    public ScImmutableMap Params() {
        return Host.root.GetMap(Key.Params).Immutable();
    }

    // panics with specified message if specified condition is not satisfied
    public void Require(boolean cond, String msg) {
        if (!cond) {
            Host.Panic(msg);
        }
    }

    public ScMutableMap Results() {
        return Host.root.GetMap(Key.Results);
    }

    public long Timestamp() {
        return Host.root.GetInt64(Key.Timestamp).Value();
    }

    public void Trace(String text) {
        Host.Trace(text);
    }

    public ScUtility Utility() {
        return new ScUtility(Host.root.GetMap(Key.Utility));
    }
}
