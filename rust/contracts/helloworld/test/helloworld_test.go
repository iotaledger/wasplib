package test

import (
	"github.com/iotaledger/wasplib/govm"
	"github.com/iotaledger/wasplib/rust/contracts/helloworld/test/helloworld"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeployHelloWorld(t *testing.T) {
	te := govm.NewTestEnv(t, helloworld.ScName)
	_, err := te.Chain.FindContract(helloworld.ScName)
	require.NoError(t, err)
}

func TestHelloWorld(t *testing.T) {
	te := govm.NewTestEnv(t, helloworld.ScName)
	_ = te.NewCallParams(helloworld.FuncHelloWorld).Post(0)
}

func TestGetHelloWorld(t *testing.T) {
	te := govm.NewTestEnv(t, helloworld.ScName)
	res := te.CallView(helloworld.ViewGetHelloWorld)
	result := te.Results(res)
	hw := result.GetString(helloworld.VarHelloWorld)
	require.EqualValues(t, "Hello, world!", hw.Value())
}
