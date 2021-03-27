## Token Transfers

There are two methods in the function context that deal with token balances. The
first one is the `balances()` method, which can be used to determine the current
total balance per token color that is governed by the smart contract. The second
one is the `incoming()` method, which can be used to determine the amounts of
incoming tokens per token color that were sent with the request to call the
smart contract function.

Both methods provide access to zero or more balances of tokens, each for a
different token color, through a special `ScBalances` map proxy. Note that the
incoming balances are provided to the smart contract function as if they have
already been deposited in the smart contract's account, but if any error occurs
which causes the function to panic these incoming tokens will be returned to
where they came from, and it will be as if they were never sent to the smart
contract.

There's also a `transfer_to_address()` method in the function context that can
transfer tokens from the smart contract account to any Tangle address. The
tokens to be transferred are provided to the method through a special
`ScTransfers` map proxy. We will be using the transfer_to_address() method in
the dividend example to disperse the incoming tokens to the member addresses.

The idea behind the dividend smart contract is that once we have set up the list
of members, consisting of address/factor pairs, and knowing the total sum of the
factors, we can automatically pay out a dividend to each of the members in the
list according to the factors involved. Whatever amount of tokens gets sent to
the 'divide' function will be divided over the members in proportion based on
their respective factors. For example, you could set it up that address A has a
factor 50, B has 30, and C has 20, for a total of 100. Then whenever an amount
of tokens gets sent to the 'divide' function, address A will receive 50/100th,
address B will receive 30/100th, and address C will receive 20/100th of that
amount.

Here is how the 'divide' function starts:

```rust
// 'divide' is a function that will take any iotas it receives and properly
// disperse them to the addresses in the member list according to the dispersion
// factors associated with these addresses.
// Anyone can send iota tokens to this function and they will automatically be
// passed on to the member list. Note that this function does not deal with
// fractions. It simply truncates the calculated amount to the nearest lower
// integer and keeps any remaining iotas in its own account. They will be added
// to any next round of tokens received prior to calculation of the new
// dispersion amounts.
pub fn func_divide(ctx: &ScFuncContext) {

    // Log initiation of the 'divide' Func in the host log.
    ctx.log("dividend.divide");

    // Create an ScBalances map proxy to the total account balances for this
    // smart contract. Note that ScBalances wraps an ScImmutableMap of token
    // color/amount combinations in a simpler to use interface.
    let balances: ScBalances = ctx.balances();

    // Retrieve the amount of plain iota tokens from the account balance
    let amount: i64 = balances.balance(&ScColor::IOTA);

    // Create an ScMutableMap proxy to the state storage map on the host.
    let state: ScMutableMap = ctx.state();

    // retrieve the pre-calculated totalFactor value from the state storage
    // through an ScmutableInt64 proxy
    let total_factor: i64 = state.get_int64(VAR_TOTAL_FACTOR).value();

    // note that it is useless to try to divide less than totalFactor iotas
    // because every member would receive zero iotas
    if amount < total_factor {
        // log the fact that we have nothing to do in the host log
        ctx.log("dividend.divide: nothing to divide");

        // And exit the function. Note that we could not have used a require()
        // statement here, because that would have indicated an error and caused
        // a panic out of the function, returning any amount of tokens that was
        // intended to be dispersed to the members. Returning normally will keep
        // these tokens in our account ready for dispersal in a next round.
        return;
    }
```

Now that we know we have determined that we have a non-zero amount of iota
tokens available to send to the members we can start transferring them:

```rust
// Create an ScMutableMap proxy to the 'members' map in the state storage.
let members: ScMutableMap = state.get_map(VAR_MEMBERS);

// Create an ScMutableAddressArray proxy to the 'memberList' Address array
// in the state storage.
let member_list: ScMutableAddressArray = state.get_address_array(VAR_MEMBER_LIST);

// Determine the current length of the memberList array.
let size: i32 = member_list.length();

// loop through all indexes of the memberList array
for i in 0..size {
    // Retrieve the next address from the memberList array through an
    // ScMutableAddress proxy that references the value at the required index.
    let address: ScAddress = member_list.get_address(i).value();

    // Retrieve the factor associated with the address from the members map
    // through an ScMutableInt64 proxy referencing the value in the map.
    let factor: i64 = members.get_int64(&address).value();

    // calculate the fair share of iotas to disperse to this member based on the
    // factor we just retrieved. Note that the result will been truncated.
    let share: i64 = amount * factor / total_factor;

    // is there anything to disperse to this member?
    if share > 0 {
        // Yes, so let's set up an ScTransfers map proxy that transfers the
        // calculated amount of iotas. Note that ScTransfers wraps an
        // ScMutableMap of token color/amount combinations in a simpler to use
        // interface. The constructor we use here creates and initializes a
        // single token color transfer in a single statement. The actual color
        // and amount values passed in will be stored in a new map on the host.
        let transfers: ScTransfers = ScTransfers::new(&ScColor::IOTA, share);

        // Perform the actual transfer of tokens from the smart contract to the
        // member address. The transfer_to_address() method receives the address
        // value and the proxy to the new transfers map on the host, and will
        // call the corresponding host sandbox function with these values.
        ctx.transfer_to_address(&address, transfers);
    }
}

// Log successful completion of the 'divide' Func in the host log.
ctx.log("dividend.divide ok");
```

This completes the logic for the 'divide' function. In the next section we will
look at View-only functions.

Next: [View-Only Functions](Views.md)
