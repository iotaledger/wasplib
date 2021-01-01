// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallBuilder struct {
	ScRequestBuilder
}

// execute the call request when finished building
func (ctx ScCallBuilder) Call() ScCallBuilder {
	ctx.exec(-1)
	return ctx
}

// specify a different contract the call request is for
func (ctx ScCallBuilder) Contract(contract string) ScCallBuilder {
	ctx.ScRequestBuilder.contract(contract)
	return ctx
}

// provide parameters for the call request
func (ctx ScCallBuilder) Params() ScMutableMap {
	return ctx.ScRequestBuilder.params()
}

// access call request results after executing the call
func (ctx ScCallBuilder) Results() ScImmutableMap {
	return ctx.ScRequestBuilder.results()
}

// transfer tokens as part of the call request
func (ctx ScCallBuilder) Transfer(color *ScColor, amount int64) ScCallBuilder {
	ctx.ScRequestBuilder.transfer(color, amount)
	return ctx
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScDeployBuilder struct {
	deploy ScMutableMap
}

// start deployment of a smart contract with the specified name and description
func NewScDeployBuilder(name string, description string) ScDeployBuilder {
	deploys := root.GetMapArray(KeyDeploys)
	deploy := deploys.GetMap(deploys.Length())
	deploy.GetString(KeyName).SetValue(name)
	deploy.GetString(KeyDescription).SetValue(description)
	return ScDeployBuilder{deploy}
}

// execute the deployment of the smart contract with the specified program hash
func (ctx ScDeployBuilder) Deploy(programHash *ScHash) {
	ctx.deploy.GetHash(KeyHash).SetValue(programHash)
}

// provide parameters for the deployment's 'init' call
func (ctx ScDeployBuilder) Params() ScMutableMap {
	return ctx.deploy.GetMap(KeyParams)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostBuilder struct {
	ScRequestBuilder
}

// specify another chain to post the request to
func (ctx ScPostBuilder) Chain(chain *ScAddress) ScPostBuilder {
	ctx.request.GetAddress(KeyChain).SetValue(chain)
	return ctx
}

// specify a different contract the post request is for
func (ctx ScPostBuilder) Contract(contract string) ScPostBuilder {
	ctx.ScRequestBuilder.contract(contract)
	return ctx
}

// provide parameters for the post request
func (ctx ScPostBuilder) Params() ScMutableMap {
	return ctx.ScRequestBuilder.params()
}

// execute the post request when finished building with the specified delay
func (ctx ScPostBuilder) Post(delay int64) {
	ctx.exec(delay)
}

// transfer tokens as part of the post request
func (ctx ScPostBuilder) Transfer(color *ScColor, amount int64) ScPostBuilder {
	ctx.ScRequestBuilder.transfer(color, amount)
	return ctx
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewBuilder struct {
	ScRequestBuilder
}

// specify a different contract the view request is for
func (ctx ScViewBuilder) Contract(contract string) ScViewBuilder {
	ctx.ScRequestBuilder.contract(contract)
	return ctx
}

// provide parameters for the view request
func (ctx ScViewBuilder) Params() ScMutableMap {
	return ctx.ScRequestBuilder.params()
}

// access view request results after executing the call
func (ctx ScViewBuilder) Results() ScImmutableMap {
	return ctx.ScRequestBuilder.results()
}

// execute the view request when finished building
func (ctx ScViewBuilder) View() ScViewBuilder {
	ctx.exec(-2)
	return ctx
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestBuilder struct {
	request ScMutableMap
}

// start building a request for the specified smart contract function
func newScRequestBuilder(key MapKey, function string) ScRequestBuilder {
	requests := root.GetMapArray(key)
	request := requests.GetMap(requests.Length())
	request.GetString(KeyFunction).SetValue(function)
	return ScRequestBuilder{request}
}

// specify a different contract the request is for
func (ctx ScRequestBuilder) contract(contract string) {
	ctx.request.GetString(KeyContract).SetValue(contract)
}

// execute the request
func (ctx ScRequestBuilder) exec(delay int64) {
	ctx.request.GetInt(KeyDelay).SetValue(delay)
}

// provide parameters for the request
func (ctx ScRequestBuilder) params() ScMutableMap {
	return ctx.request.GetMap(KeyParams)
}

// access request results after executing the request
func (ctx ScRequestBuilder) results() ScImmutableMap {
	return ctx.request.GetMap(KeyResults).Immutable()
}

// transfer tokens as part of the request
func (ctx ScRequestBuilder) transfer(color *ScColor, amount int64) {
	transfers := ctx.request.GetMap(KeyTransfers)
	transfers.GetInt(color).SetValue(amount)
}
