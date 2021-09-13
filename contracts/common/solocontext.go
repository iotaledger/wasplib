package common

import (
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/contracts/common"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/stretchr/testify/require"
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

// TODO how to pass InitParams to init
func NewSoloContext(t *testing.T, chain *solo.Chain, contract string, onLoad func(), params ...interface{}) *SoloContext {
	if chain == nil {
		chain = common.StartChain(t, "chain1")
	}
	ctx := &SoloContext{contract: contract, Chain: chain}
	ctx.Err = DeployWasmContractByName(chain, contract, params...)
	if ctx.Err != nil {
		return ctx
	}
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

func NewSoloContract(t *testing.T, contract string, onLoad func(), params ...interface{}) *SoloContext {
	ctx := NewSoloContext(t, nil, contract, onLoad, params...)
	require.NoError(t, ctx.Err)
	return ctx
}

func (s *SoloContext) Address() wasmlib.ScAddress {
	if s.keyPair == nil {
		return s.ScAddress(s.Chain.OriginatorAddress)
	}
	return s.ScAddress(ledgerstate.NewED25519Address(s.keyPair.PublicKey))
}

func (s *SoloContext) Balance(agent *SoloAgent, color ...wasmlib.ScColor) int64 {
	var account *iscp.AgentID
	if agent == nil {
		account = iscp.NewAgentID(s.Chain.ChainID.AsAddress(), iscp.Hn(s.contract))
	} else {
		account = iscp.NewAgentID(agent.address, 0)
	}
	balances := s.Chain.GetAccountBalance(account)
	switch len(color) {
	case 0:
		return int64(balances.Get(colored.IOTA))
	case 1:
		col, err := colored.ColorFromBytes(color[0].Bytes())
		require.NoError(s.Chain.Env.T, err)
		return int64(balances.Get(col))
	default:
		require.Fail(s.Chain.Env.T, "too many color arguments")
		return 0
	}
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

func (s *SoloContext) ScAgentID(agentID iscp.AgentID) wasmlib.ScAgentID {
	return wasmlib.NewScAgentIDFromBytes(agentID.Bytes())
}

func (s *SoloContext) ScColor(color colored.Color) wasmlib.ScColor {
	return wasmlib.NewScColorFromBytes(color.Bytes())
}

func (s *SoloContext) ScChainID(chainID iscp.ChainID) wasmlib.ScChainID {
	return wasmlib.NewScChainIDFromBytes(chainID.Bytes())
}

func (s *SoloContext) ScHash(hash hashing.HashValue) wasmlib.ScHash {
	return wasmlib.NewScHashFromBytes(hash.Bytes())
}

func (s *SoloContext) ScHname(hname iscp.Hname) wasmlib.ScHname {
	return wasmlib.NewScHnameFromBytes(hname.Bytes())
}

func (s *SoloContext) ScRequestID(requestID iscp.RequestID) wasmlib.ScRequestID {
	return wasmlib.NewScRequestIDFromBytes(requestID.Bytes())
}

func (s *SoloContext) Sign(agent *SoloAgent) *SoloContext {
	s.keyPair = agent.pair
	return s
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
