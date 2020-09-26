package org.iota.wasplib;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScAccount;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScLog;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class DonateWithFeedback {
	//export donate
	public static void donate() {
		ScContext ctx = new ScContext();
		ScLog tlog = ctx.TimestampedLog("l");
		ScRequest request = ctx.Request();
		DonationInfo di = new DonationInfo();
		di.seq = tlog.Length();
		di.id = request.Id();
		di.amount = request.Balance("iota");
		di.sender = request.Address();
		di.feedback = request.Params().GetString("f").Value();
		di.error = "";
		if (di.amount == 0 || di.feedback.length() == 0) {
			di.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)";
			if (di.amount > 0) {
				ctx.Transfer(di.sender, "iota", di.amount);
				di.amount = 0;
			}
		}
		byte[] data = encodeDonationInfo(di);
		tlog.Append(request.Timestamp(), data);

		ScMutableMap state = ctx.State();
		ScMutableInt maxd = state.GetInt("maxd");
		ScMutableInt total = state.GetInt("total");
		if (di.amount > maxd.Value()) {
			maxd.SetValue(di.amount);
		}
		total.SetValue(total.Value() + di.amount);
	}

	//export withdraw
	public static void withdraw() {
		ScContext ctx = new ScContext();
	}

	//export transferOwnership
	public static void transferOwnership() {
		ScContext ctx = new ScContext();
		String owner = ctx.Contract().Owner();
		ScRequest request = ctx.Request();
		if (!request.Address().equals(owner)) {
			ctx.Log("Cancel spoofed request");
			return;
		}

		ScAccount account = ctx.Account();
		long bal = account.Balance("iota");
		long withdrawSum = request.Params().GetInt("s").Value();
		if (withdrawSum == 0 || withdrawSum > bal) {
			withdrawSum = bal;
		}
		if (withdrawSum == 0) {
			ctx.Log("DonateWithFeedback: withdraw. nothing to withdraw");
			return;
		}

		ctx.Transfer(owner, "iota", withdrawSum);
	}

	public static DonationInfo decodeDonationInfo(byte[] data) {
		BytesDecoder decoder = new BytesDecoder(data);
		DonationInfo bet = new DonationInfo();
		bet.seq = decoder.Int();
		bet.id = decoder.String();
		bet.amount = decoder.Int();
		bet.sender = decoder.String();
		bet.feedback = decoder.String();
		bet.error = decoder.String();
		return bet;
	}

	public static byte[] encodeDonationInfo(DonationInfo bet) {
		return new BytesEncoder().
				Int(bet.seq).
				String(bet.id).
				Int(bet.amount).
				String(bet.sender).
				String(bet.feedback).
				String(bet.error).
				Data();
	}

	public static class DonationInfo {
		long seq;
		String id;
		long amount;
		String sender;
		String feedback;
		String error;
	}
}
