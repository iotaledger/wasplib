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

// parameter structure required for ctx.post()
pub struct PostRequestParams {
    //@formatter:off
    pub contract_id: ScContractId,              // full contract id (chain id + contract Hname)
    pub function:    ScHname,                   // Hname of the contract func or view to call
    pub params:      Option<ScMutableMap>,      // an optional map of parameters to pass to the function
    pub transfer:    Option<Box<dyn Balances>>, // optional balances to transfer as part of the call
    pub delay:       i64,                       // delay in seconds before the function will be run
    //@formatter:on
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// defines which map objects can be passed as a map of transfers to a function call or post
pub trait Balances {
    fn map_id(&self) -> i32;
}

// used to retrieve any information that is related to colored token balances
pub struct ScBalances {
    balances: ScImmutableMap,
}

impl ScBalances {
    // retrieve the balance for the specified token color
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.balances.get_int(color).value()
    }

    // retrieve a list of all token colors that have a non-zero balance
    pub fn colors(&self) -> ScImmutableColorArray {
        self.balances.get_color_array(&KEY_COLOR)
    }

    // retrieve the color of newly minted tokens
    pub fn minted(&self) -> ScColor {
        ScColor::from_bytes(&self.balances.get_bytes(&ScColor::MINT).value())
    }
}

// ScBalances can be used to transfer tokens to a function call
impl Balances for ScBalances {
    fn map_id(&self) -> i32 {
        self.balances.obj_id
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
        let balance = ScTransfers::new_transfers();
        balance.add(color, amount);
        balance
    }

    // create a new transfer object ready to add token transfers
    pub fn new_transfers() -> ScTransfers {
        ScTransfers { transfers: ScMutableMap::new() }
    }

    // add the specified token transfer to the transfer object
    pub fn add(&self, color: &ScColor, amount: i64) {
        self.transfers.get_int(color).set_value(amount);
    }
}

// ScTransfers can be used to transfer tokens to a function call
impl Balances for ScTransfers {
    fn map_id(&self) -> i32 {
        self.transfers.obj_id
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// provide access to utility functions that are handled by the host
pub struct ScUtility {
    utility: ScMutableMap,
}

impl ScUtility {
    // aggregates the specified multiple BLS signatures and public keys into a single one
    pub fn aggregate_bls_signatures(&self, pub_keys_bin: &[&[u8]], sigs_bin: &[&[u8]]) -> (Vec<u8>, Vec<u8>) {
        let mut encode = BytesEncoder::new();
        encode.int(pub_keys_bin.len() as i64);
        for pub_key in pub_keys_bin {
            encode.bytes(pub_key);
        }
        encode.int(sigs_bin.len() as i64);
        for sig in sigs_bin {
            encode.bytes(sig);
        }
        let aggregator = self.utility.get_bytes(&KEY_AGGREGATE_BLS);
        aggregator.set_value(&encode.data());
        let aggregated = aggregator.value();
        let mut decode = BytesDecoder::new(&aggregated);
        return (decode.bytes().to_vec(), decode.bytes().to_vec());
    }

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
        let rnd = self.utility.get_int(&KEY_RANDOM).value();
        (rnd as u64 % max as u64) as i64
    }

    // checks if the specified BLS signature is valid
    pub fn valid_bls_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        self.utility.get_bytes(&KEY_VALID_BLS).set_value(&encode.data());
        self.utility.get_int(&KEY_VALID).value() != 0
    }

    // checks if the specified ED25519 signature is valid
    pub fn valid_ed25519_signature(&self, data: &[u8], pub_key: &[u8], signature: &[u8]) -> bool {
        let mut encode = BytesEncoder::new();
        encode.bytes(data);
        encode.bytes(pub_key);
        encode.bytes(signature);
        self.utility.get_bytes(&KEY_VALID_ED25519).set_value(&encode.data());
        self.utility.get_int(&KEY_VALID).value() != 0
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
        ROOT.get_int(&KEY_TIMESTAMP).value()
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
    pub fn call(&self, hcontract: ScHname, hfunction: ScHname, params: Option<ScMutableMap>, transfer: Option<Box<dyn Balances>>) -> ScImmutableMap {
        let mut encode = BytesEncoder::new();
        encode.hname(&hcontract);
        encode.hname(&hfunction);
        if let Some(params) = params {
            encode.int(params.obj_id as i64);
        } else {
            encode.int(0);
        }
        if let Some(transfer) = transfer {
            encode.int(transfer.map_id() as i64);
        } else {
            encode.int(0);
        }
        ROOT.get_bytes(&KEY_CALL).set_value(&encode.data());
        ROOT.get_map(&KEY_RETURN).immutable()
    }

    // retrieve the agent id of the caller of the smart contract
    pub fn caller(&self) -> ScAgentId { ROOT.get_agent_id(&KEY_CALLER).value() }

    // shorthand to synchronously call a smart contract function on the current contract
    pub fn call_self(&self, hfunction: ScHname, params: Option<ScMutableMap>, transfer: Option<Box<dyn Balances>>) -> ScImmutableMap {
        self.call(self.contract_id().hname(), hfunction, params, transfer)
    }

    // deploys a new instance of the specified smart contract on the current chain
    // the provided parameters are passed to the smart contract "init" function
    pub fn deploy(&self, program_hash: &ScHash, name: &str, description: &str, params: Option<ScMutableMap>) {
        let mut encode = BytesEncoder::new();
        encode.hash(program_hash);
        encode.string(name);
        encode.string(description);
        if let Some(params) = params {
            encode.int(params.obj_id as i64);
        } else {
            encode.int(0);
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

    // posts a request to asynchronously invoke the specified smart
    // contract function according to the specified request parameters
    pub fn post(&self, par: &PostRequestParams) {
        let mut encode = BytesEncoder::new();
        encode.contract_id(&par.contract_id);
        encode.hname(&par.function);
        if let Some(params) = &par.params {
            encode.int(params.obj_id as i64);
        } else {
            encode.int(0);
        }
        if let Some(transfer) = &par.transfer {
            encode.int(transfer.map_id() as i64);
        } else {
            encode.int(0);
        }
        encode.int(par.delay);
        ROOT.get_bytes(&KEY_POST).set_value(&encode.data());
    }

    // access to mutable state storage
    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    // transfers the specified tokens to the specified Tangle ledger address
    pub fn transfer_to_address<T: Balances + ?Sized>(&self, address: &ScAddress, transfer: &T) {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let tx = transfers.get_map(transfers.length());
        tx.get_address(&KEY_ADDRESS).set_value(address);
        tx.get_int(&KEY_BALANCES).set_value(transfer.map_id() as i64);
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
            encode.int(params.obj_id as i64);
        } else {
            encode.int(0);
        }
        encode.int(0);
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
