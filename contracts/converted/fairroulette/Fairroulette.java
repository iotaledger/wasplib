// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts.fairroulette;

public class Fairroulette {

private static final int MaxNumber = 5;
private static final int DefaultPlayPeriod = 120;

public static void funcLockBets(ScFuncContext ctx, FuncLockBetsParams params) {
    // move all current bets to the locked_bets array
    state = ctx.State();
    bets = state.GetBytesArray(VarBets);
    lockedBets = state.GetBytesArray(VarLockedBets);
    nrBets = bets.Length();
    for (int i = 0; i < nrBets; i++) {
        bytes = bets.GetBytes(i).Value();
        lockedBets.GetBytes(i).SetValue(bytes);
    }
    bets.Clear();

    ctx.Post(PostRequestParams {
        auction.ContractId = ctx.ContractId();
        auction.Function = HFuncPayWinners;
        auction.Params = null;
        auction.Transfer = null;
        auction.Delay = 0;
    });
}

public static void funcPayWinners(ScFuncContext ctx, FuncPayWinnersParams params) {
    scId = ctx.ContractId().AsAgentId();
    winningNumber = ctx.Utility().Random(5) + 1;
    state = ctx.State();
    state.GetInt(VarLastWinningNumber).SetValue(winningNumber);

    // gather all winners and calculate some totals
    totalBetAmount = 0;
    totalWinAmount = 0;
    lockedBets = state.GetBytesArray(VarLockedBets);
    let mut winners: Vec<Bet> = Vec::new();
    nrBets = lockedBets.Length();
    for (int i = 0; i < nrBets; i++) {
        bet = Bet::fromBytes(lockedBets.GetBytes(i).Value());
        totalBetAmount += bet.Amount;
        if (bet.Number == winningNumber) {
            totalWinAmount += bet.Amount;
            winners.Push(bet);
        }
    }
    lockedBets.Clear();

    if (winners.IsEmpty()) {
        ctx.Log("Nobody wins!");
        // compact separate bet deposit UTXOs into a single one
        ctx.TransferToAddress(scId.Address(), new ScTransfers(ScColor.IOTA, totalBetAmount));
        return;
    }

    // pay out the winners proportionally to their bet amount
    totalPayout = 0;
    size = winners.Len();
    for (int i = 0; i < size; i++) {
        bet = &winners[i];
        payout = totalBetAmount * bet.Amount / totalWinAmount;
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
        remainder = totalBetAmount - totalPayout;
        text = "Remainder is " + remainder;
        ctx.Log(text);
        ctx.TransferToAddress(scId.Address(), new ScTransfers(ScColor.IOTA, remainder));
    }
}

public static void funcPlaceBet(ScFuncContext ctx, FuncPlaceBetParams params) {
    amount = ctx.Incoming().Balance(ScColor.IOTA);
    if (amount == 0) {
        ctx.Panic("Empty bet...");
    }
    number = params.Number.Value();
    if (number < 1 || number > MaxNumber) {
        ctx.Panic("Invalid number...");
    }

    Bet bet = new Bet();
         {
        bet.Better = ctx.Caller();
        bet.Amount = amount;
        bet.Number = number;
    }

    state = ctx.State();
    bets = state.GetBytesArray(VarBets);
    betNr = bets.Length();
    bets.GetBytes(betNr).SetValue(bet.ToBytes());
    if (betNr == 0) {
        playPeriod = state.GetInt(VarPlayPeriod).Value();
        if (playPeriod < 10) {
            playPeriod = DefaultPlayPeriod;
        }
        ctx.Post(PostRequestParams {
            bet.ContractId = ctx.ContractId();
            bet.Function = HFuncLockBets;
            bet.Params = null;
            bet.Transfer = null;
            bet.Delay = playPeriod;
        });
    }
}

public static void funcPlayPeriod(ScFuncContext ctx, FuncPlayPeriodParams params) {
    playPeriod = params.PlayPeriod.Value();
    if (playPeriod < 10) {
        ctx.Panic("Invalid play period...");
    }

    ctx.State().GetInt(VarPlayPeriod).SetValue(playPeriod);
}
}
