// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::hashtypes::*;
use super::immutable::*;
use super::keys::*;
use super::mutable::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScCallBuilder {
    base: ScRequestBuilder,
}

impl ScCallBuilder {
    pub fn new(function: &str) -> ScCallBuilder {
        ScCallBuilder { base: ScRequestBuilder::new(&KEY_CALLS, function) }
    }

    // execute the call request when finished building
    pub fn call(&self) -> &ScCallBuilder {
        self.base.exec(-1);
        self
    }

    // specify a different contract the call request is for
    pub fn contract(&self, contract: &str) -> &ScCallBuilder {
        self.base.contract(contract);
        self
    }

    // provide parameters for the call request
    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    // access call request results after executing the call
    pub fn results(&self) -> ScImmutableMap {
        self.base.results()
    }

    // transfer tokens as part of the call request
    pub fn transfer(&self, color: &ScColor, amount: i64) -> &ScCallBuilder {
        self.base.transfer(color, amount);
        self
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScDeployBuilder {
    deploy: ScMutableMap,
}

impl ScDeployBuilder {
    // start deployment of a smart contract with the specified name and description
    pub fn new(name: &str, description: &str) -> ScDeployBuilder {
        let deploys = ROOT.get_map_array(&KEY_DEPLOYS);
        let deploy = deploys.get_map(deploys.length());
        deploy.get_string(&KEY_NAME).set_value(name);
        deploy.get_string(&KEY_DESCRIPTION).set_value(description);
        ScDeployBuilder { deploy: deploy }
    }

    // execute the deployment of the smart contract with the specified program hash
    pub fn deploy(&self, program_hash: &ScHash) {
        self.deploy.get_hash(&KEY_HASH).set_value(program_hash);
    }

    // provide parameters for the deployment's 'init' call
    pub fn params(&self) -> ScMutableMap {
        self.deploy.get_map(&KEY_PARAMS)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScPostBuilder {
    base: ScRequestBuilder,
}

impl ScPostBuilder {
    pub fn new(function: &str) -> ScPostBuilder {
        ScPostBuilder { base: ScRequestBuilder::new(&KEY_POSTS, function) }
    }

    // specify another chain to post the request to
    pub fn chain(&self, chain: &ScAddress) -> &ScPostBuilder {
        self.base.request.get_address(&KEY_CHAIN).set_value(chain);
        self
    }

    // specify a different contract the post request is for
    pub fn contract(&self, contract: &str) -> &ScPostBuilder {
        self.base.contract(contract);
        self
    }

    // provide parameters for the post request
    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    // execute the post request when finished building with the specified delay
    pub fn post(&self, delay: i64) {
        self.base.exec(delay);
    }

    // transfer tokens as part of the post request
    pub fn transfer(&self, color: &ScColor, amount: i64) -> &ScPostBuilder {
        self.base.transfer(color, amount);
        self
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScTransferBuilder {
    transfer: ScMutableMap,
}

impl ScTransferBuilder {
    // start a transfer to the specified local chain agent account
    pub fn new_transfer(agent: &ScAgent) -> ScTransferBuilder {
        let local_chain = ROOT.get_address(&KEY_CHAIN).value();
        ScTransferBuilder::new_transfer_cross_chain(&local_chain, agent)
    }

    // start a transfer to a Tangle ledger address
    pub fn new_transfer_to_address(address: &ScAddress) -> ScTransferBuilder {
        ScTransferBuilder::new_transfer_cross_chain(&ScAddress::NULL, &address.as_agent())
    }

    // start a transfer to the specified cross chain agent account
    pub fn new_transfer_cross_chain(chain: &ScAddress, agent: &ScAgent) -> ScTransferBuilder {
        let transfers = ROOT.get_map_array(&KEY_TRANSFERS);
        let transfer = transfers.get_map(transfers.length());
        transfer.get_agent(&KEY_AGENT).set_value(agent);
        if *chain != ScAddress::NULL {
            transfer.get_address(&KEY_CHAIN).set_value(chain);
        }
        ScTransferBuilder { transfer: transfer }
    }

    // sends the complete built transfer to the node
    pub fn send(&self) {
        self.transfer.get_int(&ScColor::MINT).set_value(-1);
    }

    // transfer the specified amount of tokens of the specified color as part of this transfer
    // concatenate one of these for each separate color/amount combination
    // amount is supposed to be > 0, unique colors can appear only once
    pub fn transfer(&self, color: &ScColor, amount: i64) -> &ScTransferBuilder {
        self.transfer.get_int(color).set_value(amount);
        self
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScViewBuilder {
    base: ScRequestBuilder,
}

impl ScViewBuilder {
    pub fn new(function: &str) -> ScViewBuilder {
        ScViewBuilder { base: ScRequestBuilder::new(&KEY_VIEWS, function) }
    }

    // specify a different contract the view request is for
    pub fn contract(&self, contract: &str) -> &ScViewBuilder {
        self.base.contract(contract);
        self
    }

    // provide parameters for the view request
    pub fn params(&self) -> ScMutableMap {
        self.base.params()
    }

    // access view request results after executing the call
    pub fn results(&self) -> ScImmutableMap {
        self.base.results()
    }

    // execute the view request when finished building
    pub fn view(&self) -> &ScViewBuilder {
        self.base.exec(-2);
        self
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScRequestBuilder {
    request: ScMutableMap,
}

impl ScRequestBuilder {
    // start building a request for the specified smart contract function
    fn new<T: MapKey + ?Sized>(key: &T, function: &str) -> ScRequestBuilder {
        let requests = ROOT.get_map_array(key);
        let request = requests.get_map(requests.length());
        request.get_string(&KEY_FUNCTION).set_value(function);
        ScRequestBuilder { request: request }
    }

    // specify a different contract the request is for
    fn contract(&self, contract: &str) {
        self.request.get_string(&KEY_CONTRACT).set_value(contract);
    }

    // execute the request
    fn exec(&self, delay: i64) {
        self.request.get_int(&KEY_DELAY).set_value(delay);
    }

    // provide parameters for the request
    fn params(&self) -> ScMutableMap {
        self.request.get_map(&KEY_PARAMS)
    }

    // access request results after executing the request
    fn results(&self) -> ScImmutableMap {
        self.request.get_map(&KEY_RESULTS).immutable()
    }

    // transfer tokens as part of the request
    fn transfer(&self, color: &ScColor, amount: i64) {
        let transfers = self.request.get_map(&KEY_TRANSFERS);
        transfers.get_int(color).set_value(amount);
    }
}
