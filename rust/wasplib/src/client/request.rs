// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::hashtypes::*;
use super::immutable::*;
use super::mutable::*;

pub(crate) fn make_request(key: &str, function: &str) -> ScMutableMap {
    let root = ScMutableMap::new(1);
    let requests = root.get_map_array(key);
    let request = requests.get_map(requests.length());
    request.get_string("function").set_value(function);
    request
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScCallInfo {
    pub(crate) call: ScMutableMap,
}

impl ScCallInfo {
    pub fn call(&self) {
        self.call.get_int("delay").set_value(-1);
    }

    pub fn contract(&self, contract: &str) -> &ScCallInfo {
        self.call.get_string("contract").set_value(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.call.get_map("params")
    }

    pub fn results(&self) -> ScImmutableMap {
        self.call.get_map("results").immutable()
    }

    pub fn transfer(&self, color: &ScColor, amount: i64) {
        let transfers = self.call.get_key_map("transfers");
        transfers.get_int(&color.to_bytes()).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScPostInfo {
    pub(crate) post: ScMutableMap,
}

impl ScPostInfo {
    pub fn chain(&self, chain: &ScAddress) -> &ScPostInfo {
        self.post.get_address("chain").set_value(chain);
        self
    }

    pub fn contract(&self, contract: &str) -> &ScPostInfo {
        self.post.get_string("contract").set_value(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.post.get_map("params")
    }

    pub fn post(&self, delay: i64) {
        self.post.get_int("delay").set_value(delay);
    }

    pub fn transfer(&self, color: &ScColor, amount: i64) {
        let transfers = self.post.get_key_map("transfers");
        transfers.get_int(&color.to_bytes()).set_value(amount);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewInfo {
    pub(crate) view: ScMutableMap,
}

impl ScViewInfo {
    pub fn contract(&self, contract: &str) -> &ScViewInfo {
        self.view.get_string("contract").set_value(contract);
        self
    }

    pub fn params(&self) -> ScMutableMap {
        self.view.get_map("params")
    }

    pub fn results(&self) -> ScImmutableMap {
        self.view.get_map("results").immutable()
    }

    pub fn view(&self) {
        self.view.get_int("delay").set_value(-2);
    }
}
