package common

import (
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/chainid"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

type SoloContext struct {
	Chain    *solo.Chain
	contract string
	Err      error
	keyPair  *ed25519.KeyPair
	wasmHost wasmhost.WasmHost
}

var (
	_        wasmlib.ScFuncCallContext = &SoloContext{}
	_        wasmlib.ScViewCallContext = &SoloContext{}
	soloHost wasmlib.ScHost
)

func NewSoloContext(contract string, onLoad func(), chain *solo.Chain) *SoloContext {
	ctx := &SoloContext{contract: contract, Chain: chain}
	ctx.wasmHost.Init(chain.Log)
	ctx.wasmHost.TrackObject(wasmproc.NewNullObject(&ctx.wasmHost.KvStoreHost))
	ctx.wasmHost.TrackObject(NewSoloScContext(ctx))
	if soloHost == nil {
		soloHost = wasmlib.ConnectHost(&ctx.wasmHost)
	}
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	onLoad()
	return ctx
}

func (s *SoloContext) Address() wasmlib.ScAddress {
	if s.keyPair == nil {
		return s.ScAddress(s.Chain.OriginatorAddress)
	}
	return s.ScAddress(ledgerstate.NewED25519Address(s.keyPair.PublicKey))
}

func (s *SoloContext) CanCallFunc() {
	panic("CanCallFunc")
}

func (s *SoloContext) CanCallView() {
	panic("CanCallView")
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

func (s *SoloContext) SignWith(keyPair *ed25519.KeyPair) *SoloContext {
	s.keyPair = keyPair
	return s
}

func (s *SoloContext) Transfer() wasmlib.ScTransfers {
	return wasmlib.NewScTransfers()
}

func (s *SoloContext) WaitForRequestsThrough(numReq int, maxWait ...time.Duration) bool {
	_ = wasmlib.ConnectHost(soloHost)
	result := s.Chain.WaitForRequestsThrough(numReq, maxWait...)
	_ = wasmlib.ConnectHost(&s.wasmHost)
	return result
}
