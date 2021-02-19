// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.dividend;

import org.iota.wasp.contracts.dividend.lib.*;
import org.iota.wasp.contracts.dividend.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class Dividend {
	private static final Key KeyAddress = new Key("address");
	private static final Key KeyFactor = new Key("factor");
	private static final Key KeyMembers = new Key("members");
	private static final Key KeyTotalFactor = new Key("total_factor");

	public static void FuncDivide(ScFuncContext ctx, FuncDivideParams params) {
		long amount = ctx.Balances().Balance(ScColor.IOTA);
		if (amount == 0) {
			ctx.Panic("Nothing to divide");
		}
		ScMutableMap state = ctx.State();
		ScMutableInt totalFactor = state.GetInt(KeyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(KeyMembers);
		long parts = 0;
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = new Member(members.GetBytes(i).Value());
			long part = amount * m.Factor / total;
			if (part != 0) {
				parts += part;
				ctx.TransferToAddress(m.Address, new ScTransfers(ScColor.IOTA, part));
			}
		}
		if (parts != amount) {
			// note we truncated the calculations down to the nearest integer
			// there could be some small remainder left in the contract, but
			// that will be picked up in the next round as part of the balance
			long remainder = amount - parts;
			ctx.Log("Remainder in contract: " + remainder);
		}
	}

	public static void FuncMember(ScFuncContext ctx, FuncMemberParams params) {
		if (!ctx.Caller().equals(ctx.ContractCreator())) {
			ctx.Panic("Cancel spoofed request");
		}
		ScImmutableMap p = ctx.Params();
		ScImmutableAddress address = p.GetAddress(KeyAddress);
		if (!address.Exists()) {
			ctx.Panic("Missing address");
		}
		ScImmutableInt factor = p.GetInt(KeyFactor);
		if (!factor.Exists()) {
			ctx.Panic("Missing factor");
		}
		Member member = new Member();
		{
			member.Address = address.Value();
			member.Factor = factor.Value();
		}
		ScMutableMap state = ctx.State();
		ScMutableInt totalFactor = state.GetInt(KeyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(KeyMembers);
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = new Member(members.GetBytes(i).Value());
			if (m.Address.equals(member.Address)) {
				total -= m.Factor;
				total += member.Factor;
				totalFactor.SetValue(total);
				members.GetBytes(i).SetValue(member.toBytes());
				ctx.Log("Updated: " + member.Address);
				return;
			}
		}
		total += member.Factor;
		totalFactor.SetValue(total);
		members.GetBytes(size).SetValue(member.toBytes());
		ctx.Log("Appended: " + member.Address);
	}
}
