// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.donatewithfeedback;

public class Donatewithfeedback {

public static void funcDonate(ScFuncContext ctx, FuncDonateParams params) {
    Donation donation = new Donation();
         {
        donation.Amount = ctx.Incoming().Balance(ScColor.IOTA);
        donation.Donator = ctx.Caller();
        donation.Error = "";
        donation.Feedback = params.Feedback.Value();
        donation.Timestamp = ctx.Timestamp();
    }
    if (donation.Amount == 0 || donation.Feedback.Len() == 0) {
        donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
        if (donation.Amount > 0) {
            ctx.TransferToAddress(donation.Donator.Address(), new ScTransfers(ScColor.IOTA, donation.Amount));
            donation.Amount = 0;
        }
    }
    state = ctx.State();
    log = state.GetBytesArray(VarLog);
    log.GetBytes(log.Length()).SetValue(donation.ToBytes());

    largestDonation = state.GetInt(VarMaxDonation);
    totalDonated = state.GetInt(VarTotalDonation);
    if (donation.Amount > largestDonation.Value()) {
        largestDonation.SetValue(donation.Amount);
    }
    totalDonated.SetValue(totalDonated.Value() + donation.Amount);
}

public static void funcWithdraw(ScFuncContext ctx, FuncWithdrawParams params) {
    balance = ctx.Balances().Balance(ScColor.IOTA);
    amount = params.Amount.Value();
    if (amount == 0 || amount > balance) {
        amount = balance;
    }
    if (amount == 0) {
        ctx.Log("DonateWithFeedback: nothing to withdraw");
        return;
    }

    scCreator = ctx.ContractCreator().Address();
    ctx.TransferToAddress(scCreator, new ScTransfers(ScColor.IOTA, amount));
}

public static void viewDonations(ScViewContext ctx, ViewDonationsParams params) {
    state = ctx.State();
    largestDonation = state.GetInt(VarMaxDonation);
    totalDonated = state.GetInt(VarTotalDonation);
    log = state.GetBytesArray(VarLog);
    results = ctx.Results();
    results.GetInt(VarMaxDonation).SetValue(largestDonation.Value());
    results.GetInt(VarTotalDonation).SetValue(totalDonated.Value());
    donations = results.GetMapArray(VarDonations);
    size = log.Length();
    for (int i = 0; i < size; i++) {
        di = Donation::fromBytes(log.GetBytes(i).Value());
        donation = donations.GetMap(i);
        donation.GetInt(VarAmount).SetValue(di.Amount);
        donation.GetString(VarDonator).SetValue(di.Donator.toString());
        donation.GetString(VarError).SetValue(di.Error);
        donation.GetString(VarFeedback).SetValue(di.Feedback);
        donation.GetInt(VarTimestamp).SetValue(di.Timestamp);
    }
}
}
