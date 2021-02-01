// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const SC_NAME: &str = "dividend";
const SC_HNAME: Hname = Hname(0xcce2e239);

const PARAM_ADDRESS: &str = "address";
const PARAM_FACTOR: &str = "factor";

const VAR_MEMBERS: &str = "members";
const VAR_TOTAL_FACTOR: &str = "total_factor";

const FUNC_DIVIDE: &str = "divide";
const FUNC_MEMBER: &str = "member";

const HFUNC_DIVIDE: Hname = Hname(0xc7878107);
const HFUNC_MEMBER: Hname = Hname(0xc07da2cb);
