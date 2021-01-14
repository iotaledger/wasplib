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
	private static final Key KeyAmount = new Key("amount");
	private static final Key KeyDonations = new Key("donations");
	private static final Key KeyDonator = new Key("donator");
	private static final Key KeyError = new Key("error");
	private static final Key KeyFeedback = new Key("feedback");
	private static final Key KeyLog = new Key("log");
	private static final Key KeyMaxDonation = new Key("max_donation");
	private static final Key KeyTimestamp = new Key("timestamp");
	private static final Key KeyTotalDonation = new Key("total_donation");
	private static final Key KeyWithdrawAmount = new Key("withdraw");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("donate", DonateWithFeedback::donate);
		exports.AddCall("withdraw", DonateWithFeedback::withdraw);
		exports.AddView("view_donations", DonateWithFeedback::viewDonations);
	}

	public static void donate(ScCallContext sc) {
		DonationInfo donation = new DonationInfo();
		{
			donation.Amount = sc.Incoming().Balance(ScColor.IOTA);
			donation.Donator = sc.Caller();
			donation.Error = "";
			donation.Feedback = sc.Params().GetString(KeyFeedback).Value();
			donation.Timestamp = sc.Timestamp();
		}
		if (donation.Amount == 0 || donation.Feedback.length() == 0) {
			donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
			if (donation.Amount > 0) {
				sc.Transfer(donation.Donator, ScColor.IOTA, donation.Amount);
				donation.Amount = 0;
			}
		}
		ScMutableMap state = sc.State();
		ScMutableBytesArray log = state.GetBytesArray(KeyLog);
		log.GetBytes(log.Length()).SetValue(DonationInfo.encode(donation));

		ScMutableInt largestDonation = state.GetInt(KeyMaxDonation);
		ScMutableInt totalDonated = state.GetInt(KeyTotalDonation);
		if (donation.Amount > largestDonation.Value()) {
			largestDonation.SetValue(donation.Amount);
		}
		totalDonated.SetValue(totalDonated.Value() + donation.Amount);
	}

	public static void withdraw(ScCallContext sc) {
		ScAgent scOwner = sc.Contract().Creator();
		if (!sc.From(scOwner)) {
			sc.Panic("Cancel spoofed request");
		}

		long amount = sc.Balances().Balance(ScColor.IOTA);
		long withdrawAmount = sc.Params().GetInt(KeyWithdrawAmount).Value();
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
		ScImmutableInt largestDonation = state.GetInt(KeyMaxDonation);
		ScImmutableInt totalDonated = state.GetInt(KeyTotalDonation);
		ScImmutableBytesArray log = state.GetBytesArray(KeyLog);
		ScMutableMap results = sc.Results();
		results.GetInt(KeyMaxDonation).SetValue(largestDonation.Value());
		results.GetInt(KeyTotalDonation).SetValue(totalDonated.Value());
		ScMutableMapArray donations = results.GetMapArray(KeyDonations);
		int size = log.Length();
		for (int i = 0; i < size; i++) {
			DonationInfo di = DonationInfo.decode(log.GetBytes(i).Value());
			ScMutableMap donation = donations.GetMap(i);
			donation.GetInt(KeyAmount).SetValue(di.Amount);
			donation.GetString(KeyDonator).SetValue(di.Donator.toString());
			donation.GetString(KeyError).SetValue(di.Error);
			donation.GetString(KeyFeedback).SetValue(di.Feedback);
			donation.GetInt(KeyTimestamp).SetValue(di.Timestamp);
		}
	}
}
