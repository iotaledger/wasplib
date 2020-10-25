package org.iota.wasplib.contracts;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScExports;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableMap;

import java.util.ArrayList;

public class FairRoulette {
	private static long NUM_COLORS = 5;
	private static long PLAY_PERIOD = 120;

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.Add("placeBet");
		exports.Add("lockBets");
		exports.Add("payWinners");
		exports.AddProtected("playPeriod");
		exports.Add("nothing");
	}

	//export placeBet
	public static void placeBet() {
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		long amount = request.Balance(ScColor.IOTA);
		if (amount == 0) {
			sc.Log("Empty bet...");
			return;
		}
		long color = request.Params().GetInt("color").Value();
		if (color == 0) {
			sc.Log("No color...");
			return;
		}
		if (color < 1 || color > NUM_COLORS) {
			sc.Log("Invalid color...");
			return;
		}

		BetInfo bet = new BetInfo();
		bet.id = request.Id();
		bet.sender = request.Address();
		bet.color = color;
		bet.amount = amount;

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		int betNr = bets.Length();
		byte[] bytes = encodeBetInfo(bet);
		bets.GetBytes(betNr).SetValue(bytes);
		if (betNr == 0) {
			long playPeriod = state.GetInt("playPeriod").Value();
			if (playPeriod < 10) {
				playPeriod = PLAY_PERIOD;
			}
			sc.PostRequest(sc.Contract().Address(), "lockBets", playPeriod);
		}
	}

	//export lockBets
	public static void lockBets() {
		// can only be sent by SC itself
		ScContext sc = new ScContext();
		ScAddress scAddress = sc.Contract().Address();
		if (!sc.Request().From(scAddress)) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
		int nrBets = bets.Length();
		for (int i = 0; i < nrBets; i++) {
			byte[] bytes = bets.GetBytes(i).Value();
			lockedBets.GetBytes(i).SetValue(bytes);
		}
		bets.Clear();

		sc.PostRequest(scAddress, "payWinners", 0);
	}

	//export payWinners
	public static void payWinners() {
		// can only be sent by SC itself
		ScContext sc = new ScContext();
		ScAddress scAddress = sc.Contract().Address();
		if (!sc.Request().From(scAddress)) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long winningColor = sc.Random(5) + 1;
		ScMutableMap state = sc.State();
		state.GetInt("lastWinningColor").SetValue(winningColor);

		long totalBetAmount = 0;
		long totalWinAmount = 0;
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
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
			sc.Transfer(scAddress, ScColor.IOTA, totalBetAmount);
			return;
		}

		long totalPayout = 0;
		for (int i = 0; i < winners.size(); i++) {
			BetInfo bet = winners.get(i);
			long payout = totalBetAmount * bet.amount / totalWinAmount;
			if (payout != 0) {
				totalPayout += payout;
				sc.Transfer(bet.sender, ScColor.IOTA, payout);
			}
			String text = "Pay " + payout + " to " + bet.sender;
			sc.Log(text);
		}

		if (totalPayout != totalBetAmount) {
			long remainder = totalBetAmount - totalPayout;
			String text = "Remainder is " + remainder;
			sc.Log(text);
			sc.Transfer(scAddress, ScColor.IOTA, remainder);
		}
	}

	//export playPeriod
	public static void playPeriod() {
		// can only be sent by SC owner
		ScContext sc = new ScContext();
		if (!sc.Request().From(sc.Contract().Owner())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long playPeriod = sc.Request().Params().GetInt("playPeriod").Value();
		if (playPeriod < 10) {
			sc.Log("Invalid play period...");
			return;
		}

		sc.State().GetInt("playPeriod").SetValue(playPeriod);
	}

	public static BetInfo decodeBetInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		BetInfo bet = new BetInfo();
		bet.id = decoder.String();
		bet.sender = decoder.Address();
		bet.color = decoder.Int();
		bet.amount = decoder.Int();
		return bet;
	}

	public static byte[] encodeBetInfo(BetInfo bet) {
		return new BytesEncoder().
				String(bet.id).
				Address(bet.sender).
				Int(bet.color).
				Int(bet.amount).
				Data();
	}

	public static class BetInfo {
		String id;
		ScAddress sender;
		long color;
		long amount;
	}
}
