package test

// import (
// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/iotaledger/wasp/packages/kv/codec"
// 	"github.com/iotaledger/wasp/packages/kv/collections"
// 	"github.com/iotaledger/wasp/packages/solo"
// 	"github.com/iotaledger/wasp/packages/vm/core/root"
// 	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
// 	"github.com/stretchr/testify/require"
// )

// // TODO refactor to use blocklog

// func TestEventlogGetLast3(t *testing.T) { run2(t, testEventlogGetLast3) }

// func testEventlogGetLast3(t *testing.T, w bool) {
// 	_, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	for i := 1; i < 6; i++ {
// 		req := solo.NewCallParams(ScName, sbtestsc.FuncEventLogGenericData.Name,
// 			sbtestsc.VarCounter, i,
// 		).WithIotas(1)
// 		_, err := chain.PostRequestSync(req, nil)
// 		require.NoError(t, err)
// 	}

// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamMaxLastRecords, 3,
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)

// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)
// 	require.EqualValues(t, 3, array.MustLen())
// }

// func TestEventlogGetBetweenTs(t *testing.T) { run2(t, testEventlogGetBetweenTs) }
// func testEventlogGetBetweenTs(t *testing.T, w bool) {
// 	env, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	env.SetTimeStep(500 * time.Millisecond)
// 	var err error
// 	for i := 1; i < 6; i++ {
// 		req := solo.NewCallParams(ScName,
// 			sbtestsc.FuncEventLogGenericData.Name,
// 			sbtestsc.VarCounter, i,
// 		).WithIotas(1)
// 		_, err = chain.PostRequestSync(req, nil)
// 		require.NoError(t, err)
// 	}

// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamFromTs, 0,
// 		eventlog.ParamToTs, chain.State.Timestamp().UnixNano()-int64(1500*time.Millisecond),
// 		eventlog.ParamMaxLastRecords, 2,
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)

// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)
// 	require.EqualValues(t, 2, array.MustLen())
// }

// func TestEventLogEventData(t *testing.T) { run2(t, testEventLogEventData) }
// func testEventLogEventData(t *testing.T, w bool) {
// 	_, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	req := solo.NewCallParams(ScName,
// 		sbtestsc.FuncEventLogEventData.Name,
// 	).WithIotas(1)
// 	_, err := chain.PostRequestSync(req, nil)
// 	require.NoError(t, err)

// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamFromTs, 0,
// 		eventlog.ParamToTs, chain.State.Timestamp(),
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)
// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)

// 	require.EqualValues(t, 1, array.MustLen())

// 	str, err := chain.GetEventsForContractString(ScName)
// 	require.NoError(t, err)
// 	t.Log(str)
// }

// func TestEventLogDifferentCalls(t *testing.T) { run2(t, testEventLogDifferentCalls) }
// func testEventLogDifferentCalls(t *testing.T, w bool) {
// 	env, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)
// 	env.SetTimeStep(500 * time.Millisecond)

// 	count := 1
// 	// events
// 	for i := 1; i <= 3; i++ {
// 		req := solo.NewCallParams(ScName,
// 			sbtestsc.FuncEventLogEventData.Name,
// 			sbtestsc.VarCounter, count,
// 		).WithIotas(1)
// 		count++
// 		_, err := chain.PostRequestSync(req, nil)
// 		require.NoError(t, err)
// 	}
// 	// generic
// 	for i := 1; i <= 2; i++ {
// 		req := solo.NewCallParams(ScName,
// 			sbtestsc.FuncEventLogGenericData.Name,
// 			sbtestsc.VarCounter, count,
// 		).WithIotas(1)
// 		count++
// 		_, err := chain.PostRequestSync(req, nil)
// 		require.NoError(t, err)
// 	}
// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamFromTs, 0,
// 		eventlog.ParamToTs, chain.State.Timestamp(),
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)
// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)
// 	require.EqualValues(t, 5, array.MustLen())

// 	str, err := chain.GetEventsForContractString(ScName)
// 	require.NoError(t, err)
// 	t.Log(str)

// 	require.EqualValues(t, 0, strings.Count(str, "[req]"))
// 	require.EqualValues(t, 2, strings.Count(str, "[GenericData]"))
// 	require.EqualValues(t, 3, strings.Count(str, "[Event]"))

// 	recStr := chain.GetRequestReceiptsForBlockRangeAsStrings(0, 0)
// 	str = strings.Join(recStr, "\n")
// 	t.Logf("\n%s", str)
// 	extra := 0
// 	if w {
// 		extra = 1
// 	}
// 	require.EqualValues(t, 8+extra, strings.Count(str, "req/tx"))
// }

// func TestChainLogGetNumRecords(t *testing.T) { run2(t, testChainLogGetNumRecords) }
// func testChainLogGetNumRecords(t *testing.T, w bool) {
// 	_, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	req := solo.NewCallParams(ScName,
// 		sbtestsc.FuncEventLogGenericData.Name,
// 		sbtestsc.VarCounter, solo.Saldo,
// 	).WithIotas(1)
// 	_, err := chain.PostRequestSync(req, nil)
// 	require.NoError(t, err)

// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetNumRecords.Name,
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)

// 	v, ok, err := codec.DecodeInt64(res.MustGet(eventlog.ParamNumRecords))

// 	require.NoError(t, err)
// 	require.True(t, ok)
// 	require.EqualValues(t, 1, v)

// 	str, err := chain.GetEventsForContractString(ScName)
// 	require.NoError(t, err)
// 	t.Log(str)

// 	require.EqualValues(t, 0, strings.Count(str, "[req]"))
// 	require.EqualValues(t, 1, strings.Count(str, "[GenericData]"))
// 	require.EqualValues(t, 0, strings.Count(str, "[Event]"))
// 	require.EqualValues(t, 1, strings.Count(str, strconv.FormatUint(solo.Saldo, 10)))

// 	recStr := chain.GetRequestReceiptsForBlockRangeAsStrings(0, 0)
// 	str = strings.Join(recStr, "\n")
// 	t.Logf("\n%s", str)
// 	extra := 0
// 	if w {
// 		extra = 1
// 	}
// 	require.EqualValues(t, 4+extra, strings.Count(str, "req/tx"))
// }

// func TestChainLogSandboxDeploy(t *testing.T) { run2(t, testChainLogSandboxDeploy) }
// func testChainLogSandboxDeploy(t *testing.T, w bool) {
// 	_, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	req := solo.NewCallParams(ScName,
// 		sbtestsc.FuncEventLogDeploy.Name,
// 	).WithIotas(1)
// 	_, err := chain.PostRequestSync(req, nil)
// 	require.NoError(t, err)

// 	// This call should return only one record which should be the type of TRDeploy
// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamContractHname, root.Contract.Hname(),
// 	)
// 	require.NoError(t, err)
// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)

// 	require.EqualValues(t, 2, array.MustLen())

// 	str, err := chain.GetEventsForContractString(root.Contract.Name)
// 	require.NoError(t, err)
// 	t.Log(str)

// 	require.EqualValues(t, 2, strings.Count(str, "[deploy]"))
// 	require.EqualValues(t, 0, strings.Count(str, "[req]"))

// 	recStr := chain.GetRequestReceiptsForBlockRangeAsStrings(0, 0)
// 	str = strings.Join(recStr, "\n")
// 	t.Logf("\n%s", str)
// 	extra := 0
// 	if w {
// 		extra = 1
// 	}
// 	require.EqualValues(t, 4+extra, strings.Count(str, "req/tx"))
// }

// func TestChainLogMultiple(t *testing.T) { run2(t, testChainLogMultiple) }
// func testChainLogMultiple(t *testing.T, w bool) {
// 	_, chain := setupChain(t, nil)
// 	setupTestSandboxSC(t, chain, nil, w)

// 	req := solo.NewCallParams(ScName,
// 		sbtestsc.FuncEventLogEventData.Name,
// 	).WithIotas(1)
// 	_, err := chain.PostRequestSync(req, nil)
// 	require.NoError(t, err)

// 	req = solo.NewCallParams(ScName,
// 		sbtestsc.FuncEventLogGenericData.Name,
// 		sbtestsc.VarCounter, 33333,
// 	).WithIotas(1)
// 	_, err = chain.PostRequestSync(req, nil)
// 	require.NoError(t, err)

// 	/////Should return 4 logs records/////
// 	res, err := chain.CallView(eventlog.Contract.Name, eventlog.FuncGetRecords.Name,
// 		eventlog.ParamContractHname, HScName,
// 	)
// 	require.NoError(t, err)
// 	array := collections.NewArray16ReadOnly(res, eventlog.ParamRecords)
// 	require.EqualValues(t, 2, array.MustLen())
// 	//////////////////////////////////////

// 	strRoot, err := chain.GetEventsForContractString(root.Contract.Name)
// 	require.NoError(t, err)
// 	t.Log(strRoot)
// 	require.EqualValues(t, 0, strings.Count(strRoot, "[req]"))
// 	require.EqualValues(t, 1, strings.Count(strRoot, "[deploy]"))

// 	strTest, err := chain.GetEventsForContractString(ScName)
// 	require.NoError(t, err)
// 	t.Log(strTest)
// 	require.EqualValues(t, 0, strings.Count(strTest, "[req]"))
// 	require.EqualValues(t, 1, strings.Count(strTest, "[GenericData]"))
// 	require.EqualValues(t, 1, strings.Count(strTest, "[Event]"))
// 	require.EqualValues(t, 1, strings.Count(strTest, "33333"))

// 	recStr := chain.GetRequestReceiptsForBlockRangeAsStrings(0, 0)
// 	str := strings.Join(recStr, "\n")
// 	t.Logf("\n%s", str)
// 	extra := 0
// 	if w {
// 		extra = 1
// 	}
// 	require.EqualValues(t, 5+extra, strings.Count(str, "req/tx"))
// }
