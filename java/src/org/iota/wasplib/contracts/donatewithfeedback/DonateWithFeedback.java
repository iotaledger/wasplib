// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.donatewithfeedback;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableBytesArray;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class DonateWithFeedback {
	private static final Key keyAmount = new Key("amount");
	private static final Key keyDonations = new Key("donations");
	private static final Key keyDonator = new Key("donator");
	private static final Key keyError = new Key("error");
	private static final Key keyFeedback = new Key("feedback");
	private static final Key keyLog = new Key("log");
	private static final Key keyMaxDonation = new Key("max_donation");
	private static final Key keyTimestamp = new Key("timestamp");
	private static final Key keyTotalDonation = new Key("total_donation");
	private static final Key keyWithdrawAmount = new Key("withdraw");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("donate", DonateWithFeedback::donate);
		exports.AddCall("withdraw", DonateWithFeedback::withdraw);
		exports.AddView("view_donations", DonateWithFeedback::viewDonations);
	}

	public static void donate(ScCallContext sc) {
		DonationInfo donation = new DonationInfo();
		{
			donation.amount = sc.Incoming().Balance(ScColor.IOTA);
			donation.donator = sc.Caller();
			donation.error = "";
			donation.feedback = sc.Params().GetString(keyFeedback).Value();
			donation.timestamp = sc.Timestamp();
		}
		if (donation.amount == 0 || donation.feedback.length() == 0) {
			donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
			if (donation.amount > 0) {
				sc.Transfer(donation.donator, ScColor.IOTA, donation.amount);
				donation.amount = 0;
			}
		}
		ScMutableMap state = sc.State();
		ScMutableBytesArray log = state.GetBytesArray(keyLog);
		log.GetBytes(log.Length()).SetValue(DonationInfo.encode(donation));

		ScMutableInt largestDonation = state.GetInt(keyMaxDonation);
		ScMutableInt totalDonated = state.GetInt(keyTotalDonation);
		if (donation.amount > largestDonation.Value()) {
			largestDonation.SetValue(donation.amount);
		}
		totalDonated.SetValue(totalDonated.Value() + donation.amount);
	}

	public static void withdraw(ScCallContext sc) {
		ScAgent scOwner = sc.Contract().Creator();
		if (!sc.From(scOwner)) {
			sc.Panic("Cancel spoofed request");
		}

		long amount = sc.Balances().Balance(ScColor.IOTA);
		long withdrawAmount = sc.Params().GetInt(keyWithdrawAmount).Value();
		if (withdrawAmount == 0 || withdrawAmount > amount) {
			withdrawAmount = amount;
		}
		if (withdrawAmount == 0) {
			sc.Panic("DonateWithFeedback: nothing to withdraw");
		}

		sc.Transfer(scOwner, ScColor.IOTA, withdrawAmount);
	}

	public static void viewDonations(ScViewContext sc) {
		ScImmutableMap state = sc.State();
		ScImmutableInt largestDonation = state.GetInt(keyMaxDonation);
		ScImmutableInt totalDonated = state.GetInt(keyTotalDonation);
		ScImmutableBytesArray log = state.GetBytesArray(keyLog);
		ScMutableMap results = sc.Results();
		results.GetInt(keyMaxDonation).SetValue(largestDonation.Value());
		results.GetInt(keyTotalDonation).SetValue(totalDonated.Value());
		ScMutableMapArray donations = results.GetMapArray(keyDonations);
		int size = log.Length();
		for (int i = 0; i < size; i++) {
			DonationInfo di = DonationInfo.decode(log.GetBytes(i).Value());
			ScMutableMap donation = donations.GetMap(i);
			donation.GetInt(keyAmount).SetValue(di.amount);
			donation.GetString(keyDonator).SetValue(di.donator.toString());
			donation.GetString(keyError).SetValue(di.error);
			donation.GetString(keyFeedback).SetValue(di.feedback);
			donation.GetInt(keyTimestamp).SetValue(di.timestamp);
		}
	}
}
