package test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
)

var (
	allParams = []string{
		ParamAddress,
		ParamAgentId,
		ParamChainId,
		ParamColor,
		ParamHash,
		ParamHname,
		ParamInt16,
		ParamInt32,
		ParamInt64,
		ParamRequestId,
	}
	allLengths    = []int{33, 37, 33, 32, 32, 4, 2, 4, 8, 34}
	invalidValues = map[string][][]byte{
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
	zeroHash = make([]byte, 32)
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestNoParams(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncParamTypes).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestValidParams(t *testing.T) {
	_ = testValidParams(t)
}

func testValidParams(t *testing.T) *solo.Chain {
	chain := setupTest(t)

	chainId := chain.ChainID
	address := chainId.AsAddress()
	hname := HScName
	agentId := coretypes.NewAgentID(address, hname)
	color, _, err := ledgerstate.ColorFromBytes([]byte("RedGreenBlueYellowCyanBlackWhite"))
	require.NoError(t, err)
	hash, err := hashing.HashValueFromBytes([]byte("0123456789abcdeffedcba9876543210"))
	require.NoError(t, err)
	requestId, err := coretypes.RequestIDFromBytes([]byte("abcdefghijklmnopqrstuvwxyz123456\x00\x00"))
	require.NoError(t, err)
	req := solo.NewCallParams(ScName, FuncParamTypes,
		ParamAddress, address,
		ParamAgentId, agentId,
		ParamBytes, []byte("these are bytes"),
		ParamChainId, chainId,
		ParamColor, color,
		ParamHash, hash,
		ParamHname, hname,
		ParamInt16, int16(12345),
		ParamInt32, int32(1234567890),
		ParamInt64, int64(1234567890123456789),
		ParamRequestId, requestId,
		ParamString, "this is a string",
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)
	return chain
}

func TestValidSizeParams(t *testing.T) {
	for index, param := range allParams {
		t.Run("ValidSize "+param, func(t *testing.T) {
			chain := setupTest(t)
			req := solo.NewCallParams(ScName, FuncParamTypes,
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

			req := solo.NewCallParams(ScName, FuncParamTypes,
				param, make([]byte, 0),
			).WithIotas(1)
			_, err := chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.HasSuffix(err.Error(), "invalid type size"))

			req = solo.NewCallParams(ScName, FuncParamTypes,
				param, make([]byte, allLengths[index]-1),
			).WithIotas(1)
			_, err = chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.Contains(err.Error(), "invalid type size"))

			req = solo.NewCallParams(ScName, FuncParamTypes,
				param, make([]byte, allLengths[index]+1),
			).WithIotas(1)
			_, err = chain.PostRequestSync(req, nil)
			require.Error(t, err)
			require.True(t, strings.Contains(err.Error(), "invalid type size"))
		})
	}
}

func TestInvalidTypeParams(t *testing.T) {
	for param, values := range invalidValues {
		for index, value := range values {
			t.Run("InvalidType "+param+" "+strconv.Itoa(index), func(t *testing.T) {
				chain := setupTest(t)
				req := solo.NewCallParams(ScName, FuncParamTypes,
					param, value,
				).WithIotas(1)
				_, err := chain.PostRequestSync(req, nil)
				require.Error(t, err)
				require.True(t, strings.Contains(err.Error(), "invalid "))
			})
		}
	}
}

func TestViewBlockRecords(t *testing.T) {
	t.SkipNow()
	chain := testValidParams(t)

	res, err := chain.CallView(ScName, ViewBlockRecords, ParamBlockIndex, int32(1))
	require.NoError(t, err)
	count, exist, err := codec.DecodeInt32(res.MustGet(ResultCount))
	require.NoError(t, err)
	require.True(t, exist)
	require.EqualValues(t, 1, count)
}
