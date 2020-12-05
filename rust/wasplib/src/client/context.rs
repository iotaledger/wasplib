// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::hashtypes::*;
use super::host::set_string;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;
use super::request::*;

pub(crate) static ROOT_CALL_CONTEXT: ScCallContext = ScCallContext { root: ScMutableMap::new(1) };
pub(crate) static ROOT_VIEW_CONTEXT: ScViewContext = ScViewContext { root: ScMutableMap::new(1) };

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScBalances {
    balances: ScImmutableKeyMap,
}

impl ScBalances {
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.balances.get_int(color.to_bytes()).value()
    }

    pub fn minted(&self) -> ScColor {
        let mint_key = &ScColor::MINT.to_bytes();
        return ScColor::from_bytes(&self.balances.get_bytes(mint_key).value());
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

pub struct ScCallContext {
    root: ScMutableMap,
}

impl ScCallContext {
    pub fn balances(&self) -> ScBalances {
        ScBalances { balances: self.root.get_key_map("balances").immutable() }
    }

    pub fn call(&self, function: &str) -> ScCallInfo {
        ScCallInfo { call: make_request("calls", function) }
    }

    pub fn caller(&self) -> ScAgent { self.root.get_agent("caller").value() }

    pub fn contract(&self) -> ScContract {
        ScContract { contract: self.root.get_map("contract").immutable() }
    }

    pub fn error(&self) -> ScMutableString {
        self.root.get_string("error")
    }

    pub fn from(&self, originator: &ScAgent) -> bool {
        self.caller() == *originator
    }
    pub fn incoming(&self) -> ScBalances {
        ScBalances { balances: self.root.get_key_map("incoming").immutable() }
    }

    pub fn log(&self, text: &str) {
        set_string(1, key_log(), text)
    }

    pub fn params(&self) -> ScImmutableMap {
        self.root.get_map("params").immutable()
    }

    pub fn post(&self, function: &str) -> ScPostInfo {
        ScPostInfo { post: make_request("posts", function) }
    }

    pub fn results(&self) -> ScMutableMap {
        self.root.get_map("results")
    }

    pub fn state(&self) -> ScMutableMap {
        self.root.get_map("state")
    }

    pub fn timestamp(&self) -> i64 {
        self.root.get_int("timestamp").value()
    }

    pub fn timestamped_log(&self, key: &str) -> ScLog {
        ScLog { log: self.root.get_map("logs").get_map_array(key) }
    }

    pub fn trace(&self, text: &str) {
        set_string(1, key_trace(), text)
    }

    pub fn transfer(&self, agent: &ScAgent, color: &ScColor, amount: i64) {
        let transfers = self.root.get_map_array("transfers");
        let transfer = transfers.get_map(transfers.length());
        transfer.get_agent("agent").set_value(agent);
        transfer.get_color("color").set_value(color);
        transfer.get_int("amount").set_value(amount);
    }

    pub fn utility(&self) -> ScUtility {
        ScUtility { utility: self.root.get_map("utility") }
    }

    pub fn view(&self, function: &str) -> ScViewInfo {
        ScViewInfo { view: make_request("views", function) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewContext {
    root: ScMutableMap,
}

impl ScViewContext {
    pub fn balances(&self) -> ScBalances {
        ScBalances { balances: self.root.get_key_map("balances").immutable() }
    }

    pub fn caller(&self) -> ScAgent { self.root.get_agent("caller").value() }

    pub fn contract(&self) -> ScContract {
        ScContract { contract: self.root.get_map("contract").immutable() }
    }

    pub fn error(&self) -> ScMutableString {
        self.root.get_string("error")
    }

    pub fn from(&self, originator: &ScAgent) -> bool {
        self.caller() == *originator
    }

    pub fn log(&self, text: &str) {
        set_string(1, key_log(), text)
    }

    pub fn params(&self) -> ScImmutableMap {
        self.root.get_map("params").immutable()
    }

    pub fn results(&self) -> ScMutableMap {
        self.root.get_map("results")
    }

    pub fn state(&self) -> ScImmutableMap {
        self.root.get_map("state").immutable()
    }

    pub fn timestamp(&self) -> i64 {
        self.root.get_int("timestamp").value()
    }

    pub fn timestamped_log(&self, key: &str) -> ScImmutableMapArray {
        self.root.get_map("logs").get_map_array(key).immutable()
    }

    pub fn trace(&self, text: &str) {
        set_string(1, key_trace(), text)
    }

    pub fn utility(&self) -> ScUtility {
        ScUtility { utility: self.root.get_map("utility") }
    }

    pub fn view(&self, function: &str) -> ScViewInfo {
        ScViewInfo { view: make_request("views", function) }
    }
}
