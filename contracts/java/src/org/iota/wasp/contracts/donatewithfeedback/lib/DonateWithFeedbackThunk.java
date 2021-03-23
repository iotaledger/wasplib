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
import org.iota.wasp.wasmlib.keys.*;

public class DonateWithFeedbackThunk {
    public static void main(String[] args) {
    }

    @Export("on_load")
    public static void onLoad() {
        ScExports exports = new ScExports();
        exports.AddFunc(Consts.FuncDonate, DonateWithFeedbackThunk::funcDonateThunk);
        exports.AddFunc(Consts.FuncWithdraw, DonateWithFeedbackThunk::funcWithdrawThunk);
        exports.AddView(Consts.ViewDonations, DonateWithFeedbackThunk::viewDonationsThunk);
    }

    private static void funcDonateThunk(ScFuncContext ctx) {
        ctx.Log("donatewithfeedback.funcDonate");
        var p = ctx.Params();
        var params = new FuncDonateParams();
        params.Feedback = p.GetString(Consts.ParamFeedback);
        DonateWithFeedback.funcDonate(ctx, params);
        ctx.Log("donatewithfeedback.funcDonate ok");
    }

    private static void funcWithdrawThunk(ScFuncContext ctx) {
        ctx.Log("donatewithfeedback.funcWithdraw");
        // only SC creator can withdraw donated funds
        ctx.Require(ctx.Caller().equals(ctx.ContractCreator()), "no permission");

        var p = ctx.Params();
        var params = new FuncWithdrawParams();
        params.Amount = p.GetInt64(Consts.ParamAmount);
        DonateWithFeedback.funcWithdraw(ctx, params);
        ctx.Log("donatewithfeedback.funcWithdraw ok");
    }

    private static void viewDonationsThunk(ScViewContext ctx) {
        ctx.Log("donatewithfeedback.viewDonations");
        var params = new ViewDonationsParams();
        DonateWithFeedback.viewDonations(ctx, params);
        ctx.Log("donatewithfeedback.viewDonations ok");
    }
}
