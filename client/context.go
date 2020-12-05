// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

import "strconv"

var (
	rootCallContext = ScCallContext{root: ScMutableMap{objId: 1}}
	rootViewContext = ScViewContext{root: ScMutableMap{objId: 1}}
)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScBalances struct {
	balances ScImmutableKeyMap
}

func (ctx ScBalances) Balance(color *ScColor) int64 {
	return ctx.balances.GetInt(color.Bytes()).Value()
}

func (ctx ScBalances) Minted() *ScColor {
	mintKey := MINT.Bytes()
	return NewScColor(ctx.balances.GetBytes(mintKey).Value())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	call ScMutableMap
}

func (ctx ScCallInfo) Call() ScCallInfo {
	ctx.call.GetInt("delay").SetValue(-1)
	return ctx
}

func (ctx ScCallInfo) Params() ScMutableMap {
	return ctx.call.GetMap("params")
}

func (ctx ScCallInfo) Results() ScImmutableMap {
	return ctx.call.GetMap("results").Immutable()
}

func (ctx ScCallInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.call.GetKeyMap("transfers")
	transfers.GetInt(color.Bytes()).SetValue(amount)
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

type ScPostInfo struct {
	post ScMutableMap
}

func (ctx ScPostInfo) Params() ScMutableMap {
	return ctx.post.GetMap("params")
}

func (ctx ScPostInfo) Post(delay int64) {
	ctx.post.GetInt("delay").SetValue(delay)
}

func (ctx ScPostInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.post.GetKeyMap("transfers")
	transfers.GetInt(color.Bytes()).SetValue(amount)
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

func (ctx ScUtility) String(value int64) string {
	return strconv.FormatInt(value, 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewInfo struct {
	view ScMutableMap
}

func (ctx ScViewInfo) Params() ScMutableMap {
	return ctx.view.GetMap("params")
}

func (ctx ScViewInfo) Results() ScImmutableMap {
	return ctx.view.GetMap("results").Immutable()
}

func (ctx ScViewInfo) View() ScViewInfo {
	ctx.view.GetInt("delay").SetValue(-2)
	return ctx
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallContext struct {
	root ScMutableMap
}

func (ctx ScCallContext) Balances() ScBalances {
	return ScBalances{ctx.root.GetKeyMap("balances").Immutable()}
}

func (ctx ScCallContext) Call(contract string, function string) ScCallInfo {
	calls := ctx.root.GetMapArray("calls")
	call := calls.GetMap(calls.Length())
	call.GetString("contract").SetValue(contract)
	call.GetString("function").SetValue(function)
	return ScCallInfo{call}
}

func (ctx ScCallContext) Caller() *ScAgent {
	return ctx.root.GetAgent("caller").Value()
}

func (ctx ScCallContext) CallSelf(function string) ScCallInfo {
	return ctx.Call("", function)
}

func (ctx ScCallContext) Contract() ScContract {
	return ScContract{ctx.root.GetMap("contract").Immutable()}
}

func (ctx ScCallContext) Error() ScMutableString {
	return ctx.root.GetString("error")
}

func (ctx ScCallContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

func (ctx ScCallContext) Incoming() ScBalances {
	return ScBalances{ctx.root.GetKeyMap("incoming").Immutable()}
}

func (ctx ScCallContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScCallContext) Params() ScImmutableMap {
	return ctx.root.GetMap("params").Immutable()
}

func (ctx ScCallContext) PostGlobal(chain *ScAddress, contract string, function string) ScPostInfo {
	posts := ctx.root.GetMapArray("posts")
	post := posts.GetMap(posts.Length())
	post.GetAddress("chain").SetValue(chain)
	post.GetString("contract").SetValue(contract)
	post.GetString("function").SetValue(function)
	return ScPostInfo{post}
}

func (ctx ScCallContext) PostLocal(contract string, function string) ScPostInfo {
	posts := ctx.root.GetMapArray("posts")
	post := posts.GetMap(posts.Length())
	post.GetString("contract").SetValue(contract)
	post.GetString("function").SetValue(function)
	return ScPostInfo{post}
}

func (ctx ScCallContext) PostSelf(function string) ScPostInfo {
	return ctx.PostLocal("", function)
}

func (ctx ScCallContext) Results() ScMutableMap {
	return ctx.root.GetMap("results")
}

func (ctx ScCallContext) State() ScMutableMap {
	return ctx.root.GetMap("state")
}

func (ctx ScCallContext) Timestamp() int64 {
	return ctx.root.GetInt("timestamp").Value()
}

func (ctx ScCallContext) TimestampedLog(key string) ScLog {
	return ScLog{ctx.root.GetMap("logs").GetMapArray(key)}
}

func (ctx ScCallContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScCallContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	transfers := ctx.root.GetMapArray("transfers")
	transfer := transfers.GetMap(transfers.Length())
	transfer.GetAgent("agent").SetValue(agent)
	transfer.GetColor("color").SetValue(color)
	transfer.GetInt("amount").SetValue(amount)
}

func (ctx ScCallContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}

func (ctx ScCallContext) View(contract string, function string) ScViewInfo {
	views := ctx.root.GetMapArray("views")
	view := views.GetMap(views.Length())
	view.GetString("contract").SetValue(contract)
	view.GetString("function").SetValue(function)
	return ScViewInfo{view}
}

func (ctx ScCallContext) ViewSelf(function string) ScViewInfo {
	return ctx.View("", function)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewContext struct {
	root ScMutableMap
}

func (ctx ScViewContext) Balances() ScBalances {
	return ScBalances{ctx.root.GetKeyMap("balances").Immutable()}
}

func (ctx ScViewContext) Caller() *ScAgent {
	return ctx.root.GetAgent("caller").Value()
}

func (ctx ScViewContext) Contract() ScContract {
	return ScContract{ctx.root.GetMap("contract").Immutable()}
}

func (ctx ScViewContext) Error() ScMutableString {
	return ctx.root.GetString("error")
}

func (ctx ScViewContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

func (ctx ScViewContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScViewContext) Params() ScImmutableMap {
	return ctx.root.GetMap("params").Immutable()
}

func (ctx ScViewContext) Results() ScMutableMap {
	return ctx.root.GetMap("results")
}

func (ctx ScViewContext) State() ScImmutableMap {
	return ctx.root.GetMap("state").Immutable()
}

func (ctx ScViewContext) Timestamp() int64 {
	return ctx.root.GetInt("timestamp").Value()
}

func (ctx ScViewContext) TimestampedLog(key string) ScImmutableMapArray {
	return ctx.root.GetMap("logs").GetMapArray(key).Immutable()
}

func (ctx ScViewContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScViewContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}

func (ctx ScViewContext) View(contract string, function string) ScViewInfo {
	views := ctx.root.GetMapArray("views")
	view := views.GetMap(views.Length())
	view.GetString("contract").SetValue(contract)
	view.GetString("function").SetValue(function)
	return ScViewInfo{view}
}

func (ctx ScViewContext) ViewSelf(function string) ScViewInfo {
	return ctx.View("", function)
}
