// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableAddress;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class Dividend {
	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("member", Dividend::member);
		exports.AddCall("divide", Dividend::divide);
	}

	public static void member(ScCallContext sc) {
		ScRequest request = sc.Request();
		if (!request.From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}
		ScImmutableMap params = request.Params();
		ScImmutableAddress address = params.GetAddress("address");
		if (!address.Exists()) {
			sc.Log("Missing address");
			return;
		}
		ScImmutableInt factor = params.GetInt("factor");
		if (!factor.Exists()) {
			sc.Log("Missing factor");
			return;
		}
		Member member = new Member();
		member.address = address.Value();
		member.factor = factor.Value();
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt("totalFactor");
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray("members");
		int size = members.Length();
		for (int i = 0; i < size; i++) {
			byte[] bytes = members.GetBytes(i).Value();
			Member m = decodeMember(bytes);
			if (m.address.equals(member.address)) {
				total -= m.factor;
				total += member.factor;
				totalFactor.SetValue(total);
				bytes = encodeMember(member);
				members.GetBytes(i).SetValue(bytes);
				sc.Log("Updated: " + member.address.toString());
				return;
			}
		}
		total += member.factor;
		totalFactor.SetValue(total);
		byte[] bytes = encodeMember(member);
		members.GetBytes(size).SetValue(bytes);
		sc.Log("Appended: " + member.address.toString());
	}

	public static void divide(ScCallContext sc) {
		long amount = sc.Account().Balance(ScColor.IOTA);
		if (amount == 0) {
			sc.Log("Nothing to divide");
			return;
		}
		ScMutableMap state = sc.State();
		ScMutableInt totalFactor = state.GetInt("totalFactor");
		long total = totalFactor.Value();
		ScMutableBytesArray members = state.GetBytesArray("members");
		int size = members.Length();
		long parts = 0;
		for (int i = 0; i < size; i++) {
			byte[] bytes = members.GetBytes(i).Value();
			Member m = decodeMember(bytes);
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
			sc.Log("Remainder in contract: " + sc.Utility().String(remainder));
		}
	}

	public static Member decodeMember(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		Member bet = new Member();
		bet.address = decoder.Address();
		bet.factor = decoder.Int();
		return bet;
	}

	public static byte[] encodeMember(Member donation) {
		return new BytesEncoder().
				Address(donation.address).
				Int(donation.factor).
				Data();
	}

	public static class Member {
		ScAddress address;
		long factor;
	}
}
