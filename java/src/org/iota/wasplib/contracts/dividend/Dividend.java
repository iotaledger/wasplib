// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.dividend;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableAddress;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class Dividend {
	private static final Key KeyAddress = new Key("address");
	private static final Key KeyFactor = new Key("factor");
	private static final Key KeyMembers = new Key("members");
	private static final Key KeyTotalFactor = new Key("total_factor");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("member", Dividend::member);
		exports.AddCall("dividend", Dividend::dividend);
	}

	public static void member(ScCallContext sc) {
		if (!sc.From(sc.ContractCreator())) {
			sc.Panic("Cancel spoofed request");
		}
		ScImmutableMap params = sc.Params();
		ScImmutableAddress address = params.GetAddress(KeyAddress);
		if (!address.Exists()) {
			sc.Panic("Missing address");
		}
		ScImmutableInt factor = params.GetInt(KeyFactor);
		if (!factor.Exists()) {
			sc.Panic("Missing factor");
		}
		Member member = new Member();
		{
			member.Address = address.Value();
			member.Factor = factor.Value();
		}
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt(KeyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(KeyMembers);
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = Member.decode(members.GetBytes(i).Value());
			if (m.Address.equals(member.Address)) {
				total -= m.Factor;
				total += member.Factor;
				totalFactor.SetValue(total);
				members.GetBytes(i).SetValue(Member.encode(member));
				sc.Log("Updated: " + member.Address);
				return;
			}
		}
		total += member.Factor;
		totalFactor.SetValue(total);
		members.GetBytes(size).SetValue(Member.encode(member));
		sc.Log("Appended: " + member.Address);
	}

	public static void dividend(ScCallContext sc) {
		long amount = sc.Balances().Balance(ScColor.IOTA);
		if (amount == 0) {
			sc.Panic("Nothing to divide");
		}
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt(KeyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(KeyMembers);
		long parts = 0;
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = Member.decode(members.GetBytes(i).Value());
			long part = amount * m.Factor / total;
			if (part != 0) {
				parts += part;
				sc.Transfer(m.Address.AsAgent(), ScColor.IOTA, part);
			}
		}
		if (parts != amount) {
			// note we truncated the calculations down to the nearest integer
			// there could be some small remainder left in the contract, but
			// that will be picked up in the next round as part of the balance
			long remainder = amount - parts;
			sc.Log("Remainder in contract: " + remainder);
		}
	}
}
