## Smart Contract State

Now that we have the parameters to the function sorted out it's time to update
the smart contract state:

```rust
let state = ctx.state();

let members = state.get_map("members");
let current_factor = members.get_int64(&address);
if ! current_factor.exists() {
    // add new address to member list
    let member_list = state.get_address_array("memberList");
    member_list.get_address(member_list.length()).set_value(&address);
}

let total_factor = state.get_int64("totalFactor");
total_factor.set_value(total_factor.value() - current_factor.value() + factor);
current_factor.set_value(factor);
```

The first thing we do is access the ScMutableMap proxy object to the state
key/value store through the state() method of the function context. Note that
initially we start out with a totally empty state, and that with every update to
the state we keep track of the type of value we have set. The host will prevent
any attempt to access the data stored in the state storage as a different value
type.

Now that we can access the state we start out by creating an ScMutableMap proxy
to the map named "members" that will store the address/factor key/value pairs.
If this map does not exist on the host yet, it will be automatically created as
an empty map. We then use the address we got from the parameters as key to
create a proxy to the current value of the factor for that address. We check if
there already is a value stored there and if not then we need to add the address
to the list of members. We do that by creating an ScAddressArray proxy to an
array named "memberList" that contains all the addresses we have added as
members thus far. We append to the array by getting the ScAddress proxy for the
value stored at the index equal to the current length of the array, i.e. one
item beyond the current last element, and set the Address value stored there to
the address we got from the parameters.

The next step is to calculate the running total of all factors and store that in
the state. We create an ScInt64 proxy for the Int64 value named
"TotalFactor" in the state map, and then calculate the new total by retrieving
the current value of the total through the proxy, then subtracting the value
currently stored under the address, which in case the value was not present yet
will be the default value zero, and then add the factor we extracted from the
parameters. And finally we store the factor from the parameters as the new
current factor for that address.

This completes the logic for the "member" function.

Next: [Incoming Token Transfers](Incoming.md)
