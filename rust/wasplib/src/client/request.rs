// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::hashtypes::*;
use super::immutable::*;
use super::mutable::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScBaseInfo {
    request: ScMutableMap,
}

impl ScBaseInfo {
    fn new(key: &str, function: &str) -> ScBaseInfo {
        let requests = ROOT.get_map_array(key);
        let request = requests.get_map(requests.length());
        request.get_string("function").set_value(function);
        ScBaseInfo { request: request }
    }

    fn chain(&self, chain: &ScAddress) {
        self.request.get_address("chain").set_value(chain);
    }

    fn contract(&self, contract: &str) {
        self.request.get_string("contract").set_value(contract);
    }

    fn exec(&self, delay: i64) {
        self.request.get_int("delay").set_value(delay);
    }

    fn params(&self) -> ScMutableMap {
        self.request.get_map("params")
    }

    fn results(&self) -> ScImmutableMap {
        self.request.get_map("results").immutable()
    }

    fn transfer(&self, color: &ScColor, amount: i64) {
        let transfers = self.request.get_map("transfers");
        transfers.get_int(color).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScCallInfo {
    base: ScBaseInfo,
}

impl ScCallInfo {
    pub fn new(function: &str) -> ScCallInfo {
        ScCallInfo { base: ScBaseInfo::new("calls", function) }
    }

    pub fn call(&self) {
        self.base.exec(-1);
    }

    pub fn contract(&self, contract: &str) -> &ScCallInfo {
        self.base.contract(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    pub fn results(&self) -> ScImmutableMap {
        self.base.results()
    }

    pub fn transfer(&self, color: &ScColor, amount: i64) {
        self.base.transfer(color, amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScPostInfo {
    base: ScBaseInfo,
}

impl ScPostInfo {
    pub fn new(function: &str) -> ScPostInfo {
        ScPostInfo { base: ScBaseInfo::new("posts", function) }
    }

    pub fn chain(&self, chain: &ScAddress) -> &ScPostInfo {
        self.base.chain(chain);
        self
    }

    pub fn contract(&self, contract: &str) -> &ScPostInfo {
        self.base.contract(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    pub fn post(&self, delay: i64) {
        self.base.exec(delay);
    }

    pub fn transfer(&self, color: &ScColor, amount: i64) {
        self.base.transfer(color, amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewInfo {
    base: ScBaseInfo,
}

impl ScViewInfo {
    pub fn new(function: &str) -> ScViewInfo {
        ScViewInfo { base: ScBaseInfo::new("views", function) }
    }

    pub fn contract(&self, contract: &str) -> &ScViewInfo {
        self.base.contract(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    pub fn results(&self) -> ScImmutableMap {
        self.base.results()
    }

    pub fn view(&self) {
        self.base.exec(-2);
    }
}
