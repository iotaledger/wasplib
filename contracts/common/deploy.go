package common

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	Debug      = true
	StackTrace = true
	TraceHost  = true
)

var (
	ContractAccount coretypes.AgentID
	ContractId      coretypes.ContractID
	CreatorWallet   signaturescheme.SignatureScheme
)

func DeployContract(t *testing.T, scName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	CreatorWallet = env.NewSignatureSchemeWithFunds()
	chain := env.NewChain(CreatorWallet, "chain1")
	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	err := chain.DeployWasmContract(CreatorWallet, scName, wasmFile)
	require.NoError(t, err)
	ContractId = coretypes.NewContractID(chain.ChainID, coretypes.Hn(scName))
	ContractAccount = coretypes.NewAgentIDFromContractID(ContractId)
	return chain
}
