package test

import (
	"fmt"
	"testing"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/corecontracts/coreroot"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

//nolint:revive
const (
	DEBUG        = false
	ERC20_NAME   = "erc20"
	ERC20_SUPPLY = 100000

	// ERC20 constants
	PARAM_SUPPLY  = "s"
	PARAM_CREATOR = "c"
)

var (
// WasmFileTestcore = "sbtestsc/testcore_bg.wasm"
// WasmFileErc20    = "sbtestsc/erc20_bg.wasm"
)

// deploy the specified contract on the chain
func DeployGoContract(chain *solo.Chain, keyPair *ed25519.KeyPair, name, scName string, params ...interface{}) error {
	if *wasmsolo.GoDebug {
		wasmproc.GoWasmVM = wasmhost.NewWasmGoVM(scName, testcore.OnLoad)
		hprog, err := chain.UploadWasm(keyPair, []byte("go:"+scName))
		if err != nil {
			return err
		}
		return chain.DeployContract(keyPair, name, hprog, filterKeys(params...)...)
	}

	wasmproc.GoWasmVM = wasmhost.NewWasmTimeVM()
	wasmFile := scName + "_bg.wasm"
	wasmFile = util.LocateFile(wasmFile, scName+"/pkg")
	return chain.DeployWasmContract(keyPair, name, wasmFile, filterKeys(params...)...)
}

// filters wasmlib.Key parameters and replaces them with their proper string equivalent
func filterKeys(params ...interface{}) []interface{} {
	for i, param := range params {
		if par, ok := param.(wasmlib.Key); ok {
			params[i] = string(par)
		}
	}
	return params
}

func setupTest(t *testing.T, runWasm bool, addCreator ...bool) *wasmsolo.SoloContext {
	chain := wasmsolo.StartChain(t, "chain1")

	var creator *wasmsolo.SoloAgent
	var pair *ed25519.KeyPair
	if len(addCreator) != 0 && addCreator[0] {
		creator = wasmsolo.NewSoloAgent(chain.Env)
		pair = creator.Pair

		ctxRoot := wasmsolo.NewSoloContextForCore(t, chain, coreroot.ScName, coreroot.OnLoad)
		grant := coreroot.ScFuncs.GrantDeployPermission(ctxRoot)
		grant.Params.Deployer().SetValue(creator.ScAgentID())
		grant.Func.TransferIotas(1).Post()
		require.NoError(t, ctxRoot.Err)
	}

	if !runWasm {
		chain.Env.WithNativeContract(sbtestsc.Processor)
		err := chain.DeployContract(pair, testcore.ScName, sbtestsc.Contract.ProgramHash)
		require.NoError(t, err)
	}

	wasmsolo.SoloHost = nil
	ctx := wasmsolo.NewSoloContextForChain(t, chain, creator, testcore.ScName, testcore.OnLoad)
	require.NoError(t, ctx.Err)
	return ctx
}

func setupChain(t *testing.T, keyPairOriginator *ed25519.KeyPair) (*solo.Solo, *solo.Chain) {
	wasmhost.HostTracing = DEBUG
	wasmhost.ExtendedHostTracing = DEBUG
	core.PrintWellKnownHnames()
	env := solo.New(t, DEBUG, false).WithNativeContract(sbtestsc.Processor)
	chain := env.NewChain(keyPairOriginator, "ch1")
	return env, chain
}

func setupDeployer(t *testing.T, chain *solo.Chain) (*ed25519.KeyPair, ledgerstate.Address, *iscp.AgentID) {
	user, userAddr := chain.Env.NewKeyPairWithFunds()
	chain.Env.AssertAddressIotas(userAddr, solo.Saldo)

	req := solo.NewCallParams(root.Contract.Name, root.FuncGrantDeployPermission.Name,
		root.ParamDeployer, iscp.NewAgentID(userAddr, 0),
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
	return user, userAddr, iscp.NewAgentID(userAddr, 0)
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

func setupTestSandboxSC(t *testing.T, chain *solo.Chain, user *ed25519.KeyPair, runWasm bool) (*iscp.AgentID, uint64) {
	var err error
	var extraToken uint64
	if runWasm {
		err = DeployGoContract(chain, user, ScName, ScName)
		extraToken = 1
	} else {
		err = chain.DeployContract(user, ScName, sbtestsc.Contract.ProgramHash)
		extraToken = 0
	}
	require.NoError(t, err)

	deployed := iscp.NewAgentID(chain.ChainID.AsAddress(), HScName)
	req := solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(1)
	_, err = chain.PostRequestSync(req, user)
	require.NoError(t, err)
	t.Logf("deployed test_sandbox'%s': %s", ScName, HScName)
	return deployed, extraToken
}

//nolint:deadcode,unused
func setupERC20(t *testing.T, chain *solo.Chain, user *ed25519.KeyPair, runWasm bool) *iscp.AgentID {
	var err error
	if !runWasm {
		t.Logf("skipped %s. Only for Wasm tests, always loads %s", t.Name(), ERC20_NAME)
		return nil
	}
	var userAgentID *iscp.AgentID
	if user == nil {
		userAgentID = &chain.OriginatorAgentID
	} else {
		userAddr := ledgerstate.NewED25519Address(user.PublicKey)
		userAgentID = iscp.NewAgentID(userAddr, 0)
	}
	err = DeployGoContract(chain, user, ERC20_NAME, ERC20_NAME,
		PARAM_SUPPLY, 1000000,
		PARAM_CREATOR, userAgentID,
	)
	require.NoError(t, err)

	deployed := iscp.NewAgentID(chain.ChainID.AsAddress(), HScName)
	t.Logf("deployed erc20'%s': %s --  %s", ERC20_NAME, iscp.Hn(ERC20_NAME), deployed)
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
	// setupERC20(t, chain, user, w)
}
