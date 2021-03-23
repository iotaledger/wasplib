## Smart Contract State

The smart contract state storage on the host is similar to the parameter storage
in that it uses a key/value map using raw data bytes. The only difference is
that the state map is mutable as opposed to the immutable params map. So we will
use the proxy objects in the exact same way as we did with the params map with
only two differences. The main difference is that we can actually modify the
data stored in the state map. Another difference is that because all data in
state storage was put there by us, we can be sure that the raw data bytes
represent the value types we stored correctly. So as long as we access the data
in the same way as we stored it we will always get valid data back. To help with
this WasmLib will at runtime verify that the data is accessed using the same
type that was used to stored it, and panic if this was not the case. This will
prevent any data type interpretation errors to propagate in our code or state
storage.

Let's look at how the 'member' function of the 'dividend' smart contract goes
about accessing its state:

```rust
// Now that we have sorted out the parameters we will start using the state
// storage on the host. First we create an ScMutableMap proxy that refers to
// the state storage map on the host.
let state: ScMutableMap = ctx.state();

// We will store the address/factor combinations in a key/value sub-map inside
// the state map. We tell the state map proxy to create an ScMutableMap proxy
// to a map named 'members' in the state storage. If there is no 'members' map
// present yet this will automatically create an empty map on the host.
let members: ScMutableMap = state.get_map(VAR_MEMBERS);

// Now we create an ScMutableInt64 proxy for the value stored in the 'members'
// map under the key defined by the 'address' parameter we retrieved earlier.
let current_factor: ScMutableInt64 = members.get_int64( & address);

// Check to see if this key/value combination exists in the 'members' map
if ! current_factor.exists() {
    // If it does not exist yet then we have to add this new address to the
    // 'memberList' array. We tell the state map proxy to create an
    // ScMutableAddressArray proxy to an Address array named 'memberList' in
    // the state storage. Again, if the array was not present yet it will
    // automatically be created.
    let member_list: ScMutableAddressArray = state.get_address_array(VAR_MEMBER_LIST);
    
    // Now we will append the new address to the memberList array.
    // First we determine the current length of the array.
    let length: i32 = member_list.length();
    
    // Next we create an ScMutableAddress proxy to the Address value that lives
    // at that index in the memberList array (no value, since we're appending).
    let new_address: ScMutableAddress = member_list.get_address(length);
    
    // And finally we append the new address to the array by telling the proxy
    // to update the value it refers with the 'address' parameter.
    new_address.set_value( &address);
}
```

Note how we simply define a nested structure of containers within the state map
by using them as if they already existed. The same thing goes for values in the
containers. You can immediately start using them, and they will default to
all-zero values for fixed size value types, and to zero-length values for
variable sized value types. You will see default values in action in the
fragment below.

```rust
// Create an ScMutableInt64 proxy named 'totalFactor' for an Int64 value in
// state storage. Note that we don't care whether this value exists or not,
// because WasmLib will treat it as if it has the default value of zero.
let total_factor: ScMutableInt64 = state.get_int64(VAR_TOTAL_FACTOR);

// Now we calculate the new running total sum of factors by first getting the
// current value of 'totalFactor' from the state storage, then subtracting the
// current value of the factor associated with the 'address' parameter, if any
// exists. Again, if the associated value doesn't exist, WasmLib will assume it
// to be zero. Finally we add the factor retrieved from the parameters,
// resulting in the new totalFactor.
let new_total_factor: i64 = total_factor.value() - current_factor.value() + factor;

// Now we store the new totalFactor in the state storage
total_factor.set_value(new_total_factor);

// And we also store the factor from the parameters under the address from the
// parameters in the state storage that the proxy refers to
current_factor.set_value(factor);

// Log successful completion of the 'member' Func in the host log.
ctx.log("dividend.member ok");
```

This completes the logic for the 'member' function. In the next section we will
look at how to detect incoming token transfers and how to send tokens to Tangle
addresses.

Next: [Token Transfers](Transfers.md)
