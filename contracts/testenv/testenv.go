// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testenv

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/contracts/rust/dividend"
	"github.com/iotaledger/wasp/contracts/rust/donatewithfeedback"
	"github.com/iotaledger/wasp/contracts/rust/erc20"
	"github.com/iotaledger/wasp/contracts/rust/fairauction"
	"github.com/iotaledger/wasp/contracts/rust/fairroulette"
	"github.com/iotaledger/wasp/contracts/rust/helloworld"
	"github.com/iotaledger/wasp/contracts/rust/inccounter"
	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/contracts/rust/tokenregistry"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const (
	Debug      = false
	StackTrace = false
	TraceHost  = false

	WasmRunnerRust     = 0 // run default Rust Wasm code
	WasmRunnerGo       = 1 // run Go Wasm code instead of Rust Wasm code
	WasmRunnerGoDirect = 2 // run Go code directly, without using Wasm
)

//TODO update contracts/readme
//TODO remove hardcoded dependency in ScForGoVM

var WasmRunner = 0

var ScForGoVM = map[string]func(){
	"dividend":           dividend.OnLoad,
	"donatewithfeedback": donatewithfeedback.OnLoad,
	"erc20":              erc20.OnLoad,
	"fairauction":        fairauction.OnLoad,
	"fairroulette":       fairroulette.OnLoad,
	"helloworld":         helloworld.OnLoad,
	"inccounter":         inccounter.OnLoad,
	"testcore":           testcore.OnLoad,
	"tokenregistry":      tokenregistry.OnLoad,
}

type TestEnv struct {
	Chain           *solo.Chain
	ContractAccount coretypes.AgentID
	ContractId      coretypes.ContractID
	CreatorAgentId  coretypes.AgentID
	CreatorWallet   signaturescheme.SignatureScheme
	Env             *solo.Solo
	host            wasmlib.ScHost
	ScName          string
	req             *solo.CallParams
	T               *testing.T
	wallets         []signaturescheme.SignatureScheme
}

func NewTestEnv(t *testing.T, scName string) *TestEnv {
	wasmhost.HostTracing = TraceHost
	te := &TestEnv{T: t, ScName: scName}
	te.Env = solo.New(t, Debug, StackTrace)
	te.CreatorWallet = te.Env.NewSignatureSchemeWithFunds()
	te.CreatorAgentId = coretypes.NewAgentIDFromAddress(te.CreatorWallet.Address())
	te.Chain = te.Env.NewChain(te.CreatorWallet, "chain1")
	err := DeployGoContract(te.Chain, te.CreatorWallet, scName, scName)
	require.NoError(te.T, err)
	te.ContractId = coretypes.NewContractID(te.Chain.ChainID, coretypes.Hn(scName))
	te.ContractAccount = coretypes.NewAgentIDFromContractID(te.ContractId)
	return te
}

// returns the agent id of a wallet with preloaded funds for the agent with specified index
func (te *TestEnv) Agent(index int) coretypes.AgentID {
	return coretypes.NewAgentIDFromAddress(te.Wallet(index).Address())
}

// calls view on current contract
func (te *TestEnv) CallView(funcName string, params ...interface{}) dict.Dict {
	if te.host != nil {
		wasmlib.ConnectHost(te.host)
	}
	ret, err := te.Chain.CallView(te.ScName, funcName, filterKeys(params...)...)
	require.NoError(te.T, err)
	return ret
}

// sets up request for func or view on current contract
func (te *TestEnv) NewCallParams(funcName string, params ...interface{}) *TestEnv {
	te.req = solo.NewCallParams(te.ScName, funcName, filterKeys(params...)...)
	return te
}

func (te *TestEnv) post(iotas int64, scheme []signaturescheme.SignatureScheme) (dict.Dict, error) {
	if te.host != nil {
		wasmlib.ConnectHost(te.host)
	}
	if iotas != 0 {
		te.WithTransfer(balance.ColorIOTA, iotas)
	}
	sigScheme := signaturescheme.SignatureScheme(nil)
	if len(scheme) > 0 {
		sigScheme = scheme[0]
	}
	ret, err := te.Chain.PostRequest(te.req, sigScheme)
	return ret, err
}

