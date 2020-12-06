// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

import "strconv"

type ScBalances struct {
	balances ScImmutableMap
}

func (ctx ScBalances) Balance(color *ScColor) int64 {
	return ctx.balances.GetInt(color).Value()
}

func (ctx ScBalances) Minted() *ScColor {
	return NewScColor(ctx.balances.GetBytes(MINT).Value())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContract struct {
	contract ScImmutableMap
}

func (ctx ScContract) Color() *ScColor {
	return ctx.contract.GetColor(Key("color")).Value()
}

func (ctx ScContract) Description() string {
	return ctx.contract.GetString(Key("description")).Value()
}

func (ctx ScContract) Id() *ScAgent {
	return ctx.contract.GetAgent(Key("id")).Value()
}

func (ctx ScContract) Name() string {
	return ctx.contract.GetString(Key("name")).Value()
}

func (ctx ScContract) Owner() *ScAgent {
	return ctx.contract.GetAgent(Key("owner")).Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMapArray
}

func (ctx ScLog) Append(timestamp int64, data []byte) {
	logEntry := ctx.log.GetMap(ctx.log.Length())
	logEntry.GetInt(Key("timestamp")).SetValue(timestamp)
	logEntry.GetBytes(Key("data")).SetValue(data)
}

func (ctx ScLog) Length() int32 {
	return int32(ctx.log.Length())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Base58Decode(value string) []byte {
	decode := ctx.utility.GetString(Key("base58"))
	encode := ctx.utility.GetBytes(Key("base58"))
	decode.SetValue(value)
	return encode.Value()
}

func (ctx ScUtility) Base58Encode(value []byte) string {
	decode := ctx.utility.GetString(Key("base58"))
	encode := ctx.utility.GetBytes(Key("base58"))
	encode.SetValue(value)
	return decode.Value()
}

func (ctx ScUtility) Hash(value []byte) []byte {
	hash := ctx.utility.GetBytes(Key("hash"))
	hash.SetValue(value)
	return hash.Value()
}

func (ctx ScUtility) Random(max int64) int64 {
	rnd := ctx.utility.GetInt(Key("random")).Value()
	return int64(uint64(rnd) % uint64(max))
}

func (ctx ScUtility) String(value int64) string {
	return strconv.FormatInt(value, 10)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallContext struct {
	root ScMutableMap
}

func (ctx ScCallContext) Balances() ScBalances {
	return ScBalances{root.GetMap(Key("balances")).Immutable()}
}

func (ctx ScCallContext) Call(function string) ScCallInfo {
	return ScCallInfo{makeRequest("calls", function)}
}

func (ctx ScCallContext) Caller() *ScAgent {
	return root.GetAgent(Key("caller")).Value()
}

func (ctx ScCallContext) Contract() ScContract {
	return ScContract{root.GetMap(Key("contract")).Immutable()}
}

func (ctx ScCallContext) Error() ScMutableString {
	return root.GetString(Key("error"))
}

func (ctx ScCallContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

func (ctx ScCallContext) Incoming() ScBalances {
	return ScBalances{root.GetMap(Key("incoming")).Immutable()}
}

func (ctx ScCallContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScCallContext) Params() ScImmutableMap {
	return root.GetMap(Key("params")).Immutable()
}

func (ctx ScCallContext) Post(function string) ScPostInfo {
	return ScPostInfo{makeRequest("posts", function)}
}

func (ctx ScCallContext) Results() ScMutableMap {
	return root.GetMap(Key("results"))
}

func (ctx ScCallContext) State() ScMutableMap {
	return root.GetMap(Key("state"))
}

func (ctx ScCallContext) Timestamp() int64 {
	return root.GetInt(Key("timestamp")).Value()
}

func (ctx ScCallContext) TimestampedLog(key Key) ScLog {
	return ScLog{root.GetMap(Key("logs")).GetMapArray(key)}
}

func (ctx ScCallContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScCallContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	transfers := root.GetMapArray(Key("transfers"))
	transfer := transfers.GetMap(transfers.Length())
	transfer.GetAgent(Key("agent")).SetValue(agent)
	transfer.GetColor(Key("color")).SetValue(color)
	transfer.GetInt(Key("amount")).SetValue(amount)
}

func (ctx ScCallContext) Utility() ScUtility {
	return ScUtility{root.GetMap(Key("utility"))}
}

func (ctx ScCallContext) View(function string) ScViewInfo {
	return ScViewInfo{makeRequest("views", function)}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewContext struct {
	root ScMutableMap
}

func (ctx ScViewContext) Balances() ScBalances {
	return ScBalances{root.GetMap(Key("balances")).Immutable()}
}

func (ctx ScViewContext) Caller() *ScAgent {
	return root.GetAgent(Key("caller")).Value()
}

func (ctx ScViewContext) Contract() ScContract {
	return ScContract{root.GetMap(Key("contract")).Immutable()}
}

func (ctx ScViewContext) Error() ScMutableString {
	return root.GetString(Key("error"))
}

func (ctx ScViewContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

func (ctx ScViewContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScViewContext) Params() ScImmutableMap {
	return root.GetMap(Key("params")).Immutable()
}

func (ctx ScViewContext) Results() ScMutableMap {
	return root.GetMap(Key("results"))
}

func (ctx ScViewContext) State() ScImmutableMap {
	return root.GetMap(Key("state")).Immutable()
}

func (ctx ScViewContext) Timestamp() int64 {
	return root.GetInt(Key("timestamp")).Value()
}

func (ctx ScViewContext) TimestampedLog(key Key) ScImmutableMapArray {
	return root.GetMap(Key("logs")).GetMapArray(key).Immutable()
}

func (ctx ScViewContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScViewContext) Utility() ScUtility {
	return ScUtility{root.GetMap(Key("utility"))}
}

func (ctx ScViewContext) View(function string) ScViewInfo {
	return ScViewInfo{makeRequest("views", function)}
}
