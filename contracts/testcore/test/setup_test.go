package sandbox_tests

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	RUN_GO             = true // run go Wasm code directly, without Wasm
	DEBUG              = false
	RUN_WASM           = false
	WASM_FILE_TESTCORE = "../../../wasm/testcore_bg.wasm"
	WASM_FILE_ERC20    = "../../../wasm/erc20_bg.wasm"
	ERC20_NAME         = "erc20"
	ERC20_SUPPLY       = 100000

	// ERC20 constants
	PARAM_SUPPLY     = "s"
	PARAM_CREATOR    = "c"
	PARAM_ACCOUNT    = "ac"
	PARAM_DELEGATION = "d"
	PARAM_AMOUNT     = "am"
	PARAM_RECIPIENT  = "r"
)

var SandboxSCName = "test_sandbox"

func setupChain(t *testing.T, sigSchemeChain signaturescheme.SignatureScheme) (*solo.Solo, *solo.Chain) {
	env := solo.New(t, DEBUG, false)
	chain := env.NewChain(sigSchemeChain, "ch1")
	return env, chain
}

func setupDeployer(t *testing.T, chain *solo.Chain) signaturescheme.SignatureScheme {
	user := chain.Env.NewSignatureSchemeWithFunds()
	chain.Env.AssertAddressBalance(user.Address(), balance.ColorIOTA, testutil.RequestFundsAmount)

	req := solo.NewCall(root.Interface.Name, root.FuncGrantDeploy,
		root.ParamDeployer, coretypes.NewAgentIDFromAddress(user.Address()),
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	return user
}

func setupTestSandboxSC(t *testing.T, chain *solo.Chain, user signaturescheme.SignatureScheme) (coretypes.ContractID, int64) {
	var err error
	var extraToken int64
	if RUN_GO {
		err = govm.DeployGoContract(chain, user, SandboxSCName, "testcore")
		extraToken = 1
	} else if RUN_WASM {
		err = chain.DeployWasmContract(user, SandboxSCName, WASM_FILE_TESTCORE)
		extraToken = 1
	} else {
		err = chain.DeployContract(user, SandboxSCName, test_sandbox_sc.Interface.ProgramHash)
		extraToken = 0
	}
	require.NoError(t, err)

	deployed := coretypes.NewContractID(chain.ChainID, coretypes.Hn(test_sandbox_sc.Interface.Name))
	req := solo.NewCall(SandboxSCName, test_sandbox_sc.FuncDoNothing)
	_, err = chain.PostRequest(req, user)
	require.NoError(t, err)
	t.Logf("deployed test_sandbox'%s': %s", SandboxSCName, coretypes.Hn(SandboxSCName))
	return deployed, extraToken
}

func setupERC20(t *testing.T, chain *solo.Chain, user signaturescheme.SignatureScheme) coretypes.ContractID {
	var err error
	if !(RUN_WASM || RUN_GO) {
		// only wasm test
		t.SkipNow()
	}
	var userAgentID coretypes.AgentID
	if user == nil {
		userAgentID = chain.OriginatorAgentID
	} else {
		userAgentID = coretypes.NewAgentIDFromAddress(user.Address())
	}
	if RUN_GO {
		err = govm.DeployGoContract(chain, user, ERC20_NAME, ERC20_NAME,
			PARAM_SUPPLY, 1000000,
			PARAM_CREATOR, userAgentID,
		)
	}else {
		err = chain.DeployWasmContract(user, ERC20_NAME, WASM_FILE_ERC20,
			PARAM_SUPPLY, 1000000,
			PARAM_CREATOR, userAgentID,
		)
	}
	require.NoError(t, err)

	deployed := coretypes.NewContractID(chain.ChainID, coretypes.Hn(test_sandbox_sc.Interface.Name))
	t.Logf("deployed erc20'%s': %s --  %s", ERC20_NAME, coretypes.Hn(ERC20_NAME), deployed)
	return deployed
}

func TestSetup1(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)
}

func TestSetup2(t *testing.T) {
	_, chain := setupChain(t, nil)
	user := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user)
}

func TestSetup3(t *testing.T) {
	_, chain := setupChain(t, nil)
	user := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user)
	setupERC20(t, chain, user)
}
