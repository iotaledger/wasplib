package test

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/testwasmlib"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

var (
	ParamAddress   = string(testwasmlib.ParamAddress)
	ParamAgentId   = string(testwasmlib.ParamAgentId)
	ParamBytes     = string(testwasmlib.ParamBytes)
	ParamChainId   = string(testwasmlib.ParamChainId)
	ParamColor     = string(testwasmlib.ParamColor)
	ParamHash      = string(testwasmlib.ParamHash)
	ParamHname     = string(testwasmlib.ParamHname)
	ParamInt64     = string(testwasmlib.ParamInt64)
	ParamRequestId = string(testwasmlib.ParamRequestId)
	ParamString    = string(testwasmlib.ParamString)

	allParams = []string{
		ParamAddress,
		ParamAgentId,
		ParamChainId,
		ParamColor,
		ParamHash,
		ParamHname,
		ParamInt64,
		ParamRequestId,
	}
	allLengths = []int{33, 37, 33, 32, 32, 4, 8, 34}
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, testwasmlib.ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, testwasmlib.ScName)
	_, err := chain.FindContract(testwasmlib.ScName)
	require.NoError(t, err)
}

func TestNoParams(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestValidParams(t *testing.T) {
	chain := setupTest(t)

	chainId := chain.ChainID
	address := chainId.AsAddress()
	hname := coretypes.Hn(testwasmlib.ScName)
	agentId := coretypes.NewAgentID(address, hname)
	color, _, err := ledgerstate.ColorFromBytes([]byte("RedGreenBlueYellowCyanBlackWhite"))
	require.NoError(t, err)
	hash, err := hashing.HashValueFromBytes([]byte("0123456789abcdeffedcba9876543210"))
	require.NoError(t, err)
	//requestId,err := coretypes.RequestIDFromBytes([]byte("abcdefghijklmnopqrstuvwxyz123456\x00\x00"))
	//require.NoError(t, err)
	req := solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
		ParamAddress, address,
		ParamAgentId, agentId,
		ParamBytes, []byte("these are bytes"),
		ParamChainId, chainId,
		ParamColor, color,
		ParamHash, hash,
		ParamHname, hname,
		ParamInt64, int64(1234567890123456789),
		//ParamRequestId, requestId,
		ParamString, "this is a string",
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestValidSizeParams(t *testing.T) {
	for index, param := range allParams {
		t.Run("ValidSize "+param, func(t *testing.T) {
			chain := setupTest(t)
			req := solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
				param, make([]byte, allLengths[index]),
			).WithIotas(1)
			_, err := chain.PostRequestSync(req, nil)
			require.Error(t, err)
			if param == ParamChainId {
				require.True(t, strings.Contains(err.Error(), "invalid "))
			} else {
				require.True(t, strings.Contains(err.Error(), "mismatch: "))
			}
		})
	}
}

func TestInvalidSizeParams(t *testing.T) {
	for index, param := range allParams {
		t.Run("InvalidSize "+param, func(t *testing.T) {
			chain := setupTest(t)

			req := solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
				param, make([]byte, 0),
			).WithIotas(1)
			_, err := chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.HasSuffix(err.Error(), "invalid type size"))

			req = solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
				param, make([]byte, allLengths[index]-1),
			).WithIotas(1)
			_, err = chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.Contains(err.Error(), "invalid type size"))

			req = solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
				param, make([]byte, allLengths[index]+1),
			).WithIotas(1)
			_, err = chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.Contains(err.Error(), "invalid type size"))
		})
	}
}

var zeroHash = make([]byte, 32)
var invalidValues = map[string][][]byte{
	ParamAddress: {
		append([]byte{3}, zeroHash...),
		append([]byte{4}, zeroHash...),
		append([]byte{255}, zeroHash...),
	},
	ParamChainId: {
		append([]byte{0}, zeroHash...),
		append([]byte{1}, zeroHash...),
		append([]byte{3}, zeroHash...),
		append([]byte{4}, zeroHash...),
		append([]byte{255}, zeroHash...),
	},
	ParamRequestId: {
		append(zeroHash, []byte{128, 0}...),
		append(zeroHash, []byte{127, 1}...),
		append(zeroHash, []byte{0, 1}...),
		append(zeroHash, []byte{255, 255}...),
		append(zeroHash, []byte{4, 4}...),
	},
}

func TestInvalidTypeParams(t *testing.T) {
	for param, values := range invalidValues {
		for index, value := range values {
			t.Run("InvalidType "+param + " " + strconv.Itoa(index), func(t *testing.T) {
				chain := setupTest(t)
				req := solo.NewCallParams(testwasmlib.ScName, testwasmlib.FuncParamTypes,
					param, value,
				).WithIotas(1)
				_, err := chain.PostRequestSync(req, nil)
				require.Error(t, err)
				require.True(t, strings.Contains(err.Error(), "invalid "))
			})
		}
	}
}
