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
		ScContext sc = new ScContext();
		ScRequest request = sc.Request();
		String color = request.Hash();
		ScMutableMap state = sc.State();
		ScMutableMap registry = state.GetMap("tr");
		if (!registry.GetString(color).Value().isEmpty()) {
			sc.Log("TokenRegistry: Color already exists");
			return;
		}
		ScImmutableMap params = request.Params();
		TokenInfo token = new TokenInfo();
		token.supply = request.Balance(color);
		token.mintedBy = request.Address();
		token.owner = request.Address();
		token.created = request.Timestamp();
		token.updated = request.Timestamp();
		token.description = params.GetString("dscr").Value();
		token.userDefined = params.GetString("ud").Value();
		if (token.supply <= 0) {
			sc.Log("TokenRegistry: Insufficient supply");
			return;
		}
		if (token.description.isEmpty()) {
			token.description += "no dscr";
		}
		byte[] bytes = encodeTokenInfo(token);
		registry.GetBytes(color).SetValue(bytes);
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
		//ScContext sc = new ScContext();
	}

	//export transferOwnership
	public static void transferOwnership() {
		//ScContext sc = new ScContext();
	}

	public static TokenInfo decodeTokenInfo(byte[] bytes) {
		BytesDecoder decoder = new BytesDecoder(bytes);
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

	public static byte[] encodeTokenInfo(TokenInfo token) {
		return new BytesEncoder().
				Int(token.supply).
				String(token.mintedBy).
				String(token.owner).
				Int(token.created).
				Int(token.updated).
				String(token.description).
				String(token.userDefined).
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
