// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

package wasmlib

import (
	"strconv"
)

type PostRequestParams struct {
	ContractId ScContractId
	Function   ScHname
	Params     *ScMutableMap
	Transfer   *ScTransfers
	Delay      int64
}

// used to retrieve any information that is related to colored token balances
type ScBalances struct {
	balances ScImmutableMap
}

// retrieve the balance for the specified token color
func (ctx ScBalances) Balance(color ScColor) int64 {
	return ctx.balances.GetInt(color).Value()
}

// retrieve a list of all token colors that have a non-zero balance
func (ctx ScBalances) Colors() ScImmutableColorArray {
	return ctx.balances.GetColorArray(KeyColor)
}

// retrieve the color of newly minted tokens
func (ctx ScBalances) Minted() ScColor {
	return NewScColorFromBytes(ctx.balances.GetBytes(MINT).Value())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfers struct {
	transfers ScMutableMap
}

// special constructor for simplifying single transfers
func NewScTransfer(color ScColor, amount int64) ScTransfers {
	transfer := NewScTransfers()
	transfer.Add(color, amount)
	return transfer
}

func NewScTransfers() ScTransfers {
	return ScTransfers{transfers: *NewScMutableMap()}
}

func NewScTransfersFromBalances(balances ScBalances) ScTransfers {
	transfers := NewScTransfers()
	colors := balances.Colors()
	length := colors.Length()
	for i := int32(0); i < length; i++ {
		color := colors.GetColor(i).Value()
		transfers.Add(color, balances.Balance(color))
	}
	return transfers
}

// transfers the specified amount of tokens of the specified color
func (ctx ScTransfers) Add(color ScColor, amount int64) {
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

func (ctx ScUtility) BlsAddressFromPubKey(pubKey []byte) ScAddress {
	ctx.utility.GetBytes(KeyBlsAddress).SetValue(pubKey)
	return ctx.utility.GetAddress(KeyAddress).Value()
}

func (ctx ScUtility) BlsAggregateSignatures(pubKeys [][]byte, sigs [][]byte) ([]byte, []byte) {
	encode := NewBytesEncoder()
	encode.Int(int64(len(pubKeys)))
	for _, pubKey := range pubKeys {
		encode.Bytes(pubKey)
	}
	encode.Int(int64(len(sigs)))
	for _, sig := range sigs {
		encode.Bytes(sig)
	}
	aggregator := ctx.utility.GetBytes(KeyBlsAggregate)
	aggregator.SetValue(encode.Data())
	decode := NewBytesDecoder(aggregator.Value())
	return decode.Bytes(), decode.Bytes()
}

func (ctx ScUtility) BlsValidSignature(data []byte, pubKey []byte, signature []byte) bool {
	bytes := NewBytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature).Data()
	ctx.utility.GetBytes(KeyBlsValid).SetValue(bytes)
	return ctx.utility.GetInt(KeyValid).Value() != 0
}

func (ctx ScUtility) Ed25519AddressFromPubKey(pubKey []byte) ScAddress {
	ctx.utility.GetBytes(KeyEd25519Address).SetValue(pubKey)
	return ctx.utility.GetAddress(KeyAddress).Value()
}

func (ctx ScUtility) Ed25519ValidSignature(data []byte, pubKey []byte, signature []byte) bool {
	bytes := NewBytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature).Data()
	ctx.utility.GetBytes(KeyEd25519Valid).SetValue(bytes)
	return ctx.utility.GetInt(KeyValid).Value() != 0
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) HashBlake2b(value []byte) ScHash {
	hash := ctx.utility.GetBytes(KeyHashBlake2b)
	hash.SetValue(value)
	return NewScHashFromBytes(hash.Value())
}

