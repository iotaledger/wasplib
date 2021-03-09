// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package org.iota.wasp.wasmlib.keys;

import org.iota.wasp.wasmlib.hashtypes.*;

public class Core {

    public static final ScHname Accounts = new ScHname(0x3c4b5e02);
    public static final ScHname AccountsFuncDeposit = new ScHname(0xbdc9102d);
    public static final ScHname AccountsFuncWithdrawToAddress = new ScHname(0x26608cb5);
    public static final ScHname AccountsFuncWithdrawToChain = new ScHname(0x437bc026);
    public static final ScHname AccountsViewAccounts = new ScHname(0x3c4b5e02);
    public static final ScHname AccountsViewBalance = new ScHname(0x84168cb4);
    public static final ScHname AccountsViewTotalAssets = new ScHname(0xfab0f8d2);

    public static final Key AccountsParamAgentID = new Key("a");

    public static final ScHname Blob = new ScHname(0xfd91bc63);
    public static final ScHname BlobFuncStoreBlob = new ScHname(0xddd4c281);
    public static final ScHname BlobViewGetBlobField = new ScHname(0x1f448130);
    public static final ScHname BlobViewGetBlobInfo = new ScHname(0xfde4ab46);
    public static final ScHname BlobViewListBlobs = new ScHname(0x62ca7990);

    public static final Key BlobParamField = new Key("field");
    public static final Key BlobParamHash = new Key("hash");

    public static final ScHname Eventlog = new ScHname(0x661aa7d8);
    public static final ScHname EventlogViewGetNumRecords = new ScHname(0x2f4b4a8c);
    public static final ScHname EventlogViewGetRecords = new ScHname(0xd01a8085);

    public static final Key EventlogParamContractHname = new Key("contractHname");
    public static final Key EventlogParamFromTs = new Key("fromTs");
    public static final Key EventlogParamMaxLastRecords = new Key("maxLastRecords");
    public static final Key EventlogParamToTs = new Key("toTs");

    public static final ScHname Root = new ScHname(0xcebf5908);
    public static final ScHname RootFuncClaimChainOwnership = new ScHname(0x03ff0fc0);
    public static final ScHname RootFuncDelegateChainOwnership = new ScHname(0x93ecb6ad);
    public static final ScHname RootFuncDeployContract = new ScHname(0x28232c27);
    public static final ScHname RootFuncGrantDeployPermission = new ScHname(0xf440263a);
    public static final ScHname RootFuncRevokeDeployPermission = new ScHname(0x850744f1);
    public static final ScHname RootFuncSetContractFee = new ScHname(0x8421a42b);
    public static final ScHname RootFuncSetDefaultFee = new ScHname(0x3310ecd0);
    public static final ScHname RootViewFindContract = new ScHname(0xc145ca00);
    public static final ScHname RootViewGetChainInfo = new ScHname(0x434477e2);
    public static final ScHname RootViewGetFeeInfo = new ScHname(0x9fe54b48);

    public static final Key RootParamChainOwner = new Key("$$owner$$");
    public static final Key RootParamDeployer = new Key("$$deployer$$");
    public static final Key RootParamDescription = new Key("$$description$$");
    public static final Key RootParamHname = new Key("$$hname$$");
    public static final Key RootParamName = new Key("$$name$$");
    public static final Key RootParamOwnerFee = new Key("$$ownerfee$$");
    public static final Key RootParamProgramHash = new Key("$$proghash$$");
    public static final Key RootParamValidatorFee = new Key("$$validatorfee$$");
}