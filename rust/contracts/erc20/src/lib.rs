#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

const VAR_SUPPLY: &str = "supply";
const VAR_BALANCES: &str = "b";
const VAR_APPROVALS: &str = "a";
const VAR_SOURCE_ADDRESS: &str = "s";
const VAR_TARGET_ADDRESS: &str = "addr";
const VAR_AMOUNT: &str = "amount";

struct Delegation {
    delegated_to: ScAgent,
    amount: i64,
}

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("initSC");
    exports.add("transfer");
    exports.add("approve");
    exports.add("transfer_from");
}

#[no_mangle]
pub fn initSC() {
    let sc = ScContext::new();
    sc.log("initSC");

    let state = sc.state();
    let supplyState = state.get_int(VAR_SUPPLY);
    if supplyState.value() > 0 {
        // already initialized
        sc.log("initSC.fail: already initialized");
        return;
    }
    let params = sc.request().params();
    let supplyParam = params.get_int(VAR_SUPPLY);
    if supplyParam.value() == 0 {
        sc.log("initSC.fail: wrong 'supply' parameter");
        return;
    }
    let supply = supplyParam.value();
    supplyState.set_value(supply);
    let owner = sc.contract().owner();
    state.get_key_map(VAR_BALANCES).get_int(owner.to_bytes()).set_value(supply);

    sc.log(&("initSC.success. Supply = ".to_string() + &supply.to_string()));
    sc.log(&("initSC.success. Owner = ".to_string() + &owner.to_string()));
}

#[no_mangle]
pub fn transfer() {
    let sc = ScContext::new();
    sc.log("transfer");

    let request = sc.request();
    let params = request.params();
    let sender = request.sender();
    sc.log(&("sender: ".to_string() + &sender.to_string()));

    // TODO validate parameter address
    let target_addr = params.get_agent(VAR_TARGET_ADDRESS).value();
    let amount = params.get_int(VAR_AMOUNT).value();
    if amount <= 0 {
        sc.log("transfer.fail: wrong 'amount' parameter");
        return;
    }
    let succ = transfer_internal(&sender, &target_addr, amount);
    sc.log(if succ { "transfer.success" } else { "transfer.fail" });
}

#[no_mangle]
pub fn approve() {
    let sc = ScContext::new();
    sc.log("approve");

    let state = sc.state();
    let request = sc.request();
    let sender = request.sender();
    let delegations_data = state.get_key_map(VAR_APPROVALS).get_bytes(sender.to_bytes());
    let mut delegations = decode_delegations(&delegations_data.value());

    let params = request.params();
    let amount = params.get_int(VAR_AMOUNT);
    if amount.value() == 0 {
        sc.log("approve.fail: wrong 'amount' parameter");
        return;
    }
    let target_addr = params.get_agent(VAR_TARGET_ADDRESS);

    add_delegation(&mut delegations, target_addr.value(), amount.value());

    delegations_data.set_value(encode_delegations(&delegations).as_slice());

    sc.log("approve.success");
}

#[no_mangle]
pub fn transfer_from() {
    let sc = ScContext::new();
    sc.log("transfer_from");

    let state = sc.state();
    let request = sc.request();
    let sender = request.sender();

    // take parameters
    let params = request.params();
    let amount = params.get_int(VAR_AMOUNT);
    if amount.value() == 0 {
        sc.log("transfer_from.fail: wrong 'amount' parameter");
        return;
    }
    // TODO parameter validation
    let source_addr = params.get_agent(VAR_SOURCE_ADDRESS);
    let target_addr = params.get_agent(VAR_TARGET_ADDRESS);

    let delegations_data = state.get_key_map(VAR_APPROVALS).get_bytes(source_addr.value().to_bytes());
    let mut delegations = decode_delegations(&delegations_data.value());

    if !sub_delegation(&mut delegations, sender, amount.value()) {
        sc.log("transfer_from.fail: wrong delegation, possibly over the limit");
        return;
    }
    if !transfer_internal(&source_addr.value(), &target_addr.value(), amount.value()) {
        sc.log("transfer_from.fail: possibly not enough balance in the source address");
    }

    delegations_data.set_value(encode_delegations(&delegations).as_slice());
    sc.log("transfer_from.success");
}

fn transfer_internal(source_addr: &ScAgent, target_addr: &ScAgent, amount: i64) -> bool {
    let sc = ScContext::new();
    let balances = sc.state().get_key_map(VAR_BALANCES);
    let source_balance = balances.get_int(source_addr.to_bytes());
    sc.log(&("transfer_internal: source addr: = ".to_string() + &source_addr.to_string()));
    sc.log(&("transfer_internal: source balance: = ".to_string() + &source_balance.value().to_string()));

    let target_balance = balances.get_int(target_addr.to_bytes());
    sc.log(&("transfer_internal: target addr: = ".to_string() + &target_addr.to_string()));
    sc.log(&("transfer_internal: target balance: = ".to_string() + &target_balance.value().to_string()));

    if amount > source_balance.value() {
        return false;
    }
    target_balance.set_value(target_balance.value() + amount);
    source_balance.set_value(source_balance.value() - amount);
    true
}

fn add_delegation(lst: &mut Vec<Delegation>, delegate: ScAgent, amount: i64) {
    for d in lst.iter_mut() {
        if d.delegated_to == delegate {
            d.amount = amount;
            return;
        }
    }
    lst.push(Delegation {
        delegated_to: delegate,
        amount: amount,
    })
}

fn sub_delegation(lst: &mut Vec<Delegation>, delegate: ScAgent, amount: i64) -> bool {
    for d in lst {
        if d.delegated_to == delegate {
            return if d.amount >= amount {
                d.amount -= amount;
                true
            } else {
                false
            };
        }
    }
    false
}

fn clean_delegations(lst: Vec<Delegation>) -> Vec<Delegation> {
    let mut ret: Vec<Delegation> = Vec::new();
    for d in lst {
        if d.amount > 0 {
            ret.push(d)
        }
    }
    ret
}

fn encode_delegations(delegations: &Vec<Delegation>) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.int(delegations.len() as i64);
    for d in delegations {
        encoder.agent(&d.delegated_to);
        encoder.int(d.amount);
    }
    return encoder.data();
}

fn decode_delegations(bytes: &[u8]) -> Vec<Delegation> {
    if bytes.len() == 0 {
        return Vec::new();
    }
    let mut decoder = BytesDecoder::new(bytes);
    let size = decoder.int();
    let mut ret: Vec<Delegation> = Vec::with_capacity(size as usize);
    for _ in 0..size {
        ret.push(Delegation {
            delegated_to: decoder.agent(),
            amount: decoder.int(),
        })
    }
    ret
}
