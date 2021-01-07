package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExample1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex1")

	chainInfo, coreContracts := chain.GetInfo() // calls view root::GetInfo
	t.Logf("chainID: %s", chainInfo.ChainID)
	t.Logf("chain owner ID: %s", chainInfo.ChainOwnerID)
	for hname, rec := range coreContracts {
		t.Logf("    Core contract '%s': %s", rec.Name, coretypes.NewContractID(chain.ChainID, hname))
	}
}

func TestExample2(t *testing.T) {
	env := solo.New(t, false, false)
	userWallet := env.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := env.GetUtxodbBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	env.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)
}

func TestExample3(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex3")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	req := solo.NewCall("hello_new_world_1", "hello")
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	// call the contract to extract the value of the 'counter'. Must be equal 1
	res, err := chain.CallView("hello_new_world_1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 1, counter)
}

func TestExample4(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex4")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	// call to panic on purpose
	req := solo.NewCall("hello_new_world_1", "hello",
		"panic", 1)
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err) // expect error

	// call the contract to extract the value of the 'counter'. Must be equal 0
	res, err := chain.CallView("hello_new_world_1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 0, counter)
}

func TestExample5(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex5")

	userWallet := env.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := env.GetUtxodbBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	env.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)

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
	env.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)
}

func TestExample6(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex6")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	userWallet := env.NewSignatureSchemeWithFunds()

	contractName := "hello_new_world_1"
	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(contractName))
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	req := solo.NewCall(contractName, "hello").WithTransfer(map[balance.Color]int64{
		balance.ColorIOTA: 7,
	},
	)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	// call the contract to extract the value of the 'counter'. Must be equal 1
	res, err := chain.CallView("hello_new_world_1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 1, counter)

	// check the balance of the smart contract. Expect 7 iotas
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 7)
	// check the balance of the user address. Must be 8 iotas less
	env.AssertUtxodbBalance(userWallet.Address(), balance.ColorIOTA, 1337-8)
}

func TestExample7(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex7")

	err := chain.DeployWasmContract(nil, "hello_new_world_1", WASM_FILE)
	require.NoError(t, err)

	userWallet := env.NewSignatureSchemeWithFunds()

	contractName := "hello_new_world_1"
	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(contractName))
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	req := solo.NewCall(contractName, "hello", "panic", 1).
		WithTransfer(map[balance.Color]int64{
			balance.ColorIOTA: 7,
		},
		)
	_, err = chain.PostRequest(req, userWallet)
	require.Error(t, err)

	// check the balance of the smart contract. Expect 0 iotas
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 0)
	// check the balance of the user address. Must be 1 iotas less (for request)
	env.AssertUtxodbBalance(userWallet.Address(), balance.ColorIOTA, 1337-1)
}
