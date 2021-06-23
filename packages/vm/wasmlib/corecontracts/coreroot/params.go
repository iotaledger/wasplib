// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreroot

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ImmutableDelegateChainOwnershipParams struct {
	id int32
}

func (s ImmutableDelegateChainOwnershipParams) ChainOwner() wasmlib.ScImmutableAgentId {
	return wasmlib.NewScImmutableAgentId(s.id, ParamChainOwner.KeyId())
}

type MutableDelegateChainOwnershipParams struct {
	id int32
}

func (s MutableDelegateChainOwnershipParams) ChainOwner() wasmlib.ScMutableAgentId {
	return wasmlib.NewScMutableAgentId(s.id, ParamChainOwner.KeyId())
}

type ImmutableDeployContractParams struct {
	id int32
}

func (s ImmutableDeployContractParams) Description() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, ParamDescription.KeyId())
}

func (s ImmutableDeployContractParams) Name() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, ParamName.KeyId())
}

func (s ImmutableDeployContractParams) ProgramHash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, ParamProgramHash.KeyId())
}

type MutableDeployContractParams struct {
	id int32
}

func (s MutableDeployContractParams) Description() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, ParamDescription.KeyId())
}

func (s MutableDeployContractParams) Name() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, ParamName.KeyId())
}

func (s MutableDeployContractParams) ProgramHash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, ParamProgramHash.KeyId())
}

type ImmutableGrantDeployPermissionParams struct {
	id int32
}

func (s ImmutableGrantDeployPermissionParams) Deployer() wasmlib.ScImmutableAgentId {
	return wasmlib.NewScImmutableAgentId(s.id, ParamDeployer.KeyId())
}

type MutableGrantDeployPermissionParams struct {
	id int32
}

func (s MutableGrantDeployPermissionParams) Deployer() wasmlib.ScMutableAgentId {
	return wasmlib.NewScMutableAgentId(s.id, ParamDeployer.KeyId())
}

type ImmutableRevokeDeployPermissionParams struct {
	id int32
}

func (s ImmutableRevokeDeployPermissionParams) Deployer() wasmlib.ScImmutableAgentId {
	return wasmlib.NewScImmutableAgentId(s.id, ParamDeployer.KeyId())
}

type MutableRevokeDeployPermissionParams struct {
	id int32
}

func (s MutableRevokeDeployPermissionParams) Deployer() wasmlib.ScMutableAgentId {
	return wasmlib.NewScMutableAgentId(s.id, ParamDeployer.KeyId())
}

type ImmutableSetContractFeeParams struct {
	id int32
}

func (s ImmutableSetContractFeeParams) Hname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamHname.KeyId())
}

func (s ImmutableSetContractFeeParams) OwnerFee() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamOwnerFee.KeyId())
}

func (s ImmutableSetContractFeeParams) ValidatorFee() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamValidatorFee.KeyId())
}

type MutableSetContractFeeParams struct {
	id int32
}

func (s MutableSetContractFeeParams) Hname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamHname.KeyId())
}

func (s MutableSetContractFeeParams) OwnerFee() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamOwnerFee.KeyId())
}

func (s MutableSetContractFeeParams) ValidatorFee() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamValidatorFee.KeyId())
}

type ImmutableSetDefaultFeeParams struct {
	id int32
}

func (s ImmutableSetDefaultFeeParams) OwnerFee() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamOwnerFee.KeyId())
}

func (s ImmutableSetDefaultFeeParams) ValidatorFee() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, ParamValidatorFee.KeyId())
}

type MutableSetDefaultFeeParams struct {
	id int32
}

func (s MutableSetDefaultFeeParams) OwnerFee() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamOwnerFee.KeyId())
}

func (s MutableSetDefaultFeeParams) ValidatorFee() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, ParamValidatorFee.KeyId())
}

type ImmutableFindContractParams struct {
	id int32
}

func (s ImmutableFindContractParams) Hname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamHname.KeyId())
}

type MutableFindContractParams struct {
	id int32
}

func (s MutableFindContractParams) Hname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamHname.KeyId())
}

type ImmutableGetFeeInfoParams struct {
	id int32
}

func (s ImmutableGetFeeInfoParams) Hname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamHname.KeyId())
}

type MutableGetFeeInfoParams struct {
	id int32
}

func (s MutableGetFeeInfoParams) Hname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamHname.KeyId())
}
