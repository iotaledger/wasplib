// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::builders::*;
use super::hashtypes::*;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// used to retrieve any information related to the current smart contract
pub struct ScContract {
    contract: ScImmutableMap,
}

impl ScContract {
    // retrieve the chain id of the chain this contract lives on
    pub fn chain(&self) -> ScAddress {
        self.contract.get_address(&KEY_CHAIN).value()
    }

    // retrieve the agent id of the owner of the chain this contract lives on
    pub fn chain_owner(&self) -> ScAgent {
        self.contract.get_agent(&KEY_CHAIN_OWNER).value()
    }

    // retrieve the agent id of the creator of this contract
    pub fn creator(&self) -> ScAgent {
        self.contract.get_agent(&KEY_CREATOR).value()
    }

    // retrieve this contract's description
    pub fn description(&self) -> String {
        self.contract.get_string(&KEY_DESCRIPTION).value()
    }

    // retrieve the id of this contract
    pub fn id(&self) -> ScAgent {
        self.contract.get_agent(&KEY_ID).value()
    }

    // retrieve this contract's name
    pub fn name(&self) -> String {
        self.contract.get_string(&KEY_NAME).value()
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

pub struct ScUtility {
    utility: ScMutableMap,
}

impl ScUtility {
    // decodes the specified base58-encoded string value to its original bytes
    pub fn base58_decode(&self, value: &str) -> Vec<u8> {
        //TODO atomic set/get
        let decode = self.utility.get_string(&KEY_BASE58);
        let encode = self.utility.get_bytes(&KEY_BASE58);
        decode.set_value(value);
        encode.value()
    }

    // encodes the specified bytes to a base-58-encoded string
    pub fn base58_encode(&self, value: &[u8]) -> String {
        //TODO atomic set/get
        let decode = self.utility.get_string(&KEY_BASE58);
        let encode = self.utility.get_bytes(&KEY_BASE58);
        encode.set_value(value);
        decode.value()
    }

    // hashes the specified value bytes using blake2b hashing and returns the resulting 32-byte hash
    pub fn hash(&self, value: &[u8]) -> ScHash {
        //TODO atomic set/get
        let hash = self.utility.get_bytes(&KEY_HASH);
        hash.set_value(value);
        ScHash::from_bytes(&hash.value())
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

    // groups contract-related information under one access space
    fn contract(&self) -> ScContract {
        ScContract { contract: ROOT.get_map(&KEY_CONTRACT).immutable() }
    }

    // quick check to see if the caller of the smart contract was the specified originator agent
    fn from(&self, originator: &ScAgent) -> bool {
        self.caller() == *originator
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

    // starts a call to a smart contract view function.
    fn view(&self, function: &str) -> ScViewBuilder {
        ScViewBuilder::new(function)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with mutable access to state
pub struct ScCallContext {}

impl ScBaseContext for ScCallContext {}

impl ScCallContext {
    // starts a call to a smart contract function
    pub fn call(&self, function: &str) -> ScCallBuilder {
        ScCallBuilder::new(function)
    }

    // starts deployment of a smart contract
    pub fn deploy(&self, name: &str, description: &str) -> ScDeployBuilder {
        ScDeployBuilder::new(name, description)
    }

    // access the incoming balances for all token colors
    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_INCOMING).immutable() }
    }

    // starts a (delayed) post to a smart contract function.
    pub fn post(&self, function: &str) -> ScPostBuilder {
        ScPostBuilder::new(function)
    }

    // signals an event on the chain that entities can register for
    fn signal_event(&self, text: &str) {
        ROOT.get_string(&KEY_EVENT).set_value(text)
    }

    // access to mutable state storage
    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    // access to mutable named timestamped log
    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScLog {
        ScLog { log: ROOT.get_map(&KEY_LOGS).get_map_array(key) }
    }

    // transfer the specified amount of tokens of the specified color to the specified agent
    pub fn transfer(&self, agent: &ScAgent, color: &ScColor, amount: i64) {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let transfer = transfers.get_map(transfers.length());
        transfer.get_agent(&KEY_AGENT).set_value(agent);
        transfer.get_int(color).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// smart contract interface with immutable access to state
pub struct ScViewContext {}

impl ScBaseContext for ScViewContext {}

impl ScViewContext {
    // access to immutable state storage
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_STATE).immutable()
    }

    // access to immutable named timestamped log
    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        ROOT.get_map(&KEY_LOGS).get_map_array(key).immutable()
    }
}
