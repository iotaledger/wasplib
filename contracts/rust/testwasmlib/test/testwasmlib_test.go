package test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/testwasmlib"
	"github.com/stretchr/testify/require"
)

var (
	allParams = []string{
		ParamAddress,
		ParamAgentID,
		ParamChainID,
		ParamColor,
		ParamHash,
		ParamHname,
		ParamInt16,
		ParamInt32,
		ParamInt64,
		ParamRequestID,
	}
	allLengths    = []int{33, 37, 33, 32, 32, 4, 2, 4, 8, 34}
	invalidValues = map[string][][]byte{
		ParamAddress: {
			append([]byte{3}, zeroHash...),
			append([]byte{4}, zeroHash...),
			append([]byte{255}, zeroHash...),
		},
		ParamChainID: {
			append([]byte{0}, zeroHash...),
			append([]byte{1}, zeroHash...),
			append([]byte{3}, zeroHash...),
			append([]byte{4}, zeroHash...),
			append([]byte{255}, zeroHash...),
		},
		ParamRequestID: {
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
	return common.StartChainAndDeployWasmContractByName(t, testwasmlib.ScName)
}

func TestDeploy(t *testing.T) {
	chain := setupTest(t)
	_, err := chain.FindContract(testwasmlib.ScName)
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

	chainID := chain.ChainID
	address := chainID.AsAddress()
	hname := HScName
	agentID := iscp.NewAgentID(address, hname)
	color, err := colored.ColorFromBytes([]byte("RedGreenBlueYellowCyanBlackWhite"))
	require.NoError(t, err)
	hash, err := hashing.HashValueFromBytes([]byte("0123456789abcdeffedcba9876543210"))
	require.NoError(t, err)
	requestID, err := iscp.RequestIDFromBytes([]byte("abcdefghijklmnopqrstuvwxyz123456\x00\x00"))
	require.NoError(t, err)
	req := solo.NewCallParams(ScName, FuncParamTypes,
		ParamAddress, address,
		ParamAgentID, agentID,
		ParamBytes, []byte("these are bytes"),
		ParamChainID, chainID,
		ParamColor, color,
		ParamHash, hash,
		ParamHname, hname,
		ParamInt16, int16(12345),
		ParamInt32, int32(1234567890),
		ParamInt64, int64(1234567890123456789),
		ParamRequestID, requestID,
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
			if param == ParamChainID {
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

func TestClearArray(t *testing.T) {
	chain := testValidParams(t)

	req := solo.NewCallParams(ScName, FuncArraySet,
		ParamName, "bands",
		ParamIndex, int32(0),
		ParamValue, "Simple Minds",
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	req = solo.NewCallParams(ScName, FuncArraySet,
		ParamName, "bands",
		ParamIndex, int32(1),
		ParamValue, "Dire Straits",
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	req = solo.NewCallParams(ScName, FuncArraySet,
		ParamName, "bands",
		ParamIndex, int32(2),
		ParamValue, "ELO",
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	res, err := chain.CallView(ScName, ViewArrayLength,
		ParamName, "bands")
	require.NoError(t, err)
	length, exist, err := codec.DecodeInt32(res.MustGet(ResultLength))
	require.NoError(t, err)
	require.True(t, exist)
	require.EqualValues(t, 3, length)

	res, err = chain.CallView(ScName, ViewArrayValue,
		ParamName, "bands",
		ParamIndex, int32(1))
	require.NoError(t, err)
	value, exist, err := codec.DecodeString(res.MustGet(ResultValue))
	require.NoError(t, err)
	require.True(t, exist)
	require.EqualValues(t, "Dire Straits", value)

	req = solo.NewCallParams(ScName, FuncArrayClear,
		ParamName, "bands",
	).WithIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	res, err = chain.CallView(ScName, ViewArrayLength,
		ParamName, "bands")
	require.NoError(t, err)
	length, exist, err = codec.DecodeInt32(res.MustGet(ResultLength))
	require.NoError(t, err)
	require.True(t, exist)
	require.EqualValues(t, 0, length)

	_, err = chain.CallView(ScName, ViewArrayValue, ParamName, "bands", ParamIndex, int32(0))
	require.Error(t, err)
}
