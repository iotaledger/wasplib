// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.donatewithfeedback;

import org.iota.wasp.contracts.donatewithfeedback.lib.*;
import org.iota.wasp.contracts.donatewithfeedback.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

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

	public static void FuncDonate(ScFuncContext ctx, FuncDonateParams params) {
		Donation donation = new Donation();
		{
			donation.Amount = ctx.Incoming().Balance(ScColor.IOTA);
			donation.Donator = ctx.Caller();
			donation.Error = "";
			donation.Feedback = ctx.Params().GetString(KeyFeedback).Value();
			donation.Timestamp = ctx.Timestamp();
		}
		if (donation.Amount == 0 || donation.Feedback.length() == 0) {
			donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
			if (donation.Amount > 0) {
				ctx.TransferToAddress(donation.Donator.Address(), new ScTransfers(ScColor.IOTA, donation.Amount));
				donation.Amount = 0;
			}
		}
		ScMutableMap state = ctx.State();
		ScMutableBytesArray log = state.GetBytesArray(KeyLog);
		log.GetBytes(log.Length()).SetValue(donation.toBytes());

		ScMutableInt largestDonation = state.GetInt(KeyMaxDonation);
		ScMutableInt totalDonated = state.GetInt(KeyTotalDonation);
		if (donation.Amount > largestDonation.Value()) {
			largestDonation.SetValue(donation.Amount);
		}
		totalDonated.SetValue(totalDonated.Value() + donation.Amount);
	}

	public static void FuncWithdraw(ScFuncContext ctx, FuncWithdrawParams params) {
		ScAgentId scOwner = ctx.ContractCreator();
		if (!ctx.Caller().equals(scOwner)) {
			ctx.Panic("Cancel spoofed request");
		}

		long amount = ctx.Balances().Balance(ScColor.IOTA);
		long withdrawAmount = ctx.Params().GetInt(KeyWithdrawAmount).Value();
		if (withdrawAmount == 0 || withdrawAmount > amount) {
			withdrawAmount = amount;
		}
		if (withdrawAmount == 0) {
			ctx.Panic("DonateWithFeedback: nothing to withdraw");
		}

		ctx.TransferToAddress(scOwner.Address(), new ScTransfers(ScColor.IOTA, withdrawAmount));
	}

	public static void ViewDonations(ScViewContext ctx, ViewDonationsParams params) {
		ScImmutableMap state = ctx.State();
		ScImmutableInt largestDonation = state.GetInt(KeyMaxDonation);
		ScImmutableInt totalDonated = state.GetInt(KeyTotalDonation);
		ScImmutableBytesArray log = state.GetBytesArray(KeyLog);
		ScMutableMap results = ctx.Results();
		results.GetInt(KeyMaxDonation).SetValue(largestDonation.Value());
		results.GetInt(KeyTotalDonation).SetValue(totalDonated.Value());
		ScMutableMapArray donations = results.GetMapArray(KeyDonations);
		int size = log.Length();
		for (int i = 0; i < size; i++) {
			Donation di = new Donation(log.GetBytes(i).Value());
			ScMutableMap donation = donations.GetMap(i);
			donation.GetInt(KeyAmount).SetValue(di.Amount);
			donation.GetString(KeyDonator).SetValue(di.Donator.toString());
			donation.GetString(KeyError).SetValue(di.Error);
			donation.GetString(KeyFeedback).SetValue(di.Feedback);
			donation.GetInt(KeyTimestamp).SetValue(di.Timestamp);
		}
	}
}
