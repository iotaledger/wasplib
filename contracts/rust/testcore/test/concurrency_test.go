package test

import (
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/kvdecoder"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestCounter(t *testing.T) { run2(t, testCounter) }
func testCounter(t *testing.T, w bool) {
	ctx := setupTest(t, w)

	f := testcore.ScFuncs.IncCounter(ctx)
	f.Func.TransferIotas(1)
	for i := 0; i < 33; i++ {
		f.Func.Post()
		require.NoError(t, ctx.Err)
	}

	v := testcore.ScFuncs.GetCounter(ctx)
	v.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, 33, v.Results.Counter().Value())
}

func TestConcurrency(t *testing.T) { run2(t, testConcurrency) }
func testConcurrency(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	extra := 0
	if w {
		extra = 1
	}
	req := solo.NewCallParams(ScName, sbtestsc.FuncIncCounter.Name).WithIotas(1)

	repeats := []int{300, 100, 100, 100, 200, 100, 100}
	sum := 0
	for _, i := range repeats {
		sum += i
	}
	for r, n := range repeats {
		go func(_, n int) {
			for i := 0; i < n; i++ {
				tx, _, err := chain.RequestFromParamsToLedger(req, nil)
				require.NoError(t, err)
				chain.Env.EnqueueRequests(tx)
			}
		}(r, n)
	}
	require.True(t, chain.WaitForRequestsThrough(sum+3+extra, 20*time.Second))

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)

	deco := kvdecoder.New(ret, chain.Log)
	res := deco.MustGetInt64(sbtestsc.VarCounter)
	require.EqualValues(t, sum, res)

	extraIota := uint64(0)
	if w {
		extraIota = 1
	}
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertCommonAccountIotas(extraIota + 2)
	agentID := iscp.NewAgentID(chain.ChainID.AsAddress(), HScName)
	chain.AssertIotas(agentID, uint64(sum)+1)
}

func TestConcurrency2(t *testing.T) { run2(t, testConcurrency2) }
func testConcurrency2(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	extra := 0
	if w {
		extra = 1
	}
	req := solo.NewCallParams(ScName, sbtestsc.FuncIncCounter.Name).WithIotas(1)

	repeats := []int{300, 100, 100, 100, 200, 100, 100}
	users := make([]*ed25519.KeyPair, len(repeats))
	userAddr := make([]ledgerstate.Address, len(repeats))
	sum := 0
	for _, i := range repeats {
		sum += i
	}
	for r, n := range repeats {
		go func(r, n int) {
			users[r], userAddr[r] = chain.Env.NewKeyPairWithFunds()
			for i := 0; i < n; i++ {
				tx, _, err := chain.RequestFromParamsToLedger(req, users[r])
				require.NoError(t, err)
				chain.Env.EnqueueRequests(tx)
			}
		}(r, n)
	}

	require.True(t, chain.WaitForRequestsThrough(sum+3+extra, 20*time.Second))

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)

	deco := kvdecoder.New(ret, chain.Log)
	res := deco.MustGetInt64(sbtestsc.VarCounter)
	require.EqualValues(t, sum, res)

	for i := range users {
		chain.AssertIotas(iscp.NewAgentID(userAddr[i], 0), 0)
		chain.Env.AssertAddressIotas(userAddr[i], solo.Saldo-uint64(repeats[i]))
	}
	extraIota := uint64(0)
	if w {
		extraIota = 1
	}
	chain.AssertIotas(&chain.OriginatorAgentID, 0)
	chain.AssertCommonAccountIotas(extraIota + 2)
	agentID := iscp.NewAgentID(chain.ChainID.AsAddress(), HScName)
	chain.AssertIotas(agentID, uint64(sum)+1)
}
