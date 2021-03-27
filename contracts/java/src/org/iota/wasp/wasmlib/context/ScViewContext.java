// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScViewContext extends ScBaseContext {
    public ScViewContext() {
    }

    public ScImmutableMap Call(ScHname hContract, ScHname hFunction, ScMutableMap params) {
        BytesEncoder encode = new BytesEncoder();
        encode.Hname(hContract);
        encode.Hname(hFunction);
        if (params != null) {
            encode.Int64(params.mapId());
        } else {
            encode.Int64(0);
        }
        encode.Int64(0);
        Host.root.GetBytes(Key.Call).SetValue(encode.Data());
        return Host.root.GetMap(Key.Return).Immutable();
    }

    public ScImmutableMap CallSelf(ScHname hFunction, ScMutableMap params) {
        return Call(Contract(), hFunction, params);
    }

    public ScImmutableMap State() {
        return Host.root.GetMap(Key.State).Immutable();
    }
}
