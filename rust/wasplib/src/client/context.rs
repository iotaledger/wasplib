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

    pub fn minted(&self) -> ScColor {
        return ScColor::from_bytes(&self.balances.get_bytes(&ScColor::MINT).value());
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScContract {
    contract: ScImmutableMap,
}

impl ScContract {
    pub fn color(&self) -> ScColor {
        self.contract.get_color("color").value()
    }

    pub fn description(&self) -> String {
        self.contract.get_string("description").value()
    }

    pub fn id(&self) -> ScAgent {
        self.contract.get_agent("id").value()
    }

    pub fn name(&self) -> String {
        self.contract.get_string("name").value()
    }

    pub fn owner(&self) -> ScAgent {
        self.contract.get_agent("owner").value()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScLog {
    log: ScMutableMapArray,
}

impl ScLog {
    pub fn append(&self, timestamp: i64, data: &[u8]) {
        let log_entry = self.log.get_map(self.log.length());
        log_entry.get_int("timestamp").set_value(timestamp);
        log_entry.get_bytes("data").set_value(data);
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
        let decode = self.utility.get_string("base58");
        let encode = self.utility.get_bytes("base58");
        decode.set_value(value);
        encode.value()
    }

    pub fn base58_encode(&self, value: &[u8]) -> String {
        //TODO atomic set/get
        let decode = self.utility.get_string("base58");
        let encode = self.utility.get_bytes("base58");
        encode.set_value(value);
        decode.value()
    }

    pub fn hash(&self, value: &[u8]) -> Vec<u8> {
        //TODO atomic set/get
        let hash = self.utility.get_bytes("hash");
        hash.set_value(value);
        hash.value()
    }

    pub fn random(&self, max: i64) -> i64 {
        let rnd = self.utility.get_int("random").value();
        (rnd as u64 % max as u64) as i64
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub trait ScBaseContext {
    fn balances(&self) -> ScBalances {
        ScBalances { balances: ROOT.get_map("balances").immutable() }
    }

    fn caller(&self) -> ScAgent { ROOT.get_agent("caller").value() }

    fn contract(&self) -> ScContract {
        ScContract { contract: ROOT.get_map("contract").immutable() }
    }

    fn error(&self) -> ScMutableString {
        ROOT.get_string("error")
    }

    fn from(&self, originator: &ScAgent) -> bool {
        self.caller() == *originator
    }

    fn log(&self, text: &str) {
        set_string(1, key_log(), text)
    }

    fn params(&self) -> ScImmutableMap {
        ROOT.get_map("params").immutable()
    }

    fn results(&self) -> ScMutableMap {
        ROOT.get_map("results")
    }

    fn timestamp(&self) -> i64 {
        ROOT.get_int("timestamp").value()
    }

    fn trace(&self, text: &str) {
        set_string(1, key_trace(), text)
    }

    fn utility(&self) -> ScUtility {
        ScUtility { utility: ROOT.get_map("utility") }
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
        ScBalances { balances: ROOT.get_map("incoming").immutable() }
    }

    pub fn post(&self, function: &str) -> ScPostInfo {
        ScPostInfo::new(function)
    }

    pub fn state(&self) -> ScMutableMap {
        ROOT.get_map("state")
    }

    pub fn timestamped_log(&self, key: &str) -> ScLog {
        ScLog { log: ROOT.get_map("logs").get_map_array(key) }
    }

    pub fn transfer(&self, agent: &ScAgent, color: &ScColor, amount: i64) {
        let transfers = ROOT.get_map_array("transfers");
        let transfer = transfers.get_map(transfers.length());
        transfer.get_agent("agent").set_value(agent);
        transfer.get_color("color").set_value(color);
        transfer.get_int("amount").set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewContext {}

impl ScBaseContext for ScViewContext {}

impl ScViewContext {
    pub fn state(&self) -> ScImmutableMap {
        ROOT.get_map("state").immutable()
    }

    pub fn timestamped_log(&self, key: &str) -> ScImmutableMapArray {
        ROOT.get_map("logs").get_map_array(key).immutable()
    }
}
