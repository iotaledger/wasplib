// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.contracts.donatewithfeedback.lib;

import de.mirkosertic.bytecoder.api.*;
import org.iota.wasp.contracts.donatewithfeedback.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.immutable.*;

public class DonateWithFeedbackThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc("donate", DonateWithFeedbackThunk::funcDonateThunk);
        exports.AddFunc("withdraw", DonateWithFeedbackThunk::funcWithdrawThunk);
        exports.AddView("donations", DonateWithFeedbackThunk::viewDonationsThunk);
    }

    private static void funcDonateThunk(ScFuncContext ctx) {
        var p = ctx.Params();
        var params = new FuncDonateParams();
        params.Feedback = p.GetString(Consts.ParamFeedback);
        ctx.Log("donatewithfeedback.funcDonate");
        DonateWithFeedback.funcDonate(ctx, params);
        ctx.Log("donatewithfeedback.funcDonate ok");
    }

    private static void funcWithdrawThunk(ScFuncContext ctx) {
        // only SC creator can withdraw donated funds
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncWithdrawParams();
        params.Amount = p.GetInt64(Consts.ParamAmount);
        ctx.Log("donatewithfeedback.funcWithdraw");
        DonateWithFeedback.funcWithdraw(ctx, params);
        ctx.Log("donatewithfeedback.funcWithdraw ok");
    }

    private static void viewDonationsThunk(ScViewContext ctx) {
        var params = new ViewDonationsParams();
        ctx.Log("donatewithfeedback.viewDonations");
        DonateWithFeedback.viewDonations(ctx, params);
        ctx.Log("donatewithfeedback.viewDonations ok");
    }
}