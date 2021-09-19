// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmsolo

import (
	"flag"
	"testing"
	"time"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/iotaledger/wasplib/packages/vm/wasmvm"
	"github.com/stretchr/testify/require"
)

const (
	Debug      = true
	StackTrace = false
	TraceHost  = false
)

type SoloContext struct {
	Chain     *solo.Chain
	Convertor SoloConvertor
	Err       error
	keyPair   *ed25519.KeyPair
	scName    string
	wasmHost  wasmhost.WasmHost
}

var (
	_        wasmlib.ScFuncCallContext = &SoloContext{}
	_        wasmlib.ScViewCallContext = &SoloContext{}
	SoloHost wasmlib.ScHost
	GoDebug  = flag.Bool("godebug", false, "debug go smart contract code")
)

// NewSoloContext can be used to create a SoloContext associated with a smart contract with
// minimal information and will verify successful creation before returning ctx.
// It will start a default chain "chain1" before initializing the smart contract.
// It takes the scName and onLoad() function for the contract and optionally
// an init.Func initialized with the parameters to pass to the contract's init() function.
// Unless you want to use a different chain than the default "chain1" this will be your
// function of choice to set up a smart contract for your tests
func NewSoloContext(t *testing.T, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	ctx := NewSoloContextForChain(t, nil, scName, onLoad, init...)
	require.NoError(t, ctx.Err)
	return ctx
}

// NewSoloContextForChain can be used to create a SoloContext associated with a smart contract on a particular chain.
// When chain is nil the function will start a default chain "chain1" before initializing the smart contract.
// It takes the scName and onLoad() function for the contract and optionally
// an init.Func initialized with the parameters to pass to the contract's init() function.
// You can check for any error that occurred by checking the ctx.Err member.
func NewSoloContextForChain(t *testing.T, chain *solo.Chain, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	if chain == nil {
		chain = StartChain(t, "chain1")
	}

	ctx := &SoloContext{scName: scName, Chain: chain}
	ctx.Err = deploy(chain, scName, onLoad, init...)
	if ctx.Err != nil {
		return ctx
	}

	ctx.wasmHost.Init(chain.Log)
	ctx.wasmHost.TrackObject(wasmproc.NewNullObject(&ctx.wasmHost.KvStoreHost))
	ctx.wasmHost.TrackObject(NewSoloScContext(ctx))
	if SoloHost == nil {
		SoloHost = wasmlib.ConnectHost(&ctx.wasmHost)
	}
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	onLoad()
	return ctx
}

func deploy(chain *solo.Chain, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) error {
	var params []interface{}
	if len(init) != 0 {
		params = init[0].Params()
	}

	if SoloHost == nil {
		retDict, err := chain.CallView(root.Contract.Name, root.FuncFindContract.Name,
			root.ParamHname, iscp.Hn(scName),
		)
		if err != nil {
			return err
		}
		retBin, err := retDict.Get(root.ParamContractFound)
		if err != nil {
			return err
		}
		if len(retBin) == 1 && retBin[0] == 0xff {
			// a contract with that name already exists: probably native code
			return nil
		}
	}

	SoloHost = nil
	if *GoDebug {
		wasmproc.GoWasmVM = wasmvm.NewWasmGoVM(scName, onLoad)
		hprog, err := chain.UploadWasm(nil, []byte("go:"+scName))
		if err != nil {
			return err
		}
		return chain.DeployContract(nil, scName, hprog, params...)
	}

	// wasmproc.GoWasmVM = NewWasmTimeJavaVM()
	// wasmproc.GoWasmVM = NewWartVM()
	// wasmproc.GoWasmVM = NewWasmerVM()
	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	return chain.DeployWasmContract(nil, scName, wasmFile, params...)
}

// StartChain starts a new chain named chainName.
func StartChain(t *testing.T, chainName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	// wasmhost.ExtendedHostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	return env.NewChain(nil, chainName)
}

// AdvanceClockBy is used to forward the internal clock by the provided step duration.
func (ctx *SoloContext) AdvanceClockBy(step time.Duration) {
	ctx.Chain.Env.AdvanceClockBy(step)
}

// Balance returns the account balance of the specified agent on the chain associated with ctx.
// To return the account balance of the contract associated with ctx use nil as agent.
// The optional color parameter can be used to retrieve the balance for the specific color.
// When color is omitted, wasmlib.IOTA is assumed.
func (ctx *SoloContext) Balance(agent *SoloAgent, color ...wasmlib.ScColor) int64 {
	var account *iscp.AgentID
	if agent == nil {
		account = iscp.NewAgentID(ctx.Chain.ChainID.AsAddress(), iscp.Hn(ctx.scName))
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

// CanCallFunc is a dummy function that is required to use SoloContext as an ScFuncCallContext
func (ctx *SoloContext) CanCallFunc() {
	panic("SoloContext.CanCallFunc")
}

// CanCallView is a dummy function that is required to use SoloContext as an ScViewCallContext
func (ctx *SoloContext) CanCallView() {
	panic("SoloContext,CanCallView")
}

// ContractExists checks to see if the contract named scName exists in the chain associated with ctx.
func (ctx *SoloContext) ContractExists(scName string) error {
	_, err := ctx.Chain.FindContract(scName)
	return err
}

func (ctx *SoloContext) Host() wasmlib.ScHost {
	return nil
}

// NewSoloAgent creates a new SoloAgent with solo.Saldo tokens in its address
func (ctx *SoloContext) NewSoloAgent() *SoloAgent {
	return NewSoloAgent(ctx.Chain.Env)
}

// Sign is used to force a different agent for signing a Post() request
func (ctx *SoloContext) Sign(agent *SoloAgent) *SoloContext {
	ctx.keyPair = agent.pair
	return ctx
}

// Transfer creates a new ScTransfers proxy
func (ctx *SoloContext) Transfer() wasmlib.ScTransfers {
	return wasmlib.NewScTransfers()
}

// WaitForPendingRequests waits for expectedRequests pending requests to be processed.
// The function will wait for maxWait (default 5 seconds) duration before giving up with a timeout.
// The function returns the false in case of a timeout.
func (ctx *SoloContext) WaitForPendingRequests(expectedRequests int, maxWait ...time.Duration) bool {
	_ = wasmlib.ConnectHost(SoloHost)
	info := ctx.Chain.MempoolInfo()
	result := ctx.Chain.WaitForRequestsThrough(expectedRequests+info.OutPoolCounter, maxWait...)
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	return result
}
