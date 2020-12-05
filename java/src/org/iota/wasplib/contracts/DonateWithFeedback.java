// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScBalances;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScLog;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class DonateWithFeedback {
	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("donate", DonateWithFeedback::donate);
		exports.AddCall("withdraw", DonateWithFeedback::withdraw);
		exports.AddView("viewDonations", DonateWithFeedback::viewDonations);
	}

	public static void donate(ScCallContext sc) {
		ScLog tlog = sc.TimestampedLog("l");
		DonationInfo donation = new DonationInfo();
		donation.seq = tlog.Length();
		donation.amount = sc.Balances().Balance(ScColor.IOTA);
		donation.donator = sc.Caller();
		donation.error = "";
		donation.feedback = sc.Params().GetString("f").Value();
		if (donation.amount == 0 || donation.feedback.length() == 0) {
			donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
			if (donation.amount > 0) {
				sc.Transfer(donation.donator, ScColor.IOTA, donation.amount);
				donation.amount = 0;
			}
		}
		byte[] bytes = encodeDonationInfo(donation);
		tlog.Append(sc.Timestamp(), bytes);

		ScMutableMap state = sc.State();
		ScMutableInt largestDonation = state.GetInt("maxd");
		ScMutableInt totalDonated = state.GetInt("total");
		if (donation.amount > largestDonation.Value()) {
			largestDonation.SetValue(donation.amount);
		}
		totalDonated.SetValue(totalDonated.Value() + donation.amount);
	}

	public static void withdraw(ScCallContext sc) {
		ScAgent scOwner = sc.Contract().Owner();
		if (!sc.From(scOwner)) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScBalances account = sc.Balances();
		long amount = account.Balance(ScColor.IOTA);
		long withdrawAmount = sc.Params().GetInt("s").Value();
		if (withdrawAmount == 0 || withdrawAmount > amount) {
			withdrawAmount = amount;
		}
		if (withdrawAmount == 0) {
			sc.Log("DonateWithFeedback: withdraw. nothing to withdraw");
			return;
		}

		sc.Transfer(scOwner, ScColor.IOTA, withdrawAmount);
	}

	public static void viewDonations(ScViewContext sc) {
		ScImmutableMap state = sc.State();
		ScImmutableInt largestDonation = state.GetInt("maxd");
		ScImmutableInt totalDonated = state.GetInt("total");
		ScImmutableMapArray tlog = sc.TimestampedLog("l");
		ScMutableMap results = sc.Results();
		results.GetInt("largest").SetValue(largestDonation.Value());
		results.GetInt("total").SetValue(totalDonated.Value());
		ScMutableMapArray donations = results.GetMapArray("donations");
		int size = tlog.Length();
		for (int i = 0; i < size; i++) {
			ScImmutableMap log = tlog.GetMap(i);
			ScMutableMap donation = donations.GetMap(i);
			donation.GetInt("timestamp").SetValue(log.GetInt("timestamp").Value());
			byte[] bytes = log.GetBytes("data").Value();
			DonationInfo di = decodeDonationInfo(bytes);
			donation.GetInt("amount").SetValue(di.amount);
			donation.GetString("feedback").SetValue(di.feedback);
			donation.GetString("donator").SetValue(di.donator.toString());
			donation.GetString("error").SetValue(di.error);
		}
	}

	public static DonationInfo decodeDonationInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		DonationInfo bet = new DonationInfo();
		bet.seq = decoder.Int();
		bet.donator = decoder.Agent();
		bet.amount = decoder.Int();
		bet.feedback = decoder.String();
		bet.error = decoder.String();
		return bet;
	}

	public static byte[] encodeDonationInfo(DonationInfo donation) {
		return new BytesEncoder().
				Int(donation.seq).
				Agent(donation.donator).
				Int(donation.amount).
				String(donation.feedback).
				String(donation.error).
				Data();
	}

	public static class DonationInfo {
		long seq;
		ScAgent donator;
		long amount;
		String feedback;
		String error;
	}
}
