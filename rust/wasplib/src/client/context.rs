// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// encapsulates standard host entities into a simple interface

use super::hashtypes::*;
use super::host::set_string;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScAccount {
    account: ScImmutableMap,
}

impl ScAccount {
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.account.get_key_map("balance").get_int(color.to_bytes()).value()
    }

    pub fn colors(&self) -> ScImmutableColorArray {
        self.account.get_color_array("colors")
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

pub struct ScExports {
    exports: ScMutableStringArray,
    next: i32,
}

impl ScExports {
    pub fn new() -> ScExports {
        let root = ScMutableMap::new(1);
        ScExports { exports: root.get_string_array("exports"), next: 0 }
    }

    pub fn add(&mut self, name: &str) {
        self.next += 1;
        self.exports.get_string(self.next).set_value(name);
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

pub struct ScPostedRequest {
    request: ScMutableMap,
}

impl ScPostedRequest {
    pub fn code(&self, code: i64) {
        self.request.get_int("code").set_value(code);
    }

    pub fn contract(&self, contract: &ScAgent) {
        self.request.get_agent("contract").set_value(contract);
    }

    pub fn delay(&self, delay: i64) {
        self.request.get_int("delay").set_value(delay);
    }

    pub fn function(&self, function: &str) {
        self.request.get_string("function").set_value(function);
    }

    pub fn params(&self) -> ScMutableMap {
        self.request.get_map("params")
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScRequest {
    request: ScImmutableMap,
}

impl ScRequest {
    pub fn balance(&self, color: &ScColor) -> i64 {
        self.request.get_key_map("balance").get_int(color.to_bytes()).value()
    }

    pub fn colors(&self) -> ScImmutableColorArray {
        self.request.get_color_array("colors")
    }

    pub fn from(&self, originator: &ScAgent) -> bool {
        self.sender() == *originator
    }

    pub fn id(&self) -> ScRequestId {
        self.request.get_request_id("id").value()
    }

    pub fn minted_color(&self) -> ScColor {
        self.request.get_color("hash").value()
    }

    pub fn params(&self) -> ScImmutableMap {
        self.request.get_map("params")
    }

    pub fn sender(&self) -> ScAgent { self.request.get_agent("sender").value() }

    pub fn timestamp(&self) -> i64 {
        self.request.get_int("timestamp").value()
    }

    pub fn tx_hash(&self) -> ScTxHash {
        self.request.get_tx_hash("hash").value()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScTransfer {
    transfer: ScMutableMap,
}

impl ScTransfer {
    pub fn agent(&self, agent: &ScAgent) {
        self.transfer.get_agent("agent").set_value(agent);
    }

    pub fn amount(&self, amount: i64) {
        self.transfer.get_int("amount").set_value(amount);
    }

    pub fn color(&self, color: &ScColor) {
        self.transfer.get_color("color").set_value(color);
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

pub struct ScContext {
    root: ScMutableMap,
}

impl ScContext {
    pub fn new() -> ScContext {
        ScContext { root: ScMutableMap::new(1) }
    }

    pub fn account(&self) -> ScAccount {
        ScAccount { account: self.root.get_map("account").immutable() }
    }

    pub fn contract(&self) -> ScContract {
        ScContract { contract: self.root.get_map("contract").immutable() }
    }

    pub fn error(&self) -> ScMutableString {
        self.root.get_string("error")
    }

    pub fn log(&self, text: &str) {
        set_string(1, key_log(), text)
    }

    pub fn post_request(&self, contract: &ScAgent, function: &str, delay: i64) -> ScMutableMap {
        let posted_requests = self.root.get_map_array("postedRequests");
        let request = ScPostedRequest { request: posted_requests.get_map(posted_requests.length()) };
        request.contract(contract);
        request.function(function);
        request.delay(delay);
        request.params()
    }

    // just for compatibility with old hardcoded SCs
    pub fn post_request_with_code(&self, contract: &ScAgent, code: i64, delay: i64) -> ScMutableMap {
        let posted_requests = self.root.get_map_array("postedRequests");
        let request = ScPostedRequest { request: posted_requests.get_map(posted_requests.length()) };
        request.contract(contract);
        request.code(code);
        request.delay(delay);
        request.params()
    }

    pub fn request(&self) -> ScRequest {
        ScRequest { request: self.root.get_map("request").immutable() }
    }

    pub fn state(&self) -> ScMutableMap {
        self.root.get_map("state")
    }

    pub fn timestamped_log(&self, key: &str) -> ScLog {
        ScLog { log: self.root.get_map("logs").get_map_array(key) }
    }

    pub fn trace(&self, text: &str) {
        set_string(1, key_trace(), text)
    }

    pub fn transfer(&self, agent: &ScAgent, color: &ScColor, amount: i64) {
        let transfers = self.root.get_map_array("transfers");
        let xfer = ScTransfer { transfer: transfers.get_map(transfers.length()) };
        xfer.agent(agent);
        xfer.color(color);
        xfer.amount(amount);
    }

    pub fn utility(&self) -> ScUtility {
        ScUtility { utility: self.root.get_map("utility") }
    }
}
