// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.bytes.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScFuncContext extends ScBaseContext {
	public ScFuncContext() {
	}

	public ScImmutableMap Call(ScHname hContract, ScHname hFunction, ScMutableMap params, ScTransfers transfer) {
		BytesEncoder encode = new BytesEncoder();
		encode.Hname(hContract);
		encode.Hname(hFunction);
		if (params != null) {
			encode.Int(params.mapId());
		} else {
			encode.Int(0);
		}
		if (transfer != null) {
			encode.Int(transfer.mapId());
		} else {
			encode.Int(0);
		}
		Host.root.GetBytes(Key.Call).SetValue(encode.Data());
		return Host.root.GetMap(Key.Return).Immutable();
	}

	public ScAgentId Caller() {
		return Host.root.GetAgentId(Key.Caller).Value();
	}

	public ScImmutableMap CallSelf(ScHname hFunction, ScMutableMap params, ScTransfers transfer) {
		return Call(ContractId().Hname(), hFunction, params, transfer);
	}

	public void Deploy(ScHash programHash, String name, String description, ScMutableMap params) {
		BytesEncoder encode = new BytesEncoder();
		encode.Hash(programHash);
		encode.String(name);
		encode.String(description);
		if (params != null) {
			encode.Int(params.mapId());
		} else {
			encode.Int(0);
		}
		Host.root.GetBytes(Key.Deploy).SetValue(encode.Data());
	}

	public void Event(String text) {
		Host.root.GetString(Key.Event).SetValue(text);
	}

	public ScBalances Incoming() {
		return new ScBalances(Host.root.GetMap(Key.Incoming).Immutable());
	}

	public ScColor MintedColor() {
		return new ScColor(RequestId());
	}

	public long MintedSupply() {
		return Host.root.GetInt(Key.Minted).Value();
	}

	public void Post(PostRequestParams req) {
		BytesEncoder encode = new BytesEncoder();
		encode.ContractId(req.ContractId);
		encode.Hname(req.Function);
		if (req.Params != null) {
			encode.Int(req.Params.mapId());
		} else {
			encode.Int(0);
		}
		if (req.Transfer != null) {
			encode.Int(req.Transfer.mapId());
		} else {
			encode.Int(0);
		}
		encode.Int(req.Delay);
		Host.root.GetBytes(Key.Post).SetValue(encode.Data());
	}

	public ScRequestId RequestId() {
		return Host.root.GetRequestId(Key.RequestId).Value();
	}

	public ScMutableMap State() {
		return Host.root.GetMap(Key.State);
	}

	public void TransferToAddress(ScAddress address, ScTransfers transfer) {
		ScMutableMapArray transfers = Host.root.GetMapArray(Key.Transfers);
		ScMutableMap tx = transfers.GetMap(transfers.Length());
		tx.GetAddress(Key.Address).SetValue(address);
		tx.GetInt(Key.Balances).SetValue(transfer.mapId());
	}
}
