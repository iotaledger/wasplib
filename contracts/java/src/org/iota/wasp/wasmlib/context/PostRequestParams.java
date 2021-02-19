package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.mutable.*;

public class PostRequestParams {
	public ScContractId ContractId;
	public ScHname Function;
	public ScMutableMap Params;
	public ScTransfers Transfer;
	public long Delay;
}
