package common

import (
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

type SoloScContext struct {
	wasmproc.ScContext
	ctx *SoloContext
}

func NewSoloScContext(ctx *SoloContext) *SoloScContext {
	return &SoloScContext{ScContext: *wasmproc.NewScContext(nil, &ctx.wasmHost.KvStoreHost), ctx: ctx}
}

func (o *SoloScContext) Exists(keyID, typeID int32) bool {
	return o.GetTypeID(keyID) > 0
}

func (o *SoloScContext) GetBytes(keyID, typeID int32) []byte {
	switch keyID {
	case wasmhost.KeyChainID:
		return o.ctx.Chain.ChainID.Bytes()
	default:
		o.InvalidKey(keyID)
		return nil
	}
}

func (o *SoloScContext) GetObjectID(keyID, typeID int32) int32 {
	host := &o.ctx.wasmHost
	return wasmproc.GetMapObjectID(o, keyID, typeID, wasmproc.ObjFactories{
		// wasmhost.KeyBalances:  func() wasmproc.WaspObject { return wasmproc.NewScBalances(o.vm, keyID) },
		wasmhost.KeyExports: func() wasmproc.WaspObject { return wasmproc.NewScExports(host) },
		// wasmhost.KeyIncoming:  func() wasmproc.WaspObject { return wasmproc.NewScBalances(o.vm, keyID) },
		wasmhost.KeyMaps: func() wasmproc.WaspObject { return wasmproc.NewScMaps(&host.KvStoreHost) },
		// wasmhost.KeyMinted:    func() wasmproc.WaspObject { return wasmproc.NewScBalances(o.vm, keyID) },
		// wasmhost.KeyParams:    func() wasmproc.WaspObject { return wasmproc.NewScDict(o.host, o.vm.params()) },
		wasmhost.KeyResults: func() wasmproc.WaspObject { return wasmproc.NewScDict(&host.KvStoreHost, dict.New()) },
		wasmhost.KeyReturn:  func() wasmproc.WaspObject { return wasmproc.NewScDict(&host.KvStoreHost, dict.New()) },
		// wasmhost.KeyState:     func() wasmproc.WaspObject { return wasmproc.NewScDict(o.host, o.vm.state()) },
		// wasmhost.KeyTransfers: func() wasmproc.WaspObject { return wasmproc.NewScTransfers(o.vm) },
		// wasmhost.KeyUtility:   func() wasmproc.WaspObject { return wasmproc.NewScUtility(o.vm) },
	})
}

func (o *SoloScContext) SetBytes(keyID, typeID int32, bytes []byte) {
	switch keyID {
	case wasmhost.KeyCall:
		o.processCall(bytes)
	case wasmhost.KeyPost:
		o.processPost(bytes)
	default:
		o.ScContext.SetBytes(keyID, typeID, bytes)
	}
}

func (o *SoloScContext) processCall(bytes []byte) {
	decode := wasmproc.NewBytesDecoder(bytes)
	contract, err := iscp.HnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	function, err := iscp.HnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	paramsID := decode.Int32()
	transferID := decode.Int32()
	if transferID != 0 {
		o.postSync(contract, function, paramsID, transferID, 0)
		return
	}

	funcName := o.ctx.wasmHost.FunctionFromCode(uint32(function))
	if funcName == "" {
		o.Panic("unknown function")
	}
	o.Tracef("CALL %s.%s", o.ctx.contract, funcName)
	params := o.getParams(paramsID)
	_ = wasmlib.ConnectHost(soloHost)
	res, err := o.ctx.Chain.CallView(o.ctx.contract, funcName, params)
	_ = wasmlib.ConnectHost(&o.ctx.wasmHost)
	o.ctx.Err = err
	if err != nil {
		o.Panic("failed to invoke call: " + err.Error())
	}
	returnID := o.GetObjectID(int32(wasmlib.KeyReturn), wasmlib.TYPE_MAP)
	o.ctx.wasmHost.FindObject(returnID).(*wasmproc.ScDict).SetKvStore(res)
}

func (o *SoloScContext) processPost(bytes []byte) {
	decode := wasmproc.NewBytesDecoder(bytes)
	chainID, err := iscp.ChainIDFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	if !chainID.Equals(&o.ctx.Chain.ChainID) {
		o.Panic("invalid chainID")
	}
	contract, err := iscp.HnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	function, err := iscp.HnameFromBytes(decode.Bytes())
	if err != nil {
		o.Panic(err.Error())
	}
	paramsID := decode.Int32()
	transferID := decode.Int32()
	delay := decode.Int32()
	o.postSync(contract, function, paramsID, transferID, delay)
	//metadata := &iscp.SendMetadata{
	//	TargetContract: contract,
	//	EntryPoint:     function,
	//	Args:           params,
	//}
	//delay := decode.Int32()
	//if delay == 0 {
	//	if !o.vm.ctx.Send(chainID.AsAddress(), transfer, metadata) {
	//		o.Panic("failed to send to %s", chainID.AsAddress().String())
	//	}
	//	return
	//}
	//
	//if delay < -1 {
	//	o.Panic("invalid delay: %d", delay)
	//}
	//
	//timeLock := time.Unix(0, o.vm.ctx.GetTimestamp())
	//timeLock = timeLock.Add(time.Duration(delay) * time.Second)
	//options := iscp.SendOptions{
	//	TimeLock: uint32(timeLock.Unix()),
	//}
	//if !o.vm.ctx.Send(chainID.AsAddress(), transfer, metadata, options) {
	//	o.Panic("failed to send to %s", chainID.AsAddress().String())
	//}
}

func (o *SoloScContext) getParams(paramsID int32) dict.Dict {
	if paramsID == 0 {
		return dict.New()
	}
	params := o.ctx.wasmHost.FindObject(paramsID).(*wasmproc.ScDict).KvStore().(dict.Dict)
	params.MustIterate("", func(key kv.Key, value []byte) bool {
		o.Tracef("  PARAM '%s'", key)
		return true
	})
	return params
}

func (o *SoloScContext) getTransfer(transferID int32) colored.Balances {
	if transferID == 0 {
		return colored.NewBalances()
	}
	transfer := colored.NewBalances()
	transferDict := o.ctx.wasmHost.FindObject(transferID).(*wasmproc.ScDict).KvStore()
	transferDict.MustIterate("", func(key kv.Key, value []byte) bool {
		color, _, err := codec.DecodeColor([]byte(key))
		if err != nil {
			o.Panic(err.Error())
		}
		amount, _, err := codec.DecodeUint64(value)
		if err != nil {
			o.Panic(err.Error())
		}
		o.Tracef("  XFER %d '%s'", amount, color.String())
		transfer[color] = amount
		return true
	})
	return transfer
}

func (o *SoloScContext) postSync(contract, function iscp.Hname, paramsID, transferID, delay int32) {
	if delay != 0 {
		o.Panic("unsupported nonzero delay for SoloContext")
	}
	if contract != iscp.Hn(o.ctx.contract) {
		o.Panic("invalid contract")
	}
	funcName := o.ctx.wasmHost.FunctionFromCode(uint32(function))
	if funcName == "" {
		o.Panic("unknown function")
	}
	o.Tracef("POST %s.%s", o.ctx.contract, funcName)
	params := o.getParams(paramsID)
	req := solo.NewCallParamsFromDic(o.ctx.contract, funcName, params)
	if transferID != 0 {
		transfer := o.getTransfer(transferID)
		req.WithTransfers(transfer)
	}
	_ = wasmlib.ConnectHost(soloHost)
	res, err := o.ctx.Chain.PostRequestSync(req, o.ctx.keyPair)
	_ = wasmlib.ConnectHost(&o.ctx.wasmHost)
	o.ctx.Err = err
	if err != nil {
		return
	}
	returnID := o.GetObjectID(int32(wasmlib.KeyReturn), wasmlib.TYPE_MAP)
	o.ctx.wasmHost.FindObject(returnID).(*wasmproc.ScDict).SetKvStore(res)
}
