// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::hashtypes::*;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;

pub struct PostRequestParams {
    pub contract: ScContractId,
    pub function: Hname,
    pub params: Option<ScMutableMap>,
    pub transfer: Option<Box<dyn Balances>>,
    pub delay: i64,
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

impl Balances for ScBalances {
    fn map_id(&self) -> i32 {
        self.balances.obj_id
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScLog {
    log: ScMutableMapArray,
}

impl ScLog {
    // appends the specified timestamp and data to the timestamped log
    pub fn append(&self, timestamp: i64, data: &[u8]) {
        let log_entry = self.log.get_map(self.log.length());
        log_entry.get_int(&KEY_TIMESTAMP).set_value(timestamp);
        log_entry.get_bytes(&KEY_DATA).set_value(data);
    }

    // number of items in the timestamped log
    pub fn length(&self) -> i32 {
        self.log.length()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScTransfers {
    transfers: ScMutableMap,
}

impl ScTransfers {
    pub fn new(color: &ScColor, amount: i64) -> ScTransfers {
        let balance = ScTransfers::new_transfers();
        balance.add(color, amount);
        balance
    }

    pub fn new_transfers() -> ScTransfers {
        ScTransfers { transfers: ScMutableMap::new() }
    }

    // appends the specified timestamp and data to the timestamped log
    pub fn add(&self, color: &ScColor, amount: i64) {
        self.transfers.get_int(color).set_value(amount);
    }
}

impl Balances for ScTransfers {
    fn map_id(&self) -> i32 {
        self.transfers.obj_id
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

    // hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
    pub fn hash(&self, value: &[u8]) -> ScHash {
        let hash = self.utility.get_bytes(&KEY_HASH);
        hash.set_value(value);
        ScHash::from_bytes(&hash.value())
    }

    pub fn hname(&self, value: &str) -> Hname {
        self.utility.get_string(&KEY_NAME).set_value(value);
        Hname::from_bytes(&self.utility.get_bytes(&KEY_HNAME).value())
    }

    // generates a random value from 0 to max (exclusive max) using a deterministic RNG
    pub fn random(&self, max: i64) -> i64 {
        let rnd = self.utility.get_int(&KEY_RANDOM).value();
        (rnd as u64 % max as u64) as i64
    }
}

// wrapper for simplified use by hashtypes
pub(crate) fn base58_encode(bytes: &[u8]) -> String {
    ScCallContext {}.utility().base58_encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// shared interface part of ScCallContext and ScViewContext
pub trait ScBaseContext {
    // access the current balances for all token colors
    fn balances(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_BALANCES).immutable() }
    }

    // retrieve the agent id of the caller of the smart contract
    fn caller(&self) -> ScAgent { ROOT.get_agent(&KEY_CALLER).value() }

    // retrieve the agent id of the owner of the chain this contract lives on
    fn chain_owner(&self) -> ScAgent {
        ROOT.get_agent(&KEY_CHAIN_OWNER).value()
    }

    // retrieve the agent id of the creator of this contract
    fn contract_creator(&self) -> ScAgent {
        ROOT.get_agent(&KEY_CREATOR).value()
    }

    // retrieve the id of this contract
    fn contract_id(&self) -> ScContractId {
        ROOT.get_contract_id(&KEY_ID).value()
    }

    // quick check to see if the caller of the smart contract was the specified originator agent
    fn from(&self, originator: &ScAgent) -> bool {
        self.caller().equals(originator)
    }

    // logs informational text message
    fn log(&self, text: &str) {
        ROOT.get_string(&KEY_LOG).set_value(text)
    }

    // logs error text message and then panics
    fn panic(&self, text: &str) {
        ROOT.get_string(&KEY_PANIC).set_value(text)
    }

    // retrieve parameters passed to the smart contract function that was called
    fn params(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_PARAMS).immutable()
    }

    // panics if condition is not satisfied
    fn require(&self, cond: bool, msg: &str) {
        if !cond {
            self.panic(msg)
        }
    }

    // any results returned by the smart contract function call are returned here
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
pub struct ScCallContext {}

impl ScBaseContext for ScCallContext {}

impl ScCallContext {
    // calls a smart contract function
    pub fn call(&self, contract: Hname, function: Hname, params: Option<ScMutableMap>, transfer: Option<Box<dyn Balances>>) -> ScImmutableMap {
        let calls = ROOT.get_map_array(&KEY_CALLS);
        let call = calls.get_map(calls.length());
        call.get_hname(&KEY_CONTRACT).set_value(contract);
        call.get_hname(&KEY_FUNCTION).set_value(function);
        if let Some(params) = params {
            call.get_int(&KEY_PARAMS).set_value(params.obj_id as i64);
        }
        if let Some(transfer) = transfer {
            call.get_int(&KEY_TRANSFERS).set_value(transfer.map_id() as i64);
        }
        call.get_int(&KEY_DELAY).set_value(-1);
        call.get_map(&KEY_RESULTS).immutable()
    }

    // deploys a smart contract
    pub fn deploy(&self, program_hash: &ScHash, name: &str, description: &str, params: Option<ScMutableMap>) {
        let deploys = ROOT.get_map_array(&KEY_DEPLOYS);
        let deploy = deploys.get_map(deploys.length());
        deploy.get_string(&KEY_NAME).set_value(name);
        deploy.get_string(&KEY_DESCRIPTION).set_value(description);
        if let Some(params) = params {
            deploy.get_int(&KEY_PARAMS).set_value(params.obj_id as i64);
        }
        deploy.get_hash(&KEY_HASH).set_value(program_hash);
    }

    // signals an event on the node that external entities can subscribe to
    pub fn event(&self, text: &str) {
        ROOT.get_string(&KEY_EVENT).set_value(text)
    }

    // access the incoming balances for all token colors
    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_INCOMING).immutable() }
    }

    // (delayed) posts a smart contract function
    pub fn post(&self, par: &PostRequestParams) {
        if par.delay < 0 { self.panic("Invalid delay") }
        let calls = ROOT.get_map_array(&KEY_CALLS);
        let call = calls.get_map(calls.length());
        call.get_chain_id(&KEY_CHAIN).set_value(&par.contract.chain_id());
        call.get_hname(&KEY_CONTRACT).set_value(par.contract.hname());
        call.get_hname(&KEY_FUNCTION).set_value(par.function);
        if let Some(params) = &par.params {
            call.get_int(&KEY_PARAMS).set_value(params.obj_id as i64);
        }
        if let Some(transfer) = &par.transfer {
            call.get_int(&KEY_TRANSFERS).set_value(transfer.map_id() as i64);
        }
        call.get_int(&KEY_DELAY).set_value(par.delay);
    }

    // access to mutable state storage
    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    // access to mutable named timestamped log
    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScLog {
        ScLog { log: ROOT.get_map(&KEY_LOGS).get_map_array(key) }
    }

    // transfer colored token amounts to the specified Tangle ledger address
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

impl ScBaseContext for ScViewContext {}

impl ScViewContext {
    // calls a smart contract function
    pub fn call(&self, contract: Hname, function: Hname, params: Option<ScMutableMap>) -> ScImmutableMap {
        let calls = ROOT.get_map_array(&KEY_CALLS);
        let call = calls.get_map(calls.length());
        call.get_hname(&KEY_CONTRACT).set_value(contract);
        call.get_hname(&KEY_FUNCTION).set_value(function);
        if let Some(params) = params {
            call.get_int(&KEY_PARAMS).set_value(params.obj_id as i64);
        }
        call.get_int(&KEY_DELAY).set_value(-1);
        call.get_map(&KEY_RESULTS).immutable()
    }

    // access to immutable state storage
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_STATE).immutable()
    }

    // access to immutable named timestamped log
    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        ROOT.get_map(&KEY_LOGS).get_map_array(key).immutable()
    }
}
