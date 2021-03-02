## Incoming Token Transfers

The idea behind the dividend smart contract is that once we have set up the list
of members, consisting of address/factor pairs, and knowing the total sum of the
factors, we can automatically pay out a dividend to each of the members in the
list according to the factors involved. Whatever amount of tokens gets sent to
the "divide" function will be divided over the members in proportion based on
their respective factors. For example, you could set it up that address A has a
factor 50, B has 30, and C has 20, for a total of 100. Then whenever an amount
of tokens get sent to the "divide" function, address A will receive 50/100th, B
will receive 30/100th, and C will receive 20/100th of that amount.

Tokens provided to a smart contract function can be accesses through the
incoming() method of the function context. It contains zero or more balances of
incoming tokens, each balance for a different token color. Note that these
balances are provided to the smart contract function as if they have already
been deposited in the smart contract's account, but if any error occurs that
causes the function to panic these tokens will be returned to where they came
from. It's only after successful completion of the function that they are truly
deposited to the smart contract account, and it could be that the function
already passes them on to some other account.

Here is how the "divide" function starts:

```rust
let amount = ctx.balances().balance(&ScColor::IOTA);
if amount == 0 {
    ctx.panic("nothing to divide");
}
```

The balances() method of the context is called to create an ScBalances proxy
object that encapsulates an ScImmutableMap that contains color/amount key/value
combinations for each of the token colors that the smart contract account holds.
Note that this is similar to the ScBalances proxy returned by the incoming()
method discussed above, but it contains the total amounts in the account, i.e.
what was already there plus the incoming tokens. Next we ask the ScBalances
proxy to retrieve the amount of tokens for the color ScColor::IOTA, which is a
predefined ScColor value in WasmLib. if the amount of iotas is zero, then we
panic out of the function with an error message. Here we could have used the
require() method of the context instead to do this in a single line.

Now that we know we have a non-zero amount of iotas available for division among
the members we can take the next step:

```rust
let state = ctx.state();
let total = state.get_int64("totalFactor").value();
let members = state.get_map("members");
let member_list = state.get_address_array("memberList");
```

This will retrieve some values and create some proxies we will need next. First
we create the state ScMutableMap through the state() method of the function
context. Then we retrieve the current total factor value from the state through
the ScMutableInt proxy named "totalFactor", and then we set up the ScMutableMap
proxy to the "members" address/factor map, and the ScMutableAddressArray proxy
named "memberList" to the array of member addresses.

```rust
let size = member_list.length();
for i in 0..size {
    let address = member_list.get_address(i).value();
    let factor = members.get_int64(&address).value();
    let share = amount * factor / total;
    if share != 0 {
        let transfers = ScTransfers::new(&ScColor::IOTA, share);
        ctx.transfer_to_address(&address, transfers);
    }
}
```

Now we arrive at the heart of the function. We first determine the amount of
elements in the "memberList" array by calling the length() method. Then we loop
through every element in the array to disperse the correct amount of iotas to
every member.

First we retrieve the ScAddress value from the member list array on the host for
the current index. Then we get the associated factor Int64 value from the
"Members" map using the address as key. Then we calculate the fair share of
iotas for this member from the amount. Note that this is a truncated value. We
may ultimately end up with some iotas remaining in the smart contract account,
but they will be picked up as part of the account balance in the next dispersal
round.

Finally, when there are iotas to send to this member, we set up an ScTransfers
proxy to a token color/value map on the host, and pre-populate it with the
calculated amount of iotas. This will automatically send these values to the map
on the host. Next we use the transfer_to_address() method of the function
context to send this transfer to the member's address on the Tangle. This
function will send the address value to the host and invoke the corresponding
sandbox function on the host to perform the actual transfer.

Next: [Limiting Access](Access.md)