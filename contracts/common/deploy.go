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
	Debug      = false
	StackTrace = false
	TraceHost  = false
)

//TODO update contracts/readme
//TODO figure out how to interrupt wasmtime VM


const WasmRunner = 1

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

func StartChainAndDeployWasmContractByName(t *testing.T, scName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	CreatorWallet = env.NewSignatureSchemeWithFunds()
	chain := env.NewChain(CreatorWallet, "chain1")
	ContractId = coretypes.NewContractID(chain.ChainID, coretypes.Hn(scName))
	ContractAccount = coretypes.NewAgentIDFromContractID(ContractId)

	if WasmRunner == 1 {
		wasmproc.GoWasmVM = NewWasmGoVM(ScForGoVM)
		hprog, err := chain.UploadWasm(CreatorWallet, []byte("go:"+scName))
		require.NoError(t, err)
		err = chain.DeployContract(CreatorWallet, scName, hprog)
		require.NoError(t, err)
		return chain
	}

	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	err := chain.DeployWasmContract(CreatorWallet, scName, wasmFile)
	require.NoError(t, err)
	return chain
}
