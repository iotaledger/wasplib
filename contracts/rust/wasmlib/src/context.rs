// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use crate::bytes::*;
use crate::hashtypes::*;
use crate::immutable::*;
use crate::keys::*;
use crate::mutable::*;

// all access to the objects in host's object tree starts here
pub(crate) static ROOT: ScMutableMap = ScMutableMap { obj_id: 1 };

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// used to retrieve any information that is related to colored token balances
pub struct ScBalances {
    balances: ScImmutableMap,
}

impl ScBalances {
    // retrieve the balance for the specified token color
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.balances.get_int64(color).value()
    }

    // retrieve a list of all token colors that have a non-zero balance
    pub fn colors(&self) -> ScImmutableColorArray {
        self.balances.get_color_array(&KEY_COLOR)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// used to pass token transfer information to a function call
pub struct ScTransfers {
    transfers: ScMutableMap,
}

impl ScTransfers {
    // create a new transfers object and initialize it with the specified token transfer
    pub fn new(color: &ScColor, amount: i64) -> ScTransfers {
        let transfer = ScTransfers::new_transfers();
        transfer.add(color, amount);
        transfer
    }

    // create a new transfer object ready to add token transfers
    pub fn new_transfers() -> ScTransfers {
        ScTransfers { transfers: ScMutableMap::new() }
    }

    // create a new transfer object ready to add token transfers
    pub fn new_transfers_from_balances(balances: ScBalances) -> ScTransfers {
        let transfers = ScTransfers::new_transfers();
        let colors = balances.colors();
        for i in 0..colors.length() {
            let color = colors.get_color(i).value();
            transfers.add(&color, balances.balance(&color));
        }
        transfers
    }

    // add the specified token transfer to the transfer object
    pub fn add(&self, color: &ScColor, amount: i64) {
        self.transfers.get_int64(color).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// provide access to utility functions that are handled by the host
pub struct ScUtility {
    utility: ScMutableMap,
}

impl ScUtility {
    // decodes the specified base58-encoded string value to its original bytes
    pub fn base58_decode(&self, value: &str) -> Vec<u8> {
        self.utility.get_string(&KEY_BASE58_STRING).set_value(value);
        self.utility.get_bytes(&KEY_BASE58_BYTES).value()
    }

    // encodes the specified bytes to a base-58-encoded string
    pub fn base58_encode(&self, value: &[u8]) -> String {
        self.utility.get_bytes(&KEY_BASE58_BYTES).set_value(value);
        self.utility.get_string(&KEY_BASE58_STRING).value()
    }

    // retrieves the address for the specified ED25519 public key
    // retrieves the address for the specified BLS public key
    pub fn bls_address_from_pubkey(&self, pub_key: &[u8]) -> ScAddress {
        self.utility.get_bytes(&KEY_BLS_ADDRESS).set_value(pub_key);
        self.utility.get_address(&KEY_ADDRESS).value()
    }

    // aggregates the specified multiple BLS signatures and public keys into a single one
    pub fn bls_aggregate_signatures(&self, pub_keys_bin: &[&[u8]], sigs_bin: &[&[u8]]) -> (Vec<u8>, Vec<u8>) {
        let mut encode = BytesEncoder::new();
        encode.int64(pub_keys_bin.len() as i64);
        for pub_key in pub_keys_bin {
            encode.bytes(pub_key);
        }
        encode.int64(sigs_bin.len() as i64);
        for sig in sigs_bin {
            encode.bytes(sig);
        }
        let aggregator = self.utility.get_bytes(&KEY_BLS_AGGREGATE);
        aggregator.set_value(&encode.data());
        let aggregated = aggregator.value();
        let mut decode = BytesDecoder::new(&aggregated);
        return (decode.bytes().to_vec(), decode.bytes().to_vec());
    }

    // checks if the specified BLS signature is valid
    pub fn bls_valid_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        self.utility.get_bytes(&KEY_BLS_VALID).set_value(&encode.data());
        self.utility.get_int64(&KEY_VALID).value() != 0
    }

    pub fn ed25519_address_from_pubkey(&self, pub_key: &[u8]) -> ScAddress {
        self.utility.get_bytes(&KEY_ED25519_ADDRESS).set_value(pub_key);
        self.utility.get_address(&KEY_ADDRESS).value()
    }

    // checks if the specified ED25519 signature is valid
    pub fn ed25519_valid_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        self.utility.get_bytes(&KEY_ED25519_VALID).set_value(&encode.data());
        self.utility.get_int64(&KEY_VALID).value() != 0
    }

    // hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
    pub fn hash_blake2b(&self, value: &[u8]) -> ScHash {
        let hash = self.utility.get_bytes(&KEY_HASH_BLAKE2B);
        hash.set_value(value);
        ScHash::from_bytes(&hash.value())
    }

    // hashes the specified value bytes using sha3 hashing and returns the resulting 32-byte hash
    pub fn hash_sha3(&self, value: &[u8]) -> ScHash {
        let hash = self.utility.get_bytes(&KEY_HASH_SHA3);
        hash.set_value(value);
        ScHash::from_bytes(&hash.value())
    }

    // calculates 32-bit hash for the specified name string
    pub fn hname(&self, value: &str) -> ScHname {
        self.utility.get_string(&KEY_NAME).set_value(value);
        ScHname::from_bytes(&self.utility.get_bytes(&KEY_HNAME).value())
    }

    // generates a random value from 0 to max (exclusive max) using a deterministic RNG
    pub fn random(&self, max: i64) -> i64 {
        let rnd = self.utility.get_int64(&KEY_RANDOM).value();
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
    // access the current balances for all token colors
    fn balances(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_BALANCES).immutable() }
    }

    // retrieve the agent id of the owner of the chain this contract lives on
    fn chain_owner_id(&self) -> ScAgentId {
        ROOT.get_agent_id(&KEY_CHAIN_OWNER_ID).value()
    }

    // retrieve the agent id of the creator of this contract
    fn contract_creator(&self) -> ScAgentId {
        ROOT.get_agent_id(&KEY_CONTRACT_CREATOR).value()
    }

    // retrieve the id of this contract
    fn contract_id(&self) -> ScContractId {
        ROOT.get_contract_id(&KEY_CONTRACT_ID).value()
    }

    // logs informational text message
    fn log(&self, text: &str) {
        ROOT.get_string(&KEY_LOG).set_value(text)
    }

    // logs error text message and then panics
    fn panic(&self, text: &str) {
        ROOT.get_string(&KEY_PANIC).set_value(text)
    }

    // retrieve parameters that were passed to the smart contract function
    fn params(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_PARAMS).immutable()
    }

    // panics with specified message if specified condition is not satisfied
    fn require(&self, cond: bool, msg: &str) {
        if !cond {
            self.panic(msg)
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

    // logs debugging trace text message
    fn trace(&self, text: &str) {
        ROOT.get_string(&KEY_TRACE).set_value(text)
    }

    // access diverse utility functions
    fn utility(&self) -> ScUtility {
        ScUtility { utility: ROOT.get_map(&KEY_UTILITY) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with mutable access to state
pub struct ScFuncContext {}

// shared part of interface
impl ScBaseContext for ScFuncContext {}

impl ScFuncContext {
    // synchronously calls the specified smart contract function,
    // passing the provided parameters and token transfers to it
    pub fn call(&self, hcontract: ScHname, hfunction: ScHname, params: Option<ScMutableMap>, transfer: Option<ScTransfers>) -> ScImmutableMap {
        let mut encode = BytesEncoder::new();
        encode.hname(&hcontract);
        encode.hname(&hfunction);
        if let Some(params) = params {
            encode.int64(params.obj_id as i64);
        } else {
            encode.int64(0);
        }
        if let Some(transfers) = transfer {
            encode.int64(transfers.transfers.obj_id as i64);
        } else {
            encode.int64(0);
        }
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());
        ROOT.get_map(&KEY_RETURN).immutable()
    }

    // retrieve the agent id of the caller of the smart contract
    pub fn caller(&self) -> ScAgentId { ROOT.get_agent_id(&KEY_CALLER).value() }

    // shorthand to synchronously call a smart contract function on the current contract
    pub fn call_self(&self, hfunction: ScHname, params: Option<ScMutableMap>, transfers: Option<ScTransfers>) -> ScImmutableMap {
        self.call(self.contract_id().hname(), hfunction, params, transfers)
    }

    // deploys a new instance of the specified smart contract on the current chain
    // the provided parameters are passed to the smart contract "init" function
    pub fn deploy(&self, program_hash: &ScHash, name: &str, description: &str, params: Option<ScMutableMap>) {
        let mut encode = BytesEncoder::new();
        encode.hash(program_hash);
        encode.string(name);
        encode.string(description);
        if let Some(params) = params {
            encode.int64(params.obj_id as i64);
        } else {
            encode.int64(0);
        }
        ROOT.get_bytes(&KEY_DEPLOY).set_value(&encode.data());
    }

    // signals an event on the node that external entities can subscribe to
    pub fn event(&self, text: &str) {
        ROOT.get_string(&KEY_EVENT).set_value(text)
    }

    // access the incoming balances for all token colors
    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_INCOMING).immutable() }
    }

    // retrieve the color of the tokens that were minted in this transaction
    pub fn minted_color(&self) -> ScColor { ScColor::from_request_id(&self.request_id()) }

    // retrieve the amount of tokens that were minted in this transaction
    pub fn minted_supply(&self) -> i64 {
        ROOT.get_int64(&KEY_MINTED).value()
    }

    // asynchronously calls the specified smart contract function,
    // passing the provided parameters and token transfers to it
    pub fn post(&self, contract_id: &ScContractId, function: ScHname, params: Option<ScMutableMap>, transfer: Option<ScTransfers>, delay: i64) {
        let mut encode = BytesEncoder::new();
        encode.contract_id(contract_id);
        encode.hname(&function);
        if let Some(params) = &params {
            encode.int64(params.obj_id as i64);
        } else {
            encode.int64(0);
        }
        if let Some(transfer) = &transfer {
            encode.int64(transfer.transfers.obj_id as i64);
        } else {
            encode.int64(0);
        }
        encode.int64(delay);
        ROOT.get_bytes(&KEY_POST).set_value(&encode.data());
    }

    // shorthand to asynchronously call a smart contract function on the current contract
    pub fn post_self(&self, function: ScHname, params: Option<ScMutableMap>, transfer: Option<ScTransfers>, delay: i64) {
        self.post(&self.contract_id(), function, params, transfer, delay);
    }

    // retrieve the request id of this transaction
    fn request_id(&self) -> ScRequestId {
        ROOT.get_request_id(&KEY_REQUEST_ID).value()
    }

    // access to mutable state storage
    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    // transfers the specified tokens to the specified Tangle ledger address
    pub fn transfer_to_address(&self, address: &ScAddress, transfer: ScTransfers) {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let tx = transfers.get_map(transfers.length());
        tx.get_address(&KEY_ADDRESS).set_value(address);
        tx.get_int64(&KEY_BALANCES).set_value(transfer.transfers.obj_id as i64);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
pub struct ScViewContext {}

// shared part of interface
impl ScBaseContext for ScViewContext {}

impl ScViewContext {
    // synchronously calls the specified smart contract view,
    // passing the provided parameters to it
    pub fn call(&self, contract: ScHname, function: ScHname, params: Option<ScMutableMap>) -> ScImmutableMap {
        let mut encode = BytesEncoder::new();
        encode.hname(&contract);
        encode.hname(&function);
        if let Some(params) = params {
            encode.int64(params.obj_id as i64);
        } else {
            encode.int64(0);
        }
        encode.int64(0);
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());
        ROOT.get_map(&KEY_RETURN).immutable()
    }

    // shorthand to synchronously call a smart contract view on the current contract
    pub fn call_self(&self, function: ScHname, params: Option<ScMutableMap>) -> ScImmutableMap {
        self.call(self.contract_id().hname(), function, params)
    }

    // access to immutable state storage
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_STATE).immutable()
    }
}
