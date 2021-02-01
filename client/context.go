// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

package client

import (
	"strconv"
)

type PostRequestParams struct {
	Contract *ScContractId
	Function Hname
	Params   *ScMutableMap
	Transfer balances
	Delay    int64
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type balances interface {
	mapId() int32
}

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

// implements Balances interface
func (ctx ScBalances) mapId() int32 {
	return ctx.balances.objId
}

// retrieve the color of newly minted tokens
func (ctx ScBalances) Minted() *ScColor {
	return NewScColor(ctx.balances.GetBytes(MINT).Value())
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

type ScTransfers struct {
	transfers ScMutableMap
}

// special constructor for simplifying single transfers
func NewScTransfer(color *ScColor, amount int64) ScTransfers {
	balance := NewScTransfers()
	balance.Add(color, amount)
	return balance
}

func NewScTransfers() ScTransfers {
	return ScTransfers{transfers: *NewScMutableMap()}
}

// implements Balances interface
func (ctx ScTransfers) mapId() int32 {
	return ctx.transfers.objId
}

// transfers the specified amount of tokens of the specified color
func (ctx ScTransfers) Add(color *ScColor, amount int64) {
	ctx.transfers.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

// decodes the specified base58-encoded string value to its original bytes
func (ctx ScUtility) Base58Decode(value string) []byte {
	ctx.utility.GetString(KeyBase58String).SetValue(value)
	return ctx.utility.GetBytes(KeyBase58Bytes).Value()
}

// encodes the specified bytes to a base-58-encoded string
func (ctx ScUtility) Base58Encode(value []byte) string {
	ctx.utility.GetBytes(KeyBase58Bytes).SetValue(value)
	return ctx.utility.GetString(KeyBase58String).Value()
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) Hash(value []byte) *ScHash {
	hash := ctx.utility.GetBytes(KeyHash)
	hash.SetValue(value)
	return NewScHash(hash.Value())
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) Hname(value string) Hname {
	ctx.utility.GetString(KeyName).SetValue(value)
	return ctx.utility.GetHname(KeyHname).Value()
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

// retrieve the agent id of the owner of the chain this contract lives on
func (ctx ScBaseContext) ChainOwner() *ScAgent {
	return Root.GetAgent(KeyChainOwner).Value()
}

// retrieve the agent id of the creator of this contract
func (ctx ScBaseContext) ContractCreator() *ScAgent {
	return Root.GetAgent(KeyCreator).Value()
}

// retrieve the id of this contract
func (ctx ScBaseContext) ContractId() *ScContractId {
	return Root.GetContractId(KeyId).Value()
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

// panics if condition is not satisfied
func (ctx ScBaseContext) Require(cond bool, msg string) {
	if !cond {
		ctx.Panic(msg)
	}
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

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with mutable access to state
type ScCallContext struct {
	ScBaseContext
}

//TODO view immutable state on Wasp
//TODO parameter type checks

// calls a smart contract function
func (ctx ScCallContext) Call(contract Hname, function Hname, params *ScMutableMap, transfer balances) ScImmutableMap {
	encode := NewBytesEncoder()
	encode.Hname(contract)
	encode.Hname(function)
	if params != nil {
		encode.Int(int64(params.objId))
	} else {
		encode.Int(0)
	}
	if transfer != nil {
		encode.Int(int64(transfer.mapId()))
	} else {
		encode.Int(0)
	}
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	return Root.GetMap(KeyReturn).Immutable()
}

// deploys a smart contract
func (ctx ScCallContext) Deploy(programHash *ScHash, name string, description string, params *ScMutableMap) {
	encode := NewBytesEncoder()
	encode.Hash(programHash)
	encode.String(name)
	encode.String(description)
	if params != nil {
		encode.Int(int64(params.objId))
	} else {
		encode.Int(0)
	}
	Root.GetBytes(KeyDeploy).SetValue(encode.Data())
}

// access the incoming balances for all token colors
func (ctx ScCallContext) Incoming() ScBalances {
	return ScBalances{Root.GetMap(KeyIncoming).Immutable()}
}

// (delayed) posts a smart contract function
func (ctx ScCallContext) Post(par *PostRequestParams) {
	encode := NewBytesEncoder()
	encode.ContractId(par.Contract)
	encode.Hname(par.Function)
	if par.Params != nil {
		encode.Int(int64(par.Params.objId))
	} else {
		encode.Int(0)
	}
	if par.Transfer != nil {
		encode.Int(int64(par.Transfer.mapId()))
	} else {
		encode.Int(0)
	}
	encode.Int(par.Delay)
	Root.GetBytes(KeyPost).SetValue(encode.Data())
}

// signals an event on the node that external entities can subscribe to
func (ctx ScBaseContext) Event(text string) {
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

// transfer colored token amounts to the specified Tangle ledger address
func (ctx ScCallContext) TransferToAddress(address *ScAddress, transfer balances) {
	transfers := Root.GetMapArray(KeyTransfers)
	tx := transfers.GetMap(transfers.Length())
	tx.GetAddress(KeyAddress).SetValue(address)
	tx.GetInt(KeyBalances).SetValue(int64(transfer.mapId()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
type ScViewContext struct {
	ScBaseContext
}

// calls a smart contract function
func (ctx ScViewContext) Call(contract Hname, function Hname, params *ScMutableMap) ScImmutableMap {
	encode := NewBytesEncoder()
	encode.Hname(contract)
	encode.Hname(function)
	if params != nil {
		encode.Int(int64(params.objId))
	} else {
		encode.Int(0)
	}
	encode.Int(0)
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	return Root.GetMap(KeyReturn).Immutable()
}

// access to immutable state storage
func (ctx ScViewContext) State() ScImmutableMap {
	return Root.GetMap(KeyState).Immutable()
}

// access to immutable named timestamped log
func (ctx ScViewContext) TimestampedLog(key MapKey) ScImmutableMapArray {
	return Root.GetMap(KeyLogs).GetMapArray(key).Immutable()
}
