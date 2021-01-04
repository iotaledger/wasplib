package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExample1(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ex1")

	chainInfo, coreContracts := chain.GetInfo()   // calls view root::GetInfo
	require.EqualValues(t, 4, len(coreContracts)) // 4 core contracts deployed by default

	t.Logf("chainID: %s", chainInfo.ChainID)
	t.Logf("chain owner ID: %s", chainInfo.ChainOwnerID)
	for hname, rec := range coreContracts {
		t.Logf("    Core contract '%s': %s", rec.Name, coretypes.NewContractID(chain.ChainID, hname))
	}
}

func TestExample2(t *testing.T) {
	glb := solo.New(t, false, false)
	userWallet := glb.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := glb.GetUtxodbBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	glb.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)
}

func TestExample3(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ex3")

	userWallet := glb.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := glb.GetUtxodbBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	glb.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)

	// send 42 iotas to the own account on-chain
	req := solo.NewCall("accounts", "deposit").
		WithTransfer(map[balance.Color]int64{
			balance.ColorIOTA: 42,
		})
	_, err := chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 43) // 43!!

	// withdraw back all iotas
	req = solo.NewCall("accounts", "withdraw")
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0) // empty
	glb.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)
}
