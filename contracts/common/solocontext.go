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

func NewSoloContext(t *testing.T, chain *solo.Chain, contract string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	if chain == nil {
		chain = common.StartChain(t, "chain1")
	}
	ctx := &SoloContext{contract: contract, Chain: chain}
	ctx.Err = deploy(chain, contract, init...)
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

func NewSoloContract(t *testing.T, contract string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	ctx := NewSoloContext(t, nil, contract, onLoad, init...)
	require.NoError(t, ctx.Err)
	return ctx
}

func deploy(chain *solo.Chain, contract string, init ...*wasmlib.ScInitFunc) error {
	if len(init) == 0 {
		return DeployWasmContractByName(chain, contract)
	}
	initFunc := init[0]
	return DeployWasmContractByName(chain, contract, initFunc.Params()...)
}

func (ctx *SoloContext) Address() wasmlib.ScAddress {
	if ctx.keyPair == nil {
		return ctx.ScAddress(ctx.Chain.OriginatorAddress)
	}
	return ctx.ScAddress(ledgerstate.NewED25519Address(ctx.keyPair.PublicKey))
}

func (ctx *SoloContext) Balance(agent *SoloAgent, color ...wasmlib.ScColor) int64 {
	var account *iscp.AgentID
	if agent == nil {
		account = iscp.NewAgentID(ctx.Chain.ChainID.AsAddress(), iscp.Hn(ctx.contract))
	} else {
		account = iscp.NewAgentID(agent.address, 0)
	}
	balances := ctx.Chain.GetAccountBalance(account)
	switch len(color) {
	case 0:
		return int64(balances.Get(colored.IOTA))
	case 1:
		col, err := colored.ColorFromBytes(color[0].Bytes())
		require.NoError(ctx.Chain.Env.T, err)
		return int64(balances.Get(col))
	default:
		require.Fail(ctx.Chain.Env.T, "too many color arguments")
		return 0
	}
}

func (ctx *SoloContext) CanCallFunc() {
	panic("CanCallFunc")
}

func (ctx *SoloContext) CanCallView() {
	panic("CanCallView")
}

func (ctx *SoloContext) Host() wasmlib.ScHost {
	return nil
}

func (ctx *SoloContext) NewSoloAgent() *SoloAgent {
	return NewSoloAgent(ctx.Chain.Env)
}

func (ctx *SoloContext) ScAddress(address ledgerstate.Address) wasmlib.ScAddress {
	return wasmlib.NewScAddressFromBytes(address.Bytes())
}

func (ctx *SoloContext) ScAgentID(agentID iscp.AgentID) wasmlib.ScAgentID {
	return wasmlib.NewScAgentIDFromBytes(agentID.Bytes())
}

func (ctx *SoloContext) ScColor(color colored.Color) wasmlib.ScColor {
	return wasmlib.NewScColorFromBytes(color.Bytes())
}

func (ctx *SoloContext) ScChainID(chainID iscp.ChainID) wasmlib.ScChainID {
	return wasmlib.NewScChainIDFromBytes(chainID.Bytes())
}

func (ctx *SoloContext) ScHash(hash hashing.HashValue) wasmlib.ScHash {
	return wasmlib.NewScHashFromBytes(hash.Bytes())
}

func (ctx *SoloContext) ScHname(hname iscp.Hname) wasmlib.ScHname {
	return wasmlib.NewScHnameFromBytes(hname.Bytes())
}

func (ctx *SoloContext) ScRequestID(requestID iscp.RequestID) wasmlib.ScRequestID {
	return wasmlib.NewScRequestIDFromBytes(requestID.Bytes())
}

func (ctx *SoloContext) Sign(agent *SoloAgent) *SoloContext {
	ctx.keyPair = agent.pair
	return ctx
}

func (ctx *SoloContext) SignWith(keyPair *ed25519.KeyPair) *SoloContext {
	ctx.keyPair = keyPair
	return ctx
}

func (ctx *SoloContext) Transfer() wasmlib.ScTransfers {
	return wasmlib.NewScTransfers()
}

func (ctx *SoloContext) WaitForRequestsThrough(numReq int, maxWait ...time.Duration) bool {
	_ = wasmlib.ConnectHost(soloHost)
	result := ctx.Chain.WaitForRequestsThrough(numReq, maxWait...)
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	return result
}
