// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.donatewithfeedback;

import org.iota.wasp.contracts.donatewithfeedback.lib.*;
import org.iota.wasp.contracts.donatewithfeedback.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class DonateWithFeedback {

    public static void funcDonate(ScFuncContext ctx, FuncDonateParams params) {
        var donation = new Donation();
        {
            donation.Amount = ctx.Incoming().Balance(ScColor.IOTA);
            donation.Donator = ctx.Caller();
            donation.Error = "";
            donation.Feedback = params.Feedback.Value();
            donation.Timestamp = ctx.Timestamp();
        }
        if (donation.Amount == 0 || donation.Feedback.length() == 0) {
            donation.Error = "error: empty feedback or donated amount = 0";
            if (donation.Amount > 0) {
                ctx.TransferToAddress(donation.Donator.Address(), ScTransfers.iotas(donation.Amount));
                donation.Amount = 0;
            }
        }
        var state = ctx.State();
        var log = state.GetBytesArray(Consts.VarLog);
        log.GetBytes(log.Length()).SetValue(donation.toBytes());

        var largestDonation = state.GetInt64(Consts.VarMaxDonation);
        var totalDonated = state.GetInt64(Consts.VarTotalDonation);
        if (donation.Amount > largestDonation.Value()) {
            largestDonation.SetValue(donation.Amount);
        }
        totalDonated.SetValue(totalDonated.Value() + donation.Amount);
    }

    public static void funcWithdraw(ScFuncContext ctx, FuncWithdrawParams params) {
        var balance = ctx.Balances().Balance(ScColor.IOTA);
        var amount = params.Amount.Value();
        if (amount == 0 || amount > balance) {
            amount = balance;
        }
        if (amount == 0) {
            ctx.Log("dwf.withdraw: nothing to withdraw");
            return;
        }

        var scCreator = ctx.ContractCreator().Address();
        ctx.TransferToAddress(scCreator, ScTransfers.iotas(amount));
    }

    public static void viewDonations(ScViewContext ctx, ViewDonationsParams params) {
        var state = ctx.State();
        var largestDonation = state.GetInt64(Consts.VarMaxDonation);
        var totalDonated = state.GetInt64(Consts.VarTotalDonation);
        var log = state.GetBytesArray(Consts.VarLog);
        var results = ctx.Results();
        results.GetInt64(Consts.VarMaxDonation).SetValue(largestDonation.Value());
        results.GetInt64(Consts.VarTotalDonation).SetValue(totalDonated.Value());
        var donations = results.GetMapArray(Consts.VarDonations);
        var size = log.Length();
        for (var i = 0; i < size; i++) {
            var di = new Donation(log.GetBytes(i).Value());
            var donation = donations.GetMap(i);
            donation.GetInt64(Consts.VarAmount).SetValue(di.Amount);
            donation.GetString(Consts.VarDonator).SetValue(di.Donator.toString());
            donation.GetString(Consts.VarError).SetValue(di.Error);
            donation.GetString(Consts.VarFeedback).SetValue(di.Feedback);
            donation.GetInt64(Consts.VarTimestamp).SetValue(di.Timestamp);
        }
    }
}
