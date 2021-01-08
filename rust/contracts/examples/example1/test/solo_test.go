package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolo1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex1")

	chainInfo, coreContracts := chain.GetInfo()   // calls view root::GetInfo
	require.EqualValues(t, 4, len(coreContracts)) // 4 core contracts deployed by default

	t.Logf("chainID: %s", chainInfo.ChainID)
	t.Logf("chain owner ID: %s", chainInfo.ChainOwnerID)
	for hname, rec := range coreContracts {
		t.Logf("    Core contract '%s': %s", rec.Name, coretypes.NewContractID(chain.ChainID, hname))
	}
}

func TestSolo2(t *testing.T) {
	env := solo.New(t, false, false)
	userWallet := env.NewSignatureSchemeWithFunds() // create new wallet with 1337 iotas
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := env.GetAddressBalance(userAddress, balance.ColorIOTA) // how many iotas contains the address
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337) // assert the address has 1337 iotas
}

func TestExample3(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ex3")

	userWallet := glb.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := glb.GetAddressBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	glb.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)

	// send 42 iotas to the own account on-chain
	req := solo.NewCall("accounts", "deposit").
		WithTransfer(balance.ColorIOTA, 42)
	_, err := chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 43) // 43!!

	// withdraw back all iotas
	req = solo.NewCall("accounts", "withdraw")
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0) // empty
	glb.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)
}
