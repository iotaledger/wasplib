package common

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/chainid"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

type SoloContext struct {
	address ledgerstate.Address
	chain   *solo.Chain
	Err     error
	host    wasmlib.ScHost
	kvHost  wasmhost.KvStoreHost
}

// implements wasmlib.ScFuncContext interface
var _ wasmlib.ScHostContext = &SoloContext{}

func NewSoloContext(chain *solo.Chain, address ledgerstate.Address) *SoloContext {
	ctx := &SoloContext{chain: chain, address: address}
	ctx.kvHost.Init(chain.Log)
	ctx.kvHost.TrackObject(wasmproc.NewNullObject(&ctx.kvHost))
	// TODO ctx.kvHost.TrackObject(wasmproc.NewScContext(nil, &ctx.kvHost))
	ctx.host = wasmlib.ConnectHost(&ctx.kvHost)
	return ctx
}

func (s *SoloContext) Address() wasmlib.ScAddress {
	return s.ScAddress(s.address)
}

func (s *SoloContext) Host() wasmlib.ScHost {
	return nil
}

func (s *SoloContext) ScAddress(address ledgerstate.Address) wasmlib.ScAddress {
	return wasmlib.NewScAddressFromBytes(address.Bytes())
}

func (s *SoloContext) ScAgentID(agentID coretypes.AgentID) wasmlib.ScAgentID {
	return wasmlib.NewScAgentIDFromBytes(agentID.Bytes())
}

func (s *SoloContext) ScColor(color ledgerstate.Color) wasmlib.ScColor {
	return wasmlib.NewScColorFromBytes(color.Bytes())
}

func (s *SoloContext) ScChainID(chainID chainid.ChainID) wasmlib.ScChainID {
	return wasmlib.NewScChainIDFromBytes(chainID.Bytes())
}

func (s *SoloContext) ScHash(hash hashing.HashValue) wasmlib.ScHash {
	return wasmlib.NewScHashFromBytes(hash.Bytes())
}

func (s *SoloContext) ScHname(hname coretypes.Hname) wasmlib.ScHname {
	return wasmlib.NewScHnameFromBytes(hname.Bytes())
}

func (s *SoloContext) ScRequestID(requestID coretypes.RequestID) wasmlib.ScRequestID {
	return wasmlib.NewScRequestIDFromBytes(requestID.Bytes())
}

func (s *SoloContext) Transfer() wasmlib.ScTransfers {
	return wasmlib.NewScTransfers()
}
