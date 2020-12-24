// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::hashtypes::*;
use super::host::set_string;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;
use super::request::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScBalances {
    balances: ScImmutableMap,
}

impl ScBalances {
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.balances.get_int(color).value()
    }

    pub fn colors(&self) -> ScImmutableColorArray {
        self.balances.get_color_array(&KEY_COLOR)
    }

    pub fn minted(&self) -> ScColor {
        ScColor::from_bytes(&self.balances.get_bytes(&ScColor::MINT).value())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScContract {
    contract: ScImmutableMap,
}

impl ScContract {
    pub fn chain(&self) -> ScAddress {
        self.contract.get_address(&KEY_CHAIN).value()
    }

    pub fn description(&self) -> String {
        self.contract.get_string(&KEY_DESCRIPTION).value()
    }

    pub fn id(&self) -> ScAgent {
        self.contract.get_agent(&KEY_ID).value()
    }

    pub fn name(&self) -> String {
        self.contract.get_string(&KEY_NAME).value()
    }

    pub fn owner(&self) -> ScAgent {
        self.contract.get_agent(&KEY_OWNER).value()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScLog {
    log: ScMutableMapArray,
}

impl ScLog {
    pub fn append(&self, timestamp: i64, data: &[u8]) {
        let log_entry = self.log.get_map(self.log.length());
        log_entry.get_int(&KEY_TIMESTAMP).set_value(timestamp);
        log_entry.get_bytes(&KEY_DATA).set_value(data);
    }

    pub fn length(&self) -> i32 {
        self.log.length()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScUtility {
    utility: ScMutableMap,
}

impl ScUtility {
    pub fn base58_decode(&self, value: &str) -> Vec<u8> {
        //TODO atomic set/get
        let decode = self.utility.get_string(&KEY_BASE58);
        let encode = self.utility.get_bytes(&KEY_BASE58);
        decode.set_value(value);
        encode.value()
    }

    pub fn base58_encode(&self, value: &[u8]) -> String {
        //TODO atomic set/get
        let decode = self.utility.get_string(&KEY_BASE58);
        let encode = self.utility.get_bytes(&KEY_BASE58);
        encode.set_value(value);
        decode.value()
    }

    pub fn hash(&self, value: &[u8]) -> Vec<u8> {
        //TODO atomic set/get
        let hash = self.utility.get_bytes(&KEY_HASH);
        hash.set_value(value);
        hash.value()
    }

    pub fn random(&self, max: i64) -> i64 {
        let rnd = self.utility.get_int(&KEY_RANDOM).value();
        (rnd as u64 % max as u64) as i64
    }
}

pub(crate) fn base58_encode(bytes: &[u8]) -> String {
    ScCallContext {}.utility().base58_encode(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub trait ScBaseContext {
    fn balances(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_BALANCES).immutable() }
    }

    fn caller(&self) -> ScAgent { ROOT.get_agent(&KEY_CALLER).value() }

    fn contract(&self) -> ScContract {
        ScContract { contract: ROOT.get_map(&KEY_CONTRACT).immutable() }
    }

    fn error(&self) -> ScMutableString {
        ROOT.get_string(&KEY_ERROR)
    }

    fn from(&self, originator: &ScAgent) -> bool {
        self.caller() == *originator
    }

    fn log(&self, text: &str) {
        set_string(1, KEY_LOG, text)
    }

    fn panic(&self, text: &str) {
        set_string(1, KEY_PANIC, text)
    }

    fn params(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_PARAMS).immutable()
    }

    fn results(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_RESULTS)
    }

    fn timestamp(&self) -> i64 {
        ROOT.get_int(&KEY_TIMESTAMP).value()
    }

    fn trace(&self, text: &str) {
        set_string(1, KEY_TRACE, text)
    }

    fn utility(&self) -> ScUtility {
        ScUtility { utility: ROOT.get_map(&KEY_UTILITY) }
    }

    fn view(&self, function: &str) -> ScViewInfo {
        ScViewInfo::new(function)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScCallContext {}

impl ScBaseContext for ScCallContext {}

impl ScCallContext {
    pub fn call(&self, function: &str) -> ScCallInfo {
        ScCallInfo::new(function)
    }

    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map(&KEY_INCOMING).immutable() }
    }

    pub fn post(&self, function: &str) -> ScPostInfo {
        ScPostInfo::new(function)
    }

    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map(&KEY_STATE)
    }

    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScLog {
        ScLog { log: ROOT.get_map(&KEY_LOGS).get_map_array(key) }
    }

    pub fn transfer(&self, agent: &ScAgent, color: &ScColor, amount: i64) {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let transfer = transfers.get_map(transfers.length());
        transfer.get_agent(&KEY_AGENT).set_value(agent);
        transfer.get_int(color).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewContext {}

impl ScBaseContext for ScViewContext {}

impl ScViewContext {
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map(&KEY_STATE).immutable()
    }

    pub fn timestamped_log<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        ROOT.get_map(&KEY_LOGS).get_map_array(key).immutable()
    }
}
