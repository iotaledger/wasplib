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
