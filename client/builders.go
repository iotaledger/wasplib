// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScDeployBuilder struct {
	deploy ScMutableMap
}

// start deployment of a smart contract with the specified name and description
func NewScDeployBuilder(name string, description string) ScDeployBuilder {
	deploys := Root.GetMapArray(KeyDeploys)
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
