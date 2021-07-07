// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

package wasmlib

import (
	"encoding/binary"
	"strconv"
)

// used to retrieve any information that is related to colored token balances
type ScBalances struct {
	balances ScImmutableMap
}

// retrieve the balance for the specified token color
func (ctx ScBalances) Balance(color ScColor) int64 {
	return ctx.balances.GetInt64(color).Value()
}

// retrieve a list of all token colors that have a non-zero balance
func (ctx ScBalances) Colors() ScImmutableColorArray {
	return ctx.balances.GetColorArray(KeyColor)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfers struct {
	transfers ScMutableMap
}

// special constructor for simplifying iota transfers
func NewScTransferIotas(amount int64) ScTransfers {
	return NewScTransfer(IOTA, amount)
}

// special constructor for simplifying single transfers
func NewScTransfer(color ScColor, amount int64) ScTransfers {
	transfer := NewScTransfers()
	transfer.Set(color, amount)
	return transfer
}

func NewScTransfers() ScTransfers {
	return ScTransfers{transfers: *NewScMutableMap()}
}

func NoScTransfers() ScTransfers {
	return ScTransfers{}
}

func NewScTransfersFromBalances(balances ScBalances) ScTransfers {
	transfers := NewScTransfers()
	colors := balances.Colors()
	length := colors.Length()
	for i := int32(0); i < length; i++ {
		color := colors.GetColor(i).Value()
		transfers.Set(color, balances.Balance(color))
	}
	return transfers
}

// transfers the specified amount of tokens of the specified color
func (ctx ScTransfers) Set(color ScColor, amount int64) {
	ctx.transfers.GetInt64(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

// decodes the specified base58-encoded string value to its original bytes
func (ctx ScUtility) Base58Decode(value string) []byte {
	return ctx.utility.CallFunc(KeyBase58Decode, []byte(value))
}

// encodes the specified bytes to a base-58-encoded string
func (ctx ScUtility) Base58Encode(value []byte) string {
	return string(ctx.utility.CallFunc(KeyBase58Encode, value))
}

func (ctx ScUtility) BlsAddressFromPubKey(pubKey []byte) ScAddress {
	result := ctx.utility.CallFunc(KeyBlsAddress, pubKey)
	return NewScAddressFromBytes(result)
}

func (ctx ScUtility) BlsAggregateSignatures(pubKeys, sigs [][]byte) ([]byte, []byte) {
	encode := NewBytesEncoder()
	encode.Int32(int32(len(pubKeys)))
	for _, pubKey := range pubKeys {
		encode.Bytes(pubKey)
	}
	encode.Int32(int32(len(sigs)))
	for _, sig := range sigs {
		encode.Bytes(sig)
	}
	result := ctx.utility.CallFunc(KeyBlsAggregate, encode.Data())
	decode := NewBytesDecoder(result)
	return decode.Bytes(), decode.Bytes()
}

func (ctx ScUtility) BlsValidSignature(data, pubKey, signature []byte) bool {
	encode := NewBytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature)
	result := ctx.utility.CallFunc(KeyBlsValid, encode.Data())
	return len(result) != 0
}

func (ctx ScUtility) Ed25519AddressFromPubKey(pubKey []byte) ScAddress {
	result := ctx.utility.CallFunc(KeyEd25519Address, pubKey)
	return NewScAddressFromBytes(result)
}

func (ctx ScUtility) Ed25519ValidSignature(data, pubKey, signature []byte) bool {
	encode := NewBytesEncoder().Bytes(data).Bytes(pubKey).Bytes(signature)
	result := ctx.utility.CallFunc(KeyEd25519Valid, encode.Data())
	return len(result) != 0
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) HashBlake2b(value []byte) ScHash {
	result := ctx.utility.CallFunc(KeyHashBlake2b, value)
	return NewScHashFromBytes(result)
}

// hashes the specified value bytes using sha3 hashing and returns the resulting 32-byte hash
func (ctx ScUtility) HashSha3(value []byte) ScHash {
	result := ctx.utility.CallFunc(KeyHashSha3, value)
	return NewScHashFromBytes(result)
}

// hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
func (ctx ScUtility) Hname(value string) ScHname {
	result := ctx.utility.CallFunc(KeyHname, []byte(value))
	return NewScHnameFromBytes(result)
}

// generates a random value from 0 to max (exclusive max) using a deterministic RNG
func (ctx ScUtility) Random(max int64) int64 {
	result := ctx.utility.CallFunc(KeyRandom, nil)
	rnd := binary.LittleEndian.Uint64(result)
	return int64(rnd % uint64(max))
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
type ScBaseContext struct{}

// retrieve the agent id of this contract account
func (ctx ScBaseContext) AccountID() ScAgentID {
	return Root.GetAgentID(KeyAccountID).Value()
}

// access the current balances for all token colors
func (ctx ScBaseContext) Balances() ScBalances {
	return ScBalances{Root.GetMap(KeyBalances).Immutable()}
}

// retrieve the chain id of the chain this contract lives on
func (ctx ScBaseContext) ChainID() ScChainID {
	return Root.GetChainID(KeyChainID).Value()
}

// retrieve the agent id of the owner of the chain this contract lives on
func (ctx ScBaseContext) ChainOwnerID() ScAgentID {
	return Root.GetAgentID(KeyChainOwnerID).Value()
}

// retrieve the hname of this contract
func (ctx ScBaseContext) Contract() ScHname {
	return Root.GetHname(KeyContract).Value()
}

// retrieve the agent id of the creator of this contract
func (ctx ScBaseContext) ContractCreator() ScAgentID {
	return Root.GetAgentID(KeyContractCreator).Value()
}

// logs informational text message
func (ctx ScBaseContext) Log(text string) {
	Log(text)
}

// logs error text message and then panics
func (ctx ScBaseContext) Panic(text string) {
	Panic(text)
}

// retrieve parameters passed to the smart contract function that was called
func (ctx ScBaseContext) Params() ScImmutableMap {
	return Root.GetMap(KeyParams).Immutable()
}

// panics if condition is not satisfied
func (ctx ScBaseContext) Require(cond bool, msg string) {
	if !cond {
		Panic(msg)
	}
}

// any results returned by the smart contract function call are returned here
func (ctx ScBaseContext) Results() ScMutableMap {
	return Root.GetMap(KeyResults)
}

// deterministic time stamp fixed at the moment of calling the smart contract
func (ctx ScBaseContext) Timestamp() int64 {
	return Root.GetInt64(KeyTimestamp).Value()
}

// logs debugging trace text message
func (ctx ScBaseContext) Trace(text string) {
	Trace(text)
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

var (
	_ ScViewCallContext = &ScViewContext{}
	_ ScFuncCallContext = &ScFuncContext{}
)

// calls a smart contract function
func (ctx ScFuncContext) Call(hContract, hFunction ScHname, params *ScMutableMap, transfer *ScTransfers) ScImmutableMap {
	encode := NewBytesEncoder()
	encode.Hname(hContract)
	encode.Hname(hFunction)
	if params != nil {
		encode.Int32(params.objID)
	} else {
		encode.Int32(0)
	}
	if transfer != nil {
		encode.Int32(transfer.transfers.objID)
	} else {
		encode.Int32(0)
	}
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	return Root.GetMap(KeyReturn).Immutable()
}

// retrieve the agent id of the caller of the smart contract
func (ctx ScFuncContext) Caller() ScAgentID {
	return Root.GetAgentID(KeyCaller).Value()
}

// calls a smart contract function on the current contract
func (ctx ScFuncContext) CallSelf(hFunction ScHname, params *ScMutableMap, transfer *ScTransfers) ScImmutableMap {
	return ctx.Call(ctx.Contract(), hFunction, params, transfer)
}

func (ctx ScFuncContext) CanCallFunc() {
	panic("CanCallFunc")
}

func (ctx ScFuncContext) CanCallView() {
	panic("CanCallView")
}

// deploys a smart contract
func (ctx ScFuncContext) Deploy(programHash ScHash, name, description string, params *ScMutableMap) {
	encode := NewBytesEncoder()
	encode.Hash(programHash)
	encode.String(name)
	encode.String(description)
	if params != nil {
		encode.Int32(params.objID)
	} else {
		encode.Int32(0)
	}
	Root.GetBytes(KeyDeploy).SetValue(encode.Data())
}

// signals an event on the node that external entities can subscribe to
func (ctx ScFuncContext) Event(text string) {
	Root.GetString(KeyEvent).SetValue(text)
}

func (ctx ScFuncContext) Host() ScHost {
	return nil
}

// access the incoming balances for all token colors
func (ctx ScFuncContext) Incoming() ScBalances {
	return ScBalances{Root.GetMap(KeyIncoming).Immutable()}
}

// retrieve the tokens that were minted in this transaction
func (ctx ScFuncContext) Minted() ScBalances {
	return ScBalances{Root.GetMap(KeyMinted).Immutable()}
}

// (delayed) posts a smart contract function
func (ctx ScFuncContext) Post(chainID ScChainID, hContract, hFunction ScHname, params *ScMutableMap, transfer ScTransfers, delay int32) {
	encode := NewBytesEncoder()
	encode.ChainID(chainID)
	encode.Hname(hContract)
	encode.Hname(hFunction)
	if params != nil {
		encode.Int32(params.objID)
	} else {
		encode.Int32(0)
	}
	encode.Int32(transfer.transfers.objID)
	encode.Int32(delay)
	Root.GetBytes(KeyPost).SetValue(encode.Data())
}

func (ctx ScFuncContext) PostSelf(hFunction ScHname, params *ScMutableMap, transfer ScTransfers, delay int32) {
	ctx.Post(ctx.ChainID(), ctx.Contract(), hFunction, params, transfer, delay)
}

// retrieve the request id of this transaction
func (ctx ScFuncContext) RequestID() ScRequestID {
	return Root.GetRequestID(KeyRequestID).Value()
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
	tx.GetInt32(KeyBalances).SetValue(transfer.transfers.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
type ScViewContext struct {
	ScBaseContext
}

var _ ScViewCallContext = &ScViewContext{}

// calls a smart contract function
func (ctx ScViewContext) Call(contract, function ScHname, params *ScMutableMap) ScImmutableMap {
	encode := NewBytesEncoder()
	encode.Hname(contract)
	encode.Hname(function)
	if params != nil {
		encode.Int32(params.objID)
	} else {
		encode.Int32(0)
	}
	encode.Int32(0)
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	return Root.GetMap(KeyReturn).Immutable()
}

// calls a smart contract function on the current contract
func (ctx ScViewContext) CallSelf(function ScHname, params *ScMutableMap) ScImmutableMap {
	return ctx.Call(ctx.Contract(), function, params)
}

func (ctx ScViewContext) CanCallView() {
	panic("CanCallView")
}

// access to immutable state storage
func (ctx ScViewContext) State() ScImmutableMap {
	return Root.GetMap(KeyState).Immutable()
}