// posts the func or view request, expecting to succeed
func (te *TestEnv) Post(iotas int64, scheme ...signaturescheme.SignatureScheme) dict.Dict {
	ret, err := te.post(iotas, scheme)
	require.NoError(te.T, err)
	return ret
}

// posts the func or view request, expecting to fail
func (te *TestEnv) PostFail(iotas int64, scheme ...signaturescheme.SignatureScheme) error {
	_, err := te.post(iotas, scheme)
	require.Error(te.T, err)
	return err
}

// convert call result to wasmlib ScImmutableMap
func (te *TestEnv) Results(dict kv.KVStore) wasmlib.ScImmutableMap {
	return te.ScImmutableMap(wasmhost.KeyResults, dict)
}

// convert K/V store to wasmlib ScImmutableMap
func (te *TestEnv) ScImmutableMap(keyId int32, kvStore kv.KVStore) wasmlib.ScImmutableMap {
	logger := testutil.NewLogger(te.T, "04:05.000")
	host := &wasmhost.KvStoreHost{}
	null := wasmproc.NewNullObject(host)
	root := wasmproc.NewScDictFromKvStore(host, kvStore)
	host.Init(null, root, logger)
	root.InitObj(1, keyId, root)
	logger.Infof("Direct access to %s", host.GetKeyStringFromId(keyId))
	oldHost := wasmlib.ConnectHost(host)
	if te.host == nil {
		te.host = oldHost
	}
	return wasmlib.Root.Immutable()
}

// retrieve entire state of contract as ScImmutableMap
func (te *TestEnv) State() wasmlib.ScImmutableMap {
	ret := te.CallView("copy_all_state")
	return te.ScImmutableMap(wasmhost.KeyState, ret)
}

// process all requests until request backlog is empty
func (te *TestEnv) WaitForEmptyBacklog() {
	te.Chain.WaitForEmptyBacklog()
}

// returns a wallet with preloaded funds for the agent with specified index
func (te *TestEnv) Wallet(index int) signaturescheme.SignatureScheme {
	require.True(te.T, index <= len(te.wallets), "invalid wallet index")
	if index == len(te.wallets) {
		te.wallets = append(te.wallets, te.Env.NewSignatureSchemeWithFunds())
	}
	return te.wallets[index]
}

// add a single transfer to request
func (te *TestEnv) WithTransfer(color balance.Color, amount int64) *TestEnv {
	te.req.WithTransfer(color, amount)
	return te
}

// add multiple transfers to request
func (te *TestEnv) WithTransfers(transfer map[balance.Color]int64) *TestEnv {
	te.req.WithTransfers(transfer)
	return te
}

// deploy the specified contract on the chain
func DeployGoContract(chain *solo.Chain, sigScheme signaturescheme.SignatureScheme, name string, contractName string, params ...interface{}) error {
	if WasmRunner == WasmRunnerGoDirect {
		wasmproc.GoWasmVM = wasmhost.NewWasmGoVM(ScForGoVM)
		hprog, err := chain.UploadWasm(sigScheme, []byte("go:"+contractName))
		if err != nil {
			return err
		}
		return chain.DeployContract(sigScheme, name, hprog, filterKeys(params...)...)
	}

	wasmFile := contractName + "_bg.wasm"
	if WasmRunner == WasmRunnerGo {
		wasmFile = strings.Replace(wasmFile, "_bg", "_go", -1)
	}
	wasmFile = util.LocateFile(wasmFile, contractName + "/pkg")
	return chain.DeployWasmContract(sigScheme, name, wasmFile, filterKeys(params...)...)
}

// filters wasmlib.Key parameters and replaces them with their proper string equivalent
func filterKeys(params ...interface{}) []interface{} {
	for i, param := range params {
		switch param.(type) {
		case wasmlib.Key:
			params[i] = string(param.(wasmlib.Key))
		}
	}
	return params
}
