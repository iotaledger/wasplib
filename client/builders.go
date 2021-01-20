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

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransferBuilder struct {
	transfer ScMutableMap
}

// start a transfer to the specified local chain agent account
func NewTransfer(agent *ScAgent) ScTransferBuilder {
	localChain := Root.GetAddress(KeyChain).Value()
	return NewTransferCrossChain(localChain, agent)
}

// start a transfer to a Tangle ledger address
func NewTransferToAddress(address *ScAddress) ScTransferBuilder {
	return NewTransferCrossChain(nil, address.AsAgent())
}

// start a transfer to the specified cross chain agent account
func NewTransferCrossChain(chain *ScAddress, agent *ScAgent) ScTransferBuilder {
	transfers := Root.GetMapArray(KeyTransfers)
	transfer := transfers.GetMap(transfers.Length())
	transfer.GetAgent(KeyAgent).SetValue(agent)
	if chain != nil {
		transfer.GetAddress(KeyChain).SetValue(chain)
	}
	return ScTransferBuilder{transfer: transfer}
}

// sends the complete built transfer to the node
func (ctx ScTransferBuilder) Send() {
	ctx.transfer.GetInt(MINT).SetValue(-1)
}

// transfer the specified amount of tokens of the specified color as part of this transfer
// concatenate one of these for each separate color/amount combination
// amount is supposed to be > 0, unique colors can appear only once
func (ctx ScTransferBuilder) Transfer(color *ScColor, amount int64) ScTransferBuilder {
	ctx.transfer.GetInt(color).SetValue(amount)
	return ctx
}
