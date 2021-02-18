// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.dividend;

public class Dividend {

public static void funcDivide(ScFuncContext ctx, FuncDivideParams params) {
    amount = ctx.Balances().Balance(ScColor.IOTA);
    if (amount == 0) {
        ctx.Panic("Nothing to divide");
    }
    state = ctx.State();
    totalFactor = state.GetInt(VarTotalFactor);
    total = totalFactor.Value();
    members = state.GetBytesArray(VarMembers);
    parts = 0;
    size = members.Length();
    for (int i = 0; i < size; i++) {
        m = Member::fromBytes(members.GetBytes(i).Value());
        part = amount * m.Factor / total;
        if (part != 0) {
            parts += part;
            ctx.TransferToAddress(m.Address, new ScTransfers(ScColor.IOTA, part));
        }
    }
    if (parts != amount) {
        // note we truncated the calculations down to the nearest integer
        // there could be some small remainder left in the contract, but
        // that will be picked up in the next round as part of the balance
        remainder = amount - parts;
        ctx.Log("Remainder in contract: " + remainder);
    }
}

public static void funcMember(ScFuncContext ctx, FuncMemberParams params) {
    Member member = new Member();
         {
        member.Address = params.Address.Value();
        member.Factor = params.Factor.Value();
    }
    state = ctx.State();
    totalFactor = state.GetInt(VarTotalFactor);
    total = totalFactor.Value();
    members = state.GetBytesArray(VarMembers);
    size = members.Length();
    for (int i = 0; i < size; i++) {
        m = Member::fromBytes(members.GetBytes(i).Value());
        if (m.Address == member.Address) {
            total -= m.Factor;
            total += member.Factor;
            totalFactor.SetValue(total);
            members.GetBytes(i).SetValue(member.ToBytes());
            ctx.Log("Updated: " + member.Address);
            return;
        }
    }
    total += member.Factor;
    totalFactor.SetValue(total);
    members.GetBytes(size).SetValue(member.ToBytes());
    ctx.Log("Appended: " + member.Address);
}
}
