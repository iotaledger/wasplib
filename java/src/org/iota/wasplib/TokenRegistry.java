package org.iota.wasplib;

import org.iota.wasplib.client.bytes.BytesDecoder;
import org.iota.wasplib.client.bytes.BytesEncoder;
import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;

public class TokenRegistry {
	//export mintSupply
	public static void mintSupply() {
		ScContext ctx = new ScContext();
		ScRequest request = ctx.Request();
		String color = request.Hash();
		ScMutableMap state = ctx.State();
		ScMutableMap registry = state.GetMap("tr");
		if (!registry.GetString(color).Value().isEmpty()) {
			ctx.Log("TokenRegistry: Color already exists");
			return;
		}
		ScImmutableMap reqParams = request.Params();
		TokenInfo token = new TokenInfo();
		token.supply = request.Balance(color);
		token.mintedBy = request.Address();
		token.owner = request.Address();
		token.created = request.Timestamp();
		token.updated = request.Timestamp();
		token.description = reqParams.GetString("dscr").Value();
		token.userDefined = reqParams.GetString("ud").Value();
		if (token.supply <= 0) {
			ctx.Log("TokenRegistry: Insufficient supply");
			return;
		}
		if (token.description.isEmpty()) {
			token.description += "no dscr";
		}
		byte[] data = encodeTokenInfo(token);
		registry.GetBytes(color).SetValue(data);
		ScMutableString colors = state.GetString("lc");
		String list = colors.Value();
		if (!list.isEmpty()) {
			list += ",";
		}
		list += color;
		colors.SetValue(list);
	}

	//export updateMetadata
	public static void updateMetadata() {
		ScContext ctx = new ScContext();
	}

	//export transferOwnership
	public static void transferOwnership() {
		ScContext ctx = new ScContext();
	}

	public static TokenInfo decodeTokenInfo(byte[] data) {
		BytesDecoder decoder = new BytesDecoder(data);
		TokenInfo token = new TokenInfo();
		token.supply = decoder.Int();
		token.mintedBy = decoder.String();
		token.owner = decoder.String();
		token.created = decoder.Int();
		token.updated = decoder.Int();
		token.description = decoder.String();
		token.userDefined = decoder.String();
		return token;
	}

	public static byte[] encodeTokenInfo(TokenInfo data) {
		return new BytesEncoder().
				Int(data.supply).
				String(data.mintedBy).
				String(data.owner).
				Int(data.created).
				Int(data.updated).
				String(data.description).
				String(data.userDefined).
				Data();
	}

	public static class TokenInfo {
		long supply;
		String mintedBy;
		String owner;
		long created;
		long updated;
		String description;
		String userDefined;
	}
}
