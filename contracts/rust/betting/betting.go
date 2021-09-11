// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package betting

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"


func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    if f.Params.Owner().Exists() {
        f.State.Owner().SetValue(f.Params.Owner().Value())
        return
    }
    f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
    f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
    f.Results.Owner().SetValue(f.State.Owner().Value())
}
