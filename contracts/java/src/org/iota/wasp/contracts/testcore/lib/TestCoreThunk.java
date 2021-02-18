// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead


package org.iota.wasp.contracts.testcore.lib;

import org.iota.wasp.contracts.testcore.*;
import org.iota.wasp.wasmlib.exports.*;
import org.iota.wasp.wasmlib.hashtypes.*;

public class TestCoreThunk {
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddFunc("callOnChain", TestCore::FuncCallOnChain);
		exports.AddFunc("checkContextFromFullEP", TestCore::FuncCheckContextFromFullEP);
		exports.AddFunc("doNothing", TestCore::FuncDoNothing);
		exports.AddFunc("init", TestCore::FuncInit);
		exports.AddFunc("passTypesFull", TestCore::FuncPassTypesFull);
		exports.AddFunc("runRecursion", TestCore::FuncRunRecursion);
		exports.AddFunc("sendToAddress", TestCore::FuncSendToAddress);
		exports.AddFunc("setInt", TestCore::FuncSetInt);
		exports.AddFunc("testCallPanicFullEP", TestCore::FuncTestCallPanicFullEP);
		exports.AddFunc("testCallPanicViewEPFromFull", TestCore::FuncTestCallPanicViewEPFromFull);
		exports.AddFunc("testChainOwnerIDFull", TestCore::FuncTestChainOwnerIDFull);
		exports.AddFunc("testContractIDFull", TestCore::FuncTestContractIDFull);
		exports.AddFunc("testEventLogDeploy", TestCore::FuncTestEventLogDeploy);
		exports.AddFunc("testEventLogEventData", TestCore::FuncTestEventLogEventData);
		exports.AddFunc("testEventLogGenericData", TestCore::FuncTestEventLogGenericData);
		exports.AddFunc("testPanicFullEP", TestCore::FuncTestPanicFullEP);
		exports.AddFunc("withdrawToChain", TestCore::FuncWithdrawToChain);
		exports.AddView("checkContextFromViewEP", TestCore::ViewCheckContextFromViewEP);
		exports.AddView("fibonacci", TestCore::ViewFibonacci);
		exports.AddView("getCounter", TestCore::ViewGetCounter);
		exports.AddView("getInt", TestCore::ViewGetInt);
		exports.AddView("justView", TestCore::ViewJustView);
		exports.AddView("passTypesView", TestCore::ViewPassTypesView);
		exports.AddView("testCallPanicViewEPFromView", TestCore::ViewTestCallPanicViewEPFromView);
		exports.AddView("testChainOwnerIDView", TestCore::ViewTestChainOwnerIDView);
		exports.AddView("testContractIDView", TestCore::ViewTestContractIDView);
		exports.AddView("testPanicViewEP", TestCore::ViewTestPanicViewEP);
		exports.AddView("testSandboxCall", TestCore::ViewTestSandboxCall);
	}
}
