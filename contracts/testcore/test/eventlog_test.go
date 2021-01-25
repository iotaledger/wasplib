package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/collections"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/eventlog"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestEventlogGetLast3(t *testing.T) { run2(t, testEventlogGetLast3, true) }
func testEventlogGetLast3(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	for i := 1; i < 6; i++ {
		req := solo.NewCall(SandboxSCName, test_sandbox_sc.FuncEventLogGenericData,
			test_sandbox_sc.VarCounter, i,
		)
		_, err := chain.PostRequest(req, nil)
		require.NoError(t, err)
	}

	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamMaxLastRecords, 3,
		eventlog.ParamContractHname, coretypes.Hn(SandboxSCName),
	)
	require.NoError(t, err)

	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)
	require.EqualValues(t, 3, array.MustLen())
}

func TestEventlogGetBetweenTs(t *testing.T) { run2(t, testEventlogGetBetweenTs, true) }
func testEventlogGetBetweenTs(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	env.SetTimeStep(500 * time.Millisecond)
	var err error
	for i := 1; i < 6; i++ {
		req := solo.NewCall(SandboxSCName,
			test_sandbox_sc.FuncEventLogGenericData,
			test_sandbox_sc.VarCounter, i,
		)
		_, err = chain.PostRequest(req, nil)
		require.NoError(t, err)
	}

	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamFromTs, 0,
		eventlog.ParamToTs, chain.State.Timestamp()-int64(1500*time.Millisecond),
		eventlog.ParamMaxLastRecords, 2,
		eventlog.ParamContractHname, test_sandbox_sc.Interface.Hname(),
	)
	require.NoError(t, err)

	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)
	require.EqualValues(t, 2, array.MustLen())
}

func TestEventLogEventData(t *testing.T) { run2(t, testEventLogEventData, true) }
func testEventLogEventData(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCall(test_sandbox_sc.Interface.Name,
		test_sandbox_sc.FuncEventLogEventData,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamFromTs, 0,
		eventlog.ParamToTs, chain.State.Timestamp(),
		eventlog.ParamContractHname, test_sandbox_sc.Interface.Hname(),
	)
	require.NoError(t, err)
	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)

	require.EqualValues(t, 3, array.MustLen())

	str, err := chain.GetEventLogRecordsString(test_sandbox_sc.Interface.Name)
	require.NoError(t, err)
	t.Log(str)
}

func TestEventLogDifferentCalls(t *testing.T) { run2(t, testEventLogDifferentCalls, true) }
func testEventLogDifferentCalls(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)
	env.SetTimeStep(500 * time.Millisecond)

	count := 1
	// events
	for i := 1; i <= 3; i++ {
		req := solo.NewCall(test_sandbox_sc.Interface.Name,
			test_sandbox_sc.FuncEventLogEventData,
			test_sandbox_sc.VarCounter, count,
		)
		count++
		_, err := chain.PostRequest(req, nil)
		require.NoError(t, err)
	}
	// generic
	for i := 1; i <= 2; i++ {
		req := solo.NewCall(test_sandbox_sc.Interface.Name,
			test_sandbox_sc.FuncEventLogGenericData,
			test_sandbox_sc.VarCounter, count,
		)
		count++
		_, err := chain.PostRequest(req, nil)
		require.NoError(t, err)
	}
	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamFromTs, 0,
		eventlog.ParamToTs, chain.State.Timestamp(),
		eventlog.ParamContractHname, test_sandbox_sc.Interface.Hname(),
	)
	require.NoError(t, err)
	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)
	require.EqualValues(t, 11, array.MustLen())

	str, err := chain.GetEventLogRecordsString(test_sandbox_sc.Interface.Name)
	require.NoError(t, err)
	t.Log(str)

	require.EqualValues(t, 6, strings.Count(str, "[req]"))
	require.EqualValues(t, 2, strings.Count(str, "[GenericData]"))
	require.EqualValues(t, 3, strings.Count(str, "[Event]"))
}

func TestChainLogGetNumRecords(t *testing.T) { run2(t, testChainLogGetNumRecords, true) }
func testChainLogGetNumRecords(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCall(SandboxSCName,
		test_sandbox_sc.FuncEventLogGenericData,
		test_sandbox_sc.VarCounter, 1337,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetNumRecords,
		eventlog.ParamContractHname, test_sandbox_sc.Interface.Hname(),
	)
	require.NoError(t, err)

	v, ok, err := codec.DecodeInt64(res.MustGet(eventlog.ParamNumRecords))

	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, 3, v)

	str, err := chain.GetEventLogRecordsString(SandboxSCName)
	require.NoError(t, err)
	t.Log(str)

	require.EqualValues(t, 2, strings.Count(str, "[req]"))
	require.EqualValues(t, 1, strings.Count(str, "[GenericData]"))
	require.EqualValues(t, 0, strings.Count(str, "[Event]"))
	require.EqualValues(t, 1, strings.Count(str, "1337"))
}

func TestChainLogSandboxDeploy(t *testing.T) { run2(t, testChainLogSandboxDeploy, true) }
func testChainLogSandboxDeploy(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCall(test_sandbox_sc.Interface.Name,
		test_sandbox_sc.FuncEventLogDeploy,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	//This call should return only one record which should be the type of TRDeploy
	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamContractHname, root.Interface.Hname(),
	)
	require.NoError(t, err)
	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)

	require.EqualValues(t, 4, array.MustLen())

	str, err := chain.GetEventLogRecordsString(root.Interface.Name)
	require.NoError(t, err)
	t.Log(str)

	require.EqualValues(t, 2, strings.Count(str, "[deploy]"))
	require.EqualValues(t, 2, strings.Count(str, "[req]"))
}

func TestChainLogMultiple(t *testing.T) { run2(t, testChainLogMultiple, true) }
func testChainLogMultiple(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCall(test_sandbox_sc.Interface.Name,
		test_sandbox_sc.FuncEventLogEventData,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	req = solo.NewCall(test_sandbox_sc.Interface.Name,
		test_sandbox_sc.FuncEventLogGenericData,
		test_sandbox_sc.VarCounter, 33333,
	)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	/////Should return 4 logs records/////
	res, err := chain.CallView(eventlog.Interface.Name, eventlog.FuncGetLogRecords,
		eventlog.ParamContractHname, test_sandbox_sc.Interface.Hname(),
	)
	require.NoError(t, err)
	array := collections.NewArrayReadOnly(res, eventlog.ParamRecords)
	require.EqualValues(t, 5, array.MustLen())
	//////////////////////////////////////

	strRoot, err := chain.GetEventLogRecordsString(root.Interface.Name)
	require.NoError(t, err)
	t.Log(strRoot)
	require.EqualValues(t, 2, strings.Count(strRoot, "[req]"))
	require.EqualValues(t, 1, strings.Count(strRoot, "[deploy]"))

	strTest, err := chain.GetEventLogRecordsString(test_sandbox_sc.Interface.Name)
	require.NoError(t, err)
	t.Log(strTest)
	require.EqualValues(t, 3, strings.Count(strTest, "[req]"))
	require.EqualValues(t, 1, strings.Count(strTest, "[GenericData]"))
	require.EqualValues(t, 1, strings.Count(strTest, "[Event]"))
	require.EqualValues(t, 1, strings.Count(strTest, "33333"))
}
