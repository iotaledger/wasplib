// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

package client

import "strconv"

// used to retrieve any information that is related to colored token balances
type ScBalances struct {
	balances ScImmutableMap
}

// retrieve the balance for the specified token color
func (ctx ScBalances) Balance(color *ScColor) int64 {
	return ctx.balances.GetInt(color).Value()
}

// retrieve a list of all token colors that have a non-zero balance
func (ctx ScBalances) Colors() ScImmutableColorArray {
	return ctx.balances.GetColorArray(KeyColor)
}

// retrieve the color of newly minted tokens
func (ctx ScBalances) Minted() *ScColor {
	return NewScColor(ctx.balances.GetBytes(MINT).Value())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// used to retrieve any information related to the current smart contract
type ScContract struct {
	contract ScImmutableMap
}

// retrieve the chain id of the chain this contract lives on
func (ctx ScContract) Chain() *ScAddress {
	return ctx.contract.GetAddress(KeyChain).Value()
}

// retrieve the agent id of the owner of the chain this contract lives on
func (ctx ScContract) ChainOwner() *ScAgent {
	return ctx.contract.GetAgent(KeyChainOwner).Value()
}

// retrieve the agent id of the creator of this contract
func (ctx ScContract) Creator() *ScAgent {
	return ctx.contract.GetAgent(KeyCreator).Value()
}

// retrieve this contract's description
func (ctx ScContract) Description() string {
	return ctx.contract.GetString(KeyDescription).Value()
}

// retrieve the id of this contract
func (ctx ScContract) Id() *ScAgent {
	return ctx.contract.GetAgent(KeyId).Value()
}

// retrieve this contract's name
func (ctx ScContract) Name() string {
	return ctx.contract.GetString(KeyName).Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMapArray
}

// appends the specified timestamp and data to the timestamped log
func (ctx ScLog) Append(timestamp int64, data []byte) {
	logEntry := ctx.log.GetMap(ctx.log.Length())
	logEntry.GetInt(KeyTimestamp).SetValue(timestamp)
	logEntry.GetBytes(KeyData).SetValue(data)
}

// number of items in the timestamped log
func (ctx ScLog) Length() int32 {
	return ctx.log.Length()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

// decodes the specified base58-encoded string value to its original bytes
func (ctx ScUtility) Base58Decode(value string) []byte {
	decode := ctx.utility.GetString(KeyBase58)
	encode := ctx.utility.GetBytes(KeyBase58)
	decode.SetValue(value)
	return encode.Value()
}

// encodes the specified bytes to a base-58-encoded string
func (ctx ScUtility) Base58Encode(value []byte) string {
	decode := ctx.utility.GetString(KeyBase58)
	encode := ctx.utility.GetBytes(KeyBase58)
	encode.SetValue(value)
	return decode.Value()
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) Hash(value []byte) *ScHash {
	hash := ctx.utility.GetBytes(KeyHash)
	hash.SetValue(value)
	return NewScHash(hash.Value())
}

// generates a random value from 0 to max (exclusive max) using a deterministic RNG
func (ctx ScUtility) Random(max int64) int64 {
	rnd := ctx.utility.GetInt(KeyRandom).Value()
	return int64(uint64(rnd) % uint64(max))
}

// converts an integer to its string representation
func (ctx ScUtility) String(value int64) string {
	return strconv.FormatInt(value, 10)
}

// wrapper for simplified use by hashtypes
func base58Encode(bytes []byte) string {
	return ScCallContext{}.Utility().Base58Encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// shared interface part of ScCallContext and ScViewContext
type ScBaseContext struct {
}

// access the current balances for all token colors
func (ctx ScBaseContext) Balances() ScBalances {
	return ScBalances{Root.GetMap(KeyBalances).Immutable()}
}

// retrieve the agent id of the caller of the smart contract
func (ctx ScBaseContext) Caller() *ScAgent {
	return Root.GetAgent(KeyCaller).Value()
}

// groups contract-related information under one access space
func (ctx ScBaseContext) Contract() ScContract {
	return ScContract{Root.GetMap(KeyContract).Immutable()}
}

// quick check to see if the caller of the smart contract was the specified originator agent
func (ctx ScBaseContext) From(originator *ScAgent) bool {
	return ctx.Caller().Equals(originator)
}

// logs informational text message
func (ctx ScBaseContext) Log(text string) {
	Root.GetString(KeyLog).SetValue(text)
}

// logs error text message and then panics
func (ctx ScBaseContext) Panic(text string) {
	Root.GetString(KeyPanic).SetValue(text)
}

// retrieve parameters passed to the smart contract function that was called
func (ctx ScBaseContext) Params() ScImmutableMap {
	return Root.GetMap(KeyParams).Immutable()
}

// any results returned by the smart contract function call are returned here
func (ctx ScBaseContext) Results() ScMutableMap {
	return Root.GetMap(KeyResults)
}

// deterministic time stamp fixed at the moment of calling the smart contract
func (ctx ScBaseContext) Timestamp() int64 {
	return Root.GetInt(KeyTimestamp).Value()
}

// logs debugging trace text message
func (ctx ScBaseContext) Trace(text string) {
	Root.GetString(KeyTrace).SetValue(text)
}

// access diverse utility functions
func (ctx ScBaseContext) Utility() ScUtility {
	return ScUtility{Root.GetMap(KeyUtility)}
}

// starts a call to a smart contract view function.
func (ctx ScBaseContext) View(function string) ScViewBuilder {
	return ScViewBuilder{newScRequestBuilder(KeyViews, function)}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with mutable access to state
type ScCallContext struct {
	ScBaseContext
}

// starts a call to a smart contract function
func (ctx ScCallContext) Call(function string) ScCallBuilder {
	return ScCallBuilder{newScRequestBuilder(KeyCalls, function)}
}

// starts deployment of a smart contract
func (ctx ScCallContext) Deploy(name string, description string) ScDeployBuilder {
	return NewScDeployBuilder(name, description)
}

// access the incoming balances for all token colors
func (ctx ScCallContext) Incoming() ScBalances {
	return ScBalances{Root.GetMap(KeyIncoming).Immutable()}
}

// starts a (delayed) post to a smart contract function.
func (ctx ScCallContext) Post(function string) ScPostBuilder {
	return ScPostBuilder{newScRequestBuilder(KeyPosts, function)}
}

// signals an event on the chain that entities can register for
func (ctx ScBaseContext) SignalEvent(text string) {
	Root.GetString(KeyEvent).SetValue(text)
}

// access to mutable state storage
func (ctx ScCallContext) State() ScMutableMap {
	return Root.GetMap(KeyState)
}

// access to mutable named timestamped log
func (ctx ScCallContext) TimestampedLog(key MapKey) ScLog {
	return ScLog{Root.GetMap(KeyLogs).GetMapArray(key)}
}

// transfer the specified amount of the specified token color to the specified agent account
func (ctx ScCallContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	NewTransfer(agent).Transfer(color, amount).Send()
}

// start a transfer to the specified Tangle ledger address
func (ctx ScCallContext) TransferToAddress(address *ScAddress) ScTransferBuilder {
	return NewTransferToAddress(address)
}

// start a transfer to the specified cross chain agent account
func (ctx ScCallContext) TransferCrossChain(chain *ScAddress, agent *ScAgent) ScTransferBuilder {
	return NewTransferCrossChain(chain, agent)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
type ScViewContext struct {
	ScBaseContext
}

// access to immutable state storage
func (ctx ScViewContext) State() ScImmutableMap {
	return Root.GetMap(KeyState).Immutable()
}

// access to immutable named timestamped log
func (ctx ScViewContext) TimestampedLog(key MapKey) ScImmutableMapArray {
	return Root.GetMap(KeyLogs).GetMapArray(key).Immutable()
}
