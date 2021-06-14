// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "root";
pub const SC_DESCRIPTION: &str = "Core root contract";
pub const HSC_NAME:       ScHname = ScHname(0xcebf5908);

pub const PARAM_CHAIN_OWNER:   &str = "$$owner$$";
pub const PARAM_DEPLOYER:      &str = "$$deployer$$";
pub const PARAM_DESCRIPTION:   &str = "$$description$$";
pub const PARAM_HNAME:         &str = "$$hname$$";
pub const PARAM_NAME:          &str = "$$name$$";
pub const PARAM_OWNER_FEE:     &str = "$$ownerfee$$";
pub const PARAM_PROGRAM_HASH:  &str = "$$proghash$$";
pub const PARAM_VALIDATOR_FEE: &str = "$$validatorfee$$";

pub const RESULT_CHAIN_ID:              &str = "c";
pub const RESULT_CHAIN_OWNER_ID:        &str = "o";
pub const RESULT_CONTRACT_REGISTRY:     &str = "r";
pub const RESULT_DATA:                  &str = "dt";
pub const RESULT_DEFAULT_OWNER_FEE:     &str = "do";
pub const RESULT_DEFAULT_VALIDATOR_FEE: &str = "dv";
pub const RESULT_DESCRIPTION:           &str = "d";
pub const RESULT_FEE_COLOR:             &str = "f";
pub const RESULT_OWNER_FEE:             &str = "of";
pub const RESULT_VALIDATOR_FEE:         &str = "vf";

pub const FUNC_CLAIM_CHAIN_OWNERSHIP:    &str = "claimChainOwnership";
pub const FUNC_DELEGATE_CHAIN_OWNERSHIP: &str = "delegateChainOwnership";
pub const FUNC_DEPLOY_CONTRACT:          &str = "deployContract";
pub const FUNC_GRANT_DEPLOY_PERMISSION:  &str = "grantDeployPermission";
pub const FUNC_REVOKE_DEPLOY_PERMISSION: &str = "revokeDeployPermission";
pub const FUNC_SET_CONTRACT_FEE:         &str = "setContractFee";
pub const FUNC_SET_DEFAULT_FEE:          &str = "setDefaultFee";
pub const VIEW_FIND_CONTRACT:            &str = "findContract";
pub const VIEW_GET_CHAIN_INFO:           &str = "getChainInfo";
pub const VIEW_GET_FEE_INFO:             &str = "getFeeInfo";

pub const HFUNC_CLAIM_CHAIN_OWNERSHIP:    ScHname = ScHname(0x03ff0fc0);
pub const HFUNC_DELEGATE_CHAIN_OWNERSHIP: ScHname = ScHname(0x93ecb6ad);
pub const HFUNC_DEPLOY_CONTRACT:          ScHname = ScHname(0x28232c27);
pub const HFUNC_GRANT_DEPLOY_PERMISSION:  ScHname = ScHname(0xf440263a);
pub const HFUNC_REVOKE_DEPLOY_PERMISSION: ScHname = ScHname(0x850744f1);
pub const HFUNC_SET_CONTRACT_FEE:         ScHname = ScHname(0x8421a42b);
pub const HFUNC_SET_DEFAULT_FEE:          ScHname = ScHname(0x3310ecd0);
pub const HVIEW_FIND_CONTRACT:            ScHname = ScHname(0xc145ca00);
pub const HVIEW_GET_CHAIN_INFO:           ScHname = ScHname(0x434477e2);
pub const HVIEW_GET_FEE_INFO:             ScHname = ScHname(0x9fe54b48);

//@formatter:on
