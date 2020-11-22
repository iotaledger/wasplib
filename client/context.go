// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

var (
	rootCallContext = ScCallContext{root: ScMutableMap{objId: 1}}
	rootViewContext = ScCallContext{root: ScMutableMap{objId: 1}}
)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScAccount struct {
	account ScImmutableMap
}

func (ctx ScAccount) Balance(color *ScColor) int64 {
	return ctx.account.GetKeyMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScAccount) Colors() ScImmutableColorArray {
	return ctx.account.GetColorArray("colors")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	call ScMutableMap
}

func (ctx ScCallInfo) contract(contract string) {
	ctx.call.GetString("contract").SetValue(contract)
}

func (ctx ScCallInfo) delay(delay int64) {
	ctx.call.GetInt("delay").SetValue(delay)
}

func (ctx ScCallInfo) function(function string) {
	ctx.call.GetString("function").SetValue(function)
}

func (ctx ScCallInfo) Params() ScMutableMap {
	return ctx.call.GetMap("params")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContract struct {
	contract ScImmutableMap
}

func (ctx ScContract) Color() *ScColor {
	return ctx.contract.GetColor("color").Value()
}

func (ctx ScContract) Description() string {
	return ctx.contract.GetString("description").Value()
}

func (ctx ScContract) Id() *ScAgent {
	return ctx.contract.GetAgent("id").Value()
}

func (ctx ScContract) Name() string {
	return ctx.contract.GetString("name").Value()
}

func (ctx ScContract) Owner() *ScAgent {
	return ctx.contract.GetAgent("owner").Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMapArray
}

func (ctx ScLog) Append(timestamp int64, data []byte) {
	logEntry := ctx.log.GetMap(ctx.log.Length())
	logEntry.GetInt("timestamp").SetValue(timestamp)
	logEntry.GetBytes("data").SetValue(data)
}

func (ctx ScLog) Length() int32 {
	return int32(ctx.log.Length())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequest struct {
	request ScImmutableMap
}

func (ctx ScRequest) Balance(color *ScColor) int64 {
	return ctx.request.GetKeyMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScRequest) Colors() ScImmutableColorArray {
	return ctx.request.GetColorArray("colors")
}

func (ctx ScRequest) From(originator *ScAgent) bool {
	return ctx.Sender().Equals(originator)
}

func (ctx ScRequest) Id() *ScRequestId {
	return ctx.request.GetRequestId("id").Value()
}

func (ctx ScRequest) MintedColor() *ScColor {
	return ctx.request.GetColor("hash").Value()
}

func (ctx ScRequest) Params() ScImmutableMap {
	return ctx.request.GetMap("params")
}

func (ctx ScRequest) Sender() *ScAgent {
	return ctx.request.GetAgent("sender").Value()
}

func (ctx ScRequest) Timestamp() int64 {
	return ctx.request.GetInt("timestamp").Value()
}

func (ctx ScRequest) TxHash() *ScTxHash {
	return ctx.request.GetTxHash("hash").Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfer struct {
	transfer ScMutableMap
}

func (ctx ScTransfer) Agent(agent *ScAgent) {
	ctx.transfer.GetAgent("agent").SetValue(agent)
}

func (ctx ScTransfer) Amount(amount int64) {
	ctx.transfer.GetInt("amount").SetValue(amount)
}

func (ctx ScTransfer) Color(color *ScColor) {
	ctx.transfer.GetColor("color").SetValue(color)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Base58Decode(value string) []byte {
	decode := ctx.utility.GetString("base58")
	encode := ctx.utility.GetBytes("base58")
	decode.SetValue(value)
	return encode.Value()
}

func (ctx ScUtility) Base58Encode(value []byte) string {
	decode := ctx.utility.GetString("base58")
	encode := ctx.utility.GetBytes("base58")
	encode.SetValue(value)
	return decode.Value()
}

func (ctx ScUtility) Hash(value []byte) []byte {
	hash := ctx.utility.GetBytes("hash")
	hash.SetValue(value)
	return hash.Value()
}

func (ctx ScUtility) Random(max int64) int64 {
	rnd := ctx.utility.GetInt("random").Value()
	return int64(uint64(rnd) % uint64(max))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallContext struct {
	root ScMutableMap
}

func (ctx ScCallContext) Account() ScAccount {
	return ScAccount{ctx.root.GetMap("account").Immutable()}
}

func (ctx ScCallContext) Call(contract string, function string) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	call := ScCallInfo{calls.GetMap(calls.Length())}
	call.contract(contract)
	call.function(function)
	return call
}

func (ctx ScCallContext) CallSelf(function string) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	call := ScCallInfo{calls.GetMap(calls.Length())}
	call.function(function)
	return call
}

func (ctx ScCallContext) Contract() ScContract {
	return ScContract{ctx.root.GetMap("contract").Immutable()}
}

func (ctx ScCallContext) Error() ScMutableString {
	return ctx.root.GetString("error")
}

func (ctx ScCallContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScCallContext) Post(contract string, function string, delay int64) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	request := ScCallInfo{calls.GetMap(calls.Length())}
	request.contract(contract)
	request.function(function)
	request.delay(delay)
	return request
}

func (ctx ScCallContext) PostSelf(function string, delay int64) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	request := ScCallInfo{calls.GetMap(calls.Length())}
	request.function(function)
	request.delay(delay)
	return request
}

func (ctx ScCallContext) Request() ScRequest {
	return ScRequest{ctx.root.GetMap("request").Immutable()}
}

func (ctx ScCallContext) State() ScMutableMap {
	return ctx.root.GetMap("state")
}

func (ctx ScCallContext) TimestampedLog(key string) ScLog {
	return ScLog{ctx.root.GetMap("logs").GetMapArray(key)}
}

func (ctx ScCallContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScCallContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	transfers := ctx.root.GetMapArray("transfers")
	xfer := ScTransfer{transfers.GetMap(transfers.Length())}
	xfer.Agent(agent)
	xfer.Color(color)
	xfer.Amount(amount)
}

func (ctx ScCallContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewContext struct {
	root ScMutableMap
}

func (ctx ScViewContext) Account() ScAccount {
	return ScAccount{ctx.root.GetMap("account").Immutable()}
}

func (ctx ScViewContext) Call(contract string, function string) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	call := ScCallInfo{calls.GetMap(calls.Length())}
	call.contract(contract)
	call.function(function)
	return call
}

func (ctx ScViewContext) CallSelf(function string) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	call := ScCallInfo{calls.GetMap(calls.Length())}
	call.function(function)
	return call
}

func (ctx ScViewContext) Contract() ScContract {
	return ScContract{ctx.root.GetMap("contract").Immutable()}
}

func (ctx ScViewContext) Error() ScMutableString {
	return ctx.root.GetString("error")
}

func (ctx ScViewContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScViewContext) Request() ScRequest {
	return ScRequest{ctx.root.GetMap("request").Immutable()}
}

func (ctx ScViewContext) State() ScImmutableMap {
	return ctx.root.GetMap("state").Immutable()
}

func (ctx ScViewContext) TimestampedLog(key string) ScLog {
	return ScLog{ctx.root.GetMap("logs").GetMapArray(key)}
}

func (ctx ScViewContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScViewContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}
