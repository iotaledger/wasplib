//nolint:dupl
package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/governance"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestOffLedgerFailNoAccount(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, _ := setupTestSandboxSC(t, chain, nil, w)

		owner, ownerAddr := env.NewKeyPairWithFunds()
		ownerAgentID := iscp.NewAgentID(ownerAddr, 0)

		chain.AssertIotas(ownerAgentID, 0)
		chain.AssertIotas(cAID, 1)

		// NOTE: NO deposit into owner account
		//req := solo.NewCallParams(accounts.Interface.Name, accounts.FuncDeposit,
		//).WithIotas(10)
		//_, err := chain.PostRequestSync(req, owner)
		//require.NoError(t, err)

		chain.AssertIotas(ownerAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req := solo.NewCallParams(ScName, sbtestsc.FuncSetInt.Name,
			sbtestsc.ParamIntParamName, "ppp",
			sbtestsc.ParamIntParamValue, 314,
		)
		_, err := chain.PostRequestOffLedger(req, owner)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "unverified account"))

		chain.AssertIotas(ownerAgentID, 0)
		chain.AssertIotas(cAID, 1)
	})
}

func TestOffLedgerNoFeeNoTransfer(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, _ := setupTestSandboxSC(t, chain, nil, w)

		owner, ownerAddr := env.NewKeyPairWithFunds()
		ownerAgentID := iscp.NewAgentID(ownerAddr, 0)

		chain.AssertIotas(ownerAgentID, 0)
		chain.AssertIotas(cAID, 1)

		// deposit into owner account
		req := solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(10)
		_, err := chain.PostRequestSync(req, owner)
		require.NoError(t, err)

		chain.AssertIotas(ownerAgentID, 10)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncSetInt.Name,
			sbtestsc.ParamIntParamName, "ppp",
			sbtestsc.ParamIntParamValue, 314,
		)
		// Look, Ma! No .WithIotas() necessary when doing off-ledger request!
		_, err = chain.PostRequestOffLedger(req, owner)
		require.NoError(t, err)

		chain.AssertIotas(ownerAgentID, 10)
		chain.AssertIotas(cAID, 1)

		ret, err := chain.CallView(ScName, sbtestsc.FuncGetInt.Name,
			sbtestsc.ParamIntParamName, "ppp")
		require.NoError(t, err)

		retInt, exists, err := codec.DecodeInt64(ret.MustGet("ppp"))
		require.NoError(t, err)
		require.True(t, exists)
		require.EqualValues(t, 314, retInt)
	})
}

func TestOffLedgerFeesEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(10)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 10)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)
		env.AssertAddressIotas(userAddr, solo.Saldo-10)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 10)
	})
}

func TestOffLedgerFeesNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(9)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 9)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10)
		_, err = chain.PostRequestOffLedger(req, user)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "not enough fees"))

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)
		env.AssertAddressIotas(userAddr, solo.Saldo-9)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 9)
	})
}

func TestOffLedgerFeesExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(11)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 11)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 1)
		chain.AssertIotas(cAID, 1)
		env.AssertAddressIotas(userAddr, solo.Saldo-11)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 10)
	})
}

func TestOffLedgerTransferWithFeesEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(10 + 42)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 10+42)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10 + 42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1+42)
		env.AssertAddressIotas(userAddr, solo.Saldo-10-42)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 10)
	})
}

func TestOffLedgerTransferWithFeesNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(10 + 41)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 10+41)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10 + 42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1+41)
		env.AssertAddressIotas(userAddr, solo.Saldo-10-41)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 10)
	})
}

func TestOffLedgerTransferWithFeesExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
			root.ParamHname, HScName,
			governance.ParamOwnerFee, 10,
		).WithIotas(1)
		_, err := chain.PostRequestSync(req, nil)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(10 + 43)
		_, err = chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 10+43)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(10 + 42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 1)
		chain.AssertIotas(cAID, 1+42)
		env.AssertAddressIotas(userAddr, solo.Saldo-10-43)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
		chain.AssertCommonAccountIotas(4 + extraToken + 10)
	})
}

func TestOffLedgerTransferEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req := solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(42)
		_, err := chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 42)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1+42)
		env.AssertAddressIotas(userAddr, solo.Saldo-42)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
		chain.AssertCommonAccountIotas(3 + extraToken)
	})
}

func TestOffLedgerTransferNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req := solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(41)
		_, err := chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 41)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1+41)
		env.AssertAddressIotas(userAddr, solo.Saldo-41)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
		chain.AssertCommonAccountIotas(3 + extraToken)
	})
}

func TestOffLedgerTransferExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		env, chain := setupChain(t, nil)
		cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
		user, userAddr, userAgentID := setupDeployer(t, chain)

		chain.AssertIotas(userAgentID, 0)
		chain.AssertIotas(cAID, 1)

		req := solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(43)
		_, err := chain.PostRequestSync(req, user)
		require.NoError(t, err)

		chain.AssertIotas(userAgentID, 43)
		chain.AssertIotas(cAID, 1)

		req = solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name).WithIotas(42)
		_, err = chain.PostRequestOffLedger(req, user)
		require.NoError(t, err)

		t.Logf("dump accounts:\n%s", chain.DumpAccounts())
		chain.AssertIotas(&chain.OriginatorAgentID, 0)
		chain.AssertIotas(userAgentID, 1)
		chain.AssertIotas(cAID, 1+42)
		env.AssertAddressIotas(userAddr, solo.Saldo-43)
		env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
		chain.AssertCommonAccountIotas(3 + extraToken)
	})
}
