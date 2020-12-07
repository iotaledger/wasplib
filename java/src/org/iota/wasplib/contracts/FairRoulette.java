// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableMap;

import java.util.ArrayList;

public class FairRoulette {
	private static final Key keyBets = new Key("bets");
	private static final Key keyColor = new Key("color");
	private static final Key keyLastWinningColor = new Key("lastWinningColor");
	private static final Key keyLockedBets = new Key("lockedBets");
	private static final Key keyPlayPeriod = new Key("playPeriod");
	private static final long NUM_COLORS = 5;
	private static final long PLAY_PERIOD = 120;

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("placeBet", FairRoulette::placeBet);
		exports.AddCall("lockBets", FairRoulette::lockBets);
		exports.AddCall("payWinners", FairRoulette::payWinners);
		exports.AddCall("playPeriod", FairRoulette::playPeriod);
		exports.AddCall("nothing", ScExports::nothing);
	}

	public static void placeBet(ScCallContext sc) {
		long amount = sc.Balances().Balance(ScColor.IOTA);
		if (amount == 0) {
			sc.Log("Empty bet...");
			return;
		}
		long color = sc.Params().GetInt(keyColor).Value();
		if (color == 0) {
			sc.Log("No color...");
			return;
		}
		if (color < 1 || color > NUM_COLORS) {
			sc.Log("Invalid color...");
			return;
		}

		BetInfo bet = new BetInfo();
		bet.better = sc.Caller();
		bet.amount = amount;
		bet.color = color;

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray(keyBets);
		int betNr = bets.Length();
		byte[] bytes = encodeBetInfo(bet);
		bets.GetBytes(betNr).SetValue(bytes);
		if (betNr == 0) {
			long playPeriod = state.GetInt(keyPlayPeriod).Value();
			if (playPeriod < 10) {
				playPeriod = PLAY_PERIOD;
			}
			sc.Post("lockBets").Post(playPeriod);
		}
	}

	public static void lockBets(ScCallContext sc) {
		// can only be sent by SC itself
		if (!sc.From(sc.Contract().Id())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray(keyBets);
		ScMutableBytesArray lockedBets = state.GetBytesArray(keyLockedBets);
		int nrBets = bets.Length();
		for (int i = 0; i < nrBets; i++) {
			byte[] bytes = bets.GetBytes(i).Value();
			lockedBets.GetBytes(i).SetValue(bytes);
		}
		bets.Clear();

		sc.Post("payWinners").Post(0);
	}

	public static void payWinners(ScCallContext sc) {
		// can only be sent by SC itself
		ScAgent scId = sc.Contract().Id();
		if (!sc.From(scId)) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long winningColor = sc.Utility().Random(5) + 1;
		ScMutableMap state = sc.State();
		state.GetInt(keyLastWinningColor).SetValue(winningColor);

		long totalBetAmount = 0;
		long totalWinAmount = 0;
		ScMutableBytesArray lockedBets = state.GetBytesArray(keyLockedBets);
		ArrayList<BetInfo> winners = new ArrayList<>();
		int nrBets = lockedBets.Length();
		for (int i = 0; i < nrBets; i++) {
			byte[] bytes = lockedBets.GetBytes(i).Value();
			BetInfo bet = decodeBetInfo(bytes);
			totalBetAmount += bet.amount;
			if (bet.color == winningColor) {
				totalWinAmount += bet.amount;
				winners.add(bet);
			}
		}
		lockedBets.Clear();

		if (winners.size() == 0) {
			sc.Log("Nobody wins!");
			// compact separate UTXOs into a single one
			sc.Transfer(scId, ScColor.IOTA, totalBetAmount);
			return;
		}

		long totalPayout = 0;
		for (int i = 0; i < winners.size(); i++) {
			BetInfo bet = winners.get(i);
			long payout = totalBetAmount * bet.amount / totalWinAmount;
			if (payout != 0) {
				totalPayout += payout;
				sc.Transfer(bet.better, ScColor.IOTA, payout);
			}
			String text = "Pay " + payout + " to " + bet.better;
			sc.Log(text);
		}

		if (totalPayout != totalBetAmount) {
			long remainder = totalBetAmount - totalPayout;
			String text = "Remainder is " + remainder;
			sc.Log(text);
			sc.Transfer(scId, ScColor.IOTA, remainder);
		}
	}

	public static void playPeriod(ScCallContext sc) {
		// can only be sent by SC owner
		if (!sc.From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long playPeriod = sc.Params().GetInt(keyPlayPeriod).Value();
		if (playPeriod < 10) {
			sc.Log("Invalid play period...");
			return;
		}

		sc.State().GetInt(keyPlayPeriod).SetValue(playPeriod);
	}

	public static BetInfo decodeBetInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		BetInfo bet = new BetInfo();
		bet.better = decoder.Agent();
		bet.amount = decoder.Int();
		bet.color = decoder.Int();
		return bet;
	}

	public static byte[] encodeBetInfo(BetInfo bet) {
		return new BytesEncoder().
				Agent(bet.better).
				Int(bet.amount).
				Int(bet.color).
				Data();
	}

	public static class BetInfo {
		ScAgent better;
		long amount;
		long color;
	}
}
