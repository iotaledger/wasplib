package common

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/coretypes"
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
	"github.com/iotaledger/wasplib/contracts/rust/tokenregistry"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	Debug      = true
	StackTrace = true
	TraceHost  = true
)

//TODO update contracts/readme
//TODO figure out how to interrupt wasmtime VM

// set to 1 to run/debug go code directly instead of running Rust or TinyGo Wasm code
const WasmRunner = 0

var (
	ContractAccount coretypes.AgentID
	ContractId      coretypes.ContractID
	CreatorWallet   signaturescheme.SignatureScheme

	//TODO remove hardcoded dependency
	ScForGoVM = map[string]func(){
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
)

func DeployWasmContractByName(chain *solo.Chain, scName string, params ...interface{}) error {
	if WasmRunner == 1 {
		wasmproc.GoWasmVM = NewWasmGoVM(ScForGoVM)
		hprog, err := chain.UploadWasm(CreatorWallet, []byte("go:"+scName))
		if err != nil { return err }
		return chain.DeployContract(CreatorWallet, scName, hprog, params...)
	}

	wasmproc.GoWasmVM = NewWasmTimeJavaVM()
	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	return chain.DeployWasmContract(CreatorWallet, scName, wasmFile, params...)
}

func StartChain(t *testing.T, scName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	CreatorWallet = env.NewSignatureSchemeWithFunds()
	chain := env.NewChain(CreatorWallet, "chain1")
	ContractId = coretypes.NewContractID(chain.ChainID, coretypes.Hn(scName))
	ContractAccount = coretypes.NewAgentIDFromContractID(ContractId)
	return chain
}

func StartChainAndDeployWasmContractByName(t *testing.T, scName string, params ...interface{}) *solo.Chain {
	chain := StartChain(t, scName)
	err := DeployWasmContractByName(chain, scName, params...)
	require.NoError(t, err)
	return chain
}