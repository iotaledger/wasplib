// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::hashtypes::*;
use super::keys::*;
use super::mutable::*;

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
        if !chain.equals(&ScAddress::NULL) {
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
