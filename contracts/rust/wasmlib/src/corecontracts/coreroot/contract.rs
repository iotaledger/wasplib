// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use std::ptr;

use crate::*;
use crate::corecontracts::coreroot::*;

pub struct ClaimChainOwnershipCall {
    pub func: ScFunc,
}

impl ClaimChainOwnershipCall {
    pub fn new(_ctx: &ScFuncContext) -> ClaimChainOwnershipCall {
        ClaimChainOwnershipCall {
            func: ScFunc::new(HSC_NAME, HFUNC_CLAIM_CHAIN_OWNERSHIP),
        }
    }
}

pub struct DelegateChainOwnershipCall {
    pub func:   ScFunc,
    pub params: MutableDelegateChainOwnershipParams,
}

impl DelegateChainOwnershipCall {
    pub fn new(_ctx: &ScFuncContext) -> DelegateChainOwnershipCall {
        let mut f = DelegateChainOwnershipCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_DELEGATE_CHAIN_OWNERSHIP),
            params: MutableDelegateChainOwnershipParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct DeployContractCall {
    pub func:   ScFunc,
    pub params: MutableDeployContractParams,
}

impl DeployContractCall {
    pub fn new(_ctx: &ScFuncContext) -> DeployContractCall {
        let mut f = DeployContractCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_DEPLOY_CONTRACT),
            params: MutableDeployContractParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct GrantDeployPermissionCall {
    pub func:   ScFunc,
    pub params: MutableGrantDeployPermissionParams,
}

impl GrantDeployPermissionCall {
    pub fn new(_ctx: &ScFuncContext) -> GrantDeployPermissionCall {
        let mut f = GrantDeployPermissionCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_GRANT_DEPLOY_PERMISSION),
            params: MutableGrantDeployPermissionParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct RevokeDeployPermissionCall {
    pub func:   ScFunc,
    pub params: MutableRevokeDeployPermissionParams,
}

impl RevokeDeployPermissionCall {
    pub fn new(_ctx: &ScFuncContext) -> RevokeDeployPermissionCall {
        let mut f = RevokeDeployPermissionCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_REVOKE_DEPLOY_PERMISSION),
            params: MutableRevokeDeployPermissionParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct SetContractFeeCall {
    pub func:   ScFunc,
    pub params: MutableSetContractFeeParams,
}

impl SetContractFeeCall {
    pub fn new(_ctx: &ScFuncContext) -> SetContractFeeCall {
        let mut f = SetContractFeeCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_CONTRACT_FEE),
            params: MutableSetContractFeeParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct SetDefaultFeeCall {
    pub func:   ScFunc,
    pub params: MutableSetDefaultFeeParams,
}

impl SetDefaultFeeCall {
    pub fn new(_ctx: &ScFuncContext) -> SetDefaultFeeCall {
        let mut f = SetDefaultFeeCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_DEFAULT_FEE),
            params: MutableSetDefaultFeeParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

pub struct FindContractCall {
    pub func:    ScView,
    pub params:  MutableFindContractParams,
    pub results: ImmutableFindContractResults,
}

impl FindContractCall {
    pub fn new(_ctx: &ScFuncContext) -> FindContractCall {
        let mut f = FindContractCall {
            func:    ScView::new(HSC_NAME, HVIEW_FIND_CONTRACT),
            params:  MutableFindContractParams { id: 0 },
            results: ImmutableFindContractResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> FindContractCall {
        FindContractCall::new(&ScFuncContext {})
    }
}

pub struct GetChainInfoCall {
    pub func:    ScView,
    pub results: ImmutableGetChainInfoResults,
}

impl GetChainInfoCall {
    pub fn new(_ctx: &ScFuncContext) -> GetChainInfoCall {
        let mut f = GetChainInfoCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_CHAIN_INFO),
            results: ImmutableGetChainInfoResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetChainInfoCall {
        GetChainInfoCall::new(&ScFuncContext {})
    }
}

pub struct GetFeeInfoCall {
    pub func:    ScView,
    pub params:  MutableGetFeeInfoParams,
    pub results: ImmutableGetFeeInfoResults,
}

impl GetFeeInfoCall {
    pub fn new(_ctx: &ScFuncContext) -> GetFeeInfoCall {
        let mut f = GetFeeInfoCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_FEE_INFO),
            params:  MutableGetFeeInfoParams { id: 0 },
            results: ImmutableGetFeeInfoResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn new_from_view(_ctx: &ScViewContext) -> GetFeeInfoCall {
        GetFeeInfoCall::new(&ScFuncContext {})
    }
}

//@formatter:on
