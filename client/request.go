// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScBaseInfo struct {
	request ScMutableMap
}

func NewScBaseInfo(key MapKey, function string) ScBaseInfo {
	requests := root.GetMapArray(key)
	request := requests.GetMap(requests.Length())
	request.GetString(KeyFunction).SetValue(function)
	return ScBaseInfo{request}
}

func (ctx ScBaseInfo) Contract(contract string) ScBaseInfo {
	ctx.request.GetString(KeyContract).SetValue(contract)
	return ctx
}

func (ctx ScBaseInfo) exec(delay int64) {
	ctx.request.GetInt(KeyDelay).SetValue(delay)
}

func (ctx ScBaseInfo) Params() ScMutableMap {
	return ctx.request.GetMap(KeyParams)
}

func (ctx ScBaseInfo) results() ScImmutableMap {
	return ctx.request.GetMap(KeyResults).Immutable()
}

func (ctx ScBaseInfo) transfer(color *ScColor, amount int64) {
	transfers := ctx.request.GetMap(KeyTransfers)
	transfers.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	ScBaseInfo
}

func (ctx ScCallInfo) Call() ScCallInfo {
	ctx.exec(-1)
	return ctx
}

func (ctx ScCallInfo) Results() ScImmutableMap {
	return ctx.ScBaseInfo.results()
}

func (ctx ScCallInfo) Transfer(color *ScColor, amount int64) {
	ctx.ScBaseInfo.transfer(color, amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostInfo struct {
	ScBaseInfo
}

func (ctx ScPostInfo) Chain(chain *ScAddress) ScPostInfo {
	ctx.request.GetAddress(KeyChain).SetValue(chain)
	return ctx
}

func (ctx ScPostInfo) Post(delay int64) {
	ctx.exec(delay)
}

func (ctx ScPostInfo) Transfer(color *ScColor, amount int64) {
	ctx.ScBaseInfo.transfer(color, amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewInfo struct {
	ScBaseInfo
}

func (ctx ScViewInfo) Results() ScImmutableMap {
	return ctx.ScBaseInfo.results()
}

func (ctx ScViewInfo) View() ScViewInfo {
	ctx.exec(-2)
	return ctx
}
