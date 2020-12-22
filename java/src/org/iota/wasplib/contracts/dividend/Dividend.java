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
	private static final Key keyAddress = new Key("address");
	private static final Key keyFactor = new Key("factor");
	private static final Key keyMembers = new Key("members");
	private static final Key keyTotalFactor = new Key("total_factor");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("member", Dividend::member);
		exports.AddCall("dividend", Dividend::dividend);
	}

	public static void member(ScCallContext sc) {
		if (!sc.From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}
		ScImmutableMap params = sc.Params();
		ScImmutableAddress address = params.GetAddress(keyAddress);
		if (!address.Exists()) {
			sc.Log("Missing address");
			return;
		}
		ScImmutableInt factor = params.GetInt(keyFactor);
		if (!factor.Exists()) {
			sc.Log("Missing factor");
			return;
		}
		Member member = new Member();
		{
			member.address = address.Value();
			member.factor = factor.Value();
		}
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt(keyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(keyMembers);
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = Member.decode(members.GetBytes(i).Value());
			if (m.address.equals(member.address)) {
				total -= m.factor;
				total += member.factor;
				totalFactor.SetValue(total);
				members.GetBytes(i).SetValue(Member.encode(member));
				sc.Log("Updated: " + member.address);
				return;
			}
		}
		total += member.factor;
		totalFactor.SetValue(total);
		members.GetBytes(size).SetValue(Member.encode(member));
		sc.Log("Appended: " + member.address);
	}

	public static void dividend(ScCallContext sc) {
		long amount = sc.Balances().Balance(ScColor.IOTA);
		if (amount == 0) {
			sc.Log("Nothing to divide");
			return;
		}
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt(keyTotalFactor);
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray(keyMembers);
		long parts = 0;
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			Member m = Member.decode(members.GetBytes(i).Value());
			long part = amount * m.factor / total;
			if (part != 0) {
				parts += part;
				sc.Transfer(m.address.AsAgent(), ScColor.IOTA, part);
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
