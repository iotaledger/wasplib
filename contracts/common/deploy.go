package common

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/contracts/rust/dividend"
	"github.com/iotaledger/wasplib/contracts/rust/donatewithfeedback"
	"github.com/iotaledger/wasplib/contracts/rust/erc20"
	"github.com/iotaledger/wasplib/contracts/rust/fairauction"
	"github.com/iotaledger/wasplib/contracts/rust/fairroulette"
	"github.com/iotaledger/wasplib/contracts/rust/helloworld"
	"github.com/iotaledger/wasplib/contracts/rust/inccounter"
	"github.com/iotaledger/wasplib/contracts/rust/testcore"
	"github.com/iotaledger/wasplib/contracts/rust/testwasmlib"
	"github.com/iotaledger/wasplib/contracts/rust/tokenregistry"
	"github.com/stretchr/testify/require"
)

const (
	Debug      = true
	StackTrace = false
	TraceHost  = false
)

//TODO update contracts/readme
//TODO figure out how to interrupt wasmtime VM

// WasmRunner set to 1 to run/debug go code directly instead of running Rust or TinyGo Wasm code
const WasmRunner = 0

//TODO remove hardcoded dependency
var ScForGoVM = map[string]func(){
	"dividend":           dividend.OnLoad,
	"donatewithfeedback": donatewithfeedback.OnLoad,
	"erc20":              erc20.OnLoad,
	"fairauction":        fairauction.OnLoad,
	"fairroulette":       fairroulette.OnLoad,
	"helloworld":         helloworld.OnLoad,
	"inccounter":         inccounter.OnLoad,
	"testcore":           testcore.OnLoad,
	"testwasmlib":        testwasmlib.OnLoad,
	"tokenregistry":      tokenregistry.OnLoad,
}

func DeployWasmContractByName(chain *solo.Chain, scName string, params ...interface{}) error {
	if WasmRunner == 1 {
		wasmproc.GoWasmVM = NewWasmGoVM(ScForGoVM)
		hprog, err := chain.UploadWasm(nil, []byte("go:"+scName))
		if err != nil {
			return err
		}
		return chain.DeployContract(nil, scName, hprog, params...)
	}

	//wasmproc.GoWasmVM = NewWasmTimeJavaVM()
	//wasmproc.GoWasmVM = NewWartVM()
	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	return chain.DeployWasmContract(nil, scName, wasmFile, params...)
}

func StartChain(t *testing.T, scName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	//wasmhost.ExtendedHostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	return env.NewChain(nil, "chain1")
}

func StartChainAndDeployWasmContractByName(t *testing.T, scName string, params ...interface{}) *solo.Chain {
	chain := StartChain(t, scName)
	err := DeployWasmContractByName(chain, scName, params...)
	require.NoError(t, err)
	return chain
}
