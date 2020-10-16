package org.iota.wasplib;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScExports;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.mutable.ScMutableBytesArray;
import org.iota.wasplib.client.mutable.ScMutableMap;

import java.util.ArrayList;

public class FairAuction {
	private static long NUM_COLORS = 5;
	private static long PLAY_PERIOD = 120;

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.Add("startAuction");
		exports.Add("finalizeAuction");
		exports.Add("placeBid");
		exports.AddProtected("setOwnerMargin");
	}

	//export startAuction
	public static void startAuction() {
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		long amount = request.Balance("iota");
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

		BidInfo bet = new BidInfo();
		bet.id = request.Hash();
		bet.sender = request.Address();
		bet.color = color;
		bet.amount = amount;

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		int betNr = bets.Length();
		byte[] bytes = encodeBidInfo(bet);
		bets.GetBytes(betNr).SetValue(bytes);
		if (betNr == 0) {
			long playPeriod = state.GetInt("playPeriod").Value();
			if (playPeriod < 10) {
				playPeriod = PLAY_PERIOD;
			}
			sc.Event("", "lockBets", playPeriod);
		}
	}

	//export finalizeAuction
	public static void finalizeAuction() {
		// can only be sent by SC itself
		ScContext sc = new ScContext();
		if (!sc.Request().Address().equals(sc.Contract().Address())) {
			sc.Log("Cancel spoofed request");
			return;
		}

		ScMutableMap state = sc.State();
		ScMutableBytesArray bets = state.GetBytesArray("bets");
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
		for (int i = 0; i < bets.Length(); i++) {
			byte[] bytes = bets.GetBytes(i).Value();
			lockedBets.GetBytes(i).SetValue(bytes);
		}
		bets.Clear();

		sc.Event("", "payWinners", 0);
	}

	//export placeBid
	public static void placeBid() {
		// can only be sent by SC itself
		ScContext sc = new ScContext();
		String scAddress = sc.Contract().Address();
		if (!sc.Request().Address().equals(scAddress)) {
			sc.Log("Cancel spoofed request");
			return;
		}

		long winningColor = sc.Random(5) + 1;
		ScMutableMap state = sc.State();
		state.GetInt("lastWinningColor").SetValue(winningColor);

		long totalBetAmount = 0;
		long totalWinAmount = 0;
		ScMutableBytesArray lockedBets = state.GetBytesArray("lockedBets");
		ArrayList<BidInfo> winners = new ArrayList<>();
		for (int i = 0; i < lockedBets.Length(); i++) {
			byte[] bytes = lockedBets.GetBytes(i).Value();
			BidInfo bet = decodeBidInfo(bytes);
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
			sc.Transfer(scAddress, "iota", totalBetAmount);
			return;
		}

		long totalPayout = 0;
		for (int i = 0; i < winners.size(); i++) {
			BidInfo bet = winners.get(i);
			long payout = totalBetAmount * bet.amount / totalWinAmount;
			if (payout != 0) {
				totalPayout += payout;
				sc.Transfer(bet.sender, "iota", payout);
			}
			String text = "Pay " + payout + " to " + bet.sender;
			sc.Log(text);
		}

		if (totalPayout != totalBetAmount) {
			long remainder = totalBetAmount - totalPayout;
			String text = "Remainder is " + remainder;
			sc.Log(text);
			sc.Transfer(scAddress, "iota", remainder);
		}
	}

	//export setOwnerMargin
	public static void setOwnerMargin() {
		// can only be sent by SC owner
		ScContext sc = new ScContext();
		if (!sc.Request().Address().equals(sc.Contract().Owner())) {
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

	public static BidInfo decodeBidInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
		BidInfo bet = new BidInfo();
		bet.id = decoder.String();
		bet.sender = decoder.String();
		bet.color = decoder.Int();
		bet.amount = decoder.Int();
		return bet;
	}

	public static byte[] encodeBidInfo(BidInfo bet) {
		return new BytesEncoder().
				String(bet.id).
				String(bet.sender).
				Int(bet.color).
				Int(bet.amount).
				Data();
	}

	public static class BidInfo {
		String id;
		String sender;
		long color;
		long amount;
	}
}
