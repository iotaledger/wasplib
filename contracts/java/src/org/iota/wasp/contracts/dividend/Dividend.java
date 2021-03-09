// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.dividend;

import org.iota.wasp.contracts.dividend.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class Dividend {

    public static void funcDivide(ScFuncContext ctx, FuncDivideParams params) {
        var amount = ctx.Balances().Balance(ScColor.IOTA);
        if (amount == 0) {
            ctx.Panic("Nothing to divide");
        }
        var state = ctx.State();
        var total = state.GetInt64(Consts.VarTotalFactor).Value();
        var members = state.GetMap(Consts.VarMembers);
        var memberList = state.GetAddressArray(Consts.VarMemberList);
        var size = memberList.Length();
        var parts = 0;
        for (var i = 0; i < size; i++) {
            var address = memberList.GetAddress(i).Value();
            var factor = members.GetInt64(address).Value();
            var share = amount * factor / total;
            if (share != 0) {
                parts += share;
                var transfers = new ScTransfers(ScColor.IOTA, share);
                ctx.TransferToAddress(address, transfers);
            }
        }
        if (parts != amount) {
            // note we truncated the calculations down to the nearest integer
            // there could be some small remainder left in the contract, but
            // that will be picked up in the next round as part of the balance
            var remainder = amount - parts;
            ctx.Log("Remainder in contract: " + remainder);
        }
    }

    public static void funcMember(ScFuncContext ctx, FuncMemberParams params) {
        var state = ctx.State();
        var members = state.GetMap(Consts.VarMembers);
        var address = params.Address.Value();
        var currentFactor = members.GetInt64(address);
        if (!currentFactor.Exists()) {
            // add new address to member list
            var memberList = state.GetAddressArray(Consts.VarMemberList);
            memberList.GetAddress(memberList.Length()).SetValue(address);
        }
        var factor = params.Factor.Value();
        var totalFactor = state.GetInt64(Consts.VarTotalFactor);
        var newTotalFactor = totalFactor.Value() - currentFactor.Value() + factor;
        totalFactor.SetValue(newTotalFactor);
        currentFactor.SetValue(factor);
    }

    public static void viewGetFactor(ScViewContext ctx, ViewGetFactorParams params) {
        var address = params.Address.Value();
        var members = ctx.State().GetMap(Consts.VarMembers);
        var factor = members.GetInt64(address).Value();
        ctx.Results().GetInt64(Consts.VarFactor).SetValue(factor);
    }
}
