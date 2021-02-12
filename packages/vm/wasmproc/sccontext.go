// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmproc

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/cbalances"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

var typeIds = map[int32]int32{
	wasmhost.KeyBalances:        wasmhost.OBJTYPE_MAP,
	wasmhost.KeyCall:            wasmhost.OBJTYPE_BYTES,
	wasmhost.KeyCaller:          wasmhost.OBJTYPE_AGENT_ID,
	wasmhost.KeyChainOwnerId:    wasmhost.OBJTYPE_AGENT_ID,
	wasmhost.KeyContractCreator: wasmhost.OBJTYPE_AGENT_ID,
	wasmhost.KeyDeploy:          wasmhost.OBJTYPE_BYTES,
	wasmhost.KeyEvent:           wasmhost.OBJTYPE_STRING,
	wasmhost.KeyExports:         wasmhost.OBJTYPE_STRING | wasmhost.OBJTYPE_ARRAY,
	wasmhost.KeyContractId:      wasmhost.OBJTYPE_CONTRACT_ID,
	wasmhost.KeyIncoming:        wasmhost.OBJTYPE_MAP,
	wasmhost.KeyLog:             wasmhost.OBJTYPE_STRING,
	wasmhost.KeyMaps:            wasmhost.OBJTYPE_MAP | wasmhost.OBJTYPE_ARRAY,
	wasmhost.KeyPanic:           wasmhost.OBJTYPE_STRING,
	wasmhost.KeyParams:          wasmhost.OBJTYPE_MAP,
	wasmhost.KeyPost:            wasmhost.OBJTYPE_BYTES,
	wasmhost.KeyResults:         wasmhost.OBJTYPE_MAP,
	wasmhost.KeyReturn:          wasmhost.OBJTYPE_MAP,
	wasmhost.KeyState:           wasmhost.OBJTYPE_MAP,
	wasmhost.KeyTimestamp:       wasmhost.OBJTYPE_INT,
	wasmhost.KeyTrace:           wasmhost.OBJTYPE_STRING,
	wasmhost.KeyTransfers:       wasmhost.OBJTYPE_MAP | wasmhost.OBJTYPE_ARRAY,
	wasmhost.KeyUtility:         wasmhost.OBJTYPE_MAP,
}

type ScContext struct {
	ScSandboxObject
}

func NewScContext(vm *wasmProcessor) *ScContext {
	o := &ScContext{}
	o.vm = vm
	o.host = &vm.KvStoreHost
	o.name = "root"
	o.id = 1
	o.isRoot = true
	o.objects = make(map[int32]int32)
	return o
}

func (o *ScContext) Exists(keyId int32, typeId int32) bool {
	if keyId == wasmhost.KeyExports {
		return o.vm.ctx == nil && o.vm.ctxView == nil
	}
	return o.GetTypeId(keyId) > 0
}

func (o *ScContext) GetBytes(keyId int32, typeId int32) []byte {
	switch keyId {
	case wasmhost.KeyCaller:
		return o.vm.ctx.Caller().Bytes()
	case wasmhost.KeyChainOwnerId:
		return o.vm.chainOwnerID().Bytes()
	case wasmhost.KeyContractCreator:
		return o.vm.contractCreator().Bytes()
	case wasmhost.KeyContractId:
		return o.vm.contractID().Bytes()
	case wasmhost.KeyTimestamp:
		return codec.EncodeInt64(o.vm.ctx.GetTimestamp())
	}
	o.invalidKey(keyId)
	return nil
}

func (o *ScContext) GetObjectId(keyId int32, typeId int32) int32 {
	if keyId == wasmhost.KeyExports && (o.vm.ctx != nil || o.vm.ctxView != nil) {
		// once map has entries (after on_load) this cannot be called any more
		o.invalidKey(keyId)
		return 0
	}

	return GetMapObjectId(o, keyId, typeId, ObjFactories{
		wasmhost.KeyBalances:  func() WaspObject { return NewScBalances(o.vm, false) },
		wasmhost.KeyExports:   func() WaspObject { return NewScExports(o.vm) },
		wasmhost.KeyIncoming:  func() WaspObject { return NewScBalances(o.vm, true) },
		wasmhost.KeyMaps:      func() WaspObject { return NewScMaps(o.vm) },
		wasmhost.KeyParams:    func() WaspObject { return NewScDictFromKvStore(&o.vm.KvStoreHost, o.vm.params()) },
		wasmhost.KeyResults:   func() WaspObject { return NewScDict(o.vm) },
		wasmhost.KeyReturn:    func() WaspObject { return NewScDict(o.vm) },
		wasmhost.KeyState:     func() WaspObject { return NewScDictFromKvStore(&o.vm.KvStoreHost, o.vm.state()) },
		wasmhost.KeyTransfers: func() WaspObject { return NewScTransfers(o.vm) },
		wasmhost.KeyUtility:   func() WaspObject { return NewScUtility(o.vm) },
	})
}

