package govm

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/client"
	"github.com/iotaledger/wasplib/contracts/dividend"
	"github.com/iotaledger/wasplib/contracts/donatewithfeedback"
	"github.com/iotaledger/wasplib/contracts/dummy"
	"github.com/iotaledger/wasplib/contracts/erc20"
	"github.com/iotaledger/wasplib/contracts/example1"
	"github.com/iotaledger/wasplib/contracts/fairauction"
	"github.com/iotaledger/wasplib/contracts/fairroulette"
	"github.com/iotaledger/wasplib/contracts/helloworld"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/contracts/testcore"
	"github.com/iotaledger/wasplib/contracts/tokenregistry"
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

var WasmRunner = 0

var ScForGoVM = map[string]func(){
	"dividend":           dividend.OnLoad,
	"donatewithfeedback": donatewithfeedback.OnLoad,
	"dummy":              dummy.OnLoad,
	"erc20":              erc20.OnLoad,
	"example1":           example1.OnLoad,
	"fairauction":        fairauction.OnLoad,
	"fairroulette":       fairroulette.OnLoad,
	"helloworld":         helloworld.OnLoad,
	"inccounter":         inccounter.OnLoad,
	"testcore":           testcore.OnLoad,
	"tokenregistry":      tokenregistry.OnLoad,
}

type TestEnv struct {
	Chain  *solo.Chain
	Env    *solo.Solo
	ScName string
	req    *solo.CallParams
	T      *testing.T
}

func NewTestEnv(t *testing.T, scName string) *TestEnv {
	wasmhost.HostTracing = TraceHost
	te := &TestEnv{T: t, ScName: scName}
	te.Env = solo.New(t, Debug, StackTrace)
	te.Chain = te.Env.NewChain(nil, "chain1")
	err := DeployGoContract(te.Chain, nil, scName, scName)
	require.NoError(te.T, err)
	return te
}

func (te *TestEnv) CallView(funcName string, params ...interface{}) dict.Dict {
	ret, err := te.Chain.CallView(te.ScName, funcName, params...)
	require.NoError(te.T, err)
	return ret
}

func (te *TestEnv) NewCallParams(funcName string, params ...interface{}) *TestEnv {
	te.req = solo.NewCallParams(te.ScName, funcName, params...)
	return te
}

func (te *TestEnv) WithTransfer(color balance.Color, amount int64) *TestEnv {
	te.req.WithTransfers(map[balance.Color]int64{color: amount})
	return te
}

func (te *TestEnv) WithTransfers(transfer map[balance.Color]int64) *TestEnv {
	te.req.WithTransfers(transfer)
	return te
}

func (te *TestEnv) post(iotas int64, scheme []signaturescheme.SignatureScheme) (dict.Dict, error) {
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

func (te *TestEnv) Post(iotas int64, scheme ...signaturescheme.SignatureScheme) dict.Dict {
	ret, err := te.post(iotas, scheme)
	require.NoError(te.T, err)
	return ret
}

func (te *TestEnv) PostFail(iotas int64, scheme ...signaturescheme.SignatureScheme) error {
	_, err := te.post(iotas, scheme)
	require.Error(te.T, err)
	return err
}

func (te *TestEnv) State() client.ScImmutableMap {
	ret := te.CallView("copy_all_state")
	return GetClientMap(te.T, wasmhost.KeyResults, ret)
}

func (te *TestEnv) WaitForEmptyBacklog() {
	te.Chain.WaitForEmptyBacklog()
}

func DeployGoContract(chain *solo.Chain, sigScheme signaturescheme.SignatureScheme, name string, contractName string, params ...interface{}) error {
	if WasmRunner == WasmRunnerGoDirect {
		wasmproc.GoWasmVM = NewGoVM(ScForGoVM)
		hprog, err := chain.UploadWasm(sigScheme, []byte("go:"+contractName))
		if err != nil {
			return err
		}
		return chain.DeployContract(sigScheme, name, hprog, params...)
	}

	wasmFile := contractName + "_bg.wasm"
	if WasmRunner == WasmRunnerGo {
		wasmFile = strings.Replace(wasmFile, "_bg", "_go", -1)
	}
	wasmFile = wasmhost.WasmPath(wasmFile)
	return chain.DeployWasmContract(sigScheme, name, wasmFile, params...)
}

func GetClientMap(t *testing.T, keyId int32, kvStore kv.KVStore) client.ScImmutableMap {
	logger := testutil.NewLogger(t, "04:05.000")
	host := &wasmhost.KvStoreHost{}
	null := wasmproc.NewNullObject(host)
	root := wasmproc.NewScDictFromKvStore(host, kvStore)
	host.Init(null, root, logger)
	root.InitObj(1, keyId, root)
	logger.Info("Direct access to %s", host.GetKeyStringFromId(keyId))
	client.ConnectHost(host)
	return client.Root.Immutable()
}
