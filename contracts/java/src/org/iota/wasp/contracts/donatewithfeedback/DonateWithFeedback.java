// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.donatewithfeedback;

import org.iota.wasp.contracts.donatewithfeedback.lib.*;
import org.iota.wasp.contracts.donatewithfeedback.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class DonateWithFeedback {

	public static void funcDonate(ScFuncContext ctx, FuncDonateParams params) {
		Donation donation = new Donation();
		{
			donation.Amount = ctx.Incoming().Balance(ScColor.IOTA);
			donation.Donator = ctx.Caller();
			donation.Error = "";
			donation.Feedback = params.Feedback.Value();
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
		ScMutableBytesArray log = state.GetBytesArray(Consts.VarLog);
		log.GetBytes(log.Length()).SetValue(donation.toBytes());

		ScMutableInt largestDonation = state.GetInt(Consts.VarMaxDonation);
		ScMutableInt totalDonated = state.GetInt(Consts.VarTotalDonation);
		if (donation.Amount > largestDonation.Value()) {
			largestDonation.SetValue(donation.Amount);
		}
		totalDonated.SetValue(totalDonated.Value() + donation.Amount);
	}

	public static void funcWithdraw(ScFuncContext ctx, FuncWithdrawParams params) {
		long balance = ctx.Balances().Balance(ScColor.IOTA);
		long amount = params.Amount.Value();
		if (amount == 0 || amount > balance) {
			amount = balance;
		}
		if (amount == 0) {
			ctx.Log("DonateWithFeedback: nothing to withdraw");
			return;
		}

		ScAddress scCreator = ctx.ContractCreator().Address();
		ctx.TransferToAddress(scCreator, new ScTransfers(ScColor.IOTA, amount));
	}

	public static void viewDonations(ScViewContext ctx, ViewDonationsParams params) {
		ScImmutableMap state = ctx.State();
		ScImmutableInt largestDonation = state.GetInt(Consts.VarMaxDonation);
		ScImmutableInt totalDonated = state.GetInt(Consts.VarTotalDonation);
		ScImmutableBytesArray log = state.GetBytesArray(Consts.VarLog);
		ScMutableMap results = ctx.Results();
		results.GetInt(Consts.VarMaxDonation).SetValue(largestDonation.Value());
		results.GetInt(Consts.VarTotalDonation).SetValue(totalDonated.Value());
		ScMutableMapArray donations = results.GetMapArray(Consts.VarDonations);
		int size = log.Length();
		for (int i = 0; i < size; i++) {
			Donation di = new Donation(log.GetBytes(i).Value());
			ScMutableMap donation = donations.GetMap(i);
			donation.GetInt(Consts.VarAmount).SetValue(di.Amount);
			donation.GetString(Consts.VarDonator).SetValue(di.Donator.toString());
			donation.GetString(Consts.VarError).SetValue(di.Error);
			donation.GetString(Consts.VarFeedback).SetValue(di.Feedback);
			donation.GetInt(Consts.VarTimestamp).SetValue(di.Timestamp);
		}
	}
}
