package org.iota.wasplib;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableMap;

import java.util.ArrayList;

public class FairRoulette {
	private static long NUM_COLORS = 5;
	private static long PLAY_PERIOD = 120;

	//export placeBet
	public static void placeBet() {
		ScContext ctx = new ScContext();
		ScRequest request = ctx.Request();
		long amount = request.Balance("iota");
		if (amount == 0) {
			ctx.Log("Empty bet...");
			return;
		}
		long color = request.Params().GetInt("color").Value();
		if (color == 0) {
			ctx.Log("No color...");
			return;
		}
		if (color < 1 || color > NUM_COLORS) {
			ctx.Log("Invalid color...");
			return;
		}

		BetInfo bet = new BetInfo();
		bet.id = request.Hash();
		bet.sender = request.Address();
		bet.color = color;
		bet.amount = amount;

		ScMutableMap state = ctx.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		int betNr = bets.Length();
		byte[] data = encodeBetInfo(bet);
		bets.GetBytes(betNr).SetValue(data);
		if (betNr == 0) {
			long playPeriod = state.GetInt("playPeriod").Value();
			if (playPeriod < 10) {
				playPeriod = PLAY_PERIOD;
			}
			ctx.Event("", "lockBets", playPeriod);
		}
	}

	//export lockBets
	public static void lockBets() {
		// can only be sent by SC itself
		ScContext ctx = new ScContext();
		if (!ctx.Request().Address().equals(ctx.Contract().Address())) {
			ctx.Log("Cancel spoofed request");
			return;
		}

		ScMutableMap state = ctx.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
		for (int i = 0; i < bets.Length(); i++) {
			byte[] bet = bets.GetBytes(i).Value();
			lockedBets.GetBytes(i).SetValue(bet);
		}
		bets.Clear();

		ctx.Event("", "payWinners", 0);
	}

	//export payWinners
	public static void payWinners() {
		// can only be sent by SC itself
		ScContext ctx = new ScContext();
		String scAddress = ctx.Contract().Address();
		if (!ctx.Request().Address().equals(scAddress)) {
			ctx.Log("Cancel spoofed request");
			return;
		}

		long winningcolor = ctx.Random(5) + 1;
		ScMutableMap state = ctx.State();
		state.GetInt("lastWinningColor").SetValue(winningcolor);

		long totalBetAmount = 0;
		long totalWinAmount = 0;
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
		ArrayList<BetInfo> winners = new ArrayList<>();
		for (int i = 0; i < lockedBets.Length(); i++) {
			byte[] betData = lockedBets.GetBytes(i).Value();
			BetInfo bet = decodeBetInfo(betData);
			totalBetAmount += bet.amount;
			if (bet.color == winningcolor) {
				totalWinAmount += bet.amount;
				winners.add(bet);
			}
		}
		lockedBets.Clear();

		if (winners.size() == 0) {
			ctx.Log("Nobody wins!");
			// compact separate UTXOs into a single one
			ctx.Transfer(scAddress, "iota", totalBetAmount);
			return;
		}

		long totalPayout = 0;
		for (int i = 0; i < winners.size(); i++) {
			BetInfo bet = winners.get(i);
			long payout = totalBetAmount * bet.amount / totalWinAmount;
			if (payout != 0) {
				totalPayout += payout;
				ctx.Transfer(bet.sender, "iota", payout);
			}
			String text = "Pay " + payout + " to " + bet.sender;
			ctx.Log(text);
		}

		if (totalPayout != totalBetAmount) {
			long remainder = totalBetAmount - totalPayout;
			String text = "Remainder is " + remainder;
			ctx.Log(text);
			ctx.Transfer(scAddress, "iota", remainder);
		}
	}

	//export playPeriod
	public static void playPeriod() {
		// can only be sent by SC owner
		ScContext ctx = new ScContext();
		if (!ctx.Request().Address().equals(ctx.Contract().Owner())) {
			ctx.Log("Cancel spoofed request");
			return;
		}

		long playPeriod = ctx.Request().Params().GetInt("playPeriod").Value();
		if (playPeriod < 10) {
			ctx.Log("Invalid play period...");
			return;
		}

		ctx.State().GetInt("playPeriod").SetValue(playPeriod);
	}

	public static BetInfo decodeBetInfo(byte[] data) {
		BytesDecoder decoder = new BytesDecoder(data);
		BetInfo bet = new BetInfo();
		bet.id = decoder.String();
		bet.sender = decoder.String();
		bet.color = decoder.Int();
		bet.amount = decoder.Int();
		return bet;
	}

	public static byte[] encodeBetInfo(BetInfo bet) {
		return new BytesEncoder().
				String(bet.id).
				String(bet.sender).
				Int(bet.color).
				Int(bet.amount).
				Data();
	}

	public static class BetInfo {
		String id;
		String sender;
		long color;
		long amount;
	}
}
