// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.contracts.fairroulette;

import org.iota.wasp.contracts.fairroulette.lib.FuncLockBetsParams;
import org.iota.wasp.contracts.fairroulette.lib.FuncPayWinnersParams;
import org.iota.wasp.contracts.fairroulette.lib.FuncPlaceBetParams;
import org.iota.wasp.contracts.fairroulette.lib.FuncPlayPeriodParams;
import org.iota.wasp.contracts.fairroulette.types.Bet;
import org.iota.wasp.wasmlib.context.ScFuncContext;
import org.iota.wasp.wasmlib.hashtypes.ScAgentId;
import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.mutable.ScMutableBytesArray;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

import java.util.ArrayList;

public class FairRoulette {
	private static final Key KeyBets = new Key("bets");
	private static final Key KeyColor = new Key("color");
	private static final Key KeyLastWinningColor = new Key("last_winning_color");
	private static final Key KeyLockedBets = new Key("locked_bets");
	private static final Key KeyPlayPeriod = new Key("play_period");

	private static final int numColors = 5;
	private static final int defaultPlayPeriod = 120;

	public static void FuncLockBets(ScFuncContext ctx, FuncLockBetsParams params) {
		// can only be sent by SC itself
		if (!ctx.Caller().equals(ctx.ContractId().AsAgentId())) {
			ctx.Panic("Cancel spoofed request");
		}

		// move all current bets to the locked_bets array
		ScMutableMap state = ctx.State();
		ScMutableBytesArray bets = state.GetBytesArray(KeyBets);
		ScMutableBytesArray lockedBets = state.GetBytesArray(KeyLockedBets);
		int nrBets = bets.Length();
		for (int i = 0; i < nrBets; i++) {
			byte[] bytes = bets.GetBytes(i).Value();
			lockedBets.GetBytes(i).SetValue(bytes);
		}
		bets.Clear();

		ctx.Post("pay_winners").Post(0);
	}

	public static void FuncPayWinners(ScFuncContext ctx, FuncPayWinnersParams params) {
		// can only be sent by SC itself
		ScAgentId scId = ctx.ContractId().AsAgentId();
		if (!ctx.Caller().equals(scId)) {
			ctx.Panic("Cancel spoofed request");
		}

		long winningColor = ctx.Utility().Random(5) + 1;
		ScMutableMap state = ctx.State();
		state.GetInt(KeyLastWinningColor).SetValue(winningColor);

		// gather all winners and calculate some totals
		long totalBetAmount = 0;
		long totalWinAmount = 0;
		ScMutableBytesArray lockedBets = state.GetBytesArray(KeyLockedBets);
		ArrayList<Bet> winners = new ArrayList<Bet>();
		int nrBets = lockedBets.Length();
		for (int i = 0; i < nrBets; i++) {
			Bet bet = new Bet(lockedBets.GetBytes(i).Value());
			totalBetAmount += bet.Amount;
			if (bet.Number == winningColor) {
				totalWinAmount += bet.Amount;
				winners.add(bet);
			}
		}
		lockedBets.Clear();

		if (winners.size() == 0) {
			ctx.Log("Nobody wins!");
			// compact separate UTXOs into a single one
			ctx.Transfer(scId, ScColor.IOTA, totalBetAmount);
			return;
		}

		// pay out the winners proportionally to their bet amount
		int totalPayout = 0;
		int size = winners.size();
		String text;
		for (int i = 0; i < size; i++) {
			Bet bet = winners.get(i);
			long payout = totalBetAmount * bet.Amount / totalWinAmount;
			if (payout != 0) {
				totalPayout += payout;
				ctx.Transfer(bet.Better, ScColor.IOTA, payout);
			}
			text = "Pay " + payout + " to " + bet.Better;
			ctx.Log(text);
		}

		// any truncation left-overs are fair picking for the smart contract
		if (totalPayout != totalBetAmount) {
			long remainder = totalBetAmount - totalPayout;
			text = "Remainder is " + remainder;
			ctx.Log(text);
			ctx.Transfer(scId, ScColor.IOTA, remainder);
		}
	}

	public static void FuncPlaceBet(ScFuncContext ctx, FuncPlaceBetParams params) {
		long amount = ctx.Incoming().Balance(ScColor.IOTA);
		if (amount == 0) {
			ctx.Panic("Empty bet...");
		}
		long color = ctx.Params().GetInt(KeyColor).Value();
		if (color == 0) {
			ctx.Panic("No color...");
		}
		if (color < 1 || color > numColors) {
			ctx.Panic("Invalid color...");
		}

		Bet bet = new Bet();
		{
			bet.Better = ctx.Caller();
			bet.Amount = amount;
			bet.Number = color;
		}

		ScMutableMap state = ctx.State();
		ScMutableBytesArray bets = state.GetBytesArray(KeyBets);
		int betNr = bets.Length();
		bets.GetBytes(betNr).SetValue(bet.toBytes());
		if (betNr == 0) {
			long playPeriod = state.GetInt(KeyPlayPeriod).Value();
			if (playPeriod < 10) {
				playPeriod = defaultPlayPeriod;
			}
			ctx.Post("lock_bets").Post(playPeriod);
		}
	}

	public static void FuncPlayPeriod(ScFuncContext ctx, FuncPlayPeriodParams params) {
		// can only be sent by SC creator
		if (!ctx.Caller().equals(ctx.ContractCreator())) {
			ctx.Panic("Cancel spoofed request");
		}

		long playPeriod = ctx.Params().GetInt(KeyPlayPeriod).Value();
		if (playPeriod < 10) {
			ctx.Panic("Invalid play period...");
		}

		ctx.State().GetInt(KeyPlayPeriod).SetValue(playPeriod);
	}
}
