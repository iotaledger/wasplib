package test

import (
	"fmt"
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	DEBUG        = true
	ERC20_NAME   = "erc20"
	ERC20_SUPPLY = 100000

	// ERC20 constants
	PARAM_SUPPLY  = "s"
	PARAM_CREATOR = "c"
)

var (
	//WasmFileTestcore = "sbtestsc/testcore_bg.wasm"
	//WasmFileErc20    = "sbtestsc/erc20_bg.wasm"
	SandboxSCName    = "test_sandbox"
)

// deploy the specified contract on the chain
func DeployGoContract(chain *solo.Chain, keyPair *ed25519.KeyPair, name string, contractName string, params ...interface{}) error {
	if common.WasmRunner == 1 {
		wasmproc.GoWasmVM = common.NewWasmGoVM(common.ScForGoVM)
		hprog, err := chain.UploadWasm(keyPair, []byte("go:"+contractName))
		if err != nil {
			return err
		}
		return chain.DeployContract(keyPair, name, hprog, filterKeys(params...)...)
	}

	wasmproc.GoWasmVM = common.NewWasmTimeJavaVM()
	wasmFile := contractName + "_bg.wasm"
	wasmFile = util.LocateFile(wasmFile, contractName+"/pkg")
	return chain.DeployWasmContract(keyPair, name, wasmFile, filterKeys(params...)...)
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

func setupChain(t *testing.T, keyPairOriginator *ed25519.KeyPair) (*solo.Solo, *solo.Chain) {
	core.PrintWellKnownHnames()
	wasmhost.HostTracing = DEBUG
	wasmhost.ExtendedHostTracing = DEBUG
	env := solo.New(t, DEBUG, true)
	chain := env.NewChain(keyPairOriginator, "ch1")
	return env, chain
}

func setupDeployer(t *testing.T, chain *solo.Chain) (*ed25519.KeyPair, ledgerstate.Address, *coretypes.AgentID) {
	user, userAddr := chain.Env.NewKeyPairWithFunds()
	chain.Env.AssertAddressIotas(userAddr, solo.Saldo)

	req := solo.NewCallParams(root.Interface.Name, root.FuncGrantDeploy,
		root.ParamDeployer, coretypes.NewAgentID(userAddr, 0),
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
	return user, userAddr, coretypes.NewAgentID(userAddr, 0)
}

func run2(t *testing.T, test func(*testing.T, bool), skipWasm ...bool) {
	t.Run(fmt.Sprintf("run CORE version of %s", t.Name()), func(t *testing.T) {
		test(t, false)
	})
	if len(skipWasm) == 0 || !skipWasm[0] {
		t.Run(fmt.Sprintf("run WASM version of %s", t.Name()), func(t *testing.T) {
			test(t, true)
		})
	} else {
		t.Logf("skipped WASM version of '%s'", t.Name())
	}
}

func setupTestSandboxSC(t *testing.T, chain *solo.Chain, user *ed25519.KeyPair, runWasm bool) (*coretypes.AgentID, uint64) {
	var err error
	var extraToken uint64
	if runWasm {
		err = DeployGoContract(chain, user, SandboxSCName, "testcore")
		extraToken = 1
	} else {
		err = chain.DeployContract(user, SandboxSCName, sbtestsc.Interface.ProgramHash)
		extraToken = 0
	}
	require.NoError(t, err)

	deployed := coretypes.NewAgentID(chain.ChainID.AsAddress(), coretypes.Hn(sbtestsc.Interface.Name))
	req := solo.NewCallParams(SandboxSCName, sbtestsc.FuncDoNothing).WithIotas(1)
	_, err = chain.PostRequestSync(req, user)
	require.NoError(t, err)
	t.Logf("deployed test_sandbox'%s': %s", SandboxSCName, coretypes.Hn(SandboxSCName))
	return deployed, extraToken
}

func setupERC20(t *testing.T, chain *solo.Chain, user *ed25519.KeyPair, runWasm bool) *coretypes.AgentID {
	var err error
	if !runWasm {
		t.Logf("skipped %s. Only for Wasm tests, always loads %s", t.Name(), ERC20_NAME)
		return nil
	}
	userAddr := ledgerstate.NewED25519Address(user.PublicKey)
	var userAgentID *coretypes.AgentID
	if user == nil {
		userAgentID = &chain.OriginatorAgentID
	} else {
		userAgentID = coretypes.NewAgentID(userAddr, 0)
	}
	err = DeployGoContract(chain, user, ERC20_NAME, ERC20_NAME,
		PARAM_SUPPLY, 1000000,
		PARAM_CREATOR, userAgentID,
	)
	require.NoError(t, err)

	deployed := coretypes.NewAgentID(chain.ChainID.AsAddress(), coretypes.Hn(sbtestsc.Interface.Name))
	t.Logf("deployed erc20'%s': %s --  %s", ERC20_NAME, coretypes.Hn(ERC20_NAME), deployed)
	return deployed
}

func TestSetup1(t *testing.T) { run2(t, testSetup1) }
func testSetup1(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)
}

func TestSetup2(t *testing.T) { run2(t, testSetup2) }
func testSetup2(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	user, _, _ := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user, w)
}

func TestSetup3(t *testing.T) { run2(t, testSetup3) }
func testSetup3(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	user, _, _ := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user, w)
	setupERC20(t, chain, user, w)
}
