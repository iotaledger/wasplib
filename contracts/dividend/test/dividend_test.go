package test

import (
	"github.com/iotaledger/wasplib/contracts/dividend"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
)

const scName = "dividend"

func TestDeploy(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	_, err := te.Chain.FindContract(scName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	user1 := te.Env.NewSignatureSchemeWithFunds()
	_ = te.NewCallParams("member",
		dividend.ParamAddress, user1.Address(),
		dividend.ParamFactor, 100,
	).Post(0)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	_ = te.NewCallParams("member",
		dividend.ParamFactor, 100,
	).PostFail(0)
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	te := govm.NewTestEnv(t, scName)
	user1 := te.Env.NewSignatureSchemeWithFunds()
	_ = te.NewCallParams("member",
		dividend.ParamAddress, user1.Address(),
	).PostFail(0)
}
