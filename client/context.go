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

func (ctx ScContract) Chain() *ScAddress {
	return ctx.contract.GetAddress(KeyChain).Value()
}

func (ctx ScContract) Description() string {
	return ctx.contract.GetString(KeyDescription).Value()
}

func (ctx ScContract) Id() *ScAgent {
	return ctx.contract.GetAgent(KeyId).Value()
}

func (ctx ScContract) Name() string {
	return ctx.contract.GetString(KeyName).Value()
}

func (ctx ScContract) Owner() *ScAgent {
	return ctx.contract.GetAgent(KeyOwner).Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMapArray
}

func (ctx ScLog) Append(timestamp int64, data []byte) {
	logEntry := ctx.log.GetMap(ctx.log.Length())
	logEntry.GetInt(KeyTimestamp).SetValue(timestamp)
	logEntry.GetBytes(KeyData).SetValue(data)
}

func (ctx ScLog) Length() int32 {
	return ctx.log.Length()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Base58Decode(value string) []byte {
	decode := ctx.utility.GetString(KeyBase58)
	encode := ctx.utility.GetBytes(KeyBase58)
	decode.SetValue(value)
	return encode.Value()
}

func (ctx ScUtility) Base58Encode(value []byte) string {
	decode := ctx.utility.GetString(KeyBase58)
	encode := ctx.utility.GetBytes(KeyBase58)
	encode.SetValue(value)
	return decode.Value()
}

func (ctx ScUtility) Hash(value []byte) []byte {
	hash := ctx.utility.GetBytes(KeyHash)
	hash.SetValue(value)
	return hash.Value()
}

func (ctx ScUtility) Random(max int64) int64 {
	rnd := ctx.utility.GetInt(KeyRandom).Value()
	return int64(uint64(rnd) % uint64(max))
}

func (ctx ScUtility) String(value int64) string {
	return strconv.FormatInt(value, 10)
}

func base58Encode(bytes []byte) string {
	return ScCallContext{}.Utility().Base58Encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScBaseContext struct {
}

func (ctx ScBaseContext) Balances() ScBalances {
	return ScBalances{root.GetMap(KeyBalances).Immutable()}
}

func (ctx ScBaseContext) Caller() *ScAgent {
	return root.GetAgent(KeyCaller).Value()
}

func (ctx ScBaseContext) Contract() ScContract {
	return ScContract{root.GetMap(KeyContract).Immutable()}
}

func (ctx ScBaseContext) Error() ScMutableString {
	return root.GetString(KeyError)
}

func (ctx ScBaseContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

func (ctx ScBaseContext) Log(text string) {
	SetString(1, int32(KeyLog), text)
}

func (ctx ScBaseContext) Params() ScImmutableMap {
	return root.GetMap(KeyParams).Immutable()
}

func (ctx ScBaseContext) Results() ScMutableMap {
	return root.GetMap(KeyResults)
}

func (ctx ScBaseContext) Timestamp() int64 {
	return root.GetInt(KeyTimestamp).Value()
}

func (ctx ScBaseContext) Trace(text string) {
	SetString(1, int32(KeyTrace), text)
}

func (ctx ScBaseContext) Utility() ScUtility {
	return ScUtility{root.GetMap(KeyUtility)}
}

func (ctx ScBaseContext) View(function string) ScViewInfo {
	return ScViewInfo{NewScBaseInfo(KeyViews, function)}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallContext struct {
	ScBaseContext
}

func (ctx ScCallContext) Call(function string) ScCallInfo {
	return ScCallInfo{NewScBaseInfo(KeyCalls, function)}
}

func (ctx ScCallContext) Incoming() ScBalances {
	return ScBalances{root.GetMap(KeyIncoming).Immutable()}
}

func (ctx ScCallContext) Post(function string) ScPostInfo {
	return ScPostInfo{NewScBaseInfo(KeyPosts, function)}
}

func (ctx ScCallContext) State() ScMutableMap {
	return root.GetMap(KeyState)
}

func (ctx ScCallContext) TimestampedLog(key MapKey) ScLog {
	return ScLog{root.GetMap(KeyLogs).GetMapArray(key)}
}

func (ctx ScCallContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	transfers := root.GetMapArray(KeyTransfers)
	transfer := transfers.GetMap(transfers.Length())
	transfer.GetAgent(KeyAgent).SetValue(agent)
	transfer.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewContext struct {
	ScBaseContext
}

func (ctx ScViewContext) State() ScImmutableMap {
	return root.GetMap(KeyState).Immutable()
}

func (ctx ScViewContext) TimestampedLog(key MapKey) ScImmutableMapArray {
	return root.GetMap(KeyLogs).GetMapArray(key).Immutable()
}