// hashes the specified value bytes using sha3 hashing and returns the resulting 32-byte hash
func (ctx ScUtility) HashSha3(value []byte) ScHash {
	hash := ctx.utility.GetBytes(KeyHashSha3)
	hash.SetValue(value)
	return NewScHashFromBytes(hash.Value())
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) Hname(value string) ScHname {
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
	return ScFuncContext{}.Utility().Base58Encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// shared interface part of ScFuncContext and ScViewContext
type ScBaseContext struct {
}

// access the current balances for all token colors
func (ctx ScBaseContext) Balances() ScBalances {
	return ScBalances{Root.GetMap(KeyBalances).Immutable()}
}

// retrieve the agent id of the owner of the chain this contract lives on
func (ctx ScBaseContext) ChainOwnerId() ScAgentId {
	return Root.GetAgentId(KeyChainOwnerId).Value()
}

// retrieve the agent id of the creator of this contract
func (ctx ScBaseContext) ContractCreator() ScAgentId {
	return Root.GetAgentId(KeyContractCreator).Value()
}

// retrieve the id of this contract
func (ctx ScBaseContext) ContractId() ScContractId {
	return Root.GetContractId(KeyContractId).Value()
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
type ScFuncContext struct {
	ScBaseContext
}

// calls a smart contract function
func (ctx ScFuncContext) Call(hContract ScHname, hFunction ScHname, params *ScMutableMap, transfer *ScTransfers) ScImmutableMap {
	encode := NewBytesEncoder()
	encode.Hname(hContract)
	encode.Hname(hFunction)
	if params != nil {
		encode.Int(int64(params.objId))
	} else {
		encode.Int(0)
	}
	if transfer != nil {
		encode.Int(int64(transfer.transfers.objId))
	} else {
		encode.Int(0)
	}
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	return Root.GetMap(KeyReturn).Immutable()
}

// retrieve the agent id of the caller of the smart contract
func (ctx ScFuncContext) Caller() ScAgentId {
	return Root.GetAgentId(KeyCaller).Value()
}

// calls a smart contract function on the current contract
func (ctx ScFuncContext) CallSelf(hFunction ScHname, params *ScMutableMap, transfer *ScTransfers) ScImmutableMap {
	return ctx.Call(ctx.ContractId().Hname(), hFunction, params, transfer)
}

// deploys a smart contract
func (ctx ScFuncContext) Deploy(programHash ScHash, name string, description string, params *ScMutableMap) {
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

// signals an event on the node that external entities can subscribe to
func (ctx ScBaseContext) Event(text string) {
	Root.GetString(KeyEvent).SetValue(text)
}

// access the incoming balances for all token colors
func (ctx ScFuncContext) Incoming() ScBalances {
	return ScBalances{Root.GetMap(KeyIncoming).Immutable()}
}

// (delayed) posts a smart contract function
func (ctx ScFuncContext) Post(par *PostRequestParams) {
	encode := NewBytesEncoder()
	encode.ContractId(par.ContractId)
	encode.Hname(par.Function)
	if par.Params != nil {
		encode.Int(int64(par.Params.objId))
	} else {
		encode.Int(0)
	}
	if par.Transfer != nil {
		encode.Int(int64(par.Transfer.transfers.objId))
	} else {
		encode.Int(0)
	}
	encode.Int(par.Delay)
	Root.GetBytes(KeyPost).SetValue(encode.Data())
}

// access to mutable state storage
func (ctx ScFuncContext) State() ScMutableMap {
	return Root.GetMap(KeyState)
}

// transfer colored token amounts to the specified Tangle ledger address
func (ctx ScFuncContext) TransferToAddress(address ScAddress, transfer ScTransfers) {
	transfers := Root.GetMapArray(KeyTransfers)
	tx := transfers.GetMap(transfers.Length())
	tx.GetAddress(KeyAddress).SetValue(address)
	tx.GetInt(KeyBalances).SetValue(int64(transfer.transfers.objId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
type ScViewContext struct {
	ScBaseContext
}

// calls a smart contract function
func (ctx ScViewContext) Call(contract ScHname, function ScHname, params *ScMutableMap) ScImmutableMap {
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

// calls a smart contract function on the current contract
func (ctx ScViewContext) CallSelf(function ScHname, params *ScMutableMap) ScImmutableMap {
	return ctx.Call(ctx.ContractId().Hname(), function, params)
}

// access to immutable state storage
func (ctx ScViewContext) State() ScImmutableMap {
	return Root.GetMap(KeyState).Immutable()
}
