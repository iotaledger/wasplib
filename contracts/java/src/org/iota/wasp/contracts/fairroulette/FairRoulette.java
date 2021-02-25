// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.fairroulette;

import org.iota.wasp.contracts.fairroulette.lib.*;
import org.iota.wasp.contracts.fairroulette.types.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.util.*;

public class FairRoulette {

    private static final int MaxNumber = 5;
    private static final int DefaultPlayPeriod = 120;

    public static void funcLockBets(ScFuncContext ctx, FuncLockBetsParams params) {
        // move all current bets to the locked_bets array
        ScMutableMap state = ctx.State();
        ScMutableBytesArray bets = state.GetBytesArray(Consts.VarBets);
        ScMutableBytesArray lockedBets = state.GetBytesArray(Consts.VarLockedBets);
        int nrBets = bets.Length();
        for (int i = 0; i < nrBets; i++) {
            byte[] bytes = bets.GetBytes(i).Value();
            lockedBets.GetBytes(i).SetValue(bytes);
        }
        bets.Clear();
        ctx.PostSelf(Consts.HFuncPayWinners, null, null, 0);
    }

    public static void funcPayWinners(ScFuncContext ctx, FuncPayWinnersParams params) {
        ScAgentId scId = ctx.ContractId().AsAgentId();
        long winningNumber = ctx.Utility().Random(5) + 1;
        ScMutableMap state = ctx.State();
        state.GetInt64(Consts.VarLastWinningNumber).SetValue(winningNumber);

        // gather all winners and calculate some totals
        long totalBetAmount = 0;
        long totalWinAmount = 0;
        ScMutableBytesArray lockedBets = state.GetBytesArray(Consts.VarLockedBets);
        ArrayList<Bet> winners = new ArrayList<Bet>();
        int nrBets = lockedBets.Length();
        Bet bet;
        for (int i = 0; i < nrBets; i++) {
            bet = new Bet(lockedBets.GetBytes(i).Value());
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
        long totalPayout = 0;
        int size = winners.size();
        String text;
        for (int i = 0; i < size; i++) {
            bet = winners.get(i);
            long payout = totalBetAmount * bet.Amount / totalWinAmount;
            if (payout != 0) {
                totalPayout += payout;
                ctx.TransferToAddress(bet.Better.Address(), new ScTransfers(ScColor.IOTA, payout));
            }
            text = "Pay " + payout +
                    " to " + bet.Better;
            ctx.Log(text);
        }

        // any truncation left-overs are fair picking for the smart contract
        if (totalPayout != totalBetAmount) {
            long remainder = totalBetAmount - totalPayout;
            text = "Remainder is " + remainder;
            ctx.Log(text);
            ctx.TransferToAddress(scId.Address(), new ScTransfers(ScColor.IOTA, remainder));
        }
    }

    public static void funcPlaceBet(ScFuncContext ctx, FuncPlaceBetParams params) {
        long amount = ctx.Incoming().Balance(ScColor.IOTA);
        if (amount == 0) {
            ctx.Panic("Empty bet...");
        }
        long number = params.Number.Value();
        if (number < 1 || number > MaxNumber) {
            ctx.Panic("Invalid number...");
        }

        Bet bet = new Bet();
        {
            bet.Better = ctx.Caller();
            bet.Amount = amount;
            bet.Number = number;
        }

        ScMutableMap state = ctx.State();
        ScMutableBytesArray bets = state.GetBytesArray(Consts.VarBets);
        int betNr = bets.Length();
        bets.GetBytes(betNr).SetValue(bet.toBytes());
        if (betNr == 0) {
            long playPeriod = state.GetInt64(Consts.VarPlayPeriod).Value();
            if (playPeriod < 10) {
                playPeriod = DefaultPlayPeriod;
            }
            ctx.PostSelf(Consts.HFuncLockBets, null, null, playPeriod);
        }
    }

    public static void funcPlayPeriod(ScFuncContext ctx, FuncPlayPeriodParams params) {
        long playPeriod = params.PlayPeriod.Value();
        if (playPeriod < 10) {
            ctx.Panic("Invalid play period...");
        }

        ctx.State().GetInt64(Consts.VarPlayPeriod).SetValue(playPeriod);
    }
}
