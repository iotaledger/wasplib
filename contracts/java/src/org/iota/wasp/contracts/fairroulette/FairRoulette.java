// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.fairroulette;

import org.iota.wasp.contracts.fairroulette.lib.*;
import org.iota.wasp.contracts.fairroulette.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.util.*;

public class FairRoulette {

    private static final int MaxNumber = 5;
    private static final int DefaultPlayPeriod = 120;

    public static void funcLockBets(ScFuncContext ctx, FuncLockBetsParams params) {
        // move all current bets to the locked_bets array
        var state = ctx.State();
        var bets = state.GetBytesArray(Consts.VarBets);
        var lockedBets = state.GetBytesArray(Consts.VarLockedBets);
        var nrBets = bets.Length();
        for (var i = 0; i < nrBets; i++) {
            var bytes = bets.GetBytes(i).Value();
            lockedBets.GetBytes(i).SetValue(bytes);
        }
        bets.Clear();

        ctx.PostSelf(Consts.HFuncPayWinners, null, null, 0);
    }

    public static void funcPayWinners(ScFuncContext ctx, FuncPayWinnersParams params) {
        var scId = ctx.ContractId().AsAgentId();
        var winningNumber = ctx.Utility().Random(5) + 1;
        var state = ctx.State();
        state.GetInt64(Consts.VarLastWinningNumber).SetValue(winningNumber);

        // gather all winners and calculate some totals
        var totalBetAmount = 0;
        var totalWinAmount = 0;
        var lockedBets = state.GetBytesArray(Consts.VarLockedBets);
        var winners = new ArrayList<Bet>();
        var nrBets = lockedBets.Length();
        for (var i = 0; i < nrBets; i++) {
            var bet = new Bet(lockedBets.GetBytes(i).Value());
            totalBetAmount += bet.Amount;
            if (bet.Number == winningNumber) {
                totalWinAmount += bet.Amount;
                winners.add(bet);
            }
        }
        lockedBets.Clear();

        if (winners.isEmpty()) {
            ctx.Log("Nobody wins!");
            // compact separate bet deposit UTXOs into a single one
            ctx.TransferToAddress(scId.Address(), new ScTransfers(ScColor.IOTA, totalBetAmount));
            return;
        }

        // pay out the winners proportionally to their bet amount
        var totalPayout = 0;
        var size = winners.size();
        for (var i = 0; i < size; i++) {
            var bet = winners.get(i);
            var payout = totalBetAmount * bet.Amount / totalWinAmount;
            if (payout != 0) {
                totalPayout += payout;
                ctx.TransferToAddress(bet.Better.Address(), new ScTransfers(ScColor.IOTA, payout));
            }
            var text = "Pay " + payout +
                    " to " + bet.Better;
            ctx.Log(text);
        }

        // any truncation left-overs are fair picking for the smart contract
        if (totalPayout != totalBetAmount) {
            var remainder = totalBetAmount - totalPayout;
            var text = "Remainder is " + remainder;
            ctx.Log(text);
            ctx.TransferToAddress(scId.Address(), new ScTransfers(ScColor.IOTA, remainder));
        }
    }

    public static void funcPlaceBet(ScFuncContext ctx, FuncPlaceBetParams params) {
        var amount = ctx.Incoming().Balance(ScColor.IOTA);
        if (amount == 0) {
            ctx.Panic("Empty bet...");
        }
        var number = params.Number.Value();
        if (number < 1 || number > MaxNumber) {
            ctx.Panic("Invalid number...");
        }

        var bet = new Bet();
        {
            bet.Better = ctx.Caller();
            bet.Amount = amount;
            bet.Number = number;
        }

        var state = ctx.State();
        var bets = state.GetBytesArray(Consts.VarBets);
        var betNr = bets.Length();
        bets.GetBytes(betNr).SetValue(bet.toBytes());
        if (betNr == 0) {
            var playPeriod = state.GetInt64(Consts.VarPlayPeriod).Value();
            if (playPeriod < 10) {
                playPeriod = DefaultPlayPeriod;
            }
            ctx.PostSelf(Consts.HFuncLockBets, null, null, playPeriod);
        }
    }

    public static void funcPlayPeriod(ScFuncContext ctx, FuncPlayPeriodParams params) {
        var playPeriod = params.PlayPeriod.Value();
        ctx.Require(playPeriod >= 10, "invalid play period");
        ctx.State().GetInt64(Consts.VarPlayPeriod).SetValue(playPeriod);
    }

    public static void viewLastWinningNumber(ScViewContext ctx, ViewLastWinningNumberParams params) {
        // Create an ScImmutableMap proxy to the state storage map on the host.
        var state = ctx.State();

        // Get the 'lastWinningNumber' int64 value from state storage through
        // an ScImmutableInt64 proxy.
        var lastWinningNumber = state.GetInt64(Consts.VarLastWinningNumber).Value();

        // Create an ScMutableMap proxy to the map on the host that will store the
        // key/value pairs that we want to return from this View function
        var results = ctx.Results();

        // Set the value associated with the 'lastWinningNumber' key to the value
        // we got from state storage
        results.GetInt64(Consts.VarLastWinningNumber).SetValue(lastWinningNumber);

    }
}
