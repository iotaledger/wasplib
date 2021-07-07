// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use std::convert::TryInto;

use crate::bytes::*;
use crate::contract::*;
use crate::hashtypes::*;
use crate::host::*;
use crate::immutable::*;
use crate::keys::*;
use crate::mutable::*;

// all access to the objects in host's object tree starts here
pub(crate) static ROOT: ScMutableMap = ScMutableMap { obj_id: 1 };

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// retrieves any information that is related to colored token balances
pub struct ScBalances {
    balances: ScImmutableMap,
}

impl ScBalances {
    // retrieve the balance for the specified token color
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.balances.get_int64(color).value()
    }

    // retrieve an array of all token colors that have a non-zero balance
    pub fn colors(&self) -> ScImmutableColorArray {
        self.balances.get_color_array(&KEY_COLOR)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// passes token transfer information to a function call
#[derive(Clone, Copy)]
pub struct ScTransfers {
    pub(crate) transfers: ScMutableMap,
}

impl ScTransfers {
    // create a new transfers object and initialize it with the specified amount of iotas
    pub fn iotas(amount: i64) -> ScTransfers {
        ScTransfers::new(&ScColor::IOTA, amount)
    }

    // create a new transfers object and initialize it with the specified token transfer
    pub fn new(color: &ScColor, amount: i64) -> ScTransfers {
        let transfer = ScTransfers::new_transfers();
        transfer.set(color, amount);
        transfer
    }

    // create a new transfer object ready to add token transfers
    pub fn new_transfers() -> ScTransfers {
        ScTransfers { transfers: ScMutableMap::new() }
    }

    pub fn none() -> ScTransfers {
        ScTransfers { transfers: ScMutableMap { obj_id: 0 } }
    }

    // create a new transfer object from a balances object
    pub fn new_transfers_from_balances(balances: ScBalances) -> ScTransfers {
        let transfers = ScTransfers::new_transfers();
        let colors = balances.colors();
        for i in 0..colors.length() {
            let color = colors.get_color(i).value();
            transfers.set(&color, balances.balance(&color));
        }
        transfers
    }

    // set the specified colored token transfer in the transfer object
    // note that this will overwrite any previous amount for the specified color
    pub fn set(&self, color: &ScColor, amount: i64) {
        self.transfers.get_int64(color).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// provides access to utility functions that are handled by the host
pub struct ScUtility {
    utility: ScMutableMap,
}

impl ScUtility {
    // decodes the specified base58-encoded string value to its original bytes
    pub fn base58_decode(&self, value: &str) -> Vec<u8> {
        self.utility.call_func(KEY_BASE58_DECODE, value.as_bytes())
    }

    // encodes the specified bytes to a base-58-encoded string
    pub fn base58_encode(&self, value: &[u8]) -> String {
        let result = self.utility.call_func(KEY_BASE58_ENCODE, value);
        unsafe { String::from_utf8_unchecked(result) }
    }

    // retrieves the address for the specified BLS public key
    pub fn bls_address_from_pubkey(&self, pub_key: &[u8]) -> ScAddress {
        let result = self.utility.call_func(KEY_BLS_ADDRESS, pub_key);
        ScAddress::from_bytes(&result)
    }

    // aggregates the specified multiple BLS signatures and public keys into a single one
    pub fn bls_aggregate_signatures(&self, pub_keys_bin: &[&[u8]], sigs_bin: &[&[u8]]) -> (Vec<u8>, Vec<u8>) {
        let mut encode = BytesEncoder::new();
        encode.int32(pub_keys_bin.len() as i32);
        for pub_key in pub_keys_bin {
            encode.bytes(pub_key);
        }
        encode.int32(sigs_bin.len() as i32);
        for sig in sigs_bin {
            encode.bytes(sig);
        }
        let result = self.utility.call_func(KEY_BLS_AGGREGATE, &encode.data());
        let mut decode = BytesDecoder::new(&result);
        return (decode.bytes().to_vec(), decode.bytes().to_vec());
    }

    // checks if the specified BLS signature is valid
    pub fn bls_valid_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        let result = self.utility.call_func(KEY_BLS_VALID, &encode.data());
        (result[0] & 0x01) != 0
    }

    // retrieves the address for the specified ED25519 public key
    pub fn ed25519_address_from_pubkey(&self, pub_key: &[u8]) -> ScAddress {
        let result = self.utility.call_func(KEY_ED25519_ADDRESS, pub_key);
        ScAddress::from_bytes(&result)
    }

    // checks if the specified ED25519 signature is valid
    pub fn ed25519_valid_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        let result = self.utility.call_func(KEY_ED25519_VALID, &encode.data());
        (result[0] & 0x01) != 0
    }

    // hashes the specified value bytes using BLAKE2b hashing and returns the resulting 32-byte hash
    pub fn hash_blake2b(&self, value: &[u8]) -> ScHash {
        let hash = self.utility.call_func(KEY_HASH_BLAKE2B, value);
        ScHash::from_bytes(&hash)
    }

    // hashes the specified value bytes using SHA3 hashing and returns the resulting 32-byte hash
    pub fn hash_sha3(&self, value: &[u8]) -> ScHash {
        let hash = self.utility.call_func(KEY_HASH_SHA3, value);
        ScHash::from_bytes(&hash)
    }

    // calculates 32-bit hash for the specified name string
    pub fn hname(&self, value: &str) -> ScHname {
        let result = self.utility.call_func(KEY_HNAME, value.as_bytes());
        ScHname::from_bytes(&result)
    }

    // generates a random value from 0 to max (exclusive max) using a deterministic RNG
    pub fn random(&self, max: i64) -> i64 {
        let result = self.utility.call_func(KEY_RANDOM, &vec![0_u8; 0]);
        let rnd = i64::from_le_bytes(result.try_into().expect("invalid i64 length"));
        (rnd as u64 % max as u64) as i64
    }
}

// wrapper function for simplified internal access to base58 encoding
pub(crate) fn base58_encode(bytes: &[u8]) -> String {
    ScFuncContext {}.utility().base58_encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// shared interface part of ScFuncContext and ScViewContext
pub trait ScBaseContext {
    // retrieve the agent id of this contract account
    fn account_id(&self) -> ScAgentID {
        ROOT.get_agent_id(&KEY_ACCOUNT_ID).value()
    }

    // access the current balances for all token colors
    fn balances(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_BALANCES).immutable() }
    }

    // retrieve the chain id of the chain this contract lives on
    fn chain_id(&self) -> ScChainID {
        ROOT.get_chain_id(&KEY_CHAIN_ID).value()
    }

    // retrieve the agent id of the owner of the chain this contract lives on
    fn chain_owner_id(&self) -> ScAgentID {
        ROOT.get_agent_id(&KEY_CHAIN_OWNER_ID).value()
    }

    // retrieve the hname of this contract
    fn contract(&self) -> ScHname {
        ROOT.get_hname(&KEY_CONTRACT).value()
    }

    // retrieve the agent id of the creator of this contract
    fn contract_creator(&self) -> ScAgentID {
        ROOT.get_agent_id(&KEY_CONTRACT_CREATOR).value()
    }

    // logs informational text message in the log on the host
    fn log(&self, text: &str) {
        log(text);
    }

    // logs error text message in the log on the host and then panics
    fn panic(&self, text: &str) {
        panic(text);
    }

    // retrieve parameters that were passed to the smart contract function
    fn params(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_PARAMS).immutable()
    }

    // panics with specified message if specified condition is not satisfied
    fn require(&self, cond: bool, msg: &str) {
        if !cond {
            panic(msg);
        }
    }

    // map that holds any results returned by the smart contract function
    fn results(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_RESULTS)
    }

    // deterministic time stamp fixed at the moment of calling the smart contract
    fn timestamp(&self) -> i64 {
        ROOT.get_int64(&KEY_TIMESTAMP).value()
    }

    // logs debugging trace text message in the log on the host
    // similar to log() except this will only show in the log in a special debug mode
    fn trace(&self, text: &str) {
        trace(text);
    }

    // access diverse utility functions provided by the host
    fn utility(&self) -> ScUtility {
        ScUtility { utility: ROOT.get_map(&KEY_UTILITY) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with mutable access to state
#[derive(Clone, Copy)]
pub struct ScFuncContext {}

// reuse shared part of interface
impl ScBaseContext for ScFuncContext {}

impl ScFuncCallContext for ScFuncContext {
    fn can_call_func(&self) {
        panic!("can_call_func");
    }
}

impl ScViewCallContext for ScFuncContext {
    fn can_call_view(&self) {
        panic!("can_call_view");
    }
}

impl ScFuncContext {
    // synchronously calls the specified smart contract function,
    // passing the provided parameters and token transfers to it
    pub fn call(&self, hcontract: ScHname, hfunction: ScHname, params: Option<ScMutableMap>, transfer: Option<ScTransfers>) -> ScImmutableMap {
        let mut encode = BytesEncoder::new();
        encode.hname(&hcontract);
        encode.hname(&hfunction);
        if let Some(params) = params {
            encode.int32(params.map_id());
        } else {
            encode.int32(0);
        }
        if let Some(transfer) = transfer {
            encode.int32(transfer.transfers.map_id());
        } else {
            encode.int32(0);
        }
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());
        ROOT.get_map(&KEY_RETURN).immutable()
    }

    // shorthand to synchronously call a smart contract function of the current contract
    pub fn call_self(&self, hfunction: ScHname, params: Option<ScMutableMap>, transfers: Option<ScTransfers>) -> ScImmutableMap {
        self.call(self.contract(), hfunction, params, transfers)
    }

    // retrieve the agent id of the caller of the smart contract
    pub fn caller(&self) -> ScAgentID { ROOT.get_agent_id(&KEY_CALLER).value() }

    // deploys a new instance of the specified smart contract on the current chain
    // the provided parameters are passed to the smart contract "init" function
    pub fn deploy(&self, program_hash: &ScHash, name: &str, description: &str, params: Option<ScMutableMap>) {
        let mut encode = BytesEncoder::new();
        encode.hash(program_hash);
        encode.string(name);
        encode.string(description);
        if let Some(params) = params {
            encode.int32(params.map_id());
        } else {
            encode.int32(0);
        }
        ROOT.get_bytes(&KEY_DEPLOY).set_value(&encode.data());
    }

    // signals an event on the host that external entities can subscribe to
    pub fn event(&self, text: &str) {
        ROOT.get_string(&KEY_EVENT).set_value(text);
    }

    // access the incoming balances for all token colors
    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_INCOMING).immutable() }
    }

    // retrieve the tokens that were minted in this transaction
    pub fn minted(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_MINTED).immutable() }
    }

    // asynchronously calls the specified smart contract function,
    // passing the provided parameters and token transfers to it
    // it is possible to schedule the call for a later execution by specifying a delay
    pub fn post(&self, chain_id: &ScChainID, hcontract: ScHname, hfunction: ScHname, params: Option<ScMutableMap>, transfer: ScTransfers, delay: i32) {
        let mut encode = BytesEncoder::new();
        encode.chain_id(chain_id);
        encode.hname(&hcontract);
        encode.hname(&hfunction);
        if let Some(params) = &params {
            encode.int32(params.map_id());
        } else {
            encode.int32(0);
        }
        encode.int32(transfer.transfers.map_id());
        encode.int32(delay);
        ROOT.get_bytes(&KEY_POST).set_value(&encode.data());
    }

    // shorthand to asynchronously call a smart contract function of the current contract
    pub fn post_self(&self, hfunction: ScHname, params: Option<ScMutableMap>, transfer: ScTransfers, delay: i32) {
        self.post(&self.chain_id(), self.contract(), hfunction, params, transfer, delay);
    }

    // retrieve the request id of this transaction
    pub fn request_id(&self) -> ScRequestID {
        ROOT.get_request_id(&KEY_REQUEST_ID).value()
    }

    // access mutable state storage on the host
    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    // transfers the specified tokens to the specified Tangle ledger address
    pub fn transfer_to_address(&self, address: &ScAddress, transfer: ScTransfers) {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let tx = transfers.get_map(transfers.length());
        tx.get_address(&KEY_ADDRESS).set_value(address);
        tx.get_int32(&KEY_BALANCES).set_value(transfer.transfers.map_id());
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract view interface which has only immutable access to state
#[derive(Clone, Copy)]
pub struct ScViewContext {}

// reuse shared part of interface
impl ScBaseContext for ScViewContext {}

impl ScViewCallContext for ScViewContext {
    fn can_call_view(&self) {
        panic!("can_call_view");
    }
}

impl ScViewContext {
    // synchronously calls the specified smart contract view,
    // passing the provided parameters to it
    pub fn call(&self, hcontract: ScHname, hfunction: ScHname, params: Option<ScMutableMap>) -> ScImmutableMap {
        let mut encode = BytesEncoder::new();
        encode.hname(&hcontract);
        encode.hname(&hfunction);
        if let Some(params) = params {
            encode.int32(params.map_id());
        } else {
            encode.int32(0);
        }
        encode.int32(0);
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());
        ROOT.get_map(&KEY_RETURN).immutable()
    }

    // shorthand to synchronously call a smart contract view of the current contract
    pub fn call_self(&self, hfunction: ScHname, params: Option<ScMutableMap>) -> ScImmutableMap {
        self.call(self.contract(), hfunction, params)
    }

    // access immutable state storage on the host
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_STATE).immutable()
    }
}