func (o *ScContext) GetTypeId(keyId int32) int32 {
	return typeIds[keyId]
}

func (o *ScContext) SetBytes(keyId int32, typeId int32, bytes []byte) {
	switch keyId {
	case wasmhost.KeyCall:
		o.processCall(bytes)
	case wasmhost.KeyDeploy:
		o.processDeploy(bytes)
	case wasmhost.KeyEvent:
		o.vm.ctx.Event(string(bytes))
	case wasmhost.KeyLog:
		o.vm.log().Infof(string(bytes))
	case wasmhost.KeyTrace:
		o.vm.log().Debugf(string(bytes))
	case wasmhost.KeyPanic:
		o.vm.log().Panicf(string(bytes))
	case wasmhost.KeyPost:
		o.processPost(bytes)
	default:
		o.invalidKey(keyId)
	}
}

func (o *ScContext) processCall(bytes []byte) {
	decode := NewBytesDecoder(bytes)
	contract, err := coretypes.NewHnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	function, err := coretypes.NewHnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	params := o.getParams(int32(decode.Int()))
	transfer := o.getTransfer(int32(decode.Int()))

	o.Trace("CALL c'%s' f'%s'", contract.String(), function.String())
	var results dict.Dict
	if o.vm.ctx != nil {
		results, err = o.vm.ctx.Call(contract, function, params, transfer)
	} else {
		results, err = o.vm.ctxView.Call(contract, function, params)
	}
	if err != nil {
		o.Panic("failed to invoke call: %v", err)
	}
	resultsId := o.GetObjectId(wasmhost.KeyReturn, wasmhost.OBJTYPE_MAP)
	o.host.FindObject(resultsId).(*ScDict).kvStore = results
}

func (o *ScContext) processDeploy(bytes []byte) {
	decode := NewBytesDecoder(bytes)
	programHash, err := hashing.HashValueFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	name := string(decode.Bytes())
	description := string(decode.Bytes())
	params := o.getParams(int32(decode.Int()))
	o.Trace("DEPLOY c'%s' f'%s'", name, description)
	err = o.vm.ctx.DeployContract(programHash, name, description, params)
	if err != nil {
		o.Panic("failed to deploy: %v", err)
	}
}

func (o *ScContext) processPost(bytes []byte) {
	decode := NewBytesDecoder(bytes)
	contract, err := coretypes.NewContractIDFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	function, err := coretypes.NewHnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	o.Trace("POST c'%s' f'%s'", contract.String(), function.String())
	params := o.getParams(int32(decode.Int()))
	transfer := o.getTransfer(int32(decode.Int()))
	delay := decode.Int()
	if delay < -1 {
		o.Panic("invalid delay: %d", delay)
	}
	o.vm.ctx.PostRequest(coretypes.PostRequestParams{
		TargetContractID: contract,
		EntryPoint:       function,
		Params:           params,
		Transfer:         transfer,
		TimeLock:         uint32(delay),
	})
}

func (o *ScContext) getParams(paramsId int32) dict.Dict {
	if paramsId == 0 {
		return dict.New()
	}
	params := o.host.FindObject(paramsId).(*ScDict).kvStore.(dict.Dict)
	params.MustIterate("", func(key kv.Key, value []byte) bool {
		o.Trace("  PARAM '%s'", key)
		return true
	})
	return params
}

func (o *ScContext) getTransfer(transferId int32) coretypes.ColoredBalances {
	if transferId == 0 {
		return cbalances.NewFromMap(nil)
	}
	transfer := make(map[balance.Color]int64)
	transferDict := o.host.FindObject(transferId).(*ScDict).kvStore
	transferDict.MustIterate("", func(key kv.Key, value []byte) bool {
		color, _, err := codec.DecodeColor([]byte(key))
		if err != nil {
			o.Panic(err.Error())
		}
		amount, _, err := codec.DecodeInt64(value)
		if err != nil {
			o.Panic(err.Error())
		}
		o.Trace("  XFER %d '%s'", amount, color.String())
		transfer[color] = amount
		return true
	})
	return cbalances.NewFromMap(transfer)
}
